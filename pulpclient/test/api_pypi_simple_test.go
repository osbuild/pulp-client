/*
Pulp 3 API

Testing PypiSimpleAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package pulpclient

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/osbuild/pulp-client/pulpclient"
)

func Test_pulpclient_PypiSimpleAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PypiSimpleAPIService PypiSimpleCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var path string

		resp, httpRes, err := apiClient.PypiSimpleAPI.PypiSimpleCreate(context.Background(), path).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PypiSimpleAPIService PypiSimplePackageRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var package_ string
		var path string

		httpRes, err := apiClient.PypiSimpleAPI.PypiSimplePackageRead(context.Background(), package_, path).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PypiSimpleAPIService PypiSimpleRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var path string

		httpRes, err := apiClient.PypiSimpleAPI.PypiSimpleRead(context.Background(), path).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
