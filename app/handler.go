package main

type Handler interface {
	Execute() Response
}

type ErrorHandler struct {
	request    Request
	error_code int16
}

func (h *ErrorHandler) Execute() Response {
	return Response{corr_id: h.Execute().corr_id, length: 8, err_code: h.error_code}
}
