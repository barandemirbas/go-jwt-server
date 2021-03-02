<p align="center">

  <h3 align="center">Go JWT Server Boilerplate</h3>

  <p align="center">
    A server side Golang authentication boilerplate
    <br>
    <a href="#">Contribute</a>
    &middot;
    <a href="#">Explore</a>
    <br>
    <img src="https://img.shields.io/badge/license-MIT-green" />
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

Start [MongoDB](#mongodb-installation--setup) and run

```$ go get```

Rename ```.env.sample``` to ```.env``` and set your Mongo URI and Secret key in ```.env``` and start the backend and client side server as well as connect to the MongoDB database

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

## License

go-jwt-server is provided under the MIT license. See [LICENSE](LICENSE)
