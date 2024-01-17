# NoteMate-Server
NoteMate REST API 

This repo is a part of a larger NoteMate project. The following is a list of the relevant repos:
- [NoteMate-Server](https://github.com/ssebs/NoteMate-Server/)
  - This is the golang REST API server. Files will be saved here.
  - The main README will be here for now.
- [NoteMate-CLI](https://github.com/ssebs/NoteMate-CLI/)
  - This is the CLI for syncing a workspace / your notes to your computer.
- [NoteMate-Web](https://github.com/ssebs/NoteMate-Server/)
  - This is the web app + Electron desktop app to interact with a hosted NoteMate-Server.
- [NoteMate-Mobile](https://github.com/ssebs/NoteMate-Mobile)
  - This is the mobile app to interact with a hosted NoteMate-Server.


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
  - Web app
    - If flutter doesn't have a good MD WYSIWYG editor
    - JS lib for editing, material-ui for viewing what files you have, etc.
    - SSO?
    - WYSIWYG or MD Text (with helpers)
  - Electron app
    - Uses above web app, but allows for saving / loading files from disk.
- How it will work:
- POST /note/new with contents + author + metadata
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
