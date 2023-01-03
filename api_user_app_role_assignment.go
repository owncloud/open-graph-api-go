/*
Libre Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// UserAppRoleAssignmentApiService UserAppRoleAssignmentApi service
type UserAppRoleAssignmentApiService service

type ApiUserCreateAppRoleAssignmentsRequest struct {
	ctx context.Context
	ApiService *UserAppRoleAssignmentApiService
	userId string
	appRoleAssignment *AppRoleAssignment
}

// New app role assignment value
func (r ApiUserCreateAppRoleAssignmentsRequest) AppRoleAssignment(appRoleAssignment AppRoleAssignment) ApiUserCreateAppRoleAssignmentsRequest {
	r.appRoleAssignment = &appRoleAssignment
	return r
}

func (r ApiUserCreateAppRoleAssignmentsRequest) Execute() (*AppRoleAssignment, *http.Response, error) {
	return r.ApiService.UserCreateAppRoleAssignmentsExecute(r)
}

/*
UserCreateAppRoleAssignments Grant an appRoleAssignment to a user

Use this API to assign a global role to a user. To grant an app role assignment to a user, you need three identifiers:
* `principalId`: The `id` of the user to whom you are assigning the app role.
* `resourceId`: The `id` of the resource `servicePrincipal` or `application` that has defined the app role.
* `appRoleId`: The `id` of the `appRole` (defined on the resource service principal or application) to assign to the user.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId key: id of user
 @return ApiUserCreateAppRoleAssignmentsRequest
*/
func (a *UserAppRoleAssignmentApiService) UserCreateAppRoleAssignments(ctx context.Context, userId string) ApiUserCreateAppRoleAssignmentsRequest {
	return ApiUserCreateAppRoleAssignmentsRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return AppRoleAssignment
func (a *UserAppRoleAssignmentApiService) UserCreateAppRoleAssignmentsExecute(r ApiUserCreateAppRoleAssignmentsRequest) (*AppRoleAssignment, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AppRoleAssignment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAppRoleAssignmentApiService.UserCreateAppRoleAssignments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user-id}/appRoleAssignments"
	localVarPath = strings.Replace(localVarPath, "{"+"user-id"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.appRoleAssignment == nil {
		return localVarReturnValue, nil, reportError("appRoleAssignment is required and must be specified")
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
	localVarPostBody = r.appRoleAssignment
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiUserDeleteAppRoleAssignmentsRequest struct {
	ctx context.Context
	ApiService *UserAppRoleAssignmentApiService
	userId string
	appRoleAssignmentId string
	ifMatch *string
}

// ETag
func (r ApiUserDeleteAppRoleAssignmentsRequest) IfMatch(ifMatch string) ApiUserDeleteAppRoleAssignmentsRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiUserDeleteAppRoleAssignmentsRequest) Execute() (*http.Response, error) {
	return r.ApiService.UserDeleteAppRoleAssignmentsExecute(r)
}

/*
UserDeleteAppRoleAssignments Delete the appRoleAssignment from a user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId key: id of user
 @param appRoleAssignmentId key: id of appRoleAssignment. This is the concatenated {user-id}:{appRole-id} separated by a colon.
 @return ApiUserDeleteAppRoleAssignmentsRequest
*/
func (a *UserAppRoleAssignmentApiService) UserDeleteAppRoleAssignments(ctx context.Context, userId string, appRoleAssignmentId string) ApiUserDeleteAppRoleAssignmentsRequest {
	return ApiUserDeleteAppRoleAssignmentsRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
		appRoleAssignmentId: appRoleAssignmentId,
	}
}

// Execute executes the request
func (a *UserAppRoleAssignmentApiService) UserDeleteAppRoleAssignmentsExecute(r ApiUserDeleteAppRoleAssignmentsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAppRoleAssignmentApiService.UserDeleteAppRoleAssignments")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user-id}/appRoleAssignments/{appRoleAssignment-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"user-id"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appRoleAssignment-id"+"}", url.PathEscape(parameterToString(r.appRoleAssignmentId, "")), -1)

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
	if r.ifMatch != nil {
		localVarHeaderParams["If-Match"] = parameterToString(*r.ifMatch, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
			newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUserListAppRoleAssignmentsRequest struct {
	ctx context.Context
	ApiService *UserAppRoleAssignmentApiService
	userId string
}

func (r ApiUserListAppRoleAssignmentsRequest) Execute() (*CollectionOfAppRoleAssignments, *http.Response, error) {
	return r.ApiService.UserListAppRoleAssignmentsExecute(r)
}

/*
UserListAppRoleAssignments Get appRoleAssignments from a user

Represents the global roles a user has been granted for an application.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId key: id of user
 @return ApiUserListAppRoleAssignmentsRequest
*/
func (a *UserAppRoleAssignmentApiService) UserListAppRoleAssignments(ctx context.Context, userId string) ApiUserListAppRoleAssignmentsRequest {
	return ApiUserListAppRoleAssignmentsRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return CollectionOfAppRoleAssignments
func (a *UserAppRoleAssignmentApiService) UserListAppRoleAssignmentsExecute(r ApiUserListAppRoleAssignmentsRequest) (*CollectionOfAppRoleAssignments, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CollectionOfAppRoleAssignments
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAppRoleAssignmentApiService.UserListAppRoleAssignments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user-id}/appRoleAssignments"
	localVarPath = strings.Replace(localVarPath, "{"+"user-id"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
