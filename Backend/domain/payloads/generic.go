package payloads

type GenericMessageResponse struct {
	Message string `json:"message"`
}

type GenericsSuccessFlagResponse struct {
	Successful bool   `json:"success"`
	Message    string `json:"message,omitempty"`
}
