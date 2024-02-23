# htmltbl

A html table parser for the terminal.

## Instalation
### From source
With go install(Go 1.18 or higher required):

    go install github.com/marinmgabriel/htmltbl@latest

Or building from repository:

    git clone http://github.com/marinmgabriel/htmltbl
    cd htmltbl
    go build

## Usage
You must specify a http link and optionally the format(table, json)

    htmltbl --format table <url>
    htmltbl --format json <url>

## License
[MIT](https://github.com/marinmgabriel/htmltbl/raw/main/LICENSE)
