[![Main](https://github.com/dhenkel92/rabbitmq-message-monitor/actions/workflows/main.yml/badge.svg?branch=main)](https://github.com/dhenkel92/rabbitmq-message-monitor/actions/workflows/main.yml)
[![Release](https://github.com/dhenkel92/rabbitmq-message-monitor/actions/workflows/release.yml/badge.svg)](https://github.com/dhenkel92/rabbitmq-message-monitor/actions/workflows/release.yml)

<!-- PROJECT LOGO -->
<br />
<p align="center">
  <h3 align="center">RabbitMQ Message Monitor</h3>

  <p align="center">
    Different ways of how to monitor messages that are going through your RabbitMQ infrastructure.
    <br />
    <br />
    <a href="https://github.com/dhenkel92/rabbitmq-message-monitor/issues">Report Bug</a>
    ·
    <a href="https://github.com/dhenkel92/rabbitmq-message-monitor/issues">Request Feature</a>
  </p>
</p>

![Product Name Screen Shot](docs/main-image.png)

## Table of Contents
<!-- vim-markdown-toc GFM -->

- [Installation](#installation)
  - [MacOS](#macos)
  - [Other distributions](#other-distributions)
- [How to use it](#how-to-use-it)
  - [Try it out](#try-it-out)
  - [General options](#general-options)
  - [Monitor an exchagne](#monitor-an-exchagne)
- [Development](#development)
  - [Pre-commit hooks](#pre-commit-hooks)
  - [Roadmap](#roadmap)

<!-- vim-markdown-toc -->

## Installation

### MacOS

You can install the tool via Homebrew and the tap repository can be found [here.](https://github.com/dhenkel92/homebrew-tap)
```
brew install dhenkel92/tap/rabbitmq-message-monitor
```

In order to get a newer version, just upgrade via Homebrew
```
brew upgrade dhenkel92/tap/rabbitmq-message-monitor
```

### Other distributions

See the [Releases page](https://github.com/dhenkel92/rabbitmq-message-monitor/releases) for a list of Peek packages for various distributions.


## How to use it

Example of how to run it:
```bash
rabbitmq-message-monitor -c "amqp://guest:guest@localhost:5672/" exchange --binding "amqp.topic=#"
```
### Try it out

1) Start a test RabbitMQ inside a Docker container and start multiple publisher:
```bash
make start-test
```
2) Open a second terminal
3) Start the RMQ Message Monitor UI
```bash
make run
```
4) Stop the publisher via CTRL-C
5) Cleanup everything
```bash
make clean
```

### General options

```bash
NAME:
   RabbitMQ Message Monitor - Different ways of how to monitor messages that are going through your RabbitMQ infrastructure.

USAGE:
   main [global options] command [command options] [arguments...]

COMMANDS:
   monitor-exchange, exchange  Bind to several exchange and visualize the different messages.
   help, h                     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --connection-string value, -c value
   --consumer-name value, --name value  (default: "rmq-message-monitor")
   --help, -h                           show help (default: false)
```

### Monitor an exchagne
```bash
NAME:
   main monitor-exchange - Bind to several exchange and visualize the different messages.

USAGE:
   main monitor-exchange [command options] [arguments...]

OPTIONS:
   --binding value, -b value  Which bindings should be created for monitoring. Format: <exchange>=<routing_key>
   --queue-prefix value       Name prefix to use for temporal queues. (default: "rmq-message-monitor")
   --help, -h                 show help (default: false)
```

## Development

### Pre-commit hooks
In order to keep a certain code quality, the project includes some pre-commit hooks that should be installed and executed on the developers pc.
In this case we are using [pre-commit](https://pre-commit.com/).

Install the hooks execute the following command in the root of the project:
```shell
$ pre-commit install
```

In case you want to execute all commands on all the files again you can just execute:
```shell
$ pre-commit run -a
```

### Roadmap

The following features will be implemented in the future
* Select a routing key via enter
  * Start to record messages with this specific routing key
  * Select a message and show the data in a prettified way
* Implement a `server` command
  * Should monitor the routing keys in the same way
  * But instead of showing a UI, it should export Prometheus metrics
