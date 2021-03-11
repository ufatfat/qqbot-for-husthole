package qqbot_for_husthole_test

import (
	qqbot "github.com/ufatfat/qqbot-for-husthole"
	"testing"
	"time"
)

func TestSendReplyNotice(t *testing.T) {
	bot, _ := qqbot.InitBot("http://localhost:4000/", "http://husthole.com/", "test_account:PivotStudio@2020@tcp(rm-2zeok8s8lj4322wc3do.mysql.rds.aliyuncs.com:3306)/husthole_test?parseTime=True", "39.105.146.221:6379", "Pivot_Studio_2020", 0)
	bot.SendReplyNotice(false, 123, 123, time.Now(), "wtfwtf", "东九大宝贝", "%E4%B8%80%E4%B8%AA%E6%B5%8B%E8%AF%95%0A%E4%B8%80%E4%B8%AA%E6%B5%8B%E8%AF%95+%E4%B8%80%E4%B8%AA%E6%B5%8B%E8%AF%95")
}
