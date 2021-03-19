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

// ConfigImporter All import operations are captured as Importer instances. Users shall use this Importer mo to track the import operation progress.
type ConfigImporter struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The path to the archive in Intersight storage that has all the MOs to be imported.
	ImportPath *string `json:"ImportPath,omitempty"`
	// The source of the archive in Intersight storage that has all the MOs to be imported. * `ImageRepo` - The 'ImageRepo' source if the source of exporter archive is image repository. * `URL` - The 'URL' source if the source of exported archive is a URL.
	ImportSource *string `json:"ImportSource,omitempty"`
	// An identifier for the importer instance.
	Name *string `json:"Name,omitempty"`
	// Specifies whether integrity checks must be skipped.
	SkipIntegrityChecks *bool `json:"SkipIntegrityChecks,omitempty"`
	// Status of the import operation. * `` - The operation has not started. * `InProgress` - The operation is in progress. * `Success` - The operation has succeeded. * `Failed` - The operation has failed. * `RollBackInitiated` - The rollback has been inititiated for import failure. * `RollBackFailed` - The rollback has failed for import failure. * `RollbackCompleted` - The rollback has completed for import failure. * `RollbackAborted` - The rollback has been aborted for import failure. * `OperationTimedOut` - The operation has timed out.
	Status *string `json:"Status,omitempty"`
	// Status message associated with failures or progress indication.
	StatusMessage *string `json:"StatusMessage,omitempty"`
	// An array of relationships to configImportedItem resources.
	ImportedItems        []ConfigImportedItemRelationship      `json:"ImportedItems,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConfigImporter ConfigImporter

// NewConfigImporter instantiates a new ConfigImporter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigImporter(classId string, objectType string) *ConfigImporter {
	this := ConfigImporter{}
	this.ClassId = classId
	this.ObjectType = objectType
	var importSource string = "ImageRepo"
	this.ImportSource = &importSource
	var skipIntegrityChecks bool = false
	this.SkipIntegrityChecks = &skipIntegrityChecks
	var status string = ""
	this.Status = &status
	return &this
}

// NewConfigImporterWithDefaults instantiates a new ConfigImporter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigImporterWithDefaults() *ConfigImporter {
	this := ConfigImporter{}
	var classId string = "config.Importer"
	this.ClassId = classId
	var objectType string = "config.Importer"
	this.ObjectType = objectType
	var importSource string = "ImageRepo"
	this.ImportSource = &importSource
	var skipIntegrityChecks bool = false
	this.SkipIntegrityChecks = &skipIntegrityChecks
	var status string = ""
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConfigImporter) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConfigImporter) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConfigImporter) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConfigImporter) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConfigImporter) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConfigImporter) SetObjectType(v string) {
	o.ObjectType = v
}

// GetImportPath returns the ImportPath field value if set, zero value otherwise.
func (o *ConfigImporter) GetImportPath() string {
	if o == nil || o.ImportPath == nil {
		var ret string
		return ret
	}
	return *o.ImportPath
}

// GetImportPathOk returns a tuple with the ImportPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigImporter) GetImportPathOk() (*string, bool) {
	if o == nil || o.ImportPath == nil {
		return nil, false
	}
	return o.ImportPath, true
}

// HasImportPath returns a boolean if a field has been set.
func (o *ConfigImporter) HasImportPath() bool {
	if o != nil && o.ImportPath != nil {
		return true
	}

	return false
}

// SetImportPath gets a reference to the given string and assigns it to the ImportPath field.
func (o *ConfigImporter) SetImportPath(v string) {
	o.ImportPath = &v
}

// GetImportSource returns the ImportSource field value if set, zero value otherwise.
func (o *ConfigImporter) GetImportSource() string {
	if o == nil || o.ImportSource == nil {
		var ret string
		return ret
	}
	return *o.ImportSource
}

// GetImportSourceOk returns a tuple with the ImportSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigImporter) GetImportSourceOk() (*string, bool) {
	if o == nil || o.ImportSource == nil {
		return nil, false
	}
	return o.ImportSource, true
}

// HasImportSource returns a boolean if a field has been set.
func (o *ConfigImporter) HasImportSource() bool {
	if o != nil && o.ImportSource != nil {
		return true
	}

	return false
}

// SetImportSource gets a reference to the given string and assigns it to the ImportSource field.
func (o *ConfigImporter) SetImportSource(v string) {
	o.ImportSource = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConfigImporter) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigImporter) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConfigImporter) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConfigImporter) SetName(v string) {
	o.Name = &v
}

// GetSkipIntegrityChecks returns the SkipIntegrityChecks field value if set, zero value otherwise.
func (o *ConfigImporter) GetSkipIntegrityChecks() bool {
	if o == nil || o.SkipIntegrityChecks == nil {
		var ret bool
		return ret
	}
	return *o.SkipIntegrityChecks
}

// GetSkipIntegrityChecksOk returns a tuple with the SkipIntegrityChecks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigImporter) GetSkipIntegrityChecksOk() (*bool, bool) {
	if o == nil || o.SkipIntegrityChecks == nil {
		return nil, false
	}
	return o.SkipIntegrityChecks, true
}

// HasSkipIntegrityChecks returns a boolean if a field has been set.
func (o *ConfigImporter) HasSkipIntegrityChecks() bool {
	if o != nil && o.SkipIntegrityChecks != nil {
		return true
	}

	return false
}

// SetSkipIntegrityChecks gets a reference to the given bool and assigns it to the SkipIntegrityChecks field.
func (o *ConfigImporter) SetSkipIntegrityChecks(v bool) {
	o.SkipIntegrityChecks = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ConfigImporter) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigImporter) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ConfigImporter) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ConfigImporter) SetStatus(v string) {
	o.Status = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *ConfigImporter) GetStatusMessage() string {
	if o == nil || o.StatusMessage == nil {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigImporter) GetStatusMessageOk() (*string, bool) {
	if o == nil || o.StatusMessage == nil {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *ConfigImporter) HasStatusMessage() bool {
	if o != nil && o.StatusMessage != nil {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *ConfigImporter) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetImportedItems returns the ImportedItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigImporter) GetImportedItems() []ConfigImportedItemRelationship {
	if o == nil {
		var ret []ConfigImportedItemRelationship
		return ret
	}
	return o.ImportedItems
}

// GetImportedItemsOk returns a tuple with the ImportedItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigImporter) GetImportedItemsOk() (*[]ConfigImportedItemRelationship, bool) {
	if o == nil || o.ImportedItems == nil {
		return nil, false
	}
	return &o.ImportedItems, true
}

// HasImportedItems returns a boolean if a field has been set.
func (o *ConfigImporter) HasImportedItems() bool {
	if o != nil && o.ImportedItems != nil {
		return true
	}

	return false
}

// SetImportedItems gets a reference to the given []ConfigImportedItemRelationship and assigns it to the ImportedItems field.
func (o *ConfigImporter) SetImportedItems(v []ConfigImportedItemRelationship) {
	o.ImportedItems = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ConfigImporter) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigImporter) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ConfigImporter) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *ConfigImporter) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o ConfigImporter) MarshalJSON() ([]byte, error) {
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
	if o.ImportPath != nil {
		toSerialize["ImportPath"] = o.ImportPath
	}
	if o.ImportSource != nil {
		toSerialize["ImportSource"] = o.ImportSource
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.SkipIntegrityChecks != nil {
		toSerialize["SkipIntegrityChecks"] = o.SkipIntegrityChecks
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.StatusMessage != nil {
		toSerialize["StatusMessage"] = o.StatusMessage
	}
	if o.ImportedItems != nil {
		toSerialize["ImportedItems"] = o.ImportedItems
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConfigImporter) UnmarshalJSON(bytes []byte) (err error) {
	type ConfigImporterWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The path to the archive in Intersight storage that has all the MOs to be imported.
		ImportPath *string `json:"ImportPath,omitempty"`
		// The source of the archive in Intersight storage that has all the MOs to be imported. * `ImageRepo` - The 'ImageRepo' source if the source of exporter archive is image repository. * `URL` - The 'URL' source if the source of exported archive is a URL.
		ImportSource *string `json:"ImportSource,omitempty"`
		// An identifier for the importer instance.
		Name *string `json:"Name,omitempty"`
		// Specifies whether integrity checks must be skipped.
		SkipIntegrityChecks *bool `json:"SkipIntegrityChecks,omitempty"`
		// Status of the import operation. * `` - The operation has not started. * `InProgress` - The operation is in progress. * `Success` - The operation has succeeded. * `Failed` - The operation has failed. * `RollBackInitiated` - The rollback has been inititiated for import failure. * `RollBackFailed` - The rollback has failed for import failure. * `RollbackCompleted` - The rollback has completed for import failure. * `RollbackAborted` - The rollback has been aborted for import failure. * `OperationTimedOut` - The operation has timed out.
		Status *string `json:"Status,omitempty"`
		// Status message associated with failures or progress indication.
		StatusMessage *string `json:"StatusMessage,omitempty"`
		// An array of relationships to configImportedItem resources.
		ImportedItems []ConfigImportedItemRelationship      `json:"ImportedItems,omitempty"`
		Organization  *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varConfigImporterWithoutEmbeddedStruct := ConfigImporterWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varConfigImporterWithoutEmbeddedStruct)
	if err == nil {
		varConfigImporter := _ConfigImporter{}
		varConfigImporter.ClassId = varConfigImporterWithoutEmbeddedStruct.ClassId
		varConfigImporter.ObjectType = varConfigImporterWithoutEmbeddedStruct.ObjectType
		varConfigImporter.ImportPath = varConfigImporterWithoutEmbeddedStruct.ImportPath
		varConfigImporter.ImportSource = varConfigImporterWithoutEmbeddedStruct.ImportSource
		varConfigImporter.Name = varConfigImporterWithoutEmbeddedStruct.Name
		varConfigImporter.SkipIntegrityChecks = varConfigImporterWithoutEmbeddedStruct.SkipIntegrityChecks
		varConfigImporter.Status = varConfigImporterWithoutEmbeddedStruct.Status
		varConfigImporter.StatusMessage = varConfigImporterWithoutEmbeddedStruct.StatusMessage
		varConfigImporter.ImportedItems = varConfigImporterWithoutEmbeddedStruct.ImportedItems
		varConfigImporter.Organization = varConfigImporterWithoutEmbeddedStruct.Organization
		*o = ConfigImporter(varConfigImporter)
	} else {
		return err
	}

	varConfigImporter := _ConfigImporter{}

	err = json.Unmarshal(bytes, &varConfigImporter)
	if err == nil {
		o.MoBaseMo = varConfigImporter.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ImportPath")
		delete(additionalProperties, "ImportSource")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "SkipIntegrityChecks")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "StatusMessage")
		delete(additionalProperties, "ImportedItems")
		delete(additionalProperties, "Organization")

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

type NullableConfigImporter struct {
	value *ConfigImporter
	isSet bool
}

func (v NullableConfigImporter) Get() *ConfigImporter {
	return v.value
}

func (v *NullableConfigImporter) Set(val *ConfigImporter) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigImporter) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigImporter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigImporter(val *ConfigImporter) *NullableConfigImporter {
	return &NullableConfigImporter{value: val, isSet: true}
}

func (v NullableConfigImporter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigImporter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
