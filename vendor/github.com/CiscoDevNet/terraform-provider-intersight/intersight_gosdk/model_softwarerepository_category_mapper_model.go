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

// SoftwarerepositoryCategoryMapperModel Maps a Cisco hardware model Series to its applicable hardware models.
type SoftwarerepositoryCategoryMapperModel struct {
	CapabilityCapability
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The category of the model series.
	Category *string `json:"Category,omitempty"`
	// The distributable tag value of the model series.
	DistTag *string `json:"DistTag,omitempty"`
	// The regex that all images of this model follow.
	RegexPattern *string `json:"RegexPattern,omitempty"`
	// Cisco hardware model series.
	SeriesId             *string  `json:"SeriesId,omitempty"`
	SupportedModels      []string `json:"SupportedModels,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwarerepositoryCategoryMapperModel SoftwarerepositoryCategoryMapperModel

// NewSoftwarerepositoryCategoryMapperModel instantiates a new SoftwarerepositoryCategoryMapperModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryCategoryMapperModel(classId string, objectType string) *SoftwarerepositoryCategoryMapperModel {
	this := SoftwarerepositoryCategoryMapperModel{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSoftwarerepositoryCategoryMapperModelWithDefaults instantiates a new SoftwarerepositoryCategoryMapperModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryCategoryMapperModelWithDefaults() *SoftwarerepositoryCategoryMapperModel {
	this := SoftwarerepositoryCategoryMapperModel{}
	var classId string = "softwarerepository.CategoryMapperModel"
	this.ClassId = classId
	var objectType string = "softwarerepository.CategoryMapperModel"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwarerepositoryCategoryMapperModel) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperModel) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwarerepositoryCategoryMapperModel) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SoftwarerepositoryCategoryMapperModel) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperModel) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwarerepositoryCategoryMapperModel) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperModel) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperModel) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperModel) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *SoftwarerepositoryCategoryMapperModel) SetCategory(v string) {
	o.Category = &v
}

// GetDistTag returns the DistTag field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperModel) GetDistTag() string {
	if o == nil || o.DistTag == nil {
		var ret string
		return ret
	}
	return *o.DistTag
}

// GetDistTagOk returns a tuple with the DistTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperModel) GetDistTagOk() (*string, bool) {
	if o == nil || o.DistTag == nil {
		return nil, false
	}
	return o.DistTag, true
}

// HasDistTag returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperModel) HasDistTag() bool {
	if o != nil && o.DistTag != nil {
		return true
	}

	return false
}

// SetDistTag gets a reference to the given string and assigns it to the DistTag field.
func (o *SoftwarerepositoryCategoryMapperModel) SetDistTag(v string) {
	o.DistTag = &v
}

// GetRegexPattern returns the RegexPattern field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperModel) GetRegexPattern() string {
	if o == nil || o.RegexPattern == nil {
		var ret string
		return ret
	}
	return *o.RegexPattern
}

// GetRegexPatternOk returns a tuple with the RegexPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperModel) GetRegexPatternOk() (*string, bool) {
	if o == nil || o.RegexPattern == nil {
		return nil, false
	}
	return o.RegexPattern, true
}

// HasRegexPattern returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperModel) HasRegexPattern() bool {
	if o != nil && o.RegexPattern != nil {
		return true
	}

	return false
}

// SetRegexPattern gets a reference to the given string and assigns it to the RegexPattern field.
func (o *SoftwarerepositoryCategoryMapperModel) SetRegexPattern(v string) {
	o.RegexPattern = &v
}

// GetSeriesId returns the SeriesId field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperModel) GetSeriesId() string {
	if o == nil || o.SeriesId == nil {
		var ret string
		return ret
	}
	return *o.SeriesId
}

// GetSeriesIdOk returns a tuple with the SeriesId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperModel) GetSeriesIdOk() (*string, bool) {
	if o == nil || o.SeriesId == nil {
		return nil, false
	}
	return o.SeriesId, true
}

// HasSeriesId returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperModel) HasSeriesId() bool {
	if o != nil && o.SeriesId != nil {
		return true
	}

	return false
}

// SetSeriesId gets a reference to the given string and assigns it to the SeriesId field.
func (o *SoftwarerepositoryCategoryMapperModel) SetSeriesId(v string) {
	o.SeriesId = &v
}

// GetSupportedModels returns the SupportedModels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoftwarerepositoryCategoryMapperModel) GetSupportedModels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SupportedModels
}

// GetSupportedModelsOk returns a tuple with the SupportedModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoftwarerepositoryCategoryMapperModel) GetSupportedModelsOk() (*[]string, bool) {
	if o == nil || o.SupportedModels == nil {
		return nil, false
	}
	return &o.SupportedModels, true
}

// HasSupportedModels returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperModel) HasSupportedModels() bool {
	if o != nil && o.SupportedModels != nil {
		return true
	}

	return false
}

// SetSupportedModels gets a reference to the given []string and assigns it to the SupportedModels field.
func (o *SoftwarerepositoryCategoryMapperModel) SetSupportedModels(v []string) {
	o.SupportedModels = v
}

func (o SoftwarerepositoryCategoryMapperModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCapabilityCapability, errCapabilityCapability := json.Marshal(o.CapabilityCapability)
	if errCapabilityCapability != nil {
		return []byte{}, errCapabilityCapability
	}
	errCapabilityCapability = json.Unmarshal([]byte(serializedCapabilityCapability), &toSerialize)
	if errCapabilityCapability != nil {
		return []byte{}, errCapabilityCapability
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Category != nil {
		toSerialize["Category"] = o.Category
	}
	if o.DistTag != nil {
		toSerialize["DistTag"] = o.DistTag
	}
	if o.RegexPattern != nil {
		toSerialize["RegexPattern"] = o.RegexPattern
	}
	if o.SeriesId != nil {
		toSerialize["SeriesId"] = o.SeriesId
	}
	if o.SupportedModels != nil {
		toSerialize["SupportedModels"] = o.SupportedModels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SoftwarerepositoryCategoryMapperModel) UnmarshalJSON(bytes []byte) (err error) {
	type SoftwarerepositoryCategoryMapperModelWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The category of the model series.
		Category *string `json:"Category,omitempty"`
		// The distributable tag value of the model series.
		DistTag *string `json:"DistTag,omitempty"`
		// The regex that all images of this model follow.
		RegexPattern *string `json:"RegexPattern,omitempty"`
		// Cisco hardware model series.
		SeriesId        *string  `json:"SeriesId,omitempty"`
		SupportedModels []string `json:"SupportedModels,omitempty"`
	}

	varSoftwarerepositoryCategoryMapperModelWithoutEmbeddedStruct := SoftwarerepositoryCategoryMapperModelWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varSoftwarerepositoryCategoryMapperModelWithoutEmbeddedStruct)
	if err == nil {
		varSoftwarerepositoryCategoryMapperModel := _SoftwarerepositoryCategoryMapperModel{}
		varSoftwarerepositoryCategoryMapperModel.ClassId = varSoftwarerepositoryCategoryMapperModelWithoutEmbeddedStruct.ClassId
		varSoftwarerepositoryCategoryMapperModel.ObjectType = varSoftwarerepositoryCategoryMapperModelWithoutEmbeddedStruct.ObjectType
		varSoftwarerepositoryCategoryMapperModel.Category = varSoftwarerepositoryCategoryMapperModelWithoutEmbeddedStruct.Category
		varSoftwarerepositoryCategoryMapperModel.DistTag = varSoftwarerepositoryCategoryMapperModelWithoutEmbeddedStruct.DistTag
		varSoftwarerepositoryCategoryMapperModel.RegexPattern = varSoftwarerepositoryCategoryMapperModelWithoutEmbeddedStruct.RegexPattern
		varSoftwarerepositoryCategoryMapperModel.SeriesId = varSoftwarerepositoryCategoryMapperModelWithoutEmbeddedStruct.SeriesId
		varSoftwarerepositoryCategoryMapperModel.SupportedModels = varSoftwarerepositoryCategoryMapperModelWithoutEmbeddedStruct.SupportedModels
		*o = SoftwarerepositoryCategoryMapperModel(varSoftwarerepositoryCategoryMapperModel)
	} else {
		return err
	}

	varSoftwarerepositoryCategoryMapperModel := _SoftwarerepositoryCategoryMapperModel{}

	err = json.Unmarshal(bytes, &varSoftwarerepositoryCategoryMapperModel)
	if err == nil {
		o.CapabilityCapability = varSoftwarerepositoryCategoryMapperModel.CapabilityCapability
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Category")
		delete(additionalProperties, "DistTag")
		delete(additionalProperties, "RegexPattern")
		delete(additionalProperties, "SeriesId")
		delete(additionalProperties, "SupportedModels")

		// remove fields from embedded structs
		reflectCapabilityCapability := reflect.ValueOf(o.CapabilityCapability)
		for i := 0; i < reflectCapabilityCapability.Type().NumField(); i++ {
			t := reflectCapabilityCapability.Type().Field(i)

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

type NullableSoftwarerepositoryCategoryMapperModel struct {
	value *SoftwarerepositoryCategoryMapperModel
	isSet bool
}

func (v NullableSoftwarerepositoryCategoryMapperModel) Get() *SoftwarerepositoryCategoryMapperModel {
	return v.value
}

func (v *NullableSoftwarerepositoryCategoryMapperModel) Set(val *SoftwarerepositoryCategoryMapperModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryCategoryMapperModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryCategoryMapperModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryCategoryMapperModel(val *SoftwarerepositoryCategoryMapperModel) *NullableSoftwarerepositoryCategoryMapperModel {
	return &NullableSoftwarerepositoryCategoryMapperModel{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryCategoryMapperModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryCategoryMapperModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
