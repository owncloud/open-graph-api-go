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
)

// GroupsApiService GroupsApi service
type GroupsApiService service

type ApiCreateGroupRequest struct {
	ctx        context.Context
	ApiService *GroupsApiService
	group      *Group
}

// New entity
func (r ApiCreateGroupRequest) Group(group Group) ApiCreateGroupRequest {
	r.group = &group
	return r
}

func (r ApiCreateGroupRequest) Execute() (*Group, *http.Response, error) {
	return r.ApiService.CreateGroupExecute(r)
}

/*
CreateGroup Add new entity to groups

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateGroupRequest
*/
func (a *GroupsApiService) CreateGroup(ctx context.Context) ApiCreateGroupRequest {
	return ApiCreateGroupRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return Group
func (a *GroupsApiService) CreateGroupExecute(r ApiCreateGroupRequest) (*Group, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Group
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsApiService.CreateGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.group == nil {
		return localVarReturnValue, nil, reportError("group is required and must be specified")
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
	localVarPostBody = r.group
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

type ApiListGroupsRequest struct {
	ctx        context.Context
	ApiService *GroupsApiService
	search     *string
	orderby    *[]string
	select_    *[]string
	expand     *[]string
}

// Search items by search phrases
func (r ApiListGroupsRequest) Search(search string) ApiListGroupsRequest {
	r.search = &search
	return r
}

// Order items by property values
func (r ApiListGroupsRequest) Orderby(orderby []string) ApiListGroupsRequest {
	r.orderby = &orderby
	return r
}

// Select properties to be returned
func (r ApiListGroupsRequest) Select_(select_ []string) ApiListGroupsRequest {
	r.select_ = &select_
	return r
}

// Expand related entities
func (r ApiListGroupsRequest) Expand(expand []string) ApiListGroupsRequest {
	r.expand = &expand
	return r
}

func (r ApiListGroupsRequest) Execute() (*CollectionOfGroup, *http.Response, error) {
	return r.ApiService.ListGroupsExecute(r)
}

/*
ListGroups Get entities from groups

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListGroupsRequest
*/
func (a *GroupsApiService) ListGroups(ctx context.Context) ApiListGroupsRequest {
	return ApiListGroupsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return CollectionOfGroup
func (a *GroupsApiService) ListGroupsExecute(r ApiListGroupsRequest) (*CollectionOfGroup, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CollectionOfGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsApiService.ListGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.search != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "$search", r.search, "")
	}
	if r.orderby != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "$orderby", r.orderby, "csv")
	}
	if r.select_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "$select", r.select_, "csv")
	}
	if r.expand != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "$expand", r.expand, "csv")
	}
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
