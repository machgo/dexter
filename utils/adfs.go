package utils

import (
	"golang.org/x/oauth2"
)

func ADFSEndpoint(tenant string) oauth2.Endpoint {
	return oauth2.Endpoint{
		AuthURL:  tenant + "/oauth2/authorize",
		TokenURL: tenant + "/oauth2/token",
	}
}
