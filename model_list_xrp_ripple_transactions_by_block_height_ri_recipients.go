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

// ListXRPRippleTransactionsByBlockHeightRIRecipients struct for ListXRPRippleTransactionsByBlockHeightRIRecipients
type ListXRPRippleTransactionsByBlockHeightRIRecipients struct {
	// String representation of the receiver address
	Address string `json:"address"`
	// String representation of the amount
	Amount string `json:"amount"`
}

// NewListXRPRippleTransactionsByBlockHeightRIRecipients instantiates a new ListXRPRippleTransactionsByBlockHeightRIRecipients object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListXRPRippleTransactionsByBlockHeightRIRecipients(address string, amount string) *ListXRPRippleTransactionsByBlockHeightRIRecipients {
	this := ListXRPRippleTransactionsByBlockHeightRIRecipients{}
	this.Address = address
	this.Amount = amount
	return &this
}

// NewListXRPRippleTransactionsByBlockHeightRIRecipientsWithDefaults instantiates a new ListXRPRippleTransactionsByBlockHeightRIRecipients object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListXRPRippleTransactionsByBlockHeightRIRecipientsWithDefaults() *ListXRPRippleTransactionsByBlockHeightRIRecipients {
	this := ListXRPRippleTransactionsByBlockHeightRIRecipients{}
	return &this
}

// GetAddress returns the Address field value
func (o *ListXRPRippleTransactionsByBlockHeightRIRecipients) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByBlockHeightRIRecipients) GetAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *ListXRPRippleTransactionsByBlockHeightRIRecipients) SetAddress(v string) {
	o.Address = v
}

// GetAmount returns the Amount field value
func (o *ListXRPRippleTransactionsByBlockHeightRIRecipients) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByBlockHeightRIRecipients) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ListXRPRippleTransactionsByBlockHeightRIRecipients) SetAmount(v string) {
	o.Amount = v
}

func (o ListXRPRippleTransactionsByBlockHeightRIRecipients) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	return json.Marshal(toSerialize)
}

type NullableListXRPRippleTransactionsByBlockHeightRIRecipients struct {
	value *ListXRPRippleTransactionsByBlockHeightRIRecipients
	isSet bool
}

func (v NullableListXRPRippleTransactionsByBlockHeightRIRecipients) Get() *ListXRPRippleTransactionsByBlockHeightRIRecipients {
	return v.value
}

func (v *NullableListXRPRippleTransactionsByBlockHeightRIRecipients) Set(val *ListXRPRippleTransactionsByBlockHeightRIRecipients) {
	v.value = val
	v.isSet = true
}

func (v NullableListXRPRippleTransactionsByBlockHeightRIRecipients) IsSet() bool {
	return v.isSet
}

func (v *NullableListXRPRippleTransactionsByBlockHeightRIRecipients) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListXRPRippleTransactionsByBlockHeightRIRecipients(val *ListXRPRippleTransactionsByBlockHeightRIRecipients) *NullableListXRPRippleTransactionsByBlockHeightRIRecipients {
	return &NullableListXRPRippleTransactionsByBlockHeightRIRecipients{value: val, isSet: true}
}

func (v NullableListXRPRippleTransactionsByBlockHeightRIRecipients) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListXRPRippleTransactionsByBlockHeightRIRecipients) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


