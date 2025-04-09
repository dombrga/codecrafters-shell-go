package main

import (
	"fmt"
	"testing"
)

func TestGetInvalidPrint(t *testing.T) {
	tcs := []struct {
		input string
		want  string
	}{
		{input: "invalidcommand", want: "invalidcommand: command not found"},
		{input: "invalidcommand", want: "invalidcommand: command not found"},
	}

	for i, tc := range tcs {
		got := GetInvalidPrint(tc.input)
		if got == tc.want {
			fmt.Printf("Test %d: got \"%s\", want \"%s\" - PASSED\n", i+1, got, tc.want)
		} else {
			t.Errorf("Test %d: got \"%s\", want \"%s\" - FAILED\n", i+1, got, tc.want)
		}
	}
}

func TestGetEchoPrint(t *testing.T) {
	tcs := []struct {
		input string
		want  string
	}{
		{input: "echo foo", want: "foo"},
		{input: "echo foobar", want: "foobar"},
		{input: "echo foo bar", want: "foo bar"},
		{input: "echo foo     bar", want: "foo     bar"},
		{input: "echo   foo     bar", want: "foo     bar"},
	}

	for i, tc := range tcs {
		got := GetEchoPrint(tc.input)
		if got == tc.want {
			fmt.Printf("Test %d: got \"%s\", want \"%s\" - PASSED\n", i+1, got, tc.want)
		} else {
			t.Errorf("Test %d: got \"%s\", want \"%s\" - FAILED\n", i+1, got, tc.want)
		}
	}
}
