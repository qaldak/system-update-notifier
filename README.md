# system-update-notifier

## Description

Notify to Slack channel if system upgrades on Dietpi are available.</br>
* for Dietpi upgrades, check /run/dietpi/.updates_available</br>
* for apt upgrades, check /run/dietpi/.apt_updates


## Getting started

### Execute

#### Command line

`go run ... [--debug]`

optional parameter:

* "--debug" set loglevel to DEBUG

#### Cronjob

`05 0 * * * cd <PATH> ; go run ... [--debug]`

### Requirements

### Links

#### Tools

* [xyz](https://foo.bar/)

## Contribute

Contributions are welcome!

## Licence

Unless otherwise specified, all code is released under the [MIT License (MIT)](LICENSE).<br>
For used or linked components the respective license terms apply.