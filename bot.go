package qqbot_for_husthole

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xioxu/goreq"
	url2 "net/url"
	"strconv"
	"time"
)

type QQBot struct {
	BotServer string
	RedirectServer string
	Rdb *redis.Client
	Db *sql.DB
	ErrQQBound error
}

func InitBot (botServer, redirectServer, mysqlConn, redisConn, redisPswd string, redisDB int) (bot *QQBot, err error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: redisConn,
		Password: redisPswd,
		DB: redisDB,
	})
	db, err := sql.Open("mysql", mysqlConn)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return &QQBot{
		BotServer: botServer,
		RedirectServer: redirectServer,
		Rdb: rdb,
		Db: db,
		ErrQQBound: errors.New("QQ bound"),
	}, nil
}

/*
功能：发送回复通知

参数：

isComment: 是否为对树洞本身的回复

holeID: 树洞号

replyID: 回复ID

timestamp: 回复时间戳

encryptedEmail: 加密后的用户email

userAlias: 树洞昵称

content: 回复内容
 */
func (bot *QQBot) SendReplyNotice (isComment bool, holeID, replyID uint, timestamp time.Time, encryptedEmail, userAlias, content string) (err error) {
	userID, _ := bot.getQQByEncryptedEmail(encryptedEmail)
	noticeStr := ""
	holeIDStr := strconv.Itoa(int(holeID))
	replyIDStr := strconv.Itoa(int(replyID))

	// 回复引用
	/*if original != "" {
		noticeStr += "[CQ:reply,text=Hello%20World,qq=10086,time=3376656000,seq=5123]"
	}*/
	if isComment {
		noticeStr += userAlias + "回复了您发表的%23" + holeIDStr + "号树洞%0A"
	} else {
		noticeStr += userAlias + "回复了您在%23" + holeIDStr + "号树洞下的回复%0A"
	}
	noticeStr += "时间：" + timestamp.Format("03:04:05") + "%0A"
	noticeStr += "内容：" + content + "%0A"
	noticeStr += "查看回复：" + bot.RedirectServer + "?holeID=" + holeIDStr + "%26replyID=" + replyIDStr
	req := goreq.Req(nil)
	url := bot.BotServer + "send_private_msg?user_id=" + strconv.Itoa(int(userID)) + "&message=" + url2.QueryEscape(content)
	fmt.Println("url: ", url)
	fmt.Println("notice: ", noticeStr)
	body, _, err := req.Get(url).Do()
	fmt.Println(string(body))
	return

	// 分享卡片部分
	/*title := "%23" + strconv.Itoa(int(holeNum))
	shareStr := "[CQ:xml,data=<?xml%20version='1.0'%20encoding='UTF-8'%20standalone='yes'%20?><msg%20serviceID=\"146\"%20temn=\"web\"%20brief=\"%91分享%93%201037树洞\"%20sourceMsgId=\"0\"%20url=\"https://husthole-5gk66z7v90a0a365-1304787517.tcloudbaseapp.com/\"%20flag=\"0\"%20adverSign=\"0\"%20multiMsgFlag=\"0\"><item%20ut=\"2\"%20advertiser_id=\"0\"%20aid=\"0\"><picture%20cover=\"https://qq.ugcimg.cn/v1/e02cjjnid0najlt6pvioi05sevb0h0fko6h6te75kr7glrr2p800/7g7gqb3961mr9o0f8bb3hr7dilff4b73am4cgum8iudjtpnnhmbaocp7c79aqc517e5ks5fjolu00kqd11m3urmgg477sm5rbbqcdu8\"%20w=\"0\"%20h=\"0\"%20/><title>" + title + "</title><summary>" + content + "</summary></item><source%20name=\"QQ浏览l.cn/PWkhNu\"%20url=\"https://url.cn/UQoBHn\"%20action=\"app\"%20a_actionData=\"com.tencent.mtt://https://husthole-5gk66z7v90a0a365-1304787517.tcloudbaseapp.com/\"%20i_actionData=\"tencent100446242://https://husthole-5gk66z7v90a0a365-1304787517.tcloudbaseapp.com/\"%20appid=\"-1\"%20/></msg>,resid=146]https://husthole-5gk66z7v90a0a365-1304787517.tcloudbaseapp.com/"
	url = SERVER + "/send_private_msg?user_id=" + strconv.Itoa(int(userID)) + "&message=" + shareStr
	fmt.Println("url: ", url)
	fmt.Println("share: ", shareStr)
	body, _, _ = req.Get(url).Do()
	fmt.Println(string(body))*/
}

func (bot *QQBot) sendSuccessfullyAddedMsg (userID int64) (err error) {
	req := goreq.Req(nil)
	url := bot.BotServer + "send_private_msg?user_id=" + strconv.FormatInt(userID, 10) + "&message=恭喜您已绑定成功！"
	body, _, err := req.Get(url).Do()
	fmt.Println(string(body))
	return
}
