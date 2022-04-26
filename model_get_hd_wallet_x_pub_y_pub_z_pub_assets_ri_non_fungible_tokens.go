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

// GetHDWalletXPubYPubZPubAssetsRINonFungibleTokens struct for GetHDWalletXPubYPubZPubAssetsRINonFungibleTokens
type GetHDWalletXPubYPubZPubAssetsRINonFungibleTokens struct {
	// Represents tokens' contract address.
	Identifier string `json:"identifier"`
	// Defines the symbol of the non-fungible token.
	Symbol string `json:"symbol"`
	// Represents tokens' unique identifier.
	TokenId string `json:"tokenId"`
	// Defines the specific token type.
	Type string `json:"type"`
}

// NewGetHDWalletXPubYPubZPubAssetsRINonFungibleTokens instantiates a new GetHDWalletXPubYPubZPubAssetsRINonFungibleTokens object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetHDWalletXPubYPubZPubAssetsRINonFungibleTokens(identifier string, symbol string, tokenId string, type_ string) *GetHDWalletXPubYPubZPubAssetsRINonFungibleTokens {
	this := GetHDWalletXPubYPubZPubAssetsRINonFungibleTokens{}
	this.Identifier = identifier
	this.Symbol = symbol
	this.TokenId = tokenId
	this.Type = type_
	return &this
}

// NewGetHDWalletXPubYPubZPubAssetsRINonFungibleTokensWithDefaults instantiates a new GetHDWalletXPubYPubZPubAssetsRINonFungibleTokens object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetHDWalletXPubYPubZPubAssetsRINonFungibleTokensWithDefaults() *GetHDWalletXPubYPubZPubAssetsRINonFungibleTokens {
	this := GetHDWalletXPubYPubZPubAssetsRINonFungibleTokens{}
	return &this
}

// GetIdentifier returns the Identifier field value
func (o *GetHDWalletXPubYPubZPubAssetsRINonFungibleTokens) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *GetHDWalletXPubYPubZPubAssetsRINonFungibleTokens) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *GetHDWalletXPubYPubZPubAssetsRINonFungibleTokens) SetIdentifier(v string) {
	o.Identifier = v
}

// GetSymbol returns the Symbol field value
func (o *GetHDWalletXPubYPubZPubAssetsRINonFungibleTokens) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *GetHDWalletXPubYPubZPubAssetsRINonFungibleTokens) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *GetHDWalletXPubYPubZPubAssetsRINonFungibleTokens) SetSymbol(v string) {
	o.Symbol = v
}

// GetTokenId returns the TokenId field value
func (o *GetHDWalletXPubYPubZPubAssetsRINonFungibleTokens) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *GetHDWalletXPubYPubZPubAssetsRINonFungibleTokens) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *GetHDWalletXPubYPubZPubAssetsRINonFungibleTokens) SetTokenId(v string) {
	o.TokenId = v
}

// GetType returns the Type field value
func (o *GetHDWalletXPubYPubZPubAssetsRINonFungibleTokens) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetHDWalletXPubYPubZPubAssetsRINonFungibleTokens) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetHDWalletXPubYPubZPubAssetsRINonFungibleTokens) SetType(v string) {
	o.Type = v
}

func (o GetHDWalletXPubYPubZPubAssetsRINonFungibleTokens) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["identifier"] = o.Identifier
	}
	if true {
		toSerialize["symbol"] = o.Symbol
	}
	if true {
		toSerialize["tokenId"] = o.TokenId
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableGetHDWalletXPubYPubZPubAssetsRINonFungibleTokens struct {
	value *GetHDWalletXPubYPubZPubAssetsRINonFungibleTokens
	isSet bool
}

func (v NullableGetHDWalletXPubYPubZPubAssetsRINonFungibleTokens) Get() *GetHDWalletXPubYPubZPubAssetsRINonFungibleTokens {
	return v.value
}

func (v *NullableGetHDWalletXPubYPubZPubAssetsRINonFungibleTokens) Set(val *GetHDWalletXPubYPubZPubAssetsRINonFungibleTokens) {
	v.value = val
	v.isSet = true
}

func (v NullableGetHDWalletXPubYPubZPubAssetsRINonFungibleTokens) IsSet() bool {
	return v.isSet
}

func (v *NullableGetHDWalletXPubYPubZPubAssetsRINonFungibleTokens) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetHDWalletXPubYPubZPubAssetsRINonFungibleTokens(val *GetHDWalletXPubYPubZPubAssetsRINonFungibleTokens) *NullableGetHDWalletXPubYPubZPubAssetsRINonFungibleTokens {
	return &NullableGetHDWalletXPubYPubZPubAssetsRINonFungibleTokens{value: val, isSet: true}
}

func (v NullableGetHDWalletXPubYPubZPubAssetsRINonFungibleTokens) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetHDWalletXPubYPubZPubAssetsRINonFungibleTokens) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

