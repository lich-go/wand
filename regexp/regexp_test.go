package regexp

import "testing"

func TestMatch(t *testing.T) {
	var b bool
	b = Match("account", "account")
	print("Match Account: ")
	println(b)

	b = Match("account", "_account")
	print("Match Account: ")
	println(b)

	b = Match("email", "@163.vip.com")
	print("Match Email: ")
	println(b)

	b = Match("ipv4", "120.22.15.21")
	print("Match ipv4: ")
	println(b)
}

func TestFind(t *testing.T) {
	var b string
	b = Find("account", "59897asdsd")
	print("Find Account: ")
	println(b)
}
