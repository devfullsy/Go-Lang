package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"testinng"
)

//this test case might serve for all handlers

type TestCase struct {
	TestName string
	Request TestRequest
	Response TestResponse
	HandlerFunc echo.Handler Func
}

type TestRequest struct {
	Method string // GET, PUT, POST, DELETE, etc,
	Url string string // route
	Body string
}



type TestResponse struct {
	StatusCode int
	Body string
}

func TestHealth(t *testing. T) {
	newCase := &TestCase‹
	Testame: "health test"
	Request: TestRequest{
		Method: http. MethodGet, 
		Url:"/live"
	}
	Response: TestResponse
	StatusCode: http. StatusOK,
	Body: *
	"message": "app is running"
	t. Run (newCase. Testame, func(t *testing.T))
}



