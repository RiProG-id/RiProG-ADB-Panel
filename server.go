package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"syscall"
)

type ExecRequest struct {
	Cmd string `json:"cmd"`
}

type ExecResult struct {
	ExitCode int    `json:"errno"`
	StdOut   string `json:"stdout"`
	StdErr   string `json:"stderr"`
}

func execHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}

	var req ExecRequest
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	cmd := exec.Command("sh", "-c", req.Cmd)
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf
	err = cmd.Run()

	exitCode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				exitCode = status.ExitStatus()
			} else {
				exitCode = -1
			}
		} else {
			exitCode = -1
		}
	}

	resp := ExecResult{
		ExitCode: exitCode,
		StdOut:   stdoutBuf.String(),
		StdErr:   stderrBuf.String(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func riprogJSHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	io.WriteString(w, `
export async function exec(cmd) {
  const res = await fetch('/riprog-exec', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ cmd })
  });
  return await res.json();
}
`)
}

type caseInsensitiveMux struct {
	fs http.Handler
}

func (c *caseInsensitiveMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lowerPath := strings.ToLower(r.URL.Path)

	switch lowerPath {
	case "/riprog.js":
		riprogJSHandler(w, r)
	case "/riprog-exec":
		execHandler(w, r)
	default:
		if strings.HasPrefix(lowerPath, "/riprog-exec") {
			execHandler(w, r)
			return
		}
		if strings.HasPrefix(lowerPath, "/riprog.js") {
			riprogJSHandler(w, r)
			return
		}
		c.fs.ServeHTTP(w, r)
	}
}

func main() {
	fs := http.FileServer(http.Dir("webroot"))
	mux := &caseInsensitiveMux{fs: fs}
	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
