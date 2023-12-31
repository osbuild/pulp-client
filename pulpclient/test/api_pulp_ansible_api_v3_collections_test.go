/*
Pulp 3 API

Testing PulpAnsibleApiV3CollectionsAPIService

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

func Test_pulpclient_PulpAnsibleApiV3CollectionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PulpAnsibleApiV3CollectionsAPIService PulpAnsibleGalaxyApiV3CollectionsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string
		var path string

		resp, httpRes, err := apiClient.PulpAnsibleApiV3CollectionsAPI.PulpAnsibleGalaxyApiV3CollectionsDelete(context.Background(), name, namespace, path).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpAnsibleApiV3CollectionsAPIService PulpAnsibleGalaxyApiV3CollectionsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var path string

		resp, httpRes, err := apiClient.PulpAnsibleApiV3CollectionsAPI.PulpAnsibleGalaxyApiV3CollectionsList(context.Background(), path).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpAnsibleApiV3CollectionsAPIService PulpAnsibleGalaxyApiV3CollectionsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string
		var path string

		resp, httpRes, err := apiClient.PulpAnsibleApiV3CollectionsAPI.PulpAnsibleGalaxyApiV3CollectionsRead(context.Background(), name, namespace, path).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpAnsibleApiV3CollectionsAPIService PulpAnsibleGalaxyApiV3CollectionsUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string
		var path string

		resp, httpRes, err := apiClient.PulpAnsibleApiV3CollectionsAPI.PulpAnsibleGalaxyApiV3CollectionsUpdate(context.Background(), name, namespace, path).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
