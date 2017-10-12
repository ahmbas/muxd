package redis

import (
	"fmt"

	redisClient "github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

func (output Connection) Tx(msg []byte) {
	client := redisClient.NewClient(&redisClient.Options{
		Addr:     fmt.Sprintf("%v:%v", output.Host, output.Port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	go func() {

		err := client.Publish(
			output.Channel,
			string(msg),
		).Err()

		if err != nil {
			log.Warnf(
				"Could not forward the following data %v %v %v",
				output,
				string(msg),
				err.Error(),
			)
			return
		}
	}()
	log.Debugf("Sent %v to %v", string(msg), output)
}
