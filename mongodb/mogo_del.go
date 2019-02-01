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
		//delWhere FindByJobNameD
		//delWhere DeleteCond
		delWhere OneTime
		deleteResult *mongo.DeleteResult
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

	// 1种：删除条件{job_name:"job10"}
	//delWhere = FindByJobNameD{
	//	JobName: "job10",
	//}

	// 2种：删除条件小于当前时间的 {"timePoint.startTime":{"$lt":timestamp}}
	//delWhere = DeleteCond{
	//	BeforeCond:TimeBefore{
	//		Before:time.Now().Unix(),
	//	},
	//}

	// 3种：{"timePoint.startTime":timestamp}
	delWhere = OneTime{
		BeforeCond:1548914035,
	}

	if deleteResult, err = collection.DeleteMany(context.TODO(), delWhere); err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println(deleteResult.DeletedCount)


}

type FindByJobNameD struct {
	JobName string `bson:"job_name"`
}

// {"timePoint.startTime":{"$lt":timestamp}}
type DeleteCond struct {
	BeforeCond TimeBefore `bson:"timePoint.startTime"`
}

// {"$lt":timestamp}
type TimeBefore struct {
	Before int64 `bson:"$lt"`
}

// 1548914039

type OneTime struct {
	BeforeCond int64 `bson:"timePoint.startTime"`
}