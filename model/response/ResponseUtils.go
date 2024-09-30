package response

type ResponseUtils struct {
	Code        int         `json:"code"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

// ฟังก์ชันสร้าง Response สำเร็จ
func SuccessResponse(data interface{}) ResponseUtils {
	return ResponseUtils{
		Code:        200,
		Description: "Success",
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

// ฟังก์ชันสร้าง Response เมื่อเกิดข้อผิดพลาด
func ErrorResponse(description string) ResponseUtils {
	return ResponseUtils{
		Code:        500,
		Description: description,
		Data:        nil,
	}
}
