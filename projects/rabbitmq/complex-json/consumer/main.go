package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/harshal5-dev/go-forge/projects/rabbitmq/complex-json/models"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	amqpURI := "amqp://admin:12345@localhost:5672/"
	conn, err := amqp.Dial(amqpURI)
	failOnError(err, "Failed to connect to rabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"json_orders_queue",
		true,
		false, false, false, nil,
	)
	failOnError(err, "Failed to declare queue")

	err = ch.Qos(1, 0, false)
	failOnError(err, "Failed to set QoS")

	deliveries, err := ch.Consume(
		q.Name,
		"order-processor-worker",
		false,
		false, false, false, nil,
	)
	failOnError(err, "Failed to register consumer")

	forever := make(chan struct{})

	go func() {
		for d := range deliveries {

			if d.ContentType != "application/json" {
				log.Printf("[!]Warning Unexpected Content-Type %s", d.ContentType)
			}

			var order models.OrderEvent
			err := json.Unmarshal(d.Body, &order)
			if err != nil {
				log.Printf("[x]JSON Parse Error: %v Rejecting message.", err)
				d.Reject(false)
				continue
			}

			fmt.Println("\n--------------------------------------------------")
			fmt.Printf(" [>] Received Order: %s\n", order.OrderID)
			fmt.Printf("     Customer: %s (Tier: %s)\n",
				order.Customer.FullName, order.Customer.Tier)
			fmt.Printf("     Shipping City: %s\n", order.ShippingCity)
			fmt.Printf("     Line Items (%d):\n", len(order.Items))
			for i, it := range order.Items {
				fmt.Printf("       [%d] %-30s x%d ($%.2f each)\n",
					i+1, it.Name, it.Quantity, it.UnitPrice)
			}
			fmt.Printf("     Total Charged: $%.2f %s\n", order.Total, order.Currency)
			fmt.Printf("     Headers: client-version=%v, source=%v\n", d.Headers["client-version"], d.Headers["source"])
			fmt.Println("--------------------------------------------------")

			d.Ack(false)
			fmt.Printf(" [✓] Order %s processed and acknowledged.\n", order.OrderID)
		}
	}()

	fmt.Println(" [*] JSON Order Consumer ready. Waiting for orders...")
	<-forever
}
