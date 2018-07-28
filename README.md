# info

This is a simple containerized microservice written in go that will build a database of Michigan registered voters and their voting history and expose that data over a simple rest api for consumption by other microservices or any client you want to setup.

# Requirements

The go app is setup to be compiled inside a docker container and run from there. The project also includes a simple makefile to make the process easier.
So at a minimum docker is required, but you'd probably want make installed as well.

# Configuration

As of 7-25-2018 there isn't a configuration file. If you look at the make file we're passing env variables over the docker commands.
In the next few commits I'm planning on setting up a yaml config file to change the config. So until that is figured out you'll need to update the makefile to point to a valid MySQL server for the database import. In the future I'll have a k8s config to spin up the database with tihs.

# Auth

As of this writting there isn't any sort of security around the microservice. If this is running inside a kube cluster or something without outside access that may be okay, but in the future I'll be adding configurable jwt parsing / checking. This will require another auth microservice to hand out jwts, but I'll hopefully have an example auth microservice setup for that.

# Getting Started

Clone this repo

git clone git@github.com:nathanmentley/mi_voter_database.git

Download a raw voter dump from this site:
http://michiganvoters.info/download.html
(The latest was from 12 May 2017 as of this writing.)
Extract that into the {reporoot}/data directory.

Build the go binary by running:

make build

Build the containerized docker image by running

make pack

You can run the binary by running

make run

However it's probably easier to run

make ensure;make serve

the first command will ensure the database schema is setup and the data files are imported. The second command will spin up the api.
