// Steam doesnt use industry-standard authentication so we have to roll our own
// valve pls fix
//
// Based of code from the Goth Package by Mark Bates
// https://github.com/markbates/goth licensed under the MIT Licence

package web

import (
	"fmt"
	"github.com/th0mas/NadeStack/config"
	"net/url"
)

var (
	steamEndpoint    = "https://steamcommunity.com/openid/login"
	openIDMode       = "checkid_setup"
	openIDNs         = "http://specs.openid.net/auth/2.0"
	openIDIdentifier = "http://specs.openid.net/auth/2.0/identifier_select"
)

func generateSteamOpenIdUrl(c *config.Config) *url.URL {
	callback, err := url.Parse(c.Domain + "/auth/callback")
	if err != nil {
		panic(err)
	}

	urlVals := map[string]string{
		"openid.claimed_id": openIDIdentifier,
		"openid.identity":   openIDIdentifier,
		"openid.mode":       openIDMode,
		"openid.ns":         openIDNs,
		"openid.realm":      fmt.Sprintf("%s://%s", callback.Scheme, callback.Host),
		"openid.return_to":  callback.String(),
	}

	steamURL, err := url.Parse(steamEndpoint)
	if err != nil {
		panic(err)
	}

	vals := steamURL.Query()
	for key, val := range urlVals {
		vals.Set(key, val)
	}

	steamURL.RawQuery = vals.Encode()

	return steamURL
}
