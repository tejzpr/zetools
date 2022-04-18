[![Open Source](https://img.shields.io/badge/Open%20Source-%20-green?logo=open-source-initiative&logoColor=white&color=blue&labelColor=blue)](https://en.wikipedia.org/wiki/Open_source)
[![Golang](https://img.shields.io/badge/-Go%20Lang-blue?logo=go&logoColor=white)](https://golang.org)
[![Go Report Card](https://goreportcard.com/badge/github.com/tejzpr/zetools)](https://goreportcard.com/report/github.com/tejzpr/zetools)
[![CodeQL](https://github.com/tejzpr/zetools/actions/workflows/codeql-analysis.yml/badge.svg?branch=main)](https://github.com/tejzpr/zetools/actions/workflows/codeql-analysis.yml)

# ZETools
Command line tools for ze common tasks

# Using Docker
```docker
docker run -it ghcr.io/tejzpr/zetools:main zetools -h
```
# Using standalone CLI
Download a release, extract the archive and run
```sh
zetools -h
```

# Available Commands
## **base64**
### Usage 
```sh
zetools base64 <encode|decode> string
```
  decode [string]
  * Decode a Base64 encoded string.
        
  encode [string]
  * Base64 Encode a string.
---
## **hmac**
### Usage 
#### Hash a string
```sh
zetools hmac <sha256|sha512> -text <string> -key <your-key>
```
#### Hash a file
```sh
zetools hmac <sha256|sha512> -filename <string> -key <your-key>
```
  sha256 
  * Generate a SHA256 hash.
        
  sha512 
  * Generate a SHA512 hash.
---
## **checkPort**
### Usage
#### Check if a port is in use
```sh
zetools checkPort --port 8080 --host localhost
```
---
## **ping**
### Usage
#### Ping a remote server
```sh
zetools ping --host www.google.com
```
---
## **ls**
### Usage
#### Lists the contents of a directory
```sh
zetools ls -bs M -p ../
```
---
## **tail**
### Usage
#### Tails a file
```sh
zetools tail -f test.log -fw true
```

---
## This software wouldn't be possible without the following projects
1. [Tail](https://github.com/nxadm/tail)
2. [Ping](https://github.com/go-ping/ping)
3. [Tablewriter](https://github.com/olekukonko/tablewriter)
4. [CLI](https://github.com/urfave/cli/v2)