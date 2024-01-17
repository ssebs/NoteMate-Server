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
- 

### Wish
- 

## Architecture
- Server:
  - golang REST API to manage notes
  - dockerized
  - mounted volume, save latest file + diffs in folder
  - sqlite to keep track? Or flat file? Make this an interface
  - Basic version control (git under hood?)
  - Unit tests for every golang file, 75% coverage minimum
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
