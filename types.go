package main

type AdmissionReview struct {
	Request *AdmissionRequest `json:"request"`
}

type AdmissionRequest struct {
	UID    string `json:"uid"`
	Object Object `json:"object"`
}

type Object struct {
	Spec Spec `json:"spec"`
}

type Spec struct {
	Containers     []Container `json:"containers"`
	InitContainers []Container `json:"initContainers"`
}

type Container struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

type PatchOperation struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value string `json:"value"`
}

type AdmissionResponse struct {
	Response struct {
		UID       string           `json:"uid"`
		Allowed   bool             `json:"allowed"`
		Patch     []PatchOperation `json:"patch"`
		PatchType string           `json:"patchType"`
	} `json:"response"`
}
