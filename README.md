# Solana IDL Generator

This project is a tool for generating Go code from Solana IDL (Interface Definition Language) files. It helps in
creating Go structs and methods to interact with Solana programs.

## Installation

To install the Solana IDL Generator, you need to have Go installed on your system. Then, you can install the generator
using the following command:

```sh
go install github.com/exc-works/solana-idl-generator@latest
```

## Usage

To generate Go code from a Solana IDL file, use the following command:

```sh
solana-idl-generator <path-to-idl-file> --package <package-name>
```

Note: This tool works well when the IDL includes discriminator fields.

### Example

For example, to generate Go code for the `counter` example, you can run:

```sh
solana-idl-generator examples/counter/counter.json --package counter
```

This will generate a Go file in the `examples/counter` directory.

## Examples

You can find example IDL files and the generated Go code in the `examples` directory. Here are some examples:

- `examples/counter/counter.json` - IDL file for the counter example.
- `examples/counter/counter.go` - Generated Go code for the counter example.

## Testing

To run tests for the generated code, use the following command:

```sh
go test ./...
```

This will run all tests in the project, including those in the `examples` directory.

## License

This project is licensed under the MIT License. See the `LICENSE` file for more details.
