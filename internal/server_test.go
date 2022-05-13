package internal_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Dan6erbond/go-fx-gin-template/internal"
	"github.com/Dan6erbond/go-fx-gin-template/internal/config"
	"github.com/Dan6erbond/go-fx-gin-template/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type ServerTestSuite struct {
	suite.Suite
	router *gin.Engine
}

func (suite *ServerTestSuite) SetupTest() {
	suite.router = internal.MakeServer(&config.AppConfig{
		Env: "test",
	})
}

func (suite *ServerTestSuite) TestHelloWorldRoute() {
	w := httptest.NewRecorder()
	data := dto.HelloWorldRequest{
		Name: "John Doe",
	}
	body, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", "/api/v1/hello-world", bytes.NewBuffer(body))
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusAccepted, w.Code)

	var resp dto.HelloWorldResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	suite.Equal("Hello John Doe!", resp.Message)
}

func TestServer(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}
