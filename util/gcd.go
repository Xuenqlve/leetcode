package util

// 最大公约数
// 欧几里得算法 gcd(a,b) = gcd(b, a % b)
// a 和 b 的最大公约数 = b 和 a/b 的余数的 最大公约数
// 当 a % b 为0时，最大公约数为b
// 当 a % b 的余数不为0时，则将a = b ， b = a % b ,递归运算，直到余数为0

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
