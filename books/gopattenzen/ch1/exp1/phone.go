package main

// 不好的接口设计，这里有两个职责，连接、关闭和聊天
type IPhone interface {
	dial(phoneNumber string)
	chat(o interface{})
	hangup()
}

func main() {

}
