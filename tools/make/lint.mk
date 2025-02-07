# A wrapper to do lint checks
#
# All make targets related to lint are defined in this file.

##@ Lint

.PHONY: lint
lint: ## Run all linter of code sources, including golint, yamllint.

# lint-deps is run separately in CI to separate the tooling install logs from the actual output logs generated
# by the lint tooling.
.PHONY: lint-deps
lint-deps: ## Everything necessary to lint

.PHONY: lint.golint
lint: lint.golint
lint-deps: $(tools/golangci-lint)
lint.golint: $(tools/golangci-lint)
	@$(LOG_TARGET)
	$(tools/golangci-lint) run $(GOLANGCI_LINT_FLAGS) --build-tags=e2e,celvalidation --config=tools/linter/golangci-lint/.golangci.yml

.PHONY: lint.yamllint
lint: lint.yamllint
lint-deps: $(tools/yamllint)
lint.yamllint: $(tools/yamllint)
	@$(LOG_TARGET)
	$(tools/yamllint) --config-file=tools/linter/yamllint/.yamllint $$(git ls-files :*.yml :*.yaml | xargs -L1 dirname | sort -u)

.PHONY: lint.markdown
lint: lint.markdown
lint-deps: $(tools/markdownlint)
lint.markdown: $(tools/markdownlint)
	@$(LOG_TARGET)
	$(tools/markdownlint) -c tools/linter/markdownlint/markdown_lint_config.json docs/site/content/**

.PHONY: lint.markdown.fix
lint: lint.markdown.fix
lint-deps: $(tools/markdownlint)
lint.markdown.fix: $(tools/markdownlint)
	@$(LOG_TARGET)
	$(tools/markdownlint) -c tools/linter/markdownlint/markdown_lint_config.json --fix docs/site/content/**

.PHONY: lint.codespell
lint: lint.codespell
lint-deps: $(tools/codespell)
lint.codespell: CODESPELL_SKIP := $(shell cat tools/linter/codespell/.codespell.skip | tr \\n ',')
lint.codespell: $(tools/codespell)
	@$(LOG_TARGET)
	$(tools/codespell) --skip $(CODESPELL_SKIP) --ignore-words tools/linter/codespell/.codespell.ignorewords --check-filenames

.PHONY: lint.checklinks
lint: lint.checklinks
lint-deps: $(tools/linkinator)
lint.checklinks: $(tools/linkinator) # Check for broken links in the docs
	@$(LOG_TARGET)
	$(tools/linkinator) docs/site/public/ -r --concurrency 25 --skip $(LINKINATOR_IGNORE)

.PHONY: lint.checklicense
lint: lint.checklicense
lint-deps: $(tools/skywalking-eyes)
lint.checklicense: $(tools/skywalking-eyes) # Check for broken links in the docs
	@$(LOG_TARGET)
	$(tools/skywalking-eyes) -c tools/linter/license/.licenserc.yaml header check
