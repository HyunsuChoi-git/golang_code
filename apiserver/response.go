package apiserver

type ResponseData struct {
	// 에러 코드 (success, failure)
	Status string `json:"status"`
	// 에러 메시지
	Message string       `json:"message"`
	Data    *interface{} `json:"data,omitempty"`
	Error   *Error       `json:"error,omitempty"`
}

type Error struct {
	// 에러 코드 (에러 코드를 통해 어떤 서비스에서 발생했는지, 어떤 에러가 발생했는지 확인)
	Code int `json:"code,omitempty"`
	//에러 메시지
	Message string `json:"message,omitempty"`
	//에러 메시지 상세
	Description *string `json:"description,omitempty"`
}

func InputDescription(errStruct *Error, description string) *Error {
	errStruct.Description = &description
	return errStruct
}

func SuccessNonDataResponse() ResponseData {
	return ResponseData{
		Status:  "success",
		Message: "success",
	}
}

func SuccessWithDataResponse(data interface{}) ResponseData {
	return ResponseData{
		Status:  "success",
		Message: "success",
		Data:    &data,
	}
}

func FailureResponse(message string, data *interface{}, error *Error) ResponseData {
	if message == "" {
		message = "failure"
	}
	return ResponseData{
		Status:  "failure",
		Message: message,
		Data:    data,
		Error:   error,
	}
}
