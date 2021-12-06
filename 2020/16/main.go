package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type rangeValue struct {
	start int
	end   int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)

	}
	defer f.Close()

	fields := map[string][2]rangeValue{}
	curPart := 0
	scanner := bufio.NewScanner(f)
	yourTicketHeader := false
	yourTicket := []int{}
	nearbyTicketsHeader := false
	nearbyTickets := [][]int{}

	accError := 0
	ctLine := 0
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			curPart++
			continue
		}

		switch curPart {
		case 0: // fields
			// departure location: 32-209 or 234-963
			var s0, e0, s1, e1 int
			txtIn := strings.Split(txt, ":")
			if _, err := fmt.Sscanf(strings.TrimSpace(txtIn[1]), "%d-%d or %d-%d", &s0, &e0, &s1, &e1); err != nil {
				log.Fatal(err)
			}

			fields[txtIn[0]] = [2]rangeValue{
				{
					start: s0,
					end:   e0,
				},
				{
					start: s1,
					end:   e1,
				},
			}
		case 1: // your ticket
			if !yourTicketHeader {
				yourTicketHeader = true
				break
			}

			values := strings.Split(txt, ",")
			for _, v := range values {
				i, err := strconv.Atoi(v)
				if err != nil {
					log.Fatal(err)
				}

				yourTicket = append(yourTicket, i)
			}
		case 2: // nearby tickets
			if !nearbyTicketsHeader {
				nearbyTicketsHeader = true
				break
			}

			values := strings.Split(txt, ",")
			curTicket := []int{}
			for _, v := range values {
				i, err := strconv.Atoi(v)
				if err != nil {
					log.Fatal(err)
				}

				curTicket = append(curTicket, i)
				valid := false
				for _, ranges := range fields {
					if (ranges[0].start <= i && i <= ranges[0].end) || (ranges[1].start <= i && i <= ranges[1].end) {
						valid = true
						break
					}
				}
				if !valid {
					accError += i
					curTicket = []int{}
					break
				}
			}
			if len(curTicket) > 0 {
				nearbyTickets = append(nearbyTickets, curTicket)
			}

			ctLine++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(accError)

	// find fields structure
	selectableFields := make([]map[string]struct{}, len(fields))
	for i := 0; i < len(fields); i++ {
		selectableFields[i] = map[string]struct{}{}
		// all fields can have any name at the beginning
		for name := range fields {
			selectableFields[i][name] = struct{}{}
		}
	}
	for _, ticket := range nearbyTickets {
		for i := 0; i < len(ticket); i++ {
			v := ticket[i]
			fieldsToRemove := []string{}
			for name := range selectableFields[i] {
				ranges := fields[name]
				if (v < ranges[0].start || ranges[0].end < v) && (v < ranges[1].start && ranges[1].end < v) {
					fieldsToRemove = append(fieldsToRemove, name)
				}
			}
			for _, s := range fieldsToRemove {
				delete(selectableFields[i], s)
			}
		}
	}

	fmt.Println(selectableFields)
}
