package ocm

import (
	"fmt"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	sdkcfg "github.com/openshift-online/ocm-cli/pkg/config"
	mocks "github.com/openshift/configuration-anomaly-detection/pkg/ocm/mock"
)

var _ = Describe("OCM", func() {
	Describe("When trying to load a configuration", func() {
		var (
			config         *sdkcfg.Config
			configLocation string
		)
		// JustBeforeEach executes before each of the following "when" statements.
		// This means: We run for each when statement "newConfigFromFile(configLocation)"
		// with a different configLocation as described in the BeforeEach() statements
		// in each "when" statement.
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

	Describe("When initializing a new OCM client", func() {
		var (
			mockCtrl         *gomock.Controller
			client           *ocmClient
			mocOCMConnection *mocks.MockocmHandlerIf
		)
		BeforeEach(func() {
			mockCtrl = gomock.NewController(GinkgoT())
			mocOCMConnection = mocks.NewMockocmHandlerIf(mockCtrl)
			client = &ocmClient{comm: mocOCMConnection}
			fmt.Println(client)
		})
	})
})
