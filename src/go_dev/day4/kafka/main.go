package main

import "fmt"
import "github.com/Shopify/sarama"


func main(){
	config := sarama.NewConfig()
	config.Producer.Requirement = sarama.WaitForAll()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Successess = true

	msg := &sarama.ProducerMessage{}
	msg.Topic = "nginx.log"
	msg.Value = sarama.StringEncoder("thi is a good test my message is good")
	client,err := sarama.NewSyncProducer([]string{"192.168.1.1:9092"},config)
	if err != nil{
		fmt.Println("Producer close,err:",err)
		return
	}
	defer client.Close()
	pid,offset,err := client.SendMessage(msg)
	if err != nil{
		fmt.Println("send message failed:",err)
		return
	}
	fmt.Println("pid:%v offset:%v\n",pid,offset)





}
