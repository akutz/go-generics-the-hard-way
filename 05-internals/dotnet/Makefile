#!/usr/bin/env bash

#  Licensed under the Apache License, Version 2.0 (the "License");
#  you may not use this file except in compliance with the License.
#  You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
#  Unless required by applicable law or agreed to in writing, software
#  distributed under the License is distributed on an "AS IS" BASIS,
#  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#  See the License for the specific language governing permissions and
#  limitations under the License.

# If you update this file, please follow
# https://suva.sh/posts/well-documented-makefiles

## --------------------------------------
## General
## --------------------------------------

SHELL:=/usr/bin/env bash
.DEFAULT_GOAL:=help

# Setup some make flags.
ifeq (,$(findstring --no-print-directory,$(MAKEFLAGS)))
MAKEFLAGS += --no-print-directory
endif

# The help will print out all targets with their descriptions organized below
# their categories. The categories are represented by `##@` and the target
# descriptions by `##`.
#
# The awk commands is responsible to read the entire set of makefiles included
# in this invocation, looking for lines of the file as xyz: ## something, and
# then pretty-format the target and help. Then, if there's a line with ##@
# something, that gets pretty-printed as a category.
# 
# More info over the usage of ANSI control characters for terminal
# formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
#
# More info over awk command: http://linuxcommand.org/lc3_adv_awk.php
.PHONY: help
help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

## --------------------------------------
## Dotnet
## --------------------------------------

# Path to the program's source file.
SRC=main.cs

# Path to the built program.
PRG=bin/dotnet

# Builds the program.
#
# The flag -p:UseSharedCompilation=false keeps the dotnet
# from running a persistent build server to speedup
# subsequent compilations. For more information, please see
# https://github.com/dotnet/sdk/issues/9487#issuecomment-396324918.
$(PRG): $(SRC)
	dotnet build --debug -p:UseSharedCompilation=false -o $(@D) $(<D)

build: ## Build the program
	@$(MAKE) $(PRG)

.PHONY: methods
methods: build
methods: ## Print the methods from the program
	lldb \
	  -o "process launch -s" \
	  -o "continue" \
	  -o "continue" \
	  -o "clrstack ." \
	  -o "continue" \
	  -o "clrstack ." \
	  -o "detach" \
	  -o "quit" \
	  $(PRG)

.PHONY: locals
locals: build
locals: ## Print the local variables from the program
	lldb \
	  -o "process launch -s" \
	  -o "continue" \
	  -o "clrstack -i ints" \
	  -o "clrstack -i strs" \
	  -o "detach" \
	  -o "quit" \
	  $(PRG)
