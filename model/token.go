package model

type Token struct {
	Username   string `db:"username" json:"username"`
	Token      string `db:"token" json:"token"`
	Expiration int64  `db:"expiration" json:"expiration"`
}
