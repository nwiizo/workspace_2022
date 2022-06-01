# README #

# Installation

## redoc-cli

Install redoc-cli as described here: https://github.com/Redocly/redoc/tree/master/cli

Note that you can use the good old npm instead of npx.

## http-server

Install http-server as described here: https://github.com/indexzero/http-server

Note that http-server is only needed to serve the svg file referenced in the OpenAPI files.

# Running the examples using redoc-cli

In `code/chapter_12`:

- Launch a terminal and start http-server with `http-server -p 9090`
- Launch a second termine and start redoc-cli for one of the OpenAPI file with `redoc-cli serve <OpenAPI filename> --options.showExtensions`
- Open a web browser and go to this url: http://localhost:8080

```
redoc-cli serve 12.1-reference-documentation-openapi.yaml --options.showExtensions
redoc-cli serve 12.2-operation-manual-openapi.yaml --options.showExtensions
redoc-cli serve 12.3-implementation-spec-openapi.yaml --options.showExtensions
redoc-cli serve 12.4-change-log.yaml --options.showExtensions
```