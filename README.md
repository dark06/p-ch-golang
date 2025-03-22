# PCH Golang

## Development Environment

- macOS

## Pre-requisites

- [Golang](https://go.dev/doc/install) `v3.12.7`
- [Task](https://taskfile.dev/) `v3.42.1`

## Solution

- The container is configured with `NetworkMode: "host"`, allowing the Go client to establish a connection using the server’s private IP address `172.31.28.67`.

- The `index.html` file was not being served correctly due to insufficient file permissions. This issue was resolved by executing `chmod 644 /var/www/html/index.html` as the root user
