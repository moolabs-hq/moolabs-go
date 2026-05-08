# CustomerAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entitlements** | [**map[string]EntitlementValue**](EntitlementValue.md) | Map of entitlements the customer has access to. The key is the feature key, the value is the entitlement value + the entitlement ID. | [readonly] 

## Methods

### NewCustomerAccess

`func NewCustomerAccess(entitlements map[string]EntitlementValue, ) *CustomerAccess`

NewCustomerAccess instantiates a new CustomerAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerAccessWithDefaults

`func NewCustomerAccessWithDefaults() *CustomerAccess`

NewCustomerAccessWithDefaults instantiates a new CustomerAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlements

`func (o *CustomerAccess) GetEntitlements() map[string]EntitlementValue`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *CustomerAccess) GetEntitlementsOk() (*map[string]EntitlementValue, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *CustomerAccess) SetEntitlements(v map[string]EntitlementValue)`

SetEntitlements sets Entitlements field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


