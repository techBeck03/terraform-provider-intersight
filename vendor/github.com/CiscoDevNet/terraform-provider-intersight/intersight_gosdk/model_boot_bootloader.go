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

// BootBootloader Lists the bootloader properties that can be overriden for boot from Local disk, NVME, SD Card, PCH Storage and SAN boot.
type BootBootloader struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Carries more information about the bootloader.
	Description *string `json:"Description,omitempty"`
	// Name of the bootloader image.
	Name *string `json:"Name,omitempty"`
	// Path to the bootloader image.
	Path                 *string `json:"Path,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BootBootloader BootBootloader

// NewBootBootloader instantiates a new BootBootloader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootBootloader(classId string, objectType string) *BootBootloader {
	this := BootBootloader{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewBootBootloaderWithDefaults instantiates a new BootBootloader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootBootloaderWithDefaults() *BootBootloader {
	this := BootBootloader{}
	var classId string = "boot.Bootloader"
	this.ClassId = classId
	var objectType string = "boot.Bootloader"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *BootBootloader) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BootBootloader) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BootBootloader) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BootBootloader) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BootBootloader) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BootBootloader) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BootBootloader) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootBootloader) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BootBootloader) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BootBootloader) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BootBootloader) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootBootloader) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BootBootloader) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BootBootloader) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *BootBootloader) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootBootloader) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *BootBootloader) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *BootBootloader) SetPath(v string) {
	o.Path = &v
}

func (o BootBootloader) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Path != nil {
		toSerialize["Path"] = o.Path
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BootBootloader) UnmarshalJSON(bytes []byte) (err error) {
	type BootBootloaderWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Carries more information about the bootloader.
		Description *string `json:"Description,omitempty"`
		// Name of the bootloader image.
		Name *string `json:"Name,omitempty"`
		// Path to the bootloader image.
		Path *string `json:"Path,omitempty"`
	}

	varBootBootloaderWithoutEmbeddedStruct := BootBootloaderWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varBootBootloaderWithoutEmbeddedStruct)
	if err == nil {
		varBootBootloader := _BootBootloader{}
		varBootBootloader.ClassId = varBootBootloaderWithoutEmbeddedStruct.ClassId
		varBootBootloader.ObjectType = varBootBootloaderWithoutEmbeddedStruct.ObjectType
		varBootBootloader.Description = varBootBootloaderWithoutEmbeddedStruct.Description
		varBootBootloader.Name = varBootBootloaderWithoutEmbeddedStruct.Name
		varBootBootloader.Path = varBootBootloaderWithoutEmbeddedStruct.Path
		*o = BootBootloader(varBootBootloader)
	} else {
		return err
	}

	varBootBootloader := _BootBootloader{}

	err = json.Unmarshal(bytes, &varBootBootloader)
	if err == nil {
		o.MoBaseComplexType = varBootBootloader.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Path")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableBootBootloader struct {
	value *BootBootloader
	isSet bool
}

func (v NullableBootBootloader) Get() *BootBootloader {
	return v.value
}

func (v *NullableBootBootloader) Set(val *BootBootloader) {
	v.value = val
	v.isSet = true
}

func (v NullableBootBootloader) IsSet() bool {
	return v.isSet
}

func (v *NullableBootBootloader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootBootloader(val *BootBootloader) *NullableBootBootloader {
	return &NullableBootBootloader{value: val, isSet: true}
}

func (v NullableBootBootloader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootBootloader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
