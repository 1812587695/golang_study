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
		record *LogRecordM
		result *mongo.InsertManyResult
		docId primitive.ObjectID
		logArr []interface{}
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

	record = &LogRecordM{
		JobName:"job10",
		Command:"echo hello22",
		Err:"error",
		Content: "hello1",
		TimePoint: TimePointM{
			StartTime:time.Now().Unix(),
			EndTime:time.Now().Unix() +10,
		},
	}

	logArr = []interface{}{record,record,record}
	// 插入表
	if result,err = collection.InsertMany(context.TODO(), logArr); err != nil{
		fmt.Println(err)
		return
	}

	docId = result.InsertedIDs[0].(primitive.ObjectID)
	fmt.Println(docId.Hex())
}

// 日志
type LogRecordM struct {
	JobName string `bson:"job_name"` //任务名
	Command string `bson:"command"`//shell命令
	Err string `bson:"err"`// 错误脚本
	Content string `bson:"content"`// 脚本输出
	TimePoint TimePointM `bson:"timePoint"`
}

//时间
type TimePointM struct {
	StartTime int64 `bson:"startTime"`
	EndTime int64 `bson:"endTime"`
}