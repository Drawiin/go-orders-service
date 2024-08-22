package repository

import (
	"context"
	"database/sql"
	"testing"

	"github.com/drawiin/go-orders-service/internal/entity"
	"github.com/drawiin/go-orders-service/internal/infra/db"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	// sqlite3
	_ "github.com/mattn/go-sqlite3"
)

// Setup

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db         *db.Queries
	DbConn     *sql.DB
	Repository entity.OrderRepositoryInterface
}

type MockOrderQueries struct {
	mock.Mock
	db.Queries
}

// We need to setup the suite
func (suite *OrderRepositoryTestSuite) SetupSuite() {
	dbConn, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	dbConn.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.DbConn = dbConn
	suite.Db = db.New(suite.DbConn)
	suite.Repository = NewOrderRepository(suite.Db)
}

// We need to clean the table after each test
func (suite *OrderRepositoryTestSuite) TearDownTest() {
	suite.DbConn.Exec("DELETE FROM orders")
}

// We need to close the connection after all tests
func (suite *OrderRepositoryTestSuite) TearDownSuite() {
	suite.DbConn.Close()
}

// Tests
func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestShouldSaveOrder_WhenSave() {
	order, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())
	repo := NewOrderRepository(suite.Db)
	err = repo.SaveOrder(order)
	suite.NoError(err)

	orderResult, err := suite.Db.GetOrder(context.Background(), order.ID)

	suite.NoError(err)
	suite.Equal(order.ID, orderResult.ID)
	suite.Equal(order.Price, orderResult.Price)
	suite.Equal(order.Tax, orderResult.Tax)
	suite.Equal(order.FinalPrice, orderResult.FinalPrice)
}

func (suite *OrderRepositoryTestSuite) TestShouldGetSavedOrders_WhenGetAllOrders() {
	insertFakeOrders(suite)
	orders, err := suite.Repository.GetAllOrders()

	suite.NoError(err)
	suite.Len(orders, 3)
	suite.ElementsMatch(generateFakeOrders(), orders)
}

func (suite *OrderRepositoryTestSuite) TestShouldGetOrderById_WhenGetOrder() {
	insertFakeOrders(suite)
	order, err := suite.Repository.GetOrder("1")

	suite.NoError(err)
	suite.Equal("1", order.ID)
	suite.Equal(10.0, order.Price)
	suite.Equal(2.0, order.Tax)
	suite.Equal(12.0, order.FinalPrice)
}

func generateFakeOrders() []*entity.Order {
	return []*entity.Order{
		{ID: "1", Price: 10.0, Tax: 2.0, FinalPrice: 12.0},
		{ID: "2", Price: 20.0, Tax: 4.0, FinalPrice: 24.0},
		{ID: "3", Price: 30.0, Tax: 6.0, FinalPrice: 36.0},
	}
}

func insertFakeOrders(suite *OrderRepositoryTestSuite) {
	orders := generateFakeOrders()
	for _, order := range orders {
		order.CalculateFinalPrice()
		repo := NewOrderRepository(suite.Db)
		repo.SaveOrder(order)
	}
}
