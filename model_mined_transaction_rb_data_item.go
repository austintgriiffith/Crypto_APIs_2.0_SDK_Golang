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

// MinedTransactionRBDataItem struct for MinedTransactionRBDataItem
type MinedTransactionRBDataItem struct {
	// Specifies a flag that permits or denies the creation of duplicate addresses.
	AllowDuplicates *bool `json:"allowDuplicates,omitempty"`
	// Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. For more information please see our [Documentation](https://developers.cryptoapis.io/technical-documentation/general-information/callbacks#callback-security).
	CallbackSecretKey *string `json:"callbackSecretKey,omitempty"`
	// Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs.
	CallbackUrl string `json:"callbackUrl"`
	// Represents the unique identification string that defines the transaction.
	TransactionId string `json:"transactionId"`
}

// NewMinedTransactionRBDataItem instantiates a new MinedTransactionRBDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMinedTransactionRBDataItem(callbackUrl string, transactionId string) *MinedTransactionRBDataItem {
	this := MinedTransactionRBDataItem{}
	this.CallbackUrl = callbackUrl
	this.TransactionId = transactionId
	return &this
}

// NewMinedTransactionRBDataItemWithDefaults instantiates a new MinedTransactionRBDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMinedTransactionRBDataItemWithDefaults() *MinedTransactionRBDataItem {
	this := MinedTransactionRBDataItem{}
	return &this
}

// GetAllowDuplicates returns the AllowDuplicates field value if set, zero value otherwise.
func (o *MinedTransactionRBDataItem) GetAllowDuplicates() bool {
	if o == nil || o.AllowDuplicates == nil {
		var ret bool
		return ret
	}
	return *o.AllowDuplicates
}

// GetAllowDuplicatesOk returns a tuple with the AllowDuplicates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinedTransactionRBDataItem) GetAllowDuplicatesOk() (*bool, bool) {
	if o == nil || o.AllowDuplicates == nil {
		return nil, false
	}
	return o.AllowDuplicates, true
}

// HasAllowDuplicates returns a boolean if a field has been set.
func (o *MinedTransactionRBDataItem) HasAllowDuplicates() bool {
	if o != nil && o.AllowDuplicates != nil {
		return true
	}

	return false
}

// SetAllowDuplicates gets a reference to the given bool and assigns it to the AllowDuplicates field.
func (o *MinedTransactionRBDataItem) SetAllowDuplicates(v bool) {
	o.AllowDuplicates = &v
}

// GetCallbackSecretKey returns the CallbackSecretKey field value if set, zero value otherwise.
func (o *MinedTransactionRBDataItem) GetCallbackSecretKey() string {
	if o == nil || o.CallbackSecretKey == nil {
		var ret string
		return ret
	}
	return *o.CallbackSecretKey
}

// GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinedTransactionRBDataItem) GetCallbackSecretKeyOk() (*string, bool) {
	if o == nil || o.CallbackSecretKey == nil {
		return nil, false
	}
	return o.CallbackSecretKey, true
}

// HasCallbackSecretKey returns a boolean if a field has been set.
func (o *MinedTransactionRBDataItem) HasCallbackSecretKey() bool {
	if o != nil && o.CallbackSecretKey != nil {
		return true
	}

	return false
}

// SetCallbackSecretKey gets a reference to the given string and assigns it to the CallbackSecretKey field.
func (o *MinedTransactionRBDataItem) SetCallbackSecretKey(v string) {
	o.CallbackSecretKey = &v
}

// GetCallbackUrl returns the CallbackUrl field value
func (o *MinedTransactionRBDataItem) GetCallbackUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value
// and a boolean to check if the value has been set.
func (o *MinedTransactionRBDataItem) GetCallbackUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CallbackUrl, true
}

// SetCallbackUrl sets field value
func (o *MinedTransactionRBDataItem) SetCallbackUrl(v string) {
	o.CallbackUrl = v
}

// GetTransactionId returns the TransactionId field value
func (o *MinedTransactionRBDataItem) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *MinedTransactionRBDataItem) GetTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *MinedTransactionRBDataItem) SetTransactionId(v string) {
	o.TransactionId = v
}

func (o MinedTransactionRBDataItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowDuplicates != nil {
		toSerialize["allowDuplicates"] = o.AllowDuplicates
	}
	if o.CallbackSecretKey != nil {
		toSerialize["callbackSecretKey"] = o.CallbackSecretKey
	}
	if true {
		toSerialize["callbackUrl"] = o.CallbackUrl
	}
	if true {
		toSerialize["transactionId"] = o.TransactionId
	}
	return json.Marshal(toSerialize)
}

type NullableMinedTransactionRBDataItem struct {
	value *MinedTransactionRBDataItem
	isSet bool
}

func (v NullableMinedTransactionRBDataItem) Get() *MinedTransactionRBDataItem {
	return v.value
}

func (v *NullableMinedTransactionRBDataItem) Set(val *MinedTransactionRBDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullableMinedTransactionRBDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullableMinedTransactionRBDataItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinedTransactionRBDataItem(val *MinedTransactionRBDataItem) *NullableMinedTransactionRBDataItem {
	return &NullableMinedTransactionRBDataItem{value: val, isSet: true}
}

func (v NullableMinedTransactionRBDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinedTransactionRBDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


