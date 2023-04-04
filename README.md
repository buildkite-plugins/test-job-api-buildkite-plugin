# Job API Test Plugin

Demonstrates how to use the buildkite current job API to interact with the current job in languages other than bash.

## Usage

Ensure that you've built the golang portion of the plugin, and placed it in the `go/bin` directory:

```shell
mkdir -p $(go env GOPATH)/bin
go build -o $(go env GOPATH)/bin/job-api-test
```

This plugin is a no-op, but it demonstrates how to use the job API to add and remove environment variables from the current job. The plugin requires the built binary to be accessible to the agent, so this plugin isn't suitable for use in production environments.

## Example

Add the following to your `pipeline.yml`:

```yml
steps:
  - command: ls # or anything, really
    plugins:
      - test-job-api#main:
```

It will add two environment variables to the job during the `environment` phase, `OCEAN` and `MOUNTAIN`. It will then remove the `OCEAN `environment variable during the `post-command` phase.

## Configuration
n/a

## Contributing

1. Fork the repo
2. Make the changes
3. Run the tests
4. Commit and push your changes
5. Send a pull request
