package spec_test

import (
	//_ "github.com/adamkalnas/secret_santa/spec"
	"github.com/adamkalnas/secret_santa/secret_santa"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SecretSanta", func() {
	It("Should not e-mail anyone when no names are provided", func() {
		list_of_participants := []string{}
		matches := []string{}

		Expect(secret_santa.GetMatches(list_of_participants)).To(Equal(matches))
	})
})
