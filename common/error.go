package common

// 自定义 error 类型
// 只需要实现 Error() string 方法就算实现了 error 接口

type CustomError struct {
	Code int
	Msg  string
}

func (err *CustomError) Error() string {
	return err.Msg
}

func (err *CustomError) GetCode() int {
	return err.Code
}

func NewError(code int, errMsg string) error {
	return &CustomError{
		Code: code,
		Msg:  errMsg,
	}
}
