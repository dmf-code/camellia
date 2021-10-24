package resp

const (
	// status 200
	Ok				= 200000

	// status 400
	BadRequest		= 400000
	ParameterError  = 400001
)


var codeText = map[int]string{
	Ok: "请求成功",
	BadRequest: "失败请求",
	ParameterError: "参数错误",
}

func CodeText(code int) string {
	return codeText[code]
}
