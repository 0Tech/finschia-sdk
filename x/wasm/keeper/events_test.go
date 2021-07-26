package keeper

import (
	"context"
	"github.com/line/lbm-sdk/x/wasm/types"
	wasmvmtypes "github.com/line/wasmvm/types"
	sdk "github.com/line/lbm-sdk/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasWasmModuleEvent(t *testing.T) {
	myContractAddr := RandomAccountAddress(t)
	specs := map[string]struct {
		srcEvents []sdk.Event
		exp       bool
	}{
		"event found": {
			srcEvents: []sdk.Event{
				sdk.NewEvent(types.WasmModuleEventType, sdk.NewAttribute("_contract_address", myContractAddr.String())),
			},
			exp: true,
		},
		"different event: not found": {
			srcEvents: []sdk.Event{
				sdk.NewEvent(types.CustomContractEventPrefix, sdk.NewAttribute("_contract_address", myContractAddr.String())),
			},
			exp: false,
		},
		"event with different address: not found": {
			srcEvents: []sdk.Event{
				sdk.NewEvent(types.WasmModuleEventType, sdk.NewAttribute("_contract_address", RandomBech32AccountAddress(t))),
			},
			exp: false,
		},
		"no event": {
			srcEvents: []sdk.Event{},
			exp:       false,
		},
	}
	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			em := sdk.NewEventManager()
			em.EmitEvents(spec.srcEvents)
			ctx := sdk.Context{}.WithContext(context.Background()).WithEventManager(em)

			got := hasWasmModuleEvent(ctx, myContractAddr)
			assert.Equal(t, spec.exp, got)
		})
	}
}

func TestNewCustomEvents(t *testing.T) {
	myContract := RandomAccountAddress(t)
	specs := map[string]struct {
		src     wasmvmtypes.Events
		exp     sdk.Events
		isError bool
	}{
		"all good": {
			src: wasmvmtypes.Events{{
				Type:       "foo",
				Attributes: []wasmvmtypes.EventAttribute{{Key: "myKey", Value: "myVal"}},
			}},
			exp: sdk.Events{sdk.NewEvent("wasm-foo",
				sdk.NewAttribute("_contract_address", myContract.String()),
				sdk.NewAttribute("myKey", "myVal"))},
		},
		"multiple attributes": {
			src: wasmvmtypes.Events{{
				Type: "foo",
				Attributes: []wasmvmtypes.EventAttribute{{Key: "myKey", Value: "myVal"},
					{Key: "myOtherKey", Value: "myOtherVal"}},
			}},
			exp: sdk.Events{sdk.NewEvent("wasm-foo",
				sdk.NewAttribute("_contract_address", myContract.String()),
				sdk.NewAttribute("myKey", "myVal"),
				sdk.NewAttribute("myOtherKey", "myOtherVal"))},
		},
		"multiple events": {
			src: wasmvmtypes.Events{{
				Type:       "foo",
				Attributes: []wasmvmtypes.EventAttribute{{Key: "myKey", Value: "myVal"}},
			}, {
				Type:       "bar",
				Attributes: []wasmvmtypes.EventAttribute{{Key: "otherKey", Value: "otherVal"}},
			}},
			exp: sdk.Events{
				sdk.NewEvent("wasm-foo",
					sdk.NewAttribute("_contract_address", myContract.String()),
					sdk.NewAttribute("myKey", "myVal")),
				sdk.NewEvent("wasm-bar",
					sdk.NewAttribute("_contract_address", myContract.String()),
					sdk.NewAttribute("otherKey", "otherVal")),
			},
		},
		"without attributes": {
			src: wasmvmtypes.Events{{
				Type: "foo",
			}},
			exp: sdk.Events{sdk.NewEvent("wasm-foo",
				sdk.NewAttribute("_contract_address", myContract.String()))},
		},
		"error on short event type": {
			src: wasmvmtypes.Events{{
				Type: "f",
			}},
			isError: true,
		},
		"error on _contract_address": {
			src: wasmvmtypes.Events{{
				Type:       "foo",
				Attributes: []wasmvmtypes.EventAttribute{{Key: "_contract_address", Value: RandomBech32AccountAddress(t)}},
			}},
			isError: true,
		},
		"error on reserved prefix": {
			src: wasmvmtypes.Events{{
				Type: "wasm",
				Attributes: []wasmvmtypes.EventAttribute{
					{Key: "_reserved", Value: "is skipped"},
					{Key: "normal", Value: "is used"}},
			}},
			isError: true,
		},
		"error on empty key": {
			src: wasmvmtypes.Events{{
				Type: "boom",
				Attributes: []wasmvmtypes.EventAttribute{
					{Key: "some", Value: "data"},
					{Key: "key", Value: ""},
				},
			}},
			isError: true,
		},
		"error on empty value": {
			src: wasmvmtypes.Events{{
				Type: "boom",
				Attributes: []wasmvmtypes.EventAttribute{
					{Key: "some", Value: "data"},
					{Key: "", Value: "value"},
				},
			}},
			isError: true,
		},
	}
	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			gotEvent, err := newCustomEvents(spec.src, myContract)
			if spec.isError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, spec.exp, gotEvent)
			}
		})
	}
}

func TestNewWasmModuleEvent(t *testing.T) {
	myContract := RandomAccountAddress(t)
	specs := map[string]struct {
		src     []wasmvmtypes.EventAttribute
		exp     sdk.Events
		isError bool
	}{
		"all good": {
			src: []wasmvmtypes.EventAttribute{{Key: "myKey", Value: "myVal"}},
			exp: sdk.Events{sdk.NewEvent("wasm",
				sdk.NewAttribute("_contract_address", myContract.String()),
				sdk.NewAttribute("myKey", "myVal"))},
		},
		"multiple attributes": {
			src: []wasmvmtypes.EventAttribute{{Key: "myKey", Value: "myVal"},
				{Key: "myOtherKey", Value: "myOtherVal"}},
			exp: sdk.Events{sdk.NewEvent("wasm",
				sdk.NewAttribute("_contract_address", myContract.String()),
				sdk.NewAttribute("myKey", "myVal"),
				sdk.NewAttribute("myOtherKey", "myOtherVal"))},
		},
		"without attributes": {
			exp: sdk.Events{sdk.NewEvent("wasm",
				sdk.NewAttribute("_contract_address", myContract.String()))},
		},
		"error on _contract_address": {
			src:     []wasmvmtypes.EventAttribute{{Key: "_contract_address", Value: RandomBech32AccountAddress(t)}},
			isError: true,
		},
	}
	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			gotEvent, err := newWasmModuleEvent(spec.src, myContract)
			if spec.isError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, spec.exp, gotEvent)
			}
		})
	}
}
