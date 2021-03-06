# ListUnconfirmedTransactionsByAddressRIRecipientsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The address which receives this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one recipient. | 
**Amount** | **string** | Represents the amount received to this address. | 

## Methods

### NewListUnconfirmedTransactionsByAddressRIRecipientsInner

`func NewListUnconfirmedTransactionsByAddressRIRecipientsInner(address string, amount string, ) *ListUnconfirmedTransactionsByAddressRIRecipientsInner`

NewListUnconfirmedTransactionsByAddressRIRecipientsInner instantiates a new ListUnconfirmedTransactionsByAddressRIRecipientsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUnconfirmedTransactionsByAddressRIRecipientsInnerWithDefaults

`func NewListUnconfirmedTransactionsByAddressRIRecipientsInnerWithDefaults() *ListUnconfirmedTransactionsByAddressRIRecipientsInner`

NewListUnconfirmedTransactionsByAddressRIRecipientsInnerWithDefaults instantiates a new ListUnconfirmedTransactionsByAddressRIRecipientsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ListUnconfirmedTransactionsByAddressRIRecipientsInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ListUnconfirmedTransactionsByAddressRIRecipientsInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ListUnconfirmedTransactionsByAddressRIRecipientsInner) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAmount

`func (o *ListUnconfirmedTransactionsByAddressRIRecipientsInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ListUnconfirmedTransactionsByAddressRIRecipientsInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ListUnconfirmedTransactionsByAddressRIRecipientsInner) SetAmount(v string)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


