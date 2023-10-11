package main

import "testing"

func TestMaskingUrl(t *testing.T) {
	str := "Here's my spammy page: http://hehefouls.netHAHAHA see you."
	expected := "Here's my spammy page: http://******************* see you."

	result := MaskingUrl(str)

	if result != expected {
		t.Errorf("Incorrect result.Expect %s, got %s",
			expected, result)
	}
}
