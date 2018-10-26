# rubberduck

Take and review quick notes from the command line using your favorite editor.


## Install

[Set up your Go environment](https://golang.org/doc/install). Make sure you have the following lines in your **.profile**:

```
export GOPATH=`go env GOPATH`
export PATH="$PATH:$GOPATH/bin"
```

Then run:

`go get github.com/phrazzld/rubberduck`


## Usage

Set your editor (i.e., the command that runs to open your note file):

`rubberduck config`

Take some notes:

`rubberduck`

Review past notes (default opens notes from each year ago, six months ago, three months ago, one month ago, one week ago, and yesterday):

`rubberduck review`


## LICENSE
[MIT](https://opensource.org/licenses/MIT)
