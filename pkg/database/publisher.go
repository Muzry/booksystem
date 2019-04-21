package database

import "time"

type Publisher struct {
	ID        string    `xorm:"varchar(255) pk notnull 'id'" json:"id"`
	ISBN      string    `xorm:"varchar(255) 'isbn'" json:"isbn"`
	Name      string    `xorm:"varchar(255) 'name'" json:"name"`
	CreatedAt time.Time `xorm:"created" json:"createAt"`
}
