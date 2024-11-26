# BEER - "BackEnd Engineers Rocks"

## What is BEER?

A tool to help amazing backend developers to visualize their wonderful applications, revealing all the hidden beatiful of the backend layer to the world.

## How to use
TODO
## Setup
TODO

## Base Directory Structure

```shell
.
|── api							// API  - The REST API to interact with the CORE application (Gofiber)
|── cmd						    // CMD  - The CLI module implemented using Cobra
├── core						// CORE - The main module with their entities and services 
├── util						// UTIL - Commom functions to help in ordinary tasks
└── web							// WEB   - The user interface that consumes the REST API (Svelte)
```

## Building

First you need to install the dependencies and buid the web application:

```shell
cd web
npm install
npm run build
```

Then you can build the CLI application:

```shell
# go to the root folder
go build
```

The folder web/build contains the static files that are served by the web application. And this folder is embedded in the CLI application.

# Roadmap to v1.0.0
- [X] CORE: Identify model entities
- [ ] CORE: Create a scheme and validator to the config yml file
- [X] CLI: Read config
- [X] CLI: Load entities from config
- [ ] CLI: Support docker compose file to get Nodes, ports and volumes
- [ ] CLI: Support external Services 
- [ ] CLI: Sniff local network to get Nodes communication
- [ ] CLI: Watch Node volumes or directories
- [ ] CLI: Watch Node ports
- [ ] CLI: Watch Node db connections
- [ ] CLI: Node Healthcheck
- [ ] CLI: Node Logs
- [ ] CLI: Node Metrics
- [ ] CLI: Node CPU
- [ ] CLI: Node Memory
- [ ] CLI: Node Disk
- [ ] CLI: Node Network
- [X] CLI: Node Process
- [ ] CLI: Host System Info
- [ ] CLI: Host Environment Variables
- [X] API:  Get APPS
- [ ] API:  Get Host Info
- [ ] API:  Expose Nodes interactions via websockets or SSE in real-time
- [ ] API:  Expose Nodes metrics via websockets or SSE in real-time
- [X] UI:   Create UI
- [ ] UI:   Show Nodes related to the local apps 
- [ ] UI:   Show bottom panel with tabs to display different information and tools
- [ ] UI:   Show Nodes related to the external services (apis, dbs, etc)
- [ ] UI:   Show Nodes Info (cpu, mem, pid, process name, ports, type (service, vm or container), etc)
- [ ] UI:   Show Nodes Interactions Animations (websockets or SSE)
- [ ] UI:   Show Requests Node

# Ideas

## 1 

Add http requests to the config file to create a Node in UI with a list of requests to be made against the apps
following the syntax of ```.http``` files https://learn.microsoft.com/en-us/aspnet/core/test/http-files?view=aspnetcore-9.0#http-file-syntax


```yaml
requests: |

  GET http://localhost:4000/api/v1/apps
  Accept: application/json
  Content-Type: application/json
  
  ###

  GET /
  Accept: application/json
  Content-Type: application/json

```

or add 

```yaml
requests:
  - method: GET
    url: http://localhost:4000/api/v1/apps
    headers:
      Accept: application/json
      Content-Type: application/json
```

or even

```yaml
requestsFile: requests.http #or requests.rest
```

With the possibility to add environments 


