package repository

import (
	"github.com/pkg/errors"
	"github.com/udodinho/golangProjects/wallet-engine/internals/ports"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"time"
)

// MongoRepository MongoDB Connection Struct
type MongoRepository struct {
	Client   *mongo.Client
	Database string
	Timeout  time.Duration
}

// MongoDB Connection
func newMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}
	return client, nil

}

func NewMongoRepository(mongoURL, mongoDB string, mongoTimeout int) (ports.WalletRepository, error) {
	repo := &MongoRepository{
		Database: mongoDB,
		Timeout:  time.Duration(mongoTimeout) * time.Second,
	}
	client, err := newMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		return nil, errors.Wrap(err, "client error")
	}
	repo.Client = client
	return repo, nil
}
