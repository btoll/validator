SHELL 		= /bin/bash -O globstar
CC      	= go
PROGRAM		= validator
BUILDDIR	= build
CONTEXT		= $(shell kubectl config current-context)

ifneq (,$(findstring development,$(CONTEXT)))
	CLUSTER=development
else ifneq (,$(findstring beta,$(CONTEXT)))
	CLUSTER=beta
else ifneq (,$(findstring production,$(CONTEXT)))
	CLUSTER=production
endif

.PHONY: build clean cleanBuild validate

build: $(PROGRAM)

$(PROGRAM):
	$(CC) build

clean:
	rm -f $(PROGRAM)

cleanBuild:
	rm -rf $(BUILDDIR)

validate: cleanBuild
	for s in $$(kubectl get deploy -ocustom-columns=NAME:.metadata.name --no-headers | sort); do \
		d=$$(ag "name: $$s" -l /home/btoll/projects/veriforce/devops/gitops-test/**/*-deployment.yaml); \
		./$(PROGRAM) --file <(kubectl kustomize $$(echo "$$d" | sed 's/base.*/overlays\/$(CLUSTER)/' 2> /dev/null) 2> /dev/null | yq -o json); \
	done

