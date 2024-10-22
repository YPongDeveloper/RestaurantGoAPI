package response

type ResponseUtils struct {
	Code        int         `json:"code"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

func SuccessResponse(data interface{}) ResponseUtils {
	return ResponseUtils{
		Code:        200,
		Description: "Success",
		Data:        data,
	}
}
func SuccessInQueue(data interface{}) ResponseUtils {
	return ResponseUtils{
		Code:        201,
		Description: "Success In Queue",
		Data:        data,
	}
}
func SuccessNoResponse() ResponseUtils {
	return ResponseUtils{
		Code:        200,
		Description: "Success",
		Data:        "{}",
	}
}

func ErrorResponse(description string) ResponseUtils {
	return ResponseUtils{
		Code:        500,
		Description: description,
		Data:        nil,
	}
}
