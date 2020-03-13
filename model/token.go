package model

type Token struct {
	Username   string `db:"username"`
	Token      string `db:"token"`
	Expiration int64  `db:"expiration"`
}
