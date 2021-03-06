/*
 * GraphHopper Directions API
 *
 * You use the GraphHopper Directions API to add route planning, navigation and route optimization to your software. E.g. the Routing API has turn instructions and elevation data and the Route Optimization API solves your logistic problems and supports various constraints like time window and capacity restrictions. Also it is possible to get all distances between all locations with our fast Matrix API.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package graphhopper

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"golang.org/x/net/context"
	"encoding/json"
)

// Linger please
var (
	_ context.Context
)

type GeocodingApiService service

/* GeocodingApiService Execute a Geocoding request
This endpoint provides forward and reverse geocoding. For more details, review the official documentation at: https://graphhopper.com/api/1/docs/geocoding/ 
 * @param ctx context.Context for authentication, logging, tracing, etc.
@param key Get your key at graphhopper.com
@param optional (nil or map[string]interface{}) with one or more of:
    @param "q" (string) If you do forward geocoding, then this would be a textual description of the address you are looking for
    @param "locale" (string) Display the search results for the specified locale. Currently French (fr), English (en), German (de) and Italian (it) are supported. If the locale wasn&#39;t found the default (en) is used.
    @param "limit" (int32) Specify the maximum number of returned results
    @param "reverse" (bool) Set to true to do a reverse Geocoding request, see point parameter
    @param "point" (string) The location bias in the format &#39;latitude,longitude&#39; e.g. point&#x3D;45.93272,11.58803
    @param "provider" (string) Can be either, default, nominatim, opencagedata
@return GeocodingResponse*/
func (a *GeocodingApiService) GeocodeGet(ctx context.Context, key string, localVarOptionals map[string]interface{}) (GeocodingResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     GeocodingResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/geocode"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if err := typeCheckParameter(localVarOptionals["q"], "string", "q"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["locale"], "string", "locale"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["limit"], "int32", "limit"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["reverse"], "bool", "reverse"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["point"], "string", "point"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["provider"], "string", "provider"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["q"].(string); localVarOk {
		localVarQueryParams.Add("q", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["locale"].(string); localVarOk {
		localVarQueryParams.Add("locale", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["limit"].(int32); localVarOk {
		localVarQueryParams.Add("limit", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["reverse"].(bool); localVarOk {
		localVarQueryParams.Add("reverse", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["point"].(string); localVarOk {
		localVarQueryParams.Add("point", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["provider"].(string); localVarOk {
		localVarQueryParams.Add("provider", parameterToString(localVarTempParam, ""))
	}
	localVarQueryParams.Add("key", parameterToString(key, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}
