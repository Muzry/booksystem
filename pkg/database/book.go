package database

import "time"

type Book struct {
	ID          string    `xorm:"varchar(255) pk notnull 'id'" json:"id"`
	BarCode     string    `xorm:"varchar(255) notnull 'barCode'" json:"barCode"`
	Number      string    `xorm:"varchar(255) notnull 'number'" json:"number"`
	ClassNumber string    `xorm:"varchar(255) notnull 'classNumber'" json:"classNumber"`
	Class       string    `xorm:"varchar(255) notnull 'class'" json:"class"`
	Name        string    `xorm:"varchar(255) notnull 'name'" json:"name"`
	Author      string    `xorm:"varchar(255) notnull 'author'" json:"author"`
	Publisher   string    `xorm:"varchar(255) notnull 'publisher'" json:"publisher"`
	PageNumber  string    `xorm:"varchar(255) notnull 'pageNumber'" json:"pageNumber"`
	Size        string    `xorm:"varchar(255) notnull 'size'" json:"size"`
	Version     string    `xorm:"varchar(255) notnull 'version'" json:"version"`
	Edition     string    `xorm:"varchar(255) notnull 'edition'" json:"edition"`
	Price       string    `xorm:"varchar(255) notnull 'price'" json:"price"`
	Location    string    `xorm:"varchar(255) notnull 'location'" json:"location"`
	CreatedAt   time.Time `xorm:"created" json:"createAt"`
	Annotation  string    `xorm:"varchar(255) 'annotation'" json:"annotation"`
}
