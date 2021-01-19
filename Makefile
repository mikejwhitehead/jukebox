
all: clean update build

update:
	composer update --no-dev

build:
	box build

clean:
	rm -rf build/*
	rm -rf vendor

test:
	composer update --dev
	vendor/bin/phpunit

install-dev:
	wget "https://dl.google.com/go/go1.15.6.linux-armv6l.tar.gz"
	sudo tar -C /usr/local -xzf go1.15.6.linux-armv6l.tar.gz
	rm go1.15.6.linux-armv6l.tar.gz