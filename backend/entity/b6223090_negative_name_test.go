package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNameNotBeBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	customer := Customer{
		Name:       "", //false
		Email:      "chanaporn@gmail.com",
		CustomerID: "M6223090",
	}

	ok, err := govalidator.ValidateStruct(customer)

	g.Expect(ok).ToNot(BeTrue())

	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("Name not be blank"))

}
