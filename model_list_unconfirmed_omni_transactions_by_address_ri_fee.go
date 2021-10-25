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

// ListUnconfirmedOmniTransactionsByAddressRIFee struct for ListUnconfirmedOmniTransactionsByAddressRIFee
type ListUnconfirmedOmniTransactionsByAddressRIFee struct {
	// Defines the amount of the fee.
	Amount string `json:"amount"`
	// Defines the unit of the fee.
	Unit string `json:"unit"`
}

// NewListUnconfirmedOmniTransactionsByAddressRIFee instantiates a new ListUnconfirmedOmniTransactionsByAddressRIFee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUnconfirmedOmniTransactionsByAddressRIFee(amount string, unit string) *ListUnconfirmedOmniTransactionsByAddressRIFee {
	this := ListUnconfirmedOmniTransactionsByAddressRIFee{}
	this.Amount = amount
	this.Unit = unit
	return &this
}

// NewListUnconfirmedOmniTransactionsByAddressRIFeeWithDefaults instantiates a new ListUnconfirmedOmniTransactionsByAddressRIFee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUnconfirmedOmniTransactionsByAddressRIFeeWithDefaults() *ListUnconfirmedOmniTransactionsByAddressRIFee {
	this := ListUnconfirmedOmniTransactionsByAddressRIFee{}
	return &this
}

// GetAmount returns the Amount field value
func (o *ListUnconfirmedOmniTransactionsByAddressRIFee) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedOmniTransactionsByAddressRIFee) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ListUnconfirmedOmniTransactionsByAddressRIFee) SetAmount(v string) {
	o.Amount = v
}

// GetUnit returns the Unit field value
func (o *ListUnconfirmedOmniTransactionsByAddressRIFee) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedOmniTransactionsByAddressRIFee) GetUnitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *ListUnconfirmedOmniTransactionsByAddressRIFee) SetUnit(v string) {
	o.Unit = v
}

func (o ListUnconfirmedOmniTransactionsByAddressRIFee) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableListUnconfirmedOmniTransactionsByAddressRIFee struct {
	value *ListUnconfirmedOmniTransactionsByAddressRIFee
	isSet bool
}

func (v NullableListUnconfirmedOmniTransactionsByAddressRIFee) Get() *ListUnconfirmedOmniTransactionsByAddressRIFee {
	return v.value
}

func (v *NullableListUnconfirmedOmniTransactionsByAddressRIFee) Set(val *ListUnconfirmedOmniTransactionsByAddressRIFee) {
	v.value = val
	v.isSet = true
}

func (v NullableListUnconfirmedOmniTransactionsByAddressRIFee) IsSet() bool {
	return v.isSet
}

func (v *NullableListUnconfirmedOmniTransactionsByAddressRIFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUnconfirmedOmniTransactionsByAddressRIFee(val *ListUnconfirmedOmniTransactionsByAddressRIFee) *NullableListUnconfirmedOmniTransactionsByAddressRIFee {
	return &NullableListUnconfirmedOmniTransactionsByAddressRIFee{value: val, isSet: true}
}

func (v NullableListUnconfirmedOmniTransactionsByAddressRIFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUnconfirmedOmniTransactionsByAddressRIFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


