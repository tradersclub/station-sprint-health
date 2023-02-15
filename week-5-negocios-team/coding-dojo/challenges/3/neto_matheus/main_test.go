package main

import (
	"reflect"
	"testing"
)

var expectedCalculatePersonAge = []User{
	{
		name:      "leo",
		birthYear: 1996,
		age:       27,
	},
	{
		name:      "robertin",
		birthYear: 2000,
		age:       23,
	},
	{
		name:      "theo",
		birthYear: 2018,
		age:       5,
	},
	{
		name:      "neto",
		birthYear: 1991,
		age:       32,
	},
	{
		name:      "matheus",
		birthYear: 1994,
		age:       29,
	},
}

func Test_CalculatePersonAge(t *testing.T) {
	got := CalculatePersonAge(usersTest)
	if !reflect.DeepEqual(got, expectedCalculatePersonAge) {
		t.Error("deu ruim")
	}

}

func Benchmark_CalculatePersonAge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculatePersonAge(usersTest)
	}
}

func Test_Get18OrMore(t *testing.T) {
	want := []User{
		{
			name:      "leo",
			birthYear: 1996,
			age:       27,
		},
		{
			name:      "robertin",
			birthYear: 2000,
			age:       23,
		},
		{
			name:      "neto",
			birthYear: 1991,
			age:       32,
		},
		{
			name:      "matheus",
			birthYear: 1994,
			age:       29,
		},
	}

	got := Get18OrMore(expectedCalculatePersonAge)

	if !reflect.DeepEqual(got, want) {
		t.Error("deu ruim novamente")
	}
}

func Benchmark_Get180rMore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get18OrMore(expectedCalculatePersonAge)
	}
}
