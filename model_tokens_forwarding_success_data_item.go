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

// TokensForwardingSuccessDataItem Defines an `item` as one result.
type TokensForwardingSuccessDataItem struct {
	// Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc.
	Blockchain string `json:"blockchain"`
	// Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
	Network string `json:"network"`
	// Represents the hash of the address that provides the tokens.
	FromAddress string `json:"fromAddress"`
	// Represents the hash of the address to forward the tokens to.
	ToAddress string `json:"toAddress"`
	// Represents the amount of the fee spent for the tokens to be forwarded.
	SpentFeesAmount string `json:"spentFeesAmount"`
	// Represents the unit of the fee spent for the tokens to be forwarded, e.g. BTC.
	SpentFeesUnit string `json:"spentFeesUnit"`
	// Defines the unique Transaction ID that triggered the token forwarding.
	TriggerTransactionId string `json:"triggerTransactionId"`
	// Defines the unique Transaction ID that forwarded the tokens.
	ForwardingTransactionId string `json:"forwardingTransactionId"`
	// Defines the type of token sent with the transaction, e.g. ERC 20.
	TokenType string `json:"tokenType"`
	Token TokensForwardingSuccessToken `json:"token"`
}

// NewTokensForwardingSuccessDataItem instantiates a new TokensForwardingSuccessDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokensForwardingSuccessDataItem(blockchain string, network string, fromAddress string, toAddress string, spentFeesAmount string, spentFeesUnit string, triggerTransactionId string, forwardingTransactionId string, tokenType string, token TokensForwardingSuccessToken) *TokensForwardingSuccessDataItem {
	this := TokensForwardingSuccessDataItem{}
	this.Blockchain = blockchain
	this.Network = network
	this.FromAddress = fromAddress
	this.ToAddress = toAddress
	this.SpentFeesAmount = spentFeesAmount
	this.SpentFeesUnit = spentFeesUnit
	this.TriggerTransactionId = triggerTransactionId
	this.ForwardingTransactionId = forwardingTransactionId
	this.TokenType = tokenType
	this.Token = token
	return &this
}

// NewTokensForwardingSuccessDataItemWithDefaults instantiates a new TokensForwardingSuccessDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokensForwardingSuccessDataItemWithDefaults() *TokensForwardingSuccessDataItem {
	this := TokensForwardingSuccessDataItem{}
	return &this
}

// GetBlockchain returns the Blockchain field value
func (o *TokensForwardingSuccessDataItem) GetBlockchain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Blockchain
}

// GetBlockchainOk returns a tuple with the Blockchain field value
// and a boolean to check if the value has been set.
func (o *TokensForwardingSuccessDataItem) GetBlockchainOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Blockchain, true
}

// SetBlockchain sets field value
func (o *TokensForwardingSuccessDataItem) SetBlockchain(v string) {
	o.Blockchain = v
}

// GetNetwork returns the Network field value
func (o *TokensForwardingSuccessDataItem) GetNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *TokensForwardingSuccessDataItem) GetNetworkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *TokensForwardingSuccessDataItem) SetNetwork(v string) {
	o.Network = v
}

// GetFromAddress returns the FromAddress field value
func (o *TokensForwardingSuccessDataItem) GetFromAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromAddress
}

// GetFromAddressOk returns a tuple with the FromAddress field value
// and a boolean to check if the value has been set.
func (o *TokensForwardingSuccessDataItem) GetFromAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FromAddress, true
}

// SetFromAddress sets field value
func (o *TokensForwardingSuccessDataItem) SetFromAddress(v string) {
	o.FromAddress = v
}

// GetToAddress returns the ToAddress field value
func (o *TokensForwardingSuccessDataItem) GetToAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToAddress
}

// GetToAddressOk returns a tuple with the ToAddress field value
// and a boolean to check if the value has been set.
func (o *TokensForwardingSuccessDataItem) GetToAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ToAddress, true
}

// SetToAddress sets field value
func (o *TokensForwardingSuccessDataItem) SetToAddress(v string) {
	o.ToAddress = v
}

// GetSpentFeesAmount returns the SpentFeesAmount field value
func (o *TokensForwardingSuccessDataItem) GetSpentFeesAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpentFeesAmount
}

// GetSpentFeesAmountOk returns a tuple with the SpentFeesAmount field value
// and a boolean to check if the value has been set.
func (o *TokensForwardingSuccessDataItem) GetSpentFeesAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SpentFeesAmount, true
}

// SetSpentFeesAmount sets field value
func (o *TokensForwardingSuccessDataItem) SetSpentFeesAmount(v string) {
	o.SpentFeesAmount = v
}

// GetSpentFeesUnit returns the SpentFeesUnit field value
func (o *TokensForwardingSuccessDataItem) GetSpentFeesUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpentFeesUnit
}

// GetSpentFeesUnitOk returns a tuple with the SpentFeesUnit field value
// and a boolean to check if the value has been set.
func (o *TokensForwardingSuccessDataItem) GetSpentFeesUnitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SpentFeesUnit, true
}

// SetSpentFeesUnit sets field value
func (o *TokensForwardingSuccessDataItem) SetSpentFeesUnit(v string) {
	o.SpentFeesUnit = v
}

// GetTriggerTransactionId returns the TriggerTransactionId field value
func (o *TokensForwardingSuccessDataItem) GetTriggerTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TriggerTransactionId
}

// GetTriggerTransactionIdOk returns a tuple with the TriggerTransactionId field value
// and a boolean to check if the value has been set.
func (o *TokensForwardingSuccessDataItem) GetTriggerTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TriggerTransactionId, true
}

// SetTriggerTransactionId sets field value
func (o *TokensForwardingSuccessDataItem) SetTriggerTransactionId(v string) {
	o.TriggerTransactionId = v
}

// GetForwardingTransactionId returns the ForwardingTransactionId field value
func (o *TokensForwardingSuccessDataItem) GetForwardingTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ForwardingTransactionId
}

// GetForwardingTransactionIdOk returns a tuple with the ForwardingTransactionId field value
// and a boolean to check if the value has been set.
func (o *TokensForwardingSuccessDataItem) GetForwardingTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ForwardingTransactionId, true
}

// SetForwardingTransactionId sets field value
func (o *TokensForwardingSuccessDataItem) SetForwardingTransactionId(v string) {
	o.ForwardingTransactionId = v
}

// GetTokenType returns the TokenType field value
func (o *TokensForwardingSuccessDataItem) GetTokenType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value
// and a boolean to check if the value has been set.
func (o *TokensForwardingSuccessDataItem) GetTokenTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TokenType, true
}

// SetTokenType sets field value
func (o *TokensForwardingSuccessDataItem) SetTokenType(v string) {
	o.TokenType = v
}

// GetToken returns the Token field value
func (o *TokensForwardingSuccessDataItem) GetToken() TokensForwardingSuccessToken {
	if o == nil {
		var ret TokensForwardingSuccessToken
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *TokensForwardingSuccessDataItem) GetTokenOk() (*TokensForwardingSuccessToken, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *TokensForwardingSuccessDataItem) SetToken(v TokensForwardingSuccessToken) {
	o.Token = v
}

func (o TokensForwardingSuccessDataItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["blockchain"] = o.Blockchain
	}
	if true {
		toSerialize["network"] = o.Network
	}
	if true {
		toSerialize["fromAddress"] = o.FromAddress
	}
	if true {
		toSerialize["toAddress"] = o.ToAddress
	}
	if true {
		toSerialize["spentFeesAmount"] = o.SpentFeesAmount
	}
	if true {
		toSerialize["spentFeesUnit"] = o.SpentFeesUnit
	}
	if true {
		toSerialize["triggerTransactionId"] = o.TriggerTransactionId
	}
	if true {
		toSerialize["forwardingTransactionId"] = o.ForwardingTransactionId
	}
	if true {
		toSerialize["tokenType"] = o.TokenType
	}
	if true {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableTokensForwardingSuccessDataItem struct {
	value *TokensForwardingSuccessDataItem
	isSet bool
}

func (v NullableTokensForwardingSuccessDataItem) Get() *TokensForwardingSuccessDataItem {
	return v.value
}

func (v *NullableTokensForwardingSuccessDataItem) Set(val *TokensForwardingSuccessDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullableTokensForwardingSuccessDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullableTokensForwardingSuccessDataItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokensForwardingSuccessDataItem(val *TokensForwardingSuccessDataItem) *NullableTokensForwardingSuccessDataItem {
	return &NullableTokensForwardingSuccessDataItem{value: val, isSet: true}
}

func (v NullableTokensForwardingSuccessDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokensForwardingSuccessDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


