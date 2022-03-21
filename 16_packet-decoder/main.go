package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var puzzle = "005473C9244483004B001F79A9CE75FF9065446725685F1223600542661B7A9F4D001428C01D8C30C61210021F0663043A20042616C75868800BAC9CB59F4BC3A40232680220008542D89B114401886F1EA2DCF16CFE3BE6281060104B00C9994B83C13200AD3C0169B85FA7D3BE0A91356004824A32E6C94803A1D005E6701B2B49D76A1257EC7310C2015E7C0151006E0843F8D000086C4284910A47518CF7DD04380553C2F2D4BFEE67350DE2C9331FEFAFAD24CB282004F328C73F4E8B49C34AF094802B2B004E76762F9D9D8BA500653EEA4016CD802126B72D8F004C5F9975200C924B5065C00686467E58919F960C017F00466BB3B6B4B135D9DB5A5A93C2210050B32A9400A9497D524BEA660084EEA8EF600849E21EFB7C9F07E5C34C014C009067794BCC527794BCC424F12A67DCBC905C01B97BF8DE5ED9F7C865A4051F50024F9B9EAFA93ECE1A49A2C2E20128E4CA30037100042612C6F8B600084C1C8850BC400B8DAA01547197D6370BC8422C4A72051291E2A0803B0E2094D4BB5FDBEF6A0094F3CCC9A0002FD38E1350E7500C01A1006E3CC24884200C46389312C401F8551C63D4CC9D08035293FD6FCAFF1468B0056780A45D0C01498FBED0039925B82CCDCA7F4E20021A692CC012B00440010B8691761E0002190E21244C98EE0B0C0139297660B401A80002150E20A43C1006A0E44582A400C04A81CD994B9A1004BB1625D0648CE440E49DC402D8612BB6C9F5E97A5AC193F589A100505800ABCF5205138BD2EB527EA130008611167331AEA9B8BDCC4752B78165B39DAA1004C906740139EB0148D3CEC80662B801E60041015EE6006801364E007B801C003F1A801880350100BEC002A3000920E0079801CA00500046A800C0A001A73DFE9830059D29B5E8A51865777DCA1A2820040E4C7A49F88028B9F92DF80292E592B6B840"

var qna map[string]int = map[string]int{
	"8A004A801A8002F478":             16,
	"620080001611562C8802118E34":     12,
	"C0015000016115A2E0802F182340":   23,
	"A0016C880162017C3686B18A3D4780": 31,
}

var qna2 map[string]int = map[string]int{
	"C200B40A82":                 3,
	"04005AC33890":               54,
	"880086C3E88112":             7,
	"CE00C43D881120":             9,
	"D8005AC2A8F0":               1,
	"F600BC2D8F":                 0,
	"9C005AC2F8F0":               0,
	"9C0141080250320F1802104A08": 1,
}

var allPackets []string

var opMap map[int]string = map[int]string{
	0: "sum",
	1: "product",
	2: "minmum",
	3: "maximum",
	5: "greater",
	6: "less",
	7: "equal",
}

var hexToBinMap map[rune]string

func readHexToBinMap(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	hexToBinMap = map[rune]string{}
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		var hex rune
		var bin string
		fmt.Sscanf(scan.Text(), "%c = %s", &hex, &bin)
		hexToBinMap[hex] = bin
	}
}

func hexToBinString(hex string) string {
	var bin string
	for _, s := range hex {
		bin += hexToBinMap[s]
	}
	return bin
}

func binToDecimal(bin string) int {
	result := 0
	for i, v := range bin {
		if v == 49 {
			result += int(math.Pow(2, float64(len(bin)-1-i)))
		}
	}
	return result
}

func main() {
	readHexToBinMap("binary.txt")
	for q := range qna2 {
		packet := hexToBinString(q)
		allPackets = []string{}
		getAllPackets(&packet)
	}
}

func calculate(op string, args []int) int {
	result := 0
	switch op {
	case "sum":
		for _, v := range args {
			result += v
		}
	case "product":
		result = 1
		for _, v := range args {
			result *= v
		}
	case "minmum":
		result := 10
		for _, v := range args {
			if v < result {
				result = v
			}
		}
	case "maximum":
		for _, v := range args {
			if v > result {
				result = v
			}
		}
	case "greater":
		if args[0] < args[1] {
			result = 1
		}
	case "less":
		if args[0] > args[1] {
			result = 1
		}
	case "equal":
		if args[0] == args[1] {
			result = 1
		}
	}
	return result
}

func getAllPackets(packet *string) *string {
	tid := binToDecimal((*packet)[3:6])
	fmt.Println(opMap[tid])
	ar := []int{}
	if tid != 4 {
		label := int((*packet)[6] - '0')
		switch label {
		case 0:
			subs := byLength(packet)
			allPackets = append(allPackets, subs...)
		case 1:
			subs := byCount(packet)
			allPackets = append(allPackets, subs...)
		}
		if len(allPackets) > 0 {
			args := []int{}
			for _, p := range allPackets {
				args = append(args, readLiteral(p))
			}
			result := calculate(opMap[tid], args)
			ar = append(ar, result)
			allPackets = []string{}
			fmt.Println(opMap[tid], args, result)
		}
	} else {
		allPackets = append(allPackets, getLiteralPacket(packet))
	}
	return packet
}

func readLiteral(packet string) int {
	num := 0
	packet = packet[6:]
	multiflier := len(packet) - len(packet)/5
	for i := 1; i < len(packet); i++ {
		if i%5 != 0 {
			multiflier -= 1
			num += (int(math.Pow(2, float64(multiflier))) * int(packet[i]-'0'))
		}
	}
	return num
}

func getLiteralPacket(packet *string) string {
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

func byCount(packet *string) []string {
	subs := []string{}
	count := binToDecimal((*packet)[7 : 7+11])
	*packet = (*packet)[7+11:]
	for i := 0; i < count; i++ {
		packet = getAllPackets(packet)
	}
	return subs
}

func byLength(packet *string) []string {
	subs := []string{}
	length := binToDecimal((*packet)[7 : 7+15])

	sub := (*packet)[7+15 : 7+15+length]
	for len(sub) > 0 {
		sub = *getAllPackets(&sub)
	}

	*packet = (*packet)[7+15+length:]
	return subs
}
