# llmstxtgen

* llmstxtgen is a tool to generate llms.txt files from sitemap.
* This tool is inspired by [timakin/llmstxt-gen](https://github.com/timakin/llmstxt-gen).

## Installation

```sh
go install github.com/syumai/llmstxtgen/cmd/llmstxtgen@latest
```

## Supported Modes

* full: generate llms-full.txt

## Usage

```sh
curl http://localhost:3000/sitemap.xml | llmstxtgen -mode=full > llms-full.txt
```

## Options

* -mode: generate mode: full
* -sleep: sleep time (milli seconds, default: 100)

## Author

syumai

## License

MIT
