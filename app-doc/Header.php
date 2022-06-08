<?php

/**
 * @apiDefine headers
 * @apiHeaderExample {json} Header-App
 * {
 *      "Accept": "application/json"
 * }
 * @apiHeaderExample {json} Header-WeChat
 * {
 *      "Accept": "application/json"
 * }
 * @apiHeaderExample {json} Header-Wap
 * {
 *      "Accept": "application/json"
 * }
 */

/**
 * @apiDefine htmlHeaders
 * @apiHeaderExample {json} HTML-Header-App
 * {
 *      "Accept": "text/html"
 * }
 * @apiHeaderExample {json} HTML-Header-WeChat
 * {
 *      "Accept": "text/html"
 * }
 * @apiHeaderExample {json} HTML-Header-Wap
 * {
 *      "Accept": "text/html"
 * }
 */

/**
 * @apiDefine V1D0D0
 * @apiHeaderExample {json} Header-App
 * {
 *      "Accept": "application/json"
 * }
 * @apiHeaderExample {json} Header-WeChat
 * {
 *      "Accept": "application/json"
 * }
 * @apiHeaderExample {json} Header-Wap
 * {
 *      "Accept": "application/json"
 * }
 */

/**
 * @apiDefine CommonError Error with status code 500
 *
 * @apiError (CommonError) {string} code 请求状态码，非200000
 * @apiError (CommonError) {string} message 请求状态码描述
 * @apiError (CommonError) {object} data 请求内容
 *
 * @apiErrorExample {json} Error-Response:
 *     {
 *       "code": "500000",
 *       "message": "出错了，请再试试吧~",
 *       "data":{}
 *     }
 */

/**
 * @apiDefine urlParameter 嵌入式URL传参
 */

/**
 * @apiDefine query URL传参
 */

/**
 * @apiDefine bodyRawJson JSON传参
 */
