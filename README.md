# PadPal-Server
PadPal REST API 

This repo is a part of a larger PadPal project. The following is a list of the relevant repos:
- [PadPal-Server](https://github.com/ssebs/PadPal-Server/)
  - This is the golang REST API server. Files will be saved here.
  - The main README will be here for now.
- [PadPal-CLI](https://github.com/ssebs/PadPal-CLI/)
  - This is the CLI for syncing a workspace / your notes to your computer.
- [PadPal-Mobile](https://github.com/ssebs/PadPal-Mobile)
  - This is the mobile app to interact with a hosted PadPal-Server.

## Usage
> TODO: replace this section once the project is up and running

## Progress
> keeping track of where I am / whats next to code / PM stuff

- [x] Create REST server
- [ ] MOVE TO /api
  - [x] GET /
    - [x] Server REST API docs
      - [x] As rendered MD
  - [x] 404 / 500 error support
    - [x] Don't crash the server on 500
- [x] Deployable via Docker
- Design Data Models
  - [x] Notes
  - [ ] Tags
  - [ ] Users + Auth
- [x] Create CRUDProviders interface to support:
    - [x] Create
    - [x] Read
      - [x] List all notes
      - [x] Details for 1 note
      - [x] List versions for 1 note
      - [x] Details for version # of note
    - [x] Update
      - [x] Update to...
      - [x] Restore from version
    - [x] Delete
- [ ] Save / Load data files
  - [ ] Test/Sample Provider
    - [x] Create
    - [x] Read
    - [ ] Update
    - [ ] Delete
  - [ ] File Provider
- [ ] Implement Version history
  - [ ] git?
  - [ ] db?
- [ ] ...
- [ ] Auth
- [ ] ...

### Docker
> TODO: Update docs once volume mount is ready!

- Install [Docker](https://www.docker.com/get-started/)/[Podman](https://podman.io/docs/installation)
- Running DockerHub image [ssebs/padpal-server](https://hub.docker.com/r/ssebs/padpal-server):
  - `docker run -d -p 5000:5000 --rm ssebs/padpal-server`
- Local / build:
  - Building:
    - `git clone github.com/ssebs/padpal-server`
    - `cd padpal-server`
    - `docker build -t ssebs/padpal-server .`
  - Running:
    - `docker run -d -p 5000:5000 --rm ssebs/padpal-server`
- Unraid:
  - TODO

### Run from src
- Install [go > v1.21](https://go.dev/doc/install)
- `go get github.com/ssebs/padpal-server`
- `go run cmd/main.go`

## Feature list for PadPal-Server
### MVP
- Save notes as files in folder structure
- SaveProvider (allow to save files in diff places)
  - Local path (container volume mount)
  - Google Drive folder?
  - S3 bucket?
  - Under the hood, use interfaces to accomplish this
- REST API
  - create goroutine on each accept
    - So if there's a 500, it will not crash everything
  - support various errors, 404, 401, 500
  - JWT
  - See [REST API Doc](./REST-API.md)
- Tag support
- Keep versions of each note
  - Allow for rollback
  - Git?
- Host Web client on the server


## Architecture
- Server:
  - golang REST API to manage notes
  - dockerized
  - mounted volume, save latest file + diffs in folder
  - sqlite to keep track? Or flat file? Make this an interface
  - Basic version control (git under hood?)
  - Unit tests for every golang file, 75% coverage minimum
- Web client:
  - use tailwindcss
  - react? maybe preact
- CLI:
  - golang CLI app
  - ./cli -login 
    - Opens a browser & SSO happens
    - Or, no auth to start
  - Sync a directory (workspace?)
    - Merge conflict? Use the server version & save local as .fix-me for now
  - For use with VSCode/vim/text editor of your choosing (MD Text only)
- UI:
  - Flutter mobile app
    - Similar to google keep, if they have a MD text editor lib
    - Android home screen widget to view / open / create new notes like keeps'
    - WYSIWYG or MD Text (with helpers)
    - Exports for Android + Web + iOS?
    - If flutter doesn't have a good MD WYSIWYG editor
    - JS lib for editing, material-ui for viewing what files you have, etc.
    - SSO?
    - WYSIWYG or MD Text (with helpers)
    - Save / load files from disk
- How it will work:
- POST /notes with contents + author + metadata
- Server will save file, record version with now() in DB
- 201 response
- UI's will sync periodically, using same logic as CLI sync
  - GET /note/latest?version
    - Only get the version information! 
    - If it's the latest, do nothing
    - If not, GET /note/latest
      - Get contents + metadata
  - Pushing local changes, PUT /note/<id>
    - Server will add new version to VC, reply 201
- Wish:
  - Google SSO login?
  - Merge conflict UI
  - Collaboration

## LICENSE
[Apache License 2.0](./LICENSE)

## External references / docs
- https://pkg.go.dev/github.com/golang-jwt/jwt#example-package-GetTokenViaHTTP
- UNRAID https://selfhosters.net/docker/templating/templating/#114-shave-off-the-xml
