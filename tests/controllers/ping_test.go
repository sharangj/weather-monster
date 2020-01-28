package controllers_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sharangj/weather_monster/server"
)

var _ = Describe("Ping", func() {
	It("returns the correct response", func() {
		router := server.Init()

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/ping", nil)
		router.ServeHTTP(w, req)

		Expect(w.Code).To(Equal(200))
		Expect(w.Body.String()).To(Equal("{\"message\":\"pong\"}\n"))
	})
})
