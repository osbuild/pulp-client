/*
Pulp 3 API

Testing AnsibleCollectionsAPIService

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

func Test_pulpclient_AnsibleCollectionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AnsibleCollectionsAPIService AnsibleCollectionsAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleCollectionHref string

		resp, httpRes, err := apiClient.AnsibleCollectionsAPI.AnsibleCollectionsAddRole(context.Background(), ansibleCollectionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AnsibleCollectionsAPIService AnsibleCollectionsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AnsibleCollectionsAPI.AnsibleCollectionsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AnsibleCollectionsAPIService AnsibleCollectionsListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleCollectionHref string

		resp, httpRes, err := apiClient.AnsibleCollectionsAPI.AnsibleCollectionsListRoles(context.Background(), ansibleCollectionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AnsibleCollectionsAPIService AnsibleCollectionsMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleCollectionHref string

		resp, httpRes, err := apiClient.AnsibleCollectionsAPI.AnsibleCollectionsMyPermissions(context.Background(), ansibleCollectionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AnsibleCollectionsAPIService AnsibleCollectionsRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleCollectionHref string

		resp, httpRes, err := apiClient.AnsibleCollectionsAPI.AnsibleCollectionsRemoveRole(context.Background(), ansibleCollectionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AnsibleCollectionsAPIService UploadCollection", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AnsibleCollectionsAPI.UploadCollection(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
