#
# Makefile
# @author Lempiy Anton <lempiyada@gmail.com>
#

GOVERSION = 1.11beta2

# Parser def
PARSER_NAME = dora_parser
PARSER_VER = v0.0.1

ifeq ($(SERVICE),parser)
	name=$(PARSER_NAME)
endif

.PHONY: target

target:
ifndef name
	@echo 'Please provide SERVICE=name (posible variants: parser)'
	@exit 1
endif
ifndef VERSION
	@echo 'Please provide VERSION=version (in ex. VERSION=v0.4.2)'
	@exit 1
endif
	@docker rmi -f lempiy/$(name):$(VERSION) || true
	services/$(SERVICE)/build.sh $(VERSION)
	@exit 0

.PHONY: dev

dev:
ifndef name
	@echo 'Please provide SERVICE=name (posible variants: parser)'
	@exit 1
endif
	services/$(SERVICE)/dev.sh
	@exit 0
