# SummaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Projects** | **int64** | Number of Python projects in index | 
**Releases** | **int64** | Number of Python distribution releases in index | 
**Files** | **int64** | Number of files for all distributions in index | 

## Methods

### NewSummaryResponse

`func NewSummaryResponse(projects int64, releases int64, files int64, ) *SummaryResponse`

NewSummaryResponse instantiates a new SummaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryResponseWithDefaults

`func NewSummaryResponseWithDefaults() *SummaryResponse`

NewSummaryResponseWithDefaults instantiates a new SummaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjects

`func (o *SummaryResponse) GetProjects() int64`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *SummaryResponse) GetProjectsOk() (*int64, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *SummaryResponse) SetProjects(v int64)`

SetProjects sets Projects field to given value.


### GetReleases

`func (o *SummaryResponse) GetReleases() int64`

GetReleases returns the Releases field if non-nil, zero value otherwise.

### GetReleasesOk

`func (o *SummaryResponse) GetReleasesOk() (*int64, bool)`

GetReleasesOk returns a tuple with the Releases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleases

`func (o *SummaryResponse) SetReleases(v int64)`

SetReleases sets Releases field to given value.


### GetFiles

`func (o *SummaryResponse) GetFiles() int64`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *SummaryResponse) GetFilesOk() (*int64, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *SummaryResponse) SetFiles(v int64)`

SetFiles sets Files field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


