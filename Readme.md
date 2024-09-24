# build
`docker build -t metrics-demo .`

# run docker from cli
`docker run -d -p 8080:8080 metrics-demo`

# run docker compose
`docker compose up -d`

# run tests
`go test`

# run tests verbosely
`go test -v`

# run docker build cloud
`docker buildx build --builder cloud-ivancurkovic046-cloudbuilder .`
