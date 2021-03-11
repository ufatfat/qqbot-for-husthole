package qqbot_for_husthole

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xioxu/goreq"
	"strconv"
	"time"
)
/* 绑定QQ信息

 参数信息：

 BindingStatus：绑定状态 //0: 未绑定; 1: 正在绑定; 2: 已绑定

 BindQQ：绑定的QQ */
type BindInfo struct {
	BindingStatus int8
	BindQQ int64
}

/*
 功能: 绑定QQ至用户账号

 参数:

 QQ: 用户要绑定的qq

 encryptedEmail: 用户加密后的email
 */
func (bot *QQBot) BindQQ (QQ int64, encryptedEmail string) (err error) {
	QQStr := strconv.FormatInt(QQ, 10)
	if err = bot.Db.QueryRow("select id from qq_bind where user_id=?", QQ).Err(); err == nil {
		return errors.New("QQ bound")
	}
	key := "bindQQ:" + QQStr
	if err = bot.Rdb.Set(key, encryptedEmail, time.Hour).Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
	key = "bindEmail:" + encryptedEmail
	if err = bot.Rdb.Set(key, QQStr, time.Hour).Err(); err != nil {
		fmt.Println(err.Error())
	}
	return
}

func (bot *QQBot) getQQByEncryptedEmail (encryptedEmail string) (QQ int64, err error) {
	row := bot.Db.QueryRow("select user_id from qq_bind where email=?", encryptedEmail)
	if err = row.Scan(&QQ); err == sql.ErrNoRows {
		return
	} else if err != nil {
		fmt.Println(err.Error())
	}
	return
}
/*
 功能: 获取用户的QQ绑定状态

 参数: encryptedEmail: 加密后的email

 返回值: bindInfo BindInfo, err error
 */
func (bot *QQBot) GetBindInfo (encryptedEmail string) (bindInfo BindInfo, err error) {
	var bindQQ int64
	row := bot.Db.QueryRow("select user_id from qq_bind where email=? and is_deleted=0", encryptedEmail)
	if err = row.Scan(&bindQQ); err == sql.ErrNoRows {
		if QQ, err1 := bot.Rdb.Get("bindEmail:" + encryptedEmail).Int64(); err == redis.Nil {
			return BindInfo{
				BindQQ: 0,
				BindingStatus: 0,
			}, nil
		} else if err1 != nil {
			return
		} else {
			return BindInfo{
				BindQQ: QQ,
				BindingStatus: 1,
			}, nil
		}
	} else if err != nil {
		fmt.Println("err: ", err.Error())
		return
	} else {
		return BindInfo{
			BindQQ: bindQQ,
			BindingStatus: 2,
		}, nil
	}
}

func (bot *QQBot) approveAddFriendRequest (approve bool, flag string) {
	req := goreq.Req(nil)
	url := bot.BotServer + "set_friend_add_request?approve=" + strconv.FormatBool(approve) + "&flag=" + flag
	body, _, _ := req.Get(url).Do()
	fmt.Println(string(body))
}


func (bot *QQBot) BotEventHandler (c *gin.Context) {
	json := make(map[string]interface{})
	if err := c.BindJSON(&json); err != nil {
		fmt.Println(err.Error())
		return
	}
	switch json["post_type"].(string) {
	case "request":
		switch json["request_type"].(string) {
		case "friend":
			userID := int64(json["user_id"].(float64))
			botID := int64(json["self_id"].(float64))
			val, err := bot.Rdb.Get("bindQQ:" + strconv.FormatInt(userID, 10)).Result()
			if err == redis.Nil {
				// error handle
				fmt.Println("Add friend request error: qq not bound. ")
				bot.approveAddFriendRequest(false, json["flag"].(string))
				return
			} else if err != nil {
				fmt.Println("Add friend request error: ", err.Error())
				return
			}
			bot.approveAddFriendRequest(true, json["flag"].(string))
			if err = bot.eventAddFriendRequest(userID, botID, val); err != nil {
				// err handle
			}

			return
		}
		
	}
}