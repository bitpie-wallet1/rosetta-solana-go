// Copyright 2020 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package solanago

import (
	"github.com/coinbase/rosetta-sdk-go/types"
	bin "github.com/dfuse-io/binary"
	"github.com/dfuse-io/solana-go"
	dfuserpc "github.com/dfuse-io/solana-go/rpc"
	"github.com/ethereum/go-ethereum/params"
)

const (
	// NodeVersion is the version of geth we are using.
	NodeVersion = "1.4.17"

	// Blockchain is solanago.
	Blockchain string = "solana"

	// MainnetNetwork is the value of the network
	// in MainnetNetworkIdentifier.
	MainnetNetwork string = "mainnet"

	// TestnetNetwork is the value of the network
	// in TestnetNetworkIdentifier.
	TestnetNetwork string = "devnet"

	// Symbol is the symbol value
	// used in Currency.
	Symbol = "SOL"

	// Decimals is the decimals value
	// used in Currency.
	Decimals = 9

	// SuccessStatus is the status of any
	// Ethereum operation considered successful.
	SuccessStatus = "SUCCESS"

	// FailureStatus is the status of any
	// Ethereum operation considered unsuccessful.
	FailureStatus = "FAILURE"

	// HistoricalBalanceSupported is whether
	// historical balance is supported.
	HistoricalBalanceSupported = true

	// GenesisBlockIndex is the index of the
	// genesis block.
	GenesisBlockIndex = int64(0)

	Separator = "__"
)

//op types

const (
	System__Transfer   = "System__Transfer"
	SplToken__Transfer = "SplToken__Transfer"
	Unknown            = "Unknown"
)

var (
	// MainnetGenesisBlockIdentifier is the *types.BlockIdentifier
	// of the mainnet genesis block.
	MainnetGenesisBlockIdentifier = &types.BlockIdentifier{
		Hash:  params.MainnetGenesisHash.Hex(),
		Index: GenesisBlockIndex,
	}

	// TestnetGenesisBlockIdentifier is the *types.BlockIdentifier
	// of the testnet genesis block.
	TestnetGenesisBlockIdentifier = &types.BlockIdentifier{
		Hash:  params.RopstenGenesisHash.Hex(),
		Index: GenesisBlockIndex,
	}

	// Currency is the *types.Currency for all
	// Ethereum networks.
	Currency = &types.Currency{
		Symbol:   Symbol,
		Decimals: Decimals,
	}

	// OperationTypes are all suppoorted operation types.
	OperationTypes = []string{
		System__Transfer,
		SplToken__Transfer,
		Unknown,
	}

	// OperationStatuses are all supported operation statuses.
	OperationStatuses = []*types.OperationStatus{
		{
			Status:     SuccessStatus,
			Successful: true,
		},
		{
			Status:     FailureStatus,
			Successful: false,
		},
	}

	// CallMethods are all supported call methods.
	CallMethods = []string{
		"getProgramAccounts",
		"getClusterNodes",
	}
)

type TokenParsed struct {
	Decimals        uint64
	Amount          uint64
	MintAutority    solana.PublicKey
	FreezeAuthority solana.PublicKey
	AuthorityType   solana.PublicKey
	NewAuthority    solana.PublicKey
	M               byte
}

type ParsedInstructionMeta struct {
	Authority    string            `json:authority,omitempty`
	NewAuthority string            `json:newAuthority,omitempty`
	Source       string            `json:source,omitempty`
	Destination  string            `json:destination,omitempty`
	Mint         string            `json:mint,omitempty`
	Decimals     uint8             `json:decimals,omitempty`
	TokenAmount  OpMetaTokenAmount `json:tokenAmount,omitempty`
	Amount       string            `json:amount,omitempty`
	Lamports     uint64            `json:lamports,omitempty`
	Space        uint64            `json:space,omitempty`
}
type OpMetaTokenAmount struct {
	Amount   string  `json:amount,omitempty`
	Decimals uint64  `json:decimals,omitempty`
	UiAmount float64 `json:uiAmount,omitempty`
}

type GetConfirmedBlockResult struct {
	Blockhash         solana.PublicKey             `json:"blockhash"`
	PreviousBlockhash solana.PublicKey             `json:"previousBlockhash"` // could be zeroes if ledger was clean-up and this is unavailable
	ParentSlot        bin.Uint64                   `json:"parentSlot"`
	Transactions      []dfuserpc.TransactionParsed `json:"transactions"`
	Rewards           []dfuserpc.BlockReward       `json:"rewards"`
	BlockTime         bin.Uint64                   `json:"blockTime,omitempty"`
}
