/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-03-10T06:51:24Z.
 *
 * API version: 1.0.9-3942
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// HyperflexMapClusterIdToStSnapshotPoint ClusterId to Snapshot Point map.
type HyperflexMapClusterIdToStSnapshotPoint struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// ClusterId of the snapshot point.
	ClusterId            *string                        `json:"ClusterId,omitempty"`
	SnapshotPoint        NullableHyperflexSnapshotPoint `json:"SnapshotPoint,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexMapClusterIdToStSnapshotPoint HyperflexMapClusterIdToStSnapshotPoint

// NewHyperflexMapClusterIdToStSnapshotPoint instantiates a new HyperflexMapClusterIdToStSnapshotPoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexMapClusterIdToStSnapshotPoint(classId string, objectType string) *HyperflexMapClusterIdToStSnapshotPoint {
	this := HyperflexMapClusterIdToStSnapshotPoint{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexMapClusterIdToStSnapshotPointWithDefaults instantiates a new HyperflexMapClusterIdToStSnapshotPoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexMapClusterIdToStSnapshotPointWithDefaults() *HyperflexMapClusterIdToStSnapshotPoint {
	this := HyperflexMapClusterIdToStSnapshotPoint{}
	var classId string = "hyperflex.MapClusterIdToStSnapshotPoint"
	this.ClassId = classId
	var objectType string = "hyperflex.MapClusterIdToStSnapshotPoint"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexMapClusterIdToStSnapshotPoint) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexMapClusterIdToStSnapshotPoint) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexMapClusterIdToStSnapshotPoint) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexMapClusterIdToStSnapshotPoint) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexMapClusterIdToStSnapshotPoint) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexMapClusterIdToStSnapshotPoint) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *HyperflexMapClusterIdToStSnapshotPoint) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexMapClusterIdToStSnapshotPoint) GetClusterIdOk() (*string, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *HyperflexMapClusterIdToStSnapshotPoint) HasClusterId() bool {
	if o != nil && o.ClusterId != nil {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *HyperflexMapClusterIdToStSnapshotPoint) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetSnapshotPoint returns the SnapshotPoint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexMapClusterIdToStSnapshotPoint) GetSnapshotPoint() HyperflexSnapshotPoint {
	if o == nil || o.SnapshotPoint.Get() == nil {
		var ret HyperflexSnapshotPoint
		return ret
	}
	return *o.SnapshotPoint.Get()
}

// GetSnapshotPointOk returns a tuple with the SnapshotPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexMapClusterIdToStSnapshotPoint) GetSnapshotPointOk() (*HyperflexSnapshotPoint, bool) {
	if o == nil {
		return nil, false
	}
	return o.SnapshotPoint.Get(), o.SnapshotPoint.IsSet()
}

// HasSnapshotPoint returns a boolean if a field has been set.
func (o *HyperflexMapClusterIdToStSnapshotPoint) HasSnapshotPoint() bool {
	if o != nil && o.SnapshotPoint.IsSet() {
		return true
	}

	return false
}

// SetSnapshotPoint gets a reference to the given NullableHyperflexSnapshotPoint and assigns it to the SnapshotPoint field.
func (o *HyperflexMapClusterIdToStSnapshotPoint) SetSnapshotPoint(v HyperflexSnapshotPoint) {
	o.SnapshotPoint.Set(&v)
}

// SetSnapshotPointNil sets the value for SnapshotPoint to be an explicit nil
func (o *HyperflexMapClusterIdToStSnapshotPoint) SetSnapshotPointNil() {
	o.SnapshotPoint.Set(nil)
}

// UnsetSnapshotPoint ensures that no value is present for SnapshotPoint, not even an explicit nil
func (o *HyperflexMapClusterIdToStSnapshotPoint) UnsetSnapshotPoint() {
	o.SnapshotPoint.Unset()
}

func (o HyperflexMapClusterIdToStSnapshotPoint) MarshalJSON() ([]byte, error) {
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
	if o.ClusterId != nil {
		toSerialize["ClusterId"] = o.ClusterId
	}
	if o.SnapshotPoint.IsSet() {
		toSerialize["SnapshotPoint"] = o.SnapshotPoint.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexMapClusterIdToStSnapshotPoint) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexMapClusterIdToStSnapshotPointWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// ClusterId of the snapshot point.
		ClusterId     *string                        `json:"ClusterId,omitempty"`
		SnapshotPoint NullableHyperflexSnapshotPoint `json:"SnapshotPoint,omitempty"`
	}

	varHyperflexMapClusterIdToStSnapshotPointWithoutEmbeddedStruct := HyperflexMapClusterIdToStSnapshotPointWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexMapClusterIdToStSnapshotPointWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexMapClusterIdToStSnapshotPoint := _HyperflexMapClusterIdToStSnapshotPoint{}
		varHyperflexMapClusterIdToStSnapshotPoint.ClassId = varHyperflexMapClusterIdToStSnapshotPointWithoutEmbeddedStruct.ClassId
		varHyperflexMapClusterIdToStSnapshotPoint.ObjectType = varHyperflexMapClusterIdToStSnapshotPointWithoutEmbeddedStruct.ObjectType
		varHyperflexMapClusterIdToStSnapshotPoint.ClusterId = varHyperflexMapClusterIdToStSnapshotPointWithoutEmbeddedStruct.ClusterId
		varHyperflexMapClusterIdToStSnapshotPoint.SnapshotPoint = varHyperflexMapClusterIdToStSnapshotPointWithoutEmbeddedStruct.SnapshotPoint
		*o = HyperflexMapClusterIdToStSnapshotPoint(varHyperflexMapClusterIdToStSnapshotPoint)
	} else {
		return err
	}

	varHyperflexMapClusterIdToStSnapshotPoint := _HyperflexMapClusterIdToStSnapshotPoint{}

	err = json.Unmarshal(bytes, &varHyperflexMapClusterIdToStSnapshotPoint)
	if err == nil {
		o.MoBaseComplexType = varHyperflexMapClusterIdToStSnapshotPoint.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClusterId")
		delete(additionalProperties, "SnapshotPoint")

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

type NullableHyperflexMapClusterIdToStSnapshotPoint struct {
	value *HyperflexMapClusterIdToStSnapshotPoint
	isSet bool
}

func (v NullableHyperflexMapClusterIdToStSnapshotPoint) Get() *HyperflexMapClusterIdToStSnapshotPoint {
	return v.value
}

func (v *NullableHyperflexMapClusterIdToStSnapshotPoint) Set(val *HyperflexMapClusterIdToStSnapshotPoint) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexMapClusterIdToStSnapshotPoint) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexMapClusterIdToStSnapshotPoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexMapClusterIdToStSnapshotPoint(val *HyperflexMapClusterIdToStSnapshotPoint) *NullableHyperflexMapClusterIdToStSnapshotPoint {
	return &NullableHyperflexMapClusterIdToStSnapshotPoint{value: val, isSet: true}
}

func (v NullableHyperflexMapClusterIdToStSnapshotPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexMapClusterIdToStSnapshotPoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
