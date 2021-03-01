<p align="center">

  <h3 align="center">Go JWT Server Boilerplate</h3>

  <p align="center">
    A server side Golang authentication boilerplate
    <br>
    <a href="#">Contribute</a>
    &middot;
    <a href="#">Explore</a>
  </p>
</p>

<br>

## Go JWT Server

Simple JWT Auth Server is a quick Golang authentication boilerplate intended to get you up and running with a simple run command.

## Getting started

Get started by cloning the repository to your local machine

```
$ git clone git@github.com:barandemirbas/go-jwt-server.git
```

Start MongoDB

```
$ mongod
```
Rename ```.env.sample``` to ```.env```
Set your Mongo URI and Secret key in ```.env``` and start the backend and client side server as well as connect to the MongoDB database

```
$ go run main.go
```

## MongoDB Installation & Setup
### Docker
Pull the MongoDB Image
```
$ docker pull mongo
```
Start the Docker container
```
$ docker run --name mongo-db -p 27017:27017 -d mongo:latest
```

### Manual
Get system username by running

```
$ whoami
```

Set database and directory permissions for you by running the command

```
$ sudo chown -Rv <username> /data/db
# Enter your password
```
or for global access, run

```
$ sudo chown -r /data/db
# Enter your password
```
