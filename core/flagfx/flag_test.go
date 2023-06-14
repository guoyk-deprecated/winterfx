package flagfx

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestNewFlagSet(t *testing.T) {
	s := New()
	require.NoError(t, s.Parse([]string{"--conf", "hello"}))
	f := s.Lookup("conf")
	require.Equal(t, "hello", f.Value.String())
}

func TestParseFlagSet(t *testing.T) {
	s := New()
	_ = s.String("ignore", "", "test")
	val := s.String("hello", "", "test")
	require.NoError(t, os.Setenv("HELLO", "WORLD"))
	require.NoError(t, Parse(ParseOptions{FlagSet: s, Args: Args{"--ignore", "world"}}))
	require.Equal(t, "WORLD", *val)
}
