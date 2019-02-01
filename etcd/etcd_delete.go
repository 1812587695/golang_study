package main

import (
	"time"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"context"
	"go.etcd.io/etcd/mvcc/mvccpb"
)

func main()  {
	var(
		config clientv3.Config
		client *clientv3.Client
		kv clientv3.KV
		delResp *clientv3.DeleteResponse
		idx int
		kvpair *mvccpb.KeyValue
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

	if delResp, err = kv.Delete(context.TODO(), "/cron/jobs/job2",clientv3.WithPrevKV()); err!=nil{
		fmt.Println(err)
	} else {
		// 打印被删除之前的值
		if len(delResp.PrevKvs) != 0{
			for idx,kvpair = range delResp.PrevKvs {
				fmt.Println("删除了：" ,idx, string(kvpair.Key), string(kvpair.Value))
			}
		}
	}

}
