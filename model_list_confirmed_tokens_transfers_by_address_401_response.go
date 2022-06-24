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

// ListConfirmedTokensTransfersByAddress401Response struct for ListConfirmedTokensTransfersByAddress401Response
type ListConfirmedTokensTransfersByAddress401Response struct {
	// Specifies the version of the API that incorporates this endpoint.
	ApiVersion string `json:"apiVersion"`
	// Defines the ID of the request. The `requestId` is generated by Crypto APIs and it's unique for every request.
	RequestId string `json:"requestId"`
	// In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user.
	Context *string `json:"context,omitempty"`
	Error ListConfirmedTokensTransfersByAddressE401 `json:"error"`
}

// NewListConfirmedTokensTransfersByAddress401Response instantiates a new ListConfirmedTokensTransfersByAddress401Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConfirmedTokensTransfersByAddress401Response(apiVersion string, requestId string, error_ ListConfirmedTokensTransfersByAddressE401) *ListConfirmedTokensTransfersByAddress401Response {
	this := ListConfirmedTokensTransfersByAddress401Response{}
	this.ApiVersion = apiVersion
	this.RequestId = requestId
	this.Error = error_
	return &this
}

// NewListConfirmedTokensTransfersByAddress401ResponseWithDefaults instantiates a new ListConfirmedTokensTransfersByAddress401Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConfirmedTokensTransfersByAddress401ResponseWithDefaults() *ListConfirmedTokensTransfersByAddress401Response {
	this := ListConfirmedTokensTransfersByAddress401Response{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *ListConfirmedTokensTransfersByAddress401Response) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTokensTransfersByAddress401Response) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *ListConfirmedTokensTransfersByAddress401Response) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetRequestId returns the RequestId field value
func (o *ListConfirmedTokensTransfersByAddress401Response) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTokensTransfersByAddress401Response) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *ListConfirmedTokensTransfersByAddress401Response) SetRequestId(v string) {
	o.RequestId = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *ListConfirmedTokensTransfersByAddress401Response) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConfirmedTokensTransfersByAddress401Response) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *ListConfirmedTokensTransfersByAddress401Response) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *ListConfirmedTokensTransfersByAddress401Response) SetContext(v string) {
	o.Context = &v
}

// GetError returns the Error field value
func (o *ListConfirmedTokensTransfersByAddress401Response) GetError() ListConfirmedTokensTransfersByAddressE401 {
	if o == nil {
		var ret ListConfirmedTokensTransfersByAddressE401
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTokensTransfersByAddress401Response) GetErrorOk() (*ListConfirmedTokensTransfersByAddressE401, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *ListConfirmedTokensTransfersByAddress401Response) SetError(v ListConfirmedTokensTransfersByAddressE401) {
	o.Error = v
}

func (o ListConfirmedTokensTransfersByAddress401Response) MarshalJSON() ([]byte, error) {
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
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableListConfirmedTokensTransfersByAddress401Response struct {
	value *ListConfirmedTokensTransfersByAddress401Response
	isSet bool
}

func (v NullableListConfirmedTokensTransfersByAddress401Response) Get() *ListConfirmedTokensTransfersByAddress401Response {
	return v.value
}

func (v *NullableListConfirmedTokensTransfersByAddress401Response) Set(val *ListConfirmedTokensTransfersByAddress401Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListConfirmedTokensTransfersByAddress401Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListConfirmedTokensTransfersByAddress401Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConfirmedTokensTransfersByAddress401Response(val *ListConfirmedTokensTransfersByAddress401Response) *NullableListConfirmedTokensTransfersByAddress401Response {
	return &NullableListConfirmedTokensTransfersByAddress401Response{value: val, isSet: true}
}

func (v NullableListConfirmedTokensTransfersByAddress401Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConfirmedTokensTransfersByAddress401Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

