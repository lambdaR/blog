# Static Web UI

The Micro Blog project includes a minimalist static web interface that interacts with the REST API provided by the Web Service.

## Overview

The static web UI is located in the `web/static/` directory and consists of:

- `index.html`: Main feed, create posts, view posts and comments
- `login.html`: User login page
- `signup.html`: User registration page
- `profile.html`: User profile, posts, and comments
- `main.js`: JavaScript for all pages
- `style.css`: Styling for all pages

## Accessing the UI

When you run the Web Service (`web/main.go`), it serves both the REST API and the static web UI. You can access the UI at:

```
http://localhost:42096
```

## Pages

### Main Feed (index.html)

The main feed page displays:

- A form to create new posts (if logged in)
- A list of all posts with their titles, content, and author
- Comments for each post
- A form to add comments to posts (if logged in)
- Tags for each post
- A form to add tags to posts (if logged in)
- Links to filter posts by tag

### Login Page (login.html)

The login page allows users to:

- Enter their email and password
- Submit the login form
- Navigate to the signup page if they don't have an account

### Signup Page (signup.html)

The signup page allows users to:

- Enter their name, email, and password
- Submit the signup form
- Navigate to the login page if they already have an account

### Profile Page (profile.html)

The profile page displays:

- User information
- Posts created by the user
- Comments made by the user

## JavaScript Implementation

The `main.js` file handles:

- API calls to the REST endpoints
- Form submissions
- Dynamic content loading
- Authentication state management
- Error handling

Example of API interaction:

```javascript
// Fetch posts
async function fetchPosts() {
  try {
    const response = await fetch('/posts');
    const data = await response.json();
    return data.posts || [];
  } catch (error) {
    console.error('Error fetching posts:', error);
    return [];
  }
}

// Create a post
async function createPost(title, content) {
  try {
    const response = await fetch('/posts', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ title, content }),
    });
    const data = await response.json();
    return data.post;
  } catch (error) {
    console.error('Error creating post:', error);
    return null;
  }
}
```

## CSS Styling

The `style.css` file provides:

- Basic responsive layout
- Styling for forms, buttons, and inputs
- Card-based design for posts and comments
- Navigation styling

## Authentication Flow

The UI handles authentication through:

1. Login/signup forms that submit to the appropriate API endpoints
2. Session cookies managed by the server
3. UI state changes based on authentication status
4. Protected actions (create post, add comment) that require authentication

## Tag Management

The UI provides tag functionality:

- Displaying tags for each post
- Adding tags to posts
- Removing tags from posts
- Filtering the feed by tag
- Browsing all available tags

## Responsive Design

The UI is designed to work on:

- Desktop browsers
- Tablet devices
- Mobile phones

## Integration with the API

The static UI interacts with the REST API provided by the Web Service:

1. JavaScript makes fetch requests to the API endpoints
2. The API returns JSON responses
3. JavaScript updates the UI based on the responses
4. Forms submit data to the API endpoints

## Customization

To customize the UI:

1. Modify the HTML files in `web/static/`
2. Update the CSS in `style.css`
3. Extend the JavaScript functionality in `main.js`