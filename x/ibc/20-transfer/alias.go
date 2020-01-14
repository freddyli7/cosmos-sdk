package transfer

// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/ibc/20-transfer/keeper
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/ibc/20-transfer/types

import (
	"github.com/cosmos/cosmos-sdk/x/ibc/20-transfer/keeper"
	"github.com/cosmos/cosmos-sdk/x/ibc/20-transfer/types"
)

const (
	DefaultPacketTimeout = keeper.DefaultPacketTimeout
	AttributeKeyReceiver = types.AttributeKeyReceiver
	SubModuleName        = types.SubModuleName
	StoreKey             = types.StoreKey
	RouterKey            = types.RouterKey
	QuerierRoute         = types.QuerierRoute
)

var (
	// functions aliases
	NewKeeper            = keeper.NewKeeper
	RegisterCodec        = types.RegisterCodec
	GetEscrowAddress     = types.GetEscrowAddress
	GetDenomPrefix       = types.GetDenomPrefix
	GetModuleAccountName = types.GetModuleAccountName
	NewMsgTransfer       = types.NewMsgTransfer

	// variable aliases
	ModuleCdc              = types.ModuleCdc
	AttributeValueCategory = types.AttributeValueCategory
)

type (
	Keeper             = keeper.Keeper
	BankKeeper         = types.BankKeeper
	ChannelKeeper      = types.ChannelKeeper
	ClientKeeper       = types.ClientKeeper
	ConnectionKeeper   = types.ConnectionKeeper
	SupplyKeeper       = types.SupplyKeeper
	MsgTransfer        = types.MsgTransfer
	MsgRecvPacket      = types.MsgRecvPacket
	PacketDataTransfer = types.PacketDataTransfer
)
