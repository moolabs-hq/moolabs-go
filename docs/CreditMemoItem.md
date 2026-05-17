# CreditMemoItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**NetsuiteId** | **string** |  | 
**CustomerInternalId** | **string** |  | 
**ForeignTotal** | Pointer to **string** |  | [optional] 
**Memo** | Pointer to **string** |  | [optional] 
**TranDate** | Pointer to **string** |  | [optional] 
**NsLastModified** | Pointer to **time.Time** |  | [optional] 
**SyncedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCreditMemoItem

`func NewCreditMemoItem(id string, netsuiteId string, customerInternalId string, ) *CreditMemoItem`

NewCreditMemoItem instantiates a new CreditMemoItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditMemoItemWithDefaults

`func NewCreditMemoItemWithDefaults() *CreditMemoItem`

NewCreditMemoItemWithDefaults instantiates a new CreditMemoItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreditMemoItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreditMemoItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreditMemoItem) SetId(v string)`

SetId sets Id field to given value.


### GetNetsuiteId

`func (o *CreditMemoItem) GetNetsuiteId() string`

GetNetsuiteId returns the NetsuiteId field if non-nil, zero value otherwise.

### GetNetsuiteIdOk

`func (o *CreditMemoItem) GetNetsuiteIdOk() (*string, bool)`

GetNetsuiteIdOk returns a tuple with the NetsuiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetsuiteId

`func (o *CreditMemoItem) SetNetsuiteId(v string)`

SetNetsuiteId sets NetsuiteId field to given value.


### GetCustomerInternalId

`func (o *CreditMemoItem) GetCustomerInternalId() string`

GetCustomerInternalId returns the CustomerInternalId field if non-nil, zero value otherwise.

### GetCustomerInternalIdOk

`func (o *CreditMemoItem) GetCustomerInternalIdOk() (*string, bool)`

GetCustomerInternalIdOk returns a tuple with the CustomerInternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerInternalId

`func (o *CreditMemoItem) SetCustomerInternalId(v string)`

SetCustomerInternalId sets CustomerInternalId field to given value.


### GetForeignTotal

`func (o *CreditMemoItem) GetForeignTotal() string`

GetForeignTotal returns the ForeignTotal field if non-nil, zero value otherwise.

### GetForeignTotalOk

`func (o *CreditMemoItem) GetForeignTotalOk() (*string, bool)`

GetForeignTotalOk returns a tuple with the ForeignTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignTotal

`func (o *CreditMemoItem) SetForeignTotal(v string)`

SetForeignTotal sets ForeignTotal field to given value.

### HasForeignTotal

`func (o *CreditMemoItem) HasForeignTotal() bool`

HasForeignTotal returns a boolean if a field has been set.

### GetMemo

`func (o *CreditMemoItem) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *CreditMemoItem) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *CreditMemoItem) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *CreditMemoItem) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetTranDate

`func (o *CreditMemoItem) GetTranDate() string`

GetTranDate returns the TranDate field if non-nil, zero value otherwise.

### GetTranDateOk

`func (o *CreditMemoItem) GetTranDateOk() (*string, bool)`

GetTranDateOk returns a tuple with the TranDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranDate

`func (o *CreditMemoItem) SetTranDate(v string)`

SetTranDate sets TranDate field to given value.

### HasTranDate

`func (o *CreditMemoItem) HasTranDate() bool`

HasTranDate returns a boolean if a field has been set.

### GetNsLastModified

`func (o *CreditMemoItem) GetNsLastModified() time.Time`

GetNsLastModified returns the NsLastModified field if non-nil, zero value otherwise.

### GetNsLastModifiedOk

`func (o *CreditMemoItem) GetNsLastModifiedOk() (*time.Time, bool)`

GetNsLastModifiedOk returns a tuple with the NsLastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsLastModified

`func (o *CreditMemoItem) SetNsLastModified(v time.Time)`

SetNsLastModified sets NsLastModified field to given value.

### HasNsLastModified

`func (o *CreditMemoItem) HasNsLastModified() bool`

HasNsLastModified returns a boolean if a field has been set.

### GetSyncedAt

`func (o *CreditMemoItem) GetSyncedAt() time.Time`

GetSyncedAt returns the SyncedAt field if non-nil, zero value otherwise.

### GetSyncedAtOk

`func (o *CreditMemoItem) GetSyncedAtOk() (*time.Time, bool)`

GetSyncedAtOk returns a tuple with the SyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedAt

`func (o *CreditMemoItem) SetSyncedAt(v time.Time)`

SetSyncedAt sets SyncedAt field to given value.

### HasSyncedAt

`func (o *CreditMemoItem) HasSyncedAt() bool`

HasSyncedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


