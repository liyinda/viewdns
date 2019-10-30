package models

import (
    //"fmt"
    orm "github.com/liyinda/viewdns/backend/api/database"
)

type DNS_A struct {
    Domain_name string `json:"domain_name"`
    Ip_addr string `json:"ip_addr"`
    TTL int64  `json:"ttl"`
}

type PublicKey struct {
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
func (a_records DNS_A) Dnslist() (id string) {
    id = orm.Test()
    return

}

