// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/**
 * Polyaxon sdk
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.14.4
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.4.7
 *
 * Do not edit the class manually.
 *
 */

(function(factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/V1Auth', 'model/V1CodeRefBodyRequest', 'model/V1CodeReference', 'model/V1CredsBodyRequest', 'model/V1EntityStatusRequest', 'model/V1ListCodeRefResponse', 'model/V1ListProjectsResponse', 'model/V1ListRunsResponse', 'model/V1OwnedEntityUUIdRequest', 'model/V1OwnerBodyRequest', 'model/V1Project', 'model/V1ProjectBodyRequest', 'model/V1Run', 'model/V1RunBodyRequest', 'model/V1StatusResponse', 'model/V1Version', 'model/V1Versions', 'api/AuthServiceApi', 'api/ProjectServiceApi', 'api/RunServiceApi', 'api/VersionsServiceApi'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('./ApiClient'), require('./model/V1Auth'), require('./model/V1CodeRefBodyRequest'), require('./model/V1CodeReference'), require('./model/V1CredsBodyRequest'), require('./model/V1EntityStatusRequest'), require('./model/V1ListCodeRefResponse'), require('./model/V1ListProjectsResponse'), require('./model/V1ListRunsResponse'), require('./model/V1OwnedEntityUUIdRequest'), require('./model/V1OwnerBodyRequest'), require('./model/V1Project'), require('./model/V1ProjectBodyRequest'), require('./model/V1Run'), require('./model/V1RunBodyRequest'), require('./model/V1StatusResponse'), require('./model/V1Version'), require('./model/V1Versions'), require('./api/AuthServiceApi'), require('./api/ProjectServiceApi'), require('./api/RunServiceApi'), require('./api/VersionsServiceApi'));
  }
}(function(ApiClient, V1Auth, V1CodeRefBodyRequest, V1CodeReference, V1CredsBodyRequest, V1EntityStatusRequest, V1ListCodeRefResponse, V1ListProjectsResponse, V1ListRunsResponse, V1OwnedEntityUUIdRequest, V1OwnerBodyRequest, V1Project, V1ProjectBodyRequest, V1Run, V1RunBodyRequest, V1StatusResponse, V1Version, V1Versions, AuthServiceApi, ProjectServiceApi, RunServiceApi, VersionsServiceApi) {
  'use strict';

  /**
   * ERROR_UNKNOWN.<br>
   * The <code>index</code> module provides access to constructors for all the classes which comprise the public API.
   * <p>
   * An AMD (recommended!) or CommonJS application will generally do something equivalent to the following:
   * <pre>
   * var PolyaxonSdk = require('index'); // See note below*.
   * var xxxSvc = new PolyaxonSdk.XxxApi(); // Allocate the API class we're going to use.
   * var yyyModel = new PolyaxonSdk.Yyy(); // Construct a model instance.
   * yyyModel.someProperty = 'someValue';
   * ...
   * var zzz = xxxSvc.doSomething(yyyModel); // Invoke the service.
   * ...
   * </pre>
   * <em>*NOTE: For a top-level AMD script, use require(['index'], function(){...})
   * and put the application logic within the callback function.</em>
   * </p>
   * <p>
   * A non-AMD browser application (discouraged) might do something like this:
   * <pre>
   * var xxxSvc = new PolyaxonSdk.XxxApi(); // Allocate the API class we're going to use.
   * var yyy = new PolyaxonSdk.Yyy(); // Construct a model instance.
   * yyyModel.someProperty = 'someValue';
   * ...
   * var zzz = xxxSvc.doSomething(yyyModel); // Invoke the service.
   * ...
   * </pre>
   * </p>
   * @module index
   * @version 1.14.4
   */
  var exports = {
    /**
     * The ApiClient constructor.
     * @property {module:ApiClient}
     */
    ApiClient: ApiClient,
    /**
     * The V1Auth model constructor.
     * @property {module:model/V1Auth}
     */
    V1Auth: V1Auth,
    /**
     * The V1CodeRefBodyRequest model constructor.
     * @property {module:model/V1CodeRefBodyRequest}
     */
    V1CodeRefBodyRequest: V1CodeRefBodyRequest,
    /**
     * The V1CodeReference model constructor.
     * @property {module:model/V1CodeReference}
     */
    V1CodeReference: V1CodeReference,
    /**
     * The V1CredsBodyRequest model constructor.
     * @property {module:model/V1CredsBodyRequest}
     */
    V1CredsBodyRequest: V1CredsBodyRequest,
    /**
     * The V1EntityStatusRequest model constructor.
     * @property {module:model/V1EntityStatusRequest}
     */
    V1EntityStatusRequest: V1EntityStatusRequest,
    /**
     * The V1ListCodeRefResponse model constructor.
     * @property {module:model/V1ListCodeRefResponse}
     */
    V1ListCodeRefResponse: V1ListCodeRefResponse,
    /**
     * The V1ListProjectsResponse model constructor.
     * @property {module:model/V1ListProjectsResponse}
     */
    V1ListProjectsResponse: V1ListProjectsResponse,
    /**
     * The V1ListRunsResponse model constructor.
     * @property {module:model/V1ListRunsResponse}
     */
    V1ListRunsResponse: V1ListRunsResponse,
    /**
     * The V1OwnedEntityUUIdRequest model constructor.
     * @property {module:model/V1OwnedEntityUUIdRequest}
     */
    V1OwnedEntityUUIdRequest: V1OwnedEntityUUIdRequest,
    /**
     * The V1OwnerBodyRequest model constructor.
     * @property {module:model/V1OwnerBodyRequest}
     */
    V1OwnerBodyRequest: V1OwnerBodyRequest,
    /**
     * The V1Project model constructor.
     * @property {module:model/V1Project}
     */
    V1Project: V1Project,
    /**
     * The V1ProjectBodyRequest model constructor.
     * @property {module:model/V1ProjectBodyRequest}
     */
    V1ProjectBodyRequest: V1ProjectBodyRequest,
    /**
     * The V1Run model constructor.
     * @property {module:model/V1Run}
     */
    V1Run: V1Run,
    /**
     * The V1RunBodyRequest model constructor.
     * @property {module:model/V1RunBodyRequest}
     */
    V1RunBodyRequest: V1RunBodyRequest,
    /**
     * The V1StatusResponse model constructor.
     * @property {module:model/V1StatusResponse}
     */
    V1StatusResponse: V1StatusResponse,
    /**
     * The V1Version model constructor.
     * @property {module:model/V1Version}
     */
    V1Version: V1Version,
    /**
     * The V1Versions model constructor.
     * @property {module:model/V1Versions}
     */
    V1Versions: V1Versions,
    /**
     * The AuthServiceApi service constructor.
     * @property {module:api/AuthServiceApi}
     */
    AuthServiceApi: AuthServiceApi,
    /**
     * The ProjectServiceApi service constructor.
     * @property {module:api/ProjectServiceApi}
     */
    ProjectServiceApi: ProjectServiceApi,
    /**
     * The RunServiceApi service constructor.
     * @property {module:api/RunServiceApi}
     */
    RunServiceApi: RunServiceApi,
    /**
     * The VersionsServiceApi service constructor.
     * @property {module:api/VersionsServiceApi}
     */
    VersionsServiceApi: VersionsServiceApi
  };

  return exports;
}));
