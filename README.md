# redis-env

redis-env allows you to easily use environment variables for configuring your application. The app config is stored in a centralized Redis database.

## Why?

* Storing application config in environment variables is better than storing it in configuration files. See [the Twelve-Factor app](http://www.12factor.net/config) for more.
* Redis makes it easy to store and update the config centrally without sharing disks or pushing files out to all your servers.
* We're not all running on [Heroku](http://www.heroku.com/). :)

## Install

    git clone git://github.com/brynary/redis-env.git
    go build
    mv redis-env /usr/local/bin

## Usage

### Add config variables

    redis-env --add DATABASE_URL=mysql://user@db-host.local

### Remove config variables

    redis-env --remove DATABASE_URL

### List config variables

    redis-env --list

Prints each config variable, one per line.

### Run

    redis-env --run rackup

Loads the config as environment variables then execs the provided command. Exist with code 111 if the config could not be loaded from Redis, otherwise exits with the code from the executed command.

## Options

* `--netaddr NETADDR` -- The location of the Redis instance. Defaults to `tcp:120.0.0.1:6379`
* `--db DB_INDEX` -- The Redis database index for the config data. Defaults to 0.
* `--name NAME` -- The name of the config to read. Defaults to `default`. Useful for storing multiple app configs in the same Redis.

_Note:_ This is the first Go code I've ever written. Please be gentle.
