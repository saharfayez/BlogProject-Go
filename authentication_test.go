package main_test

import (
	"bytes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"net/http"

	_ "goproject"
)

var _ = Describe("Authentication", func() {
	Context("When successfully creating a new user", func() {
		It("Should be able to create a new user", func() {
			requestBody :=
				`{
				"username":"ameen",
				"password":"0000"
			}`
			request, _ := http.NewRequest("POST", "http://localhost:8080/signup", bytes.NewBuffer([]byte(requestBody)))
			request.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			response, err := client.Do(request)
			Expect(err).NotTo(HaveOccurred())
			Expect(response.StatusCode).To(Equal(http.StatusCreated))
		})
	})

})
