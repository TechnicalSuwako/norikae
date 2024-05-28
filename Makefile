UNAME_S != uname -s

NAME != cat main.go | grep "var sofname" | awk '{print $$4}' | sed "s/\"//g"
VERSION != cat main.go | grep "var version" | awk '{print $$4}' | sed "s/\"//g"
PREFIX = /usr/local
MANPREFIX = ${PREFIX}/share/man

.if ${UNAME_S} == "OpenBSD"
MANPREFIX = ${PREFIX}/man
.elif ${UNAME_S} == "Linux"
PREFIX = /usr
MANPREFIX = ${PREFIX}/share/man
.endif

CC = CGO_ENABLED=0 go build
RELEASE = -ldflags="-s -w" -buildvcs=false

all:
	${CC} ${RELEASE} -o ${NAME}

release:
	mkdir -p release/bin
	env GOOS=linux 	 GOARCH=amd64   ${CC} ${RELEASE} -o \
		release/bin/${NAME}-${VERSION}-linux-amd64
	env GOOS=linux   GOARCH=arm64   ${CC} ${RELEASE} -o \
		release/bin/${NAME}-${VERSION}-linux-arm64
	env GOOS=linux   GOARCH=riscv64 ${CC} ${RELEASE} -o \
		release/bin/${NAME}-${VERSION}-linux-riscv64
	env GOOS=linux 	 GOARCH=i386    ${CC} ${RELEASE} -o \
		release/bin/${NAME}-${VERSION}-linux-386
	env GOOS=linux   GOARCH=ppc64   ${CC} ${RELEASE} -o \
		release/bin/${NAME}-${VERSION}-linux-ppc64
	env GOOS=linux 	 GOARCH=mips64  ${CC} ${RELEASE} -o \
		release/bin/${NAME}-${VERSION}-linux-mips64
	env GOOS=openbsd GOARCH=amd64   ${CC} ${RELEASE} -o \
		release/bin/${NAME}-${VERSION}-openbsd-amd64
	env GOOS=openbsd GOARCH=i386    ${CC} ${RELEASE} -o \
		release/bin/${NAME}-${VERSION}-openbsd-386
	env GOOS=openbsd GOARCH=arm64   ${CC} ${RELEASE} -o \
		release/bin/${NAME}-${VERSION}-openbsd-arm64
	env GOOS=openbsd GOARCH=mips64  ${CC} ${RELEASE} -o \
		release/bin/${NAME}-${VERSION}-openbsd-mips64
	env GOOS=freebsd GOARCH=amd64   ${CC} ${RELEASE} -o \
		release/bin/${NAME}-${VERSION}-freebsd-amd64
	env GOOS=freebsd GOARCH=i386    ${CC} ${RELEASE} -o \
		release/bin/${NAME}-${VERSION}-freebsd-386
	env GOOS=freebsd GOARCH=arm64   ${CC} ${RELEASE} -o \
		release/bin/${NAME}-${VERSION}-freebsd-arm64
	env GOOS=freebsd GOARCH=riscv64 ${CC} ${RELEASE} -o \
		release/bin/${NAME}-${VERSION}-freebsd-riscv64
	env GOOS=netbsd  GOARCH=amd64   ${CC} ${RELEASE} -o \
		release/bin/${NAME}-${VERSION}-netbsd-amd64
	env GOOS=netbsd  GOARCH=i386    ${CC} ${RELEASE} -o \
		release/bin/${NAME}-${VERSION}-netbsd-386
	env GOOS=netbsd  GOARCH=arm64   ${CC} ${RELEASE} -o \
		release/bin/${NAME}-${VERSION}-netbsd-arm64
	env GOOS=illumos GOARCH=amd64   ${CC} ${RELEASE} -o \
		release/bin/${NAME}-${VERSION}-illumos-amd64

clean:
	rm -f ${NAME} ${NAME}-${VERSION}.tar.gz

dist: clean
	mkdir -p ${NAME}-${VERSION} release/src
	cp -R LICENSE.txt Makefile README.md CHANGELOG.md\
		${NAME}.1 main.go src go.mod go.sum ${NAME}-${VERSION}
	tar zcfv release/src/${NAME}-${VERSION}.tar.gz ${NAME}-${VERSION}
	rm -rf ${NAME}-${VERSION}

install: all
	mkdir -p ${DESTDIR}${PREFIX}/bin
	cp -f ${NAME} ${DESTDIR}${PREFIX}/bin
	chmod 755 ${DESTDIR}${PREFIX}/bin/${NAME}
	mkdir -p ${DESTDIR}${MANPREFIX}/man1
	sed "s/VERSION/${VERSION}/g" < ${NAME}.1 > ${DESTDIR}${MANPREFIX}/man1/${NAME}.1
	chmod 644 ${DESTDIR}${MANPREFIX}/man1/${NAME}.1

uninstall:
	rm -f ${DESTDIR}${PREFIX}/bin/${NAME}\
		${DESTDIR}${MANPREFIX}/man1/${NAME}.1

.PHONY: all release clean dist install uninstall
