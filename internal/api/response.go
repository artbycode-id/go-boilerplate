package api

import "artbycode.id/go-app/pkg"

type ResponseBody struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type ResponseErrorBody struct {
	UniqueId string `json:"unique_id"`
	Type     string `json:"type"`
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Detail   string `json:"detail,omitempty"`
}

func errorInternalServerResponseBody(msgDetail error) *ResponseErrorBody {
	return &ResponseErrorBody{
		UniqueId: pkg.GenerateUUID(),
		Code:     0,
		Message:  "internal server error",
		Detail:   msgDetail.Error(),
	}
}
