# RpmUpdateRecordResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**PulpCreated** | Pointer to **time.Time** | Timestamp of creation. | [optional] [readonly] 
**Id** | Pointer to **string** | Update id (short update name, e.g. RHEA-2013:1777) | [optional] [readonly] 
**UpdatedDate** | Pointer to **string** | Date when the update was updated (e.g. &#39;2013-12-02 00:00:00&#39;) | [optional] [readonly] 
**Description** | Pointer to **string** | Update description | [optional] [readonly] 
**IssuedDate** | Pointer to **string** | Date when the update was issued (e.g. &#39;2013-12-02 00:00:00&#39;) | [optional] [readonly] 
**Fromstr** | Pointer to **string** | Source of the update (e.g. security@redhat.com) | [optional] [readonly] 
**Status** | Pointer to **string** | Update status (&#39;final&#39;, ...) | [optional] [readonly] 
**Title** | Pointer to **string** | Update name | [optional] [readonly] 
**Summary** | Pointer to **string** | Short summary | [optional] [readonly] 
**Version** | Pointer to **string** | Update version (probably always an integer number) | [optional] [readonly] 
**Type** | Pointer to **string** | Update type (&#39;enhancement&#39;, &#39;bugfix&#39;, ...) | [optional] [readonly] 
**Severity** | Pointer to **string** | Severity | [optional] [readonly] 
**Solution** | Pointer to **string** | Solution | [optional] [readonly] 
**Release** | Pointer to **string** | Update release | [optional] [readonly] 
**Rights** | Pointer to **string** | Copyrights | [optional] [readonly] 
**Pushcount** | Pointer to **string** | Push count | [optional] [readonly] 
**Pkglist** | Pointer to [**[]RpmUpdateCollectionResponse**](RpmUpdateCollectionResponse.md) | List of packages | [optional] [readonly] 
**References** | Pointer to **[]map[string]interface{}** | List of references | [optional] [readonly] 
**RebootSuggested** | Pointer to **bool** | Reboot suggested | [optional] [readonly] 

## Methods

### NewRpmUpdateRecordResponse

`func NewRpmUpdateRecordResponse() *RpmUpdateRecordResponse`

NewRpmUpdateRecordResponse instantiates a new RpmUpdateRecordResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpmUpdateRecordResponseWithDefaults

`func NewRpmUpdateRecordResponseWithDefaults() *RpmUpdateRecordResponse`

NewRpmUpdateRecordResponseWithDefaults instantiates a new RpmUpdateRecordResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *RpmUpdateRecordResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *RpmUpdateRecordResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *RpmUpdateRecordResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *RpmUpdateRecordResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetPulpCreated

`func (o *RpmUpdateRecordResponse) GetPulpCreated() time.Time`

GetPulpCreated returns the PulpCreated field if non-nil, zero value otherwise.

### GetPulpCreatedOk

`func (o *RpmUpdateRecordResponse) GetPulpCreatedOk() (*time.Time, bool)`

GetPulpCreatedOk returns a tuple with the PulpCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpCreated

`func (o *RpmUpdateRecordResponse) SetPulpCreated(v time.Time)`

SetPulpCreated sets PulpCreated field to given value.

### HasPulpCreated

`func (o *RpmUpdateRecordResponse) HasPulpCreated() bool`

HasPulpCreated returns a boolean if a field has been set.

### GetId

`func (o *RpmUpdateRecordResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RpmUpdateRecordResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RpmUpdateRecordResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RpmUpdateRecordResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUpdatedDate

`func (o *RpmUpdateRecordResponse) GetUpdatedDate() string`

GetUpdatedDate returns the UpdatedDate field if non-nil, zero value otherwise.

### GetUpdatedDateOk

`func (o *RpmUpdateRecordResponse) GetUpdatedDateOk() (*string, bool)`

GetUpdatedDateOk returns a tuple with the UpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDate

`func (o *RpmUpdateRecordResponse) SetUpdatedDate(v string)`

SetUpdatedDate sets UpdatedDate field to given value.

### HasUpdatedDate

`func (o *RpmUpdateRecordResponse) HasUpdatedDate() bool`

HasUpdatedDate returns a boolean if a field has been set.

### GetDescription

`func (o *RpmUpdateRecordResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RpmUpdateRecordResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RpmUpdateRecordResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RpmUpdateRecordResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIssuedDate

`func (o *RpmUpdateRecordResponse) GetIssuedDate() string`

GetIssuedDate returns the IssuedDate field if non-nil, zero value otherwise.

### GetIssuedDateOk

`func (o *RpmUpdateRecordResponse) GetIssuedDateOk() (*string, bool)`

GetIssuedDateOk returns a tuple with the IssuedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedDate

`func (o *RpmUpdateRecordResponse) SetIssuedDate(v string)`

SetIssuedDate sets IssuedDate field to given value.

### HasIssuedDate

`func (o *RpmUpdateRecordResponse) HasIssuedDate() bool`

HasIssuedDate returns a boolean if a field has been set.

### GetFromstr

`func (o *RpmUpdateRecordResponse) GetFromstr() string`

GetFromstr returns the Fromstr field if non-nil, zero value otherwise.

### GetFromstrOk

`func (o *RpmUpdateRecordResponse) GetFromstrOk() (*string, bool)`

GetFromstrOk returns a tuple with the Fromstr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromstr

`func (o *RpmUpdateRecordResponse) SetFromstr(v string)`

SetFromstr sets Fromstr field to given value.

### HasFromstr

`func (o *RpmUpdateRecordResponse) HasFromstr() bool`

HasFromstr returns a boolean if a field has been set.

### GetStatus

`func (o *RpmUpdateRecordResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RpmUpdateRecordResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RpmUpdateRecordResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RpmUpdateRecordResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *RpmUpdateRecordResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RpmUpdateRecordResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RpmUpdateRecordResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RpmUpdateRecordResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSummary

`func (o *RpmUpdateRecordResponse) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *RpmUpdateRecordResponse) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *RpmUpdateRecordResponse) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *RpmUpdateRecordResponse) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetVersion

`func (o *RpmUpdateRecordResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RpmUpdateRecordResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RpmUpdateRecordResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RpmUpdateRecordResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetType

`func (o *RpmUpdateRecordResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RpmUpdateRecordResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RpmUpdateRecordResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RpmUpdateRecordResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSeverity

`func (o *RpmUpdateRecordResponse) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *RpmUpdateRecordResponse) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *RpmUpdateRecordResponse) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *RpmUpdateRecordResponse) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSolution

`func (o *RpmUpdateRecordResponse) GetSolution() string`

GetSolution returns the Solution field if non-nil, zero value otherwise.

### GetSolutionOk

`func (o *RpmUpdateRecordResponse) GetSolutionOk() (*string, bool)`

GetSolutionOk returns a tuple with the Solution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolution

`func (o *RpmUpdateRecordResponse) SetSolution(v string)`

SetSolution sets Solution field to given value.

### HasSolution

`func (o *RpmUpdateRecordResponse) HasSolution() bool`

HasSolution returns a boolean if a field has been set.

### GetRelease

`func (o *RpmUpdateRecordResponse) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *RpmUpdateRecordResponse) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *RpmUpdateRecordResponse) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *RpmUpdateRecordResponse) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetRights

`func (o *RpmUpdateRecordResponse) GetRights() string`

GetRights returns the Rights field if non-nil, zero value otherwise.

### GetRightsOk

`func (o *RpmUpdateRecordResponse) GetRightsOk() (*string, bool)`

GetRightsOk returns a tuple with the Rights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRights

`func (o *RpmUpdateRecordResponse) SetRights(v string)`

SetRights sets Rights field to given value.

### HasRights

`func (o *RpmUpdateRecordResponse) HasRights() bool`

HasRights returns a boolean if a field has been set.

### GetPushcount

`func (o *RpmUpdateRecordResponse) GetPushcount() string`

GetPushcount returns the Pushcount field if non-nil, zero value otherwise.

### GetPushcountOk

`func (o *RpmUpdateRecordResponse) GetPushcountOk() (*string, bool)`

GetPushcountOk returns a tuple with the Pushcount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushcount

`func (o *RpmUpdateRecordResponse) SetPushcount(v string)`

SetPushcount sets Pushcount field to given value.

### HasPushcount

`func (o *RpmUpdateRecordResponse) HasPushcount() bool`

HasPushcount returns a boolean if a field has been set.

### GetPkglist

`func (o *RpmUpdateRecordResponse) GetPkglist() []RpmUpdateCollectionResponse`

GetPkglist returns the Pkglist field if non-nil, zero value otherwise.

### GetPkglistOk

`func (o *RpmUpdateRecordResponse) GetPkglistOk() (*[]RpmUpdateCollectionResponse, bool)`

GetPkglistOk returns a tuple with the Pkglist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkglist

`func (o *RpmUpdateRecordResponse) SetPkglist(v []RpmUpdateCollectionResponse)`

SetPkglist sets Pkglist field to given value.

### HasPkglist

`func (o *RpmUpdateRecordResponse) HasPkglist() bool`

HasPkglist returns a boolean if a field has been set.

### GetReferences

`func (o *RpmUpdateRecordResponse) GetReferences() []map[string]interface{}`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *RpmUpdateRecordResponse) GetReferencesOk() (*[]map[string]interface{}, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *RpmUpdateRecordResponse) SetReferences(v []map[string]interface{})`

SetReferences sets References field to given value.

### HasReferences

`func (o *RpmUpdateRecordResponse) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetRebootSuggested

`func (o *RpmUpdateRecordResponse) GetRebootSuggested() bool`

GetRebootSuggested returns the RebootSuggested field if non-nil, zero value otherwise.

### GetRebootSuggestedOk

`func (o *RpmUpdateRecordResponse) GetRebootSuggestedOk() (*bool, bool)`

GetRebootSuggestedOk returns a tuple with the RebootSuggested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootSuggested

`func (o *RpmUpdateRecordResponse) SetRebootSuggested(v bool)`

SetRebootSuggested sets RebootSuggested field to given value.

### HasRebootSuggested

`func (o *RpmUpdateRecordResponse) HasRebootSuggested() bool`

HasRebootSuggested returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


