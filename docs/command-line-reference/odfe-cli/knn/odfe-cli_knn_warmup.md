## odfe-cli knn warmup

Warmup shards for given indices

### Synopsis

Warmup command loads all graphs for all of the shards (primaries and replicas) for given indices into native memory.
This is an asynchronous operation. If the command times out, the operation will still be going on in the cluster.
To monitor this, use the Elasticsearch _tasks API. Use `odfe-cli knn stats` command to verify whether indices are successfully loaded into memory.

```
odfe-cli knn warmup index ... [flags] 
```

### Options

```
  -h, --help   Help for k-NN plugin warmup command
```

### Options inherited from parent commands

```
  -c, --config string    Configuration file for odfe-cli, default is /Users/balasvij/.odfe-cli/config.yaml
  -p, --profile string   Use a specific profile from your configuration file
```

### SEE ALSO

* [odfe-cli knn](odfe-cli_knn.md)	 - Manage the k-NN plugin
