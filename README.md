# Todos CLI App

## Description

The goal of this project is to create a todos cli app in [Go](https://go.dev) to learn the basics of [Go](https://go.dev). At the beginning, I did the project with a csv file and after that I change to an sqlite database. But there is a way to export in `csv` or `json`

## Tech Stack

- [Cobra](https://cobra.dev)
- [Huh](https://github.com/charmbracelet/huh)
- [Sqlite Driver](http://github.com/mattn/go-sqlite3)
- [Tablewriter](https://github.com/olekukonko/tablewriter)
- [UUID](https://github.com/google/uuid)

## Tasks

- [x] Install all dependencies.
- [x] Create CRUD interaction with the database.
- [x] Create the `create` command (Initialize the database).
- [x] Create the `add` command (Add a todo in the database).
- [x] Create the `list` command (List all todos which come from the database)
  - [x] View only the todos which are not done.
  - [x] When there is the flag `-a` or `--all` view all todos (done and not done).
- [x] Create the `done` command.
- [x] Create the `export` command.
  - [x] Export in JSON.
  - [x] Export in CSV.
  - [x] If no format argument provided open a TUI with huh to choose the format.
- [x] Create a `remove` command (Delete a todo with his id).
  - [x] Add a confirm TUI interface.
  - [ ] Add a force flag
- [x] Create the `reset` command (Delete the database).
  - [x] Add a confirm TUI interface.
  - [ ] Add a force flag
- [x] Create a better error handling.
- [x] Create tests.
- [x] Create a Makefile.
- [x] Add Github Actions
