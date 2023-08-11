# AnsibleAnsibleNamespaceMetadataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** | Required named, only accepts lowercase, numbers and underscores. | 
**Company** | Pointer to **string** | Optional namespace company owner. | [optional] 
**Email** | Pointer to **string** | Optional namespace contact email. | [optional] 
**Description** | Pointer to **string** | Optional short description. | [optional] 
**Resources** | Pointer to **string** | Optional resource page in markdown format. | [optional] 
**Links** | Pointer to [**[]NamespaceLinkResponse**](NamespaceLinkResponse.md) | Labeled related links. | [optional] 
**AvatarSha256** | Pointer to **string** | SHA256 digest of avatar image if present. | [optional] [readonly] 
**AvatarUrl** | Pointer to **string** | Download link for avatar image if present. | [optional] [readonly] 
**MetadataSha256** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewAnsibleAnsibleNamespaceMetadataResponse

`func NewAnsibleAnsibleNamespaceMetadataResponse(name string, ) *AnsibleAnsibleNamespaceMetadataResponse`

NewAnsibleAnsibleNamespaceMetadataResponse instantiates a new AnsibleAnsibleNamespaceMetadataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnsibleAnsibleNamespaceMetadataResponseWithDefaults

`func NewAnsibleAnsibleNamespaceMetadataResponseWithDefaults() *AnsibleAnsibleNamespaceMetadataResponse`

NewAnsibleAnsibleNamespaceMetadataResponseWithDefaults instantiates a new AnsibleAnsibleNamespaceMetadataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *AnsibleAnsibleNamespaceMetadataResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *AnsibleAnsibleNamespaceMetadataResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *AnsibleAnsibleNamespaceMetadataResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *AnsibleAnsibleNamespaceMetadataResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetName

`func (o *AnsibleAnsibleNamespaceMetadataResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AnsibleAnsibleNamespaceMetadataResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AnsibleAnsibleNamespaceMetadataResponse) SetName(v string)`

SetName sets Name field to given value.


### GetCompany

`func (o *AnsibleAnsibleNamespaceMetadataResponse) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *AnsibleAnsibleNamespaceMetadataResponse) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *AnsibleAnsibleNamespaceMetadataResponse) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *AnsibleAnsibleNamespaceMetadataResponse) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetEmail

`func (o *AnsibleAnsibleNamespaceMetadataResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AnsibleAnsibleNamespaceMetadataResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AnsibleAnsibleNamespaceMetadataResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AnsibleAnsibleNamespaceMetadataResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDescription

`func (o *AnsibleAnsibleNamespaceMetadataResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AnsibleAnsibleNamespaceMetadataResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AnsibleAnsibleNamespaceMetadataResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AnsibleAnsibleNamespaceMetadataResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResources

`func (o *AnsibleAnsibleNamespaceMetadataResponse) GetResources() string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AnsibleAnsibleNamespaceMetadataResponse) GetResourcesOk() (*string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AnsibleAnsibleNamespaceMetadataResponse) SetResources(v string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *AnsibleAnsibleNamespaceMetadataResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetLinks

`func (o *AnsibleAnsibleNamespaceMetadataResponse) GetLinks() []NamespaceLinkResponse`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AnsibleAnsibleNamespaceMetadataResponse) GetLinksOk() (*[]NamespaceLinkResponse, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AnsibleAnsibleNamespaceMetadataResponse) SetLinks(v []NamespaceLinkResponse)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AnsibleAnsibleNamespaceMetadataResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAvatarSha256

`func (o *AnsibleAnsibleNamespaceMetadataResponse) GetAvatarSha256() string`

GetAvatarSha256 returns the AvatarSha256 field if non-nil, zero value otherwise.

### GetAvatarSha256Ok

`func (o *AnsibleAnsibleNamespaceMetadataResponse) GetAvatarSha256Ok() (*string, bool)`

GetAvatarSha256Ok returns a tuple with the AvatarSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarSha256

`func (o *AnsibleAnsibleNamespaceMetadataResponse) SetAvatarSha256(v string)`

SetAvatarSha256 sets AvatarSha256 field to given value.

### HasAvatarSha256

`func (o *AnsibleAnsibleNamespaceMetadataResponse) HasAvatarSha256() bool`

HasAvatarSha256 returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *AnsibleAnsibleNamespaceMetadataResponse) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *AnsibleAnsibleNamespaceMetadataResponse) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *AnsibleAnsibleNamespaceMetadataResponse) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *AnsibleAnsibleNamespaceMetadataResponse) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetMetadataSha256

`func (o *AnsibleAnsibleNamespaceMetadataResponse) GetMetadataSha256() string`

GetMetadataSha256 returns the MetadataSha256 field if non-nil, zero value otherwise.

### GetMetadataSha256Ok

`func (o *AnsibleAnsibleNamespaceMetadataResponse) GetMetadataSha256Ok() (*string, bool)`

GetMetadataSha256Ok returns a tuple with the MetadataSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataSha256

`func (o *AnsibleAnsibleNamespaceMetadataResponse) SetMetadataSha256(v string)`

SetMetadataSha256 sets MetadataSha256 field to given value.

### HasMetadataSha256

`func (o *AnsibleAnsibleNamespaceMetadataResponse) HasMetadataSha256() bool`

HasMetadataSha256 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


