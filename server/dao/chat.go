package dao

import (
	"MyTodo/engine/v1/db"
	"MyTodo/model/bo/v1"
	"MyTodo/model/po/v1"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Chat struct {
	MongoBaseDao[po.Chat]
}

func (c *Chat) Delete(id string) error {
	oid, err := db.Mongo().BindID(id)
	if err != nil {
		return err
	}
	_, err = c.Collection().UpdateOne(context.Background(), oid, bson.M{
		"$set": bson.M{"deleted_at": time.Now()},
	})
	return err
}

func (c *Chat) Create(chat *po.Chat) error {
	_, err := c.Collection().InsertOne(context.Background(), chat)
	return err
}

func (c *Chat) Find(from, to int, page *Pagination[po.Chat]) error {
	filter := bson.M{"$or": []bson.M{
		{"$and": []bson.M{
			{"from": from},
			{"to": to},
			{"deleted_at": bson.M{"$eq": time.Time{}}}}},
		{"$and": []bson.M{
			{"from": to},
			{"to": from},
			{"deleted_at": bson.M{"$eq": time.Time{}}}}},
	}}
	sortOptions := options.Find().SetSort(bson.D{{Key: "_id", Value: -1}})
	col := c.Collection()
	cur, err := col.Find(context.Background(), filter, sortOptions, page.MongoFindOption())
	if err != nil {
		return err
	}
	defer cur.Close(context.Background())

	var messages []po.Chat
	for cur.Next(context.Background()) {
		var message po.Chat
		err := cur.Decode(&message)
		if err != nil {
			return err
		}
		var reply po.Chat
		oid, err := db.Mongo().BindID(message.Reply)
		if err == nil {
			err = col.FindOne(context.TODO(), oid).Decode(&reply)
			if err == nil {
				fmt.Println(reply)
			}
		}
		messages = append(messages, message)
	}

	page.Data = messages
	if err := cur.Err(); err != nil {
		return err
	}
	return nil
}

func (c *Chat) Snapshot(uid uint) (map[uint]*bo.Snapshot, error) {
	filter := bson.M{"$or": []bson.M{
		{"from": uid},
		{"to": uid},
	}}
	col := c.Collection()
	cur, err := col.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	data := map[uint]*bo.Snapshot{}
	for cur.Next(context.Background()) {
		var message po.Chat
		err := cur.Decode(&message)
		if err != nil {
			return nil, err
		}
		if message.From == uid {
			if data[message.To] == nil {
				data[message.To] = &bo.Snapshot{}
			}
			data[message.To].Count++
			if data[message.To].LastAt.Before(message.CreatedAt) {
				data[message.To].LastAt = message.CreatedAt
				data[message.To].LastMsg = message.Content
			}
		} else if message.To == uid {
			if data[message.From] == nil {
				data[message.From] = &bo.Snapshot{}
			}
			data[message.From].Count++
			if data[message.From].LastAt.Before(message.CreatedAt) {
				data[message.From].LastAt = message.CreatedAt
				data[message.From].LastMsg = message.Content
			}
		}
	}
	return data, nil
}
