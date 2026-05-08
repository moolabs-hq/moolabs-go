# InvalidatePortalTokensRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Invalidate a portal token by ID. | [optional] 
**Subject** | Pointer to **string** | Invalidate all portal tokens for a subject. | [optional] 

## Methods

### NewInvalidatePortalTokensRequest

`func NewInvalidatePortalTokensRequest() *InvalidatePortalTokensRequest`

NewInvalidatePortalTokensRequest instantiates a new InvalidatePortalTokensRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvalidatePortalTokensRequestWithDefaults

`func NewInvalidatePortalTokensRequestWithDefaults() *InvalidatePortalTokensRequest`

NewInvalidatePortalTokensRequestWithDefaults instantiates a new InvalidatePortalTokensRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvalidatePortalTokensRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvalidatePortalTokensRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvalidatePortalTokensRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InvalidatePortalTokensRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSubject

`func (o *InvalidatePortalTokensRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *InvalidatePortalTokensRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *InvalidatePortalTokensRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *InvalidatePortalTokensRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


