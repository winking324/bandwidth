/*
Bandwidth

Testing MFAAPIService

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

func Test_openapi_MFAAPIService(t *testing.T) {

	configuration := bandwidth.NewConfiguration()
	apiClient := bandwidth.NewAPIClient(configuration)

	t.Run("Test MFAAPIService GenerateMessagingCode", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.MFAAPI.GenerateMessagingCode(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MFAAPIService GenerateVoiceCode", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.MFAAPI.GenerateVoiceCode(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MFAAPIService VerifyCode", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.MFAAPI.VerifyCode(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
