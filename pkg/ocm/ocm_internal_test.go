package ocm

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	sdkcfg "github.com/openshift-online/ocm-cli/pkg/config"
	v1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
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

	Describe("When having a valid cluster ID", func() {
		var (
			mockCtrl         *gomock.Controller
			client           *ocmClient
			mocOCMConnection *mocks.MockocmHandlerIf
		)
		BeforeEach(func() {
			mockCtrl = gomock.NewController(GinkgoT())
			mocOCMConnection = mocks.NewMockocmHandlerIf(mockCtrl)
			client = &ocmClient{comm: mocOCMConnection}
		})
		When("getting the cluster deployment with the valid cluster ID", func() {
			It("should return a valid hivev1 deployment", func() {
				// this test fails, because v1.ClusterResourcesGetResponse{} is empty, but we need a valid JSON string in GetClusterDeployment
				// I don't know how to solve this without mocking the getClusterResource() method as well. Any idea?
				mocOCMConnection.EXPECT().OcmGetResourceLive(gomock.Any()).Return(&v1.ClusterResourcesGetResponse{}, nil).Times(1)
				cd, err := client.GetClusterDeployment("valid-cluster-test-id-1")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(cd).Should(Equal(&v1.ClusterResourcesGetResponse{}))
			})
		})
	})
})
