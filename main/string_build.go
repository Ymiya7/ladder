package main

import "fmt"

func main() {
	var s string
	for {
		n, _ := fmt.Scan(&s)
		if n == 0 {
			break
		} else {

			charCnt := make(map[byte]int)
			var chars []byte
			for i := 0; i < len(s); i++ {
				char := s[i]
				if _, ok := charCnt[char]; !ok {
					chars = append(chars, char)
				}
				charCnt[char]++
			}
			var res []byte
			backTrace(s, charCnt, chars, 0, res)

			fmt.Println(string(res))
		}
	}
}

func backTrace(s string, charCnt map[byte]int, chars []byte, index int, res []byte) {
	if index == len(s) {

		return
	}
	x := s[index]
	for _, char := range chars {
		if x != char && charCnt[char] > 0 {
			res = append(res, char)
			cnt := charCnt[char]
			cnt -= 1
			if cnt == 0 {
				delete(charCnt, char)
			}
			backTrace(s, charCnt, chars, index+1, res)
			charCnt[char]++
			res = res[:len(res)-1]
		}
	}

}
