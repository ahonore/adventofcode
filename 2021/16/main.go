package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

func f1ReadPacket(bits []int) (int, int) {
	packetVersion := (bits[0] << 2) | (bits[1] << 1) | bits[2]
	typeID := (bits[3] << 2) | (bits[4] << 1) | bits[5]
	id := 6
	if typeID == 4 {
		for {
			flag := bits[id]
			id += 5
			if flag == 0 {
				break
			}
		}
		return packetVersion, id
	}

	lenTypeID := bits[id]
	id++
	if lenTypeID == 1 {
		// 11 bits = num of subpackets
		num := 0
		for i := 0; i < 11; i++ {
			num <<= 1
			num |= bits[id+i]
		}
		id += 11
		sumVersion := 0
		for i := 0; i < num; i++ {
			subVer, read := f1ReadPacket(bits[id:])
			id += read
			sumVersion += subVer
		}
		return packetVersion + sumVersion, id
	}

	// 15 bits = len of all subpackets
	len := 0
	for i := 0; i < 15; i++ {
		len <<= 1
		len |= bits[id+i]
	}
	id += 15
	idMax := id + len
	sumVersion := 0
	for id < idMax {
		subVers, read := f1ReadPacket(bits[id:])
		sumVersion += subVers
		id += read
	}
	return packetVersion + sumVersion, id
}

func f1(r io.Reader) int {
	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	// conv hex to bits array (arry of bytes with one significant bit per int)
	bits := []int{}
	for _, h := range b {
		i, _ := strconv.ParseInt(string(h), 16, 8)
		bits = append(bits, int(i>>3), int(i>>2)&1, int(i>>1)&1, int(i&1))
	}
	count, _ := f1ReadPacket(bits)
	return count
}

func f2ReadPacket(bits []int) (int, int) {
	typeID := (bits[3] << 2) | (bits[4] << 1) | bits[5]
	id := 6
	if typeID == 4 {
		// literal value
		value := 0
		for {
			flag := bits[id]
			value <<= 4
			value |= (bits[id+1] << 3) | (bits[id+2] << 2) | (bits[id+3] << 1) | bits[id+4]
			id += 5
			if flag == 0 {
				break
			}
		}
		return value, id
	}

	lenTypeID := bits[id]
	id++
	values := []int{}
	if lenTypeID == 1 {
		// 11 bits = num of subpackets
		num := 0
		for i := 0; i < 11; i++ {
			num <<= 1
			num |= bits[id+i]
		}
		id += 11
		for i := 0; i < num; i++ {
			v, read := f2ReadPacket(bits[id:])
			id += read
			values = append(values, v)
		}
	} else {
		// 15 bits = len of all subpackets
		len := 0
		for i := 0; i < 15; i++ {
			len <<= 1
			len |= bits[id+i]
		}
		id += 15
		idMax := id + len
		for id < idMax {
			v, read := f2ReadPacket(bits[id:])
			id += read
			values = append(values, v)
		}
	}

	res := 0
	switch typeID {
	case 0: // sum
		for _, v := range values {
			res += v
		}
	case 1: // product
		res = 1
		for _, v := range values {
			res *= v
		}
	case 2: // min
		res = math.MaxInt
		for _, v := range values {
			if v < res {
				res = v
			}
		}
	case 3: //max
		res = math.MinInt
		for _, v := range values {
			if v > res {
				res = v
			}
		}
	case 5: // greater than
		if values[0] > values[1] {
			res = 1
		}
	case 6: // less than
		if values[0] < values[1] {
			res = 1
		}
	case 7: // equal to
		if values[0] == values[1] {
			res = 1
		}
	}
	return res, id
}

func f2(r io.Reader) int {
	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	// conv hex to bits array (arry of bytes with one significant bit per int)
	bits := []int{}
	for _, h := range b {
		i, _ := strconv.ParseInt(string(h), 16, 8)
		bits = append(bits, int(i>>3), int(i>>2)&1, int(i>>1)&1, int(i&1))
	}
	count, _ := f2ReadPacket(bits)
	return count
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(f1(bytes.NewReader(b)))
	fmt.Println(f2(bytes.NewReader(b)))
}
