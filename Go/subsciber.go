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

	queue, err := ch.QueueDeclare(
		"test", //Name of th queue
		false,
		false,
		true,
		false,
		nil,
	)
	failOnError(err, "An error Occured on Queue declaration.")

	err = ch.QueueBind(
		queue.Name, //name of the queue
		"test",     //Routing key
		"test2",    //Name of the exchange
		false,
		nil,
	)
	failOnError(err, "An error Occured on Queue binding.")

	msgs, err := ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "An error Occured on message consumption.")

	forever := make(func name(rw http.ResponseWriter, req *http.Request) {
		
	} bool)

	go func() {
		for m := range msgs {

			log.Printf("Message : %s", m.Body)
		}

	}()

	log.Printf("CRTL + C to close. ")
	<-forever

}
