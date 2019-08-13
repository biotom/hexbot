package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type ColourStore struct {
	client *mongo.Client
}



type MongoConnector interface {
	Client() *mongo.Client
}

func NewColourStore(c MongoConnector) (*ColourStore, error) {
	return &ColourStore{
		client: c.Client(),
	}, nil
}
func (cs *ColourStore) SaveColour(ctx context.Context, hexString string) (err error) {

	return nil
}

