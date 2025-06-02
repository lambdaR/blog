# REST API

The Micro Blog system exposes a REST API through the Web Service. This API allows clients to interact with all features of the blog system.

## Base URL

All API endpoints are available at:

```
http://localhost:42096
```

## Authentication

Some endpoints require authentication. Authentication is handled through cookies:

1. Call the `/login` or `/signup` endpoint to authenticate
2. The server will set a session cookie
3. Include this cookie in subsequent requests

## API Endpoints

### Posts

#### List Posts

```
GET /posts
```

**Response:**
```json
{
  "posts": [
    {
      "id": "post-id",
      "title": "Post Title",
      "content": "Post content...",
      "author_id": "user-id",
      "author_name": "User Name",
      "created_at": 1625097600,
      "tags": ["tag1", "tag2"]
    }
  ],
  "total": 1
}
```

#### Get Post by ID

```
GET /posts/:id
```

**Response:**
```json
{
  "post": {
    "id": "post-id",
    "title": "Post Title",
    "content": "Post content...",
    "author_id": "user-id",
    "author_name": "User Name",
    "created_at": 1625097600,
    "tags": ["tag1", "tag2"]
  }
}
```

#### Create Post

```
POST /posts
```

**Request Body:**
```json
{
  "title": "Post Title",
  "content": "Post content..."
}
```

**Response:**
```json
{
  "post": {
    "id": "post-id",
    "title": "Post Title",
    "content": "Post content...",
    "author_id": "user-id",
    "author_name": "User Name",
    "created_at": 1625097600,
    "tags": []
  }
}
```

**Note:** Requires authentication.

### Comments

#### List Comments

```
GET /comments
```

**Query Parameters:**
- `post_id` (optional): Filter comments by post ID

**Response:**
```json
{
  "comments": [
    {
      "id": "comment-id",
      "content": "Comment content...",
      "author_id": "user-id",
      "author_name": "User Name",
      "post_id": "post-id",
      "created_at": 1625097600
    }
  ],
  "total": 1
}
```

#### Create Comment

```
POST /comments
```

**Request Body:**
```json
{
  "content": "Comment content...",
  "post_id": "post-id"
}
```

**Response:**
```json
{
  "comment": {
    "id": "comment-id",
    "content": "Comment content...",
    "author_id": "user-id",
    "author_name": "User Name",
    "post_id": "post-id",
    "created_at": 1625097600
  }
}
```

**Note:** Requires authentication.

### Users

#### List Users

```
GET /users
```

**Response:**
```json
{
  "users": [
    {
      "id": "user-id",
      "name": "User Name",
      "email": "user@example.com"
    }
  ],
  "total": 1
}
```

#### Get User by ID

```
GET /users/:id
```

**Response:**
```json
{
  "user": {
    "id": "user-id",
    "name": "User Name",
    "email": "user@example.com"
  }
}
```

#### Create User

```
POST /users
```

**Request Body:**
```json
{
  "name": "User Name",
  "email": "user@example.com"
}
```

**Response:**
```json
{
  "user": {
    "id": "user-id",
    "name": "User Name",
    "email": "user@example.com"
  }
}
```

### Authentication

#### Sign Up

```
POST /signup
```

**Request Body:**
```json
{
  "name": "User Name",
  "email": "user@example.com",
  "password": "password123"
}
```

**Response:**
```json
{
  "user": {
    "id": "user-id",
    "name": "User Name",
    "email": "user@example.com"
  }
}
```

**Note:** This endpoint also logs the user in (sets a session cookie).

#### Log In

```
POST /login
```

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response:**
```json
{
  "user": {
    "id": "user-id",
    "name": "User Name",
    "email": "user@example.com"
  }
}
```

#### Log Out

```
POST /logout
```

**Response:**
```json
{
  "message": "logged out"
}
```

#### Get Current User

```
GET /users/me
```

**Response (authenticated):**
```json
{
  "user": {
    "id": "user-id",
    "name": "User Name"
  }
}
```

**Response (not authenticated):**
```json
{
  "user": null
}
```

### Tags

#### Add Tag to Post

```
POST /posts/:id/tags
```

**Request Body:**
```json
{
  "tag": "tagname"
}
```

**Response:**
```json
{
  "post": {
    "id": "post-id",
    "title": "Post Title",
    "content": "Post content...",
    "author_id": "user-id",
    "author_name": "User Name",
    "created_at": 1625097600,
    "tags": ["tagname"]
  }
}
```

**Note:** Requires authentication.

#### Remove Tag from Post

```
DELETE /posts/:id/tags/:tag
```

**Response:**
```json
{
  "post": {
    "id": "post-id",
    "title": "Post Title",
    "content": "Post content...",
    "author_id": "user-id",
    "author_name": "User Name",
    "created_at": 1625097600,
    "tags": []
  }
}
```

**Note:** Requires authentication.

#### List Tags

```
GET /tags
```

**Query Parameters:**
- `post_id` (optional): Get tags for a specific post

**Response:**
```json
{
  "tags": ["tag1", "tag2", "tag3"]
}
```

#### Get Posts by Tag

```
GET /posts/by-tag/:tag
```

**Response:**
```json
{
  "posts": [
    {
      "id": "post-id",
      "title": "Post Title",
      "content": "Post content...",
      "author_id": "user-id",
      "author_name": "User Name",
      "created_at": 1625097600,
      "tags": ["tag1", "tag2"]
    }
  ],
  "total": 1,
  "tag": "tag1"
}
```

## Error Handling

All endpoints return appropriate HTTP status codes:

- `200 OK`: Successful request
- `201 Created`: Resource created successfully
- `400 Bad Request`: Invalid request parameters
- `401 Unauthorized`: Authentication required
- `404 Not Found`: Resource not found
- `500 Internal Server Error`: Server error

Error responses have the following format:

```json
{
  "error": "Error message"
}
```