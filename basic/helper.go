package basic

import "strings"

var num2char = "0123456789abcdefghijklmnopqrstuvwxyz"

func NumToBHex(num, n int) string {
	num_str := ""
	for num != 0 {
		yu := num % n
		num_str = string(num2char[yu]) + num_str
		num = num / n
	}
	return strings.ToUpper(num_str)
}
