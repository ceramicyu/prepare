package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {

	var (
		config  clientv3.Config
		client  *clientv3.Client
		err     error
		kv      clientv3.KV
		getResp *clientv3.GetResponse
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
	kv = clientv3.NewKV(client)
	fmt.Println(">>", kv)

	//Get操作
	if getResp, err = kv.Get(context.TODO(), "/cron/job1", clientv3.WithPrevKV()); err != nil {
		fmt.Println("clientv3 kv.Get err:", err)

	} else {
		fmt.Println("	getResp.Header.Revision", getResp.Header.Revision)

	}

}
