/**
 * 匿名掲示板API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * OpenAPI spec version: 1.0.0
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 *
 * OpenAPI Generator version: 3.3.4
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'));
  } else {
    // Browser globals (root is window)
    if (!root.Api) {
      root.Api = {};
    }
    root.Api.ErrorProperties = factory(root.Api.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';



  /**
   * The ErrorProperties model module.
   * @module model/ErrorProperties
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>ErrorProperties</code>.
   * @alias module:model/ErrorProperties
   * @class
   * @param title {String} 
   * @param detail {String} 
   */
  var exports = function(title, detail) {
    var _this = this;

    _this['title'] = title;
    _this['detail'] = detail;
  };

  /**
   * Constructs a <code>ErrorProperties</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/ErrorProperties} obj Optional instance to populate.
   * @return {module:model/ErrorProperties} The populated <code>ErrorProperties</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();
      if (data.hasOwnProperty('title')) {
        obj['title'] = ApiClient.convertToType(data['title'], 'String');
      }
      if (data.hasOwnProperty('detail')) {
        obj['detail'] = ApiClient.convertToType(data['detail'], 'String');
      }
    }
    return obj;
  }

  /**
   * @member {String} title
   */
  exports.prototype['title'] = undefined;
  /**
   * @member {String} detail
   */
  exports.prototype['detail'] = undefined;



  return exports;
}));


