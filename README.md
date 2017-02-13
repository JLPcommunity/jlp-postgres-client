# Postgres Client 
JLP Meetup: [Learn PostgreSQL](http://bit.ly/2kN6wvX) on Feb 14th, 2017 @ EV Hive, Jakarta.

All the instruction below, start from the folder `jlp-postgres` or any name if you renamed it when clone.

## Run Go Server
We assume you have already installed Go on your computer. Otherwise, refer to [this](https://golang.org/doc/install).
```sh
$ cd server-go
$ go run main.go
```
**Note:** It listens to `localhost:8383`

## Run Angular 2 Client
```sh
$ cd client-ng2
$ npm install
$ ng serve
```
**Note:** It usually runs at `localhost:4200`

## Test the app
Open your browser and type this URL `http://localhost:4200`.
