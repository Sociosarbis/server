package res

import (
	"time"

	"github.com/bangumi/server/internal/dal/dao"
)

type Revision struct {
	ID        uint32                    `json:"id"`
	Type      uint8                     `json:"type"`
	Summary   string                    `json:"summary"`
	Creator   Creator                   `json:"creator"`
	CreatedAt time.Time                 `json:"created_at"`
	Data      dao.GzipPhpSerializedBlob `json:"data"`
}
