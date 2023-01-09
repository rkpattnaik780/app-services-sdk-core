/*
Service Registry API

Testing ArtifactsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package registryinstanceclient

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_registryinstanceclient_ArtifactsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ArtifactsApiService CreateArtifact", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string

        resp, httpRes, err := apiClient.ArtifactsApi.CreateArtifact(context.Background(), groupId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ArtifactsApiService DeleteArtifact", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var artifactId string

        resp, httpRes, err := apiClient.ArtifactsApi.DeleteArtifact(context.Background(), groupId, artifactId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ArtifactsApiService DeleteArtifactsInGroup", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string

        resp, httpRes, err := apiClient.ArtifactsApi.DeleteArtifactsInGroup(context.Background(), groupId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ArtifactsApiService GetContentByGlobalId", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var globalId int64

        resp, httpRes, err := apiClient.ArtifactsApi.GetContentByGlobalId(context.Background(), globalId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ArtifactsApiService GetContentByHash", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var contentHash string

        resp, httpRes, err := apiClient.ArtifactsApi.GetContentByHash(context.Background(), contentHash).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ArtifactsApiService GetContentById", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var contentId int64

        resp, httpRes, err := apiClient.ArtifactsApi.GetContentById(context.Background(), contentId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ArtifactsApiService GetLatestArtifact", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var artifactId string

        resp, httpRes, err := apiClient.ArtifactsApi.GetLatestArtifact(context.Background(), groupId, artifactId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ArtifactsApiService ListArtifactsInGroup", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string

        resp, httpRes, err := apiClient.ArtifactsApi.ListArtifactsInGroup(context.Background(), groupId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ArtifactsApiService ReferencesByContentHash", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var contentHash string

        resp, httpRes, err := apiClient.ArtifactsApi.ReferencesByContentHash(context.Background(), contentHash).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ArtifactsApiService ReferencesByContentId", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var contentId int64

        resp, httpRes, err := apiClient.ArtifactsApi.ReferencesByContentId(context.Background(), contentId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ArtifactsApiService ReferencesByGlobalId", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var globalId int64

        resp, httpRes, err := apiClient.ArtifactsApi.ReferencesByGlobalId(context.Background(), globalId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ArtifactsApiService UpdateArtifact", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var artifactId string

        resp, httpRes, err := apiClient.ArtifactsApi.UpdateArtifact(context.Background(), groupId, artifactId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ArtifactsApiService UpdateArtifactState", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var groupId string
        var artifactId string

        resp, httpRes, err := apiClient.ArtifactsApi.UpdateArtifactState(context.Background(), groupId, artifactId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
