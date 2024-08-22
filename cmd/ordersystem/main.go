package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/drawiin/go-orders-service/config"
	"github.com/drawiin/go-orders-service/internal/event/handler"
	"github.com/drawiin/go-orders-service/internal/infra/web/webserver"
	"github.com/drawiin/go-orders-service/pkg/events"
	"github.com/streadway/amqp"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config := getAppConfig()

	dbConnection := getDbConnection(config)
	defer dbConnection.Close()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("order.created", handler.NewOrderCreatedHandler(getRabbitMQChannel(config)))

	// Setup WebServer
	webserver := webserver.NewWebServer(config.WebServerPort)
	// webserver.AddHandler("/orders", NewWebOrderHandler(dbQueries, eventDispatcher))

	// Start web server
	webserver.Start()
}

func getDbConnection(config *config.Config) *sql.DB {
	conn, err := sql.Open(config.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName))
	if err != nil {
		panic(err)
	}
	return conn
}

func getAppConfig() *config.Config {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	return config
}

func getRabbitMQChannel(config *config.Config) *amqp.Channel {
	connectionString := fmt.Sprintf("amqp://%s:%s@%s:%s/", config.QueueUser, config.QueuePassword, config.QueueHost, config.QueuePort)
	conn, err := amqp.Dial(connectionString)
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}
