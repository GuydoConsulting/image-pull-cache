package main

import (
	"strconv"
)

func mutate(review AdmissionReview) AdmissionResponse {
	response := AdmissionResponse{
		Response: struct {
			UID       string           `json:"uid"`
			Allowed   bool             `json:"allowed"`
			Patch     []PatchOperation `json:"patch"`
			PatchType string           `json:"patchType"`
		}{
			UID:       review.Request.UID,
			Allowed:   true,
			PatchType: "JSONPatch",
		},
	}

	prefix := config.PullCacheURL // Retrieve the environment variable
	for _, ct := range []string{"containers", "initContainers"} {
		containers := getContainers(review, ct)
		log.Debugf("processing %s", ct)
		for i, container := range containers {
			log.Debugf("processing %s against prefix %s", container.Image, prefix)
			if !startsWith(container.Image, prefix) {
				patch := PatchOperation{
					Op:    "replace",
					Path:  "/spec/" + ct + "/" + strconv.Itoa(i) + "/image",
					Value: prefix + "/" + container.Image,
				}
				response.Response.Patch = append(response.Response.Patch, patch)
			}
		}
	}

	return response
}

// getContainers returns the slice of containers based on the type ("containers" or "initContainers")
func getContainers(review AdmissionReview, containerType string) []Container {
	if containerType == "containers" {
		return review.Request.Object.Spec.Containers
	} else if containerType == "initContainers" {
		return review.Request.Object.Spec.InitContainers
	}
	return []Container{} // Return an empty slice if the container type is unknown
}
