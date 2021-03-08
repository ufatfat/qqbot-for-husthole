package qqbot_for_husthole_test

import (
	qqbot "github.com/ufatfat/qqbot-for-husthole"
	"testing"
	"time"
)

func TestSendReplyNotice(t *testing.T) {
	bot, _ := qqbot.InitBot("http://localhost:4000/", "http://husthole.com/", "test_account:PivotStudio@2020@tcp(rm-wz9zx0212vs57636pvo.mysql.rds.aliyuncs.com:3306)/husthole_test?parseTime=True", "39.105.146.221:6379", "Pivot_Studio_2020", 0)
	bot.BindQQ(570407467, "asdcas123012")
	//bot.EventAddFriendRequest("570407467", "732343768", "sadasfasdasd")
	bot.SendReplyNotice(true, 570407467, 123, 1, time.Now(), "东九大宝贝", "一个测试", "测试测试")
}
