# ListZilliqaTransactionsByBlockHashRData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** | Defines how many items should be returned in the response per page basis. | 
**Offset** | **int32** | The starting index of the response items, i.e. where the response should start listing the returned items. | 
**Total** | **int32** | Defines the total number of items returned in the response. | 
**Items** | [**[]ListZilliqaTransactionsByBlockHashRI**](ListZilliqaTransactionsByBlockHashRI.md) |  | 

## Methods

### NewListZilliqaTransactionsByBlockHashRData

`func NewListZilliqaTransactionsByBlockHashRData(limit int32, offset int32, total int32, items []ListZilliqaTransactionsByBlockHashRI, ) *ListZilliqaTransactionsByBlockHashRData`

NewListZilliqaTransactionsByBlockHashRData instantiates a new ListZilliqaTransactionsByBlockHashRData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListZilliqaTransactionsByBlockHashRDataWithDefaults

`func NewListZilliqaTransactionsByBlockHashRDataWithDefaults() *ListZilliqaTransactionsByBlockHashRData`

NewListZilliqaTransactionsByBlockHashRDataWithDefaults instantiates a new ListZilliqaTransactionsByBlockHashRData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *ListZilliqaTransactionsByBlockHashRData) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListZilliqaTransactionsByBlockHashRData) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListZilliqaTransactionsByBlockHashRData) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *ListZilliqaTransactionsByBlockHashRData) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListZilliqaTransactionsByBlockHashRData) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListZilliqaTransactionsByBlockHashRData) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetTotal

`func (o *ListZilliqaTransactionsByBlockHashRData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListZilliqaTransactionsByBlockHashRData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListZilliqaTransactionsByBlockHashRData) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetItems

`func (o *ListZilliqaTransactionsByBlockHashRData) GetItems() []ListZilliqaTransactionsByBlockHashRI`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListZilliqaTransactionsByBlockHashRData) GetItemsOk() (*[]ListZilliqaTransactionsByBlockHashRI, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListZilliqaTransactionsByBlockHashRData) SetItems(v []ListZilliqaTransactionsByBlockHashRI)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


