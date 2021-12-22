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

// CreateFungibleTokensTransactionRequestFromAddressRB struct for CreateFungibleTokensTransactionRequestFromAddressRB
type CreateFungibleTokensTransactionRequestFromAddressRB struct {
	// In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user.
	Context *string `json:"context,omitempty"`
	Data CreateFungibleTokensTransactionRequestFromAddressRBData `json:"data"`
}

// NewCreateFungibleTokensTransactionRequestFromAddressRB instantiates a new CreateFungibleTokensTransactionRequestFromAddressRB object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFungibleTokensTransactionRequestFromAddressRB(data CreateFungibleTokensTransactionRequestFromAddressRBData) *CreateFungibleTokensTransactionRequestFromAddressRB {
	this := CreateFungibleTokensTransactionRequestFromAddressRB{}
	this.Data = data
	return &this
}

// NewCreateFungibleTokensTransactionRequestFromAddressRBWithDefaults instantiates a new CreateFungibleTokensTransactionRequestFromAddressRB object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFungibleTokensTransactionRequestFromAddressRBWithDefaults() *CreateFungibleTokensTransactionRequestFromAddressRB {
	this := CreateFungibleTokensTransactionRequestFromAddressRB{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *CreateFungibleTokensTransactionRequestFromAddressRB) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFungibleTokensTransactionRequestFromAddressRB) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *CreateFungibleTokensTransactionRequestFromAddressRB) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *CreateFungibleTokensTransactionRequestFromAddressRB) SetContext(v string) {
	o.Context = &v
}

// GetData returns the Data field value
func (o *CreateFungibleTokensTransactionRequestFromAddressRB) GetData() CreateFungibleTokensTransactionRequestFromAddressRBData {
	if o == nil {
		var ret CreateFungibleTokensTransactionRequestFromAddressRBData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateFungibleTokensTransactionRequestFromAddressRB) GetDataOk() (*CreateFungibleTokensTransactionRequestFromAddressRBData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CreateFungibleTokensTransactionRequestFromAddressRB) SetData(v CreateFungibleTokensTransactionRequestFromAddressRBData) {
	o.Data = v
}

func (o CreateFungibleTokensTransactionRequestFromAddressRB) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCreateFungibleTokensTransactionRequestFromAddressRB struct {
	value *CreateFungibleTokensTransactionRequestFromAddressRB
	isSet bool
}

func (v NullableCreateFungibleTokensTransactionRequestFromAddressRB) Get() *CreateFungibleTokensTransactionRequestFromAddressRB {
	return v.value
}

func (v *NullableCreateFungibleTokensTransactionRequestFromAddressRB) Set(val *CreateFungibleTokensTransactionRequestFromAddressRB) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFungibleTokensTransactionRequestFromAddressRB) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFungibleTokensTransactionRequestFromAddressRB) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFungibleTokensTransactionRequestFromAddressRB(val *CreateFungibleTokensTransactionRequestFromAddressRB) *NullableCreateFungibleTokensTransactionRequestFromAddressRB {
	return &NullableCreateFungibleTokensTransactionRequestFromAddressRB{value: val, isSet: true}
}

func (v NullableCreateFungibleTokensTransactionRequestFromAddressRB) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFungibleTokensTransactionRequestFromAddressRB) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


