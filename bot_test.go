package qqbot_for_husthole_test

import (
	"fmt"
	qqbot "github.com/ufatfat/qqbot-for-husthole"
	"testing"
	"time"
)

func TestSendReplyNotice(t *testing.T) {
	bot, _ := qqbot.InitBot("http://localhost:4000/", "http://husthole.com/", "test_account:PivotStudio@2020@tcp(rm-2zeok8s8lj4322wc3do.mysql.rds.aliyuncs.com:3306)/husthole_test?parseTime=True", "39.105.146.221:6379", "Pivot_Studio_2020", 0)
	bot.BindQQ(570407467, "wtfwtf")
	time.Sleep(time.Second * 5)
	bindInfo, _ := bot.GetBindInfo("wtfwtf")
	fmt.Printf("%#v", bindInfo)
}
