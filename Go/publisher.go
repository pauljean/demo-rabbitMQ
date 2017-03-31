package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func failOnError(err error, msg string) {

	if err != nil {

		log.Fatalf("%s : %s", msg, err)
		panic(fmt.Sprintf("%s : %s", msg, err))

	}
}

func main() {

	connection, err := amqp.Dial("amqp://pi:pi@10.226.159.191:5672//pi")
	failOnError(err, "An error occured on connection.")

	ch, err := connection.Channel()
	failOnError(err, "An error occured on connection to channel")

	defer ch.Close()

	err = ch.ExchangeDeclare(
		"test2", //Name of the exchange
		"topic", //Exchange mode
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "An error occured on exchange declaration.")

	err = ch.Publish(
		"test2", //Name of the exchange
		"test",  // Routing key
		false,
		false,
		amqp.Publishing{

			ContentType: "text/plain",
			Body:        []byte("salut"),
		},
	)

	fmt.Println("The message as been sent")

	failOnError(err, "Failled to publish a message")

}
