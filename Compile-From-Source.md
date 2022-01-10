You'll want to compile from a MacOS or Linux machine, compiling from native Windows in theory should be possible, but none of the scripts are designed to run on Windows (you can cross compile the Windows server/client binaries from MacOS or Linux). If you only have a Windows machine see "Windows Builds" below (TL;DR use WSL).

### From Scratch (No Docker)

From scratch without Docker, requirements for compiling:

* Go v1.17 or later
* `make`, `sed`, `tar`, `curl`, `zip`, `cut` commands

#### Go Assets

Sliver embeds its own version of the Go compiler and related tools. Run the following script to download these static assets. You only need to do this once before running `make`:

```
$ ./go-assets.sh
```

Then compile and bundle server with all dependencies and assets, by default `make` will build whatever platform you're currently running on:

```
$ make
```

#### Cross-compile to Specific Platforms

You can also specify a target platform for the `make` file, though you may need cross-compilers (see below):

```
$ make macos
$ make macos-arm64
$ make linux
$ make windows
```

### Docker Build

The Docker builds are mostly designed for running unit tests, but can be useful if you want a "just do everything" build, you just need to have Docker installed on the machine. First `git clone` the Sliver repo, then run the `build.py` script (the script isn't required but has a few short cuts see `./build.py --help`). Alternatively, execute the following command: 

```
docker build -t sliver .
```

The Docker build includes mingw and Metasploit, so it can take a while to build from scratch but Docker should cache the layers effectively. Sliver will also run it's unit tests as part of the build, and that takes a few minutes too.

### Windows Builds

__NOTE:__ Starting with v1.1.0 in order to cross-compile Windows builds you'll need `mingw` on your system:
* Ubuntu `sudo apt install mingw-w64`
* MacOS `brew install mingw-w64`

If all you have is a Windows machine, the easiest way to build Sliver is using [WSL](https://docs.microsoft.com/en-us/windows/wsl/install-win10) and following the Linux/cross-compile instructions above. To cross-compile a native Windows binary use `make windows` and copy it to your Windows file system (i.e. `/mnt/c/Users/foo/Desktop`) and run it using a terminal that supports ANSI sequences such as the [Windows Terminal](https://github.com/microsoft/terminal).

## Developers

If you want to modify any of the `.proto` files you'll need to setup a few additional tools to regenerate the `.pb.go` files.
* Protoc v3.15.8
* Protoc-gen-go v1.26.0
* Protoc-gen-go-grpc v1.1.0

#### `protoc`

First install your platform's version of `protoc` v3.15.8 or compile this version from source:

https://github.com/protocolbuffers/protobuf/releases/tag/v3.15.8

Ensure that `protoc` v3.15.8 is on your `$PATH`, you can check with a simple `protoc --version`

#### `protoc-gen-go` `protoc-gen-go-gprc`

Assuming `$GOPATH/bin` is on your `$PATH` simply run the following commands to install the appropriate versions of `protoc-gen-go` and `protoc-gen-go-gprc`:

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26.0
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
```

Ensure that these are both on your `$PATH` after running the commands, if not you probably need to add `$GOPATH/bin` to your `$PATH`. To regenerate the Protobuf and gRPC files run:

```
$ make pb
```
