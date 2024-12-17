# GOTHAT

GOTHAT is a highly opinionated framework for building progessive web applications.

It differentiates itself from most “modern” SPA web frameworks in that:

* it does not utilize a large, JavaScript based framework as its core,
* it is Server Side Rendered, using GO and Templ to render HTML,
* it uses htmx to render and replace portions of the DOM easily,
* it uses AlpineJS for easier, concise JavaSCript for user interactions,
* it uses TailwindCSS (currently v3.x) which is the least unusual aspect,
* and finally, it compiles to a single binary, making it very easy to deploy.

**GOTHAT** derives its name from the five main elements of the technology stack:

**GO** **T**empl **H**tmx **A**lpinejs **T**ailwindcss

## Modularized

GOTHAT is highly modularized, and is meant to be the core of a larger "modul-ithic" project.

Individual packages, see [weather](https://github.com/bartalcorn/weather), allow dev teams to independently create, maintain and own portions of a bigger system.

## Usage

Make a copy of `.env-example`, as `.env` and edit as needed.

`OWM_API_KEY` is only needed if using the [weather](https://github.com/bartalcorn/weather) package, and can be removed otherwise.

if not already installed, install go-task:

``` zsh
brew install go-task
```

if you are on a Mac but don't already have brew installed, [Homebrew](https://brew.sh)

if you are not on a Mac, install `go-task` via: [taskfile.dev](https://taskfile.dev/installation/)

run `task` alone to see the list of available tasks and their descriptions.

then run:

``` zsh
task updair
task updtempl
task updtailwind
go mod tidy
task dev
```

## Available TASK commands

``` task
             this is your main command
dev          run all three below, in watch mode, combined

templ        run templ generate in stand alone mode
templw       run templ generate in watch mode with proxy

css          run tailwind in stand alone mode
cssw         run tailwind in watch mode

air          run air, hot reloading of GO application

build        build GO application for linux

updair      install or update air https://github.com/air-verse/air
updtempl    install or update templ https://templ.guide/
updtailwind install or update tailwind https://tailwindcss.com/
```
