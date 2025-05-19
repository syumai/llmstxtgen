# llmstxtgen

* llmstxtgen is a tool to generate llms.txt files from sitemap.
* This tool is inspired by [timakin/llmstxt-gen](https://github.com/timakin/llmstxt-gen).

## Installation

```sh
go install github.com/syumai/llmstxtgen@latest
```

## Supported Modes

* full: generate llms-full.txt

## Usage

```sh
curl http://localhost:3000/sitemap.xml | llmstxtgen -mode=full > llms-full.txt
```

## Author

syumai

## License

MIT
