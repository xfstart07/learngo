// Author: xufei
// Date: 2019/2/20

package main

import (
	"flag"
	app2 "learngo/mq/rocketmq/consumer/app"
	"log"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"

	"github.com/gogap/errors"

	"gopkg.in/olivere/elastic.v6"

	mq_http_sdk "github.com/aliyunmq/mq-http-go-sdk"
)

var (
	zlog  *zap.SugaredLogger
	esurl = flag.String("esurl", "http://weixinote.dev:9200", "esurl=http://weixinote.dev:9200")
	level = flag.String("level", "debug", "level=debug")
)

func main() {
	flag.Parse()

	if *level == "debug" {
		logg, _ := zap.NewDevelopment()
		zlog = logg.Sugar()
	} else {
		logg, _ := zap.NewProduction()
		zlog = logg.Sugar()
	}

	newClient()
	mqConsumer := newConsumer()

	esClient, err := newES(*esurl)
	if err != nil {
		panic(err)
	}

	job := make(chan string)
	app2.CreateWrite(esClient, job)

	for {
		endChan := make(chan int)
		respChan := make(chan mq_http_sdk.ConsumeMessageResponse)
		errChan := make(chan error)
		go func() {
			select {
			case resp := <-respChan:
				{
					// 处理业务逻辑
					var handles []string
					zlog.Infof("Consume %d messages---->\n", len(resp.Messages))
					for _, v := range resp.Messages {
						handles = append(handles, v.ReceiptHandle)
						zlog.Debugf("\tMessageID: %s, PublishTime: %d, MessageTag: %s\n"+
							"\tConsumedTimes: %d, FirstConsumeTime: %d, NextConsumeTime: %d\n\tBody: %s\n",
							v.MessageId, v.PublishTime, v.MessageTag, v.ConsumedTimes,
							v.FirstConsumeTime, v.NextConsumeTime, v.MessageBody)

						// write body
						job <- v.MessageBody
					}

					// NextConsumeTime前若不确认消息消费成功，则消息会重复消费
					// 消息句柄有时间戳，同一条消息每次消费拿到的都不一样
					ackerr := mqConsumer.AckMessage(handles)
					if ackerr != nil {
						// 某些消息的句柄可能超时了会导致确认不成功
						zlog.Info(ackerr)
						for _, errAckItem := range ackerr.(errors.ErrCode).Context()["Detail"].([]mq_http_sdk.ErrAckItem) {
							zlog.Errorf("\tErrorHandle:%s, ErrorCode:%s, ErrorMsg:%s\n",
								errAckItem.ErrorHandle, errAckItem.ErrorCode, errAckItem.ErrorMsg)
						}
						time.Sleep(time.Duration(3) * time.Second)
					} else {
						zlog.Debugf("Ack ---->\n\t%s\n", handles)
					}

					endChan <- 1
				}
			case err := <-errChan:
				{
					// 没有消息
					if strings.Contains(err.(errors.ErrCode).Error(), "MessageNotExist") {
						zlog.Info("\nNo new message, continue!")
					} else {
						zlog.Error(err)
						time.Sleep(time.Duration(3) * time.Second)
					}
					endChan <- 1
				}
			case <-time.After(35 * time.Second):
				{
					zlog.Warn("Timeout of consumer message ??")
					endChan <- 1
				}
			}
		}()

		// 长轮询消费消息
		// 长轮询表示如果topic没有消息则请求会在服务端挂住3s，3s内如果有消息可以消费则立即返回
		mqConsumer.ConsumeMessage(respChan, errChan,
			3, // 一次最多消费3条(最多可设置为16条)
			3, // 长轮询时间3秒（最多可设置为30秒）
		)
		<-endChan
	}
}

var clientMQ mq_http_sdk.MQClient

func newClient() {
	// 设置HTTP接入域名（此处以公共云生产环境为例）
	endpoint := ""
	// AccessKey 阿里云身份验证，在阿里云服务器管理控制台创建
	accessKey := ""
	// SecretKey 阿里云身份验证，在阿里云服务器管理控制台创建
	secretKey := ""

	clientMQ = mq_http_sdk.NewAliyunMQClient(endpoint, accessKey, secretKey, "")
}

func newConsumer() mq_http_sdk.MQConsumer {
	// 所属的 Topic
	topic := "songplay"
	// Topic所属实例ID，默认实例为空
	instanceId := ""
	// 您在控制台创建的 Consumer ID(Group ID)
	groupId := ""

	return clientMQ.GetConsumer(instanceId, topic, groupId, "")
}

func newES(url string) (*elastic.Client, error) {
	return elastic.NewClient(elastic.SetURL(url),
		elastic.SetSniff(false),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
}
