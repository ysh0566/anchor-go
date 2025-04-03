# anchor-go

`anchor-go` generates Go clients for [Solana](https://solana.com/) programs (smart contracts) written using the [anchor](https://github.com/solana-foundation/anchor) framework.
My anchor version is 0.31.0.
## Installation

To install the `anchor-go` generator, you need to have Go installed on your system. Then, you can install the generator
using the following command:

```sh
go install github.com/ysh0566/anchor-go@latest
```

## Usage

To generate Go code from a Solana IDL file, use the following command:

```sh
anchor-go <path-to-idl-file> --package <package-name>
```

Note: This tool works well when the IDL includes discriminator fields.

### Example

For example, to generate Go code for the `counter` example, you can run:

```sh
anchor-go examples/counter/counter.json --package counter
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
