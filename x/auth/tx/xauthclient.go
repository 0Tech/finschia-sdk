// Package tx 's xauthclient.go file is copy-pasted from
// https://github.com/line/lbm-sdk/blob/v0.41.3/x/auth/client/query.go
// It is duplicated as to not introduce any breaking change in 0.41.4, see PR:
// https://github.com/line/lbm-sdk/pull/8732#discussion_r584746947
package tx

import (
	"context"
	"encoding/hex"
	"errors"
	"strings"

	ctypes "github.com/line/ostracon/rpc/core/types"

	"github.com/line/lbm-sdk/client"
	sdk "github.com/line/lbm-sdk/types"
)

// QueryTxsByEvents performs a search for transactions for a given set of events
// via the Tendermint RPC. An event takes the form of:
// "{eventAttribute}.{attributeKey} = '{attributeValue}'". Each event is
// concatenated with an 'AND' operand. It returns a slice of Info object
// containing txs and metadata. An error is returned if the query fails.
// If an empty string is provided it will order txs by asc
func queryTxsByEvents(goCtx context.Context, clientCtx client.Context, events []string, prove bool, page, limit int, orderBy string) (*sdk.SearchTxsResult, error) {
	if len(events) == 0 {
		return nil, errors.New("must declare at least one event to search")
	}

	if page <= 0 {
		return nil, errors.New("page must greater than 0")
	}

	if limit <= 0 {
		return nil, errors.New("limit must greater than 0")
	}

	// XXX: implement ANY
	query := strings.Join(events, " AND ")

	node, err := clientCtx.GetNode()
	if err != nil {
		return nil, err
	}

	resTxs, err := node.TxSearch(goCtx, query, prove, &page, &limit, orderBy)
	if err != nil {
		return nil, err
	}

	resBlocks, err := getBlocksForTxResults(clientCtx, resTxs.Txs)
	if err != nil {
		return nil, err
	}

	txs, err := formatTxResults(clientCtx.TxConfig, resTxs.Txs, resBlocks)
	if err != nil {
		return nil, err
	}

	result := sdk.NewSearchTxsResult(uint64(resTxs.TotalCount), uint64(len(txs)), uint64(page), uint64(limit), txs)

	return result, nil
}

// QueryTx queries for a single transaction by a hash string in hex format. An
// error is returned if the transaction does not exist or cannot be queried.
func queryTx(goCtx context.Context, clientCtx client.Context, hashHexStr string) (*sdk.TxResponse, error) {
	hash, err := hex.DecodeString(hashHexStr)
	if err != nil {
		return nil, err
	}

	node, err := clientCtx.GetNode()
	if err != nil {
		return nil, err
	}

	//TODO: this may not always need to be proven
	// https://github.com/line/lbm-sdk/issues/6807
	resTx, err := node.Tx(goCtx, hash, true)
	if err != nil {
		return nil, err
	}

	resBlocks, err := getBlocksForTxResults(clientCtx, []*ctypes.ResultTx{resTx})
	if err != nil {
		return nil, err
	}

	out, err := mkTxResult(clientCtx.TxConfig, resTx, resBlocks[resTx.Height])
	if err != nil {
		return out, err
	}

	return out, nil
}
