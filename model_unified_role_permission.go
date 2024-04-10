/*
Libre Graph API

Libre Graph is a free API for cloud collaboration inspired by the MS Graph API.

API version: v1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"encoding/json"
)

// checks if the UnifiedRolePermission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnifiedRolePermission{}

// UnifiedRolePermission Represents a collection of allowed resource actions and the conditions that must be met for the action to be allowed. Resource actions are tasks that can be performed on a resource. For example, an application resource may support create, update, delete, and reset password actions.
type UnifiedRolePermission struct {
	// Set of tasks that can be performed on a resource. Required.  The following is the schema for resource actions:  ```    {Namespace}/{Entity}/{PropertySet}/{Action} ```   For example: `libre.graph/applications/credentials/update`   * *{Namespace}* - The services that exposes the task. For example, all tasks in libre graph use the namespace `libre.graph`.  * *{Entity}* - The logical features or components exposed by the service in libre graph. For example, `applications`, `servicePrincipals`, or `groups`.  * *{PropertySet}* - Optional. The specific properties or aspects of the entity for which access is being granted.    For example, `libre.graph/applications/authentication/read` grants the ability to read the reply URL, logout URL,    and implicit flow property on the **application** object in libre graph. The following are reserved names for common property sets:    * `allProperties` - Designates all properties of the entity, including privileged properties.      Examples include `libre.graph/applications/allProperties/read` and `libre.graph/applications/allProperties/update`.    * `basic` - Designates common read properties but excludes privileged ones.      For example, `libre.graph/applications/basic/update` includes the ability to update standard properties like display name.    * `standard` - Designates common update properties but excludes privileged ones.      For example, `libre.graph/applications/standard/read`.  * *{Actions}* - The operations being granted. In most circumstances, permissions should be expressed in terms of CRUD operations or allTasks. Actions include:    * `create` - The ability to create a new instance of the entity.    * `read` - The ability to read a given property set (including allProperties).    * `update` - The ability to update a given property set (including allProperties).    * `delete` - The ability to delete a given entity.    * `allTasks` - Represents all CRUD operations (create, read, update, and delete).   Following the CS3 API we can represent the CS3 permissions by mapping them to driveItem properties or relations like this:  | [CS3 ResourcePermission](https://cs3org.github.io/cs3apis/#cs3.storage.provider.v1beta1.ResourcePermissions) | action | comment |  | ------------------------------------------------------------------------------------------------------------ | ------ | ------- |  | `stat` | `libre.graph/driveItem/basic/read` | `basic` because it does not include versions or trashed items |  | `get_quota` | `libre.graph/driveItem/quota/read` | read only the `quota` property |  | `get_path` | `libre.graph/driveItem/path/read` | read only the `path` property |  | `move` | `libre.graph/driveItem/path/update` | allows updating the `path` property of a CS3 resource |  | `delete` | `libre.graph/driveItem/standard/delete` | `standard` because deleting is a common update operation |  | `list_container` | `libre.graph/driveItem/children/read` | |  | `create_container` | `libre.graph/driveItem/children/create` | |  | `initiate_file_download` | `libre.graph/driveItem/content/read` | `content` is the property read when initiating a download |  | `initiate_file_upload` | `libre.graph/driveItem/upload/create` | `uploads` are a separate property. postprocessing creates the `content` |  | `add_grant` | `libre.graph/driveItem/permissions/create` | |  | `list_grant` | `libre.graph/driveItem/permissions/read` | |  | `update_grant` | `libre.graph/driveItem/permissions/update` | |  | `remove_grant` | `libre.graph/driveItem/permissions/delete` | |  | `deny_grant` | `libre.graph/driveItem/permissions/deny` | uses a non CRUD action `deny` |  | `list_file_versions` | `libre.graph/driveItem/versions/read` | `versions` is a `driveItemVersion` collection |  | `restore_file_version` | `libre.graph/driveItem/versions/update` | the only `update` action is restore |  | `list_recycle` | `libre.graph/driveItem/deleted/read` | reading a driveItem `deleted` property implies listing |  | `restore_recycle_item` | `libre.graph/driveItem/deleted/update` | the only `update` action is restore |  | `purge_recycle` | `libre.graph/driveItem/deleted/delete` | allows purging deleted `driveItems` |   Managing drives would be a different entity. A space manager role could be written as `libre.graph/drive/permission/allTasks`.
	AllowedResourceActions []string `json:"allowedResourceActions,omitempty"`
	// Optional constraints that must be met for the permission to be effective. Not supported for custom roles.  Conditions define constraints that must be met. For example, a requirement that target resource must have a certain property. The following are the supported conditions:  * Drive: `exists @Resource.Drive` - The target resource must be a drive/space * Folder: `exists @Resource.Folder` - The target resource must be a folder * File: `exists @Resource.File` - The target resource must be a file  The following is an example of a role permission with a condition that the target resource is a folder: ```json   \"rolePermissions\": [       {           \"allowedResourceActions\": [               \"libre.graph/applications/basic/update\",               \"libre.graph/applications/credentials/update\"           ],           \"condition\":  \"exists @Resource.File\"       }   ] ``` Conditions aren't supported for custom roles.
	Condition *string `json:"condition,omitempty"`
}

// NewUnifiedRolePermission instantiates a new UnifiedRolePermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnifiedRolePermission() *UnifiedRolePermission {
	this := UnifiedRolePermission{}
	return &this
}

// NewUnifiedRolePermissionWithDefaults instantiates a new UnifiedRolePermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnifiedRolePermissionWithDefaults() *UnifiedRolePermission {
	this := UnifiedRolePermission{}
	return &this
}

// GetAllowedResourceActions returns the AllowedResourceActions field value if set, zero value otherwise.
func (o *UnifiedRolePermission) GetAllowedResourceActions() []string {
	if o == nil || IsNil(o.AllowedResourceActions) {
		var ret []string
		return ret
	}
	return o.AllowedResourceActions
}

// GetAllowedResourceActionsOk returns a tuple with the AllowedResourceActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnifiedRolePermission) GetAllowedResourceActionsOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedResourceActions) {
		return nil, false
	}
	return o.AllowedResourceActions, true
}

// HasAllowedResourceActions returns a boolean if a field has been set.
func (o *UnifiedRolePermission) HasAllowedResourceActions() bool {
	if o != nil && !IsNil(o.AllowedResourceActions) {
		return true
	}

	return false
}

// SetAllowedResourceActions gets a reference to the given []string and assigns it to the AllowedResourceActions field.
func (o *UnifiedRolePermission) SetAllowedResourceActions(v []string) {
	o.AllowedResourceActions = v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *UnifiedRolePermission) GetCondition() string {
	if o == nil || IsNil(o.Condition) {
		var ret string
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnifiedRolePermission) GetConditionOk() (*string, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *UnifiedRolePermission) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given string and assigns it to the Condition field.
func (o *UnifiedRolePermission) SetCondition(v string) {
	o.Condition = &v
}

func (o UnifiedRolePermission) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnifiedRolePermission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowedResourceActions) {
		toSerialize["allowedResourceActions"] = o.AllowedResourceActions
	}
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	return toSerialize, nil
}

type NullableUnifiedRolePermission struct {
	value *UnifiedRolePermission
	isSet bool
}

func (v NullableUnifiedRolePermission) Get() *UnifiedRolePermission {
	return v.value
}

func (v *NullableUnifiedRolePermission) Set(val *UnifiedRolePermission) {
	v.value = val
	v.isSet = true
}

func (v NullableUnifiedRolePermission) IsSet() bool {
	return v.isSet
}

func (v *NullableUnifiedRolePermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnifiedRolePermission(val *UnifiedRolePermission) *NullableUnifiedRolePermission {
	return &NullableUnifiedRolePermission{value: val, isSet: true}
}

func (v NullableUnifiedRolePermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnifiedRolePermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
