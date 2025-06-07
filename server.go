package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

type ExecRequest struct {
	Cmd string `json:"cmd"`
}

type ExecResponse struct {
	Success bool   `json:"success"`
	Output  string `json:"output"`
	Error   string `json:"error,omitempty"`
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
	out, err := cmd.CombinedOutput()
	resp := ExecResponse{
		Output: string(out),
	}
	if err != nil {
		resp.Success = false
		resp.Error = err.Error()
	} else {
		resp.Success = true
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
    headers: {'Content-Type': 'application/json'},
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
		r.URL.Path = "/RiProG.js"
		r.Header.Set("Content-Type", "application/javascript")
		riprogJSHandler(w, r)
		return
	case "/riprog-exec":
		r.URL.Path = "/riprog-exec"
		execHandler(w, r)
		return
	default:
		if strings.HasPrefix(lowerPath, "/riprog-exec") || strings.HasPrefix(lowerPath, "/riprog.js") {
			if lowerPath == "/riprog-exec" {
				execHandler(w, r)
				return
			} else if lowerPath == "/riprog.js" {
				riprogJSHandler(w, r)
				return
			}
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
