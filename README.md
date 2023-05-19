# system-update-notifier

## Description

Notify to Slack channel if system upgrades on Debian-based Linux are available.<br>
On DietPi: <br>
* for Dietpi upgrades, check /run/dietpi/.updates_available<br>
* for apt upgrades, check /run/dietpi/.apt_updates<br>

On Debian-based Linux checking updates by Apt package manager.<br>
<br>


## Getting started

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

* "--log" creates logfile
* "--debug" set loglevel to DEBUG

#### **Cronjob**

`05 0 * * * cd <CWD> ; go run cmd/sysup-notifier/sysup-notifier.go [--log /path/to/logfile.log] [--debug]` <br>

or as binary executable: <br>

`05 0 * * * cd <CWD> ; sysup-notifier [--log /path/to/logfile.log] [--debug]`

Use Cronjob of root. For checking apt package manager `sudo` is needed/used.<br>


#### **Logfile**
Logfile is created in the subdirectory `log/` of the working directory.

<br>

### Requirements
#### **Go installation**
For use, Go must be installed on the systems. Otherwise you have to build this code for your platform and run as a binary executable.

#### **Root persmissions**
For checking apt package manager sudo is needed. For this case, you have to use an user with sudo permissions. According your permission settings, you have to enter the password while running the service.<br>
For Cronjob it is recommended to run the job under root user.
<br>
<br>

## Contribute

Contributions are welcome!<br>
<br>

## Licence

Unless otherwise specified, all code is released under the [MIT License (MIT)](LICENSE).<br>
For used or linked components the respective license terms apply.