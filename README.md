# Example Go API

This is an educational repo, aimed at showing how to create a simple GoLang API for PHP developers. This was streamed on YouTube [here](https://youtu.be/EnWIg_IZg_8).


## Run locally

To run this locally you need to make sure you have Go installed and working, then install all the dependencies:

```bash
go mod download
```

Once you have all the dependencies installed (should only take a few seconds) you can run the program easily using:

```bash
go run cmd/web/main.go
```

This will by default run the program and make it accessible through port `8080` allowing you to visit: `http://localhost:8080` and use the API.

