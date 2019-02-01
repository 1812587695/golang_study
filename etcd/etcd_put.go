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
		putResp *clientv3.PutResponse
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

	if putResp, err = kv.Put(context.TODO(), "/cron/jobs/job2", "job2", clientv3.WithPrevKV()); err!=nil{
		fmt.Println(err)
	} else {
		// 打印当前信息
		fmt.Println(putResp.Header)
		// 打印put以前key-value信息
		if putResp.PrevKv != nil {
			fmt.Println(putResp.PrevKv)
		}
	}

}
