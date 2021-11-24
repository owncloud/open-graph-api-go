/*
Open Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package opengraph

import (
	"encoding/json"
)

// DriveItem Reprensents a resource inside a drive. Read-only.
type DriveItem struct {
	BaseItem
	// The content stream, if the item represents a file.
	Content *string `json:"content,omitempty"`
	// An eTag for the content of the item. This eTag is not changed if only the metadata is changed. Note This property is not returned if the item is a folder. Read-only.
	CTag *string `json:"cTag,omitempty"`
	Deleted *Deleted `json:"deleted,omitempty"`
	File *OpenGraphFile `json:"file,omitempty"`
	FileSystemInfo *FileSystemInfo `json:"fileSystemInfo,omitempty"`
	Folder *Folder `json:"folder,omitempty"`
	Image *Image `json:"image,omitempty"`
	// If this property is non-null, it indicates that the driveItem is the top-most driveItem in the drive.
	Root *map[string]interface{} `json:"root,omitempty"`
	// Size of the item in bytes. Read-only.
	Size *int64 `json:"size,omitempty"`
	// WebDAV compatible URL for the item. Read-only.
	WebDavUrl *string `json:"webDavUrl,omitempty"`
	// Collection containing Item objects for the immediate children of Item. Only items representing folders have children. Read-only. Nullable.
	Children *[]DriveItem `json:"children,omitempty"`
}

// NewDriveItem instantiates a new DriveItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDriveItem() *DriveItem {
	this := DriveItem{}
	return &this
}

// NewDriveItemWithDefaults instantiates a new DriveItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDriveItemWithDefaults() *DriveItem {
	this := DriveItem{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *DriveItem) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveItem) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *DriveItem) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *DriveItem) SetContent(v string) {
	o.Content = &v
}

// GetCTag returns the CTag field value if set, zero value otherwise.
func (o *DriveItem) GetCTag() string {
	if o == nil || o.CTag == nil {
		var ret string
		return ret
	}
	return *o.CTag
}

// GetCTagOk returns a tuple with the CTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveItem) GetCTagOk() (*string, bool) {
	if o == nil || o.CTag == nil {
		return nil, false
	}
	return o.CTag, true
}

// HasCTag returns a boolean if a field has been set.
func (o *DriveItem) HasCTag() bool {
	if o != nil && o.CTag != nil {
		return true
	}

	return false
}

// SetCTag gets a reference to the given string and assigns it to the CTag field.
func (o *DriveItem) SetCTag(v string) {
	o.CTag = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *DriveItem) GetDeleted() Deleted {
	if o == nil || o.Deleted == nil {
		var ret Deleted
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveItem) GetDeletedOk() (*Deleted, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *DriveItem) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given Deleted and assigns it to the Deleted field.
func (o *DriveItem) SetDeleted(v Deleted) {
	o.Deleted = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *DriveItem) GetFile() OpenGraphFile {
	if o == nil || o.File == nil {
		var ret OpenGraphFile
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveItem) GetFileOk() (*OpenGraphFile, bool) {
	if o == nil || o.File == nil {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *DriveItem) HasFile() bool {
	if o != nil && o.File != nil {
		return true
	}

	return false
}

// SetFile gets a reference to the given OpenGraphFile and assigns it to the File field.
func (o *DriveItem) SetFile(v OpenGraphFile) {
	o.File = &v
}

// GetFileSystemInfo returns the FileSystemInfo field value if set, zero value otherwise.
func (o *DriveItem) GetFileSystemInfo() FileSystemInfo {
	if o == nil || o.FileSystemInfo == nil {
		var ret FileSystemInfo
		return ret
	}
	return *o.FileSystemInfo
}

// GetFileSystemInfoOk returns a tuple with the FileSystemInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveItem) GetFileSystemInfoOk() (*FileSystemInfo, bool) {
	if o == nil || o.FileSystemInfo == nil {
		return nil, false
	}
	return o.FileSystemInfo, true
}

// HasFileSystemInfo returns a boolean if a field has been set.
func (o *DriveItem) HasFileSystemInfo() bool {
	if o != nil && o.FileSystemInfo != nil {
		return true
	}

	return false
}

// SetFileSystemInfo gets a reference to the given FileSystemInfo and assigns it to the FileSystemInfo field.
func (o *DriveItem) SetFileSystemInfo(v FileSystemInfo) {
	o.FileSystemInfo = &v
}

// GetFolder returns the Folder field value if set, zero value otherwise.
func (o *DriveItem) GetFolder() Folder {
	if o == nil || o.Folder == nil {
		var ret Folder
		return ret
	}
	return *o.Folder
}

// GetFolderOk returns a tuple with the Folder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveItem) GetFolderOk() (*Folder, bool) {
	if o == nil || o.Folder == nil {
		return nil, false
	}
	return o.Folder, true
}

// HasFolder returns a boolean if a field has been set.
func (o *DriveItem) HasFolder() bool {
	if o != nil && o.Folder != nil {
		return true
	}

	return false
}

// SetFolder gets a reference to the given Folder and assigns it to the Folder field.
func (o *DriveItem) SetFolder(v Folder) {
	o.Folder = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *DriveItem) GetImage() Image {
	if o == nil || o.Image == nil {
		var ret Image
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveItem) GetImageOk() (*Image, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *DriveItem) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given Image and assigns it to the Image field.
func (o *DriveItem) SetImage(v Image) {
	o.Image = &v
}

// GetRoot returns the Root field value if set, zero value otherwise.
func (o *DriveItem) GetRoot() map[string]interface{} {
	if o == nil || o.Root == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Root
}

// GetRootOk returns a tuple with the Root field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveItem) GetRootOk() (*map[string]interface{}, bool) {
	if o == nil || o.Root == nil {
		return nil, false
	}
	return o.Root, true
}

// HasRoot returns a boolean if a field has been set.
func (o *DriveItem) HasRoot() bool {
	if o != nil && o.Root != nil {
		return true
	}

	return false
}

// SetRoot gets a reference to the given map[string]interface{} and assigns it to the Root field.
func (o *DriveItem) SetRoot(v map[string]interface{}) {
	o.Root = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *DriveItem) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveItem) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *DriveItem) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *DriveItem) SetSize(v int64) {
	o.Size = &v
}

// GetWebDavUrl returns the WebDavUrl field value if set, zero value otherwise.
func (o *DriveItem) GetWebDavUrl() string {
	if o == nil || o.WebDavUrl == nil {
		var ret string
		return ret
	}
	return *o.WebDavUrl
}

// GetWebDavUrlOk returns a tuple with the WebDavUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveItem) GetWebDavUrlOk() (*string, bool) {
	if o == nil || o.WebDavUrl == nil {
		return nil, false
	}
	return o.WebDavUrl, true
}

// HasWebDavUrl returns a boolean if a field has been set.
func (o *DriveItem) HasWebDavUrl() bool {
	if o != nil && o.WebDavUrl != nil {
		return true
	}

	return false
}

// SetWebDavUrl gets a reference to the given string and assigns it to the WebDavUrl field.
func (o *DriveItem) SetWebDavUrl(v string) {
	o.WebDavUrl = &v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *DriveItem) GetChildren() []DriveItem {
	if o == nil || o.Children == nil {
		var ret []DriveItem
		return ret
	}
	return *o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriveItem) GetChildrenOk() (*[]DriveItem, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *DriveItem) HasChildren() bool {
	if o != nil && o.Children != nil {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []DriveItem and assigns it to the Children field.
func (o *DriveItem) SetChildren(v []DriveItem) {
	o.Children = &v
}

func (o DriveItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBaseItem, errBaseItem := json.Marshal(o.BaseItem)
	if errBaseItem != nil {
		return []byte{}, errBaseItem
	}
	errBaseItem = json.Unmarshal([]byte(serializedBaseItem), &toSerialize)
	if errBaseItem != nil {
		return []byte{}, errBaseItem
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.CTag != nil {
		toSerialize["cTag"] = o.CTag
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	if o.File != nil {
		toSerialize["file"] = o.File
	}
	if o.FileSystemInfo != nil {
		toSerialize["fileSystemInfo"] = o.FileSystemInfo
	}
	if o.Folder != nil {
		toSerialize["folder"] = o.Folder
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.Root != nil {
		toSerialize["root"] = o.Root
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.WebDavUrl != nil {
		toSerialize["webDavUrl"] = o.WebDavUrl
	}
	if o.Children != nil {
		toSerialize["children"] = o.Children
	}
	return json.Marshal(toSerialize)
}

type NullableDriveItem struct {
	value *DriveItem
	isSet bool
}

func (v NullableDriveItem) Get() *DriveItem {
	return v.value
}

func (v *NullableDriveItem) Set(val *DriveItem) {
	v.value = val
	v.isSet = true
}

func (v NullableDriveItem) IsSet() bool {
	return v.isSet
}

func (v *NullableDriveItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDriveItem(val *DriveItem) *NullableDriveItem {
	return &NullableDriveItem{value: val, isSet: true}
}

func (v NullableDriveItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDriveItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


