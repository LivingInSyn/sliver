Sliver allows an operator to extend the local client console and its features by adding new commands based on third party tools. The easiest way to install an alias or extension is using the [armory](https://github.com/BishopFox/sliver/wiki/Armory).

#### Aliases Command Parsing

__⚠️ IMPORTANT:__ It's important to understand that all alias commands have certain Sliver shell flags (those that appear in `--help`). Your Sliver shell commands are lexically parsed by the Sliver shell first, and only unnamed positional arguments are passed to the alias code. This means you may need to escape certain arguments in order for them to be correctly parsed.

For example with Seatbelt, `seatbelt -group=system` will fail because the Sliver shell will attempt to interpret the `-group` flag as a named flag (i.e., arguments that appear in `--help`). To ensure this argument is parsed as a positional argument we need to pass it as a string argument, so the correct syntax is `seatbelt \"-group=system\"`

```
[server] sliver (ROUND_ATELIER) > seatbelt -group=system
error: invalid flag: -group

[server] sliver (ROUND_ATELIER) > seatbelt \"-group=system\"

[*] seatbelt output:


                        %&&@@@&&
                        &&&&&&&%%%,                       #&&@@@@@@%%%%%%###############%
                        &%&   %&%%                        &////(((&%%%%%#%################//((((###%%%%%%%%%%%%%%%
%%%%%%%%%%%######%%%#%%####%  &%%**#                      @////(((&%%%%%%######################(((((((((((((((((((
#%#%%%%%%%#######%#%%#######  %&%,,,,,,,,,,,,,,,,         @////(((&%%%%%#%#####################(((((((((((((((((((
#%#%%%%%%#####%%#%#%%#######  %%%,,,,,,  ,,.   ,,         @////(((&%%%%%%%######################(#(((#(#((((((((((
#####%%%####################  &%%......  ...   ..         @////(((&%%%%%%%###############%######((#(#(####((((((((
#######%##########%#########  %%%......  ...   ..         @////(((&%%%%%#########################(#(#######((#####
###%##%%####################  &%%...............          @////(((&%%%%%%%%##############%#######(#########((#####
#####%######################  %%%..                       @////(((&%%%%%%%################
                        &%&   %%%%%      Seatbelt         %////(((&%%%%%%%%#############*
                        &%%&&&%%%%%        v1.1.1         ,(((&%%%%%%%%%%%%%%%%%,
                         #%%%%##,


====== AMSIProviders ======

  GUID                           : {2781761E-28E0-4109-99FE-B9D127C57AFE}
  ProviderPath                   : "C:\ProgramData\Microsoft\Windows Defender\Platform\4.18.2111.5-0\MpOav.dll"

====== AntiVirus ======

Cannot enumerate antivirus. root\SecurityCenter2 WMI namespace is not available on Windows Servers

...
```

Another trick is to provide a single empty string argument, after which all arguments will be parsed as positional e.g., `seatbelt '' -group=system`

## What's the difference between an alias and an extension? 

From an end-user perspective there's not much of a difference between the two, except that extensions' arguments will show up in `--help` and may be required.

An alias is essentially just a thin wrapper around the existing `sideload` and `execute-assembly` commands, and aliases cannot have dependencies. 

An extension is a shared library that is reflectively loaded into the Sliver implant process, and is passed several callbacks to return data to the implant. As such these extensions must implement the Sliver API. Extensions may also have dependencies, which are other extensions. For example, the COFF Loader is a DLL extension that loads and executes BOFs, in turn BOFs simply extensions that rely on the COFF Loader as a dependency. These types of extensions do not need to implement any Sliver-specific API, since the Sliver API is abstracted by their dependency.

## Aliases

A Sliver alias is nothing more than a folder with the following structure:

- an `alias.json` file
- alias binaries in one of the following formats:
  - .NET assemblies
  - shared libraries (`.so`, `.dll`, `.dylib`)

Here is an example for the `Rebeus` alias, reusing some of the public tools from [the GhostPack organisation](https://github.com/GhostPack):

```
$ tree GhostPack
Rubeus
├── alias.json
├── Rubeus.exe
```

The `alias.json` file has the following structure:

```json
{
    "name": "Rubeus",
    "version": "0.0.0",
    "command_name": "rubeus",
    "original_author": "@GhostPack (HarmJ0y)",
    "repo_url": "https://github.com/sliverarmory/Rubeus",
    "help": "Rubeus is a C# tool set for raw Kerberos interaction and abuses.",

    "entrypoint": "Main",
    "allow_args": true,
    "default_args": "",
    "is_reflective": false,
    "is_assembly": true,
    "files": [
        {
            "os": "windows",
            "arch": "amd64",
            "path": "Rubeus.exe"
        },
        {
            "os": "windows",
            "arch": "386",
            "path": "Rubeus.exe"
        }
    ]
}
```

It contains a single JSON object, which has the following fields:

### Alias Fields

| Field Name | Description |
| ---------- | ----------- |
| `name` | The stylized display name of the alias |
| `command_name` | The actual console command (primary identifier) |
| `entrypoint` | The entrypoint (only required for DLLs / Reflective DLLs) |
| `help` | A short help message for the command |
| `long_help` | A longer help message describing the command and its usage |
| `allow_args` | Specify whether the command will allow arguments or not |
| `files` | A list of of extension files |
| `is_reflective` | Indicates whether the extension is a reflective DLL or not |
| `is_assembly` | Indicates whether the extension is a .NET assembly or not |

#### Files

| Field Name | Description |
| ---------- | ----------- |
| `os` | The operating system for which the file can be used on (i.e., `GOOS` syntax)  |
| `arch` | The cpu architecture (i.e., `GOARCH` syntax) |
| `path` | Relative path to the file from the `alias.json`, parent directories are not allowed |

To load an alias in Sliver, use the `alias load` command:

```
sliver (CONCRETE_STEEL) > alias load /home/lesnuages/tools/misc/sliver-extensions/GhostPack/Rubeus

[*] Adding rubeus command: Rubeus is a C# toolset for raw Kerberos interaction and abuses.
[*] Rubeus extension has been loaded
```

The `help` command will now list the commands added by this extension:

```
sliver (CONCRETE_STEEL) > help
...
Sliver - 3rd Party extensions:
==============================
  rubeus    [GhostPack] Rubeus is a C# toolset for raw Kerberos interaction and abuses.
...
```

### Writing Aliases

To write a new alias, one must either create a shared library or a .NET assembly, then write a manifest file compliant with the description above.

As the alias support relies on Sliver side loading capabilities, please make sure to read the [Using 3rd party tools](https://github.com/BishopFox/sliver/wiki/Using-3rd-party-tools) section, to understand how shared libraries are loaded on all platforms.



## Extensions

Extensions are similar in structure to an alias, but work a little differently.