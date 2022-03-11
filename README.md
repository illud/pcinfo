# PC-INFO

\
PC-INFO is tool that gathers information on some of the main devices of your system.

![Hello World](https://raw.githubusercontent.com/saturnavt/eskolvar.github.io/main/assets/img/pcinfo.jpeg)

## Download
<a href="https://github.com/saturnavt/pcinfo/releases/download/Pre-Alpha/main.exe" target="_blank">CLICK HERE TO DOWNLOAD</a>

## Features
- BOARD name
- CPU name
- GPU name
- RAM
- HOSTNAME
- PLATFORM


## Want to help us improve
Create a new branch and make a pull request with the changes.

PC-INFO requires [Go](https://golang.org/) v1.15+ to run.

\
Install the dependencies.

```sh
go get ./...
```

Run.

```sh
go run main.go
```

Build.
```sh
go build -ldflags -H=windowsgui main.go
```

## Built with 
## https://fyne.io/


## License

MIT

PC-INFO is [MIT licensed](LICENSE).