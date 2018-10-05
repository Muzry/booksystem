package database

import "time"

type Publisher struct {
	ISBN      string    `xorm:"varchar(255) pk not null 'isbn'" json:"isbn"`
	Name      string    `xorm:"varchar(255) 'name'" json:"name"`
	CreatedAt time.Time `xorm:"created" json:"createAt"`
}
