# REST API for PadPal

This will explain what endpoints are available for PadPal-Server, and what they'll do.

## Notes
- `/notes`
  - **Required**:
    - JWT header to get userID
  - `POST`
    - Create a new note based on the user's JWT
    - **Optional**:
      - `content`
  - `GET`
    - List all notes owned by userID
- `/notes/<id>`
  - **Required**:
    - JWT header to get userID
  - `GET`
    - Get note contents
    - `?version=latest`
      - Get the latest version of a note (default)
    - `?version=<datestamp>`
      - Get specific version of a note
    - `?only=version`
      - List all versions of a note, don't include the contents of the note
  - `UPDATE`
    - **Required**:
      - `content`
    - `?version=latest`
      - Set the latest version of a note (default)
    - `?version=<datestamp>`
      - Restore specific version to latest
  - `DELETE`
    - Delete a note (all versions)
    - `?version=<datestamp>`
      - Delete a specific version

## Tags
- `/tags`
  - `GET`
    - List all tags + ids
  - `POST`
    - Create new tag
    - **Required**:
      - `tag_name`
- `/tags/<id>`
  - **Required**:
    - JWT header to get userID
  - `GET`
    - Get the tag 
  - `UPDATE`
    - Rename a tag
    - **Required**:
      - `tag_name`
  - `DELETE`
    - Delete a tag

## Users
- `/users`
  - `GET`
    - List all users
  - `POST`
    - Create new user + their workspace
    - **Required**:
      - `type`: either `SSO` or `PW`
      - `username`
      - `pass`: only if type is `PW`
- `/users/<id>`
  - **Required**:
    - JWT header to get userID
  - `GET`
    - Get user fields
  - `UPDATE`
    - Update user fields
  - `DELETE`
    - Delete a user + orphan their workspace
      - Can only delete own user, unless "admin" role?
        - Do we want RBAC for this?

## Auth
- `/auth/login`
  - `POST`
    - **Required**:
      - `username`
      - `passwd`
  - HTTP auth
  - Returns a JWT to use for other API calls
- `/auth/sso`
  - `POST`
    - **Required**:
      - TBD
  - 0auth w/ google
  - Returns a JWT to use for other API calls

