# BillingDiscountMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrelationId** | Pointer to **string** | Correlation ID for the discount.  This is used to link discounts across different invoices (progressive billing use case).  If not provided, the invoicing engine will auto-generate one. When editing an invoice line, please make sure to keep the same correlation ID of the discount or in progressive billing setups the discount amounts might be incorrect. | [optional] 

## Methods

### NewBillingDiscountMetadata

`func NewBillingDiscountMetadata() *BillingDiscountMetadata`

NewBillingDiscountMetadata instantiates a new BillingDiscountMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingDiscountMetadataWithDefaults

`func NewBillingDiscountMetadataWithDefaults() *BillingDiscountMetadata`

NewBillingDiscountMetadataWithDefaults instantiates a new BillingDiscountMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrelationId

`func (o *BillingDiscountMetadata) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *BillingDiscountMetadata) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *BillingDiscountMetadata) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *BillingDiscountMetadata) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


