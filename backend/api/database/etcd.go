package database
//package main

import (
        "time"
        "context"
        "fmt"
        //"log"
        "go.etcd.io/etcd/clientv3"
)

var ip string = "172.20.57.58"
var port string = "2379"

func Test() (e string) {
    e = "sdfsdf"
    fmt.Println(e)
    return 
}

func GetEtcdkey(key string) (value string, err error) {
//func main() {
    cli, err := clientv3.New(clientv3.Config{
        Endpoints:   []string{ ip + ":" + port },
        DialTimeout: 5 * time.Second,
    })
    if err != nil {
        // handle error!
    }
    defer cli.Close()

    ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
    resp, err := cli.Get(ctx, "/skydns/")
    cancel()
    if err != nil {
        //log.Fatal(err)
    }
    fmt.Println(resp)

    for _, ev := range resp.Kvs {
        fmt.Printf("%s : %s\n", ev.Key, ev.Value)
        value =  string(ev.Value)
    }
    return 

}


