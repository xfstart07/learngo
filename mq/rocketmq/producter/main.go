// Author: xufei
// Date: 2019/2/20

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	mq_http_sdk "github.com/aliyunmq/mq-http-go-sdk"
)

type songBase struct {
	SongCode   string `json:"song_code"`
	Player     string `json:"player"`
	Title      string `json:"title"`
	CreateTime string `json:"create_time"`
	KtvId      int    `json:"ktv_id,omitempty"`
	KtvName    string `json:"ktv_name,omitempty"`
	KtvCode    string `json:"ktv_code,omitempty"`
	RoomId     string `json:"room_id,omitempty"`
}

type message struct {
	body string
	tag  string
}

var (
	jobCount = flag.Int("job", 1000, "job=1000")
	worker   = flag.Int("worker", 10, "worker=10")
)

func main() {
	flag.Parse()

	newClient()

	bodys := []songBase{}
	for j := 1; j <= 3000; j++ {
		bodys = append(bodys, songBase{
			KtvId:      321,
			KtvName:    "测试版权盒绑定",
			KtvCode:    "K0000320",
			RoomId:     "A01",
			SongCode:   "$abse1221",
			Title:      "K歌之王",
			Player:     "陈奕迅",
			CreateTime: time.Now().Format("2006-01-02 15:04:05"),
		})
	}
	fpara, _ := json.Marshal(bodys)

	jobs := make(chan *message, *worker)
	result := make(chan error)

	for i := 0; i < *worker; i++ {
		createWorker(jobs, result)
	}

	now := time.Now()
	for i := 0; i < *jobCount; i++ {
		go func() {
			jobs <- &message{
				body: string(fpara),
				tag:  "testSong",
			}
		}()
	}
	fmt.Println("send jobs end")

	fmt.Println("start result...")
	for i := 0; i < *jobCount; i++ {
		ret := <-result
		if ret != nil {
			fmt.Println("send Error", i)
		} else {
			fmt.Println("send success", i)
		}
	}
	close(jobs) // 发送完任务，通知结束
	close(result)

	fmt.Println("耗时: ", time.Since(now))
}

func createWorker(jobs <-chan *message, result chan<- error) {
	product := newProduct()

	go func() {
		for job := range jobs {
			result <- send(product, job)
		}
	}()
}

func send(product mq_http_sdk.MQProducer, msg *message) error {
	reqMsg := mq_http_sdk.PublishMessageRequest{
		MessageBody: msg.body, //消息内容
		MessageTag:  msg.tag,  // 消息标签
	}

	ret, err := product.PublishMessage(reqMsg)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Publish ---->\n\tMessageId:%s, BodyMD5:%s, \n", ret.MessageId, ret.MessageBodyMD5)
	}

	return err
}

var clientMQ mq_http_sdk.MQClient

func newClient() {
	// 设置HTTP接入域名（此处以公共云生产环境为例）
	endpoint := os.Getenv("AliEndPoint")
	// AccessKey 阿里云身份验证，在阿里云服务器管理控制台创建
	accessKey := os.Getenv("AliAccessKey")
	// SecretKey 阿里云身份验证，在阿里云服务器管理控制台创建
	secretKey := os.Getenv("AliSecretKey")

	clientMQ = mq_http_sdk.NewAliyunMQClient(endpoint, accessKey, secretKey, "")
}

func newProduct() mq_http_sdk.MQProducer {
	// 所属的 Topic
	topic := "songplay"
	// Topic所属实例ID，默认实例为空
	instanceId := ""

	return clientMQ.GetProducer(instanceId, topic)
}
