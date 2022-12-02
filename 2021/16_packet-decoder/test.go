package main

import (
	"fmt"
	"math"
)

func main() {
	// packet := "11010001111"
	packet := "1101001000101111"
	fmt.Println(readLiteral(packet))
}

func getLiteralPacket(packet *string) string {
	fmt.Println("lit: ", *packet)
	lit := (*packet)[:6]
	for i := 6; i < len(*packet); i += 5 {
		lit += (*packet)[i : i+5]
		if (*packet)[i] == '0' {
			break
		}
	}
	*packet = (*packet)[len(lit):]
	return string(lit)
}

func readLiteral(packet string) int {
	num := 0
	packet = packet[6:]
	multiflier := len(packet) - len(packet)/5
	fmt.Println(packet)
	for i := 1; i < len(packet); i++ {
		if i%5 != 0 {
			multiflier -= 1
			num += (int(math.Pow(2, float64(multiflier))) * int(packet[i]-'0'))
		}
	}
	return num
}
