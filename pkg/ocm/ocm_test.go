package ocm

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	sdkcfg "github.com/openshift-online/ocm-cli/pkg/config"
)

var _ = Describe("OCM", func() {
	Describe("When trying to load a configuration", func() {
		var (
			config         *sdkcfg.Config
			configLocation string
		)
		JustBeforeEach(func() {
			config, _ = newConfigFromFile(configLocation)
		})

		When("the client configuration exists", func() {
			BeforeEach(func() {
				configLocation = "../../test/ocm_test.json"
			})
			It("should load the configuration successfully", func() {
				Expect(config).To(Equal(&sdkcfg.Config{
					AccessToken:  "DUMMYVALUE",
					ClientID:     "DUMMYVALUE",
					RefreshToken: "DUMMYVALUE",
					Scopes:       []string{"DUMMYVALUE"},
					TokenURL:     "DUMMYVALUE",
					URL:          "DUMMYVALUE",
				}))
			})
		})

		When("the client configuration does not exist", func() {
			BeforeEach(func() {
				configLocation = "invalid"
			})
			It("should return an empty configuration", func() {
				Expect(config).To(Equal(&sdkcfg.Config{}))
			})
		})
	})
})
