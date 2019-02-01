package main

import (
	"time"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"context"
)

func main()  {
	var(
		config clientv3.Config
		client *clientv3.Client
		kv clientv3.KV
		putOp clientv3.Op
		getOp clientv3.Op
		opResponse clientv3.OpResponse
	)
	// 配置
	config = clientv3.Config{
		Endpoints:   []string{"192.168.0.180:2379", "192.168.0.181:2379", "192.168.0.182:2379"},
		DialTimeout: 5 * time.Second,
	}
	// 链接
	client, err := clientv3.New(config)
	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()

	// 用于读写etcd的键值对
	kv = clientv3.NewKV(client)

	// 写入op
	putOp = clientv3.OpPut("/cron/jobs/job6", "job6job6job6")

	if opResponse, err = kv.Do(context.TODO(), putOp); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(opResponse.Put().Header)


	// 获取

	getOp = clientv3.OpGet("/cron/jobs/job6")

	if opResponse, err = kv.Do(context.TODO(), getOp); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(opResponse.Get().Kvs)
}
