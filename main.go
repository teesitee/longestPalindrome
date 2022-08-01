package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "fffff"
	var palin []string
	var truePalin []string
	var finalPalin []string
	if palin == nil {
		var countLetter int
		for i := range s {
			letter := s[0]
			if letter == s[i] {
				countLetter += 1
			} else {
				break
			}
		}
		if countLetter == len(s) {
			fmt.Println("palindrome :", string(s))
			return
		}
	}
	if len(s) == 1 {
		palin = append(palin, s)
		fmt.Println(palin)
	} else if len(s) == 0 {
		palin = append(palin, "")
		fmt.Println(palin)
	}
	if len(s) != 1 || len(s) != 0 {
		for i := 0; i < (len(s) - 1); i++ {
			fLetter := string(s[i])
			var sLetter string
			for j := i + 1; j <= (len(s) - 1); j++ {
				sLetter = string(s[j])
				if fLetter == sLetter {
					palin = append(palin, string(s[i:j+1]))
				}
			}
		}
		for i := range palin {
			var reversePal []string
			pal := palin[i]
			for j := range pal {
				reversePal = append(reversePal, string(pal[len(pal)-1-j]))
			}
			// fmt.Println("reversePal :", reversePal)
			joinPal := strings.Join(reversePal, "")

			if pal == joinPal {
				truePalin = append(truePalin, pal)
				// fmt.Println("focus", truePalin)
			}
			// fmt.Println("truePalin step", i, truePalin)
		}
		// fmt.Println("truePalin :", truePalin)
		checkPalin := make(map[string]bool)
		for _, item := range truePalin {
			if _, value := checkPalin[item]; !value {
				checkPalin[item] = true /*ใส่ true ก็ได้ false ก็ได้ เพราะแค่เช็คว่ามีค่า bool หรือไม่*/
				finalPalin = append(finalPalin, item)
			}
		}
		// fmt.Println("finalPalin :", finalPalin)
	}
	if len(truePalin) == 0 {
		fmt.Println("palindrome :", string(s[0]))
	} else {
		fmt.Println("final palindrome :", finalPalin)
	}
}
