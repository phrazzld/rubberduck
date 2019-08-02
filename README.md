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

### rubberduck

Open today's entry:

`rubberduck`

### config

Set your editor:

`rubberduck config`

### review

Review recent entries (default opens entries from three months ago, one month ago, one week ago, and yesterday):

`rubberduck review`

### reminisce

Review old entries: (default opens entries from each year ago, and six months ago):

`rubberduck reminisce`

### search

Search your old entries for lines containing a term, phrase, or pattern:

`rubberduck search "<pattern>"`

Note: `search` ignores case and returns the lines immediately trailing and following the line with the matched pattern. It is equivalent to running:

`egrep "<pattern>" -R $ENTRIES_PATH --ignore-case -C 1`


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
