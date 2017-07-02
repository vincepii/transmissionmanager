# transmissionmanager

A simple application that polls transmission web server to:
* remove finished torrents (removes the torrent, not the files)
* move downloaded data to a given destination folder

## Build

Build a statically linked executable:

    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' transmissionmanager.go

## Usage

    usage: ./transmissionmanager uri username password source dest
    	username: transmission web uri (e.g., http://transmission.local:8080)
    	username: transmission web username
    	password: transmission web password
    	source: path to transmission downloads
    	dest: where to move downloaded files after torrent is removed

## Docker

After building the executable, a docker image can be built with the provided Dockerfile.

A basic ansible playbook is provided to build the container and push it to a docker registry.

### Example (docker)

When using with docker, host volumes can be mapped to the container as in the following example:

    docker run -v /transmission/Downloads/:/source -v /myserver/data:/dest" -d <docker_repo>/trmanager

Then, the application can be started using the bind mounts on the container:

    ./transmissionmanager <uri> <username> <password> /source /dest
