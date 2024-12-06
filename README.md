# Multi-Audio System

Project Overview:

**Multi-Audio system: system that handle multiple audio tracks and play them at the same time and store them as session for future use.**


Main Features:

- Play multiple track ∞ . Go routines
- Change the volume of each track - settings per track.
- CRUD session w/ the track + the settings of each of them. SQLite/PostgreSQL
- MP3/WAV convertor/resampler - 44.1HZ | 48HZ | 88HZ | 192HZ
- Downloader youtube → MP3/WAV
- GUI → Desktop app [cross-platform] || website || mobile app [cross-platform]
- profile | Profiler → for the favorite audio file + format + streamed or buffered → in SQLite
- Store the tracks on the cloud → download them || stream them from it. According to the profile.
- Searcher → Search on a cloud file storage, for music and locally.
- Uploader/Downloader → upload track to the storage solution | download track from the storage solution.
- Connector → just to make sure that it’s 200 between the storage/DB, and can retrieve anything their.
- caching or buffering from the cloud locally.
- Error handling of anything and everything | missing files.