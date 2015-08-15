# Refresh

This project is a fork of [fresh](https://github.com/pilu/fresh)
written by [Andrea Franz](http://gravityblast.com)

Refresh is a command line tool that builds and (re)starts your web
application everytime you save a Go source file.

## Installation

    go get github.com/baijum/refresh

## Usage

    cd /path/to/myapp

Start Refresh:

    refresh

Refresh will watch for file events, and every time you
create/modifiy/delete a file it will build and restart the
application.  If `go build` returns an error, it will log it in the
tmp folder.

You can use the `-c` options if you want to specify a config file.  By
default, `.refresh.conf` is used:

    refresh -c .refresh.conf

Here is a sample config file with the default settings:

    root:              .
    tmp_path:          ./tmp
    build_name:        runner-build
    build_log:         runner-build-errors.log
    valid_ext:         .go, .tpl, .tmpl, .html
    build_delay:       600
    colors:            1
    log_color_main:    cyan
    log_color_build:   yellow
    log_color_runner:  green
    log_color_watcher: magenta
    log_color_app:
    exclude_dir:
