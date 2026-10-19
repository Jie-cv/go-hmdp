package response

type Response struct {
	Success bool `json:"success"`
}

type SuccessResponse[T any] struct {
	Response
	Data  T     `json:"data"`
	Total int64 `json:"total"`
}

type CompactSuccessResponse struct {
	Response
}

type ErrorResponse struct {
	Response
	ErrorMsg string `json:"errorMsg"`
}
