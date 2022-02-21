package entity

import "time"

type Token struct {
	Hash          []byte    `json:"-"`
	PlainPassword string    `json:"token"`
	UserID        int       `json:"user_id"`
	Expiry        time.Time `json:"expiry"`
	Scope         string    `json:"scope"`
}

var EmptyToken = Token{}
