package main

import (
	"wander_fare_tracker/services/kafka"
)

func main() {
	kafka.Producer()
	kafka.StartConsumer()

}
