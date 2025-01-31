/*
Bandwidth

Testing MediaAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package bandwidth

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winking324/bandwidth/pkg/bandwidth"
)

func Test_openapi_MediaAPIService(t *testing.T) {

	configuration := bandwidth.NewConfiguration()
	apiClient := bandwidth.NewAPIClient(configuration)

	t.Run("Test MediaAPIService DeleteMedia", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountId string
		var mediaId string

		httpRes, err := apiClient.MediaAPI.DeleteMedia(context.Background(), accountId, mediaId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MediaAPIService GetMedia", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountId string
		var mediaId string

		resp, httpRes, err := apiClient.MediaAPI.GetMedia(context.Background(), accountId, mediaId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MediaAPIService ListMedia", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.MediaAPI.ListMedia(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MediaAPIService UploadMedia", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountId string
		var mediaId string

		httpRes, err := apiClient.MediaAPI.UploadMedia(context.Background(), accountId, mediaId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
