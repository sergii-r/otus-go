package hw02_unpack_string //nolint:golint,stylecheck

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type test struct {
	name     string
	input    string
	expected string
	err      error
}

func TestUnpack(t *testing.T) {
	for _, tst := range [...]test{
		{
			name:     "a4bc2d5e",
			input:    "a4bc2d5e",
			expected: "aaaabccddddde",
		},
		{
			name:     "abccd",
			input:    "abccd",
			expected: "abccd",
		},
		{
			name:     "3abc",
			input:    "3abc",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			name:     "45",
			input:    "45",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			name:     "aaa10b",
			input:    "aaa10b",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			name:     "",
			input:    "",
			expected: "",
		},
		{
			name:     "aaa0b",
			input:    "aaa0b",
			expected: "aab",
		},
		{
			name:     "П2ри4ве3тQ0",
			input:    "П2ри4ве3тQ0",
			expected: "ППриииивееет",
		},
		{
			name:     "d\n5abc",
			input:    "d\n5abc",
			expected: "d\n\n\n\n\nabc",
		},
	} {
		tst := tst
		t.Run(tst.name, func(t *testing.T) {
			result, err := Unpack(tst.input)
			require.Equal(t, tst.err, err)
			require.Equal(t, tst.expected, result)
		})
	}
}

func TestUnpackWithEscape(t *testing.T) {
	for _, tst := range [...]test{
		{
			name:     `qwe\4\5`,
			input:    `qwe\4\5`,
			expected: `qwe45`,
		},
		{
			name:     `qwe\45`,
			input:    `qwe\45`,
			expected: `qwe44444`,
		},
		{
			name:     `qwe\\5`,
			input:    `qwe\\5`,
			expected: `qwe\\\\\`,
		},
		{
			name:     `qwe\\5a`,
			input:    `qwe\\5a`,
			expected: `qwe\\\\\a`,
		},
		{
			name:     `qw\ne`,
			input:    `qw\ne`,
			expected: ``,
			err:      ErrInvalidEscaping,
		},
		{
			name:     `qwe\\\3`,
			input:    `qwe\\\3`,
			expected: `qwe\3`,
		},
	} {
		tst := tst
		t.Run(tst.name, func(t *testing.T) {
			result, err := Unpack(tst.input)
			require.Equal(t, tst.err, err)
			require.Equal(t, tst.expected, result)
		})
	}
}
