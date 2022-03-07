package ocm

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	sdkcfg "github.com/openshift-online/ocm-cli/pkg/config"
	mocks "github.com/openshift/configuration-anomaly-detection/pkg/ocm/mock"
)

var _ = Describe("OCM", func() {
	Describe("When trying to load a configuration", func() {
		var (
			config           *sdkcfg.Config
			mockCtrl         *gomock.Controller
			client           *ocmClient
			mocOCMConnection *mocks.MockocmHandlerIf
			err              error
			configLocation   string
		)
		JustBeforeEach(func() {
			mockCtrl = gomock.NewController(GinkgoT())
			mocOCMConnection = mocks.NewMockocmHandlerIf(mockCtrl)
			client = &ocmClient{comm: mocOCMConnection}

			config, err = newConfigFromFile(configLocation)
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
				// config = sdkcfg.Config{blab:ba}
				// writeToFIle(config, fp)
				// loadedconfig :=loadFile(fp)
				// Except(config).To(Equal(loadedconfig))
				// rm
			})
		})

		When("the client configuration does not exist", func() {
			BeforeEach(func() {
				configLocation = "invalid"
			})
			It("should return an empty configuration", func() {
				Expect(err).Error().Should(HaveOccurred())
				Expect(config).To(Equal(nil))
			})
		})
	})
	Describe("")
})
