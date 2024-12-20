# BEER - "BackEnd Engineers Rocks"

## What is BEER?

A tool to help amazing backend developers to visualize their wonderful applications, revealing all the hidden beatiful of the backend layer to the world.

## How to use
TODO
## Setup
TODO

### Environment Variables

| Name | Description | Default |
| --- | --- | --- |
| BEER_CONFIG_PATH | The path to the project config file. If not informed, it will be searched in the current directory | "config.beer.yml" |
| BEER_HOST_INFO_INTERVAL | The interval in seconds that BEER will scan the host info and send the event "hostInfo" to the "/sse" endpoint | 3 |

## Base Directory Structure

```shell
.
|── api							// API  - The REST API to interact with the CORE application (Gofiber)
|── cmd						    // CMD  - The CLI module implemented using Cobra
├── core						// CORE - The main module with their entities and services 
├── util						// UTIL - Commom functions to help in ordinary tasks
└── web							// WEB   - The user interface that consumes the REST API (Svelte)
```

## Building and Running the Application

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
- [ ] CLI: Support docker compose file to get Apps, ports and volumes
- [ ] CLI: Support external Services 
- [ ] CLI: Sniff local network to get Apps communication
- [ ] CLI: Watch App volumes or directories
- [ ] CLI: Watch App ports
- [ ] CLI: Watch App db connections
- [ ] CLI: App Healthcheck
- [ ] CLI: App Logs
- [ ] CLI: App Metrics
- [ ] CLI: App CPU
- [ ] CLI: App Memory
- [ ] CLI: App Disk
- [ ] CLI: App Network
- [X] CLI: App Process
- [X] CLI: Host System Info
- [ ] CLI: Host Environment Variables
- [X] API:  Get APPS
- [X] API:  Get Host Info
- [ ] API:  Expose Apps interactions via websockets or SSE in real-time
- [ ] API:  Expose Apps metrics via websockets or SSE in real-time
- [X] UI:   Create UI
- [ ] UI:   Show Apps related to the local apps 
- [ ] UI:   Show bottom panel with tabs to display different information and tools
- [ ] UI:   Show Apps related to the external services (apis, dbs, etc)
- [ ] UI:   Show Apps Info (cpu, mem, pid, process name, ports, type (service, vm or container), etc)
- [ ] UI:   Show Apps Interactions Animations (websockets or SSE)
- [ ] UI:   Show Requests App
- [ ] UI:   Create logo
- [ ] UI:   Favicon
- [ ] UI:   Canvas watermark with beer logo

# Ideas

## 1 

Add http requests to the config file to create a App in UI with a list of requests to be made against the apps
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


