## odfe-cli profile create

Create profile

### Synopsis

Create named profile to save settings and credentials that you can apply to an odfe-cli command.

```
odfe-cli profile create [flags]
```

### Options

```
  -a, --auth-type string   Authentication type. Options are disabled, basic and aws-iam.
                           If security is disabled, provide --auth-type='disabled'.
                           If security uses HTTP basic authentication, provide --auth-type='basic'.
                           If security uses AWS IAM ARNs as users, provide --auth-type='aws-iam'.
                           odfe-cli asks for additional information based on your choice of authentication type.
  -e, --endpoint string    Create profile with this endpoint or host
  -h, --help               Help for create
  -m, --max-retry int      Maximum retry attempts allowed if transient problems occur.
                           You can override this value by using the ODFE_MAX_RETRY environment variable. (default 3)
  -n, --name string        Create profile with this name
  -t, --timeout int        Maximum time allowed for connection in seconds.
                           You can override this value by using the ODFE_TIMEOUT environment variable. (default 10)
```

### Options inherited from parent commands

```
  -c, --config string    Configuration file for odfe-cli, default is /Users/balasvij/.odfe-cli/config.yaml
  -p, --profile string   Use a specific profile from your configuration file
```

### SEE ALSO

* [odfe-cli profile](odfe-cli_profile.md)	 - Manage a collection of settings and credentials that you can apply to an odfe-cli command
