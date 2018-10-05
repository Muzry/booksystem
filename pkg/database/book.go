package database

import "time"

type Book struct {
	ID          string    `xorm:"varchar(255) pk not null 'id'" json:"id"`
	BarCode     string    `xorm:"varchar(255) not null 'barcode'" json:"barCode"`
	Number      string    `xorm:"varchar(255) not null 'number'" json:"number"`
	ClassNumber string    `xorm:"varchar(255) not null 'classnumber'" json:"classNumber"`
	Class       string    `xorm:"varchar(255) not null 'class'" json:"class"`
	Name        string    `xorm:"varchar(255) not null 'name'" json:"name"`
	Author      string    `xorm:"varchar(255) not null 'author'" json:"author"`
	Publisher   string    `xorm:"varchar(255) not null 'publisher'" json:"publisher"`
	PageNumber  int64     `xorm:"int not null 'pagenumber'" json:"pageNumber"`
	Size        int64     `xorm:"varchar(255) not null 'size'" json:"size"`
	Price       float64   `xorm:"float not null 'price'" json:"price"`
	Location    string    `xorm:"varchar(255) not null 'location'" json:"location"`
	CreatedAt   time.Time `xorm:"created" json:"createAt"`
	Annotation  string    `xorm:"varchar(255) 'annotation'" json:"annotation"`
}
