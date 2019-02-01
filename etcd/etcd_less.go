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
		lease clientv3.Lease
		leaseGrandResp *clientv3.LeaseGrantResponse
		leaseId clientv3.LeaseID
		getResp *clientv3.GetResponse
		keepResp *clientv3.LeaseKeepAliveResponse
		keepRespChan <-chan *clientv3.LeaseKeepAliveResponse
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



	// 申请租约
	lease = clientv3.NewLease(client)

	// 申请一个10秒的租约
	if leaseGrandResp , err = lease.Grant(context.TODO(), 10); err!=nil{
		fmt.Println(err)
		return
	}

	// 拿到租约的id
	leaseId = leaseGrandResp.ID


	// 续租
	if keepRespChan, err = lease.KeepAlive(context.TODO(),leaseId); err != nil{
		fmt.Println(err)
		return
	}

	go func() {
		for {
			select {
				case keepResp = <- keepRespChan:
					if keepRespChan == nil {
						fmt.Println("续租已经失效")
						goto END
					} else {
						fmt.Println("收到自动续租应答", keepResp.ID)
					}
			}
		}
		END:
	}()


	// 用于读写etcd的键值对
	kv = clientv3.NewKV(client)

	// 关联租约
	if putResp, err = kv.Put(context.TODO(), "/cron/jobs/job2", "", clientv3.WithLease(leaseId)); err!=nil{
		fmt.Println(err)
	} else {
		// 打印当前信息
		fmt.Println(putResp.Header)
	}

	// 定时查看key过期没有

	for {
		if getResp, err = kv.Get(context.TODO(),"/cron/jobs/job2"); err != nil {
			fmt.Println(err)
			return
		}
		if getResp.Count == 0 {
			fmt.Println("kv 过期了")
			break
		}
		fmt.Println("没过期", getResp.Kvs)
		time.Sleep(2 * time.Second)
	}

}
