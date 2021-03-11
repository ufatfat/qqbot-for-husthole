package qqbot_for_husthole

import (
	"database/sql"
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

 BindingStatus：绑定状态 //0: 无绑定; 1: 正在绑定; 2: 已绑定

 BindQQ：绑定的QQ */
type BindInfo struct {
	BindingStatus int8
	BindQQ int64
}

func (bot *QQBot) BindQQ (QQ int64, encryptedEmail string) {
	QQStr := strconv.FormatInt(QQ, 10)
	key := "bindQQ:" + QQStr
	if err := bot.Rdb.Set(key, encryptedEmail, time.Hour).Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
	key = "bindEmail:" + encryptedEmail
	if err := bot.Rdb.Set(key, QQStr, time.Hour).Err(); err != nil {
		fmt.Println(err.Error())
	}
	return
}

func (bot *QQBot) getQQByEncryptedEmail (encryptedEmail string) (QQ int64, err error) {
	row := bot.Db.QueryRow("select user_id from qq_bind where email=?", encryptedEmail)
	if err = row.Scan(&QQ); err != nil {
		fmt.Println(err.Error())
	}
	return
}

func (bot *QQBot) GetBindInfo (encryptedEmail string) (bindInfo BindInfo, err error) {
	var bindQQ int64
	row := bot.Db.QueryRow("select user_id from qq_bind where email=?", encryptedEmail)
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
			userID := json["user_id"].(int64)
			botID := json["self_id"].(int64)
			val, err := bot.Rdb.Get("bindQQ:" + json["user_id"].(string)).Result()
			if err == redis.Nil {
				// error handle
				fmt.Println("Add friend request error: qq not bind. ")
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