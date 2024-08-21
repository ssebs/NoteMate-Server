# REST API for PadPal

This will explain what endpoints are available for PadPal-Server, and what they'll do.

- **POST** `/notes`
  - Create note
- **PUT** `/notes/:id`
  - Update note
- **GET** `/notes/:id`
  - Download note
  - Defaults to download note
  - `params: (?key=val)`
    - `meta`: (true|FALSE)
      - Get metadata only (do not download file)
- **GET** `/notes`
  - List note metadata
  - Defaults to all notes
  - `params: (?key=val)`
    - `q`: (filter query)
      - Filter results, supports * wildcard
- **DELETE** `/notes/:id`
  - Delete note

TBD:
`/auth`