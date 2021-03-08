package qqbot_for_husthole

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

/*func GetQQFromEmail (email string) (recv, send int64) {

}*/

func (bot *QQBot) BindQQ (QQ int64, encryptedEmail string) (err error) {
	key := "bindQQ:" + strconv.FormatInt(QQ, 10)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if err = bot.Rdb.Set(key, encryptedEmail, time.Hour).Err(); err != nil {
		fmt.Println(err.Error())
	}
	return
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
			userID := json["user_id"].(string)
			botID := json["self_id"].(string)
			val, err := bot.Rdb.Get("bindQQ:" + json["user_id"].(string)).Result()
			if err != nil {
				// error handle
				return
			}
			if err = bot.EventAddFriendRequest(userID, botID, val); err != nil {
				// err handle
			}
			return
		}
		
	}
}