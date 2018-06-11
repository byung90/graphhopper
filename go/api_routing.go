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

type RoutingApiService service

/* RoutingApiService Routing Request
The GraphHopper Routing API allows to calculate route and implement navigation via the turn instructions
 * @param ctx context.Context for authentication, logging, tracing, etc.
@param point Specify multiple points for which the route should be calculated. The order is important. Specify at least two points.
@param pointsEncoded IMPORTANT- TODO - currently you have to pass false for the swagger client - Have not found a way to force add a parameter. If &#x60;false&#x60; the coordinates in &#x60;point&#x60; and &#x60;snapped_waypoints&#x60; are returned as array using the order [lon,lat,elevation] for every point. If &#x60;true&#x60; the coordinates will be encoded as string leading to less bandwith usage. You&#39;ll need a special handling for the decoding of this string on the client-side. We provide open source code in [Java](https://github.com/graphhopper/graphhopper/blob/d70b63660ac5200b03c38ba3406b8f93976628a6/web/src/main/java/com/graphhopper/http/WebHelper.java#L43) and [JavaScript](https://github.com/graphhopper/graphhopper/blob/d70b63660ac5200b03c38ba3406b8f93976628a6/web/src/main/webapp/js/ghrequest.js#L139). It is especially important to use no 3rd party client if you set &#x60;elevation&#x3D;true&#x60;!
@param key Get your key at graphhopper.com
@param optional (nil or map[string]interface{}) with one or more of:
    @param "locale" (string) The locale of the resulting turn instructions. E.g. &#x60;pt_PT&#x60; for Portuguese or &#x60;de&#x60; for German
    @param "instructions" (bool) If instruction should be calculated and returned
    @param "vehicle" (string) The vehicle for which the route should be calculated. Other vehicles are foot, small_truck, ...
    @param "elevation" (bool) If &#x60;true&#x60; a third dimension - the elevation - is included in the polyline or in the GeoJson. If enabled you have to use a modified version of the decoding method or set points_encoded to &#x60;false&#x60;. See the points_encoded attribute for more details. Additionally a request can fail if the vehicle does not support elevation. See the features object for every vehicle.
    @param "calcPoints" (bool) If the points for the route should be calculated at all printing out only distance and time.
    @param "pointHint" ([]string) Optional parameter. Specifies a hint for each &#x60;point&#x60; parameter to prefer a certain street for the closest location lookup. E.g. if there is an address or house with two or more neighboring streets you can control for which street the closest location is looked up.
    @param "chDisable" (bool) Use this parameter in combination with one or more parameters of this table
    @param "weighting" (string) Which kind of &#39;best&#39; route calculation you need. Other option is &#x60;shortest&#x60; (e.g. for &#x60;vehicle&#x3D;foot&#x60; or &#x60;bike&#x60;), &#x60;short_fastest&#x60; if time and distance is expensive e.g. for &#x60;vehicle&#x3D;truck&#x60;
    @param "edgeTraversal" (bool) Use &#x60;true&#x60; if you want to consider turn restrictions for bike and motor vehicles. Keep in mind that the response time is roughly 2 times slower.
    @param "algorithm" (string) The algorithm to calculate the route. Other options are &#x60;dijkstra&#x60;, &#x60;astar&#x60;, &#x60;astarbi&#x60;, &#x60;alternative_route&#x60; and &#x60;round_trip&#x60;
    @param "heading" (int32) Favour a heading direction for a certain point. Specify either one heading for the start point or as many as there are points. In this case headings are associated by their order to the specific points. Headings are given as north based clockwise angle between 0 and 360 degree. This parameter also influences the tour generated with &#x60;algorithm&#x3D;round_trip&#x60; and force the initial direction.
    @param "headingPenalty" (int32) Penalty for omitting a specified heading. The penalty corresponds to the accepted time delay in seconds in comparison to the route without a heading.
    @param "passThrough" (bool) If &#x60;true&#x60; u-turns are avoided at via-points with regard to the &#x60;heading_penalty&#x60;.
    @param "details" ([]string) List of additional trip attributes to be returned. Try some of the following: &#x60;average_speed&#x60;, &#x60;street_name&#x60;, &#x60;edge_id&#x60;, &#x60;time&#x60;, &#x60;distance&#x60;.
    @param "roundTripDistance" (int32) If &#x60;algorithm&#x3D;round_trip&#x60; this parameter configures approximative length of the resulting round trip
    @param "roundTripSeed" (int64) If &#x60;algorithm&#x3D;round_trip&#x60; this parameter introduces randomness if e.g. the first try wasn&#39;t good.
    @param "alternativeRouteMaxPaths" (int32) If &#x60;algorithm&#x3D;alternative_route&#x60; this parameter sets the number of maximum paths which should be calculated. Increasing can lead to worse alternatives.
    @param "alternativeRouteMaxWeightFactor" (int32) If &#x60;algorithm&#x3D;alternative_route&#x60; this parameter sets the factor by which the alternatives routes can be longer than the optimal route. Increasing can lead to worse alternatives.
    @param "alternativeRouteMaxShareFactor" (int32) If &#x60;algorithm&#x3D;alternative_route&#x60; this parameter specifies how much alternatives routes can have maximum in common with the optimal route. Increasing can lead to worse alternatives.
    @param "avoid" (string) comma separate list to avoid certain roads. You can avoid motorway, ferry, tunnel or track
@return RouteResponse*/
func (a *RoutingApiService) RouteGet(ctx context.Context, point []string, pointsEncoded bool, key string, localVarOptionals map[string]interface{}) (RouteResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     RouteResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/route"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if err := typeCheckParameter(localVarOptionals["locale"], "string", "locale"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["instructions"], "bool", "instructions"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["vehicle"], "string", "vehicle"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["elevation"], "bool", "elevation"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["calcPoints"], "bool", "calcPoints"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["chDisable"], "bool", "chDisable"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["weighting"], "string", "weighting"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["edgeTraversal"], "bool", "edgeTraversal"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["algorithm"], "string", "algorithm"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["heading"], "int32", "heading"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["headingPenalty"], "int32", "headingPenalty"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["passThrough"], "bool", "passThrough"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["roundTripDistance"], "int32", "roundTripDistance"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["roundTripSeed"], "int64", "roundTripSeed"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["alternativeRouteMaxPaths"], "int32", "alternativeRouteMaxPaths"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["alternativeRouteMaxWeightFactor"], "int32", "alternativeRouteMaxWeightFactor"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["alternativeRouteMaxShareFactor"], "int32", "alternativeRouteMaxShareFactor"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["avoid"], "string", "avoid"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("point", parameterToString(point, "multi"))
	if localVarTempParam, localVarOk := localVarOptionals["locale"].(string); localVarOk {
		localVarQueryParams.Add("locale", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["instructions"].(bool); localVarOk {
		localVarQueryParams.Add("instructions", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["vehicle"].(string); localVarOk {
		localVarQueryParams.Add("vehicle", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["elevation"].(bool); localVarOk {
		localVarQueryParams.Add("elevation", parameterToString(localVarTempParam, ""))
	}
	localVarQueryParams.Add("points_encoded", parameterToString(pointsEncoded, ""))
	if localVarTempParam, localVarOk := localVarOptionals["calcPoints"].(bool); localVarOk {
		localVarQueryParams.Add("calc_points", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["pointHint"].([]string); localVarOk {
		localVarQueryParams.Add("point_hint", parameterToString(localVarTempParam, "multi"))
	}
	if localVarTempParam, localVarOk := localVarOptionals["chDisable"].(bool); localVarOk {
		localVarQueryParams.Add("ch.disable", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["weighting"].(string); localVarOk {
		localVarQueryParams.Add("weighting", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["edgeTraversal"].(bool); localVarOk {
		localVarQueryParams.Add("edge_traversal", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["algorithm"].(string); localVarOk {
		localVarQueryParams.Add("algorithm", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["heading"].(int32); localVarOk {
		localVarQueryParams.Add("heading", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["headingPenalty"].(int32); localVarOk {
		localVarQueryParams.Add("heading_penalty", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["passThrough"].(bool); localVarOk {
		localVarQueryParams.Add("pass_through", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["details"].([]string); localVarOk {
		localVarQueryParams.Add("details", parameterToString(localVarTempParam, "multi"))
	}
	if localVarTempParam, localVarOk := localVarOptionals["roundTripDistance"].(int32); localVarOk {
		localVarQueryParams.Add("round_trip.distance", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["roundTripSeed"].(int64); localVarOk {
		localVarQueryParams.Add("round_trip.seed", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["alternativeRouteMaxPaths"].(int32); localVarOk {
		localVarQueryParams.Add("alternative_route.max_paths", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["alternativeRouteMaxWeightFactor"].(int32); localVarOk {
		localVarQueryParams.Add("alternative_route.max_weight_factor", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["alternativeRouteMaxShareFactor"].(int32); localVarOk {
		localVarQueryParams.Add("alternative_route.max_share_factor", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["avoid"].(string); localVarOk {
		localVarQueryParams.Add("avoid", parameterToString(localVarTempParam, ""))
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
