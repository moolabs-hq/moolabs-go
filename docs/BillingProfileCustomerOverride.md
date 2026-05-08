# BillingProfileCustomerOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Timestamp of when the resource was created. | [readonly] 
**UpdatedAt** | **time.Time** | Timestamp of when the resource was last updated. | [readonly] 
**BillingProfileId** | Pointer to **string** | The billing profile this override is associated with.  If empty the default profile is looked up dynamically. | [optional] 
**CustomerId** | **string** | The customer id this override is associated with. | 

## Methods

### NewBillingProfileCustomerOverride

`func NewBillingProfileCustomerOverride(createdAt time.Time, updatedAt time.Time, customerId string, ) *BillingProfileCustomerOverride`

NewBillingProfileCustomerOverride instantiates a new BillingProfileCustomerOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingProfileCustomerOverrideWithDefaults

`func NewBillingProfileCustomerOverrideWithDefaults() *BillingProfileCustomerOverride`

NewBillingProfileCustomerOverrideWithDefaults instantiates a new BillingProfileCustomerOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *BillingProfileCustomerOverride) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BillingProfileCustomerOverride) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BillingProfileCustomerOverride) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BillingProfileCustomerOverride) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BillingProfileCustomerOverride) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BillingProfileCustomerOverride) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetBillingProfileId

`func (o *BillingProfileCustomerOverride) GetBillingProfileId() string`

GetBillingProfileId returns the BillingProfileId field if non-nil, zero value otherwise.

### GetBillingProfileIdOk

`func (o *BillingProfileCustomerOverride) GetBillingProfileIdOk() (*string, bool)`

GetBillingProfileIdOk returns a tuple with the BillingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingProfileId

`func (o *BillingProfileCustomerOverride) SetBillingProfileId(v string)`

SetBillingProfileId sets BillingProfileId field to given value.

### HasBillingProfileId

`func (o *BillingProfileCustomerOverride) HasBillingProfileId() bool`

HasBillingProfileId returns a boolean if a field has been set.

### GetCustomerId

`func (o *BillingProfileCustomerOverride) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *BillingProfileCustomerOverride) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *BillingProfileCustomerOverride) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


