package controllers_test

import (
	"io/ioutil"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sharangj/weather_monster/config"
)

func TestControllers(t *testing.T) {
	config.Init("test")
	gin.SetMode(gin.TestMode)
	gin.DefaultWriter = ioutil.Discard
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controllers Tests")
}
