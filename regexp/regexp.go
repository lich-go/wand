package regexp

import (
	"regexp"
)

var Rule = make(map[string]string)

func init() {
	Rule["account"] = "^[a-zA-Z][a-zA-Z0-9_]{4,20}$"
	Rule["color_hex"] = "^#([a-fA-F0-9]){3}(([a-fA-F0-9]){3})?$"
	Rule["email"] = "^[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[\\w](?:[\\w-]*[\\w])?"
	Rule["ipv4"] = "^(\\d|[01]?\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d|[01]?\\d\\d|2[0-4] \\d|25[0-5])\\.(\\d|[01]?\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d|[01]?\\d\\d|2[0-4] \\d|25[0-5])$"
	Rule["mac_address"] = "^([0-9a-fA-F]{2}:){5}[0-9a-fA-F]{2}$"
	Rule["md5"] = "(\\S){32}$"
	Rule["qq"] = "^[1-9][0-9]{4,}$"
	Rule["sha1"] = "^(\\S){40}$"
	Rule["url"] = "^(ht|f)tp(s?)\\:\\/\\/[0-9a-zA-Z]([-.\\w]*[0-9a-zA-Z])*(:(0-9)*)*(\\/?)([a-zA-Z0-9\\-\\.\\?\\,\\'\\/\\\\\\+&amp;%$#_]*)?$"

	Rule["cn_cid"] = "^(\\d{6})(\\d{4})(\\d{2})(\\d{2})(\\d{3})([0-9]|X|x)$"
	Rule["cn_phone"] = "^1[3456789]\\d{9}$"
	Rule["cn_postcode"] = "^[1-9]\\d{5}(?!\\d)$"
	Rule["cn_tel"] = "\\d{3}-\\d{8}|\\d{4}-\\d{7,8}"
	Rule["cn_utf8"] = "[\\u4e00-\\u9fa5]"
}

func Match(typ string, str string) bool {
	reg := regexp.MustCompile(Rule[typ])
	match := reg.MatchString(str)
	return match
}

func Find(typ string, str string) string {
	reg := regexp.MustCompile(Rule[typ])
	data := reg.Find([]byte(str))
	return string(data)
}
