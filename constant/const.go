package constant

// 文件常量
const (
	FileSizeLimit   = 30 * 1024 * 1024
	ProductFilePath = "conf/prod.yml"
	SavedFilesPath  = "saved_files/"
)

// qq 邮箱常量
const (
	SMTPHost     = "smtp.qq.com"
	SMTPPort     = 465
	SMTPUsername = "1127862434@qq.com"
	SMTPPassword = "ofeobzafgzatbaeg"

	MailGreetingSubject = "欢迎注册，我的伙伴！"
	MailGreetingBody    = "<h1>新年快乐</h1>"
	MailContentType     = "text/html"
)

// 飞书 openapi 常量
const (
	UploadFileToLarkURL     = "https://open.feishu.cn/open-apis/drive/v1/files/upload_all"
	DownloadFileFromLarkURL = "https://open.feishu.cn/open-apis/drive/v1/files/%s/download"
	LarkUserToken           = 1
	LarkRefreshToken        = 2
	LarkTenantToken         = 3
	LarkAppToken            = 4
)

// 状态码常量
const (
	OK = 0

	ParamsBindError     = 40000
	ParamsParseError    = 40001
	ParamsValidateError = 40002

	UnknownError = 50000

	DataQueryError         = 70000
	UserDuplicateError     = 70001
	UserMissingError       = 70002
	PasswordWrongError     = 70003
	RecordNotFound         = 70004
	RecordInsertError      = 70005
	RecordUpdateError      = 70006
	RecordDeleteError      = 70007
	TransactionBeginError  = 70008
	TransactionCommitError = 70009
	DuplicateApplyError    = 70010

	FileCloseError      = 80000
	FileUploadError     = 80001
	SaveUploadFileError = 80002

	MailSendError = 90000
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
	OK:                 "成功",
	UnknownError:       "未知错误",
	ParamsBindError:    "参数绑定错误",
	DataQueryError:     "数据查询错误",
	UserDuplicateError: "用户名重复错误",
	UserMissingError:   "用户不存在",
	PasswordWrongError: "密码错误",
}
