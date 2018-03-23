package afas

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"strings"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-afas-profit-rest/" + libraryVersion
	mediaType      = "application/json"
	charset        = "utf-8"
)

var (
	BaseURL = url.URL{
		Scheme: "https",
		Host:   "{account_number}.afasonlineconnector.nl",
		Path:   "/profitrestservices/",
	}
)

// NewAPI returns a new AFAS API client
func NewAPI(httpClient *http.Client, accountNumber string, token string) *API {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	api := &API{
		http: httpClient,
	}

	api.SetAccountNumber(accountNumber)
	api.SetToken(token)
	api.SetBaseURL(BaseURL)
	api.SetDebug(false)
	api.SetUserAgent(userAgent)
	api.SetMediaType(mediaType)
	api.SetCharset(charset)

	// Services
	api.Meta = NewMetaService(api)
	api.Connector = NewConnectorService(api)

	// api.General = NewGeneralService(api)
	// api.CRM = NewCRMService(api)
	// api.Financial = NewFinancialService(api)
	// api.HR = NewHRService(api)
	// api.OrderManagement = NewOrderManagementService(api)
	// api.Projects = NewProjectsService(api)
	// api.Subscriptions = NewSubscriptionsService(api)

	return api
}

// API manages communication with AFAS API
type API struct {
	// HTTP client used to communicate with the API.
	http *http.Client

	debug   bool
	baseURL url.URL

	// credentials
	accountNumber string
	token         string

	// User agent for client
	userAgent string

	mediaType string
	charset   string

	// Optional function called after every successful request made to the DO APIs
	onRequestCompleted RequestCompletionCallback

	// Services used for communicating with the API
	Meta      *MetaService
	Connector *ConnectorService

	// General         *GeneralService
	// CRM *CRMService
	// Financial       *FinancialService
	// HR              *HRService
	// OrderManagement *OrdermanagementService
	// Projects        *ProjectsService
	// Subscriptions   *SubscriptionsService
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

func (api *API) Debug() bool {
	return api.debug
}

func (api *API) SetDebug(debug bool) {
	api.debug = debug
}

func (api *API) AccountNumber() string {
	return api.accountNumber
}

func (api *API) SetAccountNumber(accountNumber string) {
	api.accountNumber = accountNumber
}

func (api *API) Token() string {
	return api.token
}

func (api *API) SetToken(token string) {
	// Normalize token
	if !strings.Contains(token, "<token>") {
		token = fmt.Sprintf("<token><version>1</version><data>%s</data></token>", token)
	}
	api.token = token
}

func (api *API) BaseURL() url.URL {
	return api.baseURL
}

func (api *API) SetBaseURL(baseURL url.URL) {
	api.baseURL = baseURL
}

func (api *API) SetMediaType(mediaType string) {
	api.mediaType = mediaType
}

func (api *API) MediaType() string {
	return mediaType
}

func (api *API) SetCharset(charset string) {
	api.charset = charset
}

func (api *API) Charset() string {
	return charset
}

func (api *API) SetUserAgent(userAgent string) {
	api.userAgent = userAgent
}

func (api *API) UserAgent() string {
	return userAgent
}

func (api *API) GetEndpointURL(p string) url.URL {
	apiURL := api.BaseURL()
	apiURL.Host = strings.Replace(apiURL.Host, "{account_number}", api.AccountNumber(), 1)
	apiURL.Path = path.Join(apiURL.Path, p)
	return apiURL
}

func (api *API) NewRequest(ctx context.Context, method string, URL url.URL, body interface{}) (*http.Request, error) {
	// convert body struct to json
	buf := new(bytes.Buffer)
	_, empty := body.(*EmptyRequestBody)
	if body != nil && !empty {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	// create new http request
	req, err := http.NewRequest(method, URL.String(), buf)
	if err != nil {
		return nil, err
	}

	// optionally pass along context
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	// Set token header
	b64Token := base64.StdEncoding.EncodeToString([]byte(api.Token()))
	req.Header.Add("Authorization", fmt.Sprintf("AfasToken %s", b64Token))

	// set other headers
	req.Header.Add("Content-Type", fmt.Sprintf("%s; charset=%s", api.MediaType(), api.Charset()))
	req.Header.Add("Accept", api.MediaType())
	req.Header.Add("User-Agent", api.UserAgent())

	return req, nil
}

// Do sends an API request and returns the API response. The API response is json decoded and stored in the value
// pointed to by v, or returned as an error if an API error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (api *API) Do(req *http.Request, responseBody interface{}) (*http.Response, error) {
	if api.debug == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	httpResp, err := api.http.Do(req)
	if err != nil {
		return nil, err
	}

	if api.onRequestCompleted != nil {
		api.onRequestCompleted(req, httpResp)
	}

	// close body io.Reader
	defer func() {
		if rerr := httpResp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if api.debug == true {
		dump, _ := httputil.DumpResponse(httpResp, true)
		log.Println(string(dump))
	}

	// check if the response isn't an error
	err = CheckResponse(httpResp)
	if err != nil {
		return httpResp, err
	}

	// check the provided interface parameter
	if httpResp == nil {
		return httpResp, nil
	}

	// interface implements io.Writer: write Body to it
	// if w, ok := response.Envelope.(io.Writer); ok {
	// 	_, err := io.Copy(w, httpResp.Body)
	// 	return httpResp, err
	// }

	// try to decode body into interface parameter
	if responseBody != nil {
		err = json.NewDecoder(httpResp.Body).Decode(responseBody)
		if err != nil && err != io.EOF {
			// create a simple error response
			errorResponse := &ErrorResponse1{Response: httpResp, Message: err.Error()}
			return httpResp, errorResponse
		}
	}

	return httpResp, nil
}

// CheckResponse checks the API response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range. API error responses are expected to have either no response
// body, or a XML response body that maps to ErrorResponse. Any other response
// body will be silently ignored.
func CheckResponse(r *http.Response) error {
	err := checkContentType(r)
	if err != nil {
		return &ErrorResponse1{Response: r, Message: err.Error()}
	}

	// Don't check content-lenght: a created response, for example, has no body
	if r.Header.Get("Content-Length") == "0" {
		return &ErrorResponse1{Response: r, Message: "No content in response body"}
	}

	// If the statuscode is ok: don't check for errors
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	// read data and copy it back
	data, err := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewReader(data))
	if err != nil {
		return err
	}

	if len(data) == 0 {
		return errors.New("No content in response body")
	}

	// check which type of errorresponse is sent
	err1 := ErrorResponse1{Response: r}
	err = json.Unmarshal(data, &err1)
	if err1.Message != "" {
		return err1
	}

	err2 := ErrorResponse2{Response: r}
	err = json.Unmarshal(data, &err2)
	if err2.ExternalMessage != "" {
		return err2
	}

	return &ErrorResponse1{Response: r, Message: "Unknown error"}
}

// {
// 	"message":"Invalid json"
// }
type ErrorResponse1 struct {
	// HTTP response that caused this error
	Response *http.Response `json:"-"`

	Message string `json:"message"`
}

func (r ErrorResponse1) Error() string {
	return fmt.Sprintf("%d: %s", r.Response.StatusCode, r.Message)
}

// {
//   "errorNumber": -2147180996,
//   "externalMessage": "General message: Deze UpdateConnector wordt niet ondersteund of de gebruiker is niet geautoriseerd.",
//   "profitLogReference": "923A9F6F4D5875E0D7B96982F6E4E0D2"
// }
type ErrorResponse2 struct {
	// HTTP response that caused this error
	Response *http.Response `json:"-"`

	ErrorNumber        int    `json:"errorNumber"`
	ExternalMessage    string `json:"externalMessage"`
	ProfitLogReference string `json:"profitLogReference"`
}

func (r ErrorResponse2) Error() string {
	return fmt.Sprintf("%d: %s", r.Response.StatusCode, r.ExternalMessage)
}

func checkContentType(response *http.Response) error {
	header := response.Header.Get("Content-Type")
	contentType := strings.Split(header, ";")[0]
	if contentType != mediaType {
		return fmt.Errorf("Expected Content-Type \"%s\", got \"%s\"", mediaType, contentType)
	}

	return nil
}

type EmptyRequestBody struct{}

func (r EmptyRequestBody) MarshalJSON() ([]byte, error) {
	return []byte{}, nil
}
