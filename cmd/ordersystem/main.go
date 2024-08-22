package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/drawiin/go-orders-service/config"
	"github.com/drawiin/go-orders-service/internal/event/handler"
	"github.com/drawiin/go-orders-service/internal/infra/graph"
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

	go startWebServer(config, dbConnection, eventDispatcher)
	startGraphQLServer(config, dbConnection, eventDispatcher)
}

func startWebServer(config *config.Config, dbConnection *sql.DB, eventDispatcher *events.EventDispatcher) {
	webserver := webserver.NewWebServer(config.WebServerPort)
	webHandler := NewWebOrderHandler(dbConnection, eventDispatcher)
	webserver.AddHandler("/orders/create", webHandler.Create)
	webserver.AddHandler("/orders/list", webHandler.GetAll)
	webserver.AddHandler("/orders/{id}", webHandler.GetById)
	fmt.Println("Starting Web server on  port", config.WebServerPort)
	webserver.Start()
}

func startGraphQLServer(config *config.Config, dbConnection *sql.DB, eventDispatcher *events.EventDispatcher) {
	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: NewGraphQLResolver(dbConnection, eventDispatcher)}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	fmt.Println("Starting GraphQL server on port", config.GraphQLServerPort)
	http.ListenAndServe(":"+config.GraphQLServerPort, nil)
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
