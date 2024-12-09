package api

import (
	"BM8/lv4/models"
	"context"
	"fmt"
)

func Pub(ctx context.Context, channel, message string) {
	err := models.Rdb.Publish(ctx, channel, message).Err()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("success to channel:", channel, message)
}

func Sub(ctx context.Context, channel string) {
	pubsub := models.Rdb.Subscribe(ctx, channel)

	wait := pubsub.Channel()
	for msg := range wait {
		if msg.Payload == "stop\n" {
			fmt.Println("stop")
			pubsub.Close()
			return
		} else {
			fmt.Printf("receive message: %s\n", msg.Payload)
		}
	}

}
