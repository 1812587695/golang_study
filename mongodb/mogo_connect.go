package main

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"time"
	"context"
	"fmt"
)

func main(){
	var(
		client *mongo.Client
		err error
		databases *mongo.Database
		collection *mongo.Collection
	)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if client, err = mongo.Connect(ctx, "mongodb://192.168.0.180:27017");err != nil{
		fmt.Println(err)
		return
	}

	// 选择数据库
	databases = client.Database("my_db")

	// 选择表
	collection = databases.Collection("my_collection")

	fmt.Println(collection)
}
