package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// 连接kafka消费消息

func main() {
	// consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	consumer, err := sarama.NewConsumer([]string{"192.168.0.107:9092"}, nil) // 构建 consumer obj
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions("web_log") // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println("Partition list: ", partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
			}
		}(pc)
	}
	select {} // 阻止 主goroutine end, 以便 from kafka recv data
}
