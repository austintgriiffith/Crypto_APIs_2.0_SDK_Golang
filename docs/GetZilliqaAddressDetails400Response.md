# GetZilliqaAddressDetails400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | Specifies the version of the API that incorporates this endpoint. | 
**RequestId** | **string** | Defines the ID of the request. The &#x60;requestId&#x60; is generated by Crypto APIs and it&#39;s unique for every request. | 
**Context** | Pointer to **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | [optional] 
**Error** | [**GetZilliqaAddressDetailsE400**](GetZilliqaAddressDetailsE400.md) |  | 

## Methods

### NewGetZilliqaAddressDetails400Response

`func NewGetZilliqaAddressDetails400Response(apiVersion string, requestId string, error_ GetZilliqaAddressDetailsE400, ) *GetZilliqaAddressDetails400Response`

NewGetZilliqaAddressDetails400Response instantiates a new GetZilliqaAddressDetails400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetZilliqaAddressDetails400ResponseWithDefaults

`func NewGetZilliqaAddressDetails400ResponseWithDefaults() *GetZilliqaAddressDetails400Response`

NewGetZilliqaAddressDetails400ResponseWithDefaults instantiates a new GetZilliqaAddressDetails400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *GetZilliqaAddressDetails400Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *GetZilliqaAddressDetails400Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *GetZilliqaAddressDetails400Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetRequestId

`func (o *GetZilliqaAddressDetails400Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetZilliqaAddressDetails400Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetZilliqaAddressDetails400Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetContext

`func (o *GetZilliqaAddressDetails400Response) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *GetZilliqaAddressDetails400Response) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *GetZilliqaAddressDetails400Response) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *GetZilliqaAddressDetails400Response) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetError

`func (o *GetZilliqaAddressDetails400Response) GetError() GetZilliqaAddressDetailsE400`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetZilliqaAddressDetails400Response) GetErrorOk() (*GetZilliqaAddressDetailsE400, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetZilliqaAddressDetails400Response) SetError(v GetZilliqaAddressDetailsE400)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


