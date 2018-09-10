#
# Makefile
# @author Lempiy Anton <lempiyada@gmail.com>
#

GOVERSION = 1.11

# Parser def
PARSER_NAME = dora_parser
PARSER_VER = v0.0.1

# Bot def
BOT_NAME = dora_bot
BOT_VER = v0.0.1

# API def
API_NAME = dora_API
API_VER = v0.0.1

MINIKUBE_STOPPED = $(shell minikube status | grep -o Stopped)

ifeq ($(SERVICE),parser)
	name=$(PARSER_NAME)
endif
ifeq ($(SERVICE),bot)
	name=$(BOT_NAME)
endif
ifeq ($(SERVICE),api)
	name=$(API_NAME)
endif

.PHONY: target

target:
ifndef name
	@echo 'Please provide SERVICE=name (possible variants: parser, bot, api)'
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
	@echo 'Please provide SERVICE=name (posible variants: parser, bot, api)'
	@exit 1
endif
ifeq ($(MINIKUBE_STOPPED), Stopped)
	@echo minikube is down. Running minikube ...
	@minikube start
endif
	@eval $(shell minikube docker-env)
	services/$(SERVICE)/dev.sh
	@exit 0

k8s-show-url:
ifeq ($(MINIKUBE_STOPPED), Stopped)
	@echo minikube is down. Running minikube ...
	@minikube start
endif
	@minikube service dora-api-service --url

k8s-clear-dev:
ifeq ($(MINIKUBE_STOPPED), Stopped)
	@echo minikube is down. Running minikube ...
	@minikube start
endif
	@kubectl delete -f k8s/dora-dev.yaml

k8s-create-dev:
ifeq ($(MINIKUBE_STOPPED), Stopped)
	@echo minikube is down. Running minikube ...
	@minikube start
endif
	@eval $(minikube docker-env)
	@kubectl create -f k8s/dora-dev.yaml
	
k8s-show-pods:
	@kubectl get pods

k8s-dev-down:
	@minikube stop
