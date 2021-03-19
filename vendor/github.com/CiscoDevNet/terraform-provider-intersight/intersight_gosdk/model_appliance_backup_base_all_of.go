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
)

// ApplianceBackupBaseAllOf Definition of the list of properties defined in 'appliance.BackupBase', excluding properties defined in parent classes.
type ApplianceBackupBaseAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Backup filename to backup or restore.
	Filename *string `json:"Filename,omitempty"`
	// Communication protocol used by the file server (e.g. scp or sftp). * `scp` - Secure Copy Protocol (SCP) to access the file server. * `sftp` - SSH File Transfer Protocol (SFTP) to access file server.
	Protocol *string `json:"Protocol,omitempty"`
	// Hostname of the remote file server.
	RemoteHost *string `json:"RemoteHost,omitempty"`
	// File server directory to copy the file.
	RemotePath *string `json:"RemotePath,omitempty"`
	// Remote TCP port on the file server (e.g. 22 for scp).
	RemotePort *int64 `json:"RemotePort,omitempty"`
	// Username to authenticate the fileserver.
	Username             *string `json:"Username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceBackupBaseAllOf ApplianceBackupBaseAllOf

// NewApplianceBackupBaseAllOf instantiates a new ApplianceBackupBaseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceBackupBaseAllOf(classId string, objectType string) *ApplianceBackupBaseAllOf {
	this := ApplianceBackupBaseAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var protocol string = "scp"
	this.Protocol = &protocol
	return &this
}

// NewApplianceBackupBaseAllOfWithDefaults instantiates a new ApplianceBackupBaseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceBackupBaseAllOfWithDefaults() *ApplianceBackupBaseAllOf {
	this := ApplianceBackupBaseAllOf{}
	var protocol string = "scp"
	this.Protocol = &protocol
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceBackupBaseAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceBackupBaseAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceBackupBaseAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceBackupBaseAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceBackupBaseAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceBackupBaseAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *ApplianceBackupBaseAllOf) GetFilename() string {
	if o == nil || o.Filename == nil {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceBackupBaseAllOf) GetFilenameOk() (*string, bool) {
	if o == nil || o.Filename == nil {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *ApplianceBackupBaseAllOf) HasFilename() bool {
	if o != nil && o.Filename != nil {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *ApplianceBackupBaseAllOf) SetFilename(v string) {
	o.Filename = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *ApplianceBackupBaseAllOf) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceBackupBaseAllOf) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *ApplianceBackupBaseAllOf) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *ApplianceBackupBaseAllOf) SetProtocol(v string) {
	o.Protocol = &v
}

// GetRemoteHost returns the RemoteHost field value if set, zero value otherwise.
func (o *ApplianceBackupBaseAllOf) GetRemoteHost() string {
	if o == nil || o.RemoteHost == nil {
		var ret string
		return ret
	}
	return *o.RemoteHost
}

// GetRemoteHostOk returns a tuple with the RemoteHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceBackupBaseAllOf) GetRemoteHostOk() (*string, bool) {
	if o == nil || o.RemoteHost == nil {
		return nil, false
	}
	return o.RemoteHost, true
}

// HasRemoteHost returns a boolean if a field has been set.
func (o *ApplianceBackupBaseAllOf) HasRemoteHost() bool {
	if o != nil && o.RemoteHost != nil {
		return true
	}

	return false
}

// SetRemoteHost gets a reference to the given string and assigns it to the RemoteHost field.
func (o *ApplianceBackupBaseAllOf) SetRemoteHost(v string) {
	o.RemoteHost = &v
}

// GetRemotePath returns the RemotePath field value if set, zero value otherwise.
func (o *ApplianceBackupBaseAllOf) GetRemotePath() string {
	if o == nil || o.RemotePath == nil {
		var ret string
		return ret
	}
	return *o.RemotePath
}

// GetRemotePathOk returns a tuple with the RemotePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceBackupBaseAllOf) GetRemotePathOk() (*string, bool) {
	if o == nil || o.RemotePath == nil {
		return nil, false
	}
	return o.RemotePath, true
}

// HasRemotePath returns a boolean if a field has been set.
func (o *ApplianceBackupBaseAllOf) HasRemotePath() bool {
	if o != nil && o.RemotePath != nil {
		return true
	}

	return false
}

// SetRemotePath gets a reference to the given string and assigns it to the RemotePath field.
func (o *ApplianceBackupBaseAllOf) SetRemotePath(v string) {
	o.RemotePath = &v
}

// GetRemotePort returns the RemotePort field value if set, zero value otherwise.
func (o *ApplianceBackupBaseAllOf) GetRemotePort() int64 {
	if o == nil || o.RemotePort == nil {
		var ret int64
		return ret
	}
	return *o.RemotePort
}

// GetRemotePortOk returns a tuple with the RemotePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceBackupBaseAllOf) GetRemotePortOk() (*int64, bool) {
	if o == nil || o.RemotePort == nil {
		return nil, false
	}
	return o.RemotePort, true
}

// HasRemotePort returns a boolean if a field has been set.
func (o *ApplianceBackupBaseAllOf) HasRemotePort() bool {
	if o != nil && o.RemotePort != nil {
		return true
	}

	return false
}

// SetRemotePort gets a reference to the given int64 and assigns it to the RemotePort field.
func (o *ApplianceBackupBaseAllOf) SetRemotePort(v int64) {
	o.RemotePort = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ApplianceBackupBaseAllOf) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceBackupBaseAllOf) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ApplianceBackupBaseAllOf) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ApplianceBackupBaseAllOf) SetUsername(v string) {
	o.Username = &v
}

func (o ApplianceBackupBaseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Filename != nil {
		toSerialize["Filename"] = o.Filename
	}
	if o.Protocol != nil {
		toSerialize["Protocol"] = o.Protocol
	}
	if o.RemoteHost != nil {
		toSerialize["RemoteHost"] = o.RemoteHost
	}
	if o.RemotePath != nil {
		toSerialize["RemotePath"] = o.RemotePath
	}
	if o.RemotePort != nil {
		toSerialize["RemotePort"] = o.RemotePort
	}
	if o.Username != nil {
		toSerialize["Username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceBackupBaseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varApplianceBackupBaseAllOf := _ApplianceBackupBaseAllOf{}

	if err = json.Unmarshal(bytes, &varApplianceBackupBaseAllOf); err == nil {
		*o = ApplianceBackupBaseAllOf(varApplianceBackupBaseAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Filename")
		delete(additionalProperties, "Protocol")
		delete(additionalProperties, "RemoteHost")
		delete(additionalProperties, "RemotePath")
		delete(additionalProperties, "RemotePort")
		delete(additionalProperties, "Username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplianceBackupBaseAllOf struct {
	value *ApplianceBackupBaseAllOf
	isSet bool
}

func (v NullableApplianceBackupBaseAllOf) Get() *ApplianceBackupBaseAllOf {
	return v.value
}

func (v *NullableApplianceBackupBaseAllOf) Set(val *ApplianceBackupBaseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceBackupBaseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceBackupBaseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceBackupBaseAllOf(val *ApplianceBackupBaseAllOf) *NullableApplianceBackupBaseAllOf {
	return &NullableApplianceBackupBaseAllOf{value: val, isSet: true}
}

func (v NullableApplianceBackupBaseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceBackupBaseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
