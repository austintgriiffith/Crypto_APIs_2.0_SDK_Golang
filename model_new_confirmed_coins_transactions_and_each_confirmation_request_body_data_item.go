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

// NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem struct for NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem
type NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem struct {
	// Represents the address of the transaction, per which the result is returned.
	Address string `json:"address"`
	// Specifies a flag that permits or denies the creation of duplicate addresses.
	AllowDuplicates *bool `json:"allowDuplicates,omitempty"`
	// Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs.
	CallbackSecretKey *string `json:"callbackSecretKey,omitempty"`
	// Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs.
	CallbackUrl string `json:"callbackUrl"`
	// Represents the number of confirmations, i.e. the amount of blocks that have been built on top of this block.
	ConfirmationsCount int32 `json:"confirmationsCount"`
}

// NewNewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem instantiates a new NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem(address string, callbackUrl string, confirmationsCount int32) *NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem {
	this := NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem{}
	this.Address = address
	var allowDuplicates bool = false
	this.AllowDuplicates = &allowDuplicates
	this.CallbackUrl = callbackUrl
	this.ConfirmationsCount = confirmationsCount
	return &this
}

// NewNewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItemWithDefaults instantiates a new NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItemWithDefaults() *NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem {
	this := NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem{}
	var allowDuplicates bool = false
	this.AllowDuplicates = &allowDuplicates
	return &this
}

// GetAddress returns the Address field value
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem) GetAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem) SetAddress(v string) {
	o.Address = v
}

// GetAllowDuplicates returns the AllowDuplicates field value if set, zero value otherwise.
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem) GetAllowDuplicates() bool {
	if o == nil || o.AllowDuplicates == nil {
		var ret bool
		return ret
	}
	return *o.AllowDuplicates
}

// GetAllowDuplicatesOk returns a tuple with the AllowDuplicates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem) GetAllowDuplicatesOk() (*bool, bool) {
	if o == nil || o.AllowDuplicates == nil {
		return nil, false
	}
	return o.AllowDuplicates, true
}

// HasAllowDuplicates returns a boolean if a field has been set.
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem) HasAllowDuplicates() bool {
	if o != nil && o.AllowDuplicates != nil {
		return true
	}

	return false
}

// SetAllowDuplicates gets a reference to the given bool and assigns it to the AllowDuplicates field.
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem) SetAllowDuplicates(v bool) {
	o.AllowDuplicates = &v
}

// GetCallbackSecretKey returns the CallbackSecretKey field value if set, zero value otherwise.
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem) GetCallbackSecretKey() string {
	if o == nil || o.CallbackSecretKey == nil {
		var ret string
		return ret
	}
	return *o.CallbackSecretKey
}

// GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem) GetCallbackSecretKeyOk() (*string, bool) {
	if o == nil || o.CallbackSecretKey == nil {
		return nil, false
	}
	return o.CallbackSecretKey, true
}

// HasCallbackSecretKey returns a boolean if a field has been set.
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem) HasCallbackSecretKey() bool {
	if o != nil && o.CallbackSecretKey != nil {
		return true
	}

	return false
}

// SetCallbackSecretKey gets a reference to the given string and assigns it to the CallbackSecretKey field.
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem) SetCallbackSecretKey(v string) {
	o.CallbackSecretKey = &v
}

// GetCallbackUrl returns the CallbackUrl field value
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem) GetCallbackUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value
// and a boolean to check if the value has been set.
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem) GetCallbackUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CallbackUrl, true
}

// SetCallbackUrl sets field value
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem) SetCallbackUrl(v string) {
	o.CallbackUrl = v
}

// GetConfirmationsCount returns the ConfirmationsCount field value
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem) GetConfirmationsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConfirmationsCount
}

// GetConfirmationsCountOk returns a tuple with the ConfirmationsCount field value
// and a boolean to check if the value has been set.
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem) GetConfirmationsCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConfirmationsCount, true
}

// SetConfirmationsCount sets field value
func (o *NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem) SetConfirmationsCount(v int32) {
	o.ConfirmationsCount = v
}

func (o NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
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
		toSerialize["confirmationsCount"] = o.ConfirmationsCount
	}
	return json.Marshal(toSerialize)
}

type NullableNewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem struct {
	value *NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem
	isSet bool
}

func (v NullableNewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem) Get() *NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem {
	return v.value
}

func (v *NullableNewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem) Set(val *NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullableNewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullableNewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem(val *NewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem) *NullableNewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem {
	return &NullableNewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem{value: val, isSet: true}
}

func (v NullableNewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewConfirmedCoinsTransactionsAndEachConfirmationRequestBodyDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

