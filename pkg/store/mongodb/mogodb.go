package mongodb

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"myproject/pkg/types"
	"os"
	"strings"

	"github.com/caicloud/nirvana/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// DefaultDBName is default db name to store vmapp data and so on
	DefaultDBName         = "test"
	MessageCollectionName = "test"
)

// MS is global mongo driver instance
var MS *MongoStore

func init() {
	MS = NewMongoStore()
}

// MongoStore is a mongo client
type MongoStore struct {
	Client            *mongo.Client
	MessageCollection *mongo.Collection
}

// NewMongoStore news a mongo client
func NewMongoStore() *MongoStore {
	host := os.Getenv("MONGODB_URL")
	if host != "" && !strings.Contains(host, "mongodb://") {
		host = fmt.Sprintf("mongodb://%s", host)
	}
	if host == "" {
		host = "mongodb://localhost:27917"
	}
	clientOptions := options.Client().ApplyURI(host)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("MongoDB connected: %s", host)

	db := client.Database(DefaultDBName)
	return &MongoStore{
		Client:            client,
		MessageCollection: db.Collection(MessageCollectionName),
	}
}

func (s *MongoStore) Add(ctx context.Context, obj interface{}) (interface{}, error) {
	m, ok := obj.(types.Message)
	if !ok {
		return nil, errors.New("type is not message")
	}
	filter := bson.M{"id": m.ID}
	// 这么判断是否存在真的优雅吗？
	cur, err := s.MessageCollection.Find(ctx, filter)
	if cur.Next(ctx) {
		return nil, errors.New("already exist")
	}
	_, err = s.MessageCollection.InsertOne(ctx, m)

	return m, err
}

func (s *MongoStore) Del(ctx context.Context, id interface{}) error {
	name := id.(int)
	_, err := s.MessageCollection.DeleteOne(ctx, bson.M{"id": name})
	return err
}

func (s *MongoStore) Upd(ctx context.Context, obj interface{}) (interface{}, error) {
	m, ok := obj.(types.Message)
	if !ok {
		return nil, errors.New("type is not message")
	}
	filter := bson.D{{"id", m.ID}}
	r, err := s.MessageCollection.UpdateOne(ctx, filter, bson.D{{"$set", m}})
	if r.MatchedCount < 1 {
		return nil, errors.New("not found")
	}
	return m, err
}

func (s *MongoStore) Get(ctx context.Context, id interface{}) (interface{}, error) {
	name := id.(int)
	re := types.Message{}
	filter := bson.M{"id": name}
	err := s.MessageCollection.FindOne(ctx, filter).Decode(&re)
	return re, err
}

func (s *MongoStore) List(ctx context.Context) (interface{}, error) {
	var ret []types.Message
	cur, err := s.MessageCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var a types.Message
		err := cur.Decode(&a)
		if err != nil {
			return nil, err
		}
		ret = append(ret, a)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return ret, nil
}
