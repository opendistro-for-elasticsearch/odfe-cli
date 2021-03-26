## odfe-cli curl delete

Delete command to execute requests against cluster

### Synopsis

Delete command enables you to run any DELETE API against cluster

```
odfe-cli curl delete [flags] 
```

### Examples

```

# Delete a document from an index. 
odfe-cli curl delete --path         "my-index/_doc/1" \
                     --query-params "routing=odfe-node1"

```

### Options

```
  -H, --headers :             Headers for the REST API. Consists of case-insensitive name followed by a colon (:), then by its value. Use ';' to separate multiple parameters. Ex: -H "content-type:json;accept-encoding:gzip"
  -h, --help                  Help for curl delete
  -P, --path string           URL path for the REST API
  -q, --query-params string   URL query parameters (key & value) for the REST API. Use ‘&’ to separate multiple parameters. Ex: -q "v=true&s=order:desc,index_patterns"
```

### Options inherited from parent commands

```
  -c, --config string          Configuration file for odfe-cli, default is /Users/balasvij/.odfe-cli/config.yaml
  -f, --filter-path string     Filter output fields returned by Elasticsearch. Use comma ',' to separate list of filters
  -o, --output-format string   Output format if supported by cluster, else, default format by Elasticsearch. Example json, yaml
      --pretty                 Response will be formatted
  -p, --profile string         Use a specific profile from your configuration file
```

### SEE ALSO

* [odfe-cli curl](odfe-cli_curl.md)	 - Manage Elasticsearch core features
