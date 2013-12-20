package spec_test

import (
	"github.com/adamkalnas/secret_santa/secret_santa"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SecretSanta", func() {
	It("Should not match anyone when no names are provided", func() {
		list_of_participants := []string{}
		matches := []string{}

		Expect(secret_santa.GetMatches(list_of_participants)).To(Equal(matches))
	})

	It("Should not match anyone when there is only one participant", func() {
		list_of_participants := []string{"Renea"}
		matches := []string{}

		Expect(secret_santa.GetMatches(list_of_participants)).To(Equal(matches))
	})
})
