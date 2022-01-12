/*
Libre Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v0.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// MeDrivesApiService MeDrivesApi service
type MeDrivesApiService service

type ApiListMyDrivesRequest struct {
	ctx _context.Context
	ApiService *MeDrivesApiService
	top *int32
	skip *int32
	search *string
	filter *string
	count *bool
	select_ *[]string
	expand *[]string
}

// Show only the first n items
func (r ApiListMyDrivesRequest) Top(top int32) ApiListMyDrivesRequest {
	r.top = &top
	return r
}
// Skip the first n items
func (r ApiListMyDrivesRequest) Skip(skip int32) ApiListMyDrivesRequest {
	r.skip = &skip
	return r
}
// Search items by search phrases
func (r ApiListMyDrivesRequest) Search(search string) ApiListMyDrivesRequest {
	r.search = &search
	return r
}
// Filter items by property values
func (r ApiListMyDrivesRequest) Filter(filter string) ApiListMyDrivesRequest {
	r.filter = &filter
	return r
}
// Include count of items
func (r ApiListMyDrivesRequest) Count(count bool) ApiListMyDrivesRequest {
	r.count = &count
	return r
}
// Select properties to be returned
func (r ApiListMyDrivesRequest) Select_(select_ []string) ApiListMyDrivesRequest {
	r.select_ = &select_
	return r
}
// Expand related entities
func (r ApiListMyDrivesRequest) Expand(expand []string) ApiListMyDrivesRequest {
	r.expand = &expand
	return r
}

func (r ApiListMyDrivesRequest) Execute() (CollectionOfDrives, *_nethttp.Response, error) {
	return r.ApiService.ListMyDrivesExecute(r)
}

/*
ListMyDrives Get drives from me

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListMyDrivesRequest
*/
func (a *MeDrivesApiService) ListMyDrives(ctx _context.Context) ApiListMyDrivesRequest {
	return ApiListMyDrivesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CollectionOfDrives
func (a *MeDrivesApiService) ListMyDrivesExecute(r ApiListMyDrivesRequest) (CollectionOfDrives, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  CollectionOfDrives
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MeDrivesApiService.ListMyDrives")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/me/drives"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.top != nil {
		localVarQueryParams.Add("$top", parameterToString(*r.top, ""))
	}
	if r.skip != nil {
		localVarQueryParams.Add("$skip", parameterToString(*r.skip, ""))
	}
	if r.search != nil {
		localVarQueryParams.Add("$search", parameterToString(*r.search, ""))
	}
	if r.filter != nil {
		localVarQueryParams.Add("$filter", parameterToString(*r.filter, ""))
	}
	if r.count != nil {
		localVarQueryParams.Add("$count", parameterToString(*r.count, ""))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("$select", parameterToString(*r.select_, "csv"))
	}
	if r.expand != nil {
		localVarQueryParams.Add("$expand", parameterToString(*r.expand, "csv"))
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

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
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
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
