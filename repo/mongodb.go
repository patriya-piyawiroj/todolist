package repo

import (
	"context"
	"log"
	"sync"
	"todolist/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectionString = "mongodb://mongodb:27017"
	localString      = "mongodb://localhost:27017"
	dbString         = "todolist"
	tasks            = "tasks"
)

type MongoDB struct {
	// Pointer to Collection
	collection *mongo.Collection
	// Used to create singleton object of client exposed through GetMongoClient()
	// clientInstance *mongo.Client
	//Used during creation of singleton client object in GetMongoClient().
	clientInstanceError error
	//Used to execute client creation procedure only once.
	mongoOnce sync.Once
}

func NewMongoDB(ctx context.Context) *MongoDB {
	m := MongoDB{}
	m.DBConnection(ctx)
	return &m
}

// DBConnection Get Connection to DB
func (m *MongoDB) DBConnection(ctx context.Context) (*mongo.Collection, error) {
	// Open server connection
	log.Println("Attempting conn")
	m.mongoOnce.Do(func() {
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(localString))
		if err != nil {
			log.Fatal(err)
			m.clientInstanceError = err
		}
		// Check the connection
		err = client.Ping(ctx, nil)
		if err != nil {
			m.clientInstanceError = err
		}
		m.collection = client.Database(dbString).Collection(tasks)
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
