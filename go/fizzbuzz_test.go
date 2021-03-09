package main

import (
	"fmt"
	"testing"
)

var cases = []struct {
	Input    int
	Response string
}{
	{Input: 1, Response: "1"},
	{Input: 2, Response: "2"},
	{Input: 4, Response: "4"},
	{Input: 7, Response: "7"},
	{Input: 22, Response: "22"},
	{Input: 91, Response: "91"},
	{Input: 202, Response: "202"},
	{Input: 3, Response: "fizz"},
	{Input: 6, Response: "fizz"},
	{Input: 9, Response: "fizz"},
	{Input: 12, Response: "fizz"},
	{Input: 18, Response: "fizz"},
	{Input: 21, Response: "fizz"},
	{Input: 5, Response: "buzz"},
	{Input: 10, Response: "buzz"},
	{Input: 20, Response: "buzz"},
	{Input: 25, Response: "buzz"},
	{Input: 35, Response: "buzz"},
	{Input: 40, Response: "buzz"},
	{Input: 15, Response: "fizzbuzz"},
	{Input: 30, Response: "fizzbuzz"},
	{Input: 45, Response: "fizzbuzz"},
	{Input: 60, Response: "fizzbuzz"},
	{Input: 75, Response: "fizzbuzz"},
	{Input: 90, Response: "fizzbuzz"},
}

func TestFizzBuzz(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Input, test.Response), func(t *testing.T) {
			got := FizzBuzz(test.Input)
			assertEqual(t, got, test.Response)
		})
	}
}

// Benchmark FizzBuzz function
func BenchmarkFizzBuzz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FizzBuzz(i)
	}
}

// Example FizzBuzz usage
func ExampleFizzBuzz() {
	fmt.Println(FizzBuzz(3))
	//Output: fizz
}

// should create struct of given size
func TestN(t *testing.T) {
	got := N(12)
	want := 12
	if len(got) != want {
		t.Errorf("expected %d got %d", got, want)
	}
}

func BenchmarkN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		N(i)
	}
}

func assertEqual(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("expected %q got %q", got, want)
	}
}
