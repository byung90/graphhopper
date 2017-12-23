/*
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


package com.graphhopper.directions.api.client.model;

import java.util.Objects;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import com.graphhopper.directions.api.client.model.GeocodingLocation;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

/**
 * GeocodingResponse
 */

public class GeocodingResponse {
  @SerializedName("hits")
  private List<GeocodingLocation> hits = null;

  @SerializedName("locale")
  private String locale = null;

  public GeocodingResponse hits(List<GeocodingLocation> hits) {
    this.hits = hits;
    return this;
  }

  public GeocodingResponse addHitsItem(GeocodingLocation hitsItem) {
    if (this.hits == null) {
      this.hits = new ArrayList<GeocodingLocation>();
    }
    this.hits.add(hitsItem);
    return this;
  }

   /**
   * Get hits
   * @return hits
  **/
  @ApiModelProperty(value = "")
  public List<GeocodingLocation> getHits() {
    return hits;
  }

  public void setHits(List<GeocodingLocation> hits) {
    this.hits = hits;
  }

  public GeocodingResponse locale(String locale) {
    this.locale = locale;
    return this;
  }

   /**
   * Get locale
   * @return locale
  **/
  @ApiModelProperty(value = "")
  public String getLocale() {
    return locale;
  }

  public void setLocale(String locale) {
    this.locale = locale;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    GeocodingResponse geocodingResponse = (GeocodingResponse) o;
    return Objects.equals(this.hits, geocodingResponse.hits) &&
        Objects.equals(this.locale, geocodingResponse.locale);
  }

  @Override
  public int hashCode() {
    return Objects.hash(hits, locale);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class GeocodingResponse {\n");
    
    sb.append("    hits: ").append(toIndentedString(hits)).append("\n");
    sb.append("    locale: ").append(toIndentedString(locale)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(java.lang.Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }
  
}
