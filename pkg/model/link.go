package model

type Link struct {
	Id        int    `db:"id"`
	Link      string `db:"link"`
	CropLink  string `db:"crop_link"`
	CreatedAt int64  `db:"created_at"`
}
