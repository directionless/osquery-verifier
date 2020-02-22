VERSION = $(shell git describe --tags --always --dirty)
BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
REVISION = $(shell git rev-parse HEAD)
REVSHORT = $(shell git rev-parse --short HEAD)
USER = $(shell whoami)

KIT_VERSION = "\
	-X github.com/kolide/kit/version.appName=${APP_NAME} \
	-X github.com/kolide/kit/version.version=${VERSION} \
	-X github.com/kolide/kit/version.branch=${BRANCH} \
	-X github.com/kolide/kit/version.revision=${REVISION} \
	-X github.com/kolide/kit/version.buildDate=${NOW} \
	-X github.com/kolide/kit/version.buildUser=${USER} \
	-X github.com/kolide/kit/version.goVersion=${GOVERSION}"

export GO111MODULE=on

.PHONY: all
all: build/ov

build/%: APP_NAME = $*
build/%: cmd/%
	@mkdir -p build
	go build -o $@ -ldflags ${KIT_VERSION} ./$<



launch-osqueryd: export ENROLL_SECRET=secret
launch-osqueryd:
	osqueryd --verbose --ephemeral --disable_database \
	--tls_hostname localhost:4433 \
	--tls_server_certs build/server.crt \
	--config_plugin tls \
	--enroll_tls_endpoint /enroll \
	--logger_plugin tls \
	--logger_tls_endpoint /log \
	--config_tls_endpoint /config \
	--enroll_secret_env ENROLL_SECRET


build/server.crt:
	openssl req -x509 -subj '/CN=localhost' \
	  -newkey rsa:2048  -nodes -days 365 \
	  -keyout build/server.key -out build/server.crt
