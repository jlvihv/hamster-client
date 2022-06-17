package pallet

import (
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

type EventFreeResourceApplied struct {
	Phase      types.Phase
	OrderIndex types.U64
	PeerId     string
	Topics     []types.Hash
}

type EventBalance struct {
	Phase     types.Phase
	AccountId types.AccountID
	Amount    types.U128
	Topics    []types.Hash
}

type MyEventRecords struct {
	types.EventRecords
	ResourceOrder_FreeResourceProcessed []EventFreeResourceApplied
	Balances_Withdraw                   []EventBalance //org approve event
}
