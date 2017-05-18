/**
 * GraphHopper Directions API
 * With the GraphHopper Directions API you get reliable and fast web services for routing and more with world wide coverage. We offer A-to-B routing via the Routing API optionally with turn instructions and elevation data as well as route optimization with various constraints like time window and capacity restrictions. Also it is possible to get all distances between all locations with our fast Matrix API. 
 *
 * OpenAPI spec version: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */


#include "SWGSolution.h"

#include "SWGHelpers.h"

#include <QJsonDocument>
#include <QJsonArray>
#include <QObject>
#include <QDebug>

namespace Swagger {


SWGSolution::SWGSolution(QString* json) {
    init();
    this->fromJson(*json);
}

SWGSolution::SWGSolution() {
    init();
}

SWGSolution::~SWGSolution() {
    this->cleanup();
}

void
SWGSolution::init() {
    costs = 0;
    distance = 0;
    time = 0L;
    transport_time = 0L;
    max_operation_time = 0L;
    waiting_time = 0L;
    no_vehicles = 0;
    no_unassigned = 0;
    routes = new QList<SWGRoute*>();
    unassigned = new SWGSolution_unassigned();
}

void
SWGSolution::cleanup() {
    








    if(routes != nullptr) {
        QList<SWGRoute*>* arr = routes;
        foreach(SWGRoute* o, *arr) {
            delete o;
        }
        delete routes;
    }

    if(unassigned != nullptr) {
        delete unassigned;
    }
}

SWGSolution*
SWGSolution::fromJson(QString &json) {
    QByteArray array (json.toStdString().c_str());
    QJsonDocument doc = QJsonDocument::fromJson(array);
    QJsonObject jsonObject = doc.object();
    this->fromJsonObject(jsonObject);
    return this;
}

void
SWGSolution::fromJsonObject(QJsonObject &pJson) {
    ::Swagger::setValue(&costs, pJson["costs"], "qint32", "");
    ::Swagger::setValue(&distance, pJson["distance"], "qint32", "");
    ::Swagger::setValue(&time, pJson["time"], "qint64", "");
    ::Swagger::setValue(&transport_time, pJson["transport_time"], "qint64", "");
    ::Swagger::setValue(&max_operation_time, pJson["max_operation_time"], "qint64", "");
    ::Swagger::setValue(&waiting_time, pJson["waiting_time"], "qint64", "");
    ::Swagger::setValue(&no_vehicles, pJson["no_vehicles"], "qint32", "");
    ::Swagger::setValue(&no_unassigned, pJson["no_unassigned"], "qint32", "");
    
    ::Swagger::setValue(&routes, pJson["routes"], "QList", "SWGRoute");
    
    ::Swagger::setValue(&unassigned, pJson["unassigned"], "SWGSolution_unassigned", "SWGSolution_unassigned");
}

QString
SWGSolution::asJson ()
{
    QJsonObject* obj = this->asJsonObject();
    
    QJsonDocument doc(*obj);
    QByteArray bytes = doc.toJson();
    return QString(bytes);
}

QJsonObject*
SWGSolution::asJsonObject() {
    QJsonObject* obj = new QJsonObject();
    
    obj->insert("costs", QJsonValue(costs));

    obj->insert("distance", QJsonValue(distance));

    obj->insert("time", QJsonValue(time));

    obj->insert("transport_time", QJsonValue(transport_time));

    obj->insert("max_operation_time", QJsonValue(max_operation_time));

    obj->insert("waiting_time", QJsonValue(waiting_time));

    obj->insert("no_vehicles", QJsonValue(no_vehicles));

    obj->insert("no_unassigned", QJsonValue(no_unassigned));

    QJsonArray routesJsonArray;
    toJsonArray((QList<void*>*)routes, &routesJsonArray, "routes", "SWGRoute");
    obj->insert("routes", routesJsonArray);

    toJsonValue(QString("unassigned"), unassigned, obj, QString("SWGSolution_unassigned"));

    return obj;
}

qint32
SWGSolution::getCosts() {
    return costs;
}
void
SWGSolution::setCosts(qint32 costs) {
    this->costs = costs;
}

qint32
SWGSolution::getDistance() {
    return distance;
}
void
SWGSolution::setDistance(qint32 distance) {
    this->distance = distance;
}

qint64
SWGSolution::getTime() {
    return time;
}
void
SWGSolution::setTime(qint64 time) {
    this->time = time;
}

qint64
SWGSolution::getTransportTime() {
    return transport_time;
}
void
SWGSolution::setTransportTime(qint64 transport_time) {
    this->transport_time = transport_time;
}

qint64
SWGSolution::getMaxOperationTime() {
    return max_operation_time;
}
void
SWGSolution::setMaxOperationTime(qint64 max_operation_time) {
    this->max_operation_time = max_operation_time;
}

qint64
SWGSolution::getWaitingTime() {
    return waiting_time;
}
void
SWGSolution::setWaitingTime(qint64 waiting_time) {
    this->waiting_time = waiting_time;
}

qint32
SWGSolution::getNoVehicles() {
    return no_vehicles;
}
void
SWGSolution::setNoVehicles(qint32 no_vehicles) {
    this->no_vehicles = no_vehicles;
}

qint32
SWGSolution::getNoUnassigned() {
    return no_unassigned;
}
void
SWGSolution::setNoUnassigned(qint32 no_unassigned) {
    this->no_unassigned = no_unassigned;
}

QList<SWGRoute*>*
SWGSolution::getRoutes() {
    return routes;
}
void
SWGSolution::setRoutes(QList<SWGRoute*>* routes) {
    this->routes = routes;
}

SWGSolution_unassigned*
SWGSolution::getUnassigned() {
    return unassigned;
}
void
SWGSolution::setUnassigned(SWGSolution_unassigned* unassigned) {
    this->unassigned = unassigned;
}



} /* namespace Swagger */
