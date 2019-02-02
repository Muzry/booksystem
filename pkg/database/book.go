package database

import "time"

type Book struct {
	ID          string    `xorm:"varchar(255) pk not null 'id'" json:"id"`
	BarCode     string    `xorm:"varchar(255) not null 'barCode'" json:"barCode"`
	Number      string    `xorm:"varchar(255) not null 'number'" json:"number"`
	ClassNumber string    `xorm:"varchar(255) not null 'classNumber'" json:"classNumber"`
	Class       string    `xorm:"varchar(255) not null 'class'" json:"class"`
	Name        string    `xorm:"varchar(255) not null 'name'" json:"name"`
	Author      string    `xorm:"varchar(255) not null 'author'" json:"author"`
	Publisher   string    `xorm:"varchar(255) not null 'publisher'" json:"publisher"`
	PageNumber  string    `xorm:"varchar(255) not null 'pageNumber'" json:"pageNumber"`
	Size        string    `xorm:"varchar(255) not null 'size'" json:"size"`
	Version     string    `xorm:"varchar(255) not null 'version'" json:"version"`
	Edition     string    `xorm:"varchar(255) not null 'edition'" json:"edition"`
	Price       string    `xorm:"varchar(255) not null 'price'" json:"price"`
	Location    string    `xorm:"varchar(255) not null 'location'" json:"location"`
	CreatedAt   time.Time `xorm:"created" json:"createAt"`
	Annotation  string    `xorm:"varchar(255) 'annotation'" json:"annotation"`
}
