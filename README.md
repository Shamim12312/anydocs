# anydocs

**anydocs** is a CLI tool built for developers using local coding agents such as Cursor, Claude Code, Gemini CLI and GitHub Copilot.

Its main aim is to fetch documentation starting from any URL that exposes raw text: this means that all the docs you need for that vibe coding project are just a GET request away!

Not only: **anydocs** combines the performance of Go with the benefits of concurrency, executing multiple fetch operations at a time so that you can have your documentation ready even faster.

And if you do need an extra help in organizing the docs, you can count on the AI summary functionality: Claude 4 Sonnet will take the whole fetched documentation clump and will produce a neat and ready-to-go summary!

## Install

In order to install **anydocs** there are three ways:

1. Using `go`: if you already have `go` 1.23+ installed in your environment, installing **anydocs** is effortless

```bash
go install github.com/AstraBert/anydocs
```

2. Using `npm`:

```bash
npm install @cle-does-things/anydocs
```

3. Downloading the executable from the [releases page](https://github.com/AstraBert/anydocs/releases): you can download it directly from the GitHub repository or, if you do not want to leave your terminal, you can use `curl`:

```bash
curl -L -o anydocs https://github.com/AstraBert/anydocs/releases/download/<version>/anydocs_<version>_<OS>_<processor>.tar.gz ## e.g. https://github.com/AstraBert/anydocs/releases/download/0.1.1/anydocs_0.1.1_darwin_amd64.tar.gz

# make sure the downloaded binary is executable (not needed for Windows)
chmod +x anydocs
```

In this last case, be careful to specify your OS (supported: linux, windows, macos) and your processor type (supported: amd, arm).

## Run

**anydocs** has two commands, `fetch` and `gh`.

### `fetch`

**Manual**

```text
Fetch documentation contant by passing the endpoint URLs (comma-separated, flag -u,--urls) and the path to which you would like to save this documentation (flag -p, --path). Optionally, you can also decide to produce an AI summary of the documentation (flag -s, --summary).

Usage:
  anydocs fetch [flags]

Aliases:
  fetch, f

Flags:
  -h, --help          help for fetch
  -p, --path string   Pass the path you want to save your files at
  -s, --summary       Use this flag if you want to enable AI summary of fetched documentation.
  -u, --urls string   Pass a set of llms.txt endpoints, comma separated (e.g. 'https://docs.llamaindex.ai/en/latest/llms.txt,https://raw.githubusercontent.com/AstraBert/anydocs/main/README.md')
```

**Example Usage**

```bash
# with AI summary
anydocs fetch --urls 'https://raw.githubusercontent.com/AstraBert/anydocs/main/README.md' --path CLAUDE.md --summary
# without AI summary
anydocs fetch --urls 'https://raw.githubusercontent.com/AstraBert/anydocs/main/README.md' --path CLAUDE.md
```

## `gh`

**Manual**

```text
Fetch documentation content by passing URLs of GitHub files (comma-separated, flag -u,--urls) and the path to which you would like to save this documentation (flag -p, --path). Optionally, you can also decide to produce an AI summary of the documentation (flag -s, --summary).

Usage:
  anydocs gh [flags]

Aliases:
  gh, g

Flags:
  -h, --help          help for gh
  -p, --path string   Pass the path you want to save your files at
  -s, --summary       Use this flag if you want to enable AI summary of fetched documentation.
  -u, --urls string   Pass a set of GitHub URLs, comma separated (e.g. 'https://github.com/AstraBert/PdfItDown/blob/main/README.md,https://github.com/AstraBert/anydocs/tree/main/README.md')
```

**Example Usage**

```bash
anydocs gh --urls 'https://github.com/AstraBert/anydocs/blob/main/README.md' --path CLAUDE.md
# with AI summary:
anydocs gh --urls 'https://github.com/AstraBert/anydocs/blob/main/README.md' --path CLAUDE.md --summary
```

## Contributing

We welcome contributions! Please read our [Contributing Guide](./CONTRIBUTING.md) to get started.

## License

This project is licensed under the [MIT License](./LICENSE)
