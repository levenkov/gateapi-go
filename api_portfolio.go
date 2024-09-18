/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

import (
	"context"
	"github.com/antihax/optional"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Linger please
var (
	_ context.Context
)

// PortfolioApiService PortfolioApi service
type PortfolioApiService service

// ListPortfolioAccountsOpts Optional parameters for the method 'ListPortfolioAccounts'
type ListPortfolioAccountsOpts struct {
	Currency optional.String
}

/*
ListPortfolioAccounts Get portfolio account information
The assets of each currency in the account will be adjusted according to their liquidity, defined by corresponding adjustment coefficients, and then uniformly converted to USD to calculate the total asset value and position value of the account.  You can refer to the [Formula](#portfolio-account) in the documentation
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param optional nil or *ListPortfolioAccountsOpts - Optional Parameters:
  - @param "Currency" (optional.String) -  Retrieve data of the specified currency

@return PortfolioAccount
*/
func (a *PortfolioApiService) ListPortfolioAccounts(ctx context.Context, localVarOptionals *ListPortfolioAccountsOpts) (PortfolioAccount, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PortfolioAccount
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/portfolio/accounts"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Currency.IsSet() {
		localVarQueryParams.Add("currency", parameterToString(localVarOptionals.Currency.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if ctx.Value(ContextGateAPIV4) == nil {
		// for compatibility, set configuration key and secret to context if ContextGateAPIV4 value is not present
		ctx = context.WithValue(ctx, ContextGateAPIV4, GateAPIV4{
			Key:    a.client.cfg.Key,
			Secret: a.client.cfg.Secret,
		})
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status + ", " + string(localVarBody),
		}
		var gateErr GateAPIError
		if e := a.client.decode(&gateErr, localVarBody, localVarHTTPResponse.Header.Get("Content-Type")); e == nil && gateErr.Label != "" {
			gateErr.APIError = newErr
			return localVarReturnValue, localVarHTTPResponse, gateErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
ListAccountPortfolioMode Retrieve the account's portfolio mode
cross_margin - cross margin
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return map[string]bool
*/
func (a *PortfolioApiService) ListAccountPortfolioMode(ctx context.Context) (map[string]bool, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]bool
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/portfolio/account_mode"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if ctx.Value(ContextGateAPIV4) == nil {
		// for compatibility, set configuration key and secret to context if ContextGateAPIV4 value is not present
		ctx = context.WithValue(ctx, ContextGateAPIV4, GateAPIV4{
			Key:    a.client.cfg.Key,
			Secret: a.client.cfg.Secret,
		})
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status + ", " + string(localVarBody),
		}
		var gateErr GateAPIError
		if e := a.client.decode(&gateErr, localVarBody, localVarHTTPResponse.Header.Get("Content-Type")); e == nil && gateErr.Label != "" {
			gateErr.APIError = newErr
			return localVarReturnValue, localVarHTTPResponse, gateErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
SetAccountPortfolioMode Configure the account's portfolio mode.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param portfolioMode

@return map[string]bool
*/
func (a *PortfolioApiService) SetAccountPortfolioMode(ctx context.Context, portfolioMode PortfolioMode) (map[string]bool, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]bool
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/portfolio/account_mode"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = &portfolioMode
	if ctx == nil {
		ctx = context.Background()
	}
	if ctx.Value(ContextGateAPIV4) == nil {
		// for compatibility, set configuration key and secret to context if ContextGateAPIV4 value is not present
		ctx = context.WithValue(ctx, ContextGateAPIV4, GateAPIV4{
			Key:    a.client.cfg.Key,
			Secret: a.client.cfg.Secret,
		})
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status + ", " + string(localVarBody),
		}
		var gateErr GateAPIError
		if e := a.client.decode(&gateErr, localVarBody, localVarHTTPResponse.Header.Get("Content-Type")); e == nil && gateErr.Label != "" {
			gateErr.APIError = newErr
			return localVarReturnValue, localVarHTTPResponse, gateErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
GetPortfolioBorrowable Retrieve the maximum borrowable amount for the account.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param currency Retrieve data of the specified currency

@return PortfolioBorrowable
*/
func (a *PortfolioApiService) GetPortfolioBorrowable(ctx context.Context, currency string) (PortfolioBorrowable, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PortfolioBorrowable
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/portfolio/borrowable"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("currency", parameterToString(currency, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if ctx.Value(ContextGateAPIV4) == nil {
		// for compatibility, set configuration key and secret to context if ContextGateAPIV4 value is not present
		ctx = context.WithValue(ctx, ContextGateAPIV4, GateAPIV4{
			Key:    a.client.cfg.Key,
			Secret: a.client.cfg.Secret,
		})
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status + ", " + string(localVarBody),
		}
		var gateErr GateAPIError
		if e := a.client.decode(&gateErr, localVarBody, localVarHTTPResponse.Header.Get("Content-Type")); e == nil && gateErr.Label != "" {
			gateErr.APIError = newErr
			return localVarReturnValue, localVarHTTPResponse, gateErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
GetPortfolioTransferable Retrieve the maximum amount that can be transferred out from the account
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param currency Retrieve data of the specified currency

@return PortfolioTransferable
*/
func (a *PortfolioApiService) GetPortfolioTransferable(ctx context.Context, currency string) (PortfolioTransferable, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PortfolioTransferable
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/portfolio/transferable"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("currency", parameterToString(currency, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if ctx.Value(ContextGateAPIV4) == nil {
		// for compatibility, set configuration key and secret to context if ContextGateAPIV4 value is not present
		ctx = context.WithValue(ctx, ContextGateAPIV4, GateAPIV4{
			Key:    a.client.cfg.Key,
			Secret: a.client.cfg.Secret,
		})
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status + ", " + string(localVarBody),
		}
		var gateErr GateAPIError
		if e := a.client.decode(&gateErr, localVarBody, localVarHTTPResponse.Header.Get("Content-Type")); e == nil && gateErr.Label != "" {
			gateErr.APIError = newErr
			return localVarReturnValue, localVarHTTPResponse, gateErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// ListPortfolioUniLoanInterestRecordsOpts Optional parameters for the method 'ListPortfolioUniLoanInterestRecords'
type ListPortfolioUniLoanInterestRecordsOpts struct {
	Currency optional.String
	Page     optional.Int32
	Limit    optional.Int32
}

/*
ListPortfolioUniLoanInterestRecords List loans
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param optional nil or *ListPortfolioUniLoanInterestRecordsOpts - Optional Parameters:
  - @param "Currency" (optional.String) -  Retrieve data of the specified currency
  - @param "Page" (optional.Int32) -  Page number
  - @param "Limit" (optional.Int32) -  Maximum response items.  Default: 100, minimum: 1, Maximum: 100

@return []UniLoan
*/
func (a *PortfolioApiService) ListPortfolioUniLoanInterestRecords(ctx context.Context, localVarOptionals *ListPortfolioUniLoanInterestRecordsOpts) ([]UniLoan, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []UniLoan
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/portfolio/loans"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Currency.IsSet() {
		localVarQueryParams.Add("currency", parameterToString(localVarOptionals.Currency.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if ctx.Value(ContextGateAPIV4) == nil {
		// for compatibility, set configuration key and secret to context if ContextGateAPIV4 value is not present
		ctx = context.WithValue(ctx, ContextGateAPIV4, GateAPIV4{
			Key:    a.client.cfg.Key,
			Secret: a.client.cfg.Secret,
		})
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status + ", " + string(localVarBody),
		}
		var gateErr GateAPIError
		if e := a.client.decode(&gateErr, localVarBody, localVarHTTPResponse.Header.Get("Content-Type")); e == nil && gateErr.Label != "" {
			gateErr.APIError = newErr
			return localVarReturnValue, localVarHTTPResponse, gateErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
CreatePortfolioLoan Borrow or repay
When borrowing, it is essential to ensure that the borrowed amount is not below the minimum borrowing threshold for the specific cryptocurrency and does not exceed the maximum borrowing limit set by the platform and the user.  The interest on the loan will be automatically deducted from the account at regular intervals. It is the user&#39;s responsibility to manage the repayment of the borrowed amount.  For repayment, the option to repay the entire borrowed amount is available by setting the parameter &#x60;repaid_all&#x3D;true&#x60;
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param portfolioLoan
*/
func (a *PortfolioApiService) CreatePortfolioLoan(ctx context.Context, portfolioLoan PortfolioLoan) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/portfolio/loans"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = &portfolioLoan
	if ctx == nil {
		ctx = context.Background()
	}
	if ctx.Value(ContextGateAPIV4) == nil {
		// for compatibility, set configuration key and secret to context if ContextGateAPIV4 value is not present
		ctx = context.WithValue(ctx, ContextGateAPIV4, GateAPIV4{
			Key:    a.client.cfg.Key,
			Secret: a.client.cfg.Secret,
		})
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status + ", " + string(localVarBody),
		}
		var gateErr GateAPIError
		if e := a.client.decode(&gateErr, localVarBody, localVarHTTPResponse.Header.Get("Content-Type")); e == nil && gateErr.Label != "" {
			gateErr.APIError = newErr
			return localVarHTTPResponse, gateErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// ListPortfolioLoanRecordsOpts Optional parameters for the method 'ListPortfolioLoanRecords'
type ListPortfolioLoanRecordsOpts struct {
	Type_    optional.String
	Currency optional.String
	Page     optional.Int32
	Limit    optional.Int32
}

/*
ListPortfolioLoanRecords Get load records
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param optional nil or *ListPortfolioLoanRecordsOpts - Optional Parameters:
  - @param "Type_" (optional.String) -  The types of lending records, borrow - indicates the action of borrowing funds, repay - indicates the action of repaying the borrowed funds
  - @param "Currency" (optional.String) -  Retrieve data of the specified currency
  - @param "Page" (optional.Int32) -  Page number
  - @param "Limit" (optional.Int32) -  Maximum response items.  Default: 100, minimum: 1, Maximum: 100

@return []PortfolioLoanRecord
*/
func (a *PortfolioApiService) ListPortfolioLoanRecords(ctx context.Context, localVarOptionals *ListPortfolioLoanRecordsOpts) ([]PortfolioLoanRecord, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []PortfolioLoanRecord
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/portfolio/loan_records"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Currency.IsSet() {
		localVarQueryParams.Add("currency", parameterToString(localVarOptionals.Currency.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if ctx.Value(ContextGateAPIV4) == nil {
		// for compatibility, set configuration key and secret to context if ContextGateAPIV4 value is not present
		ctx = context.WithValue(ctx, ContextGateAPIV4, GateAPIV4{
			Key:    a.client.cfg.Key,
			Secret: a.client.cfg.Secret,
		})
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status + ", " + string(localVarBody),
		}
		var gateErr GateAPIError
		if e := a.client.decode(&gateErr, localVarBody, localVarHTTPResponse.Header.Get("Content-Type")); e == nil && gateErr.Label != "" {
			gateErr.APIError = newErr
			return localVarReturnValue, localVarHTTPResponse, gateErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// ListPortfolioLoanInterestRecordsOpts Optional parameters for the method 'ListPortfolioLoanInterestRecords'
type ListPortfolioLoanInterestRecordsOpts struct {
	Currency optional.String
	Page     optional.Int32
	Limit    optional.Int32
}

/*
ListPortfolioLoanInterestRecords List interest records
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param optional nil or *ListPortfolioLoanInterestRecordsOpts - Optional Parameters:
  - @param "Currency" (optional.String) -  Retrieve data of the specified currency
  - @param "Page" (optional.Int32) -  Page number
  - @param "Limit" (optional.Int32) -  Maximum response items.  Default: 100, minimum: 1, Maximum: 100

@return []UniLoanInterestRecord
*/
func (a *PortfolioApiService) ListPortfolioLoanInterestRecords(ctx context.Context, localVarOptionals *ListPortfolioLoanInterestRecordsOpts) ([]UniLoanInterestRecord, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []UniLoanInterestRecord
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/portfolio/interest_records"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Currency.IsSet() {
		localVarQueryParams.Add("currency", parameterToString(localVarOptionals.Currency.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if ctx.Value(ContextGateAPIV4) == nil {
		// for compatibility, set configuration key and secret to context if ContextGateAPIV4 value is not present
		ctx = context.WithValue(ctx, ContextGateAPIV4, GateAPIV4{
			Key:    a.client.cfg.Key,
			Secret: a.client.cfg.Secret,
		})
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status + ", " + string(localVarBody),
		}
		var gateErr GateAPIError
		if e := a.client.decode(&gateErr, localVarBody, localVarHTTPResponse.Header.Get("Content-Type")); e == nil && gateErr.Label != "" {
			gateErr.APIError = newErr
			return localVarReturnValue, localVarHTTPResponse, gateErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}