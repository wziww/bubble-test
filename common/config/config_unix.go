// +build linux freebsd openbsd darwin

package config // import "github.com/docker/docker/client"

// DefaultDockerHost defines os specific default if DOCKER_HOST is unset
const DefaultDockerHost = "unix:///var/run/docker.sock"

const defaultProto = "unix"
const defaultAddr = "/var/run/docker.sock"

// DefaultAPIVersion The default API version used to create a new docker client.
const DefaultAPIVersion = "1.39"
