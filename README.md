# system-update-notifier

## Description

Notify to Slack channel if system upgrades on Debian-based Linux are available.</br>
On DietPi: </br>
* for Dietpi upgrades, check /run/dietpi/.updates_available</br>
* for apt upgrades, check /run/dietpi/.apt_updates</br>

On Debian-based Linux checking updates by Apt package manager.</br>


## Getting started

### Execute

#### Command line

`go run ... [--debug]`

optional parameter:

* "--debug" set loglevel to DEBUG

#### Cronjob

`05 0 * * * cd <PATH> ; go run ... [--debug]`

Use Cronjob of root. For checking apt package manager `sudo` is needed/used.

### Requirements
#### Go installation
For use, Go must be installed on the systems. Otherwise you have to build this code for your platform and run as app.

#### Root persmissions
For checking apt package manager sudo is needed. For this case, you have to use an user with sudo permissions. According your permission settings, you have to enter the password while running the service.</br>
For Cronjob it is recommended to run the job under root user.

## Contribute

Contributions are welcome!

## Licence

Unless otherwise specified, all code is released under the [MIT License (MIT)](LICENSE).<br>
For used or linked components the respective license terms apply.