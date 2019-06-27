package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"time"
)

func main(){

	var (
		config clientv3.Config
		client *clientv3.Client
		err  error
		kv clientv3.KV
		delResp *clientv3.DeleteResponse
		kvPair *mvccpb.KeyValue

	)

	config=clientv3.Config{
		Endpoints:[]string{"0.0.0.0:2379"},//集群列表
		DialTimeout:5*time.Second,
	}

	//建立一个客户端
	if client,err=clientv3.New(config);err!=nil{
		fmt.Println("clientv3 config err:",err)
		return
	}

    //用于读写etcd的键值对
	kv=clientv3.NewKV(client)
	fmt.Println(">>",kv)

	kv.Put(context.TODO(),"/cron/job3","{....}")
	//Delete操作
	//option 可以追加多个
	//clientv3.WithPrefix()
	if delResp,err=kv.Delete(context.TODO(),"/cron/job2",clientv3.WithPrevKV());err!=nil{
		fmt.Println(" clientv3 kv delete err",err)
		return
	}else{

		//被删除之前的val
		if len(delResp.PrevKvs)!=0{
			for _,kvPair=range delResp.PrevKvs{
				fmt.Println("删除了：",string(kvPair.Key),string(kvPair.Value))
			}
		}
	}

}
