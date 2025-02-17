// Code generated by solana-idl-generator. DO NOT EDIT.
package pumpfun

import (
	"bytes"
	"fmt"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

var PumpProgramID = solana.MustPublicKeyFromBase58("6EF8rrecthR5Dkzon8Nwu78hRvfCKubJ14M5uBEwF6P")

// Creates the global state.
type Initialize struct {
	Accounts struct {
		Global        solana.PublicKey
		User          solana.PublicKey
		SystemProgram solana.PublicKey
	}
	RemainingAccounts solana.AccountMetaSlice
	Args              struct{}

	programID *solana.PublicKey
}

func (m *Initialize) SetProgramID(id solana.PublicKey) {
	m.programID = &id
}

func (m *Initialize) ProgramID() solana.PublicKey {
	if m.programID != nil {
		return *m.programID
	}
	return PumpProgramID
}

func (m *Initialize) Discriminator() []byte {
	return []byte{}
}

func (m *Initialize) Build() (solana.Instruction, error) {
	var generic solana.GenericInstruction
	generic.ProgID = m.ProgramID()
	generic.AccountValues = solana.AccountMetaSlice{{
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.Global,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.User,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.SystemProgram,
	}}
	for _, item := range m.RemainingAccounts {
		generic.AccountValues = append(generic.AccountValues, item)
	}
	buf := bytes.Buffer{}
	buf.Write(m.Discriminator())
	enc := bin.NewBorshEncoder(&buf)
	err := enc.Encode(&m.Args)
	if err != nil {
		return nil, err
	}
	generic.DataBytes = buf.Bytes()
	return &generic, nil
}

func (m *Initialize) Decode(accounts []solana.PublicKey, data []byte) error {
	if len(accounts) < 3 {
		return fmt.Errorf("insufficient accounts, expect at least %d, but got %d", 3, len(accounts))
	}
	m.Accounts.Global = accounts[0]
	m.Accounts.User = accounts[1]
	m.Accounts.SystemProgram = accounts[2]
	for _, item := range accounts {
		m.RemainingAccounts = append(m.RemainingAccounts, solana.Meta(item))
	}
	dec := bin.NewBorshDecoder(data[len(m.Discriminator()):])
	return dec.Decode(&m.Args)
}

// Sets the global state parameters.
type SetParams struct {
	Accounts struct {
		Global         solana.PublicKey
		User           solana.PublicKey
		SystemProgram  solana.PublicKey
		EventAuthority solana.PublicKey
		Program        solana.PublicKey
	}
	RemainingAccounts solana.AccountMetaSlice
	Args              struct {
		FeeRecipient                solana.PublicKey
		InitialVirtualTokenReserves uint64
		InitialVirtualSolReserves   uint64
		InitialRealTokenReserves    uint64
		TokenTotalSupply            uint64
		FeeBasisPoints              uint64
	}

	programID *solana.PublicKey
}

func (m *SetParams) SetProgramID(id solana.PublicKey) {
	m.programID = &id
}

func (m *SetParams) ProgramID() solana.PublicKey {
	if m.programID != nil {
		return *m.programID
	}
	return PumpProgramID
}

func (m *SetParams) Discriminator() []byte {
	return []byte{}
}

func (m *SetParams) Build() (solana.Instruction, error) {
	var generic solana.GenericInstruction
	generic.ProgID = m.ProgramID()
	generic.AccountValues = solana.AccountMetaSlice{{
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.Global,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.User,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.SystemProgram,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.EventAuthority,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.Program,
	}}
	for _, item := range m.RemainingAccounts {
		generic.AccountValues = append(generic.AccountValues, item)
	}
	buf := bytes.Buffer{}
	buf.Write(m.Discriminator())
	enc := bin.NewBorshEncoder(&buf)
	err := enc.Encode(&m.Args)
	if err != nil {
		return nil, err
	}
	generic.DataBytes = buf.Bytes()
	return &generic, nil
}

func (m *SetParams) Decode(accounts []solana.PublicKey, data []byte) error {
	if len(accounts) < 5 {
		return fmt.Errorf("insufficient accounts, expect at least %d, but got %d", 5, len(accounts))
	}
	m.Accounts.Global = accounts[0]
	m.Accounts.User = accounts[1]
	m.Accounts.SystemProgram = accounts[2]
	m.Accounts.EventAuthority = accounts[3]
	m.Accounts.Program = accounts[4]
	for _, item := range accounts {
		m.RemainingAccounts = append(m.RemainingAccounts, solana.Meta(item))
	}
	dec := bin.NewBorshDecoder(data[len(m.Discriminator()):])
	return dec.Decode(&m.Args)
}

// Creates a new coin and bonding curve.
type Create struct {
	Accounts struct {
		Mint                   solana.PublicKey
		MintAuthority          solana.PublicKey
		BondingCurve           solana.PublicKey
		AssociatedBondingCurve solana.PublicKey
		Global                 solana.PublicKey
		MplTokenMetadata       solana.PublicKey
		Metadata               solana.PublicKey
		User                   solana.PublicKey
		SystemProgram          solana.PublicKey
		TokenProgram           solana.PublicKey
		AssociatedTokenProgram solana.PublicKey
		Rent                   solana.PublicKey
		EventAuthority         solana.PublicKey
		Program                solana.PublicKey
	}
	RemainingAccounts solana.AccountMetaSlice
	Args              struct {
		Name   string
		Symbol string
		Uri    string
	}

	programID *solana.PublicKey
}

func (m *Create) SetProgramID(id solana.PublicKey) {
	m.programID = &id
}

func (m *Create) ProgramID() solana.PublicKey {
	if m.programID != nil {
		return *m.programID
	}
	return PumpProgramID
}

func (m *Create) Discriminator() []byte {
	return []byte{}
}

func (m *Create) Build() (solana.Instruction, error) {
	var generic solana.GenericInstruction
	generic.ProgID = m.ProgramID()
	generic.AccountValues = solana.AccountMetaSlice{{
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.Mint,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.MintAuthority,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.BondingCurve,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.AssociatedBondingCurve,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.Global,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.MplTokenMetadata,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.Metadata,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.User,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.SystemProgram,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.TokenProgram,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.AssociatedTokenProgram,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.Rent,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.EventAuthority,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.Program,
	}}
	for _, item := range m.RemainingAccounts {
		generic.AccountValues = append(generic.AccountValues, item)
	}
	buf := bytes.Buffer{}
	buf.Write(m.Discriminator())
	enc := bin.NewBorshEncoder(&buf)
	err := enc.Encode(&m.Args)
	if err != nil {
		return nil, err
	}
	generic.DataBytes = buf.Bytes()
	return &generic, nil
}

func (m *Create) Decode(accounts []solana.PublicKey, data []byte) error {
	if len(accounts) < 14 {
		return fmt.Errorf("insufficient accounts, expect at least %d, but got %d", 14, len(accounts))
	}
	m.Accounts.Mint = accounts[0]
	m.Accounts.MintAuthority = accounts[1]
	m.Accounts.BondingCurve = accounts[2]
	m.Accounts.AssociatedBondingCurve = accounts[3]
	m.Accounts.Global = accounts[4]
	m.Accounts.MplTokenMetadata = accounts[5]
	m.Accounts.Metadata = accounts[6]
	m.Accounts.User = accounts[7]
	m.Accounts.SystemProgram = accounts[8]
	m.Accounts.TokenProgram = accounts[9]
	m.Accounts.AssociatedTokenProgram = accounts[10]
	m.Accounts.Rent = accounts[11]
	m.Accounts.EventAuthority = accounts[12]
	m.Accounts.Program = accounts[13]
	for _, item := range accounts {
		m.RemainingAccounts = append(m.RemainingAccounts, solana.Meta(item))
	}
	dec := bin.NewBorshDecoder(data[len(m.Discriminator()):])
	return dec.Decode(&m.Args)
}

// Buys tokens from a bonding curve.
type Buy struct {
	Accounts struct {
		Global                 solana.PublicKey
		FeeRecipient           solana.PublicKey
		Mint                   solana.PublicKey
		BondingCurve           solana.PublicKey
		AssociatedBondingCurve solana.PublicKey
		AssociatedUser         solana.PublicKey
		User                   solana.PublicKey
		SystemProgram          solana.PublicKey
		TokenProgram           solana.PublicKey
		Rent                   solana.PublicKey
		EventAuthority         solana.PublicKey
		Program                solana.PublicKey
	}
	RemainingAccounts solana.AccountMetaSlice
	Args              struct {
		Amount     uint64
		MaxSolCost uint64
	}

	programID *solana.PublicKey
}

func (m *Buy) SetProgramID(id solana.PublicKey) {
	m.programID = &id
}

func (m *Buy) ProgramID() solana.PublicKey {
	if m.programID != nil {
		return *m.programID
	}
	return PumpProgramID
}

func (m *Buy) Discriminator() []byte {
	return []byte{}
}

func (m *Buy) Build() (solana.Instruction, error) {
	var generic solana.GenericInstruction
	generic.ProgID = m.ProgramID()
	generic.AccountValues = solana.AccountMetaSlice{{
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.Global,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.FeeRecipient,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.Mint,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.BondingCurve,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.AssociatedBondingCurve,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.AssociatedUser,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.User,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.SystemProgram,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.TokenProgram,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.Rent,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.EventAuthority,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.Program,
	}}
	for _, item := range m.RemainingAccounts {
		generic.AccountValues = append(generic.AccountValues, item)
	}
	buf := bytes.Buffer{}
	buf.Write(m.Discriminator())
	enc := bin.NewBorshEncoder(&buf)
	err := enc.Encode(&m.Args)
	if err != nil {
		return nil, err
	}
	generic.DataBytes = buf.Bytes()
	return &generic, nil
}

func (m *Buy) Decode(accounts []solana.PublicKey, data []byte) error {
	if len(accounts) < 12 {
		return fmt.Errorf("insufficient accounts, expect at least %d, but got %d", 12, len(accounts))
	}
	m.Accounts.Global = accounts[0]
	m.Accounts.FeeRecipient = accounts[1]
	m.Accounts.Mint = accounts[2]
	m.Accounts.BondingCurve = accounts[3]
	m.Accounts.AssociatedBondingCurve = accounts[4]
	m.Accounts.AssociatedUser = accounts[5]
	m.Accounts.User = accounts[6]
	m.Accounts.SystemProgram = accounts[7]
	m.Accounts.TokenProgram = accounts[8]
	m.Accounts.Rent = accounts[9]
	m.Accounts.EventAuthority = accounts[10]
	m.Accounts.Program = accounts[11]
	for _, item := range accounts {
		m.RemainingAccounts = append(m.RemainingAccounts, solana.Meta(item))
	}
	dec := bin.NewBorshDecoder(data[len(m.Discriminator()):])
	return dec.Decode(&m.Args)
}

// Sells tokens into a bonding curve.
type Sell struct {
	Accounts struct {
		Global                 solana.PublicKey
		FeeRecipient           solana.PublicKey
		Mint                   solana.PublicKey
		BondingCurve           solana.PublicKey
		AssociatedBondingCurve solana.PublicKey
		AssociatedUser         solana.PublicKey
		User                   solana.PublicKey
		SystemProgram          solana.PublicKey
		AssociatedTokenProgram solana.PublicKey
		TokenProgram           solana.PublicKey
		EventAuthority         solana.PublicKey
		Program                solana.PublicKey
	}
	RemainingAccounts solana.AccountMetaSlice
	Args              struct {
		Amount       uint64
		MinSolOutput uint64
	}

	programID *solana.PublicKey
}

func (m *Sell) SetProgramID(id solana.PublicKey) {
	m.programID = &id
}

func (m *Sell) ProgramID() solana.PublicKey {
	if m.programID != nil {
		return *m.programID
	}
	return PumpProgramID
}

func (m *Sell) Discriminator() []byte {
	return []byte{}
}

func (m *Sell) Build() (solana.Instruction, error) {
	var generic solana.GenericInstruction
	generic.ProgID = m.ProgramID()
	generic.AccountValues = solana.AccountMetaSlice{{
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.Global,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.FeeRecipient,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.Mint,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.BondingCurve,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.AssociatedBondingCurve,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.AssociatedUser,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.User,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.SystemProgram,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.AssociatedTokenProgram,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.TokenProgram,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.EventAuthority,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.Program,
	}}
	for _, item := range m.RemainingAccounts {
		generic.AccountValues = append(generic.AccountValues, item)
	}
	buf := bytes.Buffer{}
	buf.Write(m.Discriminator())
	enc := bin.NewBorshEncoder(&buf)
	err := enc.Encode(&m.Args)
	if err != nil {
		return nil, err
	}
	generic.DataBytes = buf.Bytes()
	return &generic, nil
}

func (m *Sell) Decode(accounts []solana.PublicKey, data []byte) error {
	if len(accounts) < 12 {
		return fmt.Errorf("insufficient accounts, expect at least %d, but got %d", 12, len(accounts))
	}
	m.Accounts.Global = accounts[0]
	m.Accounts.FeeRecipient = accounts[1]
	m.Accounts.Mint = accounts[2]
	m.Accounts.BondingCurve = accounts[3]
	m.Accounts.AssociatedBondingCurve = accounts[4]
	m.Accounts.AssociatedUser = accounts[5]
	m.Accounts.User = accounts[6]
	m.Accounts.SystemProgram = accounts[7]
	m.Accounts.AssociatedTokenProgram = accounts[8]
	m.Accounts.TokenProgram = accounts[9]
	m.Accounts.EventAuthority = accounts[10]
	m.Accounts.Program = accounts[11]
	for _, item := range accounts {
		m.RemainingAccounts = append(m.RemainingAccounts, solana.Meta(item))
	}
	dec := bin.NewBorshDecoder(data[len(m.Discriminator()):])
	return dec.Decode(&m.Args)
}

// Allows the admin to withdraw liquidity for a migration once the bonding curve completes
type Withdraw struct {
	Accounts struct {
		Global                 solana.PublicKey
		Mint                   solana.PublicKey
		BondingCurve           solana.PublicKey
		AssociatedBondingCurve solana.PublicKey
		AssociatedUser         solana.PublicKey
		User                   solana.PublicKey
		SystemProgram          solana.PublicKey
		TokenProgram           solana.PublicKey
		Rent                   solana.PublicKey
		EventAuthority         solana.PublicKey
		Program                solana.PublicKey
	}
	RemainingAccounts solana.AccountMetaSlice
	Args              struct{}

	programID *solana.PublicKey
}

func (m *Withdraw) SetProgramID(id solana.PublicKey) {
	m.programID = &id
}

func (m *Withdraw) ProgramID() solana.PublicKey {
	if m.programID != nil {
		return *m.programID
	}
	return PumpProgramID
}

func (m *Withdraw) Discriminator() []byte {
	return []byte{}
}

func (m *Withdraw) Build() (solana.Instruction, error) {
	var generic solana.GenericInstruction
	generic.ProgID = m.ProgramID()
	generic.AccountValues = solana.AccountMetaSlice{{
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.Global,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.Mint,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.BondingCurve,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.AssociatedBondingCurve,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.AssociatedUser,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.User,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.SystemProgram,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.TokenProgram,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.Rent,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.EventAuthority,
	}, {
		IsSigner:   false,
		IsWritable: false,
		PublicKey:  m.Accounts.Program,
	}}
	for _, item := range m.RemainingAccounts {
		generic.AccountValues = append(generic.AccountValues, item)
	}
	buf := bytes.Buffer{}
	buf.Write(m.Discriminator())
	enc := bin.NewBorshEncoder(&buf)
	err := enc.Encode(&m.Args)
	if err != nil {
		return nil, err
	}
	generic.DataBytes = buf.Bytes()
	return &generic, nil
}

func (m *Withdraw) Decode(accounts []solana.PublicKey, data []byte) error {
	if len(accounts) < 11 {
		return fmt.Errorf("insufficient accounts, expect at least %d, but got %d", 11, len(accounts))
	}
	m.Accounts.Global = accounts[0]
	m.Accounts.Mint = accounts[1]
	m.Accounts.BondingCurve = accounts[2]
	m.Accounts.AssociatedBondingCurve = accounts[3]
	m.Accounts.AssociatedUser = accounts[4]
	m.Accounts.User = accounts[5]
	m.Accounts.SystemProgram = accounts[6]
	m.Accounts.TokenProgram = accounts[7]
	m.Accounts.Rent = accounts[8]
	m.Accounts.EventAuthority = accounts[9]
	m.Accounts.Program = accounts[10]
	for _, item := range accounts {
		m.RemainingAccounts = append(m.RemainingAccounts, solana.Meta(item))
	}
	dec := bin.NewBorshDecoder(data[len(m.Discriminator()):])
	return dec.Decode(&m.Args)
}

func DecodePumpInstructionRaw(accounts []*solana.AccountMeta, data []byte) (decoded any, ok bool, err error) {
	mappings := make([]solana.PublicKey, len(accounts))
	for i, item := range accounts {
		mappings[i] = item.PublicKey
	}
	return DecodePumpInstruction(mappings, data)
}
func DecodePumpInstruction(accounts []solana.PublicKey, data []byte) (decoded any, ok bool, err error) {
	if len(data) < 8 {
		return nil, false, fmt.Errorf("data length is less than 8 bytes")
	}
	discriminator := [8]byte(data[0:8])
	switch discriminator {
	default:
		return nil, false, nil
	}
}

type Global struct {
	Initialized                 bool
	Authority                   solana.PublicKey
	FeeRecipient                solana.PublicKey
	InitialVirtualTokenReserves uint64
	InitialVirtualSolReserves   uint64
	InitialRealTokenReserves    uint64
	TokenTotalSupply            uint64
	FeeBasisPoints              uint64
}

func (m *Global) Discriminator() []byte {
	return []byte{}
}

func DecodeGlobal(resp *rpc.GetAccountInfoResult) (*Global, error) {
	var m Global
	data := resp.GetBinary()
	if len(data) < len(m.Discriminator()) {
		return nil, fmt.Errorf("invalid response data")
	}
	discriminator := data[0:len(m.Discriminator())]
	if !bytes.Equal(discriminator, m.Discriminator()) {
		return nil, fmt.Errorf("discriminator mismatch, expect %x, but got %x", m.Discriminator(), discriminator)
	}
	dec := bin.NewBorshDecoder(data[len(m.Discriminator()):])
	err := dec.Decode(&m)
	return &m, err
}

type BondingCurve struct {
	VirtualTokenReserves uint64
	VirtualSolReserves   uint64
	RealTokenReserves    uint64
	RealSolReserves      uint64
	TokenTotalSupply     uint64
	Complete             bool
}

func (m *BondingCurve) Discriminator() []byte {
	return []byte{}
}

func DecodeBondingCurve(resp *rpc.GetAccountInfoResult) (*BondingCurve, error) {
	var m BondingCurve
	data := resp.GetBinary()
	if len(data) < len(m.Discriminator()) {
		return nil, fmt.Errorf("invalid response data")
	}
	discriminator := data[0:len(m.Discriminator())]
	if !bytes.Equal(discriminator, m.Discriminator()) {
		return nil, fmt.Errorf("discriminator mismatch, expect %x, but got %x", m.Discriminator(), discriminator)
	}
	dec := bin.NewBorshDecoder(data[len(m.Discriminator()):])
	err := dec.Decode(&m)
	return &m, err
}
