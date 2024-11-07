![CircleCI](https://img.shields.io/circleci/build/github/markormesher/iperf-prometheus-collector)

# `iperf` Prometheus Collector

A simple Prometheus collector to provide measurements about network connection throughput for multiple hosts using the [iperf3](https://iperf.fr/) utility.

:rocket: Jump to [quick-start example](#quick-start-docker-compose-example).

:whale: See releases on [ghcr.io](https://ghcr.io/markormesher/iperf-prometheus-collector).

Note that `iperf` tests take 10+ seconds per target and are executed sequentially. Results from each test will be returned on the next call to `/metrics`.

## Measurements

See [`cmd/metric.go`](./cmd/metric.go).

## Configuration

Configuration is via the following environment variables:

| Variable             | Required? | Description                                                              | Default                 |
|----------------------|-----------|--------------------------------------------------------------------------|-------------------------|
| `TARGET_LIST`        | yes       | Comma separated list of host names or IP addresses to run tests against. | n/a                     |
| `UPDATE_INTERVAL_MS` | no        | How often to run iperf tests.                                            | 600000ms (= 10 minutes) |
| `TEST_PROTOCOL`      | no        | Test protocol, `tcp` or `udp`.                                           | `tcp`                   |
| `TEST_OPTIONS`       | no        | Options passed directly to `iperf3`, e.g. `-p 5202 -t 20`.               | none                    |
| `LISTEN_PORT`        | no        | Server port to listen on.                                                | 9030                    |

### Important Note: `TEST_OPTIONS`

- Do not use these options to set the `--udp` flag for tests - use the `TEST_PROTOCOL=udp` environment variable instead.
- Do not use these optoins to change the output format - this tool already sets `--json` and expects a JSON output.
- These options are passed directly to the command line used to run the test. If anyone else can control this argument they can easily run arbitrary commands, so make sure untrusted users are not able to make changes.

## `iperf` Server

This collector requires that an `iperf` server is running on each of the targets to be tested and is available on the default port.

You can start the tool in server mode with `$ iperf3 -s`, but for long-term use you'll want the server to be available all the time. The easiest way to do this is enable an `iperf` background daemon so that it is started at system boot - information on how to do this can be found in the `iperf` man pages, or [here](https://askubuntu.com/questions/1251443/start-iperdf3-deamon-at-startup) for Debian-based/`systemd` systems.

## Quick-Start Docker-Compose Example

```yaml
version: "3.8"

services:
  iperf-prometheus-collector:
    image: ghcr.io/markormesher/iperf-prometheus-collector:VERSION
    restart: unless-stopped
    environment:
      - TARGET_LIST=my-host-01,my-host-02,12.34.56.78
    ports:
      - 9030:9030
```
