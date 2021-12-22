package core

type CodeMapping map[string]string

func (cm CodeMapping) GetCodeInfo(code string) string {

	if v, ok := cm[code]; ok {
		return v
	}

	return ""
}

func (cm CodeMapping) AddCodeInfo(code, msg string) {
	cm[code] = msg
}

const (
	SuccessCode   = "200"
	FailCode      = "500"
	FailJsonParse = "201"
	FailInternal  = "400"
)

var DefaultCodeMapping = CodeMapping{
	SuccessCode:   "success",
	FailCode:      "service temporarily not avaible, please try later",
	FailInternal:  "Bad Request",
	FailJsonParse: "data in wrong format",
}
