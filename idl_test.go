package main

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestArrayElement_UnmarshalJSON(t *testing.T) {
	t.Run("scalar", func(t *testing.T) {
		var elements [2]ArrayElement
		err := json.Unmarshal([]byte(`["u8", 4]`), &elements)
		require.NoError(t, err)
		require.Equal(t, "u8", elements[0].Scalar)
		require.Equal(t, 4, elements[1].ArrayLength)
	})
	t.Run("defined", func(t *testing.T) {
		var elements [2]ArrayElement
		err := json.Unmarshal([]byte(`
  [
	{
	  "defined": {
		"name": "Foo"
	  }
	},
	3
  ]
`), &elements)
		require.NoError(t, err)
		require.Equal(t, "Foo", elements[0].Defined)
		require.Equal(t, 3, elements[1].ArrayLength)
	})
	t.Run("unsupported array element", func(t *testing.T) {
		var elements [2]ArrayElement
		err := json.Unmarshal([]byte(`
  [
	{
	  "unknown": {
		"name": "Foo"
	  }
	},
	3
  ]
`), &elements)
		require.ErrorContains(t, err, "unsupported array element")
	})
}

func TestTypeKind_UnmarshalJSON(t *testing.T) {
	t.Run("scalar", func(t *testing.T) {
		var kind TypeKind
		err := json.Unmarshal([]byte(`"u8"`), &kind)
		require.NoError(t, err)
		require.True(t, kind.Scalar)
		require.Equal(t, "u8", kind.ScalarKind)
	})
	t.Run("defined", func(t *testing.T) {
		var kind TypeKind
		err := json.Unmarshal([]byte(`
	{
	 "defined": {
		"name": "Foo"
	 }
	}
	`), &kind)
		require.NoError(t, err)
		require.Equal(t, "Foo", kind.Defined)
	})
	t.Run("array", func(t *testing.T) {
		var kind TypeKind
		err := json.Unmarshal([]byte(`
	{
	 "array": [
		"u8",
		4
	 ]
	}
	`), &kind)
		require.NoError(t, err)
		require.Equal(t, "u8", kind.ArrayKind[0].Scalar)
		require.Equal(t, 4, kind.ArrayKind[1].ArrayLength)
	})
	t.Run("unsupported type kind", func(t *testing.T) {
		var kind TypeKind
		err := json.Unmarshal([]byte(`
	{
	 "unknown": [
		"u8",
		4
	 ]
	}
	`), &kind)
		require.ErrorContains(t, err, "unsupported type kind")
	})
}

func TestIDL_UnmarshalJSON(t *testing.T) {
	t.Run("sample", func(t *testing.T) {
		var idl IDL
		err := json.Unmarshal([]byte(`
{
  "address": "9sMy4hnC9MML6mioESFZmzpntt3focqwUq1ymPgbMf64",
  "metadata": {
    "name": "anchor_counter",
    "version": "0.1.0",
    "spec": "0.1.0",
    "description": "Created with Anchor"
  },
  "instructions": [
    {
      "name": "increment",
      "discriminator": [11, 18, 104, 9, 104, 174, 59, 33],
      "accounts": [
        {
          "name": "counter",
          "writable": true
        },
        {
          "name": "user",
          "signer": true
        }
      ],
      "args": []
    },
    {
      "name": "initialize",
      "discriminator": [175, 175, 109, 31, 13, 152, 155, 237],
      "accounts": [
        {
          "name": "counter",
          "writable": true,
          "signer": true
        },
        {
          "name": "user",
          "writable": true,
          "signer": true
        },
        {
          "name": "system_program",
          "address": "11111111111111111111111111111111"
        }
      ],
      "args": []
    }
  ],
  "accounts": [
    {
      "name": "Counter",
      "discriminator": [255, 176, 4, 245, 188, 253, 124, 25]
    }
  ],
  "types": [
    {
      "name": "Counter",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "count",
            "type": "u64"
          }
        ]
      }
    }
  ]
}
`), &idl)
		require.NoError(t, err)
	})
}
