package database
//package main

import (
    "time"
    "context"
    "fmt"
    //"log"
    "go.etcd.io/etcd/clientv3"
    "encoding/json"
    "github.com/tidwall/gjson"
    "strings"
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

//字符串反转处理
func Reversing(src string) (str string) {
    //去掉最左边的/skydns
    src = strings.TrimLeft(src, "/skydns")
    //将字符串转数组
    slice := strings.Split(src, "/")
    //反转数组
    for i, j := 0, len(slice) - 1; i < j; i, j = i + 1, j - 1{
       slice[i], slice[j] = slice[j], slice[i]
    }
    //将数组转字符串，并将字符串以.拼接
    str = strings.Replace(strings.Trim(fmt.Sprint(slice), "[]"), " ", ".", -1)
    if string([]byte(str)[:1]) == "." {
        //去掉字符串中出现的第一个.
        str = strings.Replace(str, ".", "", 1)
        return
    }
 
    return
}


func GetEtcdPrefix() (value string, err error) {

    var (
        config  clientv3.Config
        client  *clientv3.Client
        kv      clientv3.KV
        getResp *clientv3.GetResponse
        m map[string]interface{}
        s []map[string]interface{}
    )

    //var m map[string]interface{}
    //var s []map[string]interface{}

    cli, err := clientv3.New(clientv3.Config{
        Endpoints:   []string{ ip + ":" + port },
        DialTimeout: 5 * time.Second,
    })
    if err != nil {
        // handle error!
    }
    defer cli.Close()

    //客户端配置
    config = clientv3.Config{
        Endpoints:   []string{ ip + ":" + port }, 
        DialTimeout: 1 * time.Second,
    }
    //建立客户端
    if client, err = clientv3.New(config); err != nil {
        fmt.Println(err)
    }
    defer client.Close()

    //用于读写etcd的键值对
    kv = clientv3.NewKV(client)

    //读取/skydns/为前缀的所有key
    if getResp, err = kv.Get(context.TODO(), "/skydns/", clientv3.WithPrefix()); err != nil {
        fmt.Println(err)
    } else { //获取成功，我们遍历所有kvs
        //fmt.Println(reflect.TypeOf(getResp.Kvs))
        for _, ev := range getResp.Kvs {
            fmt.Printf("%s : %s\n", ev.Key, ev.Value)

            //定义一个slice，元素是map
            //var m map[string]interface{}
            //var s []map[string]interface{}

            host := gjson.Get(string(ev.Value), "host").String()
            ttl := gjson.Get(string(ev.Value), "ttl").String()
            domainname := Reversing(string(ev.Key))
            fmt.Println(host)  

            m = make(map[string]interface{})
            m["id"] = "650000199402171139"
            m["domainname"] = domainname
            m["status"] = "deleted"
            m["type"] = "A"
            m["rdata"] = host
            m["ttl"] = ttl
            s = append(s, m)


            data, _ := json.Marshal(s)
            value = string(data)
        }
    }

    return
}

//func GetEtcdPrefix() (value string) {
//
////    value = `[{"id": "650000199402171139","domainname": "www.baidu.com","status": "deleted","type": "name","rdata": "10.10.10.10","ttl": "4316"},{"id": "650000199402171131","domainname": "sdfsdf","status": "deleted","type": "name","rdata": "2017-03-11 05:35:45","ttl": "4317"}]`
//
//type DomainName struct {
//    Id string `json:"id"`
//    Domainname string `json:"domainname"`
//    Status string `json:"status"`
//    Type string `json:"type"`
//    Ttl string `json:"ttl"`
//    Rdata string `json:"rdata"`
//}
//
//    var json DomainName
//
//    json.Id = "123"
//    json.Domainname = "www.g.cn"
//
//    fmt.Println(json) 
//
//    return
//
//}
