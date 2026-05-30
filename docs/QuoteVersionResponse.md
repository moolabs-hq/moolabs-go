# QuoteVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuoteId** | **string** |  | 
**TenantId** | **string** |  | 
**Version** | **int32** |  | 
**State** | **string** |  | 
**LineItems** | **[]map[string]interface{}** |  | 
**CommercialTerms** | **map[string]interface{}** |  | 
**CreditTerms** | **map[string]interface{}** |  | 
**SourceVersions** | Pointer to **map[string]interface{}** |  | [optional] 
**PricingSnapshotId** | **string** |  | 
**QuoteVersionDigest** | Pointer to **string** |  | [optional] 
**BookingTrigger** | Pointer to **string** |  | [optional] 
**LockedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewQuoteVersionResponse

`func NewQuoteVersionResponse(quoteId string, tenantId string, version int32, state string, lineItems []map[string]interface{}, commercialTerms map[string]interface{}, creditTerms map[string]interface{}, pricingSnapshotId string, ) *QuoteVersionResponse`

NewQuoteVersionResponse instantiates a new QuoteVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteVersionResponseWithDefaults

`func NewQuoteVersionResponseWithDefaults() *QuoteVersionResponse`

NewQuoteVersionResponseWithDefaults instantiates a new QuoteVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuoteId

`func (o *QuoteVersionResponse) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *QuoteVersionResponse) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *QuoteVersionResponse) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetTenantId

`func (o *QuoteVersionResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *QuoteVersionResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *QuoteVersionResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetVersion

`func (o *QuoteVersionResponse) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *QuoteVersionResponse) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *QuoteVersionResponse) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetState

`func (o *QuoteVersionResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *QuoteVersionResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *QuoteVersionResponse) SetState(v string)`

SetState sets State field to given value.


### GetLineItems

`func (o *QuoteVersionResponse) GetLineItems() []map[string]interface{}`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *QuoteVersionResponse) GetLineItemsOk() (*[]map[string]interface{}, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *QuoteVersionResponse) SetLineItems(v []map[string]interface{})`

SetLineItems sets LineItems field to given value.


### GetCommercialTerms

`func (o *QuoteVersionResponse) GetCommercialTerms() map[string]interface{}`

GetCommercialTerms returns the CommercialTerms field if non-nil, zero value otherwise.

### GetCommercialTermsOk

`func (o *QuoteVersionResponse) GetCommercialTermsOk() (*map[string]interface{}, bool)`

GetCommercialTermsOk returns a tuple with the CommercialTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialTerms

`func (o *QuoteVersionResponse) SetCommercialTerms(v map[string]interface{})`

SetCommercialTerms sets CommercialTerms field to given value.


### GetCreditTerms

`func (o *QuoteVersionResponse) GetCreditTerms() map[string]interface{}`

GetCreditTerms returns the CreditTerms field if non-nil, zero value otherwise.

### GetCreditTermsOk

`func (o *QuoteVersionResponse) GetCreditTermsOk() (*map[string]interface{}, bool)`

GetCreditTermsOk returns a tuple with the CreditTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditTerms

`func (o *QuoteVersionResponse) SetCreditTerms(v map[string]interface{})`

SetCreditTerms sets CreditTerms field to given value.


### GetSourceVersions

`func (o *QuoteVersionResponse) GetSourceVersions() map[string]interface{}`

GetSourceVersions returns the SourceVersions field if non-nil, zero value otherwise.

### GetSourceVersionsOk

`func (o *QuoteVersionResponse) GetSourceVersionsOk() (*map[string]interface{}, bool)`

GetSourceVersionsOk returns a tuple with the SourceVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVersions

`func (o *QuoteVersionResponse) SetSourceVersions(v map[string]interface{})`

SetSourceVersions sets SourceVersions field to given value.

### HasSourceVersions

`func (o *QuoteVersionResponse) HasSourceVersions() bool`

HasSourceVersions returns a boolean if a field has been set.

### GetPricingSnapshotId

`func (o *QuoteVersionResponse) GetPricingSnapshotId() string`

GetPricingSnapshotId returns the PricingSnapshotId field if non-nil, zero value otherwise.

### GetPricingSnapshotIdOk

`func (o *QuoteVersionResponse) GetPricingSnapshotIdOk() (*string, bool)`

GetPricingSnapshotIdOk returns a tuple with the PricingSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingSnapshotId

`func (o *QuoteVersionResponse) SetPricingSnapshotId(v string)`

SetPricingSnapshotId sets PricingSnapshotId field to given value.


### GetQuoteVersionDigest

`func (o *QuoteVersionResponse) GetQuoteVersionDigest() string`

GetQuoteVersionDigest returns the QuoteVersionDigest field if non-nil, zero value otherwise.

### GetQuoteVersionDigestOk

`func (o *QuoteVersionResponse) GetQuoteVersionDigestOk() (*string, bool)`

GetQuoteVersionDigestOk returns a tuple with the QuoteVersionDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteVersionDigest

`func (o *QuoteVersionResponse) SetQuoteVersionDigest(v string)`

SetQuoteVersionDigest sets QuoteVersionDigest field to given value.

### HasQuoteVersionDigest

`func (o *QuoteVersionResponse) HasQuoteVersionDigest() bool`

HasQuoteVersionDigest returns a boolean if a field has been set.

### GetBookingTrigger

`func (o *QuoteVersionResponse) GetBookingTrigger() string`

GetBookingTrigger returns the BookingTrigger field if non-nil, zero value otherwise.

### GetBookingTriggerOk

`func (o *QuoteVersionResponse) GetBookingTriggerOk() (*string, bool)`

GetBookingTriggerOk returns a tuple with the BookingTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookingTrigger

`func (o *QuoteVersionResponse) SetBookingTrigger(v string)`

SetBookingTrigger sets BookingTrigger field to given value.

### HasBookingTrigger

`func (o *QuoteVersionResponse) HasBookingTrigger() bool`

HasBookingTrigger returns a boolean if a field has been set.

### GetLockedAt

`func (o *QuoteVersionResponse) GetLockedAt() string`

GetLockedAt returns the LockedAt field if non-nil, zero value otherwise.

### GetLockedAtOk

`func (o *QuoteVersionResponse) GetLockedAtOk() (*string, bool)`

GetLockedAtOk returns a tuple with the LockedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedAt

`func (o *QuoteVersionResponse) SetLockedAt(v string)`

SetLockedAt sets LockedAt field to given value.

### HasLockedAt

`func (o *QuoteVersionResponse) HasLockedAt() bool`

HasLockedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


