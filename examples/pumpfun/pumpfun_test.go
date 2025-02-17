package pumpfun

import (
	"context"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPumpFunDecoder(t *testing.T) {
	cli := rpc.New("https://api.mainnet-beta.solana.com/")
	defer cli.Close()

	resp, err := cli.GetAccountInfo(context.Background(), solana.MustPublicKeyFromBase58("4wTV1YmiEkRvAtNtsSGPtUrqRYQMe5SKy2uB4Jjaxnjf"))
	require.NoError(t, err)
	global, err := DecodeGlobal(resp)
	require.NoError(t, err)
	require.NotNil(t, global)
	t.Log(global)

	resp, err = cli.GetAccountInfo(context.Background(), solana.MustPublicKeyFromBase58("6aXPQwo4kVtLNu2kjZyR2jRTH647nPMuuj4Tu8WGBBcb"))
	require.NoError(t, err)
	bondingCurve, err := DecodeBondingCurve(resp)
	require.NoError(t, err)
	require.NotNil(t, bondingCurve)
	t.Log(bondingCurve)
}
