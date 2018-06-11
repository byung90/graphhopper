/* 
 * GraphHopper Directions API
 *
 * You use the GraphHopper Directions API to add route planning, navigation and route optimization to your software. E.g. the Routing API has turn instructions and elevation data and the Route Optimization API solves your logistic problems and supports various constraints like time window and capacity restrictions. Also it is possible to get all distances between all locations with our fast Matrix API.
 *
 * OpenAPI spec version: 1.0.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

use std::rc::Rc;
use std::borrow::Borrow;
use std::borrow::Cow;

use hyper;
use serde_json;
use futures;
use futures::{Future, Stream};

use hyper::header::UserAgent;

use super::{Error, configuration};

pub struct RoutingApiClient<C: hyper::client::Connect> {
    configuration: Rc<configuration::Configuration<C>>,
}

impl<C: hyper::client::Connect> RoutingApiClient<C> {
    pub fn new(configuration: Rc<configuration::Configuration<C>>) -> RoutingApiClient<C> {
        RoutingApiClient {
            configuration: configuration,
        }
    }
}

pub trait RoutingApi {
    fn route_get(&self, point: Vec<String>, points_encoded: bool, key: &str, locale: &str, instructions: bool, vehicle: &str, elevation: bool, calc_points: bool, point_hint: Vec<String>, ch_disable: bool, weighting: &str, edge_traversal: bool, algorithm: &str, heading: i32, heading_penalty: i32, pass_through: bool, details: Vec<String>, round_trip_distance: i32, round_trip_seed: i64, alternative_route_max_paths: i32, alternative_route_max_weight_factor: i32, alternative_route_max_share_factor: i32, avoid: &str) -> Box<Future<Item = ::models::RouteResponse, Error = Error<serde_json::Value>>>;
}


impl<C: hyper::client::Connect>RoutingApi for RoutingApiClient<C> {
    fn route_get(&self, point: Vec<String>, points_encoded: bool, key: &str, locale: &str, instructions: bool, vehicle: &str, elevation: bool, calc_points: bool, point_hint: Vec<String>, ch_disable: bool, weighting: &str, edge_traversal: bool, algorithm: &str, heading: i32, heading_penalty: i32, pass_through: bool, details: Vec<String>, round_trip_distance: i32, round_trip_seed: i64, alternative_route_max_paths: i32, alternative_route_max_weight_factor: i32, alternative_route_max_share_factor: i32, avoid: &str) -> Box<Future<Item = ::models::RouteResponse, Error = Error<serde_json::Value>>> {
        let configuration: &configuration::Configuration<C> = self.configuration.borrow();

        let method = hyper::Method::Get;

        let query = ::url::form_urlencoded::Serializer::new(String::new())
            .append_pair("point", &point.join(",").to_string())
            .append_pair("locale", &locale.to_string())
            .append_pair("instructions", &instructions.to_string())
            .append_pair("vehicle", &vehicle.to_string())
            .append_pair("elevation", &elevation.to_string())
            .append_pair("points_encoded", &points_encoded.to_string())
            .append_pair("calc_points", &calc_points.to_string())
            .append_pair("point_hint", &point_hint.join(",").to_string())
            .append_pair("ch.disable", &ch_disable.to_string())
            .append_pair("weighting", &weighting.to_string())
            .append_pair("edge_traversal", &edge_traversal.to_string())
            .append_pair("algorithm", &algorithm.to_string())
            .append_pair("heading", &heading.to_string())
            .append_pair("heading_penalty", &heading_penalty.to_string())
            .append_pair("pass_through", &pass_through.to_string())
            .append_pair("details", &details.join(",").to_string())
            .append_pair("round_trip.distance", &round_trip_distance.to_string())
            .append_pair("round_trip.seed", &round_trip_seed.to_string())
            .append_pair("alternative_route.max_paths", &alternative_route_max_paths.to_string())
            .append_pair("alternative_route.max_weight_factor", &alternative_route_max_weight_factor.to_string())
            .append_pair("alternative_route.max_share_factor", &alternative_route_max_share_factor.to_string())
            .append_pair("avoid", &avoid.to_string())
            .append_pair("key", &key.to_string())
            .finish();
        let uri_str = format!("{}/route{}", configuration.base_path, query);

        let uri = uri_str.parse();
        // TODO(farcaller): handle error
        // if let Err(e) = uri {
        //     return Box::new(futures::future::err(e));
        // }
        let mut req = hyper::Request::new(method, uri.unwrap());

        if let Some(ref user_agent) = configuration.user_agent {
            req.headers_mut().set(UserAgent::new(Cow::Owned(user_agent.clone())));
        }



        // send request
        Box::new(
        configuration.client.request(req)
            .map_err(|e| Error::from(e))
            .and_then(|resp| {
                let status = resp.status();
                resp.body().concat2()
                    .and_then(move |body| Ok((status, body)))
                    .map_err(|e| Error::from(e))
            })
            .and_then(|(status, body)| {
                if status.is_success() {
                    Ok(body)
                } else {
                    Err(Error::from((status, &*body)))
                }
            })
            .and_then(|body| {
                let parsed: Result<::models::RouteResponse, _> = serde_json::from_slice(&body);
                parsed.map_err(|e| Error::from(e))
            })
        )
    }

}
