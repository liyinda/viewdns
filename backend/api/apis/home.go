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
