package registryinstance

import (
	"github.com/microsoft/kiota-abstractions-go/authentication"
	http "github.com/microsoft/kiota-http-go"
	apiv1 "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client"
	auth "github.com/rkpattnaik780/rh-kiota-go-auth"
)

func NewAPIClient(adapter auth.RHAccessTokenProvider, baseURL string) (*apiv1.RegistryInstance, error) {

	bearerTokenProvider := authentication.NewBaseBearerTokenAuthenticationProvider(adapter)

	requestAdapter, err := http.NewNetHttpRequestAdapter(bearerTokenProvider)
	if err != nil {
		return nil, err
	}

	requestAdapter.SetBaseUrl(baseURL)

	registryinstance := apiv1.NewRegistryInstance(requestAdapter)

	return registryinstance, nil

}
