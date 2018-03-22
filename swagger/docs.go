package swagger

//This file is generated automatically. Do not try to edit it manually.

var resourceListingJson = `{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "/",
    "apis": [
        {
            "path": "/swagger/handler1",
            "description": "Some endpoint description"
        },
        {
            "path": "/swagger/handler2",
            "description": "Some endpoint description"
        },
        {
            "path": "/service/swagger",
            "description": "Some API"
        }
    ],
    "info": {
        "title": "API for some service",
        "description": "My API usually works as expected.",
        "contact": "api@gmail.com"
    }
}`
var apiDescriptionsJson = map[string]string{"swagger/handler1": `{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "/",
    "resourcePath": "/swagger/handler1",
    "apis": [
        {
            "path": "/service/swagger/handler1",
            "description": "Some endpoint description",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "Return some text",
                    "type": "string",
                    "items": {},
                    "summary": "Some endpoint description",
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "string"
                        }
                    ]
                }
            ]
        }
    ]
}`, "swagger/handler2": `{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "/",
    "resourcePath": "/swagger/handler2",
    "apis": [
        {
            "path": "/service/swagger/handler2",
            "description": "Some endpoint description",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "Return some text",
                    "type": "string",
                    "items": {},
                    "summary": "Some endpoint description",
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "",
                            "responseType": "object",
                            "responseModel": "string"
                        }
                    ]
                }
            ]
        }
    ]
}`}
