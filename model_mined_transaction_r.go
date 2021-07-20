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

// MinedTransactionR struct for MinedTransactionR
type MinedTransactionR struct {
	// Specifies the version of the API that incorporates this endpoint.
	ApiVersion string `json:"apiVersion"`
	// Defines the ID of the request. The `requestId` is generated by Crypto APIs and it's unique for every request.
	RequestId string `json:"requestId"`
	// In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user.
	Context *string `json:"context,omitempty"`
	Data MinedTransactionRData `json:"data"`
}

// NewMinedTransactionR instantiates a new MinedTransactionR object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMinedTransactionR(apiVersion string, requestId string, data MinedTransactionRData) *MinedTransactionR {
	this := MinedTransactionR{}
	this.ApiVersion = apiVersion
	this.RequestId = requestId
	this.Data = data
	return &this
}

// NewMinedTransactionRWithDefaults instantiates a new MinedTransactionR object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMinedTransactionRWithDefaults() *MinedTransactionR {
	this := MinedTransactionR{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *MinedTransactionR) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *MinedTransactionR) GetApiVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *MinedTransactionR) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetRequestId returns the RequestId field value
func (o *MinedTransactionR) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *MinedTransactionR) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *MinedTransactionR) SetRequestId(v string) {
	o.RequestId = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *MinedTransactionR) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinedTransactionR) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *MinedTransactionR) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *MinedTransactionR) SetContext(v string) {
	o.Context = &v
}

// GetData returns the Data field value
func (o *MinedTransactionR) GetData() MinedTransactionRData {
	if o == nil {
		var ret MinedTransactionRData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *MinedTransactionR) GetDataOk() (*MinedTransactionRData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *MinedTransactionR) SetData(v MinedTransactionRData) {
	o.Data = v
}

func (o MinedTransactionR) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if true {
		toSerialize["requestId"] = o.RequestId
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableMinedTransactionR struct {
	value *MinedTransactionR
	isSet bool
}

func (v NullableMinedTransactionR) Get() *MinedTransactionR {
	return v.value
}

func (v *NullableMinedTransactionR) Set(val *MinedTransactionR) {
	v.value = val
	v.isSet = true
}

func (v NullableMinedTransactionR) IsSet() bool {
	return v.isSet
}

func (v *NullableMinedTransactionR) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinedTransactionR(val *MinedTransactionR) *NullableMinedTransactionR {
	return &NullableMinedTransactionR{value: val, isSet: true}
}

func (v NullableMinedTransactionR) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinedTransactionR) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

