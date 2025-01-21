# Diversion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | The reason for the diversion. Common values: unknown, user-busy, no-answer, unavailable, unconditional, time-of-day, do-not-disturb, deflection, follow-me, out-of-service, away. | [optional] 
**Privacy** | Pointer to **string** | off or full | [optional] 
**Screen** | Pointer to **string** | No if the number was provided by the user, yes if the number was provided by the network | [optional] 
**Counter** | Pointer to **string** | The number of diversions that have occurred | [optional] 
**Limit** | Pointer to **string** | The maximum number of diversions allowed for this session | [optional] 
**Unknown** | Pointer to **string** | The normal list of values is not exhaustive. Your application must be tolerant of unlisted keys and unlisted values of those keys. | [optional] 
**OrigTo** | Pointer to **string** | Always present. Indicates the last telephone number that the call was diverted from. | [optional] 

## Methods

### NewDiversion

`func NewDiversion() *Diversion`

NewDiversion instantiates a new Diversion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiversionWithDefaults

`func NewDiversionWithDefaults() *Diversion`

NewDiversionWithDefaults instantiates a new Diversion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *Diversion) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Diversion) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Diversion) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Diversion) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetPrivacy

`func (o *Diversion) GetPrivacy() string`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *Diversion) GetPrivacyOk() (*string, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *Diversion) SetPrivacy(v string)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *Diversion) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetScreen

`func (o *Diversion) GetScreen() string`

GetScreen returns the Screen field if non-nil, zero value otherwise.

### GetScreenOk

`func (o *Diversion) GetScreenOk() (*string, bool)`

GetScreenOk returns a tuple with the Screen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreen

`func (o *Diversion) SetScreen(v string)`

SetScreen sets Screen field to given value.

### HasScreen

`func (o *Diversion) HasScreen() bool`

HasScreen returns a boolean if a field has been set.

### GetCounter

`func (o *Diversion) GetCounter() string`

GetCounter returns the Counter field if non-nil, zero value otherwise.

### GetCounterOk

`func (o *Diversion) GetCounterOk() (*string, bool)`

GetCounterOk returns a tuple with the Counter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounter

`func (o *Diversion) SetCounter(v string)`

SetCounter sets Counter field to given value.

### HasCounter

`func (o *Diversion) HasCounter() bool`

HasCounter returns a boolean if a field has been set.

### GetLimit

`func (o *Diversion) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Diversion) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Diversion) SetLimit(v string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Diversion) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetUnknown

`func (o *Diversion) GetUnknown() string`

GetUnknown returns the Unknown field if non-nil, zero value otherwise.

### GetUnknownOk

`func (o *Diversion) GetUnknownOk() (*string, bool)`

GetUnknownOk returns a tuple with the Unknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknown

`func (o *Diversion) SetUnknown(v string)`

SetUnknown sets Unknown field to given value.

### HasUnknown

`func (o *Diversion) HasUnknown() bool`

HasUnknown returns a boolean if a field has been set.

### GetOrigTo

`func (o *Diversion) GetOrigTo() string`

GetOrigTo returns the OrigTo field if non-nil, zero value otherwise.

### GetOrigToOk

`func (o *Diversion) GetOrigToOk() (*string, bool)`

GetOrigToOk returns a tuple with the OrigTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigTo

`func (o *Diversion) SetOrigTo(v string)`

SetOrigTo sets OrigTo field to given value.

### HasOrigTo

`func (o *Diversion) HasOrigTo() bool`

HasOrigTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


