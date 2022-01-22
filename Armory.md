The armory is the Sliver Alias and Extension package manager, which was introduced in Sliver v1.5 it allows you to automatically install various 3rd party tools such as BOFs and more.

The armory downloads packages from `github.com` and `api.github.com` so you'll need an internet connection in order for the command to work. The command does support proxies (see `--help`) and after an alias or extension is installed an internet connection is not required to execute the alias/extension.

Aliases and extensions are installed on the "sliver client"-side, and thus are not shared among operators in multiplayer mode.

## The Official Armory

The official armory ships with Sliver binaries and is included by default in the `Makefile` when compiling from source. You can interact with the Armory using the `armory` command. Packages installed from the official armory are compiled and cryptographically signed by the Sliver authors. While we make a best effort to review 3rd party code, you are responsible for reviewing and understanding any 3rd party code before using it.

#### Installing Packages

List available packages by running the `armory` command without arguments, packages are installed using the command name: 

```
sliver > armory install rubeus

[*] Installing alias 'Rubeus' (v0.0.21) ... done!
```

#### Updating Packages

You can update all installed aliases and extensions by running `armory update` command.

```
sliver > armory update 

[*] All aliases up to date!
[*] 1 extension(s) out of date: coff-loader
[*] Installing extension 'coff-loader' (v1.0.10) ... done!
```

#### Removing Packages

You remove packages installed from the `armory` using the `aliases rm` and `extensions rm` commands depending on if the package is an alias or an extension. You can list installed aliases and extensions by running `aliases` and `extensions` respectfully.

## Private Armories

Sliver also supports self-hosted private armories, but I've not gotten around to writing the details about this so you'll have to read thru the source code to figure out how they work for now.