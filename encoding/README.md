# Example: encoding

Examples demonstrate how to use helper methods and third-party packages for JSON decoding and encoding.

## `tidwall/gjson`

The [`github.com/tidwall/gjson`](https://github.com/tidwall/gjson) package allows an
easy access to JSON properties, without parsing the payload into a data structure,
and is therefore convenient when accessing only selected parts of the response.

``` golang
var json = `{"foo":{"bar":"BAZ"}}`
fmt.Println(gjson.Get(json, "foo.bar"))
// => BAZ
```

The [**`gjson.go`**](./gjson.go) example displays data from the
[_Cluster Stats_](https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-stats.html) API.

```
go run gjson.go
```

## `mailru/easyjson`

The [`github.com/mailru/easyjson`](https://github.com/mailru/easyjson) package uses code generation to
provide fast encoding and decoding of struct types.

The [**`easyjson.go`**](./easyjson.go) example uses custom types for representing simple`Article` documents
and index and search responses, including error responses.

```
make clean setup
go run easyjson.go
```

## `esutil.JSONReader()`

The [`esutil.JSONReader()`](../../esutil/json_reader.go) helper method takes a struct, a map,
or any other serializable object, and converts it to JSON wrapped in a reader, for convenient
passing to the `WithBody()` methods:

```golang
type Document struct{ Title string }
doc := Document{Title: "Test"}
es.Search(es.Search.WithBody(esutil.NewJSONReader(&doc)))
```
