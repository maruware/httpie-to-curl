# httpie-to-curl

## Usage

```
$ httpie-to-curl http post https://example.com foo=bar
curl -X POST -d {"foo":"bar"} https://example.com
```
