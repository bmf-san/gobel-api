# API
- [API](#api)
  - [Root Endpoint](#root-endpoint)
  - [Response Format](#response-format)
    - [Success](#success)
      - [GET](#get)
      - [POST](#post)
      - [PATCH](#patch)
      - [DELETE](#delete)
    - [Error](#error)
      - [Common](#common)
      - [GET](#get-1)
      - [POST](#post-1)
      - [PATCH](#patch-1)
      - [DELETE](#delete-1)
      - [Unauthorized](#unauthorized)
      - [Forbidden](#forbidden)
- [Public API](#public-api)
  - [Authentication](#authentication)
    - [Get an access token and refresh by credentials](#get-an-access-token-and-refresh-by-credentials)
      - [Endpoint](#endpoint)
      - [Request bodies](#request-bodies)
      - [Response bodies](#response-bodies)
  - [Post](#post-2)
    - [Get all publish posts](#get-all-publish-posts)
      - [Endpoint](#endpoint-1)
      - [Query Parameters](#query-parameters)
      - [Requests](#requests)
        - [Header](#header)
        - [Body](#body)
      - [Response](#response)
        - [Header](#header-1)
        - [Body](#body-1)
    - [Get all publish posts by category](#get-all-publish-posts-by-category)
      - [Endpoint](#endpoint-2)
      - [Path Parameters](#path-parameters)
      - [Request](#request)
        - [Header](#header-2)
        - [Body](#body-2)
      - [Response](#response-1)
        - [Header](#header-3)
        - [Body](#body-3)
    - [Get all publish posts by tag](#get-all-publish-posts-by-tag)
      - [Endpoint](#endpoint-3)
      - [Path Parameters](#path-parameters-1)
      - [Request](#request-1)
        - [Header](#header-4)
        - [Body](#body-4)
      - [Response](#response-2)
        - [Header](#header-5)
        - [Body](#body-5)
    - [Get the specified post by title](#get-the-specified-post-by-title)
      - [Endpoint](#endpoint-4)
      - [Path Parameters](#path-parameters-2)
      - [Request](#request-2)
        - [Header](#header-6)
        - [Body](#body-6)
      - [Response](#response-3)
        - [Header](#header-7)
        - [Body](#body-7)
    - [Store a newly comment](#store-a-newly-comment)
      - [Endpoint](#endpoint-5)
      - [Path Parameters](#path-parameters-3)
      - [Request](#request-3)
        - [Header](#header-8)
        - [Body](#body-8)
      - [Response](#response-4)
        - [Header](#header-9)
        - [Body](#body-9)
    - [Get all categories](#get-all-categories)
      - [Endpoint](#endpoint-6)
      - [Query Parameters](#query-parameters-1)
      - [Request](#request-4)
        - [Header](#header-10)
        - [Body](#body-10)
      - [Response](#response-5)
        - [Header](#header-11)
        - [Body](#body-11)
    - [Get the specified category by name](#get-the-specified-category-by-name)
      - [Endpoint](#endpoint-7)
      - [Path Parameters](#path-parameters-4)
      - [Request](#request-5)
        - [Header](#header-12)
        - [Body](#body-12)
      - [Responsej](#responsej)
        - [Header](#header-13)
        - [Body](#body-13)
    - [Get all tags](#get-all-tags)
      - [Endpoint](#endpoint-8)
      - [Query Parameters](#query-parameters-2)
      - [Request bodies](#request-bodies-1)
        - [Header](#header-14)
        - [Body](#body-14)
      - [Response](#response-6)
        - [Header](#header-15)
        - [Body](#body-15)
    - [Get the specified tag by name](#get-the-specified-tag-by-name)
      - [Endpoint](#endpoint-9)
      - [Path Parameters](#path-parameters-5)
      - [Request](#request-6)
        - [Header](#header-16)
        - [Body](#body-16)
      - [Response](#response-7)
        - [Header](#header-17)
        - [Body](#body-17)
- [Private API](#private-api)
  - [Authentication](#authentication-1)
    - [Disable an access token](#disable-an-access-token)
      - [Endpoint](#endpoint-10)
      - [Request](#request-7)
        - [Header](#header-18)
        - [Body](#body-18)
      - [Response](#response-8)
        - [Header](#header-19)
        - [Body](#body-19)
    - [Refresh an acess token and refresh token](#refresh-an-acess-token-and-refresh-token)
      - [Endpoint](#endpoint-11)
      - [Request](#request-8)
        - [Header](#header-20)
        - [Body](#body-20)
      - [Response](#response-9)
        - [Header](#header-21)
        - [Body](#body-21)
    - [Get the specified admin by access token](#get-the-specified-admin-by-access-token)
      - [Endpoint](#endpoint-12)
      - [Request](#request-9)
        - [Header](#header-22)
        - [Body](#body-22)
      - [Response](#response-10)
        - [Header](#header-23)
        - [Body](#body-23)
    - [Get all posts](#get-all-posts)
      - [Endpoint](#endpoint-13)
      - [Query Parameters](#query-parameters-3)
      - [Request](#request-10)
        - [Header](#header-24)
        - [Body](#body-24)
      - [Response](#response-11)
        - [Header](#header-25)
        - [Body](#body-25)
    - [Get the specified post by id](#get-the-specified-post-by-id)
      - [Endpoint](#endpoint-14)
      - [Path Parameters](#path-parameters-6)
      - [Request](#request-11)
        - [Header](#header-26)
        - [Body](#body-26)
      - [Response](#response-12)
        - [Header](#header-27)
        - [Body](#body-27)
    - [Store a newly post](#store-a-newly-post)
      - [Endpoint](#endpoint-15)
      - [Request](#request-12)
        - [Header](#header-28)
        - [Body](#body-28)
      - [Response](#response-13)
        - [Header](#header-29)
        - [Body](#body-29)
    - [Update the specified post](#update-the-specified-post)
      - [Endpoint](#endpoint-16)
      - [Query Parameters](#query-parameters-4)
      - [Request](#request-13)
        - [Header](#header-30)
        - [Body](#body-30)
        - [Header](#header-31)
        - [Body](#body-31)
    - [Remove the specified post](#remove-the-specified-post)
      - [Endpoint](#endpoint-17)
      - [Query Parameters](#query-parameters-5)
      - [Request](#request-14)
        - [Header](#header-32)
        - [Body](#body-32)
      - [Response](#response-14)
        - [Header](#header-33)
        - [Body](#body-33)
    - [Get all comments](#get-all-comments)
      - [Endpoint](#endpoint-18)
      - [Query Parameters](#query-parameters-6)
      - [Request](#request-15)
        - [Header](#header-34)
        - [Body](#body-34)
      - [Response](#response-15)
        - [Header](#header-35)
        - [Body](#body-35)
    - [Get the specified comment by id](#get-the-specified-comment-by-id)
      - [Endpoint](#endpoint-19)
      - [Path Parameters](#path-parameters-7)
      - [Request](#request-16)
        - [Header](#header-36)
        - [Body](#body-36)
      - [Response](#response-16)
        - [Header](#header-37)
        - [Body](#body-37)
  - [Update the specified comment status](#update-the-specified-comment-status)
      - [Endpoint](#endpoint-20)
      - [Query Parameters](#query-parameters-7)
      - [Request](#request-17)
        - [Header](#header-38)
        - [Body](#body-38)
      - [Response](#response-17)
        - [Header](#header-39)
        - [Body](#body-39)
    - [Get all categories](#get-all-categories-1)
      - [Endpoint](#endpoint-21)
      - [Query Parameters](#query-parameters-8)
      - [Request](#request-18)
        - [Header](#header-40)
        - [Body](#body-40)
      - [Response](#response-18)
        - [Header](#header-41)
        - [Body](#body-41)
    - [Get the specified category by id](#get-the-specified-category-by-id)
      - [Endpoint](#endpoint-22)
      - [Path Parameters](#path-parameters-8)
      - [Request](#request-19)
        - [Header](#header-42)
        - [Body](#body-42)
      - [Response](#response-19)
        - [Header](#header-43)
        - [Body](#body-43)
    - [Get all tags](#get-all-tags-1)
      - [Endpoint](#endpoint-23)
      - [Query Parameters](#query-parameters-9)
      - [Response](#response-20)
        - [Header](#header-44)
        - [Body](#body-44)
      - [Request](#request-20)
        - [Header](#header-45)
        - [Body](#body-45)
    - [Get the specified tag by id](#get-the-specified-tag-by-id)
      - [Endpoint](#endpoint-24)
      - [Path Parameters](#path-parameters-9)
      - [Request](#request-21)
        - [Header](#header-46)
        - [Body](#body-46)
      - [Request](#request-22)
        - [Header](#header-47)
        - [Body](#body-47)
    - [Store a newly category](#store-a-newly-category)
      - [Endpoint](#endpoint-25)
      - [Request](#request-23)
        - [Header](#header-48)
        - [Body](#body-48)
      - [Response](#response-21)
        - [Header](#header-49)
        - [Body](#body-49)
    - [Update the specified category](#update-the-specified-category)
      - [Endpoint](#endpoint-26)
      - [Query Parameters](#query-parameters-10)
      - [Request](#request-24)
        - [Header](#header-50)
        - [Body](#body-50)
      - [Response](#response-22)
        - [Header](#header-51)
        - [Body](#body-51)
    - [Remove the specified category](#remove-the-specified-category)
      - [Endpoint](#endpoint-27)
      - [Query Parameters](#query-parameters-11)
      - [Request](#request-25)
        - [Header](#header-52)
        - [Body](#body-52)
      - [Response](#response-23)
        - [Header](#header-53)
        - [Body](#body-53)
    - [Store a newly tag](#store-a-newly-tag)
      - [Endpoint](#endpoint-28)
      - [Request](#request-26)
        - [Header](#header-54)
        - [Body](#body-54)
      - [Response](#response-24)
        - [Header](#header-55)
        - [Body](#body-55)
    - [Update the specified tag](#update-the-specified-tag)
      - [Endpoint](#endpoint-29)
      - [Query Parameters](#query-parameters-12)
      - [Request](#request-27)
        - [Header](#header-56)
        - [Body](#body-56)
      - [Response](#response-25)
        - [Header](#header-57)
        - [Body](#body-57)
    - [Remove the specified tag](#remove-the-specified-tag)
      - [Endpoint](#endpoint-30)
      - [Query Parameters](#query-parameters-13)
      - [Request](#request-28)
        - [Header](#header-58)
        - [Body](#body-58)
      - [Response](#response-26)
        - [Header](#header-59)
        - [Body](#body-59)

## Root Endpoint
`API_DOMAIN/api`

## Response Format
### Success
#### GET
```json
HTTP/1.1 200
Content-Type: application/json

{
    "id": 1,
    "name": "foo",
    "created_at": "2020-02-12 01:00:00",
    "updated_at": "2020-02-12 01:00:00",
}
```

```json
HTTP/1.1 200
Pagination-Count: 100
Pagination-PageCount: 10
Pagination-Page: 5
Pagination-Limit: 10
Content-Type: application/json

[
    {
        "id": 1,
        "name": "foo",
        "created_at": "2020-02-12 01:00:00",
        "updated_at": "2020-02-12 01:00:00",
    },
    {
        "id": 1,
        "name": "bar",
        "created_at": "2020-02-12 01:00:00",
        "updated_at": "2020-02-12 01:00:00",
    }
]
```

#### POST
```json
HTTP/1.1  201
Content-Type: application/json

{
    "message": "OK"
}
```

```json
HTTP/1.1  201
Content-Type: application/json

{
    "id": 1,
    "name": "foo",
    "created_at": "2020-02-12 01:00:00",
    "updated_at": "2020-02-12 01:00:00","
}
```

#### PATCH
If updated entity is to be sent after the update.

```json
HTTP/1.1  200
Content-Type: application/json

{
    "id": 1,
    "name": "foo",
    "created_at": "2020-02-12 01:00:00",
    "updated_at": "2020-02-12 01:00:00","
}
```

If updated entity is not to be sent after the update
```json
HTTP/1.1  204
```

#### DELETE
```json
HTTP/1.1  204

{
    "message": "OK"
}
```

### Error
#### Common
```json
HTTP/1.1 500
Content-Type: application/json

{
    "message": "An unexpected condition has occurred"
}
```

#### GET
```json
HTTP/1.1  404
Content-Type: application/json

{
    "message": "The item does not exist"
}
```

#### POST
```json
HTTP/1.1  400
Content-Type: application/json
{

	"message": "Validation Failed",
	"errors": {
        "name": [
            "name is required",
            "name must be string"
        ],
        "email": [
            "name is required",
            "name must be string"
        ]
	}
}
```

#### PATCH
```json
HTTP/1.1  400
Content-Type: application/json

{

	"message": "Validation Failed.",
	"errors": {
        "name": [
            "name is required",
            "name must be string"
        ],
        "email": [
            "name is required",
            "name must be string"
        ]
	}
}
```

```json
HTTP/1.1  404
Content-Type: application/json

{
    "message": "The item does not exist"
}
```

#### DELETE
```json
HTTP/1.1  404
Content-Type: application/json

{
    "message": "The item does not exist"
}
```

#### Unauthorized
```json
HTTP/1.1  401
Content-Type: application/json

{
    "message": "Authentication credentials were missing or incorrect"
}
```

#### Forbidden
```json
HTTP/1.1  403
Content-Type: application/json

{
    "message": "The request is understood, but it has been refused or access is not allowed"
}
```

# Public API
Public API is open api that does not required authentication.

| Method |        Endpoint         |                       Description                       |
| :----- | :---------------------- | :------------------------------------------------------ |
| POST   | /signin                 | Get an access token and a refresh token by credentials. |
| GET    | /posts                  | Get all publish posts.                                  |
| GET    | /posts/categories/:name | Get all publish posts by cagtegory.                     |
| GET    | /posts/tags/:name       | Get all publish posts by tag.                           |
| GET    | /posts/:title           | Get the specified post by title.                        |
| POST   | /posts/:title/comments  | Store a newly comment.                                  |
| GET    | /categories             | Get all categories.                                     |
| GET    | /categories/:name       | Get the specified category by name.                     |
| GET    | /tags                   | Get all tags.                                           |
| GET    | /tags/:name             | Get the specified tag by name.                          |

## Authentication
### Get an access token and refresh by credentials
#### Endpoint
`POST /signin`

#### Request bodies
```json
{
    "email": "example@example.com",
    "password": "password"
}
```

#### Response bodies
```json
{
"access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6ImJmNjQ0Yjg1LTEyNDgtNDYwYi1iNGE5LTExYTNkNGNmYjY0NyIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTYwMTIwMTQ1MCwiaWQiOjJ9.DzxvWJPaxV8sv_ZGLsiEkKmGLbcYH2hnA9n6Q2Y86AI",
"refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDE4MDUzNTAsImlkIjoyLCJyZWZyZXNoX3V1aWQiOiI3MzQyOGNjOS0zODc2LTRlZDYtOGYwOS1kYjhkZDBlYTk3M2QifQ.SrkDUuJHYVIQ9wn09E7f7xAUhez1eNEkeOqoU4bapfE"
}
```

## Post
### Get all publish posts
#### Endpoint
`GET /posts`

#### Query Parameters
| Name  |  Type   | Required |       Description        |
| :---- | :------ | :------- | :----------------------- |
| page  | integer | optional | A page number of posts.  |
| limit | integer | optional | A limit number of posts. |

#### Requests
##### Header
|         Name         |        Description         | Example |
| :------------------- | :------------------------- | :------ |
| Pagination-Count     | A count of records         | 100     |
| Pagination-PageCount | A count of page            | 10      |
| Pagination-Page      | A current page number      | 5       |
| Pagination-Limit     | A limit of number per page | 10      |

##### Body
N/A

#### Response
##### Header
|         Name         |        Description         | Example |
| :------------------- | :------------------------- | :------ |
| Pagination-Count     | A count of records         | 100     |
| Pagination-PageCount | A count of page            | 10      |
| Pagination-Page      | A current page number      | 5       |
| Pagination-Limit     | A limit of number per page | 10      |

##### Body
```json
[
    {
        "id":1,
        "admin":{
            "id":1,
            "name":"1admin"
        },
        "category":{
            "id":1,
            "name":"1category"
        },
        "tags":[
            {
                "id":1,
                "name":"1tag"
            },
            {
                "id":2,
                "name":"2tag"
            }
        ],
        "title":"1title",
        "md_body":"#md-body",
        "html_body":"<html><body>html_body</body></html>",
        "status":"publish",
        "comments":null,
        "created_at":"2020-09-28T15:10:18Z",
        "updated_at":"2020-09-28T15:10:18Z"
    },
    {
        "id":2,
        "admin":{
            "id":1,
            "name":"1admin"
        },
        "category":{
            "id":1,
            "name":"1category"
        },
        "tags":[
            {
                "id":1,
                "name":"1tag"
            },
            {
                "id":2,
                "name":"2tag"
            }
        ],
        "title": "2title",
        "md_body":"#md-body",
        "html_body":"<html><body>html_body</body></html>",
        "status":"publish",
        "comments":null,
        "created_at":"2020-09-28T15:10:21Z",
        "updated_at":"2020-09-28T15:10:21Z"
    },
    {
        "id":3,
        "admin":{
            "id":1,
            "name":"1admin"
        },
        "category":{
            "id":1,
            "name":"1category"
        },
        "tags":[
            {
                "id":1,
                "name":"1tag"
            },
            {
                "id":2,
                "name":"2tag"
            }
        ],
        "title":"3title",
        "md_body":"#md-body",
        "html_body":"<html><body>html_body</body></html>",
        "status":"publish",
        "comments":null,
        "created_at":"2020-09-28T15:10:24Z",
        "updated_at":"2020-09-28T15:10:24Z"
    }
]
```

### Get all publish posts by category
#### Endpoint
`GET /posts/categories/:name`

#### Path Parameters
| Name |  Type  | Required |      Description      |
| :--- | :----- | :------- | :-------------------- |
| name | string | required | A name of a category. |

#### Request
##### Header
N/A

##### Body
N/A

#### Response
##### Header
N/A

##### Body
```json
[
    {
        "id":1,
        "name":"1category"
    },
    {
        "id":2,
        "name":"2category"
    },
    {
        "id":3,
        "name":"3category"
    }
]
```

### Get all publish posts by tag
#### Endpoint
`GET /posts/tags/:name`

#### Path Parameters
| Name |  Type  | Required |   Description    |
| :--- | :----- | :------- | :--------------- |
| name | string | required | A name of a tag. |

#### Request
##### Header
N/A

##### Body
N/A

#### Response
##### Header
N/A

##### Body
```json
[
    {
        "id":1,
        "admin":{
            "id":1,
            "name":"1admin"
        },
        "category":{
            "id":1,
            "name":"1category"
        },
        "tags":[
            {
                "id":1,
                "name":"1tag"
            },
            {
                "id":2,
                "name":"2tag"
            }
        ],
        "title":"1title",
        "md_body":"#md-body",
        "html_body":"\u003chtml\u003e\u003cbody\u003ehtml_body\u003c/body\u003e\u003c/html\u003e",
        "status":"publish",
        "comments":null,
        "created_at":"2020-09-28T15:10:18Z",
        "updated_at":"2020-09-28T15:10:18Z"
    },
    {
        "id":2,
        "admin":{
            "id":2,
            "name":"2admin"
        },
        "category":{
            "id":1,
            "name":"1category"
        },
        "tags":[
            {
                "id":1,
                "name":"1tag"
            },
            {
                "id":2,
                "name":"2tag"
            }
        ],
        "title":"2title",
        "md_body":"#md-body",
        "html_body":"\u003chtml\u003e\u003cbody\u003ehtml_body\u003c/body\u003e\u003c/html\u003e",
        "status":"publish",
        "comments":null,
        "created_at":"2020-09-28T15:10:24Z",
        "updated_at":"2020-09-28T15:10:24Z"
    },
    {
        "id":3,
        "admin":{
            "id":3,
            "name":"3admin"
        },
        "category":{
            "id":2,
            "name":"2category"
        },
        "tags":[
            {
                "id":1,
                "name":"1tag"
            },
            {
                "id":2,
                "name":"2tag"
            }
        ],
        "title":"3title",
        "md_body":"#md-body",
        "html_body":"\u003ch1 id=\"ge\"\u003ege\u003c/h1\u003e\n",
        "status":"publish",
        "comments":null,
        "created_at":"2020-09-28T15:10:27Z",
        "updated_at":"2020-09-28T15:10:27Z"
    },
    {
        "id":4,
        "admin":{
            "id":4,
            "name":"4admin"
        },
        "category":{
            "id":1,
            "name":"1category"
        },
        "tags":[
            {
                "id":1,
                "name":"1tag"
            },
            {
                "id":2,
                "name":"2tag"
            }
        ],
        "title":"4title",
        "md_body":"#md-body",
        "html_body":"\u003chtml\u003e\u003cbody\u003ehtml_body\u003c/body\u003e\u003c/html\u003e",
        "status":"publish",
        "comments":null,
        "created_at":"2020-09-28T15:10:30Z",
        "updated_at":"2020-09-28T15:10:30Z"
    }
]
```

### Get the specified post by title
#### Endpoint
`GET /posts/:title`

#### Path Parameters
| Name  |  Type  | Required |    Description     |
| :---- | :----- | :------- | :----------------- |
| title | string | required | A title of a post. |

#### Request
##### Header
N/A

##### Body
N/A

#### Response
##### Header
N/A

##### Body
```json
{
    "id":11,
    "admin":{
        "id":2,
        "name":"2admin"
    },
    "category":{
        "id":1,
        "name":"1category"
    },
    "tags":[
        {
            "id":1,
            "name":"1tag"
        },
        {
            "id":2,
            "name":"2tag"
        }
    ],
    "title":"1title",
    "md_body":"#md-body",
    "html_body":"\u003chtml\u003e\u003cbody\u003ehtml_body\u003c/body\u003e\u003c/html\u003e",
    "status":"publish",
    "comments":null,
    "created_at":"2020-09-28T15:18:16Z",
    "updated_at":"2020-09-28T15:18:16Z"
}
```

### Store a newly comment
#### Endpoint
`POST /posts/:title/comments`

#### Path Parameters
| Name  |  Type  | Required |    Description     |
| :---- | :----- | :------- | :----------------- |
| title | string | required | A title of a post. |

#### Request
##### Header
N/A

##### Body
```json
{
	"body": "foobar"
}
```

#### Response
##### Header
N/A

##### Body
```json
{
    "id":1,
    "post_id":1,
    "body":"foobar",
    "status":"pending",
    "created_at":"2020-10-05T15:31:44Z"
}
```

### Get all categories
#### Endpoint
`GET /categories`

#### Query Parameters
| Name  |  Type   | Required |         Description          |
| :---- | :------ | :------- | :--------------------------- |
| page  | integer | optional | A page number of categories. |
| limit | integer | optional | A limit number of posts.     |

#### Request
##### Header
N/A

##### Body
N/A

#### Response
##### Header
N/A

##### Body
```json
[
    {
        "id":1,
        "name":"1category"
    },
    {
        "id":2,
        "name":"2category"
    },
    {
        "id":3,
        "name":"3category"
    }
]
```

### Get the specified category by name
#### Endpoint
`GET /categories/:name`

#### Path Parameters
| Name |  Type  | Required |      Description      |
| :--- | :----- | :------- | :-------------------- |
| name | string | required | A name of a category. |

#### Request
##### Header
N/A

##### Body
N/A


#### Responsej
##### Header
N/A

##### Body
```json
{
    "id":1,
    "name":"1category"
}
```

### Get all tags
#### Endpoint
`GET /tags`

#### Query Parameters
| Name  |  Type   | Required |       Description        |
| :---- | :------ | :------- | :----------------------- |
| page  | integer | optional | A page number of tags.   |
| limit | integer | optional | A limit number of posts. |

#### Request bodies
##### Header
N/A

##### Body
N/A

#### Response
##### Header
N/A

##### Body
```json
[
    {
        "id":1,
        "name":"1tag"
    },
    {
        "id":2,
        "name":"2tag"
    },
    {
        "id":3,
        "name":"3tag"
    }
]
```

### Get the specified tag by name
#### Endpoint
`GET /tag/:name`

#### Path Parameters
| Name |  Type  | Required |   Description    |
| :--- | :----- | :------- | :--------------- |
| name | string | required | A name of a tag. |

#### Request
##### Header
N/A

##### Body
N/A

#### Response
##### Header
N/A

##### Body
```json
{
    "id": 1,
    "name": "1tag"
}
```

# Private API
Private API is closed api that does required authentication.

| Method |           Endpoint            |                Description                 |
| :----- | :---------------------------- | :----------------------------------------- |
| POST   | /private/signout              | Disable an access token.                   |
| POST   | /private/refresh              | Refresh an access token and refresh token. |
| GET    | /private/me                   | Get the specified admin by access token.   |
| GET    | /private/posts                | Get all posts.                             |
| GET    | /private/posts/:id            | Get the specified post by id.              |
| POST   | /private/posts                | Store a newly post.                        |
| PATCH  | /private/posts/:id            | Update the specified post.                 |
| DELETE | /private//posts/:id           | Remove the specified post.                 |
| GET    | /private//comments            | Get all comments.                          |
| GET    | /private//comments/:id        | Get the specified comment by id.           |
| PATCH  | /private//comments/:id/status | Update the specified comment status.       |
| GET    | /private//categories          | Get all categories.                        |
| GET    | /private//categories/:id      | Get the specified category by id.          |
| POST   | /private//categories          | Store a newly categories.                  |
| PATCH  | /private//categories/:id      | Update the specified category.             |
| DELETE | /private//categories/:id      | Remove the specified category.             |
| GET    | /private//tags                | Get all tags.                              |
| GET    | /private//tags/:id            | Get the specified tag by id.               |
| POST   | /private//tags                | Store a newly tag.                         |
| PATCH  | /private//tags/:id            | Update the specified tag.                  |
| DELETE | /private//tags/:id            | Remove the specified tag.                  |

## Authentication
### Disable an access token
#### Endpoint
`POST /private/signout`

#### Request
##### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

##### Body
N/A

#### Response
##### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

##### Body
```json
{
    "message": "OK"
}
```

### Refresh an acess token and refresh token
#### Endpoint
`POST /private/refresh`

#### Request
##### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | A refresh token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

##### Body
N/A

#### Response
##### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | A refresh token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

##### Body
```json
{
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6IjIxNWZiMWVmLTBiMzQtNDJlMS05NzUwLTRjNTYzZjI1MjE2MSIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTYwMjA4MTkxNywiaWQiOjJ9.hbtuLiStAc6v_y5RBCHA66SEYNlwDdB1z9FOFi54cUo",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDI2ODU4MTcsImlkIjoyLCJyZWZyZXNoX3V1aWQiOiJkODAzMWQwYy1hMTg4LTRmNjgtYTc1NS1jODM1YmFmZjRjYTAifQ.EKQHYvNn8E82HOZT9cGD3WKIhHvetFSNvDT4Qt_7GeQ"
}
```

### Get the specified admin by access token
#### Endpoint
`GET /private/me`

#### Request
##### Header
|     Name      |  Description   |                   Example                   |
| :------------ | :------------- | :------------------------------------------ |
| Authorization | A access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

##### Body
N/A

#### Response
##### Header
|     Name      |  Description   |                   Example                   |
| :------------ | :------------- | :------------------------------------------ |
| Authorization | A access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

##### Body
```json
{
    "id": 1,
    "name": "1admin"
}
```

### Get all posts
#### Endpoint
`GET /private/posts`

#### Query Parameters
| Name  |  Type   | Required |       Description        |
| :---- | :------ | :------- | :----------------------- |
| page  | integer | optional | A page number of posts.  |
| limit | integer | optional | A limit number of posts. |

#### Request
##### Header
|         Name         |        Description         |                   Example                   |
| :------------------- | :------------------------- | :------------------------------------------ |
| Authorization        | An access token            | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |
| Pagination-Count     | A count of records         | 100                                         |
| Pagination-PageCount | A count of page            | 10                                          |
| Pagination-Page      | A current page number      | 5                                           |
| Pagination-Limit     | A limit of number per page | 10                                          |

##### Body
N/A

#### Response
##### Header
|         Name         |        Description         |                   Example                   |
| :------------------- | :------------------------- | :------------------------------------------ |
| Authorization        | An access token            | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |
| Pagination-Count     | A count of records         | 100                                         |
| Pagination-PageCount | A count of page            | 10                                          |
| Pagination-Page      | A current page number      | 5                                           |
| Pagination-Limit     | A limit of number per page | 10                                          |

##### Body
```json
[
    {
        "id":1,
        "admin":{
            "id":2,
            "name":"2admin"
        },
        "category":{
            "id":1,
            "name":"1category",
            "created_at":"0001-01-01T00:00:00Z",
            "updated_at":"0001-01-01T00:00:00Z"
        },
        "tags":[
            {
                "id":1,
                "name":"1tag",
                "created_at":"0001-01-01T00:00:00Z",
                "updated_at":"0001-01-01T00:00:00Z"
            },
            {
                "id":2,
                "name":"2tag",
                "created_at":"0001-01-01T00:00:00Z",
                "updated_at":"0001-01-01T00:00:00Z"
            }
        ],
        "title":"update-post",
        "md_body":"#update-md-body",
        "html_body":"\u003chtml\u003e\u003cbody\u003eupdate_html_body\u003c/body\u003e\u003c/html\u003e",
        "status":"draft",
        "comments":[
            {
                "id":1,
                "post_id":0,
                "body":"1body",
                "status":"",
                "created_at":"2020-10-04T09:23:11Z"
            },
            {
                "id":101,
                "post_id":0,
                "body":"foobar",
                "status":"",
                "created_at":"2020-10-05T12:24:36Z"
            },
            {
                "id":102,
                "post_id":0,
                "body":"foobar",
                "status":"",
                "created_at":"2020-10-05T13:09:54Z"
            }
        ],
        "created_at":"2020-10-07T13:21:57Z",
        "updated_at":"2020-10-07T13:21:58Z"
    },
    {
        "id":2,
        "admin":{
            "id":2,
            "name":"2admin"
        },
        "category":{
            "id":2,
            "name":"2category",
            "created_at":"0001-01-01T00:00:00Z",
            "updated_at":"0001-01-01T00:00:00Z"
        },
        "tags":[
            {
                "id":2,
                "name":"2tag",
                "created_at":"0001-01-01T00:00:00Z",
                "updated_at":"0001-01-01T00:00:00Z"
            }
        ],
        "title":"2title",
        "md_body":"2md_body",
        "html_body":"2html_body",
        "status":"publish",
        "comments":null,
        "created_at":"2020-10-04T16:30:46Z",
        "updated_at":"2020-10-04T16:30:46Z"
    },
    {
        "id":3,
        "admin":{
            "id":3,
            "name":"3admin"
        },
        "category":{
            "id":3,
            "name":"3category",
            "created_at":"0001-01-01T00:00:00Z",
            "updated_at":"0001-01-01T00:00:00Z"
        },
        "tags":[
            {
                "id":3,
                "name":"3tag",
                "created_at":"0001-01-01T00:00:00Z",
                "updated_at":"0001-01-01T00:00:00Z"
            }
        ],
        "title":"3title",
        "md_body":"3md_body",
        "html_body":"3html_body",
        "status":"draft",
        "comments":null,
        "created_at":"2020-10-04T09:23:11Z",
        "updated_at":"2020-10-04T09:23:11Z"
    }
]
```

### Get the specified post by id
#### Endpoint
`GET /private/posts/:id`

#### Path Parameters
| Name | Type | Required |   Description    |
| :--- | :--- | :------- | :--------------- |
| id   | int  | required | An id of a post. |

#### Request
##### Header
N/A

##### Body
N/A

#### Response
##### Header
N/A

##### Body
```json
{
    "id":11,
    "admin":{
        "id":2,
        "name":"2admin"
    },
    "category":{
        "id":1,
        "name":"1category",
        "created_at":"0001-01-01T00:00:00Z",
        "updated_at":"0001-01-01T00:00:00Z"
    },
    "tags":[
        {
            "id":1,
            "name":"1tag",
            "created_at":"0001-01-01T00:00:00Z",
            "updated_at":"0001-01-01T00:00:00Z"
        },
        {
            "id":2,
            "name":"2tag",
            "created_at":"0001-01-01T00:00:00Z",
            "updated_at":"0001-01-01T00:00:00Z"
        }
    ],
    "title":"1title",
    "md_body":"#md-body",
    "html_body":"\u003chtml\u003e\u003cbody\u003ehtml_body\u003c/body\u003e\u003c/html\u003e",
    "status":"publish",
    "comments":null,
    "created_at":"2020-09-28T15:18:16Z",
    "updated_at":"2020-09-28T15:18:16Z"
}
```

### Store a newly post
#### Endpoint
`POST /private/posts`

#### Request
##### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

##### Body
```json
{
    "admin_id": 1,
	"category_id": 1,
	"tags": [
		{
			"id": 1
		},
		{
			"id": 2
		}
	],
	"title": "foobar",
	"body": "foo-bar-body",
	"md_body": "#md-body",
	"html_body": "<html><body>html_body</body></html>",
	"status": "draft|publish"
}
```

#### Response
##### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

##### Body
```json
{
    "id":101,
    "admin":{
        "id":2,
        "name":"2admin"
    },
    "category":{
        "id":1,
        "name":"1category",
        "created_at":"0001-01-01T00:00:00Z",
        "updated_at":"0001-01-01T00:00:00Z"
    },
    "tags":[
        {
            "id":1,
            "name":"1tag",
            "created_at":"0001-01-01T00:00:00Z",
            "updated_at":"0001-01-01T00:00:00Z"
        },
        {
            "id":2,
            "name":"2tag",
            "created_at":"0001-01-01T00:00:00Z",
            "updated_at":"0001-01-01T00:00:00Z"
        }
    ],
    "title":"jwt",
    "md_body":"#md-body",
    "html_body":"\u003chtml\u003e\u003cbody\u003ehtml_body\u003c/body\u003e\u003c/html\u003e",
    "status":"draft",
    "comments":null,
    "created_at":"2020-10-07T13:10:33Z",
    "updated_at":"2020-10-07T13:10:33Z"
}
```

### Update the specified post
#### Endpoint
`PATCH /private/posts/:id`

#### Query Parameters
| Name | Type | Required |   Description    |
| :--- | :--- | :------- | :--------------- |
| id   | int  | required | An id of a post. |

#### Request
##### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

##### Body
```json
{
    "admin_id": 1,
	"category_id": 1,
	"tags": [
		{
			"id": 1
		},
		{
			"id": 2
		}
	],
	"title": "foobar",
	"body": "foo-bar-body",
	"md_body": "#md-body",
	"html_body": "<html><body>html_body</body></html>",
	"status": "draft|publish"
}
```

##### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

##### Body
```json
{
    "id":1,
    "admin":{
        "id":2,
        "name":"2admin"
    },
    "category":{
        "id":1,
        "name":"1category",
        "created_at":"0001-01-01T00:00:00Z",
        "updated_at":"0001-01-01T00:00:00Z"
    },
    "tags":[
        {
            "id":1,
            "name":"1tag",
            "created_at":"0001-01-01T00:00:00Z",
            "updated_at":"0001-01-01T00:00:00Z"
        },
        {
            "id":2,
            "name":"2tag",
            "created_at":"0001-01-01T00:00:00Z",
            "updated_at":"0001-01-01T00:00:00Z"
        }
    ],
    "title":"update-post",
    "md_body":"#update-md-body",
    "html_body":"\u003chtml\u003e\u003cbody\u003eupdate_html_body\u003c/body\u003e\u003c/html\u003e",
    "status":"draft",
    "comments":null,
    "created_at":"2020-10-07T13:21:57Z",
    "updated_at":"2020-10-07T13:21:58Z"
}
```

### Remove the specified post
#### Endpoint
`DELETE /private/posts/:id`

#### Query Parameters
| Name | Type | Required |   Description    |
| :--- | :--- | :------- | :--------------- |
| id   | int  | required | An id of a post. |
j
#### Request
##### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

##### Body
N/A

#### Response
##### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

##### Body
```json
{
    "message": "OK"
}
```

### Get all comments
#### Endpoint
`GET /private/comments`

#### Query Parameters
| Name  |  Type   | Required |         Description         |
| :---- | :------ | :------- | :-------------------------- |
| page  | integer | optional | A page number of comments.  |
| limit | integer | optional | A limit number of comments. |

#### Request
##### Header
|         Name         |        Description         | Example |
| :------------------- | :------------------------- | :------ |
| Pagination-Count     | A count of records         | 100     |
| Pagination-PageCount | A count of page            | 10      |
| Pagination-Page      | A current page number      | 5       |
| Pagination-Limit     | A limit of number per page | 10      |

##### Body
N/A

#### Response
##### Header
|         Name         |        Description         | Example |
| :------------------- | :------------------------- | :------ |
| Pagination-Count     | A count of records         | 100     |
| Pagination-PageCount | A count of page            | 10      |
| Pagination-Page      | A current page number      | 5       |
| Pagination-Limit     | A limit of number per page | 10      |

##### Body
```json
[
    {
        "id":1,
        "post_id":11,
        "body":"foobar1",
        "status":"pending",
        "created_at":"2020-09-28T15:20:03Z"
    },
    {
        "id":2,
        "post_id":11,
        "body":"foobar2",
        "status":"pending",
        "created_at":"2020-09-28T15:20:03Z"
    },
    {
        "id":3,
        "post_id":11,
        "body":"foobar3",
        "status":"approval",
        "created_at":"2020-09-28T15:20:03Z"
    }
]
```

### Get the specified comment by id
#### Endpoint
`GET /private/comments/:id`

#### Path Parameters
| Name | Type | Required |     Description     |
| :--- | :--- | :------- | :------------------ |
| id   | int  | required | An id of a comment. |

#### Request
##### Header
N/A

##### Body
N/A

#### Response
##### Header
N/A

##### Body
```json
{
    "id":1,
    "post_id":11,
    "body":"foobar",
    "status":"pending",
    "created_at":"2020-09-28T15:20:03Z"
}
```

## Update the specified comment status
#### Endpoint
`PATCH /private/comments/:id/status`

#### Query Parameters
| Name |  Type   | Required |     Description     |
| :--- | :------ | :------- | :------------------ |
| id   | integer | required | An id of a comment. |

#### Request
##### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

##### Body
```json
{
	"status": "pending|approval"
}
```

#### Response
##### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

##### Body
```json
{
    "id": 1,
    "post_id": 1,
    "body": "body",
    "status": "pending",
    "created_at": "2020-10-07T14:11:24Z"
}
```

### Get all categories
#### Endpoint
`GET /private/categories`

#### Query Parameters
| Name  |  Type   | Required |          Description          |
| :---- | :------ | :------- | :---------------------------- |
| page  | integer | optional | A page number of categories.  |
| limit | integer | optional | A limit number of categories. |

#### Request
##### Header
|         Name         |        Description         | Example |
| :------------------- | :------------------------- | :------ |
| Pagination-Count     | A count of records         | 100     |
| Pagination-PageCount | A count of page            | 10      |
| Pagination-Page      | A current page number      | 5       |
| Pagination-Limit     | A limit of number per page | 10      |

##### Body
N/A

#### Response
##### Header
|         Name         |        Description         | Example |
| :------------------- | :------------------------- | :------ |
| Pagination-Count     | A count of records         | 100     |
| Pagination-PageCount | A count of page            | 10      |
| Pagination-Page      | A current page number      | 5       |
| Pagination-Limit     | A limit of number per page | 10      |

##### Body
```json
[
    {
        "id":1,
        "name":"1category",
        "created_at":"2020-02-23T07:10:53Z",
        "updated_at":"2020-02-23T07:10:53Z"
    },
    {
        "id":2,
        "name":"2category",
        "created_at":"2020-02-23T07:10:53Z",
        "updated_at":"2020-02-23T07:10:53Z"
    },
    {
        "id":3,
        "name":"3category",
        "created_at":"2020-02-23T07:10:53Z",
        "updated_at":"2020-02-23T07:10:53Z"
    }
]
```

### Get the specified category by id
#### Endpoint
`GET /private/categories/:id`

#### Path Parameters
| Name | Type | Required |     Description      |
| :--- | :--- | :------- | :------------------- |
| id   | int  | required | An id of a category. |

#### Request
##### Header
N/A

##### Body
N/A

#### Response
##### Header
N/A

##### Body
```json
{
    "id":1,
    "name":"1category",
    "created_at":"2020-02-23T07:10:53Z",
    "updated_at":"2020-02-23T07:10:53Z"
}
```

### Get all tags
#### Endpoint
`GET /private/tags`

#### Query Parameters
| Name  |  Type   | Required |       Description       |
| :---- | :------ | :------- | :---------------------- |
| page  | integer | optional | A page number of tags.  |
| limit | integer | optional | A limit number of tags. |

#### Response
##### Header
|         Name         |        Description         | Example |
| :------------------- | :------------------------- | :------ |
| Pagination-Count     | A count of records         | 100     |
| Pagination-PageCount | A count of page            | 10      |
| Pagination-Page      | A current page number      | 5       |
| Pagination-Limit     | A limit of number per page | 10      |

##### Body
N/A

#### Request
##### Header
|         Name         |        Description         | Example |
| :------------------- | :------------------------- | :------ |
| Pagination-Count     | A count of records         | 100     |
| Pagination-PageCount | A count of page            | 10      |
| Pagination-Page      | A current page number      | 5       |
| Pagination-Limit     | A limit of number per page | 10      |

##### Body
```json
[
    {
        "id":1,
        "name":"1tag",
        "created_at":"2020-09-28T15:17:26Z",
        "updated_at":"2020-09-28T15:17:26Z"
    },
    {
        "id":2,
        "name":"2tag",
        "created_at":"2020-02-23T07:24:19Z",
        "updated_at":"2020-02-23T07:24:19Z"
    },
    {
        "id":3,
        "name":"3tag",
        "created_at":"2020-02-23T07:24:19Z",
        "updated_at":"2020-02-23T07:24:19Z"
    }
]
```

### Get the specified tag by id
#### Endpoint
`GET /private/tags/:id`

#### Path Parameters
| Name | Type | Required |   Description   |
| :--- | :--- | :------- | :-------------- |
| id   | int  | required | An id of a tag. |

#### Request
##### Header
N/A

##### Body
N/A

#### Request
##### Header
N/A

##### Body
```json
{
    "id":1,
    "name":"1tag",
    "created_at":"2020-09-28T15:17:26Z",
    "updated_at":"2020-09-28T15:17:26Z"
}
```

### Store a newly category
#### Endpoint
`POST /private/categories`

#### Request
##### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

##### Body
```json
{
    "name": "foo"
}
```

#### Response
##### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

##### Body
```json
{
    "id":1,
    "name":"new-category",
    "created_at": "2020-10-04T09:23:11Z",
    "updated_at": "2020-10-04T09:23:11Z"
}
```

### Update the specified category
#### Endpoint
`PATCH /private/categories/:id`

#### Query Parameters
| Name | Type | Required |     Description      |
| :--- | :--- | :------- | :------------------- |
| id   | int  | required | An id of a category. |

#### Request
##### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

##### Body
```json
{
    "id": 1,
    "name": "update-category",
    "created_at": "2020-10-06T16:06:50Z",
    "updated_at": "2020-10-06T16:06:50Z"
}
```

#### Response
##### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

##### Body
```json
{
    "id": 1,
    "name": "update-category",
    "created_at": "2020-10-07T14:32:01Z",
    "updated_at": "2020-10-07T14:32:01Z"
}
```

### Remove the specified category
#### Endpoint
`PATCH /private/categories/:id`

#### Query Parameters
| Name | Type | Required |     Description      |
| :--- | :--- | :------- | :------------------- |
| id   | int  | required | An id of a category. |

#### Request
##### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

##### Body
N/A

#### Response
##### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

##### Body
```json
{
    "message": "OK"
}
```

### Store a newly tag
#### Endpoint
`POST /private/tags`

#### Request
##### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

##### Body
```json
{
    "name": "foo"
}
```

#### Response
##### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

##### Body
```json
{
    "id": 1,
    "name": "new-tag",
    "created_at": "2020-10-06T16:24:10Z",
    "updated_at": "2020-10-06T16:24:10Z"
}
```

### Update the specified tag
#### Endpoint
`PATCH /private/tags/:id`

#### Query Parameters
| Name | Type | Required |   Description   |
| :--- | :--- | :------- | :-------------- |
| id   | int  | required | An id of a tag. |

#### Request
##### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

##### Body
```json
{
    "name": "foo"
}
```

#### Response
##### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

##### Body
```json
{
    "id": 1,
    "name": "update-tag",
    "created_at": "2020-10-06T16:27:01Z",
    "updated_at": "2020-10-06T16:27:02Z"
}
```

### Remove the specified tag
#### Endpoint
`DELETE /private/tags/:id`

#### Query Parameters
| Name | Type | Required |   Description   |
| :--- | :--- | :------- | :-------------- |
| id   | int  | required | An id of a tag. |

#### Request
##### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

##### Body
N/A

#### Response
##### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

##### Body
```json
{
    "message": "OK"
}
```

