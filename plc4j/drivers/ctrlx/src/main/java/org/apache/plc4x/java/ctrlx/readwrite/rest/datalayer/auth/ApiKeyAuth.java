/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * ctrlX CORE - Data Layer API
 * This is the base API for the ctrlX Data Layer access on ctrlX CORE devices <ul> <li>Click 'Authorize' to open the 'Available authorizations' dialog.</li> <li>Enter 'username' and 'password'. The 'Client credentials location' selector together with the 'client_id' and 'client_secret' fields as well as the 'Bearer' section can be ignored.</li> <li>Click 'Authorize' and then 'Close' to close the 'Available authorizations' dialog.</li> <li>Try out those GET, PUT, ... operations you're interested in.</li> </ul>
 *
 * The version of the OpenAPI document: 2.1.0
 * Contact: support@boschrexroth.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.apache.plc4x.java.ctrlx.readwrite.rest.datalayer.auth;

import org.apache.plc4x.java.ctrlx.readwrite.rest.datalayer.Pair;

import java.util.List;
import java.util.Map;

@jakarta.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-11-18T13:34:36.056861+01:00[Europe/Berlin]")
public class ApiKeyAuth implements Authentication {
  private final String location;
  private final String paramName;

  private String apiKey;
  private String apiKeyPrefix;

  public ApiKeyAuth(String location, String paramName) {
    this.location = location;
    this.paramName = paramName;
  }

  public String getLocation() {
    return location;
  }

  public String getParamName() {
    return paramName;
  }

  public String getApiKey() {
    return apiKey;
  }

  public void setApiKey(String apiKey) {
    this.apiKey = apiKey;
  }

  public String getApiKeyPrefix() {
    return apiKeyPrefix;
  }

  public void setApiKeyPrefix(String apiKeyPrefix) {
    this.apiKeyPrefix = apiKeyPrefix;
  }

  @Override
  public void applyToParams(List<Pair> queryParams, Map<String, String> headerParams, Map<String, String> cookieParams) {
    if (apiKey == null) {
      return;
    }
    String value;
    if (apiKeyPrefix != null) {
      value = apiKeyPrefix + " " + apiKey;
    } else {
      value = apiKey;
    }
    if ("query".equals(location)) {
      queryParams.add(new Pair(paramName, value));
    } else if ("header".equals(location)) {
      headerParams.put(paramName, value);
    } else if ("cookie".equals(location)) {
      cookieParams.put(paramName, value);
    }
  }
}
