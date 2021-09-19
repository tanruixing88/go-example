package main

import "fmt"

func multiply(num1 string, num2 string) string {
	if len(num1) == 0 || len(num2) == 0 {
		return ""
	}

	if num1 == "0" || num2 == "0" {
		return "0"
	}

	num3 := make([]byte, len(num1)+len(num2)+1)
	for i := 0; i < len(num3); i++ {
		num3[i] = '0'
	}

	for j := len(num2)-1; j >= 0; j-- {
		for i := len(num1)-1; i >= 0; i-- {
			v := int(num1[i] - '0') * int(num2[j] - '0') + int(num3[i+j+2] - '0')
			//fmt.Printf("i:%d j:%d v:%d \r\n", i, j, v)
			//len(num3)-1-(len(num2)-1-j)-(len(num1)-1-i)=len(num1)+len(num2)+1-1-len(num2)+1+j-len(num1)+1+i
			num3[i+j+2] = byte(v % 10) + '0'
			num3[i+j+1] += byte(v / 10)
		}
	}

	i := 0
	for i < len(num3) {
		if num3[i] > '0' {
			break
		}
		i++
	}

	return string(num3[i:])
}

func main() {
	r := multiply("12", "3")
	fmt.Printf("multiply:%s\r\n", r)
	r = multiply("123", "456")
	fmt.Printf("multiply:%s\r\n", r)
}
