package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/line/lbm-sdk/crypto/keys/secp256k1"
	"github.com/line/lbm-sdk/codec"
	codectypes "github.com/line/lbm-sdk/codec/types"
	cryptocodec "github.com/line/lbm-sdk/crypto/codec"
	sdk "github.com/line/lbm-sdk/types"
	"github.com/line/lbm-sdk/types/tx/signing"
	authsigning "github.com/line/lbm-sdk/x/auth/signing"
	"github.com/line/lbm-sdk/x/auth/tx"
	authtx "github.com/line/lbm-sdk/x/auth/tx"
	authtypes "github.com/line/lbm-sdk/x/auth/types"
	banktypes "github.com/line/lbm-sdk/x/bank/types"
	collection "github.com/line/lbm-sdk/x/collection"
)

func main() {
	args := os.Args[1:]

	numMandatoryArgs := 2
	if len(args) < numMandatoryArgs {
		words := []string{
			os.Args[0],
			"tx_bytes",
			"chain_id",
			"[acc_num]",
			"[pubkey_bytes]",
		}
		panic(strings.Join(words, " "))
	}

	txBytesEncoded := args[0]
	chainID := args[1]

	var accNum int
	var pubkeyBytesEncoded string
	if len(args) > numMandatoryArgs {
		var err error
		accNum, err = strconv.Atoi(args[2])
		if err != nil {
			panic(err)
		}
		pubkeyBytesEncoded = args[3]
	}

	// prepare codec
	registry := codectypes.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(registry)
	txConfig := authtx.NewTxConfig(marshaler, tx.DefaultSignModes)

	// register
	sdk.RegisterInterfaces(registry)
	cryptocodec.RegisterInterfaces(registry)
	authtypes.RegisterInterfaces(registry)
	banktypes.RegisterInterfaces(registry)
	collection.RegisterInterfaces(registry)

	// decode tx
	txBytes, err := base64.StdEncoding.DecodeString(txBytesEncoded)
	if err != nil {
		panic(err)
	}
	decoder := authtx.DefaultTxDecoder(marshaler)
	tx, err := decoder(txBytes)
	if err != nil {
		panic(err)
	}

	// get signature
	sigTx := tx.(authsigning.SigVerifiableTx)
	signer := sigTx.GetSigners()[0]
	sigs, err := sigTx.GetSignaturesV2()
	if err != nil {
		panic(err)
	}
	sig := sigs[0]

	if mode := sig.Data.(*signing.SingleSignatureData).SignMode; mode == signing.SignMode_SIGN_MODE_DIRECT {
		println(successOutput(signer, sig.Sequence))
		return
	}

	// decode pubkey
	pubkeybytes, err := base64.StdEncoding.DecodeString(pubkeyBytesEncoded)
	if err != nil {
		panic(err)
	}
	pubkey := &secp256k1.PubKey{
		Key: pubkeybytes,
	}

	// find sequence
	seqBegin := uint64(0)
	seqEnd := seqBegin + 1000000
	for i := seqBegin; i < seqEnd; i++ {
		signerData := authsigning.SignerData{
			ChainID: chainID,
			AccountNumber: uint64(accNum),
			Sequence: i,
		}
		if err := authsigning.VerifySignature(pubkey, signerData, sig.Data, txConfig.SignModeHandler(), tx); err == nil {
			println(successOutput(signer, i))
			return
		}
	}
	panic(fmt.Sprintf("signer: %s, sequence not in [%d, %d)", signer, seqBegin, seqEnd))
}

func successOutput(signer sdk.AccAddress, sequence uint64) string {
	return fmt.Sprintf("signer: %s, sequence: %d", signer, sequence)
}
