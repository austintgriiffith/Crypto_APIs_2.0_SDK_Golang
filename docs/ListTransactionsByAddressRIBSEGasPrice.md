# ListTransactionsByAddressRIBSEGasPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Represents the price offered to the miner to purchase this amount of gas | 
**Unit** | **string** | Defines the unit of the gas price amount, e.g. BTC, ETH, XRP. | 

## Methods

### NewListTransactionsByAddressRIBSEGasPrice

`func NewListTransactionsByAddressRIBSEGasPrice(amount string, unit string, ) *ListTransactionsByAddressRIBSEGasPrice`

NewListTransactionsByAddressRIBSEGasPrice instantiates a new ListTransactionsByAddressRIBSEGasPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByAddressRIBSEGasPriceWithDefaults

`func NewListTransactionsByAddressRIBSEGasPriceWithDefaults() *ListTransactionsByAddressRIBSEGasPrice`

NewListTransactionsByAddressRIBSEGasPriceWithDefaults instantiates a new ListTransactionsByAddressRIBSEGasPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ListTransactionsByAddressRIBSEGasPrice) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ListTransactionsByAddressRIBSEGasPrice) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ListTransactionsByAddressRIBSEGasPrice) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetUnit

`func (o *ListTransactionsByAddressRIBSEGasPrice) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ListTransactionsByAddressRIBSEGasPrice) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ListTransactionsByAddressRIBSEGasPrice) SetUnit(v string)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

