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
		lease clientv3.Lease
		leaseGrandResp *clientv3.LeaseGrantResponse
		leaseId clientv3.LeaseID
		keepRespChan <-chan *clientv3.LeaseKeepAliveResponse
		keepResp *clientv3.LeaseKeepAliveResponse
		ctx context.Context
		cancelFunc context.CancelFunc
		kv clientv3.KV
		txn clientv3.Txn
		txnResp *clientv3.TxnResponse
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


	// lease 自动过期
	// op操作
	// txn事物：if else then


	// 1，上锁（创建锁，自动续租，拿着租约去抢占一个key）
	lease = clientv3.NewLease(client)

	// 申请一个5秒的租约
	if leaseGrandResp , err = lease.Grant(context.TODO(), 5); err!=nil{
		fmt.Println(err)
		return
	}
	// 拿到租约的id
	leaseId = leaseGrandResp.ID

	// 准备一个取消自动续租的context
	ctx, cancelFunc = context.WithCancel(context.TODO())

	// 确保函数退出，续租自动退出
	defer cancelFunc()
	defer lease.Revoke(context.TODO(), leaseId)

	// 5秒后自动取消续租
	if keepRespChan, err = lease.KeepAlive(ctx,leaseId); err != nil{
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


	// 如果不纯在key，then设置它，else抢锁失败
	kv = clientv3.NewKV(client)

	// 创建事物
	txn = kv.Txn(context.TODO())

	// 定义事物
	// 如果key不存，就put创建并且续租
	txn.If(clientv3.Compare(clientv3.CreateRevision("/cron/lock/job9"), "=", 0)).
		Then(clientv3.OpPut("/cron/lock/job9", "999999999", clientv3.WithLease(leaseId))).
		Else(clientv3.OpGet("/cron/lock/job9")) // 否则抢锁失败

	if txnResp, err = txn.Commit(); err != nil {
		fmt.Println(err)
		return
	}

	// 判断是否抢到了锁
	if !txnResp.Succeeded { // 如果没有抢到
		fmt.Println("锁被抢占：", txnResp.Responses[0])
		fmt.Println("锁被抢占：", txnResp.Responses[0].GetResponseRange())
		fmt.Println("锁被抢占：", string(txnResp.Responses[0].GetResponseRange().Kvs[0].Value))
		return
	}



	// 2，处理事物

	fmt.Println("处理事物")
	time.Sleep(50 * time.Second)

	// 3，释放锁（取消自动租约，释放租约）



}
