package repository

import (
	"fmt"
	"github.com/Munovv/go-url-crop/pkg/model"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"math/rand"
	"time"
)

type LinkMysql struct {
	db *sqlx.DB
}

func NewLinkMysql(db *sqlx.DB) *LinkMysql {
	return &LinkMysql{db: db}
}

func (r *LinkMysql) CropLink(link string) (string, error) {
	code := r.GenerateCode()
	domain := viper.GetString("server.domain") + ":" + viper.GetString("server.port")

	newLink := model.Link{
		Link:      link,
		CropLink:  domain + "/go/" + code,
		CreatedAt: time.Now().Unix(),
	}

	query := fmt.Sprintf("INSERT INTO %s (link, crop_link, created_at) VALUES (:1, :2, :3)",
		linkTable)
	_, err := r.db.NamedExec(query, map[string]interface{}{
		"1": newLink.Link,
		"2": newLink.CropLink,
		"3": newLink.CreatedAt,
	})
	if err != nil {
		logrus.Fatalf("error related to sending a SQL request: %s", err.Error())
	}

	return newLink.CropLink, nil
}

func (r *LinkMysql) GetLink(link string) (model.Link, error) {
	return model.Link{}, nil
}

func (r *LinkMysql) GenerateCode() string {
	rand.Seed(time.Now().UnixNano())

	lenght := 8
	digits := "0123456789"
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyz" + digits

	buffer := make([]byte, lenght)
	buffer[0] = digits[rand.Intn(len(digits))]
	for i := 1; i < lenght; i++ {
		buffer[i] = chars[rand.Intn(len(chars))]
	}
	rand.Shuffle(len(buffer), func(i, j int) {
		buffer[i], buffer[j] = buffer[j], buffer[i]
	})
	str := string(buffer)

	return str
}

func (r *LinkMysql) RedirectLink(code string) (string, error) {
	var link model.Link

	domain := viper.GetString("server.domain") + ":" +
		viper.GetString("server.port") + "/go/"
	cropLink := domain + code

	query := fmt.Sprintf("SELECT * FROM %s WHERE crop_link = ?", linkTable)
	if err := r.db.Get(&link, query, cropLink); err != nil {
		logrus.Fatalf("error related to sending a SQL request: %s", err.Error())
		return link.Link, err
	}

	return link.Link, nil
}
