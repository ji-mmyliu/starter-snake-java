# Battlesnake Java Starter Project

An **un**official Battlesnake template for **Java users**, written in Go.

If you'd like to use repl.it to run your Battlesnake server, click the link below and press **"fork"** on the top right.

**[Run on Replit](https://replit.com/@JimmyLiu3/starter-snake-java)**

## Run Your Battlesnake (Repl.it)

To start your Battlesnake, press the "Start" button at the top of your repl.it server!

You should see the following at the top right of your repl.it app:

```json
{"apiversion":"1","author":"","color":"#888888","head":"default","tail":"default"}
```

## Run Your Battlesnake (Local)

To start your Battlesnake, press the "Start" button at the top of your repl.it server or run:

```sh
go run ./server
```

You should see the following output once it is running

```sh
Running your Battlesnake at http://0.0.0.0:8000
```

Open [localhost:8000](http://localhost:8000) and you should see

```json
{"apiversion":"1","author":"","color":"#888888","head":"default","tail":"default"}
```

## Verbose mode
To output the board layout at every turn, run the following command in the console:

```sh
go run ./server verbose
```

## Building your Battlesnake
Head to `Main.java` to start building your snake! There is more information provided in the file.

Try to make your Battlesnake smarter by avoiding other snakes and the edges of the board, going towards food, and suviving the arena as long as possible!

**Note:** To play games and tournaments, you'll need to deploy your Battlesnake to a live web server. Here is a [repl.it template](https://replit.com/@JimmyLiu3/starter-snake-java) that you can fork to get your server live.

You can also use [ngrok](https://ngrok.com/) if you choose to run the server locally.
