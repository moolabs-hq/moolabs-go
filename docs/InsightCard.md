# InsightCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | [**CardType**](CardType.md) |  | 
**Title** | **string** |  | 
**Description** | **string** |  | 
**CtaLabel** | **string** |  | 
**CtaRoute** | **string** |  | 
**Priority** | [**Priority**](Priority.md) |  | 
**Icon** | **string** |  | 
**IconColor** | **string** |  | 
**IconBg** | **string** |  | 
**Persona** | **[]string** |  | 
**MetricValue** | Pointer to **string** |  | [optional] 

## Methods

### NewInsightCard

`func NewInsightCard(id string, type_ CardType, title string, description string, ctaLabel string, ctaRoute string, priority Priority, icon string, iconColor string, iconBg string, persona []string, ) *InsightCard`

NewInsightCard instantiates a new InsightCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightCardWithDefaults

`func NewInsightCardWithDefaults() *InsightCard`

NewInsightCardWithDefaults instantiates a new InsightCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InsightCard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InsightCard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InsightCard) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *InsightCard) GetType() CardType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InsightCard) GetTypeOk() (*CardType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InsightCard) SetType(v CardType)`

SetType sets Type field to given value.


### GetTitle

`func (o *InsightCard) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InsightCard) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InsightCard) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *InsightCard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InsightCard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InsightCard) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCtaLabel

`func (o *InsightCard) GetCtaLabel() string`

GetCtaLabel returns the CtaLabel field if non-nil, zero value otherwise.

### GetCtaLabelOk

`func (o *InsightCard) GetCtaLabelOk() (*string, bool)`

GetCtaLabelOk returns a tuple with the CtaLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtaLabel

`func (o *InsightCard) SetCtaLabel(v string)`

SetCtaLabel sets CtaLabel field to given value.


### GetCtaRoute

`func (o *InsightCard) GetCtaRoute() string`

GetCtaRoute returns the CtaRoute field if non-nil, zero value otherwise.

### GetCtaRouteOk

`func (o *InsightCard) GetCtaRouteOk() (*string, bool)`

GetCtaRouteOk returns a tuple with the CtaRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtaRoute

`func (o *InsightCard) SetCtaRoute(v string)`

SetCtaRoute sets CtaRoute field to given value.


### GetPriority

`func (o *InsightCard) GetPriority() Priority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *InsightCard) GetPriorityOk() (*Priority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *InsightCard) SetPriority(v Priority)`

SetPriority sets Priority field to given value.


### GetIcon

`func (o *InsightCard) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *InsightCard) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *InsightCard) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetIconColor

`func (o *InsightCard) GetIconColor() string`

GetIconColor returns the IconColor field if non-nil, zero value otherwise.

### GetIconColorOk

`func (o *InsightCard) GetIconColorOk() (*string, bool)`

GetIconColorOk returns a tuple with the IconColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconColor

`func (o *InsightCard) SetIconColor(v string)`

SetIconColor sets IconColor field to given value.


### GetIconBg

`func (o *InsightCard) GetIconBg() string`

GetIconBg returns the IconBg field if non-nil, zero value otherwise.

### GetIconBgOk

`func (o *InsightCard) GetIconBgOk() (*string, bool)`

GetIconBgOk returns a tuple with the IconBg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconBg

`func (o *InsightCard) SetIconBg(v string)`

SetIconBg sets IconBg field to given value.


### GetPersona

`func (o *InsightCard) GetPersona() []string`

GetPersona returns the Persona field if non-nil, zero value otherwise.

### GetPersonaOk

`func (o *InsightCard) GetPersonaOk() (*[]string, bool)`

GetPersonaOk returns a tuple with the Persona field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersona

`func (o *InsightCard) SetPersona(v []string)`

SetPersona sets Persona field to given value.


### GetMetricValue

`func (o *InsightCard) GetMetricValue() string`

GetMetricValue returns the MetricValue field if non-nil, zero value otherwise.

### GetMetricValueOk

`func (o *InsightCard) GetMetricValueOk() (*string, bool)`

GetMetricValueOk returns a tuple with the MetricValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricValue

`func (o *InsightCard) SetMetricValue(v string)`

SetMetricValue sets MetricValue field to given value.

### HasMetricValue

`func (o *InsightCard) HasMetricValue() bool`

HasMetricValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


