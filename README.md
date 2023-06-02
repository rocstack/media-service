# Media Service

This is going to my attempt of starting my own self-hosted media service. Initially it will focus on music.

## Aims

- The service will scan a target directory for music and store it's findings in a database in order to index
- Will use services such as musicbrainz or last.fm to fetch metadata and store this in the database
- The service will be able to keep the real-time state of currently playing music
- There will be a frontend to see the music library and play music

## Server

Built in GoLang for performance and safe multi-threading possibilities

## UI

A frontend application to see all indexed music data and ability to play/stream music from the media server
