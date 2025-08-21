# gogenapidemo

A demo project showcasing the integration between [TypeSpec](https://typespec.io/) and [ogen](https://ogen.dev/) for building type-safe APIs. This combination allows you to:

- Define your API using TypeSpec's clean and familiar TypeScript-like syntax
- Generate OpenAPI specifications automatically
- Generate Go server code using ogen's highly optimized, reflection-free implementation

The demo includes a simple Widget service implementation showcasing CRUD operations and custom analysis endpoints.

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
   mise run server
   ```

4. Access the API documentation at [http://localhost:8080/docs/](http://localhost:8080/docs/)

## Development

When making changes to the API specification in `apispec/`, you'll need to regenerate the API code:

```bash
mise run generate
```

This will:
1. Generate OpenAPI JSON spec from TypeSpec
2. Copy the spec to the service directory
3. Generate Go code using ogen

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