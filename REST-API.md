# REST API for PadPal

This will explain what endpoints are available for PadPal-Server, and what they'll do.

## Auth
- `/auth/login`
  - POST
  - HTTP auth
- `/auth/sso`
  - POST
  - 0auth w/ google

## Notes
- `/notes/new`
  - **Required**:
    - JWT header to get userID
  - POST
    - Create a new note based on the user's JWT
  - Optional:
    - Note content
- `/notes`
  - **Required**:
    - JWT header to get userID
  - `GET`
    - List all notes owned by userID
- `/notes/<id>`
  - **Required**:
    - JWT header to get userID
  - `GET`
    - Get note contents
    - `?version=<datestamp>`
      - Get specific version of a note
  - `UPDATE`
    - Set latest note contents
    - `?version=<datestamp>`
      - Restore specific version to latest
  - `DELETE`
    - Delete a note
    - `?version=<datestamp>`
      - Delete a specific version


## Users
- `/users/new`
  - **Required**:
    - `type`: either SSO or PW
    - `username`
    - `pass`: if type=PW
  - `POST`
    - Create new user + their workspace
- `/users/<id>`
  - `GET`
    - Get user fields
  - `UPDATE`
    - Update user fields
  - `DELETE`
    - Delete a user + orphan their workspace
      - Can only delete own user, unless "admin" role?
        - Do we want RBAC for this?
- `/users`
  - `GET`
    - List all users