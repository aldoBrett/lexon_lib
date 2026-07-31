
# Lex Engine
Lex is law in latin.

## Objective
The objective of this engine is to create the state machine (or Petri Net) that will be used to model the legal processes of Mexican Law (for now).

## Technology
The library will be coded in GoLang in order to achieve the maximum performance. Also because we're going to import it to each of the API servers that are going to need it.
It's going to have a data layer that uses PostgreSQL as a backend in order to save the current state of the network. Since it's a discrete machine that will be moving depending on the user signals.


# DB

## Migrations
To run the migration use the following command:

`sh
go run ./cmd/migrate/
make migrate        # I don't know how to run this command right now
`

## Seeds
To run the seeds use the command:

`sh
go run ./cmd/seed
`
