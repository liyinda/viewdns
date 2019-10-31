package models

import (
    //"fmt"
    orm "github.com/liyinda/viewdns/backend/api/database"
    "encoding/json"
)


type DomainName struct {
    Id string `json:"id"`
    Domainname string `json:"domainname"`
    Status string `json:"status"`
    Type string `json:"type"`
    Ttl string `json:"ttl"`
    Rdata string `json:"rdata"`
}

//添加dns的a记录
/*func (a_records DNS_A) Dnsadd() (id int64, err error) {
    //etcdctl put 

}
*/


//列出etcd中所有列信息
//etcdctl get /skydns --prefix --keys-only=true
//如果读取不到etcd信息返回异常
func (domainnames *DomainName) Dnslist() (domains []DomainName, err error) {

    //s := `[{"id": "650000199402171139","domainname": "www.baidu.com","status": "deleted","type": "name","rdata": "10.10.10.10","ttl": "4316"},{"id": "650000199402171131","domainname": "sdfsdf","status": "deleted","type": "name","rdata": "2017-03-11 05:35:45","ttl": "4317"}]`
    s, _ := orm.GetEtcdPrefix()

    err = json.Unmarshal([]byte(s), &domains)
    if err != nil {
        return
    }
    return

}

