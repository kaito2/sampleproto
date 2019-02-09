## Generate Go Code

```bash
$ rm -r go-protos

# protoc
$ for file in `find protos -name '*.proto'`; do  protoc --go_out=plugins=grpc:. $file; done

$ mv github.com/kaito2/sampleproto/go-protos/ .
$ rm -r github.com/
```