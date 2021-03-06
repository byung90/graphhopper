/**
* GraphHopper Directions API
* You use the GraphHopper Directions API to add route planning, navigation and route optimization to your software. E.g. the Routing API has turn instructions and elevation data and the Route Optimization API solves your logistic problems and supports various constraints like time window and capacity restrictions. Also it is possible to get all distances between all locations with our fast Matrix API.
*
* OpenAPI spec version: 1.0.0
* 
*
* NOTE: This class is auto generated by the swagger code generator program.
* https://github.com/swagger-api/swagger-codegen.git
* Do not edit the class manually.
*/
package io.swagger.client.models

import io.swagger.client.models.Configuration
import io.swagger.client.models.CostMatrix
import io.swagger.client.models.Objective
import io.swagger.client.models.Relation
import io.swagger.client.models.Service
import io.swagger.client.models.Shipment
import io.swagger.client.models.Vehicle
import io.swagger.client.models.VehicleType

/**
 * 
 * @param vehicles An array of vehicles that can be employed
 * @param vehicle_types An array of vehicle types
 * @param services An array of services
 * @param shipments An array of shipments
 * @param relations An array of relations
 * @param objectives An array of objectives
 * @param cost_matrices An array of cost matrices
 * @param configuration 
 */
data class Request (
    /* An array of vehicles that can be employed */
    val vehicles: kotlin.Array<Vehicle>? = null,
    /* An array of vehicle types */
    val vehicle_types: kotlin.Array<VehicleType>? = null,
    /* An array of services */
    val services: kotlin.Array<Service>? = null,
    /* An array of shipments */
    val shipments: kotlin.Array<Shipment>? = null,
    /* An array of relations */
    val relations: kotlin.Array<Relation>? = null,
    /* An array of objectives */
    val objectives: kotlin.Array<Objective>? = null,
    /* An array of cost matrices */
    val cost_matrices: kotlin.Array<CostMatrix>? = null,
    val configuration: Configuration? = null
) {

}

