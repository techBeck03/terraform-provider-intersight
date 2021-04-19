/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-04-12T05:47:20Z.
 *
 * API version: 1.0.9-4240
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// SoftwarerepositoryLocalMachine The user's local machine that is being used as the source for an image.
type SoftwarerepositoryLocalMachine struct {
	SoftwarerepositoryFileServer
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// When the import action in the file MO is updated with 'GeneratePreSignedDownloadUrl', Intersight returns a pre-signed URL in this property as part of the patch response. The user is expected to subsequently download the file using this URL.
	DownloadUrl *string `json:"DownloadUrl,omitempty"`
	// The chunk size (in bytes) for each part of the file to be uploaded.
	PartSize *int64 `json:"PartSize,omitempty"`
	// When the import action in file MO is updated with 'GeneratePreSignedUploadUrl', Intersight shall return a upload Id in this property as part of the PATCH response.
	UploadId *string `json:"UploadId,omitempty"`
	// When a file MO is created with 'LocalMachine' as the source, Intersight returns a pre-signed URL in this property as part of the POST response. The user is expected to subsequently upload the file content using this URL. Once the upload is completed, the user is expected to patch the uploader object's transfer state to success.
	UploadUrl            *string  `json:"UploadUrl,omitempty"`
	UploadUrls           []string `json:"UploadUrls,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwarerepositoryLocalMachine SoftwarerepositoryLocalMachine

// NewSoftwarerepositoryLocalMachine instantiates a new SoftwarerepositoryLocalMachine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryLocalMachine(classId string, objectType string) *SoftwarerepositoryLocalMachine {
	this := SoftwarerepositoryLocalMachine{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSoftwarerepositoryLocalMachineWithDefaults instantiates a new SoftwarerepositoryLocalMachine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryLocalMachineWithDefaults() *SoftwarerepositoryLocalMachine {
	this := SoftwarerepositoryLocalMachine{}
	var classId string = "softwarerepository.LocalMachine"
	this.ClassId = classId
	var objectType string = "softwarerepository.LocalMachine"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwarerepositoryLocalMachine) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryLocalMachine) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwarerepositoryLocalMachine) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SoftwarerepositoryLocalMachine) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryLocalMachine) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwarerepositoryLocalMachine) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDownloadUrl returns the DownloadUrl field value if set, zero value otherwise.
func (o *SoftwarerepositoryLocalMachine) GetDownloadUrl() string {
	if o == nil || o.DownloadUrl == nil {
		var ret string
		return ret
	}
	return *o.DownloadUrl
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryLocalMachine) GetDownloadUrlOk() (*string, bool) {
	if o == nil || o.DownloadUrl == nil {
		return nil, false
	}
	return o.DownloadUrl, true
}

// HasDownloadUrl returns a boolean if a field has been set.
func (o *SoftwarerepositoryLocalMachine) HasDownloadUrl() bool {
	if o != nil && o.DownloadUrl != nil {
		return true
	}

	return false
}

// SetDownloadUrl gets a reference to the given string and assigns it to the DownloadUrl field.
func (o *SoftwarerepositoryLocalMachine) SetDownloadUrl(v string) {
	o.DownloadUrl = &v
}

// GetPartSize returns the PartSize field value if set, zero value otherwise.
func (o *SoftwarerepositoryLocalMachine) GetPartSize() int64 {
	if o == nil || o.PartSize == nil {
		var ret int64
		return ret
	}
	return *o.PartSize
}

// GetPartSizeOk returns a tuple with the PartSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryLocalMachine) GetPartSizeOk() (*int64, bool) {
	if o == nil || o.PartSize == nil {
		return nil, false
	}
	return o.PartSize, true
}

// HasPartSize returns a boolean if a field has been set.
func (o *SoftwarerepositoryLocalMachine) HasPartSize() bool {
	if o != nil && o.PartSize != nil {
		return true
	}

	return false
}

// SetPartSize gets a reference to the given int64 and assigns it to the PartSize field.
func (o *SoftwarerepositoryLocalMachine) SetPartSize(v int64) {
	o.PartSize = &v
}

// GetUploadId returns the UploadId field value if set, zero value otherwise.
func (o *SoftwarerepositoryLocalMachine) GetUploadId() string {
	if o == nil || o.UploadId == nil {
		var ret string
		return ret
	}
	return *o.UploadId
}

// GetUploadIdOk returns a tuple with the UploadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryLocalMachine) GetUploadIdOk() (*string, bool) {
	if o == nil || o.UploadId == nil {
		return nil, false
	}
	return o.UploadId, true
}

// HasUploadId returns a boolean if a field has been set.
func (o *SoftwarerepositoryLocalMachine) HasUploadId() bool {
	if o != nil && o.UploadId != nil {
		return true
	}

	return false
}

// SetUploadId gets a reference to the given string and assigns it to the UploadId field.
func (o *SoftwarerepositoryLocalMachine) SetUploadId(v string) {
	o.UploadId = &v
}

// GetUploadUrl returns the UploadUrl field value if set, zero value otherwise.
func (o *SoftwarerepositoryLocalMachine) GetUploadUrl() string {
	if o == nil || o.UploadUrl == nil {
		var ret string
		return ret
	}
	return *o.UploadUrl
}

// GetUploadUrlOk returns a tuple with the UploadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryLocalMachine) GetUploadUrlOk() (*string, bool) {
	if o == nil || o.UploadUrl == nil {
		return nil, false
	}
	return o.UploadUrl, true
}

// HasUploadUrl returns a boolean if a field has been set.
func (o *SoftwarerepositoryLocalMachine) HasUploadUrl() bool {
	if o != nil && o.UploadUrl != nil {
		return true
	}

	return false
}

// SetUploadUrl gets a reference to the given string and assigns it to the UploadUrl field.
func (o *SoftwarerepositoryLocalMachine) SetUploadUrl(v string) {
	o.UploadUrl = &v
}

// GetUploadUrls returns the UploadUrls field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoftwarerepositoryLocalMachine) GetUploadUrls() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.UploadUrls
}

// GetUploadUrlsOk returns a tuple with the UploadUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoftwarerepositoryLocalMachine) GetUploadUrlsOk() (*[]string, bool) {
	if o == nil || o.UploadUrls == nil {
		return nil, false
	}
	return &o.UploadUrls, true
}

// HasUploadUrls returns a boolean if a field has been set.
func (o *SoftwarerepositoryLocalMachine) HasUploadUrls() bool {
	if o != nil && o.UploadUrls != nil {
		return true
	}

	return false
}

// SetUploadUrls gets a reference to the given []string and assigns it to the UploadUrls field.
func (o *SoftwarerepositoryLocalMachine) SetUploadUrls(v []string) {
	o.UploadUrls = v
}

func (o SoftwarerepositoryLocalMachine) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedSoftwarerepositoryFileServer, errSoftwarerepositoryFileServer := json.Marshal(o.SoftwarerepositoryFileServer)
	if errSoftwarerepositoryFileServer != nil {
		return []byte{}, errSoftwarerepositoryFileServer
	}
	errSoftwarerepositoryFileServer = json.Unmarshal([]byte(serializedSoftwarerepositoryFileServer), &toSerialize)
	if errSoftwarerepositoryFileServer != nil {
		return []byte{}, errSoftwarerepositoryFileServer
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DownloadUrl != nil {
		toSerialize["DownloadUrl"] = o.DownloadUrl
	}
	if o.PartSize != nil {
		toSerialize["PartSize"] = o.PartSize
	}
	if o.UploadId != nil {
		toSerialize["UploadId"] = o.UploadId
	}
	if o.UploadUrl != nil {
		toSerialize["UploadUrl"] = o.UploadUrl
	}
	if o.UploadUrls != nil {
		toSerialize["UploadUrls"] = o.UploadUrls
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SoftwarerepositoryLocalMachine) UnmarshalJSON(bytes []byte) (err error) {
	type SoftwarerepositoryLocalMachineWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// When the import action in the file MO is updated with 'GeneratePreSignedDownloadUrl', Intersight returns a pre-signed URL in this property as part of the patch response. The user is expected to subsequently download the file using this URL.
		DownloadUrl *string `json:"DownloadUrl,omitempty"`
		// The chunk size (in bytes) for each part of the file to be uploaded.
		PartSize *int64 `json:"PartSize,omitempty"`
		// When the import action in file MO is updated with 'GeneratePreSignedUploadUrl', Intersight shall return a upload Id in this property as part of the PATCH response.
		UploadId *string `json:"UploadId,omitempty"`
		// When a file MO is created with 'LocalMachine' as the source, Intersight returns a pre-signed URL in this property as part of the POST response. The user is expected to subsequently upload the file content using this URL. Once the upload is completed, the user is expected to patch the uploader object's transfer state to success.
		UploadUrl  *string  `json:"UploadUrl,omitempty"`
		UploadUrls []string `json:"UploadUrls,omitempty"`
	}

	varSoftwarerepositoryLocalMachineWithoutEmbeddedStruct := SoftwarerepositoryLocalMachineWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varSoftwarerepositoryLocalMachineWithoutEmbeddedStruct)
	if err == nil {
		varSoftwarerepositoryLocalMachine := _SoftwarerepositoryLocalMachine{}
		varSoftwarerepositoryLocalMachine.ClassId = varSoftwarerepositoryLocalMachineWithoutEmbeddedStruct.ClassId
		varSoftwarerepositoryLocalMachine.ObjectType = varSoftwarerepositoryLocalMachineWithoutEmbeddedStruct.ObjectType
		varSoftwarerepositoryLocalMachine.DownloadUrl = varSoftwarerepositoryLocalMachineWithoutEmbeddedStruct.DownloadUrl
		varSoftwarerepositoryLocalMachine.PartSize = varSoftwarerepositoryLocalMachineWithoutEmbeddedStruct.PartSize
		varSoftwarerepositoryLocalMachine.UploadId = varSoftwarerepositoryLocalMachineWithoutEmbeddedStruct.UploadId
		varSoftwarerepositoryLocalMachine.UploadUrl = varSoftwarerepositoryLocalMachineWithoutEmbeddedStruct.UploadUrl
		varSoftwarerepositoryLocalMachine.UploadUrls = varSoftwarerepositoryLocalMachineWithoutEmbeddedStruct.UploadUrls
		*o = SoftwarerepositoryLocalMachine(varSoftwarerepositoryLocalMachine)
	} else {
		return err
	}

	varSoftwarerepositoryLocalMachine := _SoftwarerepositoryLocalMachine{}

	err = json.Unmarshal(bytes, &varSoftwarerepositoryLocalMachine)
	if err == nil {
		o.SoftwarerepositoryFileServer = varSoftwarerepositoryLocalMachine.SoftwarerepositoryFileServer
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DownloadUrl")
		delete(additionalProperties, "PartSize")
		delete(additionalProperties, "UploadId")
		delete(additionalProperties, "UploadUrl")
		delete(additionalProperties, "UploadUrls")

		// remove fields from embedded structs
		reflectSoftwarerepositoryFileServer := reflect.ValueOf(o.SoftwarerepositoryFileServer)
		for i := 0; i < reflectSoftwarerepositoryFileServer.Type().NumField(); i++ {
			t := reflectSoftwarerepositoryFileServer.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSoftwarerepositoryLocalMachine struct {
	value *SoftwarerepositoryLocalMachine
	isSet bool
}

func (v NullableSoftwarerepositoryLocalMachine) Get() *SoftwarerepositoryLocalMachine {
	return v.value
}

func (v *NullableSoftwarerepositoryLocalMachine) Set(val *SoftwarerepositoryLocalMachine) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryLocalMachine) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryLocalMachine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryLocalMachine(val *SoftwarerepositoryLocalMachine) *NullableSoftwarerepositoryLocalMachine {
	return &NullableSoftwarerepositoryLocalMachine{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryLocalMachine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryLocalMachine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
