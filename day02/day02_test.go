package day02

import (
	"adventofcode2022/utils/input"
	"testing"
)

type SampleReader struct{}

func (SampleReader) GetInput(day int, isTest ...bool) string {
	return input.Reader.GetInput(input.Reader{}, 2, true)
}

func TestPart1(t *testing.T) {
	expected := 15
	got, err := Part1(SampleReader{})

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if expected != got {
		t.Logf("Expected %v but got %v", expected, got)
		t.Fail()
	}
}

func TestPart2(t *testing.T) {
	expected := 12
	got, err := Part2(SampleReader{})

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if expected != got {
		t.Logf("Expected %v but got %v", expected, got)
		t.Fail()
	}
}
