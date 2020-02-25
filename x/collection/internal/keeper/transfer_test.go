package keeper

import (
	"testing"

	"github.com/line/link/x/collection/internal/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestTransferFTScenario(t *testing.T) {
	ctx := cacheKeeper()

	// issue idf token
	require.NoError(t, keeper.CreateCollection(ctx, types.NewCollection(defaultSymbol, defaultName), addr1))
	collection, err := keeper.GetCollection(ctx, defaultSymbol)
	require.NoError(t, err)
	err = keeper.IssueFT(ctx, addr1, types.NewFT(collection, defaultName, defaultTokenURI, sdk.NewInt(defaultDecimals), true), sdk.NewInt(defaultAmount))
	require.NoError(t, err)

	//
	// transfer success cases
	//
	require.NoError(t, keeper.TransferFT(ctx, addr1, addr2, defaultSymbol, defaultTokenIDFT, sdk.NewInt(defaultAmount/2)))

	//
	// transfer failure cases
	//
	// Insufficient coins
	require.EqualError(t, keeper.TransferFT(ctx, addr1, addr2, defaultSymbol, defaultTokenIDFT, sdk.NewInt(defaultAmount)), sdk.ErrInsufficientCoins("insufficient account funds; 500token00100010000 < 1000token00100010000").Error())
}

func TestTransferNFTScenario(t *testing.T) {
	ctx := cacheKeeper()
	prepareCollectionTokens(ctx, t)

	require.NoError(t, keeper.Attach(ctx, addr1, defaultSymbol, defaultTokenID1, defaultTokenID2))
	require.NoError(t, keeper.Attach(ctx, addr1, defaultSymbol, defaultTokenID2, defaultTokenID3))
	require.NoError(t, keeper.Attach(ctx, addr1, defaultSymbol, defaultTokenID1, defaultTokenID4))

	//
	// transfer failure cases
	//

	// transfer non-exist token : failure
	require.EqualError(t, keeper.TransferNFT(ctx, addr1, addr2, defaultSymbol, defaultTokenID8), types.ErrCollectionTokenNotExist(types.DefaultCodespace, defaultSymbol, defaultTokenID8).Error())

	// transfer a child : failure
	require.EqualError(t, keeper.TransferNFT(ctx, addr1, addr2, defaultSymbol, defaultTokenID2), types.ErrTokenCannotTransferChildToken(types.DefaultCodespace, defaultSymbol+defaultTokenID2).Error())
	require.EqualError(t, keeper.TransferNFT(ctx, addr1, addr2, defaultSymbol, defaultTokenID3), types.ErrTokenCannotTransferChildToken(types.DefaultCodespace, defaultSymbol+defaultTokenID3).Error())
	require.EqualError(t, keeper.TransferNFT(ctx, addr1, addr2, defaultSymbol, defaultTokenID4), types.ErrTokenCannotTransferChildToken(types.DefaultCodespace, defaultSymbol+defaultTokenID4).Error())

	// transfer non-mine : failure
	require.EqualError(t, keeper.TransferNFT(ctx, addr1, addr2, defaultSymbol, defaultTokenID5), types.ErrTokenNotOwnedBy(types.DefaultCodespace, defaultSymbol+defaultTokenID5, addr1).Error())

	// transfer-nft ft : failure
	require.EqualError(t, keeper.TransferNFT(ctx, addr1, addr2, defaultSymbol, defaultTokenIDFT), types.ErrTokenNotNFT(types.DefaultCodespace, defaultSymbol+defaultTokenIDFT).Error())

	//
	// transfer success cases
	//
	require.NoError(t, keeper.TransferNFT(ctx, addr1, addr2, defaultSymbol, defaultTokenID1))
	require.NoError(t, keeper.TransferNFT(ctx, addr2, addr1, defaultSymbol, defaultTokenID1))
	require.NoError(t, keeper.TransferNFT(ctx, addr1, addr2, defaultSymbol, defaultTokenID1))

	// verify the owner of transferred tokens
	// owner of token1 is addr2
	token1, err1 := keeper.GetToken(ctx, defaultSymbol, defaultTokenID1)
	require.NoError(t, err1)
	require.Equal(t, token1.(types.NFT).GetOwner(), addr2)

	// owner of token2 is addr2
	token2, err2 := keeper.GetToken(ctx, defaultSymbol, defaultTokenID2)
	require.NoError(t, err2)
	require.Equal(t, token2.(types.NFT).GetOwner(), addr2)

	// owner of token3 is addr2
	token3, err3 := keeper.GetToken(ctx, defaultSymbol, defaultTokenID3)
	require.NoError(t, err3)
	require.Equal(t, token3.(types.NFT).GetOwner(), addr2)

	// owner of token4 is addr2
	token4, err4 := keeper.GetToken(ctx, defaultSymbol, defaultTokenID4)
	require.NoError(t, err4)
	require.Equal(t, token4.(types.NFT).GetOwner(), addr2)
}
