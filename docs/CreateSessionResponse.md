# CreateSessionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**QuoteId** | **string** |  | 
**QuoteVersion** | **int32** |  | 
**State** | **string** |  | 
**Scratchpad** | **map[string]interface{}** |  | 

## Methods

### NewCreateSessionResponse

`func NewCreateSessionResponse(id string, tenantId string, quoteId string, quoteVersion int32, state string, scratchpad map[string]interface{}, ) *CreateSessionResponse`

NewCreateSessionResponse instantiates a new CreateSessionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSessionResponseWithDefaults

`func NewCreateSessionResponseWithDefaults() *CreateSessionResponse`

NewCreateSessionResponseWithDefaults instantiates a new CreateSessionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateSessionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateSessionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateSessionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *CreateSessionResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateSessionResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateSessionResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetQuoteId

`func (o *CreateSessionResponse) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *CreateSessionResponse) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *CreateSessionResponse) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetQuoteVersion

`func (o *CreateSessionResponse) GetQuoteVersion() int32`

GetQuoteVersion returns the QuoteVersion field if non-nil, zero value otherwise.

### GetQuoteVersionOk

`func (o *CreateSessionResponse) GetQuoteVersionOk() (*int32, bool)`

GetQuoteVersionOk returns a tuple with the QuoteVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteVersion

`func (o *CreateSessionResponse) SetQuoteVersion(v int32)`

SetQuoteVersion sets QuoteVersion field to given value.


### GetState

`func (o *CreateSessionResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CreateSessionResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CreateSessionResponse) SetState(v string)`

SetState sets State field to given value.


### GetScratchpad

`func (o *CreateSessionResponse) GetScratchpad() map[string]interface{}`

GetScratchpad returns the Scratchpad field if non-nil, zero value otherwise.

### GetScratchpadOk

`func (o *CreateSessionResponse) GetScratchpadOk() (*map[string]interface{}, bool)`

GetScratchpadOk returns a tuple with the Scratchpad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScratchpad

`func (o *CreateSessionResponse) SetScratchpad(v map[string]interface{})`

SetScratchpad sets Scratchpad field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


