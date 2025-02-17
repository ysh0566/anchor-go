package main

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {
	t.Run("counter", func(t *testing.T) {
		bz, err := os.ReadFile("examples/counter/counter.json")
		require.NoError(t, err)
		var idl IDL
		err = json.Unmarshal(bz, &idl)
		require.NoError(t, err)

		src, err := generate(&idl, "counter", true, true, true)
		require.NoError(t, err)

		err = writeToFile("examples/counter/counter.json", "examples/counter", src)
		require.NoError(t, err)
	})
	t.Run("pump.fun", func(t *testing.T) {
		bz, err := os.ReadFile("examples/pumpfun/pumpfun.json")
		require.NoError(t, err)
		var idl IDL
		err = json.Unmarshal(bz, &idl)
		require.NoError(t, err)

		src, err := generate(&idl, "pumpfun", true, true, true)
		require.NoError(t, err)

		err = writeToFile("examples/pumpfun/pumpfun.json", "examples/pumpfun", src)
		require.NoError(t, err)
	})
}
