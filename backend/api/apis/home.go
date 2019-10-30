package apis

import (
    "github.com/gin-gonic/gin"
    "github.com/liyinda/viewdns/backend/api/models"
    //"github.com/gin-gonic/contrib/sessions"
    "net/http"
    //"fmt"
    //"strings"
    //"github.com/liyinda/viewdns/backend/pkg/util"
    //"github.com/liyinda/viewdns/backend/pkg/e"
    //"strconv"
    "encoding/json"
)





//查看etcd信息
func Dnslist(c *gin.Context) {
    //获取session中的user信息
    /*session := sessions.Default(c)
    user := session.Get("user")
    code := e.INVALID_PARAMS
    if user == nil {
        code = e.ERROR_AUTH_SESSION
    } else {
        code = e.SUCCESS

    }
    */
    //获取POST中json参数
    var json models.DNS_A

    /*if err := c.ShouldBindJSON(&json); err != nil {
        code = e.ERROR_NOT_JSON
    }
    */
    //获取url中token, page, limit
    //token := c.Request.URL.Query().Get("token")
    /*
    page := c.Request.URL.Query().Get("page")
    limit := c.Request.URL.Query().Get("limit")
 
    pageint, _ := strconv.ParseInt(page, 10, 64)
    limitint, _ := strconv.ParseInt(limit, 10, 64)
    */


    //获取用户信息表
    result := json.Dnslist()
    /*if err != nil{
        code = e.ERROR
    } else {
        code = e.SUCCESS
    }
    */
    //获取总用户数量
    /*
    count, err := json.Usercount()
    if err != nil{
        code = e.ERROR
    } else {
        code = e.SUCCESS
    }
    */

    c.JSON(http.StatusOK, gin.H{
        "msg": result,
    })
}


func Table(c *gin.Context) {
    //获取session中的user信息
    /*session := sessions.Default(c)
    user := session.Get("user")
    code := e.INVALID_PARAMS
    if user == nil {
        code = e.ERROR_AUTH_SESSION
    } else {
        code = e.SUCCESS

    }
    */


    num := 2


    s := `[{"id": "650000199402171139","domainname": "www.baidu.com","status": "deleted","type": "name","rdata": "10.10.10.10","ttl": "4316"},{"id": "650000199402171131","domainname": "sdfsdf","status": "deleted","type": "name","rdata": "2017-03-11 05:35:45","ttl": "4317"}]`

    keys := make([]models.PublicKey,0)
    err := json.Unmarshal([]byte(s), &keys)
    if err == nil {

        data := gin.H{
            "items": keys,
            "total": num,
        }
    
        c.JSON(http.StatusOK, gin.H{
            "code": 20000,
            "data": data,
        })

    }
}
