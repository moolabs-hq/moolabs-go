# DashboardOverviewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aging** | [**AgingReportResponse**](AgingReportResponse.md) |  | 
**Dso** | [**DSOResponse**](DSOResponse.md) |  | 
**Cei** | [**CEIResponse**](CEIResponse.md) |  | 
**DisputeSummary** | [**DisputeSummaryResponse**](DisputeSummaryResponse.md) |  | 
**PtpSummary** | [**PTPSummaryResponse**](PTPSummaryResponse.md) |  | 

## Methods

### NewDashboardOverviewResponse

`func NewDashboardOverviewResponse(aging AgingReportResponse, dso DSOResponse, cei CEIResponse, disputeSummary DisputeSummaryResponse, ptpSummary PTPSummaryResponse, ) *DashboardOverviewResponse`

NewDashboardOverviewResponse instantiates a new DashboardOverviewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardOverviewResponseWithDefaults

`func NewDashboardOverviewResponseWithDefaults() *DashboardOverviewResponse`

NewDashboardOverviewResponseWithDefaults instantiates a new DashboardOverviewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAging

`func (o *DashboardOverviewResponse) GetAging() AgingReportResponse`

GetAging returns the Aging field if non-nil, zero value otherwise.

### GetAgingOk

`func (o *DashboardOverviewResponse) GetAgingOk() (*AgingReportResponse, bool)`

GetAgingOk returns a tuple with the Aging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAging

`func (o *DashboardOverviewResponse) SetAging(v AgingReportResponse)`

SetAging sets Aging field to given value.


### GetDso

`func (o *DashboardOverviewResponse) GetDso() DSOResponse`

GetDso returns the Dso field if non-nil, zero value otherwise.

### GetDsoOk

`func (o *DashboardOverviewResponse) GetDsoOk() (*DSOResponse, bool)`

GetDsoOk returns a tuple with the Dso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDso

`func (o *DashboardOverviewResponse) SetDso(v DSOResponse)`

SetDso sets Dso field to given value.


### GetCei

`func (o *DashboardOverviewResponse) GetCei() CEIResponse`

GetCei returns the Cei field if non-nil, zero value otherwise.

### GetCeiOk

`func (o *DashboardOverviewResponse) GetCeiOk() (*CEIResponse, bool)`

GetCeiOk returns a tuple with the Cei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCei

`func (o *DashboardOverviewResponse) SetCei(v CEIResponse)`

SetCei sets Cei field to given value.


### GetDisputeSummary

`func (o *DashboardOverviewResponse) GetDisputeSummary() DisputeSummaryResponse`

GetDisputeSummary returns the DisputeSummary field if non-nil, zero value otherwise.

### GetDisputeSummaryOk

`func (o *DashboardOverviewResponse) GetDisputeSummaryOk() (*DisputeSummaryResponse, bool)`

GetDisputeSummaryOk returns a tuple with the DisputeSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeSummary

`func (o *DashboardOverviewResponse) SetDisputeSummary(v DisputeSummaryResponse)`

SetDisputeSummary sets DisputeSummary field to given value.


### GetPtpSummary

`func (o *DashboardOverviewResponse) GetPtpSummary() PTPSummaryResponse`

GetPtpSummary returns the PtpSummary field if non-nil, zero value otherwise.

### GetPtpSummaryOk

`func (o *DashboardOverviewResponse) GetPtpSummaryOk() (*PTPSummaryResponse, bool)`

GetPtpSummaryOk returns a tuple with the PtpSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtpSummary

`func (o *DashboardOverviewResponse) SetPtpSummary(v PTPSummaryResponse)`

SetPtpSummary sets PtpSummary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


