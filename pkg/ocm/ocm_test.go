package ocm_test

import (
	"errors"
	"fmt"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	awsv1alpha1 "github.com/openshift/aws-account-operator/pkg/apis/aws/v1alpha1"
	hivev1 "github.com/openshift/hive/apis/hive/v1"
	"k8s.io/apimachinery/pkg/util/json"

	"github.com/openshift/configuration-anomaly-detection/pkg/ocm"
	mocks "github.com/openshift/configuration-anomaly-detection/pkg/ocm/mock"
)

var _ = Describe("OCM", func() {
	var (
		mockCtrl          *gomock.Controller
		client            *ocm.OcmClient
		mocOCMConnection  *mocks.MockocmHandlerIf
		clusterDeployment *hivev1.ClusterDeployment
		awsAccountClaim   *awsv1alpha1.AccountClaim
		err               error
		clustername       string
		supportRoleArn    string
	)
	Describe("When fetching a ClusterDeployment", func() {
		BeforeEach(func() {
			mockCtrl = gomock.NewController(GinkgoT())
			mocOCMConnection = mocks.NewMockocmHandlerIf(mockCtrl)
			client = &ocm.OcmClient{Comm: mocOCMConnection}
			err = fmt.Errorf("some error")
			clustername = "test-cluster"
			clusterDeployment = &hivev1.ClusterDeployment{Spec: hivev1.ClusterDeploymentSpec{ClusterName: clustername}}
		})
		When("cluster ID is valid", func() {
			It("should return a valid ClusterDeployment", func() {
				cdstring, _ := json.Marshal(clusterDeployment)
				mocOCMConnection.EXPECT().OcmGetResourceLive(clustername, "cluster_deployment").
					Return(string(cdstring), nil).Times(1)

				cd_out, err := client.GetClusterDeployment(clustername)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(clusterDeployment).Should(Equal(cd_out))
			})
		})
		When("the cluster doesn't exist or any other error happends in the sdk", func() {
			It("wrapps the error message and returns nil", func() {
				mocOCMConnection.EXPECT().OcmGetResourceLive(clustername, "cluster_deployment").
					Return("", err).Times(1)

				cd_out, err_out := client.GetClusterDeployment(clustername)
				Expect(err_out).Should(HaveOccurred())
				Expect(errors.Unwrap(err_out)).Should(Equal(err))
				Expect(cd_out).Should(BeNil())
			})
		})
		When("the cluster exists but the specified resource doesn not exist", func() {
			It("will return an error", func() {
				mocOCMConnection.EXPECT().OcmGetResourceLive(clustername, "cluster_deployment").
					Return("", nil).Times(1)

				cd_out, err_out := client.GetClusterDeployment(clustername)
				Expect(err_out).Should(HaveOccurred())
				Expect(cd_out).Should(BeNil())
			})
		})
	})
	Describe("When fetching AWSAccountClaim", func() {
		BeforeEach(func() {
			mockCtrl = gomock.NewController(GinkgoT())
			mocOCMConnection = mocks.NewMockocmHandlerIf(mockCtrl)
			client = &ocm.OcmClient{Comm: mocOCMConnection}
			err = fmt.Errorf("some error")
			clustername = "test-cluster"
			awsAccountClaim = &awsv1alpha1.AccountClaim{Spec: awsv1alpha1.AccountClaimSpec{STSRoleARN: "support-role"}}
		})
		When("cluster ID is valid", func() {
			It("should return a valid AWSAccountClaim", func() {
				cdstring, _ := json.Marshal(awsAccountClaim)
				mocOCMConnection.EXPECT().OcmGetResourceLive(clustername, "aws_account_claim").
					Return(string(cdstring), nil).Times(1)

				ac_out, err := client.GetAWSAccountClaim(clustername)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(awsAccountClaim).Should(Equal(ac_out))
			})
		})
		When("the cluster doesn't exist or any other error happends in the sdk", func() {
			It("wrapps the error and returns nil", func() {
				mocOCMConnection.EXPECT().OcmGetResourceLive(clustername, "aws_account_claim").
					Return("", err).Times(1)

				ac_out, err_out := client.GetAWSAccountClaim(clustername)
				Expect(err_out).Should(HaveOccurred())
				Expect(errors.Unwrap(err_out)).Should(Equal(err))
				Expect(ac_out).Should(BeNil())
			})
		})
		When("the cluster exists but the specified resource doesn not exist", func() {
			It("will return an error", func() {
				mocOCMConnection.EXPECT().OcmGetResourceLive(clustername, "aws_account_claim").
					Return("", nil).Times(1)

				ac_out, err_out := client.GetAWSAccountClaim(clustername)
				Expect(err_out).Should(HaveOccurred())
				Expect(ac_out).Should(BeNil())
			})
		})
	})
	Describe("When fetching SupportRoleArn", func() {
		BeforeEach(func() {
			mockCtrl = gomock.NewController(GinkgoT())
			mocOCMConnection = mocks.NewMockocmHandlerIf(mockCtrl)
			client = &ocm.OcmClient{Comm: mocOCMConnection}
			err = fmt.Errorf("some error")
			clustername = "test-cluster"
			supportRoleArn = "support-role"
			awsAccountClaim = &awsv1alpha1.AccountClaim{Spec: awsv1alpha1.AccountClaimSpec{SupportRoleARN: supportRoleArn}}
		})
		When("cluster ID is valid", func() {
			It("should return a valid ARN", func() {
				acstring, _ := json.Marshal(awsAccountClaim)
				mocOCMConnection.EXPECT().OcmGetResourceLive(clustername, "aws_account_claim").
					Return(string(acstring), nil).Times(1)

				arn_out, err := client.GetSupportRoleARN(clustername)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(supportRoleArn).Should(Equal(arn_out))
			})
		})
		When("cluster ID is valid but the ARN is not specified", func() {
			It("should return an error", func() {
				awsAccountClaim.Spec.SupportRoleARN = ""
				acstring, _ := json.Marshal(awsAccountClaim)
				mocOCMConnection.EXPECT().OcmGetResourceLive(clustername, "aws_account_claim").
					Return(string(acstring), nil).Times(1)

				arn_out, err := client.GetSupportRoleARN(clustername)
				Expect(err).Should(HaveOccurred())
				Expect(arn_out).Should(Equal(""))
			})
		})
		When("the cluster doesn't exist or any other error happends in the sdk", func() {
			It("wrapps the error and returns empty string", func() {
				mocOCMConnection.EXPECT().OcmGetResourceLive(clustername, "aws_account_claim").
					Return("", err).Times(1)

				arn_out, err_out := client.GetSupportRoleARN(clustername)
				Expect(err_out).Should(HaveOccurred())
				Expect(arn_out).Should(Equal(""))
			})
		})
		When("the cluster exists but the specified resource doesn not exist", func() {
			It("will return an error", func() {
				mocOCMConnection.EXPECT().OcmGetResourceLive(clustername, "aws_account_claim").
					Return("", nil).Times(1)

				arn_out, err_out := client.GetSupportRoleARN(clustername)
				Expect(err_out).Should(HaveOccurred())
				Expect(arn_out).Should(Equal(""))
			})
		})
	})
})
