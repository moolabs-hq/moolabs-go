# NotificationEventEntitlementValuePayloadBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entitlement** | [**EntitlementMetered**](EntitlementMetered.md) |  | [readonly] 
**Feature** | [**Feature**](Feature.md) |  | [readonly] 
**Subject** | [**Subject**](Subject.md) |  | [readonly] 
**Value** | [**EntitlementValue**](EntitlementValue.md) |  | [readonly] 
**Customer** | Pointer to [**Customer**](Customer.md) |  | [optional] [readonly] 

## Methods

### NewNotificationEventEntitlementValuePayloadBase

`func NewNotificationEventEntitlementValuePayloadBase(entitlement EntitlementMetered, feature Feature, subject Subject, value EntitlementValue, ) *NotificationEventEntitlementValuePayloadBase`

NewNotificationEventEntitlementValuePayloadBase instantiates a new NotificationEventEntitlementValuePayloadBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationEventEntitlementValuePayloadBaseWithDefaults

`func NewNotificationEventEntitlementValuePayloadBaseWithDefaults() *NotificationEventEntitlementValuePayloadBase`

NewNotificationEventEntitlementValuePayloadBaseWithDefaults instantiates a new NotificationEventEntitlementValuePayloadBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlement

`func (o *NotificationEventEntitlementValuePayloadBase) GetEntitlement() EntitlementMetered`

GetEntitlement returns the Entitlement field if non-nil, zero value otherwise.

### GetEntitlementOk

`func (o *NotificationEventEntitlementValuePayloadBase) GetEntitlementOk() (*EntitlementMetered, bool)`

GetEntitlementOk returns a tuple with the Entitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlement

`func (o *NotificationEventEntitlementValuePayloadBase) SetEntitlement(v EntitlementMetered)`

SetEntitlement sets Entitlement field to given value.


### GetFeature

`func (o *NotificationEventEntitlementValuePayloadBase) GetFeature() Feature`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *NotificationEventEntitlementValuePayloadBase) GetFeatureOk() (*Feature, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *NotificationEventEntitlementValuePayloadBase) SetFeature(v Feature)`

SetFeature sets Feature field to given value.


### GetSubject

`func (o *NotificationEventEntitlementValuePayloadBase) GetSubject() Subject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *NotificationEventEntitlementValuePayloadBase) GetSubjectOk() (*Subject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *NotificationEventEntitlementValuePayloadBase) SetSubject(v Subject)`

SetSubject sets Subject field to given value.


### GetValue

`func (o *NotificationEventEntitlementValuePayloadBase) GetValue() EntitlementValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NotificationEventEntitlementValuePayloadBase) GetValueOk() (*EntitlementValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NotificationEventEntitlementValuePayloadBase) SetValue(v EntitlementValue)`

SetValue sets Value field to given value.


### GetCustomer

`func (o *NotificationEventEntitlementValuePayloadBase) GetCustomer() Customer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *NotificationEventEntitlementValuePayloadBase) GetCustomerOk() (*Customer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *NotificationEventEntitlementValuePayloadBase) SetCustomer(v Customer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *NotificationEventEntitlementValuePayloadBase) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


