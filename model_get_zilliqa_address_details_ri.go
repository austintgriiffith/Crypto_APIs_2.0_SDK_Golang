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

// GetZilliqaAddressDetailsRI struct for GetZilliqaAddressDetailsRI
type GetZilliqaAddressDetailsRI struct {
	Balance GetZilliqaAddressDetailsRIBalance `json:"balance"`
	// Defines the received transaction count to the address.
	IncomingTransactionsCount int32 `json:"incomingTransactionsCount"`
	// Defines the sent transaction count from the address.
	OutgoingTransactionsCount int32 `json:"outgoingTransactionsCount"`
	// Defines the entire count of the transactions.
	TransactionsCount int32 `json:"transactionsCount"`
}

// NewGetZilliqaAddressDetailsRI instantiates a new GetZilliqaAddressDetailsRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetZilliqaAddressDetailsRI(balance GetZilliqaAddressDetailsRIBalance, incomingTransactionsCount int32, outgoingTransactionsCount int32, transactionsCount int32) *GetZilliqaAddressDetailsRI {
	this := GetZilliqaAddressDetailsRI{}
	this.Balance = balance
	this.IncomingTransactionsCount = incomingTransactionsCount
	this.OutgoingTransactionsCount = outgoingTransactionsCount
	this.TransactionsCount = transactionsCount
	return &this
}

// NewGetZilliqaAddressDetailsRIWithDefaults instantiates a new GetZilliqaAddressDetailsRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetZilliqaAddressDetailsRIWithDefaults() *GetZilliqaAddressDetailsRI {
	this := GetZilliqaAddressDetailsRI{}
	return &this
}

// GetBalance returns the Balance field value
func (o *GetZilliqaAddressDetailsRI) GetBalance() GetZilliqaAddressDetailsRIBalance {
	if o == nil {
		var ret GetZilliqaAddressDetailsRIBalance
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *GetZilliqaAddressDetailsRI) GetBalanceOk() (*GetZilliqaAddressDetailsRIBalance, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *GetZilliqaAddressDetailsRI) SetBalance(v GetZilliqaAddressDetailsRIBalance) {
	o.Balance = v
}

// GetIncomingTransactionsCount returns the IncomingTransactionsCount field value
func (o *GetZilliqaAddressDetailsRI) GetIncomingTransactionsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IncomingTransactionsCount
}

// GetIncomingTransactionsCountOk returns a tuple with the IncomingTransactionsCount field value
// and a boolean to check if the value has been set.
func (o *GetZilliqaAddressDetailsRI) GetIncomingTransactionsCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IncomingTransactionsCount, true
}

// SetIncomingTransactionsCount sets field value
func (o *GetZilliqaAddressDetailsRI) SetIncomingTransactionsCount(v int32) {
	o.IncomingTransactionsCount = v
}

// GetOutgoingTransactionsCount returns the OutgoingTransactionsCount field value
func (o *GetZilliqaAddressDetailsRI) GetOutgoingTransactionsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OutgoingTransactionsCount
}

// GetOutgoingTransactionsCountOk returns a tuple with the OutgoingTransactionsCount field value
// and a boolean to check if the value has been set.
func (o *GetZilliqaAddressDetailsRI) GetOutgoingTransactionsCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OutgoingTransactionsCount, true
}

// SetOutgoingTransactionsCount sets field value
func (o *GetZilliqaAddressDetailsRI) SetOutgoingTransactionsCount(v int32) {
	o.OutgoingTransactionsCount = v
}

// GetTransactionsCount returns the TransactionsCount field value
func (o *GetZilliqaAddressDetailsRI) GetTransactionsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TransactionsCount
}

// GetTransactionsCountOk returns a tuple with the TransactionsCount field value
// and a boolean to check if the value has been set.
func (o *GetZilliqaAddressDetailsRI) GetTransactionsCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionsCount, true
}

// SetTransactionsCount sets field value
func (o *GetZilliqaAddressDetailsRI) SetTransactionsCount(v int32) {
	o.TransactionsCount = v
}

func (o GetZilliqaAddressDetailsRI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["balance"] = o.Balance
	}
	if true {
		toSerialize["incomingTransactionsCount"] = o.IncomingTransactionsCount
	}
	if true {
		toSerialize["outgoingTransactionsCount"] = o.OutgoingTransactionsCount
	}
	if true {
		toSerialize["transactionsCount"] = o.TransactionsCount
	}
	return json.Marshal(toSerialize)
}

type NullableGetZilliqaAddressDetailsRI struct {
	value *GetZilliqaAddressDetailsRI
	isSet bool
}

func (v NullableGetZilliqaAddressDetailsRI) Get() *GetZilliqaAddressDetailsRI {
	return v.value
}

func (v *NullableGetZilliqaAddressDetailsRI) Set(val *GetZilliqaAddressDetailsRI) {
	v.value = val
	v.isSet = true
}

func (v NullableGetZilliqaAddressDetailsRI) IsSet() bool {
	return v.isSet
}

func (v *NullableGetZilliqaAddressDetailsRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetZilliqaAddressDetailsRI(val *GetZilliqaAddressDetailsRI) *NullableGetZilliqaAddressDetailsRI {
	return &NullableGetZilliqaAddressDetailsRI{value: val, isSet: true}
}

func (v NullableGetZilliqaAddressDetailsRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetZilliqaAddressDetailsRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


