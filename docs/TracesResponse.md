# TracesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ledger** | **int32** |  | 
**Outbox** | **int32** |  | 
**Dlq** | **int32** |  | 
**Unpriced** | **int32** |  | 
**Warnings** | **int32** |  | 

## Methods

### NewTracesResponse

`func NewTracesResponse(ledger int32, outbox int32, dlq int32, unpriced int32, warnings int32, ) *TracesResponse`

NewTracesResponse instantiates a new TracesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTracesResponseWithDefaults

`func NewTracesResponseWithDefaults() *TracesResponse`

NewTracesResponseWithDefaults instantiates a new TracesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLedger

`func (o *TracesResponse) GetLedger() int32`

GetLedger returns the Ledger field if non-nil, zero value otherwise.

### GetLedgerOk

`func (o *TracesResponse) GetLedgerOk() (*int32, bool)`

GetLedgerOk returns a tuple with the Ledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedger

`func (o *TracesResponse) SetLedger(v int32)`

SetLedger sets Ledger field to given value.


### GetOutbox

`func (o *TracesResponse) GetOutbox() int32`

GetOutbox returns the Outbox field if non-nil, zero value otherwise.

### GetOutboxOk

`func (o *TracesResponse) GetOutboxOk() (*int32, bool)`

GetOutboxOk returns a tuple with the Outbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbox

`func (o *TracesResponse) SetOutbox(v int32)`

SetOutbox sets Outbox field to given value.


### GetDlq

`func (o *TracesResponse) GetDlq() int32`

GetDlq returns the Dlq field if non-nil, zero value otherwise.

### GetDlqOk

`func (o *TracesResponse) GetDlqOk() (*int32, bool)`

GetDlqOk returns a tuple with the Dlq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlq

`func (o *TracesResponse) SetDlq(v int32)`

SetDlq sets Dlq field to given value.


### GetUnpriced

`func (o *TracesResponse) GetUnpriced() int32`

GetUnpriced returns the Unpriced field if non-nil, zero value otherwise.

### GetUnpricedOk

`func (o *TracesResponse) GetUnpricedOk() (*int32, bool)`

GetUnpricedOk returns a tuple with the Unpriced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnpriced

`func (o *TracesResponse) SetUnpriced(v int32)`

SetUnpriced sets Unpriced field to given value.


### GetWarnings

`func (o *TracesResponse) GetWarnings() int32`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *TracesResponse) GetWarningsOk() (*int32, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *TracesResponse) SetWarnings(v int32)`

SetWarnings sets Warnings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


