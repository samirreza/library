package valueobject

import (
	"database/sql/driver"
	"errors"
	"fmt"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ULID ulid.ULID

func NewULIDFromString(u string) (ULID, error) {
	ulid, err := ulid.Parse(u)
	return ULID(ulid), err
}

func NewFromULID(u ulid.ULID) ULID {
	return ULID(u)
}

func (u ULID) String() string {
	return ulid.ULID(u).String()
}

func (u ULID) Value() (driver.Value, error) {
	return u.String(), nil
}

func (u *ULID) Scan(value interface{}) error {
	var us string
	switch v := value.(type) {
	case []byte:
		us = string(v)
	case string:
		us = v
	default:
		return errors.New(fmt.Sprint("Failed to parse URL:", value))
	}
	uu, err := ulid.Parse(us)
	if err != nil {
		return err
	}
	*u = ULID(uu)
	return nil
}

func (ULID) GormDataType() string {
	return "ulid"
}

func (ULID) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return "VARCHAR(26)"
}
