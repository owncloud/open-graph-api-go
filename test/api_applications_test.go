/*
Libre Graph API

Testing ApplicationsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package libregraph

import (
	"context"
	openapiclient "github.com/owncloud/libre-graph-api-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_libregraph_ApplicationsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ApplicationsApiService GetApplication", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationId string

		resp, httpRes, err := apiClient.ApplicationsApi.GetApplication(context.Background(), applicationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsApiService ListApplications", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ApplicationsApi.ListApplications(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
