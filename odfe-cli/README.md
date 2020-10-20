## ODFE Command Line Interface

ODFE Command Line Interface (odfe-cli) is an open source tool that enables customers to manage their ODFE  clusters 
using CLI or command-line-shell, additionally CLI will allow customers  to configure and manage ODFE plugins 
like **Anomaly Detections**, **Alerting**, **SQL** and access **Elasticsearch** features from 
the command line with prominence on automation.

With minimal configuration, odfe-cli enables users to execute commands runs from Kibana as well as
additional ad hoc commands used to manage an Elasticsearch cluster.

odfe-cli is better suited when you want to quickly combine a couple of commands to perform a task. These commands are 
powerful and concise. ecause shell commands can be strung together in scripts, a user can easily build a complex
sequence of commands to automate frequently run operations.

## To start developing odfe-cli

### Minimum Requirements

odfe-cli shares the same [minimum requirements](https://github.com/golang/go/wiki/MinimumRequirements#minimum-requirements)
  as Go:
- Linux kernel version 2.6.23 or later
- Windows 7 or later
- FreeBSD 11.2 or later
- MacOS 10.11 El Capitan or later

### Installation:

You can download the binaries directly from the [downloads](https://opendistro.github.io/for-elasticsearch/downloads.html) page
or from the [releases](https://github.com/opendistro-for-elasticsearch/es-cli/releases) section.

#### From Source:
odfe-cli requires Go version 1.14 or newer
1. Install [Go](https://golang.org/doc/install) > = 1.14
2. Clone the repository
    ```
    cd $GOPATH/src
    git clone git@github.com:opendistro-for-elasticsearch/es-cli.git
    ```
3. Run build from source directory to generate binary
   ```
   cd es-cli/odfe-cli
   go build .
   ```
4. Make binary executable
    ```
    chmod +x ./odfe-cli
    ```

## How to use it:

See usage with:

```
odfe-cli --help
```

#### Create default profile

```
$ odfe-cli profile create
Enter profile's name: default
Elasticsearch Endpoint: https://localhost:9200  
User Name: admin
Password: 
```

#### List existing profile

```
$ odfe-cli profile list -l
Name         UserName            Endpoint-url             
----         --------            ------------              
default      admin               https://localhost:9200   
prod         admin               https://odfe-node1:9200
                 
```

#### Using profile with odfe-cli command

1. As Parameter to ODFE Commands

     To use a named profile, add the --profile profile-name option to your command.

    **Example**
    
    The following example stops detector “invalid-logins” using the credentials and settings defined in the prod profile.
    ```
    $ odfe ad stop-detector invalid-logins --profile prod
    ```
    
2. As An Environment Variable

    To use a named profile for multiple commands in same session, you can avoid specifying the profile in every command by setting the ODFE_PROFILE environment variable at the command line. This value will be available until the end of your shell session, or until you set the variable to a different value.

    **Example**

    The following example will set prod as default profile in current session

    Linux or macOS
    ```
    $ export ODFE_PROFILE=prod
    ```
    
    Windows
    ```
    C:\> setx ODFE_PROFILE prod
    ```
    
## Security

See [CONTRIBUTING](https://github.com/opendistro-for-elasticsearch/es-cli/blob/main/CONTRIBUTING.md#security-issue-notifications) for more information.

## License

This project is licensed under the Apache-2.0 License.

