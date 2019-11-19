package main

import (
    "time"
    "context"
    "fmt"
    //"log"
    "go.etcd.io/etcd/clientv3"
    "strings"
)

var ip string = "172.20.57.58"
var port string = "2379"

//删除etcd数据
func DelEtcdKV(key string) int64 {
    cli, err := clientv3.New(clientv3.Config{
        Endpoints:   []string{ ip + ":" + port },
        DialTimeout: 5 * time.Second,
    })
    if err != nil {
        // handle error!
    }
    defer cli.Close()

    //ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
    //将域名转换为etcd的key值
    key = ReversingDomain(key)

    res, err := cli.Delete(context.TODO(), key, clientv3.WithPrefix())
    if err != nil {
        fmt.Println("errrr")
        return 0
    }
    return res.Deleted
}

//字符串反转处理,将域名反转为etcd的key值
func ReversingDomain(src string) (str string) {
    //将字符串转数组
    slice := strings.Split(src, ".")
    //反转数组
    for i, j := 0, len(slice) - 1; i < j; i, j = i + 1, j - 1{
       slice[i], slice[j] = slice[j], slice[i]
    }
    //将数组转字符串，并将字符串以/拼接
    str = strings.Replace(strings.Trim(fmt.Sprint(slice), "[]"), " ", "/", -1)
    if string([]byte(str)[:1]) == "/" {
        //去掉字符串中出现的第一个/
        str = strings.Replace(str, "/", "", 1)
    }

    //str = "/skydns/" + str + "/"
    str = "/skydns/" + str
    return
}

func main() {
    s := DelEtcdKV("www1.skydns.local")
    fmt.Println(s)
}
