package rbac

import (
	"context"

	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/mesh-projects/pkg/api/external/istio/rbac/v1alpha1"
	v1 "github.com/solo-io/mesh-projects/pkg/api/v1"
)

type testOptions struct {
	ctx            context.Context
	cluster1       string
	installationNs string
}

var _ = Describe("Rbac Translation", func() {
	var (
		testOpts *testOptions
	)
	BeforeEach(func() {
		testOpts = &testOptions{
			ctx:            context.TODO(),
			cluster1:       "someCluster1",
			installationNs: "someNs",
		}
	})
	Context("Istio", func() {
		verifyHandleIstioRbac := func(inputMesh, expectedMesh *v1.Mesh, inputRbac, expectedRbac *v1alpha1.RbacConfig) {
			computedMesh, computedRbacConfig := handleIstioRbac(testOpts.ctx, inputMesh, inputRbac)
			ExpectWithOffset(1, computedMesh.String()).To(Equal(expectedMesh.String()))
			ExpectWithOffset(1, computedRbacConfig.String()).To(Equal(expectedRbac.String()))
			ExpectWithOffset(1, computedMesh).To(Equal(expectedMesh))
			ExpectWithOffset(1, computedRbacConfig).To(Equal(expectedRbac))

		}
		Context("inactive mode", func() {
			It("should not add new rbac configs when none exist", func() {
				inputMesh := meshWithRbac(testOpts, &v1.RbacMode{
					Mode: &v1.RbacMode_Unspecified_{Unspecified: &v1.RbacMode_Unspecified{}},
				})
				expectedMesh := meshWithRbac(testOpts, &v1.RbacMode{
					Mode: &v1.RbacMode_Unspecified_{Unspecified: &v1.RbacMode_Unspecified{}},
					Status: &v1.RbacStatus{
						Code:    v1.RbacStatusCode_OK,
						Message: statusInactive,
					},
				})
				verifyHandleIstioRbac(inputMesh, expectedMesh, nil, nil)
			})
			It("should not remove existing rbac configs", func() {
				inputMesh := meshWithRbac(testOpts, &v1.RbacMode{
					Mode: &v1.RbacMode_Unspecified_{Unspecified: &v1.RbacMode_Unspecified{}},
				})
				expectedMesh := meshWithRbac(testOpts, &v1.RbacMode{
					Mode: &v1.RbacMode_Unspecified_{Unspecified: &v1.RbacMode_Unspecified{}},
					Status: &v1.RbacStatus{
						Code:    v1.RbacStatusCode_OK,
						Message: statusInactive,
					},
				})
				inputRbac := &v1alpha1.RbacConfig{
					Status: core.Status{},
					Metadata: core.Metadata{
						Name:      "default",
						Namespace: testOpts.installationNs,
						Cluster:   testOpts.cluster1,
					},
					Mode: v1alpha1.RbacConfig_ON,
				}
				expectedRbac := &v1alpha1.RbacConfig{
					Status: core.Status{},
					Metadata: core.Metadata{
						Name:      "default",
						Namespace: testOpts.installationNs,
						Cluster:   testOpts.cluster1,
					},
					Mode: v1alpha1.RbacConfig_ON,
				}
				verifyHandleIstioRbac(inputMesh, expectedMesh, inputRbac, expectedRbac)
			})
		})
		It("should handle disable mode", func() {
			inputMesh := meshWithRbac(testOpts, &v1.RbacMode{
				Mode: &v1.RbacMode_Disable_{Disable: &v1.RbacMode_Disable{}},
			})
			expectedMesh := meshWithRbac(testOpts, &v1.RbacMode{
				Mode: &v1.RbacMode_Disable_{Disable: &v1.RbacMode_Disable{}},
				Status: &v1.RbacStatus{
					Code:    v1.RbacStatusCode_OK,
					Message: statusDisable,
				},
			})
			expectedRbac := &v1alpha1.RbacConfig{
				Status: core.Status{},
				Metadata: core.Metadata{
					Name:      "default",
					Namespace: testOpts.installationNs,
					Cluster:   testOpts.cluster1,
				},
				Mode: v1alpha1.RbacConfig_OFF,
			}
			verifyHandleIstioRbac(inputMesh, expectedMesh, nil, expectedRbac)
		})
		It("should handle isolate mode", func() {
			inputMesh := meshWithRbac(testOpts, &v1.RbacMode{
				Mode: &v1.RbacMode_Enable_{Enable: &v1.RbacMode_Enable{}},
			})
			expectedMesh := meshWithRbac(testOpts, &v1.RbacMode{
				Mode: &v1.RbacMode_Enable_{Enable: &v1.RbacMode_Enable{}},
				Status: &v1.RbacStatus{
					Code:    v1.RbacStatusCode_OK,
					Message: statusIsolate,
				},
			})
			expectedRbac := &v1alpha1.RbacConfig{
				Status: core.Status{},
				Metadata: core.Metadata{
					Name:      "default",
					Namespace: testOpts.installationNs,
					Cluster:   testOpts.cluster1,
				},
				Mode: v1alpha1.RbacConfig_ON,
			}
			verifyHandleIstioRbac(inputMesh, expectedMesh, nil, expectedRbac)
		})
	})
})

func meshWithRbac(opts *testOptions, rbacCfg *v1.RbacMode) *v1.Mesh {
	return &v1.Mesh{
		MeshType: &v1.Mesh_Istio{Istio: &v1.IstioMesh{
			InstallationNamespace: opts.installationNs,
			Version:               "1",
		}},
		DiscoveryMetadata: &v1.DiscoveryMetadata{
			Cluster: opts.cluster1,
		},
		Rbac: rbacCfg,
	}
}
