package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringPad_BlankTest(t *testing.T) {
	result := StringPad("", "", 0, PAD_TYPE_LEFT)
	if !assert.Equal(t, "", result) {
		t.Fail()
	}
}

func TestStringPad_NPadGtz(t *testing.T) {
	result := StringPad("", "0", 4, PAD_TYPE_LEFT)
	if !assert.Equal(t, "0000", result) {
		t.Fail()
	}
}

func TestStringPad_PadLeft(t *testing.T) {
	result := StringPad("1", "0", 4, PAD_TYPE_LEFT)
	if !assert.Equal(t, "0001", result) {
		t.Fail()
	}
}

func TestStringPad_PadLeft_X(t *testing.T) {
	result := StringPad("1", "X", 2, PAD_TYPE_LEFT)
	if !assert.Equal(t, "X1", result) {
		t.Fail()
	}
}

func TestStringPad_PadRight(t *testing.T) {
	result := StringPad("1", "0", 4, PAD_TYPE_RIGHT)
	if !assert.Equal(t, "1000", result) {
		t.Fail()
	}
}

func TestStringPad_PadRight_X(t *testing.T) {
	result := StringPad("1", "X", 4, PAD_TYPE_RIGHT)
	if !assert.Equal(t, "1XXX", result) {
		t.Fail()
	}
}

func TestStringPad_PadLeft_NeedleExceeds(t *testing.T) {
	result := StringPad("00001", "0", 4, PAD_TYPE_LEFT)
	if !assert.Equal(t, "00001", result) {
		t.Fail()
	}
}

func TestStringPad_PadRight_NeedleExceeds(t *testing.T) {
	result := StringPad("100000", "0", 4, PAD_TYPE_RIGHT)
	if !assert.Equal(t, "100000", result) {
		t.Fail()
	}
}
