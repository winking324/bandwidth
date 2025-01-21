# CallState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **string** | The application id associated with the call. | [optional] 
**AccountId** | Pointer to **string** | The account id associated with the call. | [optional] 
**CallId** | Pointer to **string** | The programmable voice API call ID. | [optional] 
**ParentCallId** | Pointer to **NullableString** | The A-leg call id, set only if this call is the B-leg of a [&#x60;&lt;Transfer&gt;&#x60;](/docs/voice/bxml/transfer). | [optional] 
**To** | Pointer to **string** | The phone number that received the call, in E.164 format (e.g. +15555555555), or if the call was to a SIP URI, the SIP URI. | [optional] 
**From** | Pointer to **string** | The phone number that made the call, in E.164 format (e.g. +15555555555). | [optional] 
**Direction** | Pointer to [**CallDirectionEnum**](CallDirectionEnum.md) |  | [optional] 
**State** | Pointer to **string** | The current state of the call. Current possible values are &#x60;queued&#x60;, &#x60;initiated&#x60;, &#x60;answered&#x60; and &#x60;disconnected&#x60;. Additional states may be added in the future, so your application must be tolerant of unknown values. | [optional] 
**StirShaken** | Pointer to **map[string]string** | For inbound calls, the Bandwidth STIR/SHAKEN implementation will verify the information provided in the inbound invite request &#x60;Identity&#x60; header. The verification status is stored in the call state &#x60;stirShaken&#x60; property as follows.  | Property          | Description | |:------------------|:------------| | verstat | (optional) The verification status indicating whether the verification was successful or not. Possible values are &#x60;TN-Verification-Passed&#x60; or &#x60;TN-Verification-Failed&#x60;. | | attestationIndicator | (optional) The attestation level verified by Bandwidth. Possible values are &#x60;A&#x60; (full), &#x60;B&#x60; (partial) or &#x60;C&#x60; (gateway). | | originatingId | (optional) A unique origination identifier. |  Note that these are common properties but that the &#x60;stirShaken&#x60; object is free form and can contain other key-value pairs.  More information: [Understanding STIR/SHAKEN](https://www.bandwidth.com/regulations/stir-shaken). | [optional] 
**Identity** | Pointer to **NullableString** | The value of the &#x60;Identity&#x60; header from the inbound invite request. Only present for inbound calls and if the account is configured to forward this header. | [optional] 
**EnqueuedTime** | Pointer to **NullableTime** | The time this call was placed in queue. | [optional] 
**StartTime** | Pointer to **NullableTime** | The time the call was initiated, in ISO 8601 format. &#x60;null&#x60; if the call is still in your queue. | [optional] 
**AnswerTime** | Pointer to **NullableTime** | Populated once the call has been answered, with the time in ISO 8601 format. | [optional] 
**EndTime** | Pointer to **NullableTime** | Populated once the call has ended, with the time in ISO 8601 format. | [optional] 
**DisconnectCause** | Pointer to **NullableString** | | Cause | Description | |:------|:------------| | &#x60;hangup&#x60;| One party hung up the call, a [&#x60;&lt;Hangup&gt;&#x60;](../../bxml/verbs/hangup.md) verb was executed, or there was no more BXML to execute; it indicates that the call ended normally. | | &#x60;busy&#x60; | Callee was busy. | | &#x60;timeout&#x60; | Call wasn&#39;t answered before the &#x60;callTimeout&#x60; was reached. | | &#x60;cancel&#x60; | Call was cancelled by its originator while it was ringing. | | &#x60;rejected&#x60; | Call was rejected by the callee. | | &#x60;callback-error&#x60; | BXML callback couldn&#39;t be delivered to your callback server. | | &#x60;invalid-bxml&#x60; | Invalid BXML was returned in response to a callback. | | &#x60;application-error&#x60; | An unsupported action was tried on the call, e.g. trying to play a .ogg audio. | | &#x60;account-limit&#x60; | Account rate limits were reached. | | &#x60;node-capacity-exceeded&#x60; | System maximum capacity was reached. | | &#x60;error&#x60; | Some error not described in any of the other causes happened on the call. | | &#x60;unknown&#x60; | Unknown error happened on the call. |  Note: This list is not exhaustive and other values can appear in the future. | [optional] 
**ErrorMessage** | Pointer to **NullableString** | Populated only if the call ended with an error, with text explaining the reason. | [optional] 
**ErrorId** | Pointer to **NullableString** | Populated only if the call ended with an error, with a Bandwidth internal id that references the error event. | [optional] 
**LastUpdate** | Pointer to **time.Time** | The last time the call had a state update, in ISO 8601 format. | [optional] 

## Methods

### NewCallState

`func NewCallState() *CallState`

NewCallState instantiates a new CallState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallStateWithDefaults

`func NewCallStateWithDefaults() *CallState`

NewCallStateWithDefaults instantiates a new CallState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *CallState) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CallState) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CallState) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *CallState) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetAccountId

`func (o *CallState) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CallState) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CallState) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CallState) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCallId

`func (o *CallState) GetCallId() string`

GetCallId returns the CallId field if non-nil, zero value otherwise.

### GetCallIdOk

`func (o *CallState) GetCallIdOk() (*string, bool)`

GetCallIdOk returns a tuple with the CallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallId

`func (o *CallState) SetCallId(v string)`

SetCallId sets CallId field to given value.

### HasCallId

`func (o *CallState) HasCallId() bool`

HasCallId returns a boolean if a field has been set.

### GetParentCallId

`func (o *CallState) GetParentCallId() string`

GetParentCallId returns the ParentCallId field if non-nil, zero value otherwise.

### GetParentCallIdOk

`func (o *CallState) GetParentCallIdOk() (*string, bool)`

GetParentCallIdOk returns a tuple with the ParentCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCallId

`func (o *CallState) SetParentCallId(v string)`

SetParentCallId sets ParentCallId field to given value.

### HasParentCallId

`func (o *CallState) HasParentCallId() bool`

HasParentCallId returns a boolean if a field has been set.

### SetParentCallIdNil

`func (o *CallState) SetParentCallIdNil(b bool)`

 SetParentCallIdNil sets the value for ParentCallId to be an explicit nil

### UnsetParentCallId
`func (o *CallState) UnsetParentCallId()`

UnsetParentCallId ensures that no value is present for ParentCallId, not even an explicit nil
### GetTo

`func (o *CallState) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CallState) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CallState) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *CallState) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetFrom

`func (o *CallState) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CallState) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CallState) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *CallState) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetDirection

`func (o *CallState) GetDirection() CallDirectionEnum`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *CallState) GetDirectionOk() (*CallDirectionEnum, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *CallState) SetDirection(v CallDirectionEnum)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *CallState) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetState

`func (o *CallState) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CallState) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CallState) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CallState) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStirShaken

`func (o *CallState) GetStirShaken() map[string]string`

GetStirShaken returns the StirShaken field if non-nil, zero value otherwise.

### GetStirShakenOk

`func (o *CallState) GetStirShakenOk() (*map[string]string, bool)`

GetStirShakenOk returns a tuple with the StirShaken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStirShaken

`func (o *CallState) SetStirShaken(v map[string]string)`

SetStirShaken sets StirShaken field to given value.

### HasStirShaken

`func (o *CallState) HasStirShaken() bool`

HasStirShaken returns a boolean if a field has been set.

### SetStirShakenNil

`func (o *CallState) SetStirShakenNil(b bool)`

 SetStirShakenNil sets the value for StirShaken to be an explicit nil

### UnsetStirShaken
`func (o *CallState) UnsetStirShaken()`

UnsetStirShaken ensures that no value is present for StirShaken, not even an explicit nil
### GetIdentity

`func (o *CallState) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *CallState) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *CallState) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *CallState) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### SetIdentityNil

`func (o *CallState) SetIdentityNil(b bool)`

 SetIdentityNil sets the value for Identity to be an explicit nil

### UnsetIdentity
`func (o *CallState) UnsetIdentity()`

UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
### GetEnqueuedTime

`func (o *CallState) GetEnqueuedTime() time.Time`

GetEnqueuedTime returns the EnqueuedTime field if non-nil, zero value otherwise.

### GetEnqueuedTimeOk

`func (o *CallState) GetEnqueuedTimeOk() (*time.Time, bool)`

GetEnqueuedTimeOk returns a tuple with the EnqueuedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnqueuedTime

`func (o *CallState) SetEnqueuedTime(v time.Time)`

SetEnqueuedTime sets EnqueuedTime field to given value.

### HasEnqueuedTime

`func (o *CallState) HasEnqueuedTime() bool`

HasEnqueuedTime returns a boolean if a field has been set.

### SetEnqueuedTimeNil

`func (o *CallState) SetEnqueuedTimeNil(b bool)`

 SetEnqueuedTimeNil sets the value for EnqueuedTime to be an explicit nil

### UnsetEnqueuedTime
`func (o *CallState) UnsetEnqueuedTime()`

UnsetEnqueuedTime ensures that no value is present for EnqueuedTime, not even an explicit nil
### GetStartTime

`func (o *CallState) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CallState) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CallState) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *CallState) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *CallState) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *CallState) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetAnswerTime

`func (o *CallState) GetAnswerTime() time.Time`

GetAnswerTime returns the AnswerTime field if non-nil, zero value otherwise.

### GetAnswerTimeOk

`func (o *CallState) GetAnswerTimeOk() (*time.Time, bool)`

GetAnswerTimeOk returns a tuple with the AnswerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerTime

`func (o *CallState) SetAnswerTime(v time.Time)`

SetAnswerTime sets AnswerTime field to given value.

### HasAnswerTime

`func (o *CallState) HasAnswerTime() bool`

HasAnswerTime returns a boolean if a field has been set.

### SetAnswerTimeNil

`func (o *CallState) SetAnswerTimeNil(b bool)`

 SetAnswerTimeNil sets the value for AnswerTime to be an explicit nil

### UnsetAnswerTime
`func (o *CallState) UnsetAnswerTime()`

UnsetAnswerTime ensures that no value is present for AnswerTime, not even an explicit nil
### GetEndTime

`func (o *CallState) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CallState) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CallState) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *CallState) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *CallState) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *CallState) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetDisconnectCause

`func (o *CallState) GetDisconnectCause() string`

GetDisconnectCause returns the DisconnectCause field if non-nil, zero value otherwise.

### GetDisconnectCauseOk

`func (o *CallState) GetDisconnectCauseOk() (*string, bool)`

GetDisconnectCauseOk returns a tuple with the DisconnectCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectCause

`func (o *CallState) SetDisconnectCause(v string)`

SetDisconnectCause sets DisconnectCause field to given value.

### HasDisconnectCause

`func (o *CallState) HasDisconnectCause() bool`

HasDisconnectCause returns a boolean if a field has been set.

### SetDisconnectCauseNil

`func (o *CallState) SetDisconnectCauseNil(b bool)`

 SetDisconnectCauseNil sets the value for DisconnectCause to be an explicit nil

### UnsetDisconnectCause
`func (o *CallState) UnsetDisconnectCause()`

UnsetDisconnectCause ensures that no value is present for DisconnectCause, not even an explicit nil
### GetErrorMessage

`func (o *CallState) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *CallState) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *CallState) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *CallState) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *CallState) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *CallState) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetErrorId

`func (o *CallState) GetErrorId() string`

GetErrorId returns the ErrorId field if non-nil, zero value otherwise.

### GetErrorIdOk

`func (o *CallState) GetErrorIdOk() (*string, bool)`

GetErrorIdOk returns a tuple with the ErrorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorId

`func (o *CallState) SetErrorId(v string)`

SetErrorId sets ErrorId field to given value.

### HasErrorId

`func (o *CallState) HasErrorId() bool`

HasErrorId returns a boolean if a field has been set.

### SetErrorIdNil

`func (o *CallState) SetErrorIdNil(b bool)`

 SetErrorIdNil sets the value for ErrorId to be an explicit nil

### UnsetErrorId
`func (o *CallState) UnsetErrorId()`

UnsetErrorId ensures that no value is present for ErrorId, not even an explicit nil
### GetLastUpdate

`func (o *CallState) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *CallState) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *CallState) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *CallState) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


