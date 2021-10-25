/*
CryptoAPIs

Crypto APIs 2.0 is a complex and innovative infrastructure layer that radically simplifies the development of any Blockchain and Crypto related applications. Organized around REST, Crypto APIs 2.0 can assist both novice Bitcoin/Ethereum enthusiasts and crypto experts with the development of their blockchain applications. Crypto APIs 2.0 provides unified endpoints and data, raw data, automatic tokens and coins forwardings, callback functionalities, and much more.

API version: 2.0.0
Contact: developers@cryptoapis.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cryptoapis

import (
	"encoding/json"
)

// TransactionRequestRejectionDataItem Defines an `item` as one result.
type TransactionRequestRejectionDataItem struct {
	// Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc.
	Blockchain string `json:"blockchain"`
	// Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\", \"rinkeby\" are test networks.
	Network string `json:"network"`
	// The required number of approvals needed to approve the transaction.
	RequiredApprovals int32 `json:"requiredApprovals"`
	// The required number of rejections needed to reject the transaction.
	RequiredRejections int32 `json:"requiredRejections"`
	// The current number of approvals given for the transaction.
	CurrentApprovals int32 `json:"currentApprovals"`
	// The current number of rejections given for the transaction.
	CurrentRejections int32 `json:"currentRejections"`
}

// NewTransactionRequestRejectionDataItem instantiates a new TransactionRequestRejectionDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionRequestRejectionDataItem(blockchain string, network string, requiredApprovals int32, requiredRejections int32, currentApprovals int32, currentRejections int32) *TransactionRequestRejectionDataItem {
	this := TransactionRequestRejectionDataItem{}
	this.Blockchain = blockchain
	this.Network = network
	this.RequiredApprovals = requiredApprovals
	this.RequiredRejections = requiredRejections
	this.CurrentApprovals = currentApprovals
	this.CurrentRejections = currentRejections
	return &this
}

// NewTransactionRequestRejectionDataItemWithDefaults instantiates a new TransactionRequestRejectionDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionRequestRejectionDataItemWithDefaults() *TransactionRequestRejectionDataItem {
	this := TransactionRequestRejectionDataItem{}
	return &this
}

// GetBlockchain returns the Blockchain field value
func (o *TransactionRequestRejectionDataItem) GetBlockchain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Blockchain
}

// GetBlockchainOk returns a tuple with the Blockchain field value
// and a boolean to check if the value has been set.
func (o *TransactionRequestRejectionDataItem) GetBlockchainOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Blockchain, true
}

// SetBlockchain sets field value
func (o *TransactionRequestRejectionDataItem) SetBlockchain(v string) {
	o.Blockchain = v
}

// GetNetwork returns the Network field value
func (o *TransactionRequestRejectionDataItem) GetNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *TransactionRequestRejectionDataItem) GetNetworkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *TransactionRequestRejectionDataItem) SetNetwork(v string) {
	o.Network = v
}

// GetRequiredApprovals returns the RequiredApprovals field value
func (o *TransactionRequestRejectionDataItem) GetRequiredApprovals() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RequiredApprovals
}

// GetRequiredApprovalsOk returns a tuple with the RequiredApprovals field value
// and a boolean to check if the value has been set.
func (o *TransactionRequestRejectionDataItem) GetRequiredApprovalsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequiredApprovals, true
}

// SetRequiredApprovals sets field value
func (o *TransactionRequestRejectionDataItem) SetRequiredApprovals(v int32) {
	o.RequiredApprovals = v
}

// GetRequiredRejections returns the RequiredRejections field value
func (o *TransactionRequestRejectionDataItem) GetRequiredRejections() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RequiredRejections
}

// GetRequiredRejectionsOk returns a tuple with the RequiredRejections field value
// and a boolean to check if the value has been set.
func (o *TransactionRequestRejectionDataItem) GetRequiredRejectionsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequiredRejections, true
}

// SetRequiredRejections sets field value
func (o *TransactionRequestRejectionDataItem) SetRequiredRejections(v int32) {
	o.RequiredRejections = v
}

// GetCurrentApprovals returns the CurrentApprovals field value
func (o *TransactionRequestRejectionDataItem) GetCurrentApprovals() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CurrentApprovals
}

// GetCurrentApprovalsOk returns a tuple with the CurrentApprovals field value
// and a boolean to check if the value has been set.
func (o *TransactionRequestRejectionDataItem) GetCurrentApprovalsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CurrentApprovals, true
}

// SetCurrentApprovals sets field value
func (o *TransactionRequestRejectionDataItem) SetCurrentApprovals(v int32) {
	o.CurrentApprovals = v
}

// GetCurrentRejections returns the CurrentRejections field value
func (o *TransactionRequestRejectionDataItem) GetCurrentRejections() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CurrentRejections
}

// GetCurrentRejectionsOk returns a tuple with the CurrentRejections field value
// and a boolean to check if the value has been set.
func (o *TransactionRequestRejectionDataItem) GetCurrentRejectionsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CurrentRejections, true
}

// SetCurrentRejections sets field value
func (o *TransactionRequestRejectionDataItem) SetCurrentRejections(v int32) {
	o.CurrentRejections = v
}

func (o TransactionRequestRejectionDataItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["blockchain"] = o.Blockchain
	}
	if true {
		toSerialize["network"] = o.Network
	}
	if true {
		toSerialize["requiredApprovals"] = o.RequiredApprovals
	}
	if true {
		toSerialize["requiredRejections"] = o.RequiredRejections
	}
	if true {
		toSerialize["currentApprovals"] = o.CurrentApprovals
	}
	if true {
		toSerialize["currentRejections"] = o.CurrentRejections
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionRequestRejectionDataItem struct {
	value *TransactionRequestRejectionDataItem
	isSet bool
}

func (v NullableTransactionRequestRejectionDataItem) Get() *TransactionRequestRejectionDataItem {
	return v.value
}

func (v *NullableTransactionRequestRejectionDataItem) Set(val *TransactionRequestRejectionDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionRequestRejectionDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionRequestRejectionDataItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionRequestRejectionDataItem(val *TransactionRequestRejectionDataItem) *NullableTransactionRequestRejectionDataItem {
	return &NullableTransactionRequestRejectionDataItem{value: val, isSet: true}
}

func (v NullableTransactionRequestRejectionDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionRequestRejectionDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

