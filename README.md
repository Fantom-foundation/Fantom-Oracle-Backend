# Fantom Oracle Watchdog

The repository contains implementation of high performance blockchain watchdog 
for oracle contracts off-chain world interaction.

The watchdog is responsible for monitoring oracle smart contracts activity on block chain, 
especially for emitted events on contracts, and respond with relevant data from off-chain 
world needed to perform on-chain actions. Special modules can also feed on-chain contracts 
with external data, based on specified criteria, timer, or API response.  

## Building the source

Building **Watchdog** requires GIT package and Go (version 1.14 or later is recommended). You can install
it using your favourite package manager. The latest version of Go can be installed directly 
from [GoLang Website](https://golang.org/). 

Once you have the Go environment ready, clone the Watchdog repository from GitHub 
and build the binary package:

```shell
git clone https://github.com/Fantom-foundation/oracle-watchdog.git
go build -o ./build/watchdog ./cmd/watchdog
```

The build output is `build/watchdog` executable.

You don't need to clone the project into `$GOPATH` due to Go Modules tooling, 
use any suitable location. We recommend moving the built Watchdog binary 
to your `bin` path and using `Systemd` unit to manage the Watchdog as a service 
for production use.  

## Running the Watchdog server

You need access to an RPC interface of an **Opera Lachesis** node to run the **Watchdog** server. 
Please follow [Lachesis](https://github.com/Fantom-foundation/go-lachesis) instructions 
to build and run the node. You can obtain access to a remotely running instance
of Lachesis, too. 

We recommend using local IPC channel for communication between a Lachesis node and the 
Watchdog server for performance and security reasons. Please consider security implications 
of opening Lachesis RPC to outside world access.

### System.d Service unit file

To run the **Watchdog** as a system service on Linux, create a service unit file on appropriate location. 
The actual place for putting the service file may vary by Linux distribution. For example, you can use
`/etc/systemd/system/watchdog.service` file path on Ubuntu systems.

We assume you want to use `/var/opera/watchdog` as the working directory for the Watchdog and that 
you copied the Watchdog binary to `/usr/bin/watchdog/`. In that case, the recommended 
`.service` file  content is:

```
[Unit]
Description=Fantom oracle Watchdog service
After=network.target auditd.service

[Service]
Type=simple
User=opera
Group=opera
WorkingDirectory=/var/opera/watchdog
ExecStart=/usr/bin/watchdog \
            --rpc /var/opera/lachesis/data/lachesis.ipc \
            --contracts /var/opera/watchdog/contracts.json
OOMScoreAdjust=-900
Restart=on-failure
RestartSec=5s

[Install]
WantedBy=multi-user.target
Alias=watchdog.service
```

Adjust the service unit file to match your path and configuration details for Opera RPC interface,
work path and Watchdog binary file location.

Don't forget to update the System.d status to be able to use the new service file to start and stop 
the Watchdog: `systemctl daemon-reload`. Manage the service start/stop using usual System.d commands, 
i.e. `systemctl start watchdog.service`.
