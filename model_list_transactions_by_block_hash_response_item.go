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

// ListTransactionsByBlockHashResponseItem struct for ListTransactionsByBlockHashResponseItem
type ListTransactionsByBlockHashResponseItem struct {
	// Represents the index position of the transaction in the specific block.
	Index int32 `json:"index"`
	// Represents the hash of the block where this transaction was mined/confirmed for first time. The hash is defined as a cryptographic digital fingerprint made by hashing the block header twice through the SHA256 algorithm.
	MinedInBlockHash string `json:"minedInBlockHash"`
	// Represents the hight of the block where this transaction was mined/confirmed for first time. The height is defined as the number of blocks in the blockchain preceding this specific block.
	MinedInBlockHeight int32 `json:"minedInBlockHeight"`
	// Represents a list of recipient addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list.
	Recipients []ListTransactionsByBlockHashResponseItemRecipients `json:"recipients"`
	// Represents a list of sender addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list.
	Senders []ListTransactionsByBlockHashResponseItemSenders `json:"senders"`
	// Defines the exact date/time in Unix Timestamp when this transaction was mined, confirmed or first seen in Mempool, if it is unconfirmed.
	Timestamp int32 `json:"timestamp"`
	// Represents the same as `transactionId` for account-based protocols like Ethereum, while it could be different in UTXO-based protocols like Bitcoin. E.g., in UTXO-based protocols `hash` is different from `transactionId` for SegWit transactions.
	TransactionHash string `json:"transactionHash"`
	// Represents the unique identifier of a transaction, i.e. it could be `transactionId` in UTXO-based protocols like Bitcoin, and transaction `hash` in Ethereum blockchain.
	TransactionId string `json:"transactionId"`
	Fee ListTransactionsByBlockHashResponseItemFee `json:"fee"`
	BlockchainSpecific ListTransactionsByBlockHashResponseItemBlockchainSpecific `json:"blockchainSpecific"`
}

// NewListTransactionsByBlockHashResponseItem instantiates a new ListTransactionsByBlockHashResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactionsByBlockHashResponseItem(index int32, minedInBlockHash string, minedInBlockHeight int32, recipients []ListTransactionsByBlockHashResponseItemRecipients, senders []ListTransactionsByBlockHashResponseItemSenders, timestamp int32, transactionHash string, transactionId string, fee ListTransactionsByBlockHashResponseItemFee, blockchainSpecific ListTransactionsByBlockHashResponseItemBlockchainSpecific) *ListTransactionsByBlockHashResponseItem {
	this := ListTransactionsByBlockHashResponseItem{}
	this.Index = index
	this.MinedInBlockHash = minedInBlockHash
	this.MinedInBlockHeight = minedInBlockHeight
	this.Recipients = recipients
	this.Senders = senders
	this.Timestamp = timestamp
	this.TransactionHash = transactionHash
	this.TransactionId = transactionId
	this.Fee = fee
	this.BlockchainSpecific = blockchainSpecific
	return &this
}

// NewListTransactionsByBlockHashResponseItemWithDefaults instantiates a new ListTransactionsByBlockHashResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactionsByBlockHashResponseItemWithDefaults() *ListTransactionsByBlockHashResponseItem {
	this := ListTransactionsByBlockHashResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *ListTransactionsByBlockHashResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *ListTransactionsByBlockHashResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetMinedInBlockHash returns the MinedInBlockHash field value
func (o *ListTransactionsByBlockHashResponseItem) GetMinedInBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinedInBlockHash
}

// GetMinedInBlockHashOk returns a tuple with the MinedInBlockHash field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashResponseItem) GetMinedInBlockHashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MinedInBlockHash, true
}

// SetMinedInBlockHash sets field value
func (o *ListTransactionsByBlockHashResponseItem) SetMinedInBlockHash(v string) {
	o.MinedInBlockHash = v
}

// GetMinedInBlockHeight returns the MinedInBlockHeight field value
func (o *ListTransactionsByBlockHashResponseItem) GetMinedInBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinedInBlockHeight
}

// GetMinedInBlockHeightOk returns a tuple with the MinedInBlockHeight field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashResponseItem) GetMinedInBlockHeightOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MinedInBlockHeight, true
}

// SetMinedInBlockHeight sets field value
func (o *ListTransactionsByBlockHashResponseItem) SetMinedInBlockHeight(v int32) {
	o.MinedInBlockHeight = v
}

// GetRecipients returns the Recipients field value
func (o *ListTransactionsByBlockHashResponseItem) GetRecipients() []ListTransactionsByBlockHashResponseItemRecipients {
	if o == nil {
		var ret []ListTransactionsByBlockHashResponseItemRecipients
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashResponseItem) GetRecipientsOk() (*[]ListTransactionsByBlockHashResponseItemRecipients, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Recipients, true
}

// SetRecipients sets field value
func (o *ListTransactionsByBlockHashResponseItem) SetRecipients(v []ListTransactionsByBlockHashResponseItemRecipients) {
	o.Recipients = v
}

// GetSenders returns the Senders field value
func (o *ListTransactionsByBlockHashResponseItem) GetSenders() []ListTransactionsByBlockHashResponseItemSenders {
	if o == nil {
		var ret []ListTransactionsByBlockHashResponseItemSenders
		return ret
	}

	return o.Senders
}

// GetSendersOk returns a tuple with the Senders field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashResponseItem) GetSendersOk() (*[]ListTransactionsByBlockHashResponseItemSenders, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Senders, true
}

// SetSenders sets field value
func (o *ListTransactionsByBlockHashResponseItem) SetSenders(v []ListTransactionsByBlockHashResponseItemSenders) {
	o.Senders = v
}

// GetTimestamp returns the Timestamp field value
func (o *ListTransactionsByBlockHashResponseItem) GetTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashResponseItem) GetTimestampOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *ListTransactionsByBlockHashResponseItem) SetTimestamp(v int32) {
	o.Timestamp = v
}

// GetTransactionHash returns the TransactionHash field value
func (o *ListTransactionsByBlockHashResponseItem) GetTransactionHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionHash
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashResponseItem) GetTransactionHashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionHash, true
}

// SetTransactionHash sets field value
func (o *ListTransactionsByBlockHashResponseItem) SetTransactionHash(v string) {
	o.TransactionHash = v
}

// GetTransactionId returns the TransactionId field value
func (o *ListTransactionsByBlockHashResponseItem) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashResponseItem) GetTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *ListTransactionsByBlockHashResponseItem) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetFee returns the Fee field value
func (o *ListTransactionsByBlockHashResponseItem) GetFee() ListTransactionsByBlockHashResponseItemFee {
	if o == nil {
		var ret ListTransactionsByBlockHashResponseItemFee
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashResponseItem) GetFeeOk() (*ListTransactionsByBlockHashResponseItemFee, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *ListTransactionsByBlockHashResponseItem) SetFee(v ListTransactionsByBlockHashResponseItemFee) {
	o.Fee = v
}

// GetBlockchainSpecific returns the BlockchainSpecific field value
func (o *ListTransactionsByBlockHashResponseItem) GetBlockchainSpecific() ListTransactionsByBlockHashResponseItemBlockchainSpecific {
	if o == nil {
		var ret ListTransactionsByBlockHashResponseItemBlockchainSpecific
		return ret
	}

	return o.BlockchainSpecific
}

// GetBlockchainSpecificOk returns a tuple with the BlockchainSpecific field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashResponseItem) GetBlockchainSpecificOk() (*ListTransactionsByBlockHashResponseItemBlockchainSpecific, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BlockchainSpecific, true
}

// SetBlockchainSpecific sets field value
func (o *ListTransactionsByBlockHashResponseItem) SetBlockchainSpecific(v ListTransactionsByBlockHashResponseItemBlockchainSpecific) {
	o.BlockchainSpecific = v
}

func (o ListTransactionsByBlockHashResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["index"] = o.Index
	}
	if true {
		toSerialize["minedInBlockHash"] = o.MinedInBlockHash
	}
	if true {
		toSerialize["minedInBlockHeight"] = o.MinedInBlockHeight
	}
	if true {
		toSerialize["recipients"] = o.Recipients
	}
	if true {
		toSerialize["senders"] = o.Senders
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["transactionHash"] = o.TransactionHash
	}
	if true {
		toSerialize["transactionId"] = o.TransactionId
	}
	if true {
		toSerialize["fee"] = o.Fee
	}
	if true {
		toSerialize["blockchainSpecific"] = o.BlockchainSpecific
	}
	return json.Marshal(toSerialize)
}

type NullableListTransactionsByBlockHashResponseItem struct {
	value *ListTransactionsByBlockHashResponseItem
	isSet bool
}

func (v NullableListTransactionsByBlockHashResponseItem) Get() *ListTransactionsByBlockHashResponseItem {
	return v.value
}

func (v *NullableListTransactionsByBlockHashResponseItem) Set(val *ListTransactionsByBlockHashResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactionsByBlockHashResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactionsByBlockHashResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactionsByBlockHashResponseItem(val *ListTransactionsByBlockHashResponseItem) *NullableListTransactionsByBlockHashResponseItem {
	return &NullableListTransactionsByBlockHashResponseItem{value: val, isSet: true}
}

func (v NullableListTransactionsByBlockHashResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactionsByBlockHashResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

