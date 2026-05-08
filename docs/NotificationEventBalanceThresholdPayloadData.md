# NotificationEventBalanceThresholdPayloadData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entitlement** | [**EntitlementMetered**](EntitlementMetered.md) |  | [readonly] 
**Feature** | [**Feature**](Feature.md) |  | [readonly] 
**Subject** | [**Subject**](Subject.md) |  | [readonly] 
**Value** | [**EntitlementValue**](EntitlementValue.md) |  | [readonly] 
**Customer** | Pointer to [**Customer**](Customer.md) |  | [optional] [readonly] 
**Threshold** | [**NotificationRuleBalanceThresholdValue**](NotificationRuleBalanceThresholdValue.md) |  | [readonly] 

## Methods

### NewNotificationEventBalanceThresholdPayloadData

`func NewNotificationEventBalanceThresholdPayloadData(entitlement EntitlementMetered, feature Feature, subject Subject, value EntitlementValue, threshold NotificationRuleBalanceThresholdValue, ) *NotificationEventBalanceThresholdPayloadData`

NewNotificationEventBalanceThresholdPayloadData instantiates a new NotificationEventBalanceThresholdPayloadData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationEventBalanceThresholdPayloadDataWithDefaults

`func NewNotificationEventBalanceThresholdPayloadDataWithDefaults() *NotificationEventBalanceThresholdPayloadData`

NewNotificationEventBalanceThresholdPayloadDataWithDefaults instantiates a new NotificationEventBalanceThresholdPayloadData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlement

`func (o *NotificationEventBalanceThresholdPayloadData) GetEntitlement() EntitlementMetered`

GetEntitlement returns the Entitlement field if non-nil, zero value otherwise.

### GetEntitlementOk

`func (o *NotificationEventBalanceThresholdPayloadData) GetEntitlementOk() (*EntitlementMetered, bool)`

GetEntitlementOk returns a tuple with the Entitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlement

`func (o *NotificationEventBalanceThresholdPayloadData) SetEntitlement(v EntitlementMetered)`

SetEntitlement sets Entitlement field to given value.


### GetFeature

`func (o *NotificationEventBalanceThresholdPayloadData) GetFeature() Feature`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *NotificationEventBalanceThresholdPayloadData) GetFeatureOk() (*Feature, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *NotificationEventBalanceThresholdPayloadData) SetFeature(v Feature)`

SetFeature sets Feature field to given value.


### GetSubject

`func (o *NotificationEventBalanceThresholdPayloadData) GetSubject() Subject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *NotificationEventBalanceThresholdPayloadData) GetSubjectOk() (*Subject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *NotificationEventBalanceThresholdPayloadData) SetSubject(v Subject)`

SetSubject sets Subject field to given value.


### GetValue

`func (o *NotificationEventBalanceThresholdPayloadData) GetValue() EntitlementValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NotificationEventBalanceThresholdPayloadData) GetValueOk() (*EntitlementValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NotificationEventBalanceThresholdPayloadData) SetValue(v EntitlementValue)`

SetValue sets Value field to given value.


### GetCustomer

`func (o *NotificationEventBalanceThresholdPayloadData) GetCustomer() Customer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *NotificationEventBalanceThresholdPayloadData) GetCustomerOk() (*Customer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *NotificationEventBalanceThresholdPayloadData) SetCustomer(v Customer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *NotificationEventBalanceThresholdPayloadData) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetThreshold

`func (o *NotificationEventBalanceThresholdPayloadData) GetThreshold() NotificationRuleBalanceThresholdValue`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *NotificationEventBalanceThresholdPayloadData) GetThresholdOk() (*NotificationRuleBalanceThresholdValue, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *NotificationEventBalanceThresholdPayloadData) SetThreshold(v NotificationRuleBalanceThresholdValue)`

SetThreshold sets Threshold field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


