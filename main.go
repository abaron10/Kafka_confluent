package main

import (
	"context"
	"kafka-example2/consumer"
)

func main() {
	// create a new context
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	//go producer.Produce(ctx)
	consumer.Consume(ctx)
}
