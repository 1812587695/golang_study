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
		getResp *clientv3.GetResponse
		watcheStartRevision int64
		watcher clientv3.Watcher
		watchRespChan <-chan clientv3.WatchResponse
		watchResp clientv3.WatchResponse
		event *clientv3.Event
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



	// 模拟etcd中kv的变化
	go func() {
		for {
			kv.Put(context.TODO(), "/cron/jobs/job3", "job4")

			kv.Delete(context.TODO(), "/cron/jobs/job3")

			time.Sleep(1 * time.Second)
		}
	}()


	// 先get到kv值，并监听变化
	if getResp, err = kv.Get(context.TODO(), "/cron/jobs/job3"); err!=nil{
		fmt.Println(err)
	}

	if len(getResp.Kvs) != 0 {
		fmt.Println(string(getResp.Kvs[0].Value))
	}


	// 当前etcd集群事物ID，递增的
	watcheStartRevision = getResp.Header.Revision + 1

	// 创建一个watch
	watcher = clientv3.NewWatcher(client)

	// 启动监听
	fmt.Println("从该版本向后监听：", watcheStartRevision)

	watchRespChan = watcher.Watch(context.TODO(), "/cron/jobs/job3", clientv3.WithRev(watcheStartRevision))

	for watchResp = range watchRespChan {
		for _, event = range watchResp.Events {
			switch event.Type {
			case mvccpb.PUT:
				fmt.Println("修改为：", string(event.Kv.Value), "Revision:", event.Kv.CreateRevision, event.Kv.ModRevision)
			case mvccpb.DELETE:
				fmt.Println("删除了：", "Revision:", event.Kv.ModRevision)
			}
		}
	}

}
