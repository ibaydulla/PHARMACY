package models

type ErrorResponse struct{
	Success string `json:"success"`
	Errormassage string `json:"error_msg"`
	ErrorCode string `json:"errorcode"`
}
