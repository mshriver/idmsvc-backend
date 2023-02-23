##
# Rules related with the generation of plantuml diagrams.
#
# NOTE: Keep in mind that they don't need to be added to the
#       repository as it can be seen at the link below:
#       https://blog.anoff.io/2018-07-31-diagrams-with-plantuml/
#
# NOTE: You need to install plantuml by hand:
#       In Fedora systems you can do that by:
#       # dnf install -y plantuml
##

PLANTER=$(BIN)/planter

.PHONY: generate-diagrams
generate-diagrams: PLANTUML ?= $(shell command -v plantuml 2>/dev/null)
generate-diagrams: PLANTUML ?= false
generate-diagrams: $(patsubst docs/%.puml,docs/%.svg,$(wildcard docs/*.puml)) ## Generate diagrams

.PHONY: docs/db-model.puml
docs/db-model.puml: $(PLANTER)
	$(PLANTER) postgres://$(DATABASE_USER):$(DATABASE_PASSWORD)@$(DATABASE_HOST)/$(DATABASE_NAME)?sslmode=disable -o $@

.PHONY: install-planter
install-planter: $(PLANTER)

$(PLANTER):
	@{\
		export GOPATH="$(shell mktemp -d "$(PROJECT_DIR)/tmp.XXXXXXXX" 2>/dev/null)" ; \
		echo "Using GOPATH='$${GOPATH}'" ; \
		[ "$${GOPATH}" != "" ] || { echo "error:GOPATH is empty"; exit 1; } ; \
		export GOBIN="$(dir $(PLANTER))" ; \
		echo "Installing 'planter' at '$(PLANTER)'" ; \
		go install github.com/achiku/planter@latest ; \
		find "$${GOPATH}" -type d -exec chmod u+w {} \; ; \
		rm -rf "$${GOPATH}" ; \
	}

# General rule to generate a diagram in SVG format for
# each .puml file found at docs/ directory
docs/%.svg: docs/%.puml
	$(PLANTUML) -tsvg $<
