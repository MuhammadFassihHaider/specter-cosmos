package keeper

import (
	"context"

	"specter/x/specter/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (k msgServer) RequestLoan(goCtx context.Context, msg *types.MsgRequestLoan) (*types.MsgRequestLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var loan = types.Loan{
		Amount:     msg.Amount,
		Fee:        msg.Fee,
		Collateral: msg.Collateral,
		Deadline:   msg.Deadline,
		State:      "requested",
		Borrower:   msg.Creator,
	}

	usersWithAbilityToTransact := []string{
		"cosmos1dmj4q5p62vq9fpum5m2jg54sqv4p7qy4qdefhf", "cosmos1metcp63eg595s7jd3ja089uh70qh6wh8mkznps", "cosmos1sxky6hx50uxs3gdtmmqx4nntgjv5xxnj0w3gne",
	}

	if !stringInSlice(msg.Creator, usersWithAbilityToTransact) {
		panic("Not allowed!")
	}

	borrower, _ := sdk.AccAddressFromBech32(msg.Creator)

	// Get the collateral as sdk.Coins
	collateral, err := sdk.ParseCoinsNormalized(loan.Collateral)
	if err != nil {
		panic(err)
	}

	// Use the module account as escrow account
	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, borrower, types.ModuleName, collateral)
	if sdkError != nil {
		return nil, sdkError
	}

	// Add the loan to the keeper
	k.AppendLoan(
		ctx,
		loan,
	)

	return &types.MsgRequestLoanResponse{}, nil
}
