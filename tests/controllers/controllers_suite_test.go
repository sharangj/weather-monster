package controllers_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sharangj/weather_monster/config"
)

func TestControllers(t *testing.T) {
	config.Init("test")
	gin.SetMode(gin.TestMode)
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controllers Tests")
}
