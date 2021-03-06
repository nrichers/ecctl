## ecctl platform deployment-template list

Lists the platform deployment templates

### Synopsis

Lists the platform deployment templates

```
ecctl platform deployment-template list [flags]
```

### Options

```
      --filter string                  Optional key/value pair in the form of key:value that will act as a filter and exclude any templates that do not have a matching metadata item associated
  -h, --help                           help for list
      --show-instance-configurations   Shows instance configurations - only visible when using the JSON output
      --stack-version string           If present, it will cause the returned deployment templates to be adapted to return only the elements allowed in that version.
```

### Options inherited from parent commands

```
      --apikey string      API key to use to authenticate (If empty will look for EC_APIKEY environment variable)
      --config string      Config name, used to have multiple configs in $HOME/.ecctl/<env> (default "config")
      --force              Do not ask for confirmation
      --format string      Formats the output using a Go template
      --host string        Base URL to use (default "https://api.elastic-cloud.com")
      --insecure           Skips all TLS validation
      --message string     A message to set on cluster operation
      --output string      Output format [text|json] (default "text")
      --pass string        Password to use to authenticate (If empty will look for EC_PASS environment variable)
      --pprof              Enables pprofing and saves the profile to pprof-20060102150405
  -q, --quiet              Suppresses the configuration file used for the run, if any
      --region string      Elastic Cloud region
      --timeout duration   Timeout to use on all HTTP calls (default 30s)
      --trace              Enables tracing saves the trace to trace-20060102150405
      --user string        Username to use to authenticate (If empty will look for EC_USER environment variable)
      --verbose            Enable verbose mode
```

### SEE ALSO

* [ecctl platform deployment-template](ecctl_platform_deployment-template.md)	 - Manages deployment templates

