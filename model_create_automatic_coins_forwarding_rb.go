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

// CreateAutomaticCoinsForwardingRB struct for CreateAutomaticCoinsForwardingRB
type CreateAutomaticCoinsForwardingRB struct {
	// In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user.
	Context *string `json:"context,omitempty"`
	Data CreateAutomaticCoinsForwardingRBData `json:"data"`
}

// NewCreateAutomaticCoinsForwardingRB instantiates a new CreateAutomaticCoinsForwardingRB object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAutomaticCoinsForwardingRB(data CreateAutomaticCoinsForwardingRBData) *CreateAutomaticCoinsForwardingRB {
	this := CreateAutomaticCoinsForwardingRB{}
	this.Data = data
	return &this
}

// NewCreateAutomaticCoinsForwardingRBWithDefaults instantiates a new CreateAutomaticCoinsForwardingRB object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAutomaticCoinsForwardingRBWithDefaults() *CreateAutomaticCoinsForwardingRB {
	this := CreateAutomaticCoinsForwardingRB{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *CreateAutomaticCoinsForwardingRB) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAutomaticCoinsForwardingRB) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *CreateAutomaticCoinsForwardingRB) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *CreateAutomaticCoinsForwardingRB) SetContext(v string) {
	o.Context = &v
}

// GetData returns the Data field value
func (o *CreateAutomaticCoinsForwardingRB) GetData() CreateAutomaticCoinsForwardingRBData {
	if o == nil {
		var ret CreateAutomaticCoinsForwardingRBData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateAutomaticCoinsForwardingRB) GetDataOk() (*CreateAutomaticCoinsForwardingRBData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CreateAutomaticCoinsForwardingRB) SetData(v CreateAutomaticCoinsForwardingRBData) {
	o.Data = v
}

func (o CreateAutomaticCoinsForwardingRB) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCreateAutomaticCoinsForwardingRB struct {
	value *CreateAutomaticCoinsForwardingRB
	isSet bool
}

func (v NullableCreateAutomaticCoinsForwardingRB) Get() *CreateAutomaticCoinsForwardingRB {
	return v.value
}

func (v *NullableCreateAutomaticCoinsForwardingRB) Set(val *CreateAutomaticCoinsForwardingRB) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAutomaticCoinsForwardingRB) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAutomaticCoinsForwardingRB) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAutomaticCoinsForwardingRB(val *CreateAutomaticCoinsForwardingRB) *NullableCreateAutomaticCoinsForwardingRB {
	return &NullableCreateAutomaticCoinsForwardingRB{value: val, isSet: true}
}

func (v NullableCreateAutomaticCoinsForwardingRB) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAutomaticCoinsForwardingRB) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


