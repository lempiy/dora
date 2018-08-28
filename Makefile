#
# Makefile
# @author Lempiy Anton <lempiyada@gmail.com>
#

GOVERSION = 1.11beta2

# Parser def
PARSER_NAME = dora_parser
PARSER_VER = v0.0.1

MINIKUBE_STOPPED = $(shell minikube status | grep -o Stopped)

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
	@eval $(minikube docker-env)
	services/$(SERVICE)/dev.sh
	@exit 0

k8s-show-url:
ifeq ($(MINIKUBE_STOPPED), Stopped)
	@echo minikube is down. Running minikube ...
	@minikube start
endif
	@minikube service dora-parser-service --url

k8s-create-dev:
ifeq ($(MINIKUBE_STOPPED), Stopped)
	@echo minikube is down. Running minikube ...
	@minikube start
endif
	@eval $(minikube docker-env)
	@kubectl create -f k8s/dora-dev.yaml
	
k8s-show-pods:
	@kubectl get pods
