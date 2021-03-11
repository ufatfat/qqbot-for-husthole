# qqbot-for-husthole
> 华科树洞的QQ机器人功能封装
### 如何使用
1. 获取包
```go
go get github.com/ufatfat/qqbot-for-husthole
```
2. 引入包
```go
import qqbot "github.com/ufatfat/qqbot-for-husthole"
```
3. 将事件监听中间件加入路由树
```go
router.Post("/event_listener", qqbot.BotEventHandler)
```
4. 初始化bot
```go
bot, err := qqbot.InitBot("{botServer}", "{redirectServer}", "{mysqlConn}", "{redisConn}", "{redisPassword}", {redisDB})
```
如：
```go
bot, _ := qqbot.InitBot("http://localhost:2333/", "http://husthole.com/", "YOUR_ACCOUNT:YOUR_PASSWORD@tcp(YOUR_MYSQL_DATABASE:3306)/YOUR_DATABASE?parseTime=True", "YOUR_REDIS_IP:6379", "YOUR_REDIS_PASSWORD", YOUR_REDIS_DB)
```

### API Docs
##### 发送回复通知
原型：
```go
func (bot *QQBot) SendReplyNotice (isComment bool, userID uint64, holeID, replyID uint, timestamp time.Time, userAlias, content, original string) (err error)
```
使用：
```go
bot, _ := qqbot.InitBot("http://localhost:2333/", "http://husthole.com/", "YOUR_ACCOUNT:YOUR_PASSWORD@tcp(YOUR_MYSQL_DATABASE:3306)/YOUR_DATABASE?parseTime=True", "YOUR_REDIS_IP:6379", "YOUR_REDIS_PASSWORD", YOUR_REDIS_DB)

_ = bot.SendReplyNotice(true, 570407467, 123, 1, time.Now(), "东九大宝贝", "一个测试", "测试测试")
```
收到内容：
```
东九大宝贝回复了您发表的#123号树洞
时间：04:23:33
内容：一个测试
查看回复：http://husthole.com/?holeID=123&replyID=1
```