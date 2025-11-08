package utils

type ResponseJson[T any] struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Result  T      `json:"result"`
	Success bool   `json:"success"`
}

func JsonResponse[T any](error, message string, result T, success bool) ResponseJson[T] {
	return ResponseJson[T]{
		Error:   error,
		Message: message,
		Result:  result,
		Success: success,
	}
}
