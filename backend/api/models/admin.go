package models

import (
    //"fmt"
    orm "github.com/liyinda/viewdns/backend/api/database"
)

// 用户登录 from JSON
type LoginJson struct {
        Loginname     string `form:"loginname" json:"loginname" xml:"loginname"  binding:"required"`
        Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

//管理员结构体
type Admin struct {
    Login_name string `json:"login_name"`
    Password string `json:"password"`
}

//获取管理员密码
func (admins *Admin) GetPassword(login_name string) (password string, err error) {
    orm.GetEtcdkey(login_name)
    return

}
