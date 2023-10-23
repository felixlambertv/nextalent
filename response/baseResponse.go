package response

type BaseResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
	Data    any    `json:"data,omitempty"`
}
