package model

import (
	"github.com/samirreza/library/internal/valueobject"
)

type Book struct {
	ID          valueobject.ULID `gorm:"type:ulid;primary_key"`
	Title       string
	Description string
}
