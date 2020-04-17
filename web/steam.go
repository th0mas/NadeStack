// Steam doesnt use industry-standard authentication so we have to roll our own
// valve pls fix
//
// Based of code from the Goth Package by Mark Bates & contributors
// https://github.com/markbates/goth licensed under the MIT Licence

package web

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/th0mas/NadeStack/config"
)

var (
	steamEndpoint    = "https://steamcommunity.com/openid/login"
	openIDMode       = "checkid_setup"
	openIDNs         = "http://specs.openid.net/auth/2.0"
	openIDIdentifier = "http://specs.openid.net/auth/2.0/identifier_select"
)

func generateSteamOpenIDUrl(c *config.Config, uniqueID string) *url.URL {
	callback, err := url.Parse(c.Domain + "/verify/" + uniqueID)
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

	vals := make(url.Values)
	for key, val := range urlVals {
		vals.Set(key, val)
	}

	steamURL.RawQuery = vals.Encode()

	return steamURL
}

func verifySteamCallback(ctx *gin.Context, c *config.Config, d string) (string, error) {
	if ctx.Query("openid.mode") != "id_res" {
		return "", errors.New("wrong openID mode")
	}
	fmt.Println(ctx.Query("openid.return_to"))
	fmt.Println(c.Domain + "/verify/" + d)
	if ctx.Query("openid.return_to") != c.Domain+"/verify/"+d {
		return "", errors.New("return_to has wrong discord_id")
	}

	vals := map[string]string{
		"openid.assoc_handle": ctx.Query("openid.assoc_handle"),
		"openid.signed":       ctx.Query("openid.signed"),
		"openid.sig":          ctx.Query("openid.sig"),
		"openid.ns":           ctx.Query("openid.ns"),
	}

	v := make(url.Values)
	split := strings.Split(ctx.Query("openid.signed"), ",")

	for key, val := range vals {
		v.Set(key, val)
	}

	for _, val := range split {
		v.Set("openid."+val, ctx.Query("openid."+val))
	}

	v.Set("openid.mode", "check_authentication")

	resp, err := http.PostForm(steamEndpoint, v)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	content := strings.Split(string(body), "\n")

	if content[0] != "ns:"+openIDNs {
		return "", errors.New("wrong ns in the response")
	}

	if content[1] == "is_valid:false" {
		return "", errors.New("unable validate openId")
	}

	steamID := regexp.MustCompile("\\D+").ReplaceAllString(ctx.Query("openid.claimed_id"), "")

	return steamID, nil

}

// ConvertSteamID64toSteamID3 returns a SteamID3, e.g. STEAM_X:Y:Z when
// given a steamID64
func ConvertSteamID64toSteamID3(steamID64 uint64) string {
	x := (steamID64 >> 56) & 0xFF
	if x == 1 {
		x = 0
	} // Universe 1 is unusable

	y := steamID64 & 1

	z := (steamID64 >> 1) & 0x7FFFFFF

	return fmt.Sprintf("STEAM_%d:%d:%d", x, y, z)

}
