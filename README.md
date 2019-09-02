# consul-wrapper

This program allows to run and monitor a command and regisiter it in consul.
It also create TTL health check to make sure that the status of the service is accurate in consul.
If by any change the command stops working it will be deregistered automatically.

## Usage

```
Usage:
  consul-wrapper [flags]
  consul-wrapper [command]

Available Commands:
  help        Help about any command
  json        Create a consul service based on config file

Flags:
  -f, --frequency duration   Health Check Frequency (in seconds) (default 30s)
  -h, --help                 help for consul-wrapper
  -p, --port int             Service port
  -s, --service string       Consul Service Name
  -t, --token string         Consul token used for registration
```

## Example

```
 ./consul-wrapper --service plop --port 8080 --frequency 10s "consul monitor -log-level trace"
2019/04/11 20:16:44 [ConsulWrapper] Registering 'plop' in consul
...
2019/04/11 20:17:14 [ConsulWrapper] Process is running
...
2019/04/11 20:17:40 [ConsulWrapper] Error:  signal: killed
2019/04/11 20:17:40 [ConsulWrapper] Process stopped running. Exit code:  -1
2019/04/11 20:17:40 [ConsulWrapper] Deregistering 'plop' from consul
```
# Todo

* Service definition via json file
