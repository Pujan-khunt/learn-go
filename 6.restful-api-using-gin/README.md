# Restful API using Gin

Data is going to be stored in-memory, meaning the data will be destroyed when server stops
and recreated when the server restarts.

A production grade server would typically use a Database for this job.

## Endpoints
1. /albums

    - GET: Gets the list of all albums, returned as JSON.
    - POST: Add a new album from the request data sent as JSON.

2. /albums/:id

    - Get an album by its ID, returning the album data as JSON.
