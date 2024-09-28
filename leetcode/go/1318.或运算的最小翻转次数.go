package main

func main() {

}

func minFlips(a int, b int, c int) int {
	// bitA,bitB,bitC
	// 如果bitC == 1, bitA 和 bitB 有一个为1则不需要翻转，否则需要1次
	// 如果bitC == 0，bitA 和 bitB 都必须为0 ，需要翻转bitA + bitB次

	ans := 0

	// 题设 a,b,c < 10 ^9
	for i := 0; i < 31; i++ {
		bitA := (a >> i) & 1
		bitB := (b >> i) & 1
		bitC := (c >> i) & 1

		if bitC == 0 {
			ans += bitA + bitB
		} else {
			if bitA == 0 && bitB == 0 {
				ans += 1
			}

		}
	}
	return ans
}
