package tools

import (
	"WarpGPT/pkg/env"
	"WarpGPT/pkg/funcaptcha"
	"WarpGPT/pkg/logger"
	"encoding/json"
	"fmt"
	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
	"io"
	"net/url"
	"os"
	"regexp"
	"strings"
)

type Error struct {
	Location   string
	StatusCode int
	Details    string
	Error      error
}

func NewError(location string, statusCode int, details string, err error) *Error {
	return &Error{
		Location:   location,
		StatusCode: statusCode,
		Details:    details,
		Error:      err,
	}
}

type Authenticator struct {
	EmailAddress       string
	Password           string
	Proxy              string
	Session            tls_client.HttpClient
	UserAgent          string
	State              string
	URL                string
	PUID               string
	Verifier_code      string
	Verifier_challenge string
	AuthResult         AuthResult
}
type ArkoseToken struct {
	Token              string  `json:"token"`
	ChallengeURL       string  `json:"challenge_url"`
	ChallengeURLCDN    string  `json:"challenge_url_cdn"`
	ChallengeURLCDNSRI *string `json:"challenge_url_cdn_sri"`
}
type AuthResult struct {
	AccessToken map[string]interface{} `json:"access_token"`
	PUID        string                 `json:"puid"`
	FreshToken  string                 `json:"fresh_token"`
	Model       map[string]interface{} `json:"model"`
}

func NewAuthenticator(emailAddress, password string, puid string) *Authenticator {
	auth := &Authenticator{
		EmailAddress: emailAddress,
		Password:     password,
		Proxy:        os.Getenv("proxy"),
		PUID:         puid,
		UserAgent:    "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
	}
	jar := tls_client.NewCookieJar()
	cookie := &http.Cookie{
		Name:   "_puid",
		Value:  puid,
		Path:   "/",
		Domain: ".openai.com",
	}
	urls, _ := url.Parse("https://openai.com")
	jar.SetCookies(urls, []*http.Cookie{cookie})
	options := []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(20),
		tls_client.WithClientProfile(profiles.Chrome_109),
		tls_client.WithNotFollowRedirects(),
		tls_client.WithCookieJar(jar),
		tls_client.WithProxyUrl(env.E.Proxy),
	}
	auth.Session, _ = tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	return auth
}

func (auth *Authenticator) URLEncode(str string) string {
	return url.QueryEscape(str)
}

func (auth *Authenticator) Begin() *Error {
	logger.Log.Debug("Auth Begin")

	target := "https://" + env.E.OpenaiHost + "/api/auth/csrf"
	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		return NewError("begin", 0, "", err)
	}

	req.Header.Set("Host", ""+env.E.OpenaiHost+"")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("User-Agent", auth.UserAgent)
	req.Header.Set("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Set("Referer", "https://"+env.E.OpenaiHost+"/auth/login")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")

	resp, err := auth.Session.Do(req)
	if err != nil {
		return NewError("begin", 0, "", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return NewError("begin", 0, "", err)
	}

	if resp.StatusCode == 200 && strings.Contains(resp.Header.Get("Content-Type"), "json") {

		var csrfTokenResponse struct {
			CsrfToken string `json:"csrfToken"`
		}
		err = json.Unmarshal(body, &csrfTokenResponse)
		if err != nil {
			return NewError("begin", 0, "", err)
		}

		csrfToken := csrfTokenResponse.CsrfToken
		return auth.partOne(csrfToken)
	} else {
		err := NewError("begin", resp.StatusCode, string(body), fmt.Errorf("error: Check details"))
		return err
	}
}

func (auth *Authenticator) partOne(csrfToken string) *Error {
	logger.Log.Debug("Auth One")

	auth_url := "https://" + env.E.OpenaiHost + "/api/auth/signin/auth0?prompt=login"
	headers := map[string]string{
		"Host":            "" + env.E.OpenaiHost + "",
		"User-Agent":      auth.UserAgent,
		"Content-Type":    "application/x-www-form-urlencoded",
		"Accept":          "*/*",
		"Sec-Gpc":         "1",
		"Accept-Language": "en-US,en;q=0.8",
		"Origin":          "https://" + env.E.OpenaiHost + "",
		"Sec-Fetch-Site":  "same-origin",
		"Sec-Fetch-Mode":  "cors",
		"Sec-Fetch-Dest":  "empty",
		"Referer":         "https://" + env.E.OpenaiHost + "/auth/login",
		"Accept-Encoding": "gzip, deflate",
	}

	// Construct payload
	payload := fmt.Sprintf("callbackUrl=%%2F&csrfToken=%s&json=true", csrfToken)
	req, _ := http.NewRequest("POST", auth_url, strings.NewReader(payload))

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := auth.Session.Do(req)
	if err != nil {
		return NewError("part_one", 0, "Failed to send request", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return NewError("part_one", 0, "Failed to read requestbody", err)
	}

	if resp.StatusCode == 200 && strings.Contains(resp.Header.Get("Content-Type"), "json") {
		var urlResponse struct {
			URL string `json:"url"`
		}
		err = json.Unmarshal(body, &urlResponse)
		if err != nil {
			return NewError("part_one", 0, "Failed to decode JSON", err)
		}
		if urlResponse.URL == "https://"+env.E.OpenaiHost+"/api/auth/error?error=OAuthSignin" || strings.Contains(urlResponse.URL, "error") {
			err := NewError("part_one", resp.StatusCode, "You have been rate limited. Please try again later.", fmt.Errorf("error: Check details"))
			return err
		}
		return auth.partTwo(urlResponse.URL)
	} else {
		return NewError("part_one", resp.StatusCode, string(body), fmt.Errorf("error: Check details"))
	}
}

func (auth *Authenticator) partTwo(target string) *Error {
	logger.Log.Debug("Auth Two")

	headers := map[string]string{
		"Host":                        "auth0.openai.com",
		"Accept":                      "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
		"Connection":                  "keep-alive",
		"User-Agent":                  auth.UserAgent,
		"Accept-Language":             "en-US,en;q=0.9",
		"Referer":                     "https://chat.openai.com/",
		"Sec-Ch-Ua":                   "\"Not A(Brand\";v=\"99\", \"Google Chrome\";v=\"121\", \"Chromium\";v=\"121\"",
		"Sec-Ch-Ua-Arch":              "\"x86\"",
		"Sec-Ch-Ua-Bitness":           "\"64\"",
		"Sec-Ch-Ua-Full-Version":      "\"121.0.6167.161\"",
		"Sec-Ch-Ua-Full-Version-List": "\"Not A(Brand\";v=\"99.0.0.0\", \"Google Chrome\";v=\"121.0.6167.161\", \"Chromium\";v=\"121.0.6167.161\"",
	}

	req, _ := http.NewRequest("GET", target, nil)
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := auth.Session.Do(req)
	if err != nil {
		return NewError("part_two", 0, "Failed to make request", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode == 302 || resp.StatusCode == 200 {
		stateRegex := regexp.MustCompile(`state=(.*)`)
		stateMatch := stateRegex.FindStringSubmatch(string(body))
		if len(stateMatch) < 2 {
			return NewError("part_two", 0, "Could not find state in response", fmt.Errorf("error: Check details"))
		}

		state := strings.Split(stateMatch[1], `"`)[0]
		return auth.partThree(state)
	} else {
		return NewError("part_two", resp.StatusCode, string(body), fmt.Errorf("error: Check details"))

	}
}
func (auth *Authenticator) partThree(state string) *Error {
	logger.Log.Debug("Auth Three")

	target := fmt.Sprintf("https://auth0.openai.com/u/login/identifier?state=%s", state)
	emailURLEncoded := auth.URLEncode(auth.EmailAddress)

	payload := fmt.Sprintf(
		"state=%s&username=%s&js-available=false&webauthn-available=true&is-brave=false&webauthn-platform-available=true&action=default",
		state, emailURLEncoded,
	)

	headers := map[string]string{
		"Host":            "auth0.openai.com",
		"Origin":          "https://auth0.openai.com",
		"Connection":      "keep-alive",
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
		"User-Agent":      auth.UserAgent,
		"Referer":         fmt.Sprintf("https://auth0.openai.com/u/login/identifier?state=%s", state),
		"Accept-Language": "en-US,en;q=0.9",
		"Content-Type":    "application/x-www-form-urlencoded",
	}

	req, _ := http.NewRequest("POST", target, strings.NewReader(payload))

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := auth.Session.Do(req)
	if err != nil {
		return NewError("part_three", 0, "Failed to send request", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 302 || resp.StatusCode == 200 {
		return auth.partFour(state)
	} else {
		return NewError("part_three", resp.StatusCode, "Your email address is invalid.", fmt.Errorf("error: Check details"))

	}

}
func (auth *Authenticator) partFour(state string) *Error {
	logger.Log.Debug("Auth Four")

	target := fmt.Sprintf("https://auth0.openai.com/u/login/password?state=%s", state)
	emailURLEncoded := auth.URLEncode(auth.EmailAddress)
	passwordURLEncoded := auth.URLEncode(auth.Password)
	payload := fmt.Sprintf("state=%s&username=%s&password=%s", state, emailURLEncoded, passwordURLEncoded)

	headers := map[string]string{
		"Host":            "auth0.openai.com",
		"Origin":          "https://auth0.openai.com",
		"Connection":      "keep-alive",
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"User-Agent":      auth.UserAgent,
		"Referer":         fmt.Sprintf("https://auth0.openai.com/u/login/password?state=%s", state),
		"Accept-Language": "en-US,en;q=0.9",
		"Content-Type":    "application/x-www-form-urlencoded",
	}

	req, _ := http.NewRequest("POST", target, strings.NewReader(payload))

	for k, v := range headers {
		req.Header.Set(k, v)
	}
	token, err := funcaptcha.GetOpenAIArkoseToken(0, auth.PUID)
	if err != nil {
		return NewError("part_four", 0, "get arkose_token failed", err)
	}
	cookie := &http.Cookie{
		Name:  "arkoseToken",
		Value: token,
		Path:  "/",
	}
	req.AddCookie(cookie)
	resp, err := auth.Session.Do(req)
	if err != nil {
		return NewError("part_four", 0, "Failed to send request", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 302 {
		redirectURL := resp.Header.Get("Location")
		println(redirectURL)
		return auth.partFive(state, redirectURL)
	} else {
		var body interface{}
		if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
			return NewError("part_four", 0, "", err)
		}
		return NewError("part_four", resp.StatusCode, body.(string), fmt.Errorf("error: Check details"))

	}

}
func (auth *Authenticator) partFive(oldState string, redirectURL string) *Error {
	logger.Log.Debug("Auth Five")

	target := "https://auth0.openai.com" + redirectURL

	headers := map[string]string{
		"Host":            "auth0.openai.com",
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
		"Connection":      "keep-alive",
		"User-Agent":      auth.UserAgent,
		"Accept-Language": "en-GB,en-US;q=0.9,en;q=0.8",
		"Referer":         fmt.Sprintf("https://auth0.openai.com/u/login/password?state=%s", oldState),
	}

	req, _ := http.NewRequest("GET", target, nil)

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := auth.Session.Do(req)
	if err != nil {
		return NewError("part_five", 0, "Failed to send request", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 302 {
		return auth.partSix(resp.Header.Get("Location"), target)
	} else {
		return NewError("part_five", resp.StatusCode, resp.Status, fmt.Errorf("error: Check details"))

	}

}
func (auth *Authenticator) partSix(urls, redirect_url string) *Error {
	logger.Log.Debug("Auth Six")
	req, _ := http.NewRequest("GET", urls, nil)
	for k, v := range map[string]string{
		"Host":            "" + env.E.OpenaiHost + "",
		"Accept":          "application/json",
		"Connection":      "keep-alive",
		"User-Agent":      auth.UserAgent,
		"Accept-Language": "en-GB,en-US;q=0.9,en;q=0.8",
		"Referer":         redirect_url,
	} {
		req.Header.Set(k, v)
	}
	resp, err := auth.Session.Do(req)
	if err != nil {
		return NewError("part_six", 0, "Failed to send request", err)
	}
	defer resp.Body.Close()
	if err != nil {
		return NewError("part_six", 0, "Response was not JSON", err)
	}
	if resp.StatusCode != 302 {
		return NewError("part_six", resp.StatusCode, urls, fmt.Errorf("incorrect response code"))
	}
	// Check location header
	if location := resp.Header.Get("Location"); location != "https://"+env.E.OpenaiHost+"/" {
		return NewError("part_six", resp.StatusCode, location, fmt.Errorf("incorrect redirect"))
	}

	sessionUrl := "https://" + env.E.OpenaiHost + "/api/auth/session"

	req, _ = http.NewRequest("GET", sessionUrl, nil)

	// Set user agent
	req.Header.Set("User-Agent", auth.UserAgent)

	resp, err = auth.Session.Do(req)
	if err != nil {
		return NewError("get_access_token", 0, "Failed to send request", err)
	}

	if resp.StatusCode != 200 {
		return NewError("get_access_token", resp.StatusCode, "Incorrect response code", fmt.Errorf("error: Check details"))
	}
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return NewError("get_access_token", 0, "", err)
	}

	// Check if access token in data
	if _, ok := result["accessToken"]; !ok {
		resultString := fmt.Sprintf("%v", result)
		return NewError("part_six", 0, resultString, fmt.Errorf("missing access token"))
	}
	cookieUrl, _ := url.Parse("https://" + env.E.OpenaiHost + "")
	jar := auth.Session.GetCookies(cookieUrl)
	auth.AuthResult.AccessToken = result
	for _, cookie := range jar {
		if cookie.Name == "__Secure-next-auth.session-token" {
			auth.AuthResult.FreshToken = cookie.Value
		}
	}

	return nil
}

func (auth *Authenticator) GetAccessTokenByRefreshToken(freshToken string) *Error {
	logger.Log.Debug("GetAccessTokenByRefreshToken")
	sessionUrl := "https://" + env.E.OpenaiHost + "/api/auth/session"

	req, _ := http.NewRequest("GET", sessionUrl, nil)
	cookies := &http.Cookie{
		Name:  "__Secure-next-auth.session-token",
		Value: freshToken,
	}
	req.AddCookie(cookies)

	// Set user agent
	req.Header.Set("User-Agent", auth.UserAgent)

	resp, err := auth.Session.Do(req)
	if err != nil {
		return NewError("GetAccessTokenByRefreshToken", 0, "Failed to send request", err)
	}

	if resp.StatusCode != 200 {
		return NewError("GetAccessTokenByRefreshToken", resp.StatusCode, "Incorrect response code", fmt.Errorf("error: Check details"))
	}
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return NewError("GetAccessTokenByRefreshToken", 0, "", err)
	}

	// Check if access token in data
	if _, ok := result["accessToken"]; !ok {
		resultString := fmt.Sprintf("%v", result)
		return NewError("GetAccessTokenByRefreshToken", 0, resultString, fmt.Errorf("missing access token"))
	}
	cookieUrl, _ := url.Parse("https://" + env.E.OpenaiHost + "")
	jar := auth.Session.GetCookies(cookieUrl)
	auth.AuthResult.AccessToken = result
	for _, cookie := range jar {
		if cookie.Name == "__Secure-next-auth.session-token" {
			auth.AuthResult.FreshToken = cookie.Value
		}
	}
	return nil
}

func (auth *Authenticator) GetAccessToken() map[string]interface{} {
	logger.Log.Debug("GetAccessToken")
	return auth.AuthResult.AccessToken
}

func (auth *Authenticator) GetRefreshToken() string {
	logger.Log.Debug("GetRefreshToken")
	return auth.AuthResult.FreshToken
}
func (auth *Authenticator) GetModels() (map[string]interface{}, *Error) {
	logger.Log.Debug("GetModels")
	if len(auth.AuthResult.AccessToken) == 0 {
		return nil, NewError("get_model", 0, "Missing access token", fmt.Errorf("error: Check details"))
	}
	// Make request to https://"+common.E.OpenAI_HOST+"/backend-api/models
	req, _ := http.NewRequest("GET", "https://"+env.E.OpenaiHost+"/backend-api/models", nil)
	// Add headers
	req.Header.Add("Authorization", "Bearer "+auth.AuthResult.AccessToken["accessToken"].(string))
	req.Header.Add("User-Agent", auth.UserAgent)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9")
	req.Header.Add("Referer", "https://"+env.E.OpenaiHost+"/")
	req.Header.Add("Origin", "https://"+env.E.OpenaiHost+"")
	req.Header.Add("Connection", "keep-alive")

	resp, err := auth.Session.Do(req)
	if err != nil {
		return nil, NewError("get_model", 0, "Failed to make request", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, NewError("get_model", resp.StatusCode, "Failed to make request", fmt.Errorf("error: Check details"))
	}
	var responseBody map[string]interface{}
	r, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, NewError("get_model", resp.StatusCode, "Failed to get response", fmt.Errorf("error: Check details"))
	}
	if err := json.Unmarshal(r, &responseBody); err != nil {
		return nil, NewError("get_model", resp.StatusCode, "Failed to get response", fmt.Errorf("error: Check details"))
	}
	auth.AuthResult.Model = responseBody
	return responseBody, nil
}

func (auth *Authenticator) GetPUID() (string, *Error) {
	logger.Log.Debug("GetPUID")
	// Check if user has access token
	if len(auth.AuthResult.AccessToken) == 0 {
		return "", NewError("get_puid", 0, "Missing access token", fmt.Errorf("error: Check details"))
	}
	// Make request to https://"+common.E.OpenAI_HOST+"/backend-api/models
	req, _ := http.NewRequest("GET", "https://"+env.E.OpenaiHost+"/backend-api/models", nil)
	// Add headers
	req.Header.Add("Authorization", "Bearer "+auth.AuthResult.AccessToken["accessToken"].(string))
	req.Header.Add("User-Agent", auth.UserAgent)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9")
	req.Header.Add("Referer", "https://"+env.E.OpenaiHost+"/")
	req.Header.Add("Origin", "https://"+env.E.OpenaiHost+"")
	req.Header.Add("Connection", "keep-alive")

	resp, err := auth.Session.Do(req)
	if err != nil {
		return "", NewError("get_puid", 0, "Failed to make request", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", NewError("get_puid", resp.StatusCode, "Failed to make request", fmt.Errorf("error: Check details"))
	}
	// Find `_puid` cookie in response
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "_puid" {
			auth.AuthResult.PUID = cookie.Value
			return cookie.Value, nil
		}
	}
	// If cookie not found, return error
	return "", NewError("get_puid", 0, "PUID cookie not found", fmt.Errorf("error: Check details"))
}

func (auth *Authenticator) GetAuthResult() AuthResult {
	logger.Log.Debug("GetAuthResult")
	return auth.AuthResult
}
