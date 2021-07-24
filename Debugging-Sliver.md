Debugging Sliver binaries (server, client and implant) can be done using the [delve](https://github.com/go-delve/delve) debugger.

Delve can be installed with the following command:

```shell
go get github.com/go-delve/delve/cmd/dlv
```

The following examples are specific to Visual Studio Code, but other IDEs using delve probably have similar configuration options.

# Debugging the server

Debugging the Sliver needs to be done via delve [remote debugging](https://github.com/golang/vscode-go/blob/master/docs/debugging.md#remote-debugging) feature. This means you need to first start the binary with the following command:

```shell
dlv debug \
--build-flags="-tags osusergo,netgo,sqlite_omit_load_extension,server -ldflags='-X github.com/bishopfox/sliver/client/version.Version=1.1.2 -X github.com/bishopfox/sliver/client/version.CompiledAt=Never -X github.com/bishopfox/sliver/client/version.GithubReleasesURL=github.com -X github.com/bishopfox/sliver/client/version.GitCommit=aabbcc -X github.com/bishopfox/sliver/client/version.GitDirty=Dirty'" \
--headless \
--listen=:2345 \
--api-version=2 \
--log \
github.com/bishopfox/sliver/server
```

To simplify things, you can add this command as a VSCode task, by adding the following to your `.vscode/tasks.json` file:

```json
        {
            "label": "Run Debug Server",
            "type": "shell",
            "command": [
                "dlv"
            ],
            "args": [
                "debug",
                "--build-flags=\"-tags osusergo,netgo,sqlite_omit_load_extension,server -ldflags='-X github.com/bishopfox/sliver/client/version.Version=1.1.2 -X github.com/bishopfox/sliver/client/version.CompiledAt=Never -X github.com/bishopfox/sliver/client/version.GithubReleasesURL=github.com -X github.com/bishopfox/sliver/client/version.GitCommit=aabbcc -X github.com/bishopfox/sliver/client/version.GitDirty=Dirty'\"",
                "--headless",
                "--listen=:2345",
                "--api-version=2",
                "--log",
                "github.com/bishopfox/sliver/server"
            ],
            "presentation": {
                "echo": true,
                "reveal": "always",
                "focus": true,
                "showReuseMessage": false,
                "clear": true,
                "panel": "new",
            },
            "problemMatcher": [
                "$go"
            ]
        }
```

Then, you need to create your `.vscode/launch.json` file containing the following:

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Debug Server",
            "type": "go",
            "request": "attach",
            "mode": "remote",
            "remotePath": "${workspaceFolder}",
            "port": 2345,
            "host": "127.0.0.1",
            "trace": "log"
        },
    ]
}
```
Once you're all set, start by running the `Run Debug Server` task, and then hit `F5` (or use the UI to start the debugging task).

# Debugging the implant

To debug an implant, first make sure you built one by passing the `--debug` flag to the `generate` command. Then, add the following debug configuration to your `.vscode/launch.json` file:

```json
 {
            "name": "Debug Implant",
            "type": "go",
            "request": "attach",
            "mode": "remote",
            "remotePath": "",
            "port": REMOTE_PORT, // replace this
            "host": "REMOTE_HOST" // replace this
}
```
The `REMOTE_HOST` and `REMOTE_PORT` placeholders will need to be replaced to match the ones you specified on your delve server.

You will need to install the delve debugger on the target host. Once installed, run the delve server using the `exec` mode on your generated implant binary:

```shell
dlv exec --api-version 2 --headless --listen REMOTE_HOST:REMOTE_PORT --log .\implant.exe
```

Once the server is running on your target, select the `Debug Implant` debug configuration in VSCode and click on `Run`.