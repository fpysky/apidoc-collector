<?php

/**
 * @api {get} / 0 接口说明
 * @apiUse V1D0D0
 * @apiDescription 对接口中的一些公共特性进行说明
 * @apiName basedocs
 * @apiGroup A
 * @apiVersion 1.0.0
 *
 * @apiSuccess {string}       code    状态码，以业务码为准
 * @apiSuccess {String}       message     出错时提示的信息
 * @apiSuccess {object|array} data    数据
 *
 * @apiSuccessExample 请求成功
 * HTTP/2.0 200 OK
 *      {
 *          "code": "200000",
 *          "message": "",
 *          "data": {},
 *      }
 *
 * @apiError {string}       code 状态码
 * @apiError {string}       msg  提示信息
 * @apiError {object|array} data 数据
 *
 * @apiErrorExample 资源不存在
 * HTTP/2.0 404 Not Found
 *      {
 *          "code": "404000",
 *          "message": "资源不存在",
 *          "data": {},
 *      }
 **/
