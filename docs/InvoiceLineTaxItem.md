# InvoiceLineTaxItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**TaxConfig**](TaxConfig.md) | Tax provider configuration. | [optional] [readonly] 
**Percent** | Pointer to **float64** | Percent defines the percentage set manually or determined from the rate key (calculated if rate present). A nil percent implies that this tax combo is **exempt** from tax.\&quot;) | [optional] [readonly] 
**Surcharge** | Pointer to **string** | Some countries require an additional surcharge (calculated if rate present). | [optional] [readonly] 
**Behavior** | Pointer to [**InvoiceLineTaxBehavior**](InvoiceLineTaxBehavior.md) | Is the tax item inclusive or exclusive of the base amount. | [optional] [readonly] 

## Methods

### NewInvoiceLineTaxItem

`func NewInvoiceLineTaxItem() *InvoiceLineTaxItem`

NewInvoiceLineTaxItem instantiates a new InvoiceLineTaxItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceLineTaxItemWithDefaults

`func NewInvoiceLineTaxItemWithDefaults() *InvoiceLineTaxItem`

NewInvoiceLineTaxItemWithDefaults instantiates a new InvoiceLineTaxItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *InvoiceLineTaxItem) GetConfig() TaxConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *InvoiceLineTaxItem) GetConfigOk() (*TaxConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *InvoiceLineTaxItem) SetConfig(v TaxConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *InvoiceLineTaxItem) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetPercent

`func (o *InvoiceLineTaxItem) GetPercent() float64`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *InvoiceLineTaxItem) GetPercentOk() (*float64, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *InvoiceLineTaxItem) SetPercent(v float64)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *InvoiceLineTaxItem) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetSurcharge

`func (o *InvoiceLineTaxItem) GetSurcharge() string`

GetSurcharge returns the Surcharge field if non-nil, zero value otherwise.

### GetSurchargeOk

`func (o *InvoiceLineTaxItem) GetSurchargeOk() (*string, bool)`

GetSurchargeOk returns a tuple with the Surcharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurcharge

`func (o *InvoiceLineTaxItem) SetSurcharge(v string)`

SetSurcharge sets Surcharge field to given value.

### HasSurcharge

`func (o *InvoiceLineTaxItem) HasSurcharge() bool`

HasSurcharge returns a boolean if a field has been set.

### GetBehavior

`func (o *InvoiceLineTaxItem) GetBehavior() InvoiceLineTaxBehavior`

GetBehavior returns the Behavior field if non-nil, zero value otherwise.

### GetBehaviorOk

`func (o *InvoiceLineTaxItem) GetBehaviorOk() (*InvoiceLineTaxBehavior, bool)`

GetBehaviorOk returns a tuple with the Behavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehavior

`func (o *InvoiceLineTaxItem) SetBehavior(v InvoiceLineTaxBehavior)`

SetBehavior sets Behavior field to given value.

### HasBehavior

`func (o *InvoiceLineTaxItem) HasBehavior() bool`

HasBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


