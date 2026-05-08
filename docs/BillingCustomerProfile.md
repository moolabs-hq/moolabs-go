# BillingCustomerProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supplier** | [**BillingParty**](BillingParty.md) | The name and contact information for the supplier this billing profile represents | [readonly] 
**Workflow** | [**BillingWorkflow**](BillingWorkflow.md) | The billing workflow settings for this profile | [readonly] 
**Apps** | [**BillingProfileAppsOrReference**](BillingProfileAppsOrReference.md) | The applications used by this billing profile.  Expand settings govern if this includes the whole app object or just the ID references. | [readonly] 

## Methods

### NewBillingCustomerProfile

`func NewBillingCustomerProfile(supplier BillingParty, workflow BillingWorkflow, apps BillingProfileAppsOrReference, ) *BillingCustomerProfile`

NewBillingCustomerProfile instantiates a new BillingCustomerProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingCustomerProfileWithDefaults

`func NewBillingCustomerProfileWithDefaults() *BillingCustomerProfile`

NewBillingCustomerProfileWithDefaults instantiates a new BillingCustomerProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupplier

`func (o *BillingCustomerProfile) GetSupplier() BillingParty`

GetSupplier returns the Supplier field if non-nil, zero value otherwise.

### GetSupplierOk

`func (o *BillingCustomerProfile) GetSupplierOk() (*BillingParty, bool)`

GetSupplierOk returns a tuple with the Supplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplier

`func (o *BillingCustomerProfile) SetSupplier(v BillingParty)`

SetSupplier sets Supplier field to given value.


### GetWorkflow

`func (o *BillingCustomerProfile) GetWorkflow() BillingWorkflow`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *BillingCustomerProfile) GetWorkflowOk() (*BillingWorkflow, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *BillingCustomerProfile) SetWorkflow(v BillingWorkflow)`

SetWorkflow sets Workflow field to given value.


### GetApps

`func (o *BillingCustomerProfile) GetApps() BillingProfileAppsOrReference`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *BillingCustomerProfile) GetAppsOk() (*BillingProfileAppsOrReference, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *BillingCustomerProfile) SetApps(v BillingProfileAppsOrReference)`

SetApps sets Apps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


