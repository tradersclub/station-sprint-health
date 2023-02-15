package main

import "testing"

func Test_CalculatePersonAge(t *testing.T) {
	expected := []User{
		{
			name:      "leo",
			birthYear: 1996,
		},
		{
			name:      "robertin",
			birthYear: 2000,
		},
		{
			name:      "neto",
			birthYear: 1991,
		},
		{
			name:      "matheus",
			birthYear: 1994,
		},
	}
	usrs := CalculatePersonAge(usersTest)

}

func Benchmark_CalculatePersonAge(t *testing.B) {}
