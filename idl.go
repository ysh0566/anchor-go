package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/dave/jennifer/jen"
	"github.com/gagliardetto/solana-go"
)

type Pda struct {
	Seeds []struct {
		Kind  string `json:"kind"`
		Value []byte `json:"value,omitempty"`
		Path  string `json:"path,omitempty"`
	} `json:"seeds"`
	Program *Program `json:"program,omitempty"`
}

type Program struct {
	Kind  string `json:"kind"`
	Value []byte `json:"value"`
}

type IDL struct {
	Address solana.PublicKey `json:"address"`

	Name    string `json:"name"`    // compatible with earlier versions
	Version string `json:"version"` // compatible with earlier versions

	Metadata struct {
		Name        string `json:"name"`
		Version     string `json:"version"`
		Spec        string `json:"spec"`
		Description string `json:"description"`

		Address solana.PublicKey `json:"address"` // compatible with earlier versions
	} `json:"metadata"`
	Instructions []struct {
		Name          string `json:"name"`
		Discriminator []byte `json:"discriminator"`
		Accounts      []struct {
			Name      string   `json:"name"`
			Relations []string `json:"relations,omitempty"`
			Writable  bool     `json:"writable,omitempty"`
			Pda       *Pda     `json:"pda,omitempty"`
			Signer    bool     `json:"signer,omitempty"`
			Address   string   `json:"address,omitempty"`
			Docs      []string `json:"docs,omitempty"`
		} `json:"accounts"`
		Args []struct {
			Name string   `json:"name"`
			Type TypeKind `json:"type"`
		} `json:"args"`
		Docs []string `json:"docs,omitempty"`
	} `json:"instructions"`
	Accounts []struct {
		Name          string `json:"name"`
		Discriminator []byte `json:"discriminator"`

		Type Type `json:"type"` // compatible with earlier versions
	} `json:"accounts"`
	Errors []struct {
		Code int    `json:"code"`
		Name string `json:"name"`
		Msg  string `json:"msg"`
	} `json:"errors"`
	Types []struct {
		Name string `json:"name"`
		Type Type   `json:"type"`
	} `json:"types"`
}

func (idl *IDL) Normalize() {
	if idl.Name == "" {
		return
	}
	idl.Address = idl.Metadata.Address
	idl.Metadata.Name = idl.Name
	idl.Metadata.Version = idl.Version
}

type Type struct {
	Kind   string `json:"kind"`
	Fields []struct {
		Name string   `json:"name"`
		Type TypeKind `json:"type"`
		Docs []string `json:"docs,omitempty"`
	} `json:"fields"`
}

type TypeKind struct {
	Scalar     bool // indicates that the type is a scalar, e.g. u64, i64, etc.
	ScalarKind string

	Vec       bool            // indicates that the type is an vector
	ArrayKind [2]ArrayElement `json:"array"`

	Defined string

	Option bool
}

func (kind TypeKind) GoType() jen.Code {
	scalarMapping := func(scalar string) jen.Code {
		switch scalar {
		case "bool":
			return jen.Bool()
		case "u8":
			return jen.Uint8()
		case "u16":
			return jen.Uint16()
		case "u32":
			return jen.Uint32()
		case "u64":
			return jen.Uint64()
		case "u128":
			return jen.Op("*").Qual("math/big", "Int")
		case "i8":
			return jen.Int8()
		case "i16":
			return jen.Int16()
		case "i32":
			return jen.Int32()
		case "i64":
			return jen.Int64()
		case "i128":
			return nil // Not supported yet
		case "f32", "f64":
			return jen.Id(fmt.Sprintf("float%s", string([]byte(scalar)[1:])))
		case "string":
			return jen.String()
		case "pubkey", "publicKey":
			return jen.Qual("github.com/gagliardetto/solana-go", "PublicKey")
		case "bytes":
			return jen.Index().Byte()
		default:
			panic(fmt.Sprintf("unsupported scalar type: %s", scalar))
		}
	}
	var res = &jen.Statement{}
	if kind.Option {
		res = res.Op("*")
	}
	if kind.Scalar {
		return res.Add(scalarMapping(kind.ScalarKind))
	}
	if kind.Defined != "" {
		return res.Add(jen.Id(kind.Defined))
	}
	if kind.Vec {
		return res.Add(jen.Index().Add(scalarMapping(kind.ScalarKind)))
	}

	var arrayType jen.Code
	if kind.ArrayKind[0].Scalar == "" {
		arrayType = jen.Id(kind.ArrayKind[0].Defined)
	} else {
		arrayType = scalarMapping(kind.ArrayKind[0].Scalar)
		if arrayType == nil {
			return nil // Not supported yet
		}
	}
	return res.Add(jen.Index(jen.Id(strconv.Itoa(kind.ArrayKind[1].ArrayLength))).Add(arrayType))
}

func (kind *TypeKind) UnmarshalJSON(bz []byte) error {
	var s string
	if err := json.Unmarshal(bz, &s); err == nil {
		kind.Scalar = true
		kind.ScalarKind = s
		return nil
	}

	obj := struct {
		Defined struct {
			Name string `json:"name"`
		} `json:"defined"`
		Array  [2]ArrayElement `json:"array"`
		Option json.RawMessage `json:"option"`
		Vec    json.RawMessage `json:"vec"`
	}{}
	if err := json.Unmarshal(bz, &obj); err != nil {
		return err
	}
	if obj.Option != nil {
		if err := kind.UnmarshalJSON(obj.Option); err != nil {
			return err
		}
		kind.Option = true
		return nil
	}
	if obj.Vec != nil {
		if err := kind.UnmarshalJSON(obj.Vec); err != nil {
			return err
		}
		kind.Vec = true
		return nil
	}

	if obj.Defined.Name == "" && obj.Array[1].ArrayLength == 0 {
		return fmt.Errorf("unsupported type kind: %s", string(bz))
	}
	kind.Defined = obj.Defined.Name
	kind.ArrayKind = obj.Array
	return nil
}

type ArrayElement struct {
	Scalar      string
	Defined     string
	ArrayLength int
}

func (element *ArrayElement) UnmarshalJSON(bz []byte) error {
	var s string
	if err := json.Unmarshal(bz, &s); err == nil { // scalar type, e.g. u8, i8, etc.
		element.Scalar = s
		return nil
	}
	var i int
	if err := json.Unmarshal(bz, &i); err == nil { // array length
		element.ArrayLength = i
		return nil
	}
	obj := struct {
		Defined struct {
			Name string `json:"name"`
		} `json:"defined"`
	}{}
	if err := json.Unmarshal(bz, &obj); err != nil {
		return err
	}
	if obj.Defined.Name == "" {
		return fmt.Errorf("unsupported array element: %s", string(bz))
	}
	element.Defined = obj.Defined.Name
	return nil
}
