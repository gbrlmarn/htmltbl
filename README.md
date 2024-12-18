# htmltbl
A html table parser for the terminal.

## Installation
### From source
With go install(Go 1.18 or higher required):

    go install github.com/gbrlmarn/htmltbl@latest

Or building from repository:

    git clone http://github.com/gbrlmarn/htmltbl
    cd htmltbl
    go build

## Usage
You must specify a http link and optionally the format(table, json)

    htmltbl --format table <url>
    htmltbl --format json <url>
    htmltbl --format json-indent <url>

## License
[MIT](https://github.com/gblmrn/htmltbl/raw/main/LICENSE)
