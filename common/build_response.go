package common

import (
	"advanced_programming/constant"
)

// BuildResp 根据传入的状态码和信息构造 http 响应
func BuildResp(statusCode int, message string, data interface{}) interface{} {

	// 如果信息为空，获取状态码对应的信息
	if message == "" {
		message, _ = constant.StatusMessageMap[statusCode]
	}

	// 如果数据为空，构造一个 empty json: map[string]interface{}
	if data == nil {
		data = make(map[string]interface{})
	}

	// 构造响应体
	ret := map[string]interface{}{
		"msg":  message,
		"code": statusCode,
		"data": data,
	}
	return ret
}

// BuildRespByErr 根据错误构造 http 响应
func BuildRespByErr(err error) interface{} {

	// 如果是自定义错误，返回对应的信息
	if err, ok := err.(*CustomError); ok {
		if err.Error() != "" {
			return BuildResp(err.GetCode(), err.Error(), nil)
		}
		return BuildResp(err.GetCode(), constant.StatusMessageMap[err.GetCode()], nil)
	}

	// 否则返回统一的内部错误信息
	return BuildResp(constant.UnknownError, "系统繁忙请稍后再试", nil)
}

// BuildRespByCodeMsg 根据 code 和 msg 构造 http 响应
func BuildRespByCodeMsg(statusCode int, message string) interface{} {
	return BuildResp(statusCode, message, nil)
}
