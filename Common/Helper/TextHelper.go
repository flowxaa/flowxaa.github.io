package Helper

import (
	"regexp"
	"strings"
)

func GetLetterString(str string) (result string) {
	re := regexp.MustCompile(`[0-9]`)
	result = re.ReplaceAllString(str, "")
	return result
}

// FilterSpecial 过滤特殊字符
//如果字符串为空，直接返回。
func FilterSpecial(strHtml string) (result string) {
	if strHtml == "" {
		return strHtml
	}
	aryReg := []string{"'", "?", "<", ">", "*", "%", "$", "#", ">=", "=<", ";", "||", "[", "]", "&", "|", "''"}
	for _, k := range aryReg {
		strHtml = strings.Replace(strHtml, k, "", -1)
	}
	return strHtml
}
