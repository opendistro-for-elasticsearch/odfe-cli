## odfe-cli curl post

Post command to execute requests against cluster

### Synopsis

Post command enables you to run any POST API against cluster

```
odfe-cli curl post [flags] 
```

### Examples

```

# change the allocation of shards in a cluster.
odfe-cli curl post --path "_cluster/reroute" \
                 --data '
                 {
                    "commands": [
                    {
                        "move": {
                            "index": "odfe-cli", "shard": 0,
                            "from_node": "odfe-node1", "to_node": "odfe-node2"
                        }
                    },
                    {
                        "allocate_replica": {
                            "index": "test", "shard": 1,
                            "node": "odfe-node3"
                        }
                    }
                ]}' \
				--pretty

# insert a document to an index 
odfe-cli curl post --path "my-index-01/_doc" \
                   --data '
                    {
                        "message": "insert document",
                        "ip": {
                            "address": "127.0.0.1"
                        }
                    }'


```

### Options

```
  -d, --data string           Data for the REST API. If value starts with '@', the rest should be a file name to read the data from.
  -H, --headers :             Headers for the REST API. Consists of case-insensitive name followed by a colon (:), then by its value. Use ';' to separate multiple parameters. Ex: -H "content-type:json;accept-encoding:gzip"
  -h, --help                  Help for curl post
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
