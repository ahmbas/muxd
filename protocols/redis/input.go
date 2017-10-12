package redis

import redisClient "github.com/go-redis/redis"
import "fmt"
import log "github.com/sirupsen/logrus"

func (input Connection) Rx(c chan []byte) chan []byte {

	log.Debugf("Starting Input on %v", input)

	go func() {
		client := redisClient.NewClient(&redisClient.Options{
			Addr: fmt.Sprintf("%v:%v", input.Host, input.Port),
		})
		pubsub := client.Subscribe(input.Channel)
		for {
			msg, err := pubsub.ReceiveMessage()
			if err != nil {
				log.Fatalf("Error reading %v %v", input, err.Error())
			}

			log.Debugf("Got %v", msg.Payload)
			c <- []byte(msg.Payload)
		}

	}()
	return c
}
