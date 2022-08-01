package main

import "fmt"

func main() {
	s := "xxxxxxxxx"
	var palin []string
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
	} else {
		for i := 0; i < (len(s) - 1); i++ {
			fLetter := string(s[i])
			var sLetter string
			for j := i + 1; j <= (len(s) - 1); j++ {
				sLetter = string(s[j])
				if fLetter == sLetter {
					palin = append(palin, string(s[i:j+1]))
					break
				}

			}

		}

	}
	if len(palin) == 0 {
		fmt.Println("palindrome :", string(s[0]))
	} else {
		fmt.Println("palindrome total :", palin, "\npalindrome :", palin[0])
	}
}
