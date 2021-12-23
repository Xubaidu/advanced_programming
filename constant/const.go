package constant

const (
	ProductFilePath = "conf/prod.yml"
)

// 状态码常量
const (
	OK              = 0
	ParamsBindError = 40000
	UnknownError    = 50000
	DataQueryError  = 70000
)

// 数据库字段常量
const (
	UserID = "user_id"
	BlogID = "blog_id"
	Limit  = "limit"
)

var ParamMap = map[string]string{
	BlogID: "id",
	UserID: "id",
	Limit:  "limit",
}

var StatusMessageMap = map[int]string{
	OK:              "成功",
	UnknownError:    "未知错误",
	ParamsBindError: "参数绑定错误",
	DataQueryError:  "数据查询错误",
}
