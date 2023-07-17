package registryinstance

import (
	"context"

	u "net/url"

	"github.com/microsoft/kiota-abstractions-go/authentication"
	http "github.com/microsoft/kiota-http-go"
	"github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client"
	apiv1 "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client"
)

type RHAccessTokenProvider struct {
}

func (r *RHAccessTokenProvider) GetAuthorizationToken(context context.Context, url *u.URL, additionalAuthenticationContext map[string]interface{}) (string, error) {

	return "", nil
}

func (r *RHAccessTokenProvider) GetAllowedHostsValidator() *authentication.AllowedHostsValidator {

	allowedHostsValidator := authentication.NewAllowedHostsValidator([]string{"http://localhost:8080"})

	return &allowedHostsValidator

}

func NewAPIClient(adapter authentication.AccessTokenProvider) (*apiv1.RegistryInstance, error) {

	// RHAS := &RHAccessTokenProvider{}

	// var adapter authentication.AccessTokenProvider = RHAS

	bearerTokenProvider := authentication.NewBaseBearerTokenAuthenticationProvider(adapter)

	adapter2, err := http.NewNetHttpRequestAdapter(bearerTokenProvider)
	if err != nil {
		return nil, err
	}

	registryinstance := client.NewRegistryInstance(adapter2)

	return registryinstance, nil

}
