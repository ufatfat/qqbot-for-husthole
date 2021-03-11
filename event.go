package qqbot_for_husthole

import (
	"fmt"
	"strconv"
)

func (bot *QQBot) eventAddFriendRequest (userID, botID int64, encryptedEmail string) (err error) {
	if _, err = bot.Db.Exec("insert into qq_bind (email, user_id, bot_id) values (?,?,?)", encryptedEmail, userID, botID); err != nil {
		fmt.Println("err: ", err.Error())
		return
	}
	bot.Rdb.Del("bindQQ:" + strconv.FormatInt(userID, 10), "bindEmail:" + encryptedEmail)
	if err = bot.sendSuccessfullyAddedMsg(userID); err != nil {
		fmt.Println("err: ", err.Error())
	}
	return
}