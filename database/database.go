package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/mgo.v2/bson"
)

type Mongodata interface {
	GetById(id string) (*mongo.Cursor, error)
	GetTotal(params *GetCustomersRequest) (*mongo.Cursor, error)
	GetCantTotal() (int, error)
	Search(params interface{}) (*mongo.Cursor, error)
	Insert(params interface{}) (string, error)
	Update(id string, param interface{}) (int, error)
	Delete(CustomerId string) (int, error)
}

// MongoCustomerRepo stores database connection information
type mongodata struct {
	db *mongo.Database
	//client *mongo.Client
}

//------------------Connect method------------------//
func ConnectDB() (*mongo.Database, *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() //will run when  we're finished main

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(defaultDatabaseURI))
	if err != nil {
		log.Fatal(err)
	}
	client.Connect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database(defaultDbName)

	return db, client
}

//------------------Disconnect method------------------//
func DisconnectDB(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := client.Disconnect(ctx)

	return err
}

func NewDataBase(dbconection *mongo.Database) Mongodata {
	return &mongodata{db: dbconection}
}

//------------------Get customer Id------------------//*Cursor
func (repo *mongodata) GetById(id string) (*mongo.Cursor, error) {
	ObjID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	collection := repo.db.Collection(defaultTable)
	row, err := collection.Find(context.TODO(), bson.M{"_id": ObjID})
	if err != nil {
		return nil, err
	}
	return row, nil
}

//------------------Get total of customer------------------//
func (repo *mongodata) GetTotal(params *GetCustomersRequest) (*mongo.Cursor, error) {
	collection := repo.db.Collection(defaultTable)
	options := options.Find()
	options.SetSkip(int64(params.Offset))
	options.SetLimit(int64(params.Limit))
	row, err := collection.Find(context.TODO(), bson.M{}, options)
	if err != nil {
		return nil, err
	}

	return row, err
}

//------------------Get total of customer------------------//
func (repo *mongodata) GetCantTotal() (int, error) {
	collection := repo.db.Collection(defaultTable)
	total, err := collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return 0, err
	}
	return int(total), nil
}

//------------------Search------------------//
func (repo *mongodata) Search(params interface{}) (*mongo.Cursor, error) {
	collection := repo.db.Collection(defaultTable)
	row, err := collection.Find(context.TODO(), params)
	if err != nil {
		return nil, err
	}
	return row, err
}

//------------------Insert customer------------------//
func (repo *mongodata) Insert(customer interface{}) (string, error) {
	collection := repo.db.Collection(defaultTable)
	// Insert One Document.
	newCustomer, err := collection.InsertOne(context.TODO(), customer)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", newCustomer.InsertedID), nil //return customer.ID.String()

}

//------------------Update customer------------------//
func (repo *mongodata) Update(id string, param interface{}) (int, error) {
	collection := repo.db.Collection(defaultTable)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}
	resultUpdate, err := collection.UpdateOne(context.TODO(),
		bson.M{"_id": objID},
		bson.M{
			"$set": param,
		},
	)
	if err != nil {
		return 0, err
	}
	return int(resultUpdate.ModifiedCount), nil // output: 1

}

//------------------Delete customer------------------//
func (repo *mongodata) Delete(CustomerId string) (int, error) {
	collection := repo.db.Collection(defaultTable)

	objID, err := primitive.ObjectIDFromHex(CustomerId)
	if err != nil {
		return 0, err
	}
	resultDelete, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		return 0, err
	}

	return int(resultDelete.DeletedCount), nil
}
