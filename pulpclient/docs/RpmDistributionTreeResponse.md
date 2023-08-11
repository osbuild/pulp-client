# RpmDistributionTreeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PulpHref** | Pointer to **string** |  | [optional] [readonly] 
**HeaderVersion** | **string** | Header Version. | 
**ReleaseName** | **string** | Release name. | 
**ReleaseShort** | **string** | Release short name. | 
**ReleaseVersion** | **string** | Release version. | 
**ReleaseIsLayered** | **bool** | Typically False for an operating system, True otherwise. | 
**BaseProductName** | **NullableString** | Base Product name. | 
**BaseProductShort** | **NullableString** | Base Product short name. | 
**BaseProductVersion** | **NullableString** | Base Product version. | 
**Arch** | **string** | Tree architecturerch. | 
**BuildTimestamp** | **float64** | Tree build time timestamp. | 
**Instimage** | **NullableString** | Relative path to Anaconda instimage. | 
**Mainimage** | **NullableString** | Relative path to Anaconda stage2 image. | 
**Discnum** | **NullableInt64** | Disc number. | 
**Totaldiscs** | **NullableInt64** | Number of discs in media set. | 
**Addons** | [**[]AddonResponse**](AddonResponse.md) |  | 
**Checksums** | [**[]ChecksumResponse**](ChecksumResponse.md) |  | 
**Images** | [**[]ImageResponse**](ImageResponse.md) |  | 
**Variants** | [**[]VariantResponse**](VariantResponse.md) |  | 

## Methods

### NewRpmDistributionTreeResponse

`func NewRpmDistributionTreeResponse(headerVersion string, releaseName string, releaseShort string, releaseVersion string, releaseIsLayered bool, baseProductName NullableString, baseProductShort NullableString, baseProductVersion NullableString, arch string, buildTimestamp float64, instimage NullableString, mainimage NullableString, discnum NullableInt64, totaldiscs NullableInt64, addons []AddonResponse, checksums []ChecksumResponse, images []ImageResponse, variants []VariantResponse, ) *RpmDistributionTreeResponse`

NewRpmDistributionTreeResponse instantiates a new RpmDistributionTreeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpmDistributionTreeResponseWithDefaults

`func NewRpmDistributionTreeResponseWithDefaults() *RpmDistributionTreeResponse`

NewRpmDistributionTreeResponseWithDefaults instantiates a new RpmDistributionTreeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPulpHref

`func (o *RpmDistributionTreeResponse) GetPulpHref() string`

GetPulpHref returns the PulpHref field if non-nil, zero value otherwise.

### GetPulpHrefOk

`func (o *RpmDistributionTreeResponse) GetPulpHrefOk() (*string, bool)`

GetPulpHrefOk returns a tuple with the PulpHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpHref

`func (o *RpmDistributionTreeResponse) SetPulpHref(v string)`

SetPulpHref sets PulpHref field to given value.

### HasPulpHref

`func (o *RpmDistributionTreeResponse) HasPulpHref() bool`

HasPulpHref returns a boolean if a field has been set.

### GetHeaderVersion

`func (o *RpmDistributionTreeResponse) GetHeaderVersion() string`

GetHeaderVersion returns the HeaderVersion field if non-nil, zero value otherwise.

### GetHeaderVersionOk

`func (o *RpmDistributionTreeResponse) GetHeaderVersionOk() (*string, bool)`

GetHeaderVersionOk returns a tuple with the HeaderVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderVersion

`func (o *RpmDistributionTreeResponse) SetHeaderVersion(v string)`

SetHeaderVersion sets HeaderVersion field to given value.


### GetReleaseName

`func (o *RpmDistributionTreeResponse) GetReleaseName() string`

GetReleaseName returns the ReleaseName field if non-nil, zero value otherwise.

### GetReleaseNameOk

`func (o *RpmDistributionTreeResponse) GetReleaseNameOk() (*string, bool)`

GetReleaseNameOk returns a tuple with the ReleaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseName

`func (o *RpmDistributionTreeResponse) SetReleaseName(v string)`

SetReleaseName sets ReleaseName field to given value.


### GetReleaseShort

`func (o *RpmDistributionTreeResponse) GetReleaseShort() string`

GetReleaseShort returns the ReleaseShort field if non-nil, zero value otherwise.

### GetReleaseShortOk

`func (o *RpmDistributionTreeResponse) GetReleaseShortOk() (*string, bool)`

GetReleaseShortOk returns a tuple with the ReleaseShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseShort

`func (o *RpmDistributionTreeResponse) SetReleaseShort(v string)`

SetReleaseShort sets ReleaseShort field to given value.


### GetReleaseVersion

`func (o *RpmDistributionTreeResponse) GetReleaseVersion() string`

GetReleaseVersion returns the ReleaseVersion field if non-nil, zero value otherwise.

### GetReleaseVersionOk

`func (o *RpmDistributionTreeResponse) GetReleaseVersionOk() (*string, bool)`

GetReleaseVersionOk returns a tuple with the ReleaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseVersion

`func (o *RpmDistributionTreeResponse) SetReleaseVersion(v string)`

SetReleaseVersion sets ReleaseVersion field to given value.


### GetReleaseIsLayered

`func (o *RpmDistributionTreeResponse) GetReleaseIsLayered() bool`

GetReleaseIsLayered returns the ReleaseIsLayered field if non-nil, zero value otherwise.

### GetReleaseIsLayeredOk

`func (o *RpmDistributionTreeResponse) GetReleaseIsLayeredOk() (*bool, bool)`

GetReleaseIsLayeredOk returns a tuple with the ReleaseIsLayered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseIsLayered

`func (o *RpmDistributionTreeResponse) SetReleaseIsLayered(v bool)`

SetReleaseIsLayered sets ReleaseIsLayered field to given value.


### GetBaseProductName

`func (o *RpmDistributionTreeResponse) GetBaseProductName() string`

GetBaseProductName returns the BaseProductName field if non-nil, zero value otherwise.

### GetBaseProductNameOk

`func (o *RpmDistributionTreeResponse) GetBaseProductNameOk() (*string, bool)`

GetBaseProductNameOk returns a tuple with the BaseProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseProductName

`func (o *RpmDistributionTreeResponse) SetBaseProductName(v string)`

SetBaseProductName sets BaseProductName field to given value.


### SetBaseProductNameNil

`func (o *RpmDistributionTreeResponse) SetBaseProductNameNil(b bool)`

 SetBaseProductNameNil sets the value for BaseProductName to be an explicit nil

### UnsetBaseProductName
`func (o *RpmDistributionTreeResponse) UnsetBaseProductName()`

UnsetBaseProductName ensures that no value is present for BaseProductName, not even an explicit nil
### GetBaseProductShort

`func (o *RpmDistributionTreeResponse) GetBaseProductShort() string`

GetBaseProductShort returns the BaseProductShort field if non-nil, zero value otherwise.

### GetBaseProductShortOk

`func (o *RpmDistributionTreeResponse) GetBaseProductShortOk() (*string, bool)`

GetBaseProductShortOk returns a tuple with the BaseProductShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseProductShort

`func (o *RpmDistributionTreeResponse) SetBaseProductShort(v string)`

SetBaseProductShort sets BaseProductShort field to given value.


### SetBaseProductShortNil

`func (o *RpmDistributionTreeResponse) SetBaseProductShortNil(b bool)`

 SetBaseProductShortNil sets the value for BaseProductShort to be an explicit nil

### UnsetBaseProductShort
`func (o *RpmDistributionTreeResponse) UnsetBaseProductShort()`

UnsetBaseProductShort ensures that no value is present for BaseProductShort, not even an explicit nil
### GetBaseProductVersion

`func (o *RpmDistributionTreeResponse) GetBaseProductVersion() string`

GetBaseProductVersion returns the BaseProductVersion field if non-nil, zero value otherwise.

### GetBaseProductVersionOk

`func (o *RpmDistributionTreeResponse) GetBaseProductVersionOk() (*string, bool)`

GetBaseProductVersionOk returns a tuple with the BaseProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseProductVersion

`func (o *RpmDistributionTreeResponse) SetBaseProductVersion(v string)`

SetBaseProductVersion sets BaseProductVersion field to given value.


### SetBaseProductVersionNil

`func (o *RpmDistributionTreeResponse) SetBaseProductVersionNil(b bool)`

 SetBaseProductVersionNil sets the value for BaseProductVersion to be an explicit nil

### UnsetBaseProductVersion
`func (o *RpmDistributionTreeResponse) UnsetBaseProductVersion()`

UnsetBaseProductVersion ensures that no value is present for BaseProductVersion, not even an explicit nil
### GetArch

`func (o *RpmDistributionTreeResponse) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *RpmDistributionTreeResponse) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *RpmDistributionTreeResponse) SetArch(v string)`

SetArch sets Arch field to given value.


### GetBuildTimestamp

`func (o *RpmDistributionTreeResponse) GetBuildTimestamp() float64`

GetBuildTimestamp returns the BuildTimestamp field if non-nil, zero value otherwise.

### GetBuildTimestampOk

`func (o *RpmDistributionTreeResponse) GetBuildTimestampOk() (*float64, bool)`

GetBuildTimestampOk returns a tuple with the BuildTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildTimestamp

`func (o *RpmDistributionTreeResponse) SetBuildTimestamp(v float64)`

SetBuildTimestamp sets BuildTimestamp field to given value.


### GetInstimage

`func (o *RpmDistributionTreeResponse) GetInstimage() string`

GetInstimage returns the Instimage field if non-nil, zero value otherwise.

### GetInstimageOk

`func (o *RpmDistributionTreeResponse) GetInstimageOk() (*string, bool)`

GetInstimageOk returns a tuple with the Instimage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstimage

`func (o *RpmDistributionTreeResponse) SetInstimage(v string)`

SetInstimage sets Instimage field to given value.


### SetInstimageNil

`func (o *RpmDistributionTreeResponse) SetInstimageNil(b bool)`

 SetInstimageNil sets the value for Instimage to be an explicit nil

### UnsetInstimage
`func (o *RpmDistributionTreeResponse) UnsetInstimage()`

UnsetInstimage ensures that no value is present for Instimage, not even an explicit nil
### GetMainimage

`func (o *RpmDistributionTreeResponse) GetMainimage() string`

GetMainimage returns the Mainimage field if non-nil, zero value otherwise.

### GetMainimageOk

`func (o *RpmDistributionTreeResponse) GetMainimageOk() (*string, bool)`

GetMainimageOk returns a tuple with the Mainimage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainimage

`func (o *RpmDistributionTreeResponse) SetMainimage(v string)`

SetMainimage sets Mainimage field to given value.


### SetMainimageNil

`func (o *RpmDistributionTreeResponse) SetMainimageNil(b bool)`

 SetMainimageNil sets the value for Mainimage to be an explicit nil

### UnsetMainimage
`func (o *RpmDistributionTreeResponse) UnsetMainimage()`

UnsetMainimage ensures that no value is present for Mainimage, not even an explicit nil
### GetDiscnum

`func (o *RpmDistributionTreeResponse) GetDiscnum() int64`

GetDiscnum returns the Discnum field if non-nil, zero value otherwise.

### GetDiscnumOk

`func (o *RpmDistributionTreeResponse) GetDiscnumOk() (*int64, bool)`

GetDiscnumOk returns a tuple with the Discnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscnum

`func (o *RpmDistributionTreeResponse) SetDiscnum(v int64)`

SetDiscnum sets Discnum field to given value.


### SetDiscnumNil

`func (o *RpmDistributionTreeResponse) SetDiscnumNil(b bool)`

 SetDiscnumNil sets the value for Discnum to be an explicit nil

### UnsetDiscnum
`func (o *RpmDistributionTreeResponse) UnsetDiscnum()`

UnsetDiscnum ensures that no value is present for Discnum, not even an explicit nil
### GetTotaldiscs

`func (o *RpmDistributionTreeResponse) GetTotaldiscs() int64`

GetTotaldiscs returns the Totaldiscs field if non-nil, zero value otherwise.

### GetTotaldiscsOk

`func (o *RpmDistributionTreeResponse) GetTotaldiscsOk() (*int64, bool)`

GetTotaldiscsOk returns a tuple with the Totaldiscs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotaldiscs

`func (o *RpmDistributionTreeResponse) SetTotaldiscs(v int64)`

SetTotaldiscs sets Totaldiscs field to given value.


### SetTotaldiscsNil

`func (o *RpmDistributionTreeResponse) SetTotaldiscsNil(b bool)`

 SetTotaldiscsNil sets the value for Totaldiscs to be an explicit nil

### UnsetTotaldiscs
`func (o *RpmDistributionTreeResponse) UnsetTotaldiscs()`

UnsetTotaldiscs ensures that no value is present for Totaldiscs, not even an explicit nil
### GetAddons

`func (o *RpmDistributionTreeResponse) GetAddons() []AddonResponse`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *RpmDistributionTreeResponse) GetAddonsOk() (*[]AddonResponse, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *RpmDistributionTreeResponse) SetAddons(v []AddonResponse)`

SetAddons sets Addons field to given value.


### GetChecksums

`func (o *RpmDistributionTreeResponse) GetChecksums() []ChecksumResponse`

GetChecksums returns the Checksums field if non-nil, zero value otherwise.

### GetChecksumsOk

`func (o *RpmDistributionTreeResponse) GetChecksumsOk() (*[]ChecksumResponse, bool)`

GetChecksumsOk returns a tuple with the Checksums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksums

`func (o *RpmDistributionTreeResponse) SetChecksums(v []ChecksumResponse)`

SetChecksums sets Checksums field to given value.


### GetImages

`func (o *RpmDistributionTreeResponse) GetImages() []ImageResponse`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *RpmDistributionTreeResponse) GetImagesOk() (*[]ImageResponse, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *RpmDistributionTreeResponse) SetImages(v []ImageResponse)`

SetImages sets Images field to given value.


### GetVariants

`func (o *RpmDistributionTreeResponse) GetVariants() []VariantResponse`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *RpmDistributionTreeResponse) GetVariantsOk() (*[]VariantResponse, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *RpmDistributionTreeResponse) SetVariants(v []VariantResponse)`

SetVariants sets Variants field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


