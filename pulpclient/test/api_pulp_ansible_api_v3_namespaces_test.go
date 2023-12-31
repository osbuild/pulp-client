/*
Pulp 3 API

Testing PulpAnsibleApiV3NamespacesAPIService

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

func Test_pulpclient_PulpAnsibleApiV3NamespacesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PulpAnsibleApiV3NamespacesAPIService PulpAnsibleGalaxyApiV3NamespacesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var path string

		resp, httpRes, err := apiClient.PulpAnsibleApiV3NamespacesAPI.PulpAnsibleGalaxyApiV3NamespacesList(context.Background(), path).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpAnsibleApiV3NamespacesAPIService PulpAnsibleGalaxyApiV3NamespacesRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var path string

		resp, httpRes, err := apiClient.PulpAnsibleApiV3NamespacesAPI.PulpAnsibleGalaxyApiV3NamespacesRead(context.Background(), name, path).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
