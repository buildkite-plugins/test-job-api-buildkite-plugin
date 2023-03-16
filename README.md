# Job API Test Plugin

Demonstrates how to use the buildkite current job API to interact with the current job in languages other than bash

## Example

Add the following to your `pipeline.yml`:

```yml
steps:
  - command: ls
    plugins:
      - moskyb/job-api-test#main:
```

## Configuration
n/a

## Developing

To run the tests:

```shell
docker-compose run --rm tests
```

## Contributing

1. Fork the repo
2. Make the changes
3. Run the tests
4. Commit and push your changes
5. Send a pull request
