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

// AddressCoinsTransactionConfirmed struct for AddressCoinsTransactionConfirmed
type AddressCoinsTransactionConfirmed struct {
	// Specifies the version of the API that incorporates this endpoint.
	ApiVersion string `json:"apiVersion"`
	// Represents a unique identifier that serves as reference to the specific request which prompts a callback, e.g. Blockchain Events Subscription, Blockchain Automation, etc.
	ReferenceId string `json:"referenceId"`
	// Specifies a unique ID generated by the system and attached to each callback. It is used by the server to recognize consecutive requests with the same data with the purpose not to perform the same operation twice.
	IdempotencyKey string `json:"idempotencyKey"`
	Data AddressCoinsTransactionConfirmedData `json:"data"`
}

// NewAddressCoinsTransactionConfirmed instantiates a new AddressCoinsTransactionConfirmed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressCoinsTransactionConfirmed(apiVersion string, referenceId string, idempotencyKey string, data AddressCoinsTransactionConfirmedData) *AddressCoinsTransactionConfirmed {
	this := AddressCoinsTransactionConfirmed{}
	this.ApiVersion = apiVersion
	this.ReferenceId = referenceId
	this.IdempotencyKey = idempotencyKey
	this.Data = data
	return &this
}

// NewAddressCoinsTransactionConfirmedWithDefaults instantiates a new AddressCoinsTransactionConfirmed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressCoinsTransactionConfirmedWithDefaults() *AddressCoinsTransactionConfirmed {
	this := AddressCoinsTransactionConfirmed{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *AddressCoinsTransactionConfirmed) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *AddressCoinsTransactionConfirmed) GetApiVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *AddressCoinsTransactionConfirmed) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetReferenceId returns the ReferenceId field value
func (o *AddressCoinsTransactionConfirmed) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *AddressCoinsTransactionConfirmed) GetReferenceIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *AddressCoinsTransactionConfirmed) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetIdempotencyKey returns the IdempotencyKey field value
func (o *AddressCoinsTransactionConfirmed) GetIdempotencyKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdempotencyKey
}

// GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field value
// and a boolean to check if the value has been set.
func (o *AddressCoinsTransactionConfirmed) GetIdempotencyKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IdempotencyKey, true
}

// SetIdempotencyKey sets field value
func (o *AddressCoinsTransactionConfirmed) SetIdempotencyKey(v string) {
	o.IdempotencyKey = v
}

// GetData returns the Data field value
func (o *AddressCoinsTransactionConfirmed) GetData() AddressCoinsTransactionConfirmedData {
	if o == nil {
		var ret AddressCoinsTransactionConfirmedData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AddressCoinsTransactionConfirmed) GetDataOk() (*AddressCoinsTransactionConfirmedData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AddressCoinsTransactionConfirmed) SetData(v AddressCoinsTransactionConfirmedData) {
	o.Data = v
}

func (o AddressCoinsTransactionConfirmed) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if true {
		toSerialize["referenceId"] = o.ReferenceId
	}
	if true {
		toSerialize["idempotencyKey"] = o.IdempotencyKey
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAddressCoinsTransactionConfirmed struct {
	value *AddressCoinsTransactionConfirmed
	isSet bool
}

func (v NullableAddressCoinsTransactionConfirmed) Get() *AddressCoinsTransactionConfirmed {
	return v.value
}

func (v *NullableAddressCoinsTransactionConfirmed) Set(val *AddressCoinsTransactionConfirmed) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressCoinsTransactionConfirmed) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressCoinsTransactionConfirmed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressCoinsTransactionConfirmed(val *AddressCoinsTransactionConfirmed) *NullableAddressCoinsTransactionConfirmed {
	return &NullableAddressCoinsTransactionConfirmed{value: val, isSet: true}
}

func (v NullableAddressCoinsTransactionConfirmed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressCoinsTransactionConfirmed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


