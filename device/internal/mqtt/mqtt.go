package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"iot-platform/models"
	"log"
	"strings"
)

var topic = "/sys/#"

func NewMqttServer(mqttBroker string) {
	opt := mqtt.NewClientOptions().AddBroker(mqttBroker).SetClientID("go-mqtt-server-client-id").
		SetUsername("get").SetPassword("123456")

	// 回调
	opt.SetDefaultPublishHandler(PublishHandler)
	c := mqtt.NewClient(opt)

	// 连接
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// 订阅主题
	if token := c.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	defer func() {
		// 取消订阅
		if token := c.Unsubscribe(topic); token.Wait() && token.Error() != nil {
			log.Println("[ERROR] : ", token.Error())
		}
		// 关闭连接
		c.Disconnect(250)
	}()

	select {}

}

func PublishHandler(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("MESSAGE : %s\n", message.Payload())
	fmt.Printf("TOPIC : %s\n", message.Topic())
	topicArray := strings.Split(strings.TrimPrefix(message.Topic(), "/"), "/")
	if len(topicArray) >= 4 {
		if topicArray[3] == "ping" {
			err := models.UpdateDeviceOnlineTime(topicArray[1], topicArray[2])
			if err != nil {
				log.Printf("[DB ERROR : %v\n", err)
			}
		}

	}
}