package publisher

import (
	"booksystem/pkg/database"
	"fmt"
)

const (
	Create = "create"
	Update = "update"
)

func GetPublisherInfo(publisherID string) (*database.Publisher, error) {

	publisher := database.Publisher{}
	db, err := database.GetConnection()
	defer db.Close()
	if err != nil {
		return nil, err
	}

	has, err := db.ID(publisherID).Get(&publisher)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, fmt.Errorf("no such result from publisher table")
	}
	return &publisher, nil
}

func DeletePublisher(publisherID string) error {
	publisher := database.Publisher{}
	db, err := database.GetConnection()
	defer db.Close()

	if err != nil {
		return err
	}

	_, err = db.ID(publisherID).Delete(publisher)
	if err != nil {
		return err
	}
	return nil
}

func GetPublishers(limit, start int) (*[]database.Publisher, error) {
	publishers := make([]database.Publisher, 0)
	db, err := database.GetConnection()
	defer db.Close()

	if err != nil {
		return nil, err
	}

	err = db.Limit(limit, start).Find(&publishers)

	if err != nil {
		return nil, err
	}
	return &publishers, nil
}

func CreateORUpdatePublisher(publisher database.Publisher, mothodType, publisherID string) (*database.Publisher, error) {
	db, err := database.GetConnection()
	defer db.Close()

	if err != nil {
		return nil, err
	}

	if mothodType == Create {
		publisher.ID = publisherID
		_, err = db.Insert(publisher)
		if err != nil {
			return nil, err
		}
	} else if mothodType == Update {
		_, err = db.ID(publisherID).Update(publisher)
		if err != nil {
			return nil, err
		}
	}
	return &publisher, nil
}
