# PaymentItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**NetsuiteId** | **string** |  | 
**CustomerInternalId** | **string** |  | 
**PaymentAmount** | Pointer to **string** |  | [optional] 
**AppliedInvoiceId** | **string** |  | 
**AppliedAmount** | Pointer to **string** |  | [optional] 
**TranDate** | Pointer to **string** |  | [optional] 
**NsLastModified** | Pointer to **time.Time** |  | [optional] 
**SyncedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPaymentItem

`func NewPaymentItem(id string, netsuiteId string, customerInternalId string, appliedInvoiceId string, ) *PaymentItem`

NewPaymentItem instantiates a new PaymentItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentItemWithDefaults

`func NewPaymentItemWithDefaults() *PaymentItem`

NewPaymentItemWithDefaults instantiates a new PaymentItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentItem) SetId(v string)`

SetId sets Id field to given value.


### GetNetsuiteId

`func (o *PaymentItem) GetNetsuiteId() string`

GetNetsuiteId returns the NetsuiteId field if non-nil, zero value otherwise.

### GetNetsuiteIdOk

`func (o *PaymentItem) GetNetsuiteIdOk() (*string, bool)`

GetNetsuiteIdOk returns a tuple with the NetsuiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetsuiteId

`func (o *PaymentItem) SetNetsuiteId(v string)`

SetNetsuiteId sets NetsuiteId field to given value.


### GetCustomerInternalId

`func (o *PaymentItem) GetCustomerInternalId() string`

GetCustomerInternalId returns the CustomerInternalId field if non-nil, zero value otherwise.

### GetCustomerInternalIdOk

`func (o *PaymentItem) GetCustomerInternalIdOk() (*string, bool)`

GetCustomerInternalIdOk returns a tuple with the CustomerInternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerInternalId

`func (o *PaymentItem) SetCustomerInternalId(v string)`

SetCustomerInternalId sets CustomerInternalId field to given value.


### GetPaymentAmount

`func (o *PaymentItem) GetPaymentAmount() string`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *PaymentItem) GetPaymentAmountOk() (*string, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *PaymentItem) SetPaymentAmount(v string)`

SetPaymentAmount sets PaymentAmount field to given value.

### HasPaymentAmount

`func (o *PaymentItem) HasPaymentAmount() bool`

HasPaymentAmount returns a boolean if a field has been set.

### GetAppliedInvoiceId

`func (o *PaymentItem) GetAppliedInvoiceId() string`

GetAppliedInvoiceId returns the AppliedInvoiceId field if non-nil, zero value otherwise.

### GetAppliedInvoiceIdOk

`func (o *PaymentItem) GetAppliedInvoiceIdOk() (*string, bool)`

GetAppliedInvoiceIdOk returns a tuple with the AppliedInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedInvoiceId

`func (o *PaymentItem) SetAppliedInvoiceId(v string)`

SetAppliedInvoiceId sets AppliedInvoiceId field to given value.


### GetAppliedAmount

`func (o *PaymentItem) GetAppliedAmount() string`

GetAppliedAmount returns the AppliedAmount field if non-nil, zero value otherwise.

### GetAppliedAmountOk

`func (o *PaymentItem) GetAppliedAmountOk() (*string, bool)`

GetAppliedAmountOk returns a tuple with the AppliedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedAmount

`func (o *PaymentItem) SetAppliedAmount(v string)`

SetAppliedAmount sets AppliedAmount field to given value.

### HasAppliedAmount

`func (o *PaymentItem) HasAppliedAmount() bool`

HasAppliedAmount returns a boolean if a field has been set.

### GetTranDate

`func (o *PaymentItem) GetTranDate() string`

GetTranDate returns the TranDate field if non-nil, zero value otherwise.

### GetTranDateOk

`func (o *PaymentItem) GetTranDateOk() (*string, bool)`

GetTranDateOk returns a tuple with the TranDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranDate

`func (o *PaymentItem) SetTranDate(v string)`

SetTranDate sets TranDate field to given value.

### HasTranDate

`func (o *PaymentItem) HasTranDate() bool`

HasTranDate returns a boolean if a field has been set.

### GetNsLastModified

`func (o *PaymentItem) GetNsLastModified() time.Time`

GetNsLastModified returns the NsLastModified field if non-nil, zero value otherwise.

### GetNsLastModifiedOk

`func (o *PaymentItem) GetNsLastModifiedOk() (*time.Time, bool)`

GetNsLastModifiedOk returns a tuple with the NsLastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsLastModified

`func (o *PaymentItem) SetNsLastModified(v time.Time)`

SetNsLastModified sets NsLastModified field to given value.

### HasNsLastModified

`func (o *PaymentItem) HasNsLastModified() bool`

HasNsLastModified returns a boolean if a field has been set.

### GetSyncedAt

`func (o *PaymentItem) GetSyncedAt() time.Time`

GetSyncedAt returns the SyncedAt field if non-nil, zero value otherwise.

### GetSyncedAtOk

`func (o *PaymentItem) GetSyncedAtOk() (*time.Time, bool)`

GetSyncedAtOk returns a tuple with the SyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedAt

`func (o *PaymentItem) SetSyncedAt(v time.Time)`

SetSyncedAt sets SyncedAt field to given value.

### HasSyncedAt

`func (o *PaymentItem) HasSyncedAt() bool`

HasSyncedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


