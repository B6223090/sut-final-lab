package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestCorretAll(t *testing.T) {
	g := NewGomegaWithT(t)

	customer := Customer{
		Name:       "chanaporn",
		Email:      "chanaporn@gmail.com",
		CustomerID: "M6223090",
	}

	ok, err := govalidator.ValidateStruct(customer)

	g.Expect(ok).To(BeTrue())

	g.Expect(err).To(BeNil())

}
