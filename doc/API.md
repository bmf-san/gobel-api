# API
Refered to [cryptlex/rest-api-response-format](https://github.com/cryptlex/rest-api-response-format).

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
  - [Post](#post-2)
    - [Get all publish posts](#get-all-publish-posts)
      - [Endpoint](#endpoint-1)
      - [Header](#header)
      - [Query Parameters](#query-parameters)
    - [Get all publish posts by category](#get-all-publish-posts-by-category)
      - [Endpoint](#endpoint-2)
      - [Path Parameters](#path-parameters)
    - [Get all publish posts by tag](#get-all-publish-posts-by-tag)
      - [Endpoint](#endpoint-3)
      - [Path Parameters](#path-parameters-1)
    - [Get the specified post by title](#get-the-specified-post-by-title)
      - [Endpoint](#endpoint-4)
      - [Path Parameters](#path-parameters-2)
    - [Store a newly comment](#store-a-newly-comment)
      - [Endpoint](#endpoint-5)
      - [Request bodies](#request-bodies-1)
      - [Path Parameters](#path-parameters-3)
    - [Get all categories](#get-all-categories)
      - [Endpoint](#endpoint-6)
      - [Query Parameters](#query-parameters-1)
    - [Get the specified category by name](#get-the-specified-category-by-name)
      - [Endpoint](#endpoint-7)
      - [Path Parameters](#path-parameters-4)
    - [Get all tags](#get-all-tags)
      - [Endpoint](#endpoint-8)
      - [Query Parameters](#query-parameters-2)
    - [Get the specified tag by name](#get-the-specified-tag-by-name)
      - [Endpoint](#endpoint-9)
      - [Path Parameters](#path-parameters-5)
- [Private API](#private-api)
  - [Authentication](#authentication-1)
    - [Disable an access token](#disable-an-access-token)
      - [Endpoint](#endpoint-10)
      - [Header](#header-1)
    - [Refresh an acess token and refresh token](#refresh-an-acess-token-and-refresh-token)
      - [Endpoint](#endpoint-11)
      - [Header](#header-2)
    - [Get the specified admin by access token](#get-the-specified-admin-by-access-token)
      - [Endpoint](#endpoint-12)
      - [Header](#header-3)
    - [Get all posts](#get-all-posts)
      - [Endpoint](#endpoint-13)
      - [Header](#header-4)
      - [Query Parameters](#query-parameters-3)
    - [Get the specified post by id](#get-the-specified-post-by-id)
      - [Endpoint](#endpoint-14)
      - [Path Parameters](#path-parameters-6)
    - [Store a newly comment](#store-a-newly-comment-1)
      - [Endpoint](#endpoint-15)
      - [Request bodies](#request-bodies-2)
    - [Store a newly post](#store-a-newly-post)
      - [Endpoint](#endpoint-16)
      - [Header](#header-5)
      - [Request bodies](#request-bodies-3)
    - [Update the specified post](#update-the-specified-post)
      - [Endpoint](#endpoint-17)
      - [Header](#header-6)
      - [Query Parameters](#query-parameters-4)
      - [Request bodies](#request-bodies-4)
    - [Remove the specified post](#remove-the-specified-post)
      - [Endpoint](#endpoint-18)
      - [Header](#header-7)
      - [Query Parameters](#query-parameters-5)
    - [Get all comments](#get-all-comments)
      - [Endpoint](#endpoint-19)
      - [Header](#header-8)
      - [Query Parameters](#query-parameters-6)
    - [Get the specified comment by id](#get-the-specified-comment-by-id)
      - [Endpoint](#endpoint-20)
      - [Path Parameters](#path-parameters-7)
  - [Update the specified comment status](#update-the-specified-comment-status)
      - [Endpoint](#endpoint-21)
      - [Header](#header-9)
      - [Query Parameters](#query-parameters-7)
      - [Request bodies](#request-bodies-5)
    - [Get all categories](#get-all-categories-1)
      - [Endpoint](#endpoint-22)
      - [Header](#header-10)
      - [Query Parameters](#query-parameters-8)
    - [Get the specified category by id](#get-the-specified-category-by-id)
      - [Endpoint](#endpoint-23)
      - [Path Parameters](#path-parameters-8)
    - [Get all tags](#get-all-tags-1)
      - [Endpoint](#endpoint-24)
      - [Header](#header-11)
      - [Query Parameters](#query-parameters-9)
    - [Get the specified tag by id](#get-the-specified-tag-by-id)
      - [Endpoint](#endpoint-25)
      - [Path Parameters](#path-parameters-9)
    - [Store a newly category](#store-a-newly-category)
      - [Endpoint](#endpoint-26)
      - [Header](#header-12)
      - [Request bodies](#request-bodies-6)
    - [Update the specified category](#update-the-specified-category)
      - [Endpoint](#endpoint-27)
      - [Header](#header-13)
      - [Query Parameters](#query-parameters-10)
      - [Request bodies](#request-bodies-7)
    - [Remove the specified category](#remove-the-specified-category)
      - [Endpoint](#endpoint-28)
      - [Header](#header-14)
      - [Query Parameters](#query-parameters-11)
    - [Store a newly tag](#store-a-newly-tag)
      - [Endpoint](#endpoint-29)
      - [Header](#header-15)
      - [Request bodies](#request-bodies-8)
    - [Update the specified tag](#update-the-specified-tag)
      - [Endpoint](#endpoint-30)
      - [Header](#header-16)
      - [Query Parameters](#query-parameters-12)
      - [Request bodies](#request-bodies-9)
    - [Remove the specified tag](#remove-the-specified-tag)
      - [Endpoint](#endpoint-31)
      - [Header](#header-17)
      - [Query Parameters](#query-parameters-13)

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
Location: /posts/
Content-Type: application/json

{
    "message": "The item was created successfully"
}
```

#### PATCH
If updated entity is to be sent after the update.

```json
HTTP/1.1  200
Content-Type: application/json

{
    "message": "The item was updated successfully"
}
```

If updated entity is not to be sent after the update
```json
HTTP/1.1  204
```

#### DELETE
```json
HTTP/1.1  204
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

## Post
### Get all publish posts
#### Endpoint
`GET /posts`

#### Header
|         Name         |        Description         | Example |
| :------------------- | :------------------------- | :------ |
| Pagination-Count     | A count of records         | 100     |
| Pagination-PageCount | A count of page            | 10      |
| Pagination-Page      | A current page number      | 5       |
| Pagination-Limit     | A limit of number per page | 10      |

#### Query Parameters
| Name  |  Type   | Required |       Description        |
| :---- | :------ | :------- | :----------------------- |
| page  | integer | optional | A page number of posts.  |
| limit | integer | optional | A limit number of posts. |

### Get all publish posts by category
#### Endpoint
`GET /posts/categories/:name`

#### Path Parameters
| Name |  Type  | Required |      Description      |
| :--- | :----- | :------- | :-------------------- |
| name | string | required | A name of a category. |

### Get all publish posts by tag
#### Endpoint
`GET /posts/tags/:name`

#### Path Parameters
| Name |  Type  | Required |   Description    |
| :--- | :----- | :------- | :--------------- |
| name | string | required | A name of a tag. |

### Get the specified post by title
#### Endpoint
`GET /posts/:title`

#### Path Parameters
| Name  |  Type  | Required |    Description     |
| :---- | :----- | :------- | :----------------- |
| title | string | required | A title of a post. |

### Store a newly comment
#### Endpoint
`POST /posts/:title/comments`

#### Request bodies
```json
{
	"body": "foobar"
}
```

#### Path Parameters
| Name  |  Type  | Required |    Description     |
| :---- | :----- | :------- | :----------------- |
| title | string | required | A title of a post. |


### Get all categories
#### Endpoint
`GET /categories`

#### Query Parameters
| Name  |  Type   | Required |         Description          |
| :---- | :------ | :------- | :--------------------------- |
| page  | integer | optional | A page number of categories. |
| limit | integer | optional | A limit number of posts.     |

### Get the specified category by name
#### Endpoint
`GET /categories/:name`

#### Path Parameters
| Name |  Type  | Required |      Description      |
| :--- | :----- | :------- | :-------------------- |
| name | string | required | A name of a category. |

### Get all tags
#### Endpoint
`GET /tags`

#### Query Parameters
| Name  |  Type   | Required |       Description        |
| :---- | :------ | :------- | :----------------------- |
| page  | integer | optional | A page number of tags.   |
| limit | integer | optional | A limit number of posts. |

### Get the specified tag by name
#### Endpoint
`GET /tag/:name`

#### Path Parameters
| Name |  Type  | Required |   Description    |
| :--- | :----- | :------- | :--------------- |
| name | string | required | A name of a tag. |

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

#### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

### Refresh an acess token and refresh token
#### Endpoint
`POST /private/refresh`

#### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | A refresh token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

### Get the specified admin by access token
#### Endpoint
`GET /private/me`

#### Header
|     Name      |  Description   |                   Example                   |
| :------------ | :------------- | :------------------------------------------ |
| Authorization | A access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

### Get all posts
#### Endpoint
`GET /private/posts`

#### Header
|         Name         |        Description         |                   Example                   |
| :------------------- | :------------------------- | :------------------------------------------ |
| Authorization        | An access token            | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |
| Pagination-Count     | A count of records         | 100                                         |
| Pagination-PageCount | A count of page            | 10                                          |
| Pagination-Page      | A current page number      | 5                                           |
| Pagination-Limit     | A limit of number per page | 10                                          |

#### Query Parameters
| Name  |  Type   | Required |       Description        |
| :---- | :------ | :------- | :----------------------- |
| page  | integer | optional | A page number of posts.  |
| limit | integer | optional | A limit number of posts. |

### Get the specified post by id
#### Endpoint
`GET /private/posts/:id`

#### Path Parameters
| Name | Type | Required |   Description    |
| :--- | :--- | :------- | :--------------- |
| id   | int  | required | An id of a post. |

### Store a newly comment
#### Endpoint
`POST /private/comment`

#### Request bodies
```json
{
    "post_id": 1,
	"body": "foobar",
}
```
### Store a newly post
#### Endpoint
`POST /private/posts`

#### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

#### Request bodies
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

### Update the specified post
#### Endpoint
`PATCH /private/posts/:id`

#### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

#### Query Parameters
| Name | Type | Required |   Description    |
| :--- | :--- | :------- | :--------------- |
| id   | int  | required | An id of a post. |

#### Request bodies
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

### Remove the specified post
#### Endpoint
`DELETE /private/posts/:id`

#### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

#### Query Parameters
| Name | Type | Required |   Description    |
| :--- | :--- | :------- | :--------------- |
| id   | int  | required | An id of a post. |

### Get all comments
#### Endpoint
`GET /private/comments`

#### Header
|         Name         |        Description         | Example |
| :------------------- | :------------------------- | :------ |
| Pagination-Count     | A count of records         | 100     |
| Pagination-PageCount | A count of page            | 10      |
| Pagination-Page      | A current page number      | 5       |
| Pagination-Limit     | A limit of number per page | 10      |

#### Query Parameters
| Name  |  Type   | Required |         Description         |
| :---- | :------ | :------- | :-------------------------- |
| page  | integer | optional | A page number of comments.  |
| limit | integer | optional | A limit number of comments. |

### Get the specified comment by id
#### Endpoint
`GET /private/comments/:id`

#### Path Parameters
| Name | Type | Required |     Description     |
| :--- | :--- | :------- | :------------------ |
| id   | int  | required | An id of a comment. |

## Update the specified comment status
#### Endpoint
`PATCH /private/comments/:id/status`

#### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

#### Query Parameters
| Name |  Type   | Required |     Description     |
| :--- | :------ | :------- | :------------------ |
| id   | integer | required | An id of a comment. |

#### Request bodies
```json
{
    "status": "approval|pending"
}
```

### Get all categories
#### Endpoint
`GET /private/categories`

#### Header
|         Name         |        Description         | Example |
| :------------------- | :------------------------- | :------ |
| Pagination-Count     | A count of records         | 100     |
| Pagination-PageCount | A count of page            | 10      |
| Pagination-Page      | A current page number      | 5       |
| Pagination-Limit     | A limit of number per page | 10      |

#### Query Parameters
| Name  |  Type   | Required |          Description          |
| :---- | :------ | :------- | :---------------------------- |
| page  | integer | optional | A page number of categories.  |
| limit | integer | optional | A limit number of categories. |

### Get the specified category by id
#### Endpoint
`GET /private/categories/:id`

#### Path Parameters
| Name | Type | Required |     Description      |
| :--- | :--- | :------- | :------------------- |
| id   | int  | required | An id of a category. |

### Get all tags
#### Endpoint
`GET /private/tags`

#### Header
|         Name         |        Description         | Example |
| :------------------- | :------------------------- | :------ |
| Pagination-Count     | A count of records         | 100     |
| Pagination-PageCount | A count of page            | 10      |
| Pagination-Page      | A current page number      | 5       |
| Pagination-Limit     | A limit of number per page | 10      |

#### Query Parameters
| Name  |  Type   | Required |       Description       |
| :---- | :------ | :------- | :---------------------- |
| page  | integer | optional | A page number of tags.  |
| limit | integer | optional | A limit number of tags. |

### Get the specified tag by id
#### Endpoint
`GET /private/tags/:id`

#### Path Parameters
| Name | Type | Required |   Description   |
| :--- | :--- | :------- | :-------------- |
| id   | int  | required | An id of a tag. |

### Store a newly category
#### Endpoint
`POST /private/categories`

#### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

#### Request bodies
```json
{
    "name": "foo"
}
```

### Update the specified category
#### Endpoint
`PATCH /private/categories/:id`

#### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

#### Query Parameters
| Name | Type | Required |     Description      |
| :--- | :--- | :------- | :------------------- |
| id   | int  | required | An id of a category. |

#### Request bodies
```json
{
    "name": "foo"
}
```

### Remove the specified category
#### Endpoint
`PATCH /private/categories/:id`

#### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

#### Query Parameters
| Name | Type | Required |     Description      |
| :--- | :--- | :------- | :------------------- |
| id   | int  | required | An id of a category. |

### Store a newly tag
#### Endpoint
`POST /private/tags`

#### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

#### Request bodies
```json
{
    "name": "foo"
}
```

### Update the specified tag
#### Endpoint
`PATCH /private/tags/:id`

#### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

#### Query Parameters
| Name | Type | Required |   Description   |
| :--- | :--- | :------- | :-------------- |
| id   | int  | required | An id of a tag. |

#### Request bodies
```json
{
    "name": "foo"
}
```

### Remove the specified tag
#### Endpoint
`DELETE /private/tags/:id`

#### Header
|     Name      |   Description   |                   Example                   |
| :------------ | :-------------- | :------------------------------------------ |
| Authorization | An access token | Bearer e856e7bd-2572-4890-b9e0-a79ea09cd431 |

#### Query Parameters
| Name | Type | Required |   Description   |
| :--- | :--- | :------- | :-------------- |
| id   | int  | required | An id of a tag. |
