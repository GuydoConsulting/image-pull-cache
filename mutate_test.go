package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestAdmissionController(t *testing.T) {
	os.Setenv("PULL_CACHE_URL", "harbor.core.domain")
	config = newConfig()
	log.SetLevel(logrus.DebugLevel)
	testCases := []struct {
		Name              string
		AdmissionReview   AdmissionReview
		AdmissionResponse AdmissionResponse
	}{
		{
			Name: "Single container, Single initContainer",
			AdmissionReview: AdmissionReview{
				Request: &AdmissionRequest{
					UID: "12345",
					Object: Object{
						Spec: Spec{
							Containers: []Container{
								{Name: "main-container", Image: "nginx:latest"},
							},
							InitContainers: []Container{
								{Name: "init-container", Image: "k8s.gcr.io/test"},
							},
						},
					},
				},
			},
		},
		{
			Name: "Two containers, Two initContainers",
			AdmissionReview: AdmissionReview{
				Request: &AdmissionRequest{
					UID: "67890",
					Object: Object{
						Spec: Spec{
							Containers: []Container{
								{Name: "container-1", Image: "docker.io/library/nginx:latest"},
								{Name: "container-2", Image: "public.ecr.aws/test:latest"},
							},
							InitContainers: []Container{
								{Name: "init-container-1", Image: "busybox:latest"},
								{Name: "init-container-2", Image: "nginx:latest"},
							},
						},
					},
				},
			},
		},
		{
			Name: "Two containers, No initContainers",
			AdmissionReview: AdmissionReview{
				Request: &AdmissionRequest{
					UID: "67890",
					Object: Object{
						Spec: Spec{
							Containers: []Container{
								{Name: "container-1", Image: "docker.io/library/nginx:latest"},
								{Name: "container-2", Image: "public.ecr.aws/test:latest"},
							},
						},
					},
				},
			},
		},
		{
			Name: "No containers, Two initContainers",
			AdmissionReview: AdmissionReview{
				Request: &AdmissionRequest{
					UID: "67890",
					Object: Object{
						Spec: Spec{
							InitContainers: []Container{
								{Name: "container-1", Image: "docker.io/library/nginx:latest"},
								{Name: "container-2", Image: "public.ecr.aws/test:latest"},
							},
						},
					},
				},
			},
		},
		{
			Name: "Nothing",
			AdmissionReview: AdmissionReview{
				Request: &AdmissionRequest{
					UID: "67890",
					Object: Object{
						Spec: Spec{},
					},
				},
			},
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			tc := tc
			t.Parallel() // Run sub-tests in parallel if possible
			// Mock your admission controller logic using tc.AdmissionReview
			// For demonstration, we'll just print the admission request
			fmt.Printf("Received Admission Review: %+v\n", tc.AdmissionReview.Request)
			// Call your admission controller function with tc.AdmissionReview
			fmt.Printf("Received Admission Response: %+v\n", mutate(tc.AdmissionReview))
			// Assert the expected behavior
		})
	}
	os.Unsetenv("PULL_CACHE_URL")
}
