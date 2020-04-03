# API
Refered to [cryptlex/rest-api-response-format](https://github.com/cryptlex/rest-api-response-format).

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

| Method |        Endpoint         |             Description              |
| :----- | :---------------------- | :----------------------------------- |
| POST   | /authenticate           | Get a json web token by credentials. |
| GET    | /posts                  | Get all publish posts.               |
| GET    | /posts/categories/:name | Get all publish posts by cagtegory. |
| GET    | /posts/tags/:name       | Get all publish posts by tag.       |
| GET    | /posts/:title           | Get the specified post by title.     |
| POST   | /posts/:title/comments  | Store a newly comment.               |
| GET    | /categories             | Get all categories.                  |
| GET    | /categories/:name       | Get the specified category by name.  |
| GET    | /tags                   | Get all tags.                        |
| GET    | /tags/:name             | Get the specified tag by name.       |

// TODO: search api

## Authentication
### Get a json web token by credentials
#### Endpoint
`POST /authenticate`

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

| Method |       Endpoint       |             Description              |
| :----- | :------------------- | :----------------------------------- |
| GET    | /posts               | Get all posts.                       |
| GET    | /posts/:id           | Get the specified post by id.        |
| POST   | /posts               | Store a newly post.                  |
| PATCH  | /posts/:id           | Update the specified post.           |
| DELETE | /posts/:id           | Remove the specified post.           |
| GET    | /comments            | Get all comments.                    |
| GET    | /comments/:id        | Get the specified comment by id.     |
| PATCH  | /comments/:id/status | Update the specified comment status. |
| GET    | /categories          | Get all categories.                  |
| GET    | /categories/:id      | Get the specified category by id.    |
| POST   | /categories          | Store a newly categories.            |
| PATCH  | /categories/:id      | Update the specified category.       |
| DELETE | /categories/:id      | Remove the specified category.       |
| GET    | /tags                | Get all tags.                        |
| GET    | /tags/:id            | Get the specified tag by id.         |
| POST   | /tags                | Store a newly tag.                   |
| PATCH  | /tags/:id            | Update the specified tag.            |
| DELETE | /tags/:id            | Remove the specified tag.            |

### Get all posts
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

### Get the specified post by id
#### Endpoint
`GET /posts/:id`

#### Path Parameters
| Name | Type | Required |   Description    |
| :--- | :--- | :------- | :--------------- |
| id   | int  | required | An id of a post. |

### Store a newly comment
#### Endpoint
`POST /comment`

#### Request bodies
```json
{
    "post_id": 1,
	"body": "foobar",
}
```
### Store a newly post
#### Endpoint
`POST /posts`

#### Header
|     Name      |            Description            |        Example        |
| :------------ | :-------------------------------- | :-------------------- |
| Authorization | An bearer token for authorization | Bearer abcd.efgh.ijkl |

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
`PATCH /posts/:id`

#### Header
|     Name      |            Description            |        Example        |
| :------------ | :-------------------------------- | :-------------------- |
| Authorization | An bearer token for authorization | Bearer abcd.efgh.ijkl |

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
`DELETE /posts/:id`

#### Header
|     Name      |            Description            |        Example        |
| :------------ | :-------------------------------- | :-------------------- |
| Authorization | An bearer token for authorization | Bearer abcd.efgh.ijkl |

#### Query Parameters
| Name | Type | Required |   Description    |
| :--- | :--- | :------- | :--------------- |
| id   | int  | required | An id of a post. |

### Get all comments
#### Endpoint
`GET /comments`

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
`GET /comments/:id`

#### Path Parameters
| Name | Type | Required |     Description     |
| :--- | :--- | :------- | :------------------ |
| id   | int  | required | An id of a comment. |

## Update the specified comment status
#### Endpoint
`PATCH /comments/:id/status`

#### Header
|     Name      |            Description            |        Example        |
| :------------ | :-------------------------------- | :-------------------- |
| Authorization | An bearer token for authorization | Bearer abcd.efgh.ijkl |

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
`GET /categories`

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
`GET /categories/:id`

#### Path Parameters
| Name | Type | Required |     Description      |
| :--- | :--- | :------- | :------------------- |
| id   | int  | required | An id of a category. |

### Get all tags
#### Endpoint
`GET /tags`

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
`GET /tags/:id`

#### Path Parameters
| Name | Type | Required |   Description   |
| :--- | :--- | :------- | :-------------- |
| id   | int  | required | An id of a tag. |

### Store a newly category
#### Endpoint
`POST /categories`

#### Header
|     Name      |            Description            |        Example        |
| :------------ | :-------------------------------- | :-------------------- |
| Authorization | An bearer token for authorization | Bearer abcd.efgh.ijkl |

#### Request bodies
```json
{
    "name": "foo"
}
```

### Update the specified category
#### Endpoint
`PATCH /categories/:id`

#### Header
|     Name      |            Description            |        Example        |
| :------------ | :-------------------------------- | :-------------------- |
| Authorization | An bearer token for authorization | Bearer abcd.efgh.ijkl |

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
`PATCH /categories/:id`

#### Header
|     Name      |            Description            |        Example        |
| :------------ | :-------------------------------- | :-------------------- |
| Authorization | An bearer token for authorization | Bearer abcd.efgh.ijkl |

#### Query Parameters
| Name | Type | Required |     Description      |
| :--- | :--- | :------- | :------------------- |
| id   | int  | required | An id of a category. |

### Store a newly tag
#### Endpoint
`POST /tags`

#### Header
|     Name      |            Description            |        Example        |
| :------------ | :-------------------------------- | :-------------------- |
| Authorization | An bearer token for authorization | Bearer abcd.efgh.ijkl |

#### Request bodies
```json
{
    "name": "foo"
}
```

### Update the specified tag
#### Endpoint
`PATCH /tags/:id`

#### Header
|     Name      |            Description            |        Example        |
| :------------ | :-------------------------------- | :-------------------- |
| Authorization | An bearer token for authorization | Bearer abcd.efgh.ijkl |

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
`DELETE /tags/:id`

#### Header
|     Name      |            Description            |        Example        |
| :------------ | :-------------------------------- | :-------------------- |
| Authorization | An bearer token for authorization | Bearer abcd.efgh.ijkl |

#### Query Parameters
| Name | Type | Required |   Description   |
| :--- | :--- | :------- | :-------------- |
| id   | int  | required | An id of a tag. |
