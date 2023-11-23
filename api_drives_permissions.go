/*
Libre Graph API

Libre Graph is a free API for cloud collaboration inspired by the MS Graph API.

API version: v1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// DrivesPermissionsApiService DrivesPermissionsApi service
type DrivesPermissionsApiService service

type ApiCreateLinkRequest struct {
	ctx                 context.Context
	ApiService          *DrivesPermissionsApiService
	driveId             string
	itemId              string
	driveItemCreateLink *DriveItemCreateLink
}

// In the request body, provide a JSON object with the following parameters.
func (r ApiCreateLinkRequest) DriveItemCreateLink(driveItemCreateLink DriveItemCreateLink) ApiCreateLinkRequest {
	r.driveItemCreateLink = &driveItemCreateLink
	return r
}

func (r ApiCreateLinkRequest) Execute() (*Permission, *http.Response, error) {
	return r.ApiService.CreateLinkExecute(r)
}

/*
CreateLink Create a sharing link for a DriveItem

You can use the createLink action to share a driveItem via a sharing link.

The response will be a permission object with the link facet containing the created link details.

## Link types

For now, The following values are allowed for the type parameter.

| Value          | Display name      | Description                                                     |
| -------------- | ----------------- | --------------------------------------------------------------- |
| view           | View              | Creates a read-only link to the driveItem.                      |
| upload         | Upload            | Creates a read-write link to the folder driveItem.              |
| edit           | Edit              | Creates a read-write link to the driveItem.                     |
| createOnly     | File Drop         | Creates an upload-only link to the folder driveItem.            |
| blocksDownload | Secure View       | Creates a read-only link that blocks download to the driveItem. |


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param driveId key: id of drive
 @param itemId key: id of item
 @return ApiCreateLinkRequest
*/
func (a *DrivesPermissionsApiService) CreateLink(ctx context.Context, driveId string, itemId string) ApiCreateLinkRequest {
	return ApiCreateLinkRequest{
		ApiService: a,
		ctx:        ctx,
		driveId:    driveId,
		itemId:     itemId,
	}
}

// Execute executes the request
//  @return Permission
func (a *DrivesPermissionsApiService) CreateLinkExecute(r ApiCreateLinkRequest) (*Permission, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Permission
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DrivesPermissionsApiService.CreateLink")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1beta1/drives/{drive-id}/items/{item-id}/createLink"
	localVarPath = strings.Replace(localVarPath, "{"+"drive-id"+"}", url.PathEscape(parameterValueToString(r.driveId, "driveId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"item-id"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.driveItemCreateLink
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v OdataError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeletePermissionRequest struct {
	ctx        context.Context
	ApiService *DrivesPermissionsApiService
	driveId    string
	itemId     string
	permId     string
}

func (r ApiDeletePermissionRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeletePermissionExecute(r)
}

/*
DeletePermission Remove access to a DriveItem

Remove access to a DriveItem.

Only sharing permissions that are not inherited can be deleted. The `inheritedFrom` property must be `null`.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param driveId key: id of drive
 @param itemId key: id of item
 @param permId key: id of permission
 @return ApiDeletePermissionRequest
*/
func (a *DrivesPermissionsApiService) DeletePermission(ctx context.Context, driveId string, itemId string, permId string) ApiDeletePermissionRequest {
	return ApiDeletePermissionRequest{
		ApiService: a,
		ctx:        ctx,
		driveId:    driveId,
		itemId:     itemId,
		permId:     permId,
	}
}

// Execute executes the request
func (a *DrivesPermissionsApiService) DeletePermissionExecute(r ApiDeletePermissionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DrivesPermissionsApiService.DeletePermission")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1beta1/drives/{drive-id}/items/{item-id}/permissions/{perm-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"drive-id"+"}", url.PathEscape(parameterValueToString(r.driveId, "driveId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"item-id"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"perm-id"+"}", url.PathEscape(parameterValueToString(r.permId, "permId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v OdataError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetPermissionRequest struct {
	ctx        context.Context
	ApiService *DrivesPermissionsApiService
	driveId    string
	itemId     string
	permId     string
}

func (r ApiGetPermissionRequest) Execute() (*Permission, *http.Response, error) {
	return r.ApiService.GetPermissionExecute(r)
}

/*
GetPermission Get sharing permission for a file or folder

Return the effective sharing permission for a particular permission resource.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param driveId key: id of drive
 @param itemId key: id of item
 @param permId key: id of permission
 @return ApiGetPermissionRequest
*/
func (a *DrivesPermissionsApiService) GetPermission(ctx context.Context, driveId string, itemId string, permId string) ApiGetPermissionRequest {
	return ApiGetPermissionRequest{
		ApiService: a,
		ctx:        ctx,
		driveId:    driveId,
		itemId:     itemId,
		permId:     permId,
	}
}

// Execute executes the request
//  @return Permission
func (a *DrivesPermissionsApiService) GetPermissionExecute(r ApiGetPermissionRequest) (*Permission, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Permission
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DrivesPermissionsApiService.GetPermission")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1beta1/drives/{drive-id}/items/{item-id}/permissions/{perm-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"drive-id"+"}", url.PathEscape(parameterValueToString(r.driveId, "driveId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"item-id"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"perm-id"+"}", url.PathEscape(parameterValueToString(r.permId, "permId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v OdataError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiInviteRequest struct {
	ctx             context.Context
	ApiService      *DrivesPermissionsApiService
	driveId         string
	itemId          string
	driveItemInvite *DriveItemInvite
}

// In the request body, provide a JSON object with the following parameters. To create a custom role submit a list of actions instead of roles.
func (r ApiInviteRequest) DriveItemInvite(driveItemInvite DriveItemInvite) ApiInviteRequest {
	r.driveItemInvite = &driveItemInvite
	return r
}

func (r ApiInviteRequest) Execute() (*CollectionOfPermissions, *http.Response, error) {
	return r.ApiService.InviteExecute(r)
}

/*
Invite Send a sharing invitation

Sends a sharing invitation for a `driveItem`. A sharing invitation provides permissions to the
recipients and optionally sends them an email with a sharing link.

The response will be a permission object with the grantedToV2 property containing the created grant details.

## Roles property values
For now, roles are only identified by a uuid. There are no hardcoded aliases like `read` or `write` because role actions can be completely customized.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param driveId key: id of drive
 @param itemId key: id of item
 @return ApiInviteRequest
*/
func (a *DrivesPermissionsApiService) Invite(ctx context.Context, driveId string, itemId string) ApiInviteRequest {
	return ApiInviteRequest{
		ApiService: a,
		ctx:        ctx,
		driveId:    driveId,
		itemId:     itemId,
	}
}

// Execute executes the request
//  @return CollectionOfPermissions
func (a *DrivesPermissionsApiService) InviteExecute(r ApiInviteRequest) (*CollectionOfPermissions, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CollectionOfPermissions
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DrivesPermissionsApiService.Invite")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1beta1/drives/{drive-id}/items/{item-id}/invite"
	localVarPath = strings.Replace(localVarPath, "{"+"drive-id"+"}", url.PathEscape(parameterValueToString(r.driveId, "driveId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"item-id"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.driveItemInvite
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v OdataError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v OdataError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListPermissionsRequest struct {
	ctx        context.Context
	ApiService *DrivesPermissionsApiService
	driveId    string
	itemId     string
}

func (r ApiListPermissionsRequest) Execute() (*CollectionOfPermissionsWithAllowedValues, *http.Response, error) {
	return r.ApiService.ListPermissionsExecute(r)
}

/*
ListPermissions List the effective sharing permissions on a driveItem.

The permissions collection includes potentially sensitive information and may not be available for every caller.

* For the owner of the item, all sharing permissions will be returned. This includes co-owners.
* For a non-owner caller, only the sharing permissions that apply to the caller are returned.
* Sharing permission properties that contain secrets (e.g. `webUrl`) are only returned for callers that are able to create the sharing permission.

All permission objects have an `id`. A permission representing
* a link has the `link` facet filled with details.
* a share has the `roles` property set and the `grantedToV2` property filled with the grant recipient details.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param driveId key: id of drive
 @param itemId key: id of item
 @return ApiListPermissionsRequest
*/
func (a *DrivesPermissionsApiService) ListPermissions(ctx context.Context, driveId string, itemId string) ApiListPermissionsRequest {
	return ApiListPermissionsRequest{
		ApiService: a,
		ctx:        ctx,
		driveId:    driveId,
		itemId:     itemId,
	}
}

// Execute executes the request
//  @return CollectionOfPermissionsWithAllowedValues
func (a *DrivesPermissionsApiService) ListPermissionsExecute(r ApiListPermissionsRequest) (*CollectionOfPermissionsWithAllowedValues, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CollectionOfPermissionsWithAllowedValues
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DrivesPermissionsApiService.ListPermissions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1beta1/drives/{drive-id}/items/{item-id}/permissions"
	localVarPath = strings.Replace(localVarPath, "{"+"drive-id"+"}", url.PathEscape(parameterValueToString(r.driveId, "driveId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"item-id"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v OdataError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSetPermissionPasswordRequest struct {
	ctx                 context.Context
	ApiService          *DrivesPermissionsApiService
	driveId             string
	itemId              string
	permId              string
	sharingLinkPassword *SharingLinkPassword
}

// New password value
func (r ApiSetPermissionPasswordRequest) SharingLinkPassword(sharingLinkPassword SharingLinkPassword) ApiSetPermissionPasswordRequest {
	r.sharingLinkPassword = &sharingLinkPassword
	return r
}

func (r ApiSetPermissionPasswordRequest) Execute() (*Permission, *http.Response, error) {
	return r.ApiService.SetPermissionPasswordExecute(r)
}

/*
SetPermissionPassword Set sharing link password

Set the password of a sharing permission.

Only the `password` property can be modified this way.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param driveId key: id of drive
 @param itemId key: id of item
 @param permId key: id of permission
 @return ApiSetPermissionPasswordRequest
*/
func (a *DrivesPermissionsApiService) SetPermissionPassword(ctx context.Context, driveId string, itemId string, permId string) ApiSetPermissionPasswordRequest {
	return ApiSetPermissionPasswordRequest{
		ApiService: a,
		ctx:        ctx,
		driveId:    driveId,
		itemId:     itemId,
		permId:     permId,
	}
}

// Execute executes the request
//  @return Permission
func (a *DrivesPermissionsApiService) SetPermissionPasswordExecute(r ApiSetPermissionPasswordRequest) (*Permission, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Permission
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DrivesPermissionsApiService.SetPermissionPassword")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1beta1/drives/{drive-id}/items/{item-id}/permissions/{perm-id}/setPassword"
	localVarPath = strings.Replace(localVarPath, "{"+"drive-id"+"}", url.PathEscape(parameterValueToString(r.driveId, "driveId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"item-id"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"perm-id"+"}", url.PathEscape(parameterValueToString(r.permId, "permId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.sharingLinkPassword == nil {
		return localVarReturnValue, nil, reportError("sharingLinkPassword is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.sharingLinkPassword
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v OdataError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdatePermissionRequest struct {
	ctx        context.Context
	ApiService *DrivesPermissionsApiService
	driveId    string
	itemId     string
	permId     string
	permission *Permission
}

// New property values
func (r ApiUpdatePermissionRequest) Permission(permission Permission) ApiUpdatePermissionRequest {
	r.permission = &permission
	return r
}

func (r ApiUpdatePermissionRequest) Execute() (*Permission, *http.Response, error) {
	return r.ApiService.UpdatePermissionExecute(r)
}

/*
UpdatePermission Update sharing permission

Update the properties of a sharing permission by patching the permission resource.

Only the `roles`, `expirationDateTime` and `password` properties can be modified this way.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param driveId key: id of drive
 @param itemId key: id of item
 @param permId key: id of permission
 @return ApiUpdatePermissionRequest
*/
func (a *DrivesPermissionsApiService) UpdatePermission(ctx context.Context, driveId string, itemId string, permId string) ApiUpdatePermissionRequest {
	return ApiUpdatePermissionRequest{
		ApiService: a,
		ctx:        ctx,
		driveId:    driveId,
		itemId:     itemId,
		permId:     permId,
	}
}

// Execute executes the request
//  @return Permission
func (a *DrivesPermissionsApiService) UpdatePermissionExecute(r ApiUpdatePermissionRequest) (*Permission, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Permission
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DrivesPermissionsApiService.UpdatePermission")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1beta1/drives/{drive-id}/items/{item-id}/permissions/{perm-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"drive-id"+"}", url.PathEscape(parameterValueToString(r.driveId, "driveId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"item-id"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"perm-id"+"}", url.PathEscape(parameterValueToString(r.permId, "permId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.permission == nil {
		return localVarReturnValue, nil, reportError("permission is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.permission
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v OdataError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
