OUT := server
SRC := server.go

GREEN  := \033[0;32m
YELLOW := \033[1;33m
RESET  := \033[0m

all: clean fmt build

fmt:
	@echo "$(YELLOW)[FMT]$(RESET) Formatting source files..."
	@gofmt -w $(SRC)

build:
	@echo "$(YELLOW)===[ RiProG Compiler ]===$(RESET)"
	@echo "$(YELLOW)Source files: $(RESET)$(SRC)"
	@echo "$(YELLOW)Output file:  $(RESET)$(OUT)"
	@echo "$(YELLOW)Host Arch:    $(RESET)$$(uname -m)"
	@echo "$(GREEN)[BUILD]$(RESET) Compiling $(SRC) -> $(OUT)..."
	@go build -ldflags="-s -w -buildid=" -o $(OUT) $(SRC)
	@echo "$(GREEN)[DONE]$(RESET) Build complete."

clean:
	@echo "$(YELLOW)[CLEAN]$(RESET) Removing $(OUT)..."
	@rm -f $(OUT)
