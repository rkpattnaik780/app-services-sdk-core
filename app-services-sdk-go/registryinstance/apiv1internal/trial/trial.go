package main

import (
	"fmt"

	registryinstance "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal"
	"github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client/admin"
	auth "github.com/rkpattnaik780/rh-kiota-go-auth"
)

func main() {

	token := map[string]string{
		"refresh-token": "hfkdldjdjfj",
	}

	RHAS := auth.NewRHAccessTokenProvider(token, []string{"http://localhost:8080", "https://api.openshift.com"})

	registry, err := registryinstance.NewAPIClient(RHAS)
	if err != nil {
		fmt.Println(err)
	}

	// config := *admin.RulesRequestBuilderGetRequestConfiguration{

	// }

	config := admin.NewRulesRequestBuilder("http://localhost:8080", RHAS)

	registry.Admin().Rules().Get(config)

}
