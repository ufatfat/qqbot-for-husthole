package qqbot_for_husthole_test

import (
	qqbot "github.com/ufatfat/qqbot-for-husthole"
	"testing"
	"time"
)

func TestSendReplyNotice(t *testing.T) {
	qqbot.SendReplyNotice(false, 570407467, 123, 0, time.Now(), "东九大宝贝", "一个测试")
}
