package param

// Optional
// 获取不定参数中的第一位作为可选参数
// 假如不存在则返回空字符串
func Optional(params []string) string {
	if len(params) > 0 {
		return params[0]
	} else {
		return ""
	}
}
