package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		wantErr  bool
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde", wantErr: false},
		{input: "abccd", expected: "abccd", wantErr: false},
		{input: "", expected: "", wantErr: false},
		{input: "aaa0b", expected: "aab", wantErr: false},
		{input: "56daf", expected: "", wantErr: true},
		{input: "1bdsaw", expected: "", wantErr: true},
		{input: "sdt79aq", expected: "", wantErr: true},
		{input: `d\n5abc`, expected: "d\n\n\n\n\nabc", wantErr: false},
		{input: `a\ngs`, expected: "a\ngs", wantErr: false},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			if tc.wantErr {
				require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, result)
			}
		})
	}
}
