package main

import "fmt"

func main() {
	a := "Here's my spammy page: http://hehefouls.netHAHAHA see you."
	b := MaskingUrl(a)

	fmt.Println(a)
	fmt.Println(b)
}

func MaskingUrl(str string) string {
	var startUrlIndex, isMasking = 0, false
	buffer := []byte(str)

	for i := range buffer {
		if buffer[i] == 'h' && string(buffer[i:i+7]) == "http://" {
			startUrlIndex = i + 7
			isMasking = true
		}

		if startUrlIndex != 0 && i >= startUrlIndex && isMasking == true {
			if buffer[i] == ' ' {
				isMasking = false
			} else {
				buffer[i] = '*'
			}
		}
	}

	return string(buffer)
}
