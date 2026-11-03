package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

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
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare queue")

	order := models.OrderEvent{
		OrderID:    "ORD-2026-9812",
		TrackingID: "TRK-FEDEX-88392",
		Customer: models.CustomerInfo{
			CustomerID: "CUST-402",
			FullName:   "Jane Doe",
			Email:      "jane.doe@example.com",
			Tier:       "VIP",
		},
		Items: []models.OrderLineItem{
			{
				ProductID:   "PROD-KB-01",
				Name:        "Mechanical Gaming Keyboard",
				Quantity:    1,
				UnitPrice:   129.99,
				TotalAmount: 129.99,
			},
			{
				ProductID:   "PROD-MS-02",
				Name:        "Wireless Ergonomic Mouse",
				Quantity:    2,
				UnitPrice:   49.50,
				TotalAmount: 99.00,
			},
		},
		Subtotal:     228.99,
		Tax:          18.32,
		Total:        247.31,
		Currency:     "USD",
		ShippingCity: "San Francisco",
		CreatedAt:    time.Now().UTC(),
	}

	payloadBytes, err := json.Marshal(order)
	failOnError(err, "Failed to marshal order info JSON")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(
		ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode:    amqp.Persistent,
			ContentType:     "application/json",
			ContentEncoding: "utf-8",
			MessageId:       order.OrderID,
			Timestamp:       order.CreatedAt,

			Headers: amqp.Table{
				"client-version": "v1.4.0",
				"source":         "web-checkout",
				"customer-tier":  order.Customer.Tier,
			},
			Body: payloadBytes,
		},
	)
	failOnError(err, "Failed to publish JSON order")

	fmt.Println("==================================================")
	fmt.Printf(" [✓] Successfully published JSON Order: %s\n", order.OrderID)
	fmt.Printf("     Customer: %s (%s)\n", order.Customer.FullName, order.Customer.Email)
	fmt.Printf("     Items: %d item(s), Total: $%.2f %s\n",
		len(order.Items), order.Total, order.Currency)
	fmt.Println("==================================================")
}
