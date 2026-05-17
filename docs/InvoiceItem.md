# InvoiceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**NetsuiteId** | **string** |  | 
**TranId** | Pointer to **string** |  | [optional] 
**TranDate** | Pointer to **string** |  | [optional] 
**DueDate** | Pointer to **string** |  | [optional] 
**CustomerInternalId** | **string** |  | 
**ForeignTotal** | Pointer to **string** |  | [optional] 
**ForeignAmountUnpaid** | Pointer to **string** |  | [optional] 
**StatusLabel** | Pointer to **string** |  | [optional] 
**CurrencyCode** | Pointer to **string** |  | [optional] 
**Memo** | Pointer to **string** |  | [optional] 
**NsLastModified** | Pointer to **time.Time** |  | [optional] 
**SyncedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewInvoiceItem

`func NewInvoiceItem(id string, netsuiteId string, customerInternalId string, ) *InvoiceItem`

NewInvoiceItem instantiates a new InvoiceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceItemWithDefaults

`func NewInvoiceItemWithDefaults() *InvoiceItem`

NewInvoiceItemWithDefaults instantiates a new InvoiceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceItem) SetId(v string)`

SetId sets Id field to given value.


### GetNetsuiteId

`func (o *InvoiceItem) GetNetsuiteId() string`

GetNetsuiteId returns the NetsuiteId field if non-nil, zero value otherwise.

### GetNetsuiteIdOk

`func (o *InvoiceItem) GetNetsuiteIdOk() (*string, bool)`

GetNetsuiteIdOk returns a tuple with the NetsuiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetsuiteId

`func (o *InvoiceItem) SetNetsuiteId(v string)`

SetNetsuiteId sets NetsuiteId field to given value.


### GetTranId

`func (o *InvoiceItem) GetTranId() string`

GetTranId returns the TranId field if non-nil, zero value otherwise.

### GetTranIdOk

`func (o *InvoiceItem) GetTranIdOk() (*string, bool)`

GetTranIdOk returns a tuple with the TranId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranId

`func (o *InvoiceItem) SetTranId(v string)`

SetTranId sets TranId field to given value.

### HasTranId

`func (o *InvoiceItem) HasTranId() bool`

HasTranId returns a boolean if a field has been set.

### GetTranDate

`func (o *InvoiceItem) GetTranDate() string`

GetTranDate returns the TranDate field if non-nil, zero value otherwise.

### GetTranDateOk

`func (o *InvoiceItem) GetTranDateOk() (*string, bool)`

GetTranDateOk returns a tuple with the TranDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranDate

`func (o *InvoiceItem) SetTranDate(v string)`

SetTranDate sets TranDate field to given value.

### HasTranDate

`func (o *InvoiceItem) HasTranDate() bool`

HasTranDate returns a boolean if a field has been set.

### GetDueDate

`func (o *InvoiceItem) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *InvoiceItem) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *InvoiceItem) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *InvoiceItem) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetCustomerInternalId

`func (o *InvoiceItem) GetCustomerInternalId() string`

GetCustomerInternalId returns the CustomerInternalId field if non-nil, zero value otherwise.

### GetCustomerInternalIdOk

`func (o *InvoiceItem) GetCustomerInternalIdOk() (*string, bool)`

GetCustomerInternalIdOk returns a tuple with the CustomerInternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerInternalId

`func (o *InvoiceItem) SetCustomerInternalId(v string)`

SetCustomerInternalId sets CustomerInternalId field to given value.


### GetForeignTotal

`func (o *InvoiceItem) GetForeignTotal() string`

GetForeignTotal returns the ForeignTotal field if non-nil, zero value otherwise.

### GetForeignTotalOk

`func (o *InvoiceItem) GetForeignTotalOk() (*string, bool)`

GetForeignTotalOk returns a tuple with the ForeignTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignTotal

`func (o *InvoiceItem) SetForeignTotal(v string)`

SetForeignTotal sets ForeignTotal field to given value.

### HasForeignTotal

`func (o *InvoiceItem) HasForeignTotal() bool`

HasForeignTotal returns a boolean if a field has been set.

### GetForeignAmountUnpaid

`func (o *InvoiceItem) GetForeignAmountUnpaid() string`

GetForeignAmountUnpaid returns the ForeignAmountUnpaid field if non-nil, zero value otherwise.

### GetForeignAmountUnpaidOk

`func (o *InvoiceItem) GetForeignAmountUnpaidOk() (*string, bool)`

GetForeignAmountUnpaidOk returns a tuple with the ForeignAmountUnpaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignAmountUnpaid

`func (o *InvoiceItem) SetForeignAmountUnpaid(v string)`

SetForeignAmountUnpaid sets ForeignAmountUnpaid field to given value.

### HasForeignAmountUnpaid

`func (o *InvoiceItem) HasForeignAmountUnpaid() bool`

HasForeignAmountUnpaid returns a boolean if a field has been set.

### GetStatusLabel

`func (o *InvoiceItem) GetStatusLabel() string`

GetStatusLabel returns the StatusLabel field if non-nil, zero value otherwise.

### GetStatusLabelOk

`func (o *InvoiceItem) GetStatusLabelOk() (*string, bool)`

GetStatusLabelOk returns a tuple with the StatusLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusLabel

`func (o *InvoiceItem) SetStatusLabel(v string)`

SetStatusLabel sets StatusLabel field to given value.

### HasStatusLabel

`func (o *InvoiceItem) HasStatusLabel() bool`

HasStatusLabel returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *InvoiceItem) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *InvoiceItem) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *InvoiceItem) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *InvoiceItem) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetMemo

`func (o *InvoiceItem) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *InvoiceItem) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *InvoiceItem) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *InvoiceItem) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetNsLastModified

`func (o *InvoiceItem) GetNsLastModified() time.Time`

GetNsLastModified returns the NsLastModified field if non-nil, zero value otherwise.

### GetNsLastModifiedOk

`func (o *InvoiceItem) GetNsLastModifiedOk() (*time.Time, bool)`

GetNsLastModifiedOk returns a tuple with the NsLastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsLastModified

`func (o *InvoiceItem) SetNsLastModified(v time.Time)`

SetNsLastModified sets NsLastModified field to given value.

### HasNsLastModified

`func (o *InvoiceItem) HasNsLastModified() bool`

HasNsLastModified returns a boolean if a field has been set.

### GetSyncedAt

`func (o *InvoiceItem) GetSyncedAt() time.Time`

GetSyncedAt returns the SyncedAt field if non-nil, zero value otherwise.

### GetSyncedAtOk

`func (o *InvoiceItem) GetSyncedAtOk() (*time.Time, bool)`

GetSyncedAtOk returns a tuple with the SyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedAt

`func (o *InvoiceItem) SetSyncedAt(v time.Time)`

SetSyncedAt sets SyncedAt field to given value.

### HasSyncedAt

`func (o *InvoiceItem) HasSyncedAt() bool`

HasSyncedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


