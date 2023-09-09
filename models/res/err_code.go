package res

type ErrorCode int

const (
	SettingsError ErrorCode = 1001
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingsError: "系统错误",
	}
)
