# gogenapidemo

A demo project showcasing the integration between [TypeSpec](https://typespec.io/) and [ogen](https://ogen.dev/) for building type-safe APIs. This combination allows you to:

- Define your API using TypeSpec's clean and familiar TypeScript-like syntax
- Generate OpenAPI specifications automatically
- Generate Go server code using ogen's highly optimized, reflection-free implementation

## Dependencies

This project uses [mise](https://mise.jdx.dev/) for managing tool versions. You'll need:

- mise

## Getting Started

1. Clone the repository
2. Install dependencies:
   ```bash
   mise install
   ```

3. Run the server:
   ```bash
   mise run serve
   ```

## Development

When making changes to the API specification in `apispec/`, you'll need to regenerate the API code:

```bash
mise run generate
```

## Project Structure

- `apispec/` - Contains the TypeSpec API definition
- `api/` - Contains the generated ogen code
- `service/` - Contains the service implementation
- `cmd/` - Contains the server entrypoint

## Why TypeSpec + ogen?

This combination gives you the best of both worlds:
- TypeSpec's excellent developer experience and type-safe API design
- ogen's high-performance Go implementation with:
  - No reflection
  - No runtime overhead
  - Native Go type support
  - Built-in OpenTelemetry support