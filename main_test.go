package main_test

import (
	models "a21hc3NpZ25tZW50/models"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {

	Describe("Connet", func() {
		It("should connect to the database management", func() {
			_, err := models.ConnectDatabase()

			Expect(err).To(BeNil())
		})
	})
	Describe("Connet", func() {
		It("should connect to the database management", func() {
			a, _ := models.ConnectDatabase()

			Expect(a).To(Equal(models.DB))
		})
	})
})
