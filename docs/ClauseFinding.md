# ClauseFinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FamilyId** | **string** |  | 
**FoundSpan** | Pointer to **map[string]interface{}** |  | [optional] 
**Verdict** | **string** |  | 
**MatchedTier** | Pointer to **string** |  | [optional] 
**Delta** | Pointer to **string** |  | [optional] 
**Confidence** | **float32** |  | 
**RequiresContext** | Pointer to **bool** |  | [optional] [default to false]
**DealContextHit** | Pointer to **string** |  | [optional] 
**ClarifyingQuestion** | Pointer to **string** |  | [optional] 

## Methods

### NewClauseFinding

`func NewClauseFinding(familyId string, verdict string, confidence float32, ) *ClauseFinding`

NewClauseFinding instantiates a new ClauseFinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClauseFindingWithDefaults

`func NewClauseFindingWithDefaults() *ClauseFinding`

NewClauseFindingWithDefaults instantiates a new ClauseFinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFamilyId

`func (o *ClauseFinding) GetFamilyId() string`

GetFamilyId returns the FamilyId field if non-nil, zero value otherwise.

### GetFamilyIdOk

`func (o *ClauseFinding) GetFamilyIdOk() (*string, bool)`

GetFamilyIdOk returns a tuple with the FamilyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyId

`func (o *ClauseFinding) SetFamilyId(v string)`

SetFamilyId sets FamilyId field to given value.


### GetFoundSpan

`func (o *ClauseFinding) GetFoundSpan() map[string]interface{}`

GetFoundSpan returns the FoundSpan field if non-nil, zero value otherwise.

### GetFoundSpanOk

`func (o *ClauseFinding) GetFoundSpanOk() (*map[string]interface{}, bool)`

GetFoundSpanOk returns a tuple with the FoundSpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundSpan

`func (o *ClauseFinding) SetFoundSpan(v map[string]interface{})`

SetFoundSpan sets FoundSpan field to given value.

### HasFoundSpan

`func (o *ClauseFinding) HasFoundSpan() bool`

HasFoundSpan returns a boolean if a field has been set.

### GetVerdict

`func (o *ClauseFinding) GetVerdict() string`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *ClauseFinding) GetVerdictOk() (*string, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *ClauseFinding) SetVerdict(v string)`

SetVerdict sets Verdict field to given value.


### GetMatchedTier

`func (o *ClauseFinding) GetMatchedTier() string`

GetMatchedTier returns the MatchedTier field if non-nil, zero value otherwise.

### GetMatchedTierOk

`func (o *ClauseFinding) GetMatchedTierOk() (*string, bool)`

GetMatchedTierOk returns a tuple with the MatchedTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedTier

`func (o *ClauseFinding) SetMatchedTier(v string)`

SetMatchedTier sets MatchedTier field to given value.

### HasMatchedTier

`func (o *ClauseFinding) HasMatchedTier() bool`

HasMatchedTier returns a boolean if a field has been set.

### GetDelta

`func (o *ClauseFinding) GetDelta() string`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *ClauseFinding) GetDeltaOk() (*string, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *ClauseFinding) SetDelta(v string)`

SetDelta sets Delta field to given value.

### HasDelta

`func (o *ClauseFinding) HasDelta() bool`

HasDelta returns a boolean if a field has been set.

### GetConfidence

`func (o *ClauseFinding) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *ClauseFinding) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *ClauseFinding) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.


### GetRequiresContext

`func (o *ClauseFinding) GetRequiresContext() bool`

GetRequiresContext returns the RequiresContext field if non-nil, zero value otherwise.

### GetRequiresContextOk

`func (o *ClauseFinding) GetRequiresContextOk() (*bool, bool)`

GetRequiresContextOk returns a tuple with the RequiresContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresContext

`func (o *ClauseFinding) SetRequiresContext(v bool)`

SetRequiresContext sets RequiresContext field to given value.

### HasRequiresContext

`func (o *ClauseFinding) HasRequiresContext() bool`

HasRequiresContext returns a boolean if a field has been set.

### GetDealContextHit

`func (o *ClauseFinding) GetDealContextHit() string`

GetDealContextHit returns the DealContextHit field if non-nil, zero value otherwise.

### GetDealContextHitOk

`func (o *ClauseFinding) GetDealContextHitOk() (*string, bool)`

GetDealContextHitOk returns a tuple with the DealContextHit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealContextHit

`func (o *ClauseFinding) SetDealContextHit(v string)`

SetDealContextHit sets DealContextHit field to given value.

### HasDealContextHit

`func (o *ClauseFinding) HasDealContextHit() bool`

HasDealContextHit returns a boolean if a field has been set.

### GetClarifyingQuestion

`func (o *ClauseFinding) GetClarifyingQuestion() string`

GetClarifyingQuestion returns the ClarifyingQuestion field if non-nil, zero value otherwise.

### GetClarifyingQuestionOk

`func (o *ClauseFinding) GetClarifyingQuestionOk() (*string, bool)`

GetClarifyingQuestionOk returns a tuple with the ClarifyingQuestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClarifyingQuestion

`func (o *ClauseFinding) SetClarifyingQuestion(v string)`

SetClarifyingQuestion sets ClarifyingQuestion field to given value.

### HasClarifyingQuestion

`func (o *ClauseFinding) HasClarifyingQuestion() bool`

HasClarifyingQuestion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


