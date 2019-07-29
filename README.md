# rubberduck

Ad-hoc journaling at the command line.


## Install

[Set up your Go environment](https://golang.org/doc/install). Make sure you have the following lines in your **.profile**:

```
export GOPATH=`go env GOPATH`
export PATH="$PATH:$GOPATH/bin"
```

Then run:

`go get github.com/phrazzld/rubberduck`


## Usage

Set your editor:

`rubberduck config`

Open today's entry:

`rubberduck`

Review recent entries: (default opens entries from three months ago, one month ago, one week ago, and yesterday):

`rubberduck review`

Review old entries: (default opens entries from each year ago, and six months ago):

`rubberduck reminisce`


## Testing

```
go test
go test -bench .
go test -cover
go test -coverprofile c.out
go tool -cover -html=c.out
```

## License
[MIT](https://opensource.org/licenses/MIT)
