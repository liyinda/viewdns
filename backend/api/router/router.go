package router

import (
    "github.com/gin-gonic/gin"
    . "github.com/liyinda/viewdns/backend/api/apis"
    "net/http"
    //"fmt"
    //"github.com/liyinda/viewdns/backend/middleware/jwt"
    //"github.com/gin-gonic/contrib/sessions"
    "github.com/gin-contrib/cors"
)

func InitRouter() *gin.Engine {
    router := gin.Default()

    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "OPTIONS", "DELETE"},
        AllowHeaders:     []string{"Content-Type","Authorization","X-Token","Access-Control-Allow-Origin"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        AllowOriginFunc: func(origin string) bool {
            return origin == "*"
        },
    }))


    //用户管理入口
    home := router.Group("/home")
    {
        home.GET("/dnslist", Dnslist)
    }



    //定义默认路由
    router.NoRoute(func(c *gin.Context) {
        c.JSON(http.StatusNotFound, gin.H{
            "status": 404,
            "error":  "404, page not exists!",
        })
    })


    return router
}

