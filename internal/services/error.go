package services

type ErrorServiceDTO struct {
	UniqueId string `json:"unique_id"`
	Type     string `json:"type"`
	Code     int    `json:"code"`
	Message  string `json:"error"`
	Detail   string `json:"detail,omitempty"`
}

func (e *ErrorServiceDTO) Error() string {
	return e.Message
}
