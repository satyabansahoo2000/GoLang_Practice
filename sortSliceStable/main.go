package main

import (
	"fmt"
	"sort"
)

type dob struct {
	day   int
	month int
	year  int
}

type people struct {
	name  string
	email string
	dob   dob
}

var members map[int]people

func main() {
	members = make(map[int]people)
	members[1] = people{
		name:  "test1",
		email: "test1@xyz.com",
		dob: dob{
			day:   28,
			month: 02,
			year:  2000,
		},
	}
	members[2] = people{
		name:  "test2",
		email: "test2@xyz.com",
		dob: dob{
			day:   05,
			month: 05,
			year:  1999,
		},
	}
	members[3] = people{
		name:  "test3",
		email: "test3@xyz.com",
		dob: dob{
			day:   07,
			month: 07,
			year:  2001,
		},
	}
	
	var keys []int
	for key := range members {
		keys = append(keys, key)
	}

	var memberValues []people
	for _, k := range keys {
		memberValues = append(memberValues, members[k])
	}

	sort.SliceStable(memberValues, func(i, j int) bool {
		if memberValues[i].dob.year != memberValues[j].dob.year {
			return memberValues[i].dob.year < memberValues[j].dob.year
		}
		if memberValues[i].dob.month != memberValues[j].dob.month {
			return memberValues[i].dob.month < memberValues[j].dob.month
		}
		return memberValues[i].dob.day < memberValues[j].dob.day
	})

	for _, v := range memberValues {
		fmt.Println(v.name, v.email, v.dob)
	}
}