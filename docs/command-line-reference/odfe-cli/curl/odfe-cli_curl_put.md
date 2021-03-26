## odfe-cli curl put

Put command to execute requests against cluster

### Synopsis

Put command enables you to run any PUT API against cluster

```
odfe-cli curl put [flags] 
```

### Examples

```

# Create a knn index from mapping setting saved in file "knn-mapping.json"
odfe-cli curl put --path "my-knn-index"  \
                  --data "@some-location/knn-mapping.json" \
                  --pretty

# Update cluster settings transiently
odfe-cli curl put --path "_cluster/settings" \
                  --query-params "flat_settings=true"  \
                  --data '
                    {
                      "transient" : {
                        "indices.recovery.max_bytes_per_sec" : "20mb"
                      }
                    }' \
                  --pretty


```

### Options

```
  -d, --data string           Data for the REST API. If value starts with '@', the rest should be a file name to read the data from.
  -H, --headers :             Headers for the REST API. Consists of case-insensitive name followed by a colon (:), then by its value. Use ';' to separate multiple parameters. Ex: -H "content-type:json;accept-encoding:gzip"
  -h, --help                  Help for curl put
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
