package qqbot_for_husthole

import (
"fmt"
"github.com/xioxu/goreq"
"strconv"
"time"
)

func SendReplyNotice (isComment bool, userID uint64, holeID, replyID uint, timestamp time.Time, userAlias, content string) {
	noticeStr := ""
	holeIDStr := strconv.Itoa(int(holeID))
	replyIDStr := strconv.Itoa(int(replyID))
	if isComment {
		noticeStr += userAlias + "回复了您发表的%23" + holeIDStr + "号树洞%0A"
	} else {
		noticeStr += userAlias + "回复了您在%23" + holeIDStr + "号树洞下的回复%0A"
	}
	noticeStr += "时间：" + timestamp.Format("03:04:05") + "%0A"
	noticeStr += "内容：" + content + "%0A"
	noticeStr += "查看回复：" + REDIRECT_SERVER + "?holeID=" + holeIDStr + "%26replyID=" + replyIDStr
	req := goreq.Req(nil)
	url := BOT_SERVER + "send_private_msg?user_id=" + strconv.Itoa(int(userID)) + "&message=" + noticeStr
	fmt.Println("url: ", url)
	fmt.Println("notice: ", noticeStr)
	body, _, _ := req.Get(url).Do()
	fmt.Println(string(body))

	// 分享卡片部分
	/*title := "%23" + strconv.Itoa(int(holeNum))
	shareStr := "[CQ:xml,data=<?xml%20version='1.0'%20encoding='UTF-8'%20standalone='yes'%20?><msg%20serviceID=\"146\"%20temn=\"web\"%20brief=\"%91分享%93%201037树洞\"%20sourceMsgId=\"0\"%20url=\"https://husthole-5gk66z7v90a0a365-1304787517.tcloudbaseapp.com/\"%20flag=\"0\"%20adverSign=\"0\"%20multiMsgFlag=\"0\"><item%20ut=\"2\"%20advertiser_id=\"0\"%20aid=\"0\"><picture%20cover=\"https://qq.ugcimg.cn/v1/e02cjjnid0najlt6pvioi05sevb0h0fko6h6te75kr7glrr2p800/7g7gqb3961mr9o0f8bb3hr7dilff4b73am4cgum8iudjtpnnhmbaocp7c79aqc517e5ks5fjolu00kqd11m3urmgg477sm5rbbqcdu8\"%20w=\"0\"%20h=\"0\"%20/><title>" + title + "</title><summary>" + content + "</summary></item><source%20name=\"QQ浏览l.cn/PWkhNu\"%20url=\"https://url.cn/UQoBHn\"%20action=\"app\"%20a_actionData=\"com.tencent.mtt://https://husthole-5gk66z7v90a0a365-1304787517.tcloudbaseapp.com/\"%20i_actionData=\"tencent100446242://https://husthole-5gk66z7v90a0a365-1304787517.tcloudbaseapp.com/\"%20appid=\"-1\"%20/></msg>,resid=146]https://husthole-5gk66z7v90a0a365-1304787517.tcloudbaseapp.com/"
	url = SERVER + "/send_private_msg?user_id=" + strconv.Itoa(int(userID)) + "&message=" + shareStr
	fmt.Println("url: ", url)
	fmt.Println("share: ", shareStr)
	body, _, _ = req.Get(url).Do()
	fmt.Println(string(body))*/
}