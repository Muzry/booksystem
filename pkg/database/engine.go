package database

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

const DataBaseFile = "database.db"

//var Engine *xorm.Engine

func GetConnection() (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("sqlite3", DataBaseFile)
	if err != nil {
		return nil, err
	}
	engine.ShowSQL(true)
	return engine, nil
}

func InitTable(engine *xorm.Engine) error {

	//Engine = engine
	err := engine.Sync2(Book{}, Publisher{})
	if err != nil {
		return err
	}
	return nil
}
