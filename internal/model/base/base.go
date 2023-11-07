package base

import (
	"fmt"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/snowflake"
	"time"
)

type Format interface {
	IDString() string          // IDString returns the ID as a string
	CreatedAtUnixMilli() int64 // CreatedAtUnixMilli returns the created time in milliseconds
	UpdatedAtUnixMilli() int64 //
}

var _ Format = (*Model)(nil)

type Model struct {
	ID        int64     `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"column:created_at;autoUpdateTime;comment:创建时间"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime;comment:更新时间"`
}

func (c Model) IDString() string {
	return fmt.Sprintf("%d", c.ID)
}

func (c Model) CreatedAtUnixMilli() int64 {
	return c.CreatedAt.UnixMilli()
}

func (c Model) UpdatedAtUnixMilli() int64 {
	return c.UpdatedAt.UnixMilli()
}

func NewModel() Model {
	return Model{
		ID:        snowflake.CenterIDNode.Generate().Int64(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
