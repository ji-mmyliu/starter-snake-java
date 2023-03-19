# Battlesnake Java Starter Project

An **un**official Battlesnake template for **Java users**, written in Go.

[![Run on Replit](https://replit.com/@JimmyLiu3/starter-snake-java)](https://replit.com/@JimmyLiu3/starter-snake-java)


## Run Your Battlesnake

To start your Battlesnake, press the "Start" button at the top of your repl.it server or run:

```sh
go run ./server
```

You should see the following output once it is running

```sh
Running your Battlesnake at http://0.0.0.0:8000
```

Open [localhost:8000](http://localhost:8000) in your browser and you should see

```json
{"apiversion":"1","author":"","color":"#888888","head":"default","tail":"default"}
```

## Verbose mode
To output the board layout at every turn, run the following command in the console:

```sh
go run ./server verbose
```

## Next Steps

**Note:** To play games and tournaments, you'll need to deploy your Battlesnake to a live web server such as [repl.it](https://repl.it)