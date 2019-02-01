package main

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"time"
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

func main(){
	var(
		client *mongo.Client
		err error
		databases *mongo.Database
		collection *mongo.Collection
		where *FindByJobName
		cursor mongo.Cursor
		record *LogRecordS
		op *options.FindOptions
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


	where = &FindByJobName{
		JobName:"job10",
	}

	//op = new(options.FindOptions)  // op = &options.FindOptions{}
	op = &options.FindOptions{}

	if cursor, err = collection.Find(context.TODO(),where,op.SetSkip(1),op.SetLimit(5));err != nil {
		fmt.Println(err)
		return
	}
	// 关闭资源
	defer cursor.Close(context.TODO())

	// 遍历结果
	for cursor.Next(context.TODO()) {
		record = &LogRecordS{}

		// 反序列化 并 把数据放到record中
		if err = cursor.Decode(record); err != nil{
			fmt.Println(err)
			return
		}

		fmt.Println((*record))
		//fmt.Println((*record).JobName)
		//fmt.Println((*record).TimePoint.EndTime)

	}

}

// 日志
type LogRecordS struct {
	JobName string `bson:"job_name"` //任务名
	Command string `bson:"command"`//shell命令
	Err string `bson:"err"`// 错误脚本
	Content string `bson:"content"`// 脚本输出
	TimePoint TimePointS `bson:"timePoint"`
}

//时间
type TimePointS struct {
	StartTime int64 `bson:"startTime"`
	EndTime int64 `bson:"endTime"`
}

type FindByJobName struct {
	JobName string `bson:"job_name"`
}