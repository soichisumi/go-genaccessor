package example

import (
	"encoding"

	adamant "github.com/GincoInc/protobuf/gen/go/gincoinc/adamant/constant/v1"
	constant "github.com/GincoInc/protobuf/gen/go/gincoinc/constant/v1"
)

//go:generate go-genaccessor

type Person struct {
	id   string                 `getter:""`
	name string                 `getter:"" setter:"Rename"`
	tags []string               `getter:"" setter:""`
	text encoding.TextMarshaler `getter:"" setter:""`
}

type Wallet struct {
	walletType  adamant.WalletType   `getter:""`
	addressType constant.AddressType `getter:""`
}
