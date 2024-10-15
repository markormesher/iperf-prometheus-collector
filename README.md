![CircleCI](https://img.shields.io/circleci/build/github/markormesher/iperf-prometheus-collector)

# `iperf` Prometheus Collector

A simple Prometheus collector to provide measurements about network connection throughput for multiple hosts using the [iperf3](https://iperf.fr/) utility.

:rocket: Jump to [quick-start example](#quick-start-docker-compose-example).

:whale: See releases on [ghcr.io](https://ghcr.io/markormesher/iperf-prometheus-collector).

Note that `iperf` tests take 10+ seconds per target and are executed sequentially, so this collector runs asynchronously. Tests are run on a configurable interval and every request to the `/metrics` endpoint will return the most recent results. Emitted metrics are timestamped, so this approach does not result in out of date data being logged.

## Measurements

| Measurement              | Description                                             | Labels   |
| ------------------------ | ------------------------------------------------------- | -------- |
| `iperf_sent_bytes`       | Total number of bytes sent during the test.             | `target` |
| `iperf_sent_seconds`     | Duration of the test on the sending side, in seconds.   | `target` |
| `iperf_received_bytes`   | Total number of bytes received during the test.         | `target` |
| `iperf_received_seconds` | Duration of the test on the receiving side, in seconds. | `target` |

These metrics can be combined to show the throughput in bps with the following example Prometheus query:

```
avg by (target) (iperf_received_bytes / iperf_received_seconds * 8)
```

## Configuration

Configuration is via the following environment variables:

| Variable           | Required? | Description                                                                                                    | Default                 |
|--------------------|-----------|----------------------------------------------------------------------------------------------------------------|-------------------------|
| `TARGET_LIST`      | yes       | Comma separated list of host names or IP addresses to run tests against.                                       | n/a                     |
| `TEST_INTERVAL_MS` | no        | How often to run iperf tests.                                                                                  | 600000ms (= 10 minutes) |
| `OPTIONS`          | no        | Additional [iperf options](https://github.com/esnet/iperf/blob/master/docs/invoking.rst) (e.g. `--bitrate 1k`) | none                    |

### `iperf` Server

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
