package e

import "chatgpt_x/pkg/logger"

var MsgFlags = map[int]string{
	SUCCESS:       "success",
	ERROR:         "系统内部未知错误",
	InvalidParams: "请求参数错误",
	ErrorAuthFail: "认证失败或权限不足",

	ErrorUserIsExist:                 "用户已经存在",
	ErrorUserCreateFail:              "用户创建失败",
	ErrorIncorrectUsernameOrPassword: "用户名或密码错误",
	ErrorGenerateTokenFail:           "生成token失败",
	ErrorUserIsDisabled:              "用户已被封禁",
	ErrorUserDeleteFail:              "用户删除失败",
	ErrorAiModelIsExist:              "AI模型已经存在",
	ErrorAiModelMapCreateFail:        "AI模型关系映射创建失败",
	ErrorAiModelMapUpdateFail:        "AI模型关系映射更新失败",
	ErrorAiModelMapSelectFail:        "AI模型关系映射查询失败",
	ErrorAiModelMapDeleteFail:        "AI模型关系映射删除失败",
}

// GetMsg get error information based on Code.
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}

// HasError any errors will be saved to the log.
func HasError(err error) bool {
	if err != nil {
		logger.Error(err)
		return true
	}
	return false
}
