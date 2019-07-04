package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {

	var (
		config       clientv3.Config
		client       *clientv3.Client
		err          error
		kv           clientv3.KV
		lease        clientv3.Lease
		leaseResp    *clientv3.LeaseGrantResponse
		leaseID      clientv3.LeaseID
		putResp      *clientv3.PutResponse
		getResp      *clientv3.GetResponse
		keepResp     *clientv3.LeaseKeepAliveResponse
		keepRespChan <-chan *clientv3.LeaseKeepAliveResponse
	)

	config = clientv3.Config{
		Endpoints:   []string{"0.0.0.0:2379"}, //集群列表
		DialTimeout: 5 * time.Second,
	}

	//建立一个客户端
	if client, err = clientv3.New(config); err != nil {
		fmt.Println("clientv3 config err:", err)
		return
	}

	//用于读写etcd的键值对

	//申请一个lease
	lease = clientv3.NewLease(client)

	//创建一个10秒的租约
	if leaseResp, err = lease.Grant(context.TODO(), 10); err != nil {
		fmt.Println("clientv3 lease grant err ", err)
		return
	}

	//拿到租约的ID
	leaseID = leaseResp.ID

	//自动续租
	//	ctx,_:=context.WithTimeout(context.TODO(),5*time.Second)

	if keepRespChan, err = lease.KeepAlive(context.TODO(), leaseID); err != nil {
		fmt.Println("clientv3 lease leppalive err ", err)
		return
	}

	go func() {
		for {
			select {
			case keepResp = <-keepRespChan:
				if keepRespChan == nil {
					fmt.Println("租约已经过期")
					goto END
				} else {
					fmt.Println("收到自动续租应答：", keepResp.ID)
				}
			}
		}
	END:
	}()

	//获取KV API子集
	kv = clientv3.NewKV(client)

	if putResp, err = kv.Put(context.TODO(), "/cron/job2", "{....}", clientv3.WithLease(leaseID)); err != nil {
		fmt.Println("clientv3 kv put err:", err)
		return
	}

	fmt.Println("写入成功：", putResp.Header.Revision)

	//
	for {
		if getResp, err = kv.Get(context.TODO(), "/cron/job2"); err != nil {
			fmt.Println(" clientv3 kv get err", err)
			return
		}
		if getResp.Count == 0 {
			fmt.Println("KV 过期了")
			break
		}
		fmt.Println("KV 还没过期", getResp.Kvs)
		time.Sleep(2 * time.Second)
	}

}
