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

// HyperflexFeatureLimitInternal The HyperFlex installer feature limits for internal system use.
type HyperflexFeatureLimitInternal struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                           `json:"ObjectType"`
	FeatureLimitEntries  []HyperflexFeatureLimitEntry     `json:"FeatureLimitEntries,omitempty"`
	AppCatalog           *HyperflexAppCatalogRelationship `json:"AppCatalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexFeatureLimitInternal HyperflexFeatureLimitInternal

// NewHyperflexFeatureLimitInternal instantiates a new HyperflexFeatureLimitInternal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexFeatureLimitInternal(classId string, objectType string) *HyperflexFeatureLimitInternal {
	this := HyperflexFeatureLimitInternal{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexFeatureLimitInternalWithDefaults instantiates a new HyperflexFeatureLimitInternal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexFeatureLimitInternalWithDefaults() *HyperflexFeatureLimitInternal {
	this := HyperflexFeatureLimitInternal{}
	var classId string = "hyperflex.FeatureLimitInternal"
	this.ClassId = classId
	var objectType string = "hyperflex.FeatureLimitInternal"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexFeatureLimitInternal) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexFeatureLimitInternal) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexFeatureLimitInternal) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexFeatureLimitInternal) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexFeatureLimitInternal) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexFeatureLimitInternal) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFeatureLimitEntries returns the FeatureLimitEntries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexFeatureLimitInternal) GetFeatureLimitEntries() []HyperflexFeatureLimitEntry {
	if o == nil {
		var ret []HyperflexFeatureLimitEntry
		return ret
	}
	return o.FeatureLimitEntries
}

// GetFeatureLimitEntriesOk returns a tuple with the FeatureLimitEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexFeatureLimitInternal) GetFeatureLimitEntriesOk() (*[]HyperflexFeatureLimitEntry, bool) {
	if o == nil || o.FeatureLimitEntries == nil {
		return nil, false
	}
	return &o.FeatureLimitEntries, true
}

// HasFeatureLimitEntries returns a boolean if a field has been set.
func (o *HyperflexFeatureLimitInternal) HasFeatureLimitEntries() bool {
	if o != nil && o.FeatureLimitEntries != nil {
		return true
	}

	return false
}

// SetFeatureLimitEntries gets a reference to the given []HyperflexFeatureLimitEntry and assigns it to the FeatureLimitEntries field.
func (o *HyperflexFeatureLimitInternal) SetFeatureLimitEntries(v []HyperflexFeatureLimitEntry) {
	o.FeatureLimitEntries = v
}

// GetAppCatalog returns the AppCatalog field value if set, zero value otherwise.
func (o *HyperflexFeatureLimitInternal) GetAppCatalog() HyperflexAppCatalogRelationship {
	if o == nil || o.AppCatalog == nil {
		var ret HyperflexAppCatalogRelationship
		return ret
	}
	return *o.AppCatalog
}

// GetAppCatalogOk returns a tuple with the AppCatalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexFeatureLimitInternal) GetAppCatalogOk() (*HyperflexAppCatalogRelationship, bool) {
	if o == nil || o.AppCatalog == nil {
		return nil, false
	}
	return o.AppCatalog, true
}

// HasAppCatalog returns a boolean if a field has been set.
func (o *HyperflexFeatureLimitInternal) HasAppCatalog() bool {
	if o != nil && o.AppCatalog != nil {
		return true
	}

	return false
}

// SetAppCatalog gets a reference to the given HyperflexAppCatalogRelationship and assigns it to the AppCatalog field.
func (o *HyperflexFeatureLimitInternal) SetAppCatalog(v HyperflexAppCatalogRelationship) {
	o.AppCatalog = &v
}

func (o HyperflexFeatureLimitInternal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FeatureLimitEntries != nil {
		toSerialize["FeatureLimitEntries"] = o.FeatureLimitEntries
	}
	if o.AppCatalog != nil {
		toSerialize["AppCatalog"] = o.AppCatalog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexFeatureLimitInternal) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexFeatureLimitInternalWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType          string                           `json:"ObjectType"`
		FeatureLimitEntries []HyperflexFeatureLimitEntry     `json:"FeatureLimitEntries,omitempty"`
		AppCatalog          *HyperflexAppCatalogRelationship `json:"AppCatalog,omitempty"`
	}

	varHyperflexFeatureLimitInternalWithoutEmbeddedStruct := HyperflexFeatureLimitInternalWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexFeatureLimitInternalWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexFeatureLimitInternal := _HyperflexFeatureLimitInternal{}
		varHyperflexFeatureLimitInternal.ClassId = varHyperflexFeatureLimitInternalWithoutEmbeddedStruct.ClassId
		varHyperflexFeatureLimitInternal.ObjectType = varHyperflexFeatureLimitInternalWithoutEmbeddedStruct.ObjectType
		varHyperflexFeatureLimitInternal.FeatureLimitEntries = varHyperflexFeatureLimitInternalWithoutEmbeddedStruct.FeatureLimitEntries
		varHyperflexFeatureLimitInternal.AppCatalog = varHyperflexFeatureLimitInternalWithoutEmbeddedStruct.AppCatalog
		*o = HyperflexFeatureLimitInternal(varHyperflexFeatureLimitInternal)
	} else {
		return err
	}

	varHyperflexFeatureLimitInternal := _HyperflexFeatureLimitInternal{}

	err = json.Unmarshal(bytes, &varHyperflexFeatureLimitInternal)
	if err == nil {
		o.MoBaseMo = varHyperflexFeatureLimitInternal.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FeatureLimitEntries")
		delete(additionalProperties, "AppCatalog")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableHyperflexFeatureLimitInternal struct {
	value *HyperflexFeatureLimitInternal
	isSet bool
}

func (v NullableHyperflexFeatureLimitInternal) Get() *HyperflexFeatureLimitInternal {
	return v.value
}

func (v *NullableHyperflexFeatureLimitInternal) Set(val *HyperflexFeatureLimitInternal) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexFeatureLimitInternal) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexFeatureLimitInternal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexFeatureLimitInternal(val *HyperflexFeatureLimitInternal) *NullableHyperflexFeatureLimitInternal {
	return &NullableHyperflexFeatureLimitInternal{value: val, isSet: true}
}

func (v NullableHyperflexFeatureLimitInternal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexFeatureLimitInternal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
