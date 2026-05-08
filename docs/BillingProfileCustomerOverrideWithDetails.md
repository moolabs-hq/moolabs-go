# BillingProfileCustomerOverrideWithDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerOverride** | Pointer to [**BillingProfileCustomerOverride**](BillingProfileCustomerOverride.md) | The customer override values.  If empty the merged values are calculated based on the default profile. | [optional] 
**BaseBillingProfileId** | **string** | The billing profile the customerProfile is associated with at the time of query.  customerOverride contains the explicit mapping set in the customer override object. If that is empty, then the baseBillingProfileId is the default profile. | 
**CustomerProfile** | Pointer to [**BillingCustomerProfile**](BillingCustomerProfile.md) | Merged billing profile with the customer specific overrides. | [optional] 
**Customer** | Pointer to [**Customer**](Customer.md) | The customer this override belongs to. | [optional] 

## Methods

### NewBillingProfileCustomerOverrideWithDetails

`func NewBillingProfileCustomerOverrideWithDetails(baseBillingProfileId string, ) *BillingProfileCustomerOverrideWithDetails`

NewBillingProfileCustomerOverrideWithDetails instantiates a new BillingProfileCustomerOverrideWithDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingProfileCustomerOverrideWithDetailsWithDefaults

`func NewBillingProfileCustomerOverrideWithDetailsWithDefaults() *BillingProfileCustomerOverrideWithDetails`

NewBillingProfileCustomerOverrideWithDetailsWithDefaults instantiates a new BillingProfileCustomerOverrideWithDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerOverride

`func (o *BillingProfileCustomerOverrideWithDetails) GetCustomerOverride() BillingProfileCustomerOverride`

GetCustomerOverride returns the CustomerOverride field if non-nil, zero value otherwise.

### GetCustomerOverrideOk

`func (o *BillingProfileCustomerOverrideWithDetails) GetCustomerOverrideOk() (*BillingProfileCustomerOverride, bool)`

GetCustomerOverrideOk returns a tuple with the CustomerOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerOverride

`func (o *BillingProfileCustomerOverrideWithDetails) SetCustomerOverride(v BillingProfileCustomerOverride)`

SetCustomerOverride sets CustomerOverride field to given value.

### HasCustomerOverride

`func (o *BillingProfileCustomerOverrideWithDetails) HasCustomerOverride() bool`

HasCustomerOverride returns a boolean if a field has been set.

### GetBaseBillingProfileId

`func (o *BillingProfileCustomerOverrideWithDetails) GetBaseBillingProfileId() string`

GetBaseBillingProfileId returns the BaseBillingProfileId field if non-nil, zero value otherwise.

### GetBaseBillingProfileIdOk

`func (o *BillingProfileCustomerOverrideWithDetails) GetBaseBillingProfileIdOk() (*string, bool)`

GetBaseBillingProfileIdOk returns a tuple with the BaseBillingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseBillingProfileId

`func (o *BillingProfileCustomerOverrideWithDetails) SetBaseBillingProfileId(v string)`

SetBaseBillingProfileId sets BaseBillingProfileId field to given value.


### GetCustomerProfile

`func (o *BillingProfileCustomerOverrideWithDetails) GetCustomerProfile() BillingCustomerProfile`

GetCustomerProfile returns the CustomerProfile field if non-nil, zero value otherwise.

### GetCustomerProfileOk

`func (o *BillingProfileCustomerOverrideWithDetails) GetCustomerProfileOk() (*BillingCustomerProfile, bool)`

GetCustomerProfileOk returns a tuple with the CustomerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerProfile

`func (o *BillingProfileCustomerOverrideWithDetails) SetCustomerProfile(v BillingCustomerProfile)`

SetCustomerProfile sets CustomerProfile field to given value.

### HasCustomerProfile

`func (o *BillingProfileCustomerOverrideWithDetails) HasCustomerProfile() bool`

HasCustomerProfile returns a boolean if a field has been set.

### GetCustomer

`func (o *BillingProfileCustomerOverrideWithDetails) GetCustomer() Customer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *BillingProfileCustomerOverrideWithDetails) GetCustomerOk() (*Customer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *BillingProfileCustomerOverrideWithDetails) SetCustomer(v Customer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *BillingProfileCustomerOverrideWithDetails) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


