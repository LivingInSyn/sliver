__NOTE:__ Any platform can cross-compile a standalone executable to any other platform out of the box, you only need cross-compilers when using `--format shared`.

Sliver can tell you which platforms it can likely target based on the server's platform and available cross-compilers by running the `generate info` command in the console.


## From Linux to MacOS/Windows

To compile Windows shared library and shellcode implants from Linux, install mingw from your local package manager: 

```
sudo apt install mingw-w64
```

To compile MacOS shared library implants from Linux, we recommend using https://github.com/tpoechtrager/osxcross by default Sliver will look in `/opt/osxcross` but you can override this via [environment variables](https://github.com/BishopFox/sliver/wiki/Environment-Variables). If you do not have a MacOS based machine you can use GitHub Actions' MacOS instances to build OSXCross.

Sliver automatically looks in the default paths for these cross compilers, once installed simply use the `generate` command with the desired `--os` and `--arch`, check `~/.sliver/logs/sliver.log` for build errors. You can override any cross compiler location via the appropriate [environment variables](https://github.com/BishopFox/sliver/wiki/Environment-Variables).

## From MacOS to Linux/Windows

To compile Windows shared library and shellcode implants from MacOS install mingw from brew:

```
brew install mingw-w64
```

For Linux, we recommend `musl-cross` to target 64-bit Linux, which can be installed via brew:

```
brew install FiloSottile/musl-cross/musl-cross
brew install mingw-w64
```

I'm not aware of any good options to target 32-bit Linux from MacOS. Sliver automatically looks in the default paths for these cross compilers, once installed simply use the `generate` command with the desired `--os` and `--arch`, check `~/.sliver/logs/sliver.log` for build errors. You can override any cross compiler location via the appropriate [environment variables](https://github.com/BishopFox/sliver/wiki/Environment-Variables).

## From Windows to MacOS/Linux

Good luck.