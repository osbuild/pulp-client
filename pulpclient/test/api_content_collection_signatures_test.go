/*
Pulp 3 API

Testing ContentCollectionSignaturesAPIService

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

func Test_pulpclient_ContentCollectionSignaturesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentCollectionSignaturesAPIService ContentAnsibleCollectionSignaturesCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentCollectionSignaturesAPI.ContentAnsibleCollectionSignaturesCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentCollectionSignaturesAPIService ContentAnsibleCollectionSignaturesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentCollectionSignaturesAPI.ContentAnsibleCollectionSignaturesList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentCollectionSignaturesAPIService ContentAnsibleCollectionSignaturesRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleCollectionVersionSignatureHref string

		resp, httpRes, err := apiClient.ContentCollectionSignaturesAPI.ContentAnsibleCollectionSignaturesRead(context.Background(), ansibleCollectionVersionSignatureHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
