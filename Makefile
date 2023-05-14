GOOS_LINUX = linux
# GOOS_WINDOWS = windows
GOOS_MACOS = darwin

GOARCH_LINUX = amd64
# GOARCH_WINDOWS = amd64
GOARCH_MACOS = arm64
MOD_PERM = chmod +x ./builder.sh

install_linux:
	${MOD_PERM}
	env GOOS=${GOOS_LINUX} GOARCH=${GOARCH_LINUX} ./builder.sh

install_macos:
	${MOD_PERM}
	env GOOS=${GOOS_MACOS} GOARCH=${GOARCH_MACOS} ./builder.sh

uninstall:
	chmod +x ./uninstall.sh
	./uninstall.sh
