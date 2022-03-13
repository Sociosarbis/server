package model

import (
	"time"
)

type Creator struct {
	Username string
	Nickname string
}

type Revision struct {
	ID        uint32
	Type      uint8
	Summary   string
	Creator   Creator
	CreatedAt time.Time
	Data      []interface{}
}
