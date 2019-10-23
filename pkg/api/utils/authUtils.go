package utils

import (
	"net/http"
	"strings"

	"github.com/keptn/go-utils/pkg/api/models"
)

// AuthHandler handles projects
type AuthHandler struct {
	BaseURL    string
	AuthToken  string
	AuthHeader string
	HTTPClient *http.Client
	Scheme     string
}

// NewAuthHandler returns a new AuthHandler
func NewAuthHandler(baseURL string) *AuthHandler {
	baseURL = strings.TrimPrefix(baseURL, "http://")
	baseURL = strings.TrimPrefix(baseURL, "https://")
	return &AuthHandler{
		BaseURL:    baseURL,
		AuthHeader: "",
		AuthToken:  "",
		HTTPClient: &http.Client{},
		Scheme:     "https",
	}
}

// NewAuthenticatedAuthHandler returns a new AuthHandler that authenticates at the endpoint via the provided token
func NewAuthenticatedAuthHandler(baseURL string, authToken string, authHeader string, httpClient *http.Client, scheme string) *AuthHandler {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL = strings.TrimPrefix(baseURL, "http://")
	baseURL = strings.TrimPrefix(baseURL, "https://")
	return &AuthHandler{
		BaseURL:    baseURL,
		AuthHeader: authHeader,
		AuthToken:  authToken,
		HTTPClient: httpClient,
		Scheme:     scheme,
	}
}

func (a *AuthHandler) getBaseURL() string {
	return a.BaseURL
}

func (a *AuthHandler) getAuthToken() string {
	return a.AuthToken
}

func (a *AuthHandler) getAuthHeader() string {
	return a.AuthHeader
}

func (a *AuthHandler) getHTTPClient() *http.Client {
	return a.HTTPClient
}

// Authenticate authenticates the client request against the server
func (a *AuthHandler) Authenticate() (*models.EventContext, *models.Error) {
	return post(a.Scheme+"://"+a.getBaseURL()+"/v1/auth", nil, a)
}