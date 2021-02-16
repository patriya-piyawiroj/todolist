package repo

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"sync"
	"todolist/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type MongoDB struct {
	addrString       string
	dbString         string
	collectionString string
	// Pointer to Collection
	collection *mongo.Collection
	// Used to create singleton object of client exposed through GetMongoClient()
	// clientInstance *mongo.Client
	//Used during creation of singleton client object in GetMongoClient().
	clientInstanceError error
	//Used to execute client creation procedure only once.
	mongoOnce sync.Once
}

func NewRepo(ctx context.Context, addr string, db string, collection string) *MongoDB {
	m := MongoDB{
		addrString:       addr,
		dbString:         db,
		collectionString: collection,
	}
	m.DBConnection(ctx)

	// Sample index creation
	mod := mongo.IndexModel{
		Keys: bson.M{
			"CreatedAt": 1, // index in ascending order
		}, Options: nil,
	}
	ind, err := m.collection.Indexes().CreateOne(ctx, mod)
	if err != nil {
		fmt.Println("Indexes().CreateOne() ERROR:", err)
	} else {
		fmt.Println("CreateOne() index:", ind)
		fmt.Println("CreateOne() type:", reflect.TypeOf(ind), "\n")
	}

	return &m
}

// DBConnection Get Connection to DB
func (m *MongoDB) DBConnection(ctx context.Context) (*mongo.Collection, error) {
	// Open server connection
	log.Println("Attempting conn")
	m.mongoOnce.Do(func() {
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.addrString))
		if err != nil {
			log.Fatal(err)
			m.clientInstanceError = err
		}
		// Check the connection
		err = client.Ping(ctx, nil)
		if err != nil {
			m.clientInstanceError = err
		}
		m.collection = client.Database(m.dbString).Collection(m.collectionString)
		log.Println("Connected to mongo client")
	})
	return m.collection, m.clientInstanceError
}

// Insert in to DB
func (m *MongoDB) Insert(ctx context.Context, t *models.Task) error {
	log.Println("Attempting insert")
	var err error
	m.collection, err = m.DBConnection(ctx)
	if err != nil {
		return err
	}
	res, err := m.collection.InsertOne(ctx, t)
	if err != nil {
		return err
	}
	newID := res.InsertedID
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		t.OID = oid
	}

	log.Println("inserted document with ID", newID)
	return nil
}
