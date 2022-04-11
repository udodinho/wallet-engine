package repository

import (
	"fmt"
	"github.com/udodinho/golangProjects/wallet-engine/internals/core/domain/wallet"
	"github.com/udodinho/golangProjects/wallet-engine/internals/core/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"log"
)

func (m *MongoRepository) CreateWallet(wallet *wallet.User) (*wallet.User, error) {
	coll := m.Client.Database("opay").Collection("opay-collection")
	_, err := coll.InsertOne(context.TODO(), wallet)
	if err != nil {
		helpers.LogEvent("INFO", "Creating wallet failed.")
		return nil, helpers.PrintErrorMessage("500", err.Error())
	}

	helpers.LogEvent("INFO", fmt.Sprintf("Wallet %v created successfully", wallet))
	return wallet, nil
}

func (m *MongoRepository) GetUserByEmail(email string) ([]*wallet.User, error) {
	helpers.LogEvent("INFO", fmt.Sprintf("Getting user by email..."))
	coll := m.Client.Database("opay").Collection("opay-collection")
	filter := bson.D{{"email", email}}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var data []*wallet.User
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		log.Fatal(err)
	}
	helpers.LogEvent("INFO", fmt.Sprintf("Found %v number of users", email))
	return data, err
}

func (m *MongoRepository) CheckPassword(userRef string) ([]*wallet.User, error) {
	helpers.LogEvent("INFO", fmt.Sprintf("Checking user password..."))
	coll := m.Client.Database("opay").Collection("opay-collection")
	filter := bson.D{{"reference", userRef}}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var data []*wallet.User
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		log.Fatal(err)
	}
	helpers.LogEvent("INFO", fmt.Sprintf("Found %v number of users", userRef))
	return data, err
}

func (m *MongoRepository) GetBalance(userID string) ([]*wallet.Wallet, error) {
	helpers.LogEvent("INFO", fmt.Sprintf("Getting balance for user..."))
	coll := m.Client.Database("opay").Collection("account-balance")
	filter := bson.D{{"user_id", userID}}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var data []*wallet.Wallet
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		log.Fatal(err)
	}
	helpers.LogEvent("INFO", fmt.Sprintf("Got the balance of %v successfully", userID))
	return data, err
}

func (m *MongoRepository) ChangeStatus(isActive bool, userRef string) (interface{}, error) {
	coll := m.Client.Database("opay").Collection("opay-collection")
	filter := bson.D{{"reference", userRef}}
	opts := options.Update().SetUpsert(true)
	update := bson.D{
		{"$set", bson.D{
			{"is_active", isActive},
		}},
	}
	value, err := coll.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		helpers.LogEvent("INFO", "Changing status failed.")
		return nil, helpers.PrintErrorMessage("500", err.Error())
	}

	if value.MatchedCount != 0 {
		helpers.LogEvent("INFO", "Status changed.")
	}

	if value.UpsertedCount != 0 {
		helpers.LogEvent("INFO", fmt.Sprintf("Inserted new value with %v\n changed successfully.", value.UpsertedID))
	}

	return value.UpsertedID, err
}

func (m *MongoRepository) PostTransaction(transaction *wallet.Wallet) (interface{}, error) {
	coll := m.Client.Database("opay").Collection("account-balance")
	_, err := coll.InsertOne(context.TODO(), transaction)
	if err != nil {
		helpers.LogEvent("INFO", "Posting transaction failed.")
		return nil, helpers.PrintErrorMessage("500", err.Error())
	}

	helpers.LogEvent("INFO", "Transaction posted successfully.")
	return transaction, nil
}

func (m *MongoRepository) SaveTransaction(transaction *wallet.Transaction) (interface{}, error) {
	coll := m.Client.Database("opay").Collection("transaction-service")
	_, err := coll.InsertOne(context.TODO(), transaction)
	if err != nil {
		helpers.LogEvent("INFO", "Saving transaction failed.")
		return nil, helpers.PrintErrorMessage("500", err.Error())
	}

	helpers.LogEvent("INFO", fmt.Sprintf("Transaction saved successfully"))
	return transaction, err
}
