# RemittanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**ExternalReference** | Pointer to **string** |  | [optional] 
**Source** | **string** |  | 
**PayerName** | Pointer to **string** |  | [optional] 
**PayerAccountRef** | Pointer to **string** |  | [optional] 
**RawReferenceText** | Pointer to **string** |  | [optional] 
**AmountMicros** | **int32** |  | 
**CurrencyCode** | **string** |  | 
**ReceivedAt** | **time.Time** |  | 
**Status** | **string** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewRemittanceResponse

`func NewRemittanceResponse(id string, tenantId string, source string, amountMicros int32, currencyCode string, receivedAt time.Time, status string, createdAt time.Time, updatedAt time.Time, ) *RemittanceResponse`

NewRemittanceResponse instantiates a new RemittanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemittanceResponseWithDefaults

`func NewRemittanceResponseWithDefaults() *RemittanceResponse`

NewRemittanceResponseWithDefaults instantiates a new RemittanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RemittanceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemittanceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemittanceResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *RemittanceResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *RemittanceResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *RemittanceResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetExternalReference

`func (o *RemittanceResponse) GetExternalReference() string`

GetExternalReference returns the ExternalReference field if non-nil, zero value otherwise.

### GetExternalReferenceOk

`func (o *RemittanceResponse) GetExternalReferenceOk() (*string, bool)`

GetExternalReferenceOk returns a tuple with the ExternalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReference

`func (o *RemittanceResponse) SetExternalReference(v string)`

SetExternalReference sets ExternalReference field to given value.

### HasExternalReference

`func (o *RemittanceResponse) HasExternalReference() bool`

HasExternalReference returns a boolean if a field has been set.

### GetSource

`func (o *RemittanceResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RemittanceResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RemittanceResponse) SetSource(v string)`

SetSource sets Source field to given value.


### GetPayerName

`func (o *RemittanceResponse) GetPayerName() string`

GetPayerName returns the PayerName field if non-nil, zero value otherwise.

### GetPayerNameOk

`func (o *RemittanceResponse) GetPayerNameOk() (*string, bool)`

GetPayerNameOk returns a tuple with the PayerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerName

`func (o *RemittanceResponse) SetPayerName(v string)`

SetPayerName sets PayerName field to given value.

### HasPayerName

`func (o *RemittanceResponse) HasPayerName() bool`

HasPayerName returns a boolean if a field has been set.

### GetPayerAccountRef

`func (o *RemittanceResponse) GetPayerAccountRef() string`

GetPayerAccountRef returns the PayerAccountRef field if non-nil, zero value otherwise.

### GetPayerAccountRefOk

`func (o *RemittanceResponse) GetPayerAccountRefOk() (*string, bool)`

GetPayerAccountRefOk returns a tuple with the PayerAccountRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerAccountRef

`func (o *RemittanceResponse) SetPayerAccountRef(v string)`

SetPayerAccountRef sets PayerAccountRef field to given value.

### HasPayerAccountRef

`func (o *RemittanceResponse) HasPayerAccountRef() bool`

HasPayerAccountRef returns a boolean if a field has been set.

### GetRawReferenceText

`func (o *RemittanceResponse) GetRawReferenceText() string`

GetRawReferenceText returns the RawReferenceText field if non-nil, zero value otherwise.

### GetRawReferenceTextOk

`func (o *RemittanceResponse) GetRawReferenceTextOk() (*string, bool)`

GetRawReferenceTextOk returns a tuple with the RawReferenceText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawReferenceText

`func (o *RemittanceResponse) SetRawReferenceText(v string)`

SetRawReferenceText sets RawReferenceText field to given value.

### HasRawReferenceText

`func (o *RemittanceResponse) HasRawReferenceText() bool`

HasRawReferenceText returns a boolean if a field has been set.

### GetAmountMicros

`func (o *RemittanceResponse) GetAmountMicros() int32`

GetAmountMicros returns the AmountMicros field if non-nil, zero value otherwise.

### GetAmountMicrosOk

`func (o *RemittanceResponse) GetAmountMicrosOk() (*int32, bool)`

GetAmountMicrosOk returns a tuple with the AmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMicros

`func (o *RemittanceResponse) SetAmountMicros(v int32)`

SetAmountMicros sets AmountMicros field to given value.


### GetCurrencyCode

`func (o *RemittanceResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *RemittanceResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *RemittanceResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetReceivedAt

`func (o *RemittanceResponse) GetReceivedAt() time.Time`

GetReceivedAt returns the ReceivedAt field if non-nil, zero value otherwise.

### GetReceivedAtOk

`func (o *RemittanceResponse) GetReceivedAtOk() (*time.Time, bool)`

GetReceivedAtOk returns a tuple with the ReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAt

`func (o *RemittanceResponse) SetReceivedAt(v time.Time)`

SetReceivedAt sets ReceivedAt field to given value.


### GetStatus

`func (o *RemittanceResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RemittanceResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RemittanceResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMetadata

`func (o *RemittanceResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RemittanceResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RemittanceResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RemittanceResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RemittanceResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RemittanceResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RemittanceResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *RemittanceResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RemittanceResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RemittanceResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


