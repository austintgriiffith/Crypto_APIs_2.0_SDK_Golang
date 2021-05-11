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

// MinedTransactionRequestBodyDataItem struct for MinedTransactionRequestBodyDataItem
type MinedTransactionRequestBodyDataItem struct {
	// Specifies a flag that permits or denies the creation of duplicate addresses.
	AllowDuplicates *bool `json:"allowDuplicates,omitempty"`
	// Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs.
	CallbackSecretKey *string `json:"callbackSecretKey,omitempty"`
	// Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs.
	CallbackUrl string `json:"callbackUrl"`
	// Represents the unique identification string that defines the transaction.
	TransactionId string `json:"transactionId"`
}

// NewMinedTransactionRequestBodyDataItem instantiates a new MinedTransactionRequestBodyDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMinedTransactionRequestBodyDataItem(callbackUrl string, transactionId string) *MinedTransactionRequestBodyDataItem {
	this := MinedTransactionRequestBodyDataItem{}
	this.CallbackUrl = callbackUrl
	this.TransactionId = transactionId
	return &this
}

// NewMinedTransactionRequestBodyDataItemWithDefaults instantiates a new MinedTransactionRequestBodyDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMinedTransactionRequestBodyDataItemWithDefaults() *MinedTransactionRequestBodyDataItem {
	this := MinedTransactionRequestBodyDataItem{}
	return &this
}

// GetAllowDuplicates returns the AllowDuplicates field value if set, zero value otherwise.
func (o *MinedTransactionRequestBodyDataItem) GetAllowDuplicates() bool {
	if o == nil || o.AllowDuplicates == nil {
		var ret bool
		return ret
	}
	return *o.AllowDuplicates
}

// GetAllowDuplicatesOk returns a tuple with the AllowDuplicates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinedTransactionRequestBodyDataItem) GetAllowDuplicatesOk() (*bool, bool) {
	if o == nil || o.AllowDuplicates == nil {
		return nil, false
	}
	return o.AllowDuplicates, true
}

// HasAllowDuplicates returns a boolean if a field has been set.
func (o *MinedTransactionRequestBodyDataItem) HasAllowDuplicates() bool {
	if o != nil && o.AllowDuplicates != nil {
		return true
	}

	return false
}

// SetAllowDuplicates gets a reference to the given bool and assigns it to the AllowDuplicates field.
func (o *MinedTransactionRequestBodyDataItem) SetAllowDuplicates(v bool) {
	o.AllowDuplicates = &v
}

// GetCallbackSecretKey returns the CallbackSecretKey field value if set, zero value otherwise.
func (o *MinedTransactionRequestBodyDataItem) GetCallbackSecretKey() string {
	if o == nil || o.CallbackSecretKey == nil {
		var ret string
		return ret
	}
	return *o.CallbackSecretKey
}

// GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinedTransactionRequestBodyDataItem) GetCallbackSecretKeyOk() (*string, bool) {
	if o == nil || o.CallbackSecretKey == nil {
		return nil, false
	}
	return o.CallbackSecretKey, true
}

// HasCallbackSecretKey returns a boolean if a field has been set.
func (o *MinedTransactionRequestBodyDataItem) HasCallbackSecretKey() bool {
	if o != nil && o.CallbackSecretKey != nil {
		return true
	}

	return false
}

// SetCallbackSecretKey gets a reference to the given string and assigns it to the CallbackSecretKey field.
func (o *MinedTransactionRequestBodyDataItem) SetCallbackSecretKey(v string) {
	o.CallbackSecretKey = &v
}

// GetCallbackUrl returns the CallbackUrl field value
func (o *MinedTransactionRequestBodyDataItem) GetCallbackUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value
// and a boolean to check if the value has been set.
func (o *MinedTransactionRequestBodyDataItem) GetCallbackUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CallbackUrl, true
}

// SetCallbackUrl sets field value
func (o *MinedTransactionRequestBodyDataItem) SetCallbackUrl(v string) {
	o.CallbackUrl = v
}

// GetTransactionId returns the TransactionId field value
func (o *MinedTransactionRequestBodyDataItem) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *MinedTransactionRequestBodyDataItem) GetTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *MinedTransactionRequestBodyDataItem) SetTransactionId(v string) {
	o.TransactionId = v
}

func (o MinedTransactionRequestBodyDataItem) MarshalJSON() ([]byte, error) {
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

type NullableMinedTransactionRequestBodyDataItem struct {
	value *MinedTransactionRequestBodyDataItem
	isSet bool
}

func (v NullableMinedTransactionRequestBodyDataItem) Get() *MinedTransactionRequestBodyDataItem {
	return v.value
}

func (v *NullableMinedTransactionRequestBodyDataItem) Set(val *MinedTransactionRequestBodyDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullableMinedTransactionRequestBodyDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullableMinedTransactionRequestBodyDataItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinedTransactionRequestBodyDataItem(val *MinedTransactionRequestBodyDataItem) *NullableMinedTransactionRequestBodyDataItem {
	return &NullableMinedTransactionRequestBodyDataItem{value: val, isSet: true}
}

func (v NullableMinedTransactionRequestBodyDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinedTransactionRequestBodyDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

