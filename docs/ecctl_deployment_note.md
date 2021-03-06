## ecctl deployment note

Manages a deployment's notes

### Synopsis

Manages a deployment's notes

```
ecctl deployment note [flags]
```

### Options

```
  -h, --help   help for note
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

* [ecctl deployment](ecctl_deployment.md)	 - Manages deployments
* [ecctl deployment note create](ecctl_deployment_note_create.md)	 - Adds a note to a deployment
* [ecctl deployment note list](ecctl_deployment_note_list.md)	 - Lists the deployment notes
* [ecctl deployment note show](ecctl_deployment_note_show.md)	 - Shows a deployment note
* [ecctl deployment note update](ecctl_deployment_note_update.md)	 - Updates the deployment notes

