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
## base64
### Usage 
```sh
zetools base64 <encode|decode> string
```
  decode [string]
  * Decode a Base64 encoded string.
        
  encode [string]
  * Base64 Encode a string.