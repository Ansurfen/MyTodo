package mongo

import (
	"MyTodo/conf"
	interfaces "MyTodo/interface"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type TodoMongo struct {
	*mongo.Client
}

func New(opt conf.MongoOption) *TodoMongo {
	cli, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(fmt.Sprintf("mongodb://%s:%d",
			opt.Host, opt.Port)).SetAuth(options.Credential{
		Username: opt.Username,
		Password: opt.Password,
	}))
	if err != nil {
		panic(err)
	}
	return &TodoMongo{cli}
}

func (c *TodoMongo) Ping() error {
	return c.Client.Ping(context.TODO(), readpref.Primary())
}

func (c *TodoMongo) BindID(id string) (bson.M, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return bson.M{"_id": oid}, nil
}

func (c *TodoMongo) Database(database string) *mongo.Database {
	return c.Client.Database(database)
}

func (c *TodoMongo) Collection(database string, tbl interfaces.Table) *mongo.Collection {
	return c.Database(database).Collection(tbl.TableName())
}
