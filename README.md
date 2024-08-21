# k8s-keepalive
![Go Test Workflow](https://github.com/mike-callahan/k8s-keepalive/actions/workflows/go.yml/badge.svg)
![Docker Build Workflow](https://github.com/mike-callahan/k8s-keepalive/actions/workflows/docker-image.yml/badge.svg)
## Introduction
Over time I've found myself needing a way to keep a k8s pod alive either for debugging purposes or for interactive use. 

The typical recommendation on the internet is to use `tail -f /dev/null` or `/usr/bin/sleep infinity`.

I've found a few issues with these two approaches:
1. Both of those commands assume you have `tail` or `sleep` available in your image.
2. If you don't have tail or sleep you *probably* need a package manager to install them.
3. Unless you're careful, a package manager means more than just a single binary ending up in the image layer.
4. Depending on configuration, `tail` and `sleep` dont respect `SIGINT` making killing the container annoying.
5. There is no way robust way to configure Liveliness, Startup, or Readiness probes with `tail` and `sleep`.

This binary attempts to solve all of the aforementioned challenges. The binary is written in golang and is statically linked, making it compatible with a wide range of environments.

## Additional features

For health checking you can point to `:5000/` or `:5000/healthz` to recieve a 200 OK. You can also type any valid HTTP status code after the `/` to recieve that status code back.

Example: `:5000/404` will send back a 404 in the response header and body.

## Usage
### Installation options
#### Run
To run it standalone:

`docker run -p 8080:5000 mikecallahan/k8s-keepalive:1.0.0`

#### Copy the binary
Download the latest release, `COPY` it into your Dockerfile, and use it in your `ENTRYPOINT [""]` or `CMD [""]` when you need it.

#### Public image
Use the public image in a multi-stage build:
```
FROM mikecallahan/k8s-keepalive:1.0.0 as build

--your docker file starts here ---
FROM ...
COPY --from=build /usr/bin/k8s-keepalive /usr/bin/k8s-keepalive
```

#### Update a yaml file
If you have an existing image already deployed in k8s, and your image has `wget` or `curl`, update your deployment.yaml:

```
spec:
  containers:
  - name: my-existing-container
    image: ubuntu:latest
    command: ["/bin/sh"]
    args: ["-c", "curl -OL https://github.com/mike-callahan/k8s-keepalive/releases/download/v1.0.0/k8s-keepalive && ./k8s-keepalive"]
```
## Roadmap
- [x] Keep container running
- [x] Support health checks
- [x] Statically link binary (net package)
- [x] Release Dockerfile and image
- [x] Support native `SIGINT` and `SIGTERM`
- [ ] Cross-compile for Windows
