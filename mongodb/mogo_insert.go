package main

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"time"
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

func main(){
	var(
		client *mongo.Client
		err error
		databases *mongo.Database
		collection *mongo.Collection
		record *LogRecord
		result *mongo.InsertOneResult
		docId primitive.ObjectID
	)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if client, err = mongo.Connect(ctx, "mongodb://192.168.0.180:27017");err != nil{
		fmt.Println(err)
		return
	}

	// 选择数据库
	databases = client.Database("cron")

	// 选择表
	collection = databases.Collection("log")

	record = &LogRecord{
		JobName:"job10",
		Command:"echo hello1",
		Err:"error",
		Content: "hello1",
		TimePoint: TimePoint{
			StartTime:time.Now().Unix(),
			EndTime:time.Now().Unix() +10,
		},
	}
	// 插入表
	if result,err = collection.InsertOne(context.TODO(), record); err != nil{
		fmt.Println(err)
		return
	}

	docId = result.InsertedID.(primitive.ObjectID)
	fmt.Println(docId.Hex())
}

// 日志
type LogRecord struct {
	JobName string `bson:"job_name"` //任务名
	Command string `bson:"command"`//shell命令
	Err string `bson:"err"`// 错误脚本
	Content string `bson:"content"`// 脚本输出
	TimePoint TimePoint `bson:"timePoint"`
}

//时间
type TimePoint struct {
	StartTime int64 `bson:"startTime"`
	EndTime int64 `bson:"endTime"`
}