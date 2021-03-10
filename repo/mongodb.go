package repo

import (
	"context"
	"fmt"
	"log"
	"sync"
	"todolist/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

// Repository Interface for db
type Repository interface {
	Insert(ctx context.Context, task *models.Task) error
	GetByID(ctx context.Context, id primitive.ObjectID) (models.Task, error)
	GetAll(ctx context.Context) ([]models.Task, error)
	Delete(ctx context.Context, id primitive.ObjectID) (string, error)
	Update(ctx context.Context, id primitive.ObjectID, t *models.Task) (models.Task, error)
}

// MongoDB struct
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

// NewRepo instantiates new MongoDB
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
	mod.Options.SetBackground(true) // must run in background
	ind, err := m.collection.Indexes().CreateOne(ctx, mod)
	if err != nil {
		fmt.Println("Indexes().CreateOne() ERROR:", err)
	} else {
		fmt.Println("CreateOne() index:", ind)
	}

	return &m
}

// DBConnection Get Connection to DB
func (m *MongoDB) DBConnection(ctx context.Context) (*mongo.Collection, error) { // mongo options, connection pools
	// Open server connection
	log.Println("Attempting to connect to ", m.addrString)
	// m.mongoOnce.Do(func() {
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
	// })
	return m.collection, m.clientInstanceError
}

// Insert in to DB
func (m *MongoDB) Insert(ctx context.Context, t *models.Task) error {
	res, err := m.collection.InsertOne(ctx, t)
	if err != nil {
		return err
	}
	newID := res.InsertedID
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		t.OID = oid
	}
	log.Println("Inserted document with ID", newID)
	return nil
}

// GetByID returns tasks and error if not found
func (m *MongoDB) GetByID(ctx context.Context, id primitive.ObjectID) (models.Task, error) {
	var res models.Task
	err := m.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&res)
	log.Println("Found document:", res)
	return res, err
}

// GetAll returns all tasks in collection
func (m *MongoDB) GetAll(ctx context.Context) ([]models.Task, error) {
	var res []models.Task
	// empty case, no documents
	cursor, err := m.collection.Find(ctx, bson.M{})
	if err != nil {
		return res, err
	}
	err = cursor.All(ctx, &res)
	return res, nil
}

// Delete object with OID, returns id of deleted object and error if not found
func (m *MongoDB) Delete(ctx context.Context, id primitive.ObjectID) (string, error) {
	// Handle document not found
	_, err := m.collection.DeleteOne(ctx, bson.M{"_id": id})
	log.Println("deleted document with id", id.String())
	return id.String(), err
}

// Update object at OID, returns document found and error if not found
func (m *MongoDB) Update(ctx context.Context, id primitive.ObjectID, t *models.Task) (models.Task, error) {
	// timestamp updateAt, createAt
	update := bson.M{
		"$set": bson.M{
			"name":        t.Name,
			"description": t.Description,
			"status":      t.Status,
		},
	}
	var res models.Task
	err := m.collection.FindOneAndUpdate(ctx, bson.M{"_id": id}, update).Decode(&res)
	return res, err
}
