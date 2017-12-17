A program to URL encode a string from the terminal. This is program is inspired by Eric Meyer's [URL Decoder/Encoder](https://meyerweb.com/eric/tools/dencoder/) but adapted to be used from the terminal rather than from a browser.

This program uses Go's native `net/url` package to do the encoding of the URL which provides a pretty smart encoding algorithm. For example, it won't encode the `http://` at the beginning of a URL but it would encode it if it is found as a query string parameter (e.g. `http://localhost?q=http://hello` would be encoded as `http://localhost?q=http%3A%2F%2Fhello`). Likewise, the Go library is smart enough to encode spaces as `+` if they are on the query string but as `%20` if they are anywhere else on the URL.

## Samples of usage
A very basic URL with a space on the query string:

```
$ ue "http://localhost/something?q=hello world"

# outputs
# http://localhost/something?q=hello+world
```

A more complex URL with multiple query string parameters that include spaces and quotes.

```
$ ue "http://localhost:8983/solr/bibdata/select?debugQuery=false&q=title:\"silver buckle\""

# outputs
# http://localhost:8983/solr/bibdata/select?debugQuery=false&q=title%3A%22silver+buckle%22
```

Passing the encoded result to cURL:

```
$ curl $(ue "http://localhost/something?q=hello world")
```

Passing the encoded result to cURL via `xargs`:

```
$ ue "http://localhost/something?q=hello world" | xargs curl
```

Notice that these examples assume the executable `ue` is in your PATH. If that is *not* the case for you use `./ue` instead.


## Compiling the source code
```
git clone https://github.com/hectorcorrea/urlenc
go build -o ue
./ue -help
```

## Downloading the executable
Download the executable for your operating system from the [Releases Tab](https://github.com/hectorcorrea/urlenc/releases).


## Final notes
There are many solutions to URL encoding on the terminal as described on [this Stack Overflow thread](https://stackoverflow.com/questions/296536/how-to-urlencode-data-for-curl-command). However, I didn't quite like that many of them encode the text in its entirely, for example, some of them encode `http://localhost` as `http%3A%2F%2Flocalhost` whereas this program is smart enough to only encode what needs to be encoded rather than blindly encoding each character. So, as xkcd would have guessed, now we have [yet another tool](https://xkcd.com/927/) for it.

An unfortunately side effect of using Go for this project is that the resulting executable is 1.3M in size, whereas cURL itself is only 181K. In Linux this is easy to address by compiling the executable with `gccgo` which results in a 33K executable, but I could not figure out how to use `gccgo` on the Mac so we are stuck with a 1.3M executable for now.
