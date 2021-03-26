## odfe-cli ad create

Create detectors based on JSON files

### Synopsis

Create detectors based on a local JSON file
To begin, use `odfe-cli ad create --generate-template` to generate a sample configuration. Save this template locally and update it for your use case. Then use `odfe-cli ad create file-path` to create detector.

```
odfe-cli ad create json-file-path ... [flags]
```

### Options

```
  -g, --generate-template   Output sample detector configuration
  -h, --help                Help for create
```

### Options inherited from parent commands

```
  -c, --config string    Configuration file for odfe-cli, default is /Users/balasvij/.odfe-cli/config.yaml
  -p, --profile string   Use a specific profile from your configuration file
```

### SEE ALSO

* [odfe-cli ad](odfe-cli_ad.md)	 - Manage the Anomaly Detection plugin
