package qqbot_for_husthole

func (bot *QQBot) EventAddFriendRequest (userID, botID, encryptedEmail string) (err error) {
	if _, err = bot.Db.Exec("insert into qq_bind (email, user_id, bot_id) values (?,?,?)", encryptedEmail, userID, botID); err != nil {
		// error handle
		return
	}
	return
}