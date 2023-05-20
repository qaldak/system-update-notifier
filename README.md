# system-update-notifier

[![Go Release Builder](https://github.com/qaldak/system-update-notifier/actions/workflows/go-release-builder.yml/badge.svg)](https://github.com/qaldak/system-update-notifier/actions/workflows/go-release-builder.yml)

[![Go Builder](https://github.com/qaldak/system-update-notifier/actions/workflows/go-builder.yml/badge.svg?branch=main)](https://github.com/qaldak/system-update-notifier/actions/workflows/go-builder.yml)

## Description

Notify to Slack channel if system upgrades on Debian-based Linux are available.<br>
On DietPi: <br>

- for Dietpi upgrades, check /run/dietpi/.updates_available<br>
- for apt upgrades, check /run/dietpi/.apt_updates<br>

On Debian-based Linux checking updates by Apt package manager.<br>
<br>

## Getting started

### Download binary

Sysup-notifier binaries are available for **amd64, arm7 or arm64**. Ready to download from release. <br>
example: <br>

```
wget -c https://github.com/qaldak/system-update-notifier/releases/download/1.0.0/system-update-notifier_1.0.0_linux_amd64.tar.gz
```

### Configuration

#### **.env File**

Following variables has to be defined in `.env` file for Slack notifications:<br>

```
SLACK_AUTH_TOKEN={Foo}
SLACK_CHANNEL_ID={Bar}
```

The `.env` file must be **located in the working directory**.
<br>

### Execute

#### **Command line**

`go run cmd/sysup-notifier/sysup-notifier.go [--log /path/to/logfile.log] [--debug]`<br>

or as binary executable: <br>

`./sysup-notifier [--log /path/to/logfile.log] [--debug]`

optional parameter:

- "--log" creates logfile
- "--debug" set loglevel to DEBUG

#### **Cronjob**

`05 0 * * * cd <CWD> ; go run cmd/sysup-notifier/sysup-notifier.go [--log /path/to/logfile.log] [--debug]` <br>

or as binary executable: <br>

`05 0 * * * cd <CWD> ; sysup-notifier [--log /path/to/logfile.log] [--debug]`

Use Cronjob of root. For checking apt package manager `sudo` is needed/used.<br>

#### **Logfile**

By default, Logfile is created in the subdirectory `log/` of the working directory.<br>
Logrotation is not supported. Use an external program like `logrotate`.<br>
<br>
In case you don't want create a logfile, you can set the option `--log none`. Then no logfile will be created, but `Stdout` is used.

<br>

### Requirements

#### **Go installation**

For use, Go must be installed on the systems. Otherwise you have to build this code for your platform and run as a binary executable.

#### **Root permissions**

For checking apt package manager sudo permissions are needed. For this case, you have to use an user with sudo permissions. According your permission settings, you have to enter password for running the program..<br>
For Cronjob: test first user settings or run the job under root user.
<br>
<br>

## Contribute

Contributions are welcome!<br>
<br>

## Licence

Unless otherwise specified, all code is released under the [MIT License (MIT)](LICENSE).<br>
For used or linked components the respective license terms apply.
