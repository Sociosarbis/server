package model

import (
	"time"

	"github.com/bangumi/server/internal/dal/dao"
)

type Creator struct {
	Username string
	Nickname string
}

type Revision struct {
	ID        uint32
	Type      uint8
	Summary   string
	CreatorID uint32
	CreatedAt time.Time
	Data      dao.GzipPhpSerializedBlob
}
