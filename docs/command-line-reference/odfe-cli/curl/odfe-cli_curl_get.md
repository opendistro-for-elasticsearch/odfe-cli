## odfe-cli curl get

Get command to execute requests against cluster

### Synopsis

Get command enables you to run any GET API against cluster

```
odfe-cli curl get [flags] 
```

### Examples

```

# get document count for an index
odfe-cli curl get --path "_cat/count/my-index-01" --query-params "v=true" --pretty

# get health status of a cluster.
odfe-cli curl get --path "_cluster/health" --pretty --filter-path "status"

# get explanation for cluster allocation for a given index and shard number
odfe-cli curl get --path "_cluster/allocation/explain" \
                  --data '{
                    "index": "my-index-01",
                    "shard": 0,
                    "primary": false,
                    "current_node": "nodeA"                         
                  }'

```

### Options

```
  -d, --data string           Data for the REST API. If value starts with '@', the rest should be a file name to read the data from.
  -H, --headers :             Headers for the REST API. Consists of case-insensitive name followed by a colon (:), then by its value. Use ';' to separate multiple parameters. Ex: -H "content-type:json;accept-encoding:gzip"
  -h, --help                  Help for curl get
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
