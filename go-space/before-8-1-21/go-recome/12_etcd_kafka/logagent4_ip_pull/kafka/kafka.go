package kafka

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

// 专门往kafka写日志的模块
type logData struct {
	topic string
	data  string
}

var (
	client      sarama.SyncProducer // 声明一个全局的连接kafka的生产者client
	logDataChan chan *logData
)

// Init 初始化client
func Init(addrs []string, maxSize int) (err error) {
	config := sarama.NewConfig()
	// tailf包使⽤
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出⼀个 partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 连接kafka
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}

	// Init logDataChan
	logDataChan = make(chan *logData, maxSize)
	// 开启后台的goroutine从通道logDataChan中取数据发往kafka
	go sendToKafka()
	return
}

// 给外部暴露的一个函数，该函数只把日志数据发送到一个内部的channel中
func SendToChan(topic, data string) {
	msg := &logData{
		topic: topic,
		data:  data,
	}

	//将从taillog获取 topic, line data, 发送到当下kafka内部channel
	logDataChan <- msg
}

// 从logDataChan通道中读取数据 真正往kafka发送日志 function
func sendToKafka() {
	for {
		select {
		case ld := <-logDataChan: // get logData struct
			// 构造⼀个消息
			msg := &sarama.ProducerMessage{}
			msg.Topic = ld.topic
			msg.Value = sarama.StringEncoder(ld.data)
			// 发送到kafka
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				fmt.Println("send msg failed, err:", err)
				return
			}
			fmt.Printf("pid:%v offset:%v\n", pid, offset)
		default:
			time.Sleep(time.Millisecond * 5) // wait back handle
		}
	}
}
