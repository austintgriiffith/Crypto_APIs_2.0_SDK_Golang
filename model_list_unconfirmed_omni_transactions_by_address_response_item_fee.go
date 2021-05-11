/*
 * CryptoAPIs
 *
 * Crypto APIs 2.0 is a complex and innovative infrastructure layer that radically simplifies the development of any Blockchain and Crypto related applications. Organized around REST, Crypto APIs 2.0 can assist both novice Bitcoin/Ethereum enthusiasts and crypto experts with the development of their blockchain applications. Crypto APIs 2.0 provides unified endpoints and data, raw data, automatic tokens and coins forwardings, callback functionalities, and much more.
 *
 * API version: 2.0.0
 * Contact: developers@cryptoapis.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cryptoapis

import (
	"encoding/json"
)

// ListUnconfirmedOmniTransactionsByAddressResponseItemFee struct for ListUnconfirmedOmniTransactionsByAddressResponseItemFee
type ListUnconfirmedOmniTransactionsByAddressResponseItemFee struct {
	// Defines the amount of the fee.
	Amount string `json:"amount"`
	// Defines the unit of the fee.
	Unit string `json:"unit"`
}

// NewListUnconfirmedOmniTransactionsByAddressResponseItemFee instantiates a new ListUnconfirmedOmniTransactionsByAddressResponseItemFee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUnconfirmedOmniTransactionsByAddressResponseItemFee(amount string, unit string) *ListUnconfirmedOmniTransactionsByAddressResponseItemFee {
	this := ListUnconfirmedOmniTransactionsByAddressResponseItemFee{}
	this.Amount = amount
	this.Unit = unit
	return &this
}

// NewListUnconfirmedOmniTransactionsByAddressResponseItemFeeWithDefaults instantiates a new ListUnconfirmedOmniTransactionsByAddressResponseItemFee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUnconfirmedOmniTransactionsByAddressResponseItemFeeWithDefaults() *ListUnconfirmedOmniTransactionsByAddressResponseItemFee {
	this := ListUnconfirmedOmniTransactionsByAddressResponseItemFee{}
	return &this
}

// GetAmount returns the Amount field value
func (o *ListUnconfirmedOmniTransactionsByAddressResponseItemFee) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedOmniTransactionsByAddressResponseItemFee) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ListUnconfirmedOmniTransactionsByAddressResponseItemFee) SetAmount(v string) {
	o.Amount = v
}

// GetUnit returns the Unit field value
func (o *ListUnconfirmedOmniTransactionsByAddressResponseItemFee) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedOmniTransactionsByAddressResponseItemFee) GetUnitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *ListUnconfirmedOmniTransactionsByAddressResponseItemFee) SetUnit(v string) {
	o.Unit = v
}

func (o ListUnconfirmedOmniTransactionsByAddressResponseItemFee) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableListUnconfirmedOmniTransactionsByAddressResponseItemFee struct {
	value *ListUnconfirmedOmniTransactionsByAddressResponseItemFee
	isSet bool
}

func (v NullableListUnconfirmedOmniTransactionsByAddressResponseItemFee) Get() *ListUnconfirmedOmniTransactionsByAddressResponseItemFee {
	return v.value
}

func (v *NullableListUnconfirmedOmniTransactionsByAddressResponseItemFee) Set(val *ListUnconfirmedOmniTransactionsByAddressResponseItemFee) {
	v.value = val
	v.isSet = true
}

func (v NullableListUnconfirmedOmniTransactionsByAddressResponseItemFee) IsSet() bool {
	return v.isSet
}

func (v *NullableListUnconfirmedOmniTransactionsByAddressResponseItemFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUnconfirmedOmniTransactionsByAddressResponseItemFee(val *ListUnconfirmedOmniTransactionsByAddressResponseItemFee) *NullableListUnconfirmedOmniTransactionsByAddressResponseItemFee {
	return &NullableListUnconfirmedOmniTransactionsByAddressResponseItemFee{value: val, isSet: true}
}

func (v NullableListUnconfirmedOmniTransactionsByAddressResponseItemFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUnconfirmedOmniTransactionsByAddressResponseItemFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

