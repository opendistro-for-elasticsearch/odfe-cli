## odfe-cli curl

Manage Elasticsearch core features

### Synopsis

Use the curl command to execute any REST API calls against the cluster.

### Options

```
  -f, --filter-path string     Filter output fields returned by Elasticsearch. Use comma ',' to separate list of filters
  -h, --help                   Help for curl command
  -o, --output-format string   Output format if supported by cluster, else, default format by Elasticsearch. Example json, yaml
      --pretty                 Response will be formatted
```

### Options inherited from parent commands

```
  -c, --config string    Configuration file for odfe-cli, default is /Users/balasvij/.odfe-cli/config.yaml
  -p, --profile string   Use a specific profile from your configuration file
```

### SEE ALSO

* [odfe-cli](../odfe-cli.md)	 - odfe-cli is a unified command line interface for managing ODFE clusters
* [odfe-cli curl delete](odfe-cli_curl_delete.md)	 - Delete command to execute requests against cluster
* [odfe-cli curl get](odfe-cli_curl_get.md)	 - Get command to execute requests against cluster
* [odfe-cli curl post](odfe-cli_curl_post.md)	 - Post command to execute requests against cluster
* [odfe-cli curl put](odfe-cli_curl_put.md)	 - Put command to execute requests against cluster
