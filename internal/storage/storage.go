package storage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	db *mongo.Database
}

const dbName = "scraper"
const collName = "products"

func ConnectStorage(storagePath string) (*Storage, error) {
	clientOptions := options.Client().ApplyURI(storagePath)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return &Storage{db: client.Database(dbName)}, nil
}

func (s *Storage) CloseStorage() error {
	err := s.db.Client().Disconnect(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) InsertData(data []byte) error {
	coll := s.db.Collection(collName)

	_, err := coll.InsertOne(context.Background(), data)
	if err != nil {
		if dupErr := isDuplicate(err); dupErr != nil {
			return dupErr
		}
		return err
	}
	return nil

}

func isDuplicate(err error) error {
	if mongoErr, ok := err.(mongo.WriteException); ok {
		if mongoErr.WriteErrors[0].Code == 11000 {
			return fmt.Errorf("Duplicate value error")

		}
	}
	return nil
}
