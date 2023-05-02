package dev02

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWithoutEscapeExamples(t *testing.T) {
	require.Equal(t, "aaaabccddddde", Run("a4bc2d5e"), "check a4bc2d5e")
	require.Equal(t, "abcd", Run("abcd"), "check abcd")
	require.Equal(t, "", Run("45"), "check 45")
	require.Equal(t, "", Run(""), "check emty")
}

func TestWithEscapeExamples(t *testing.T) {
	require.Equal(t, `qwe45`, Run(`qwe\4\5`), `check qwe\4\5`)
	require.Equal(t, `qwe44444`, Run(`qwe\45`), `check qwe\45`)
	require.Equal(t, `qwe\\\\\`, Run(`qwe\\5`), `check qwe\\5`)
}
