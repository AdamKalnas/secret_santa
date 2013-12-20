package spec_test

import (
	"github.com/adamkalnas/secret_santa/secret_santa"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SecretSanta", func() {
	It("Should not match anyone when no names are provided", func() {
		list_of_participants := []string{}
		matches := make(map[string]string)

		Expect(secret_santa.GetMatches(list_of_participants)).To(Equal(matches))
	})

	It("Should not match anyone when there is only one participant", func() {
		list_of_participants := []string{"Renea"}
		matches := make(map[string]string)

		Expect(secret_santa.GetMatches(list_of_participants)).To(Equal(matches))
	})

	It("Should match two people to eachother when there are only two participants", func() {
		list_of_participants := []string{"Renea", "Rachel"}
		matches := make(map[string]string)
		matches["Renea"] = "Rachel"
		matches["Rachel"] = "Renea"

		Expect(secret_santa.GetMatches(list_of_participants)).To(Equal(matches))
	})

	It("Should match three participants to recipients other than themselves", func() {
		list_of_participants := []string{"Renea", "Rachel", "Ryan"}
		results := secret_santa.GetMatches(list_of_participants)
		Expect(len(results)).To(Equal(3))
		for _, participant := range list_of_participants {
			match, found := results[participant]
			Expect(match).ToNot(Equal(participant))
			Expect(found).To(Equal(true))
		}
	})
})
