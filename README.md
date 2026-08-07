![logo](codewars-pretty-stats-baner.svg)

# Codewars pretty stats

A free API for retrieving great stats from your CodeWars account to spruce up your GitHub account, website, or anything else.

[![codewars_stats](https://codewars-pretty-stats.selector0073.com/?size=2&username=Selector0073)](https://github.com/Selector0073/codewars-pretty-stats/)

## Run Locally

Clone the project

```bash
  git clone https://github.com/Selector0073/codewars-pretty-stats
```

Go to the project directory

```bash
  cd codewars-pretty-stats
```

Create a `.env` file and fill it with the structure `.env.example`

Start the server

```bash
  go run cmd/codewars-api/main.go
```

Make a request to url with the correct [parameters](#paramaters)
```
http://localhost:4322/?size=1&username=Selector0073
```


## Running Tests

To run tests, run the following command

```bash
  go test ./...
```


## Usage/Examples

To use the projects, simply make a request with your [parameters](#paramaters) to
`https://codewars-pretty-stats.selector0073.com`
for example:

```Markdown
[![codewars_stats](https://codewars-pretty-stats.selector0073.com/?size=2&username=Selector0073)](https://github.com/Selector0073/codewars-pretty-stats/)
```
or
```HTML
<img src="https://codewars-pretty-stats.selector0073.com/?size=1&username=Selector0073">
```


## Parameters
Parameters should be passed as query params
```
https://codewars-pretty-stats.selector0073.com/?size=1&username=Selector0073
```

`size` is the size multiplier. This value must be an float from 0.5 to 5.

`username` is your codewars username.


## Contributing

Contributions are always welcome! You can fork the repository and after finishing work add a pull request describing the changes


## Roadmap

- [ ] Add caching
- [ ] Get the font locally, not from the internet
- [ ] Add codewars logo


## License

[MIT](https://choosealicense.com/licenses/mit/)

