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
)

// FirmwareDistributableAllOf Definition of the list of properties defined in 'firmware.Distributable', excluding properties defined in parent classes.
type FirmwareDistributableAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The file location of the distributable.
	FileLocation *string `json:"FileLocation,omitempty"`
	// The category into which the distributable falls into according to the supported platform series. For e.g.; C-Series/B-Series/Infrastructure.
	ImageCategory *string `json:"ImageCategory,omitempty"`
	// The source of the distributable. If it has been created by the user or system. * `System` - The distributable has been created by the System. * `User` - The distributable has been created by the User.
	Origin               *string                                `json:"Origin,omitempty"`
	Catalog              *SoftwarerepositoryCatalogRelationship `json:"Catalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareDistributableAllOf FirmwareDistributableAllOf

// NewFirmwareDistributableAllOf instantiates a new FirmwareDistributableAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareDistributableAllOf(classId string, objectType string) *FirmwareDistributableAllOf {
	this := FirmwareDistributableAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var origin string = "System"
	this.Origin = &origin
	return &this
}

// NewFirmwareDistributableAllOfWithDefaults instantiates a new FirmwareDistributableAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareDistributableAllOfWithDefaults() *FirmwareDistributableAllOf {
	this := FirmwareDistributableAllOf{}
	var classId string = "firmware.Distributable"
	this.ClassId = classId
	var objectType string = "firmware.Distributable"
	this.ObjectType = objectType
	var origin string = "System"
	this.Origin = &origin
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareDistributableAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareDistributableAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareDistributableAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareDistributableAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareDistributableAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareDistributableAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFileLocation returns the FileLocation field value if set, zero value otherwise.
func (o *FirmwareDistributableAllOf) GetFileLocation() string {
	if o == nil || o.FileLocation == nil {
		var ret string
		return ret
	}
	return *o.FileLocation
}

// GetFileLocationOk returns a tuple with the FileLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareDistributableAllOf) GetFileLocationOk() (*string, bool) {
	if o == nil || o.FileLocation == nil {
		return nil, false
	}
	return o.FileLocation, true
}

// HasFileLocation returns a boolean if a field has been set.
func (o *FirmwareDistributableAllOf) HasFileLocation() bool {
	if o != nil && o.FileLocation != nil {
		return true
	}

	return false
}

// SetFileLocation gets a reference to the given string and assigns it to the FileLocation field.
func (o *FirmwareDistributableAllOf) SetFileLocation(v string) {
	o.FileLocation = &v
}

// GetImageCategory returns the ImageCategory field value if set, zero value otherwise.
func (o *FirmwareDistributableAllOf) GetImageCategory() string {
	if o == nil || o.ImageCategory == nil {
		var ret string
		return ret
	}
	return *o.ImageCategory
}

// GetImageCategoryOk returns a tuple with the ImageCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareDistributableAllOf) GetImageCategoryOk() (*string, bool) {
	if o == nil || o.ImageCategory == nil {
		return nil, false
	}
	return o.ImageCategory, true
}

// HasImageCategory returns a boolean if a field has been set.
func (o *FirmwareDistributableAllOf) HasImageCategory() bool {
	if o != nil && o.ImageCategory != nil {
		return true
	}

	return false
}

// SetImageCategory gets a reference to the given string and assigns it to the ImageCategory field.
func (o *FirmwareDistributableAllOf) SetImageCategory(v string) {
	o.ImageCategory = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *FirmwareDistributableAllOf) GetOrigin() string {
	if o == nil || o.Origin == nil {
		var ret string
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareDistributableAllOf) GetOriginOk() (*string, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *FirmwareDistributableAllOf) HasOrigin() bool {
	if o != nil && o.Origin != nil {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given string and assigns it to the Origin field.
func (o *FirmwareDistributableAllOf) SetOrigin(v string) {
	o.Origin = &v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *FirmwareDistributableAllOf) GetCatalog() SoftwarerepositoryCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret SoftwarerepositoryCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareDistributableAllOf) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *FirmwareDistributableAllOf) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given SoftwarerepositoryCatalogRelationship and assigns it to the Catalog field.
func (o *FirmwareDistributableAllOf) SetCatalog(v SoftwarerepositoryCatalogRelationship) {
	o.Catalog = &v
}

func (o FirmwareDistributableAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FileLocation != nil {
		toSerialize["FileLocation"] = o.FileLocation
	}
	if o.ImageCategory != nil {
		toSerialize["ImageCategory"] = o.ImageCategory
	}
	if o.Origin != nil {
		toSerialize["Origin"] = o.Origin
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareDistributableAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFirmwareDistributableAllOf := _FirmwareDistributableAllOf{}

	if err = json.Unmarshal(bytes, &varFirmwareDistributableAllOf); err == nil {
		*o = FirmwareDistributableAllOf(varFirmwareDistributableAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FileLocation")
		delete(additionalProperties, "ImageCategory")
		delete(additionalProperties, "Origin")
		delete(additionalProperties, "Catalog")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFirmwareDistributableAllOf struct {
	value *FirmwareDistributableAllOf
	isSet bool
}

func (v NullableFirmwareDistributableAllOf) Get() *FirmwareDistributableAllOf {
	return v.value
}

func (v *NullableFirmwareDistributableAllOf) Set(val *FirmwareDistributableAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareDistributableAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareDistributableAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareDistributableAllOf(val *FirmwareDistributableAllOf) *NullableFirmwareDistributableAllOf {
	return &NullableFirmwareDistributableAllOf{value: val, isSet: true}
}

func (v NullableFirmwareDistributableAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareDistributableAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
