## odfe-cli profile

Manage a collection of settings and credentials that you can apply to an odfe-cli command

### Synopsis

A named profile is a collection of settings and credentials that you can apply to an odfe-cli command. When you specify a profile for a command (e.g. `odfe-cli <command> --profile <profile_name>`), odfe-cli uses the profile's settings and credentials to run the given command.
To configure a default profile for commands, either specify the default profile name in an environment variable (`ODFE_PROFILE`) or create a profile named `default`.

### Options

```
  -h, --help   Help for profile
```

### Options inherited from parent commands

```
  -c, --config string    Configuration file for odfe-cli, default is /Users/balasvij/.odfe-cli/config.yaml
  -p, --profile string   Use a specific profile from your configuration file
```

### SEE ALSO

* [odfe-cli](../odfe-cli.md)	 - odfe-cli is a unified command line interface for managing ODFE clusters
* [odfe-cli profile create](odfe-cli_profile_create.md)	 - Create profile
* [odfe-cli profile delete](odfe-cli_profile_delete.md)	 - Delete profiles by names
* [odfe-cli profile list](odfe-cli_profile_list.md)	 - List profiles from the config file
