[![Continuous Integration Test Status](https://github.com/elfabri/learn-cicd-starter/actions/workflows/ci.yml/badge.svg)](https://github.com/elfabri/learn-cicd-starter/actions/workflows/ci.yml)
[![Continuous Deployment Status](https://github.com/elfabri/learn-cicd-starter/actions/workflows/cd.yml/badge.svg)](https://github.com/elfabri/learn-cicd-starter/actions/workflows/cd.yml)

# learn-cicd-starter (Notely)

This repo contains the starter code for the "Notely" application for the "Learn CICD" course on [Boot.dev](https://boot.dev).

## Local Development

Make sure you're on Go version 1.22+.

Create a `.env` file in the root of the project with the following contents:

```bash
PORT="8080"
```

Run the server:

```bash
go build -o notely && ./notely
```

*This starts the server in non-database mode.* It will serve a simple webpage at `http://localhost:8080`.

You do *not* need to set up a database or any interactivity on the webpage yet. Instructions for that will come later in the course!

Fabri's version of Boot.dev's Notely app.

# Conclusions

## continuous integrations

Automatic tests that are going to be run **on the creation of pull request to main**. This check the code for **style**, to keep a standard or follow a convention, and **run unit tests**.

## continuous integrations

Automatic commands to be run when there is a push into the **main** branch, when a pull request **is merged** into **main**. The commands are the normal ones to build the app.

I couldn't create a google cloud acc so I did not deploy the app there, might try that later. Also did not create acc with Turso, to host the database of the app. It should be to update the cd to deploy into GCP, link the database with an env var created into github secrets, and also into my local dev env, and install goose to run migrations both local and on the cd yml.

