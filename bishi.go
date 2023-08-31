package main

import "fmt"

func main1() {
	a := "aaabbaaac"
	fmt.Println(getR(a))

}
func getR(str string) float64 {
	count := 0
	totalLength := 0
	currLength := 0

	for i := 0; i < len(str); i++ {
		if i == 0 || str[i] == str[i-1] {
			currLength++
		} else {
			totalLength += currLength
			count++
			currLength = 1
		}
	}
	totalLength += currLength
	count++
	avgLength := float64(totalLength) / float64(count)
	return avgLength
}

//func isP(num int) bool {
//	reversed := 0
//	origin := num
//	for num > 0 {
//		remainder := num % 10
//		reversed = reversed*10 + remainder
//		num /= 10
//	}
//	return origin == reversed
//
//}
//func generatePa(n int) []int {
//	p := []int{}
//	for i := 0; i <= n; i++ {
//		if isP(i) {
//			p = append(p, i)
//		}
//
//	}
//	return p
//}
