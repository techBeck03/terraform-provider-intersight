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
	"time"
)

// NiaapiRevisionInfoAllOf Definition of the list of properties defined in 'niaapi.RevisionInfo', excluding properties defined in parent classes.
type NiaapiRevisionInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The date the revision is made.
	DatePublished *time.Time `json:"DatePublished,omitempty"`
	// The changes made in this revision.
	RevisionComment *string `json:"RevisionComment,omitempty"`
	// The Revision No. of this revision.
	RevisionNo           *string `json:"RevisionNo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiaapiRevisionInfoAllOf NiaapiRevisionInfoAllOf

// NewNiaapiRevisionInfoAllOf instantiates a new NiaapiRevisionInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiRevisionInfoAllOf(classId string, objectType string) *NiaapiRevisionInfoAllOf {
	this := NiaapiRevisionInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiaapiRevisionInfoAllOfWithDefaults instantiates a new NiaapiRevisionInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiRevisionInfoAllOfWithDefaults() *NiaapiRevisionInfoAllOf {
	this := NiaapiRevisionInfoAllOf{}
	var classId string = "niaapi.RevisionInfo"
	this.ClassId = classId
	var objectType string = "niaapi.RevisionInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiaapiRevisionInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiaapiRevisionInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiaapiRevisionInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiaapiRevisionInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiaapiRevisionInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiaapiRevisionInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDatePublished returns the DatePublished field value if set, zero value otherwise.
func (o *NiaapiRevisionInfoAllOf) GetDatePublished() time.Time {
	if o == nil || o.DatePublished == nil {
		var ret time.Time
		return ret
	}
	return *o.DatePublished
}

// GetDatePublishedOk returns a tuple with the DatePublished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiRevisionInfoAllOf) GetDatePublishedOk() (*time.Time, bool) {
	if o == nil || o.DatePublished == nil {
		return nil, false
	}
	return o.DatePublished, true
}

// HasDatePublished returns a boolean if a field has been set.
func (o *NiaapiRevisionInfoAllOf) HasDatePublished() bool {
	if o != nil && o.DatePublished != nil {
		return true
	}

	return false
}

// SetDatePublished gets a reference to the given time.Time and assigns it to the DatePublished field.
func (o *NiaapiRevisionInfoAllOf) SetDatePublished(v time.Time) {
	o.DatePublished = &v
}

// GetRevisionComment returns the RevisionComment field value if set, zero value otherwise.
func (o *NiaapiRevisionInfoAllOf) GetRevisionComment() string {
	if o == nil || o.RevisionComment == nil {
		var ret string
		return ret
	}
	return *o.RevisionComment
}

// GetRevisionCommentOk returns a tuple with the RevisionComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiRevisionInfoAllOf) GetRevisionCommentOk() (*string, bool) {
	if o == nil || o.RevisionComment == nil {
		return nil, false
	}
	return o.RevisionComment, true
}

// HasRevisionComment returns a boolean if a field has been set.
func (o *NiaapiRevisionInfoAllOf) HasRevisionComment() bool {
	if o != nil && o.RevisionComment != nil {
		return true
	}

	return false
}

// SetRevisionComment gets a reference to the given string and assigns it to the RevisionComment field.
func (o *NiaapiRevisionInfoAllOf) SetRevisionComment(v string) {
	o.RevisionComment = &v
}

// GetRevisionNo returns the RevisionNo field value if set, zero value otherwise.
func (o *NiaapiRevisionInfoAllOf) GetRevisionNo() string {
	if o == nil || o.RevisionNo == nil {
		var ret string
		return ret
	}
	return *o.RevisionNo
}

// GetRevisionNoOk returns a tuple with the RevisionNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiRevisionInfoAllOf) GetRevisionNoOk() (*string, bool) {
	if o == nil || o.RevisionNo == nil {
		return nil, false
	}
	return o.RevisionNo, true
}

// HasRevisionNo returns a boolean if a field has been set.
func (o *NiaapiRevisionInfoAllOf) HasRevisionNo() bool {
	if o != nil && o.RevisionNo != nil {
		return true
	}

	return false
}

// SetRevisionNo gets a reference to the given string and assigns it to the RevisionNo field.
func (o *NiaapiRevisionInfoAllOf) SetRevisionNo(v string) {
	o.RevisionNo = &v
}

func (o NiaapiRevisionInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DatePublished != nil {
		toSerialize["DatePublished"] = o.DatePublished
	}
	if o.RevisionComment != nil {
		toSerialize["RevisionComment"] = o.RevisionComment
	}
	if o.RevisionNo != nil {
		toSerialize["RevisionNo"] = o.RevisionNo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiaapiRevisionInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiaapiRevisionInfoAllOf := _NiaapiRevisionInfoAllOf{}

	if err = json.Unmarshal(bytes, &varNiaapiRevisionInfoAllOf); err == nil {
		*o = NiaapiRevisionInfoAllOf(varNiaapiRevisionInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DatePublished")
		delete(additionalProperties, "RevisionComment")
		delete(additionalProperties, "RevisionNo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiaapiRevisionInfoAllOf struct {
	value *NiaapiRevisionInfoAllOf
	isSet bool
}

func (v NullableNiaapiRevisionInfoAllOf) Get() *NiaapiRevisionInfoAllOf {
	return v.value
}

func (v *NullableNiaapiRevisionInfoAllOf) Set(val *NiaapiRevisionInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiRevisionInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiRevisionInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiRevisionInfoAllOf(val *NiaapiRevisionInfoAllOf) *NullableNiaapiRevisionInfoAllOf {
	return &NullableNiaapiRevisionInfoAllOf{value: val, isSet: true}
}

func (v NullableNiaapiRevisionInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiRevisionInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
