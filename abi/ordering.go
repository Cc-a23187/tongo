package abi
// Code autogenerated. DO NOT EDIT. 

import (
"context"
"github.com/tonkeeper/tongo"
)


type ContractInterface string

// more wallet-related contract interfaces are defined in wallet.go
const (
	Auction             ContractInterface = "auction"
	Domain              ContractInterface = "domain"
	NftEditable         ContractInterface = "nft_editable"
	NftSale             ContractInterface = "nft_sale"
	NftSaleGetgems      ContractInterface = "nft_sale_getgems"
	PaymentChannel      ContractInterface = "payment_channel"
	StorageContract     ContractInterface = "storage_contract"
	StorageProvider     ContractInterface = "storage_provider"
	Subscription        ContractInterface = "subscription"
	Telemint            ContractInterface = "telemint"
	Tep62Collection     ContractInterface = "tep62_collection"
	Tep62Item           ContractInterface = "tep62_item"
	Tep66               ContractInterface = "tep66"
	Tep74               ContractInterface = "tep74"
	Tep85               ContractInterface = "tep85"
	TfNominator         ContractInterface = "tf_nominator"
	TonstakePool        ContractInterface = "tonstake_pool"
	ValidatorController ContractInterface = "validator_controller"
	Wallet              ContractInterface = "wallet"
	WalletV4R2          ContractInterface = "wallet_v4r2"
	WhalesNominators    ContractInterface = "whales_nominators"
)

type InvokeFn func(ctx context.Context, executor Executor, reqAccountID tongo.AccountID) (string, any, error)

// MethodDescription describes a particular method and provides a function to execute it.
type MethodDescription struct {
	Name string
	// InvokeFn executes this method on a contract and returns parsed execution results.
	InvokeFn InvokeFn
	// ImplementedBy is a list of contract interfaces that implement this method.
	// All contract interfaces share the same method with the same output type.
	ImplementedBy []ContractInterface
	// ImplementedByFn returns an implemented contract interface based on a type hint from InvokeFn.
	// Contract interfaces share the same method name but output is different for each contract interface.
	// Check GetSaleData out as an example.
	ImplementedByFn func(typeName string) ContractInterface
}

var methodInvocationOrder = []MethodDescription{
	{
		Name:          "get_auction_info",
		InvokeFn:      GetAuctionInfo,
		ImplementedBy: []ContractInterface{Auction},
	},
	{
		Name:          "get_authority_address",
		InvokeFn:      GetAuthorityAddress,
		ImplementedBy: []ContractInterface{Tep85},
	},
	{
		Name:          "get_channel_state",
		InvokeFn:      GetChannelState,
		ImplementedBy: []ContractInterface{PaymentChannel},
	},
	{
		Name:          "get_collection_data",
		InvokeFn:      GetCollectionData,
		ImplementedBy: []ContractInterface{Tep62Collection},
	},
	{
		Name:          "get_domain",
		InvokeFn:      GetDomain,
		ImplementedBy: []ContractInterface{Domain},
	},
	{
		Name:          "get_editor",
		InvokeFn:      GetEditor,
		ImplementedBy: []ContractInterface{NftEditable},
	},
	{
		Name:          "get_full_domain",
		InvokeFn:      GetFullDomain,
		ImplementedBy: []ContractInterface{Domain},
	},
	{
		Name:          "get_jetton_data",
		InvokeFn:      GetJettonData,
		ImplementedBy: []ContractInterface{Tep74},
	},
	{
		Name:          "get_last_fill_up_time",
		InvokeFn:      GetLastFillUpTime,
		ImplementedBy: []ContractInterface{Domain},
	},
	{
		Name:     "get_members_raw",
		InvokeFn: GetMembersRaw,
		ImplementedByFn: func(typeHint string) ContractInterface {
			switch typeHint {
			case "GetMembersRaw_WhalesNominatorResult":
				return WhalesNominators
			}
			return ""
		},
	},
	{
		Name:          "get_next_proof_info",
		InvokeFn:      GetNextProofInfo,
		ImplementedBy: []ContractInterface{StorageContract},
	},
	{
		Name:          "get_nft_data",
		InvokeFn:      GetNftData,
		ImplementedBy: []ContractInterface{Tep62Item},
	},
	{
		Name:     "get_params",
		InvokeFn: GetParams,
		ImplementedByFn: func(typeHint string) ContractInterface {
			switch typeHint {
			case "GetParams_WhalesNominatorResult":
				return WhalesNominators
			}
			return ""
		},
	},
	{
		Name:          "get_plugin_list",
		InvokeFn:      GetPluginList,
		ImplementedBy: []ContractInterface{WalletV4R2},
	},
	{
		Name:          "get_pool_data",
		InvokeFn:      GetPoolData,
		ImplementedBy: []ContractInterface{TfNominator},
	},
	{
		Name:          "get_pool_full_data",
		InvokeFn:      GetPoolFullData,
		ImplementedBy: []ContractInterface{TonstakePool},
	},
	{
		Name:          "get_pool_status",
		InvokeFn:      GetPoolStatus,
		ImplementedBy: []ContractInterface{WhalesNominators},
	},
	{
		Name:          "get_public_key",
		InvokeFn:      GetPublicKey,
		ImplementedBy: []ContractInterface{StorageProvider, Wallet},
	},
	{
		Name:          "get_revoked_time",
		InvokeFn:      GetRevokedTime,
		ImplementedBy: []ContractInterface{Tep85},
	},
	{
		Name:     "get_sale_data",
		InvokeFn: GetSaleData,
		ImplementedByFn: func(typeHint string) ContractInterface {
			switch typeHint {
			case "GetSaleData_BasicResult":
				return NftSale
			case "GetSaleData_GetgemsAuctionResult":
				return NftSaleGetgems
			case "GetSaleData_GetgemsResult":
				return NftSaleGetgems
			}
			return ""
		},
	},
	{
		Name:          "get_staking_status",
		InvokeFn:      GetStakingStatus,
		ImplementedBy: []ContractInterface{WhalesNominators},
	},
	{
		Name:          "get_storage_contract_data",
		InvokeFn:      GetStorageContractData,
		ImplementedBy: []ContractInterface{StorageContract},
	},
	{
		Name:          "get_storage_params",
		InvokeFn:      GetStorageParams,
		ImplementedBy: []ContractInterface{StorageProvider},
	},
	{
		Name:          "get_subscription_data",
		InvokeFn:      GetSubscriptionData,
		ImplementedBy: []ContractInterface{Subscription},
	},
	{
		Name:          "get_subwallet_id",
		InvokeFn:      GetSubwalletId,
		ImplementedBy: []ContractInterface{WalletV4R2},
	},
	{
		Name:          "get_telemint_auction_config",
		InvokeFn:      GetTelemintAuctionConfig,
		ImplementedBy: []ContractInterface{Telemint},
	},
	{
		Name:          "get_telemint_auction_state",
		InvokeFn:      GetTelemintAuctionState,
		ImplementedBy: []ContractInterface{Telemint},
	},
	{
		Name:          "get_telemint_token_name",
		InvokeFn:      GetTelemintTokenName,
		ImplementedBy: []ContractInterface{Telemint},
	},
	{
		Name:          "get_torrent_hash",
		InvokeFn:      GetTorrentHash,
		ImplementedBy: []ContractInterface{StorageContract},
	},
	{
		Name:          "get_validator_controller_data",
		InvokeFn:      GetValidatorControllerData,
		ImplementedBy: []ContractInterface{ValidatorController},
	},
	{
		Name:          "get_wallet_data",
		InvokeFn:      GetWalletData,
		ImplementedBy: []ContractInterface{Tep74},
	},
	{
		Name:          "get_wallet_params",
		InvokeFn:      GetWalletParams,
		ImplementedBy: []ContractInterface{StorageProvider},
	},
	{
		Name:          "is_active",
		InvokeFn:      IsActive,
		ImplementedBy: []ContractInterface{StorageContract},
	},
	{
		Name:          "list_nominators",
		InvokeFn:      ListNominators,
		ImplementedBy: []ContractInterface{TfNominator},
	},
	{
		Name:          "list_votes",
		InvokeFn:      ListVotes,
		ImplementedBy: []ContractInterface{TfNominator},
	},
	{
		Name:          "royalty_params",
		InvokeFn:      RoyaltyParams,
		ImplementedBy: []ContractInterface{Tep66},
	},
	{
		Name:          "seqno",
		InvokeFn:      Seqno,
		ImplementedBy: []ContractInterface{StorageProvider, Wallet},
	},
}
