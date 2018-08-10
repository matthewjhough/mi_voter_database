# info

This is a simple containerized microservice written in go that will build a database of Michigan registered voters, their voting history, and expose that data over a simple rest api for consumption by other microservices or any client you want to setup.

# Requirements

You'll want docker and kubernetes installed to spin this up. There is also a makefile to keep things simple... so having make installed would be a benefit.

The app is written in Go, but all the building is done inside a docker container... so there is no need to install any of that on your development or production machine.

# Configuration

There is a basic config file. For testing you can simply copy the config.sample.yaml in the config dir and save it as config.yaml.

# Auth

The api is wrapper in a JWT middleware. There isn't any code that generates jwts. So you'll either need ot manually build them or setup an Auth Microservice to hand jwts out.

For testing if you set the jwt key to "test" you can use this JWT:

eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.5mhBHqs5_DTLdINd9p5m7ZJ6XD0Xc55kIaCRY5r6HRA

Obviously, if you're exposing the api to the realworld you'll want to setup a real secret key.

# Getting Started

Clone this repo

git clone git@github.com:nathanmentley/mi_voter_database.git

Download a raw voter dump from this site:
http://michiganvoters.info/download.html
(The latest was from 12 May 2017 as of this writing.)
Extract that into the {reporoot}/data directory.

Before starting anything you'll want a docker image registry running for kubernetes to pull from. The makefile makes that easy to start up. Run:

$ make setup

That will run a docker image registry locally inside docker exposed under port 5000.

To build the go app you can simply run:

$ make build

Once you have the go binary build you can package it in a docker image by running:

$ make pack

If the image is built without error you'll need to upload it to your docker registry. Run:

$ make upload

Now that you have the app build, packaged, and uploaded. You can simply deploy it using kubernetes. Run:

$ make deploy

Once you have everything running you can easily change the go source code and deploy an update by running:

$ make build;make pack;make upload;make deploy

Once you're done developing you can trash all the kubernetes stuff by running:

$ make stop
** WARNING: This trashes everything. Even things not related to this project. If you're runnign more stuff in kube you should probably not do this.


# Importing the voter registration data

Once you have the app running you'll quickly find it's pretty useless without data.

The kubernetes deploy sets up a kube cronjob that runs at midnight to import / ensure the data. However, If you want to import the data right away you can run:

make ensure

** WARNING: Right now there is a PersistentVolumeClaim that the data importer will be trying to load the data from. It's called voter-service-data-pv-claim. You'll need to manually copy the data from the {repoRoot}/data directory. In the future I'll have the "make deploy" logic do that.

# Example Request

localhost/voter?query={"Limit":2,"Offset":0,"Include":[],"Filters":[{"Field":"FirstName","Value":"NATHAN"},{"Field":"LastName","Value":"MENTLEY"}]}

curl -X GET 'http://localhost/voter?query={%22Limit%22:2,%22Offset%22:0,%22Include%22:[%22VoterHistories%22],%22Filters%22:[{%22Field%22:%22FirstName%22,%22Value%22:%22NATHAN%22},{%22Field%22:%22LastName%22,%22Value%22:%22MENTLEY%22}]}' -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.5mhBHqs5_DTLdINd9p5m7ZJ6XD0Xc55kIaCRY5r6HRA'
