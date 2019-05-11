package handlers

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Health Check Handler", func() {
	var (
		rr *httptest.ResponseRecorder
	)
	BeforeEach(func() {
		req, err := http.NewRequest("GET", "/health-check", nil)
		Expect(err).NotTo(HaveOccurred())

		rr = httptest.NewRecorder()
		handler := http.HandlerFunc(HealthCheckHandler)

		handler.ServeHTTP(rr, req)
	})

	It("returns 200", func() {
		Expect(rr.Code).To(Equal(http.StatusOK))
	})

	It("returns 'alive' payload", func() {
		Expect(rr.Body.String()).To(Equal(`{ "alive": true }`))
	})
})
