## odfe-cli ad start

Start detectors based on a list of IDs, names, or name regex patterns

### Synopsis

Start detectors based on a list of IDs, names, or name regex patterns.
Wrap regex patterns in quotation marks to prevent the terminal from matching patterns against the files in the current directory.
The default input is detector name. Use the `--id` flag if input is detector ID instead of name

```
odfe-cli ad start detector_name ... [flags] 
```

### Options

```
  -h, --help   Help for start
      --id     Input is detector ID
```

### Options inherited from parent commands

```
  -c, --config string    Configuration file for odfe-cli, default is /Users/balasvij/.odfe-cli/config.yaml
  -p, --profile string   Use a specific profile from your configuration file
```

### SEE ALSO

* [odfe-cli ad](odfe-cli_ad.md)	 - Manage the Anomaly Detection plugin
