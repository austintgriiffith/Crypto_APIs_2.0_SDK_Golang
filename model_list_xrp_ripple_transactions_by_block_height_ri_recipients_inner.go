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

// ListXRPRippleTransactionsByBlockHeightRIRecipientsInner struct for ListXRPRippleTransactionsByBlockHeightRIRecipientsInner
type ListXRPRippleTransactionsByBlockHeightRIRecipientsInner struct {
	// String representation of the receiver address
	Address string `json:"address"`
	// String representation of the amount
	Amount string `json:"amount"`
}

// NewListXRPRippleTransactionsByBlockHeightRIRecipientsInner instantiates a new ListXRPRippleTransactionsByBlockHeightRIRecipientsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListXRPRippleTransactionsByBlockHeightRIRecipientsInner(address string, amount string) *ListXRPRippleTransactionsByBlockHeightRIRecipientsInner {
	this := ListXRPRippleTransactionsByBlockHeightRIRecipientsInner{}
	this.Address = address
	this.Amount = amount
	return &this
}

// NewListXRPRippleTransactionsByBlockHeightRIRecipientsInnerWithDefaults instantiates a new ListXRPRippleTransactionsByBlockHeightRIRecipientsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListXRPRippleTransactionsByBlockHeightRIRecipientsInnerWithDefaults() *ListXRPRippleTransactionsByBlockHeightRIRecipientsInner {
	this := ListXRPRippleTransactionsByBlockHeightRIRecipientsInner{}
	return &this
}

// GetAddress returns the Address field value
func (o *ListXRPRippleTransactionsByBlockHeightRIRecipientsInner) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByBlockHeightRIRecipientsInner) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *ListXRPRippleTransactionsByBlockHeightRIRecipientsInner) SetAddress(v string) {
	o.Address = v
}

// GetAmount returns the Amount field value
func (o *ListXRPRippleTransactionsByBlockHeightRIRecipientsInner) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByBlockHeightRIRecipientsInner) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ListXRPRippleTransactionsByBlockHeightRIRecipientsInner) SetAmount(v string) {
	o.Amount = v
}

func (o ListXRPRippleTransactionsByBlockHeightRIRecipientsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	return json.Marshal(toSerialize)
}

type NullableListXRPRippleTransactionsByBlockHeightRIRecipientsInner struct {
	value *ListXRPRippleTransactionsByBlockHeightRIRecipientsInner
	isSet bool
}

func (v NullableListXRPRippleTransactionsByBlockHeightRIRecipientsInner) Get() *ListXRPRippleTransactionsByBlockHeightRIRecipientsInner {
	return v.value
}

func (v *NullableListXRPRippleTransactionsByBlockHeightRIRecipientsInner) Set(val *ListXRPRippleTransactionsByBlockHeightRIRecipientsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListXRPRippleTransactionsByBlockHeightRIRecipientsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListXRPRippleTransactionsByBlockHeightRIRecipientsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListXRPRippleTransactionsByBlockHeightRIRecipientsInner(val *ListXRPRippleTransactionsByBlockHeightRIRecipientsInner) *NullableListXRPRippleTransactionsByBlockHeightRIRecipientsInner {
	return &NullableListXRPRippleTransactionsByBlockHeightRIRecipientsInner{value: val, isSet: true}
}

func (v NullableListXRPRippleTransactionsByBlockHeightRIRecipientsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListXRPRippleTransactionsByBlockHeightRIRecipientsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

