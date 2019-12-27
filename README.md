# WatcherUpload
This repo is an experimentation of watching a folder and dynamically uploading files to a server.

The project is divided in two parts: the `server` and the `watcher`

## Watcher
The watcher is looking to any changes in a specific directory.
It interacts with the [server](#Server) to synchronise the specified directory with the server.

The files actions supported are:
* Created
* Written
* Moved
* Renamed
* Removed

The watcher takes the action needed to synchronise the local directory and the server directory.

macOS `.DS_Store` are ignored.

The file permissions are not synced.

## Server
The server is a simple REST API receiving diverse information from the watcher about the status of a folder.

File actions supported by the server:
* Upload
* Remove
* Move
* Rename

It will create folders if needed to support nested directory.

An empty directory can't be upload. A file must exist in the directory.
