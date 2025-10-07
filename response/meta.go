package response

type MetaData struct {
	Status    string `json:"status"`
	RequestID string `json:"request_id"`
	Code      string `json:"code"`
	Message   string `json:"message"`
}
