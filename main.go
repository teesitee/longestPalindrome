package main

import "fmt"

func main() {
	s := "babad"
	var palin []string
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
		fmt.Println("first palindrome :", string(s[0]))
	} else {
		fmt.Println("palindrome total :", palin, "\nfirst palindrome :", palin[0])
	}
}
