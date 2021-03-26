## odfe-cli ad update

Update detectors based on JSON files

### Synopsis

Update detectors based on JSON files.
To begin, use `odfe-cli ad get detector-name > detector_to_be_updated.json` to download the detector. Modify the file, and then use `odfe-cli ad update file-path` to update the detector.

```
odfe-cli ad update json-file-path ... [flags]
```

### Options

```
  -f, --force   Stop detector and update forcefully
  -h, --help    Help for update
  -s, --start   Start detector if update is successful
```

### Options inherited from parent commands

```
  -c, --config string    Configuration file for odfe-cli, default is /Users/balasvij/.odfe-cli/config.yaml
  -p, --profile string   Use a specific profile from your configuration file
```

### SEE ALSO

* [odfe-cli ad](odfe-cli_ad.md)	 - Manage the Anomaly Detection plugin
