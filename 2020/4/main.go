package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// byr (Birth Year)
// iyr (Issue Year)
// eyr (Expiration Year)
// hgt (Height)
// hcl (Hair Color)
// ecl (Eye Color)
// pid (Passport ID)
// cid (Country ID)

func genPassportKeysMap() map[string]struct{} {
	return map[string]struct{}{
		"byr": {},
		"iyr": {},
		"eyr": {},
		"hgt": {},
		"hcl": {},
		"ecl": {},
		"pid": {},
		"cid": {},
	}
}

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
// hgt (Height) - a number followed by either cm or in:
// 		If cm, the number must be at least 150 and at most 193.
// 		If in, the number must be at least 59 and at most 76.
// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
// pid (Passport ID) - a nine-digit number, including leading zeroes.
// cid (Country ID) - ignored, missing or not.

func isKeyValid(key, value string) bool {
	switch key {
	case "byr":
		re := regexp.MustCompile(`^[0-9]{4}$`)
		if !re.MatchString(value) {
			return false
		}

		i, _ := strconv.Atoi(value)
		return i >= 1920 && i <= 2002
	case "iyr":
		re := regexp.MustCompile(`^[0-9]{4}$`)
		if !re.MatchString(value) {
			return false
		}

		i, _ := strconv.Atoi(value)
		return i >= 2010 && i <= 2020
	case "eyr":
		re := regexp.MustCompile(`^[0-9]{4}$`)
		if !re.MatchString(value) {
			return false
		}

		i, _ := strconv.Atoi(value)
		return i >= 2020 && i <= 2030
	case "hgt":
		re := regexp.MustCompile(`^[0-9]+(in|cm)$`)
		if !re.MatchString(value) {
			return false
		}

		i, _ := strconv.Atoi(value[:len(value)-2])
		switch value[len(value)-2:] {
		case "cm":
			return i >= 150 && i <= 193
		case "in":
			return i >= 59 && i <= 76
		}
		return false
	case "hcl":
		re := regexp.MustCompile(`^#[0-9a-f]{6}$`)
		return re.MatchString(value)
	case "ecl":
		re := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
		return re.MatchString(value)
	case "pid":
		re := regexp.MustCompile(`^[0-9]{9}$`)
		return re.MatchString(value)
	}
	return true
}

func main2() {
	f, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()

	ct := 0
	buf := bufio.NewReader(f)
	keys := genPassportKeysMap()
	for l, _, _ := buf.ReadLine(); l != nil; l, _, _ = buf.ReadLine() {
		entry := string(l)
		if entry == "" {
			delete(keys, "cid")
			if len(keys) == 0 {
				ct++
			}

			keys = genPassportKeysMap()
			continue
		}

		strs := strings.Split(string(l), " ")
		for _, s := range strs {
			delete(keys, s[:3])
		}
	}
	delete(keys, "cid")
	if len(keys) == 0 {
		ct++
	}
	fmt.Println(ct)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()

	ct := 0
	buf := bufio.NewReader(f)
	keys := genPassportKeysMap()
	for l, _, _ := buf.ReadLine(); l != nil; l, _, _ = buf.ReadLine() {
		entry := string(l)
		if entry == "" {
			delete(keys, "cid")
			if len(keys) == 0 {
				ct++
			}

			keys = genPassportKeysMap()
			continue
		}

		strs := strings.Split(string(l), " ")
		for _, s := range strs {
			if isKeyValid(s[:3], s[4:]) {
				delete(keys, s[:3])
			}
		}
	}
	delete(keys, "cid")
	if len(keys) == 0 {
		ct++
	}
	fmt.Println(ct)
}
