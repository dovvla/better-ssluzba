package models

import "time"

type Request struct {
	CreatedAt time.Time `gorm:"autoCreateTime:true"`
}
