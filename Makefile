NAME=norikae
VERSION := $(shell cat main.go | grep "var version" | awk '{print $4}' | sed "s/\"//g")
# Linux、Solaris、Haiku
PREFIX=/usr
# FreeBSDとOpenBSD
#PREFIX=/usr/local
# NetBSD
#PREFIX=/usr/pkg
MANPREFIX=${PREFIX}/share/man
# Linux、OpenBSD、Solaris、Haiku
CNFPREFIX=/etc
# FreeBSD
#CNFPREFIX=/usr/local/etc
# NetBSD
#CNFPREFIX=/usr/pkg/etc
CC=CGO_ENABLED=0 go build
# リリース。なし＝デバッグ。
RELEASE=-ldflags="-s -w" -buildvcs=false

all:
	${CC} ${RELEASE} -o ${NAME}

clean:
	rm -f ${NAME} ${NAME}-${VERSION}.tar.gz

dist: clean
	mkdir -p ${NAME}${VERSION}
	cp -R LICENSE.txt Makefile README.md CHANGELOG.md\
		view static logo.jpg\
		${NAME}.1 *.go *.json ${NAME}-${VERSION}
	tar -zcfv ${NAME}-${VERSION}.tar.gz ${NAME}-${VERSION}
	rm -rf ${NAME}-${VERSION}

install: all
	mkdir -p ${DESTDIR}${PREFIX}/bin
	cp -f ${NAME} ${DESTDIR}${PREFIX}/bin
	chmod 755 ${DESTDIR}${PREFIX}/bin/${NAME}
	mkdir -p ${DESTDIR}${MANPREFIX}/man1
	sed "s/VERSION/${VERSION}/g" < ${NAME}.1 > ${DESTDIR}${MANPREFIX}/man1/${NAME}.1
	chmod 644 ${DESTDIR}${MANPREFIX}/man1/${NAME}.1
	mkdir -p ${DESTDIR}${CNFPREFIX}/${NAME}
	chmod 755 ${DESTDIR}${CNFPREFIX}/${NAME}

uninstall:
	rm -f ${DESTDIOR}${PREFIX}/bin/${NAME}\
		${DESTDIR}${MANPREFIX}/man1/${NAME}.1\
		${DESTDIR}${CNFPREFIX}/${NAME}\
		${DESTDIR}${PREFIX}/share/${NAME}

.PHONY: all options clean dist install uninstall
