package helpers

type ServerResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Body    map[string]interface{} `json:"body"`
}

func NewServerResponse() *ServerResponse {
	return &ServerResponse{}
}
