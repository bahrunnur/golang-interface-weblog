# Golang Interface Demonstration

there are 2 folders that demonstrate how 2 consumers is using database implementation:
1. `happy-consumer` define their own interface to `database` package
2. `slightly-unhappy-consumer` use a defined interface from `database` package

## Running the test
I use Makefile, so execute this:
```bash
make test
```