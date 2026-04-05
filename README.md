# Status
A lightweight and minimalistic status page written in Go.

## Configuration
To run the service, you need to create a configuration file named `config.toml`

Example:
```toml
title = "title" # title in the page header (required)
description = "description" # description in the page header (required)
icon = "link of a icon" # icon in the page header (optional, default : https://placehold.co/600x400)
interval = 2 # checking interval, in seconds (optional, default : 2)
port = 3333 # listening port (optional, default : 3333)

[[service]]
name = "Forgejo" # name of the service (required)
url = "https://git.oriondev.fr" # url of the service, used in checking requests (required)
show_url = true # use hyperlink on the service's name (optional, default: false)
```

## Command Line Options
### Usage
`status [options]`

### Options
- `--port`/`-p` `<port>`: Listening port
- `--interval`/`-i` `<interval>`: Checking interval (in seconds)

## Installation
### Alpine Linux
**These commands need to be run as root.**

```sh
echo "https://git.oriondev.fr/api/packages/orion/alpine/stable/master" >> /etc/apk/repositories
cd /etc/apk/keys/ 
curl -JO https://git.oriondev.fr/api/packages/orion/alpine/key
apk add status
```

### Other
```sh
git clone https://git.oriondev.fr/orion/status.git
cd status
make build
cp bin/status /usr/bin
```

## License
This project uses [MIT](LICENSE) License.
