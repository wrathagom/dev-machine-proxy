package terminal

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sync"

	"github.com/creack/pty"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for local dev use
	},
}

// Handler handles WebSocket terminal connections
type Handler struct{}

// NewHandler creates a new terminal handler
func NewHandler() *Handler {
	return &Handler{}
}

// ServeWS handles WebSocket connections for terminal sessions
func (h *Handler) ServeWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}

	// Get user's default shell
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "/bin/bash"
	}

	// Start a new shell process with PTY
	cmd := exec.Command(shell)
	cmd.Env = append(os.Environ(), "TERM=xterm-256color")

	ptmx, err := pty.Start(cmd)
	if err != nil {
		log.Printf("PTY start error: %v", err)
		conn.WriteMessage(websocket.TextMessage, []byte("Failed to start terminal: "+err.Error()))
		conn.Close()
		return
	}

	var once sync.Once
	cleanup := func() {
		once.Do(func() {
			cmd.Process.Kill()
			ptmx.Close()
			conn.Close()
		})
	}
	defer cleanup()

	var wg sync.WaitGroup
	wg.Add(2)

	// Read from PTY, write to WebSocket
	go func() {
		defer wg.Done()
		defer cleanup()
		buf := make([]byte, 4096)
		for {
			n, err := ptmx.Read(buf)
			if err != nil {
				if err != io.EOF {
					log.Printf("PTY read error: %v", err)
				}
				return
			}
			if err := conn.WriteMessage(websocket.BinaryMessage, buf[:n]); err != nil {
				return
			}
		}
	}()

	// Read from WebSocket, write to PTY
	go func() {
		defer wg.Done()
		defer cleanup()
		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			// Handle resize messages (JSON starting with {"cols":)
			if msgType == websocket.TextMessage && len(msg) > 0 && msg[0] == '{' {
				h.handleResize(ptmx, msg)
				continue
			}

			if _, err := ptmx.Write(msg); err != nil {
				return
			}
		}
	}()

	// Wait for shell to exit
	cmd.Wait()
	cleanup()
	wg.Wait()
}

// handleResize handles terminal resize messages
func (h *Handler) handleResize(ptmx *os.File, msg []byte) {
	// Parse simple JSON: {"cols":80,"rows":24}
	var cols, rows uint16
	_, err := parseResizeMsg(msg, &cols, &rows)
	if err != nil {
		return
	}

	if cols > 0 && rows > 0 {
		pty.Setsize(ptmx, &pty.Winsize{
			Cols: cols,
			Rows: rows,
		})
	}
}

// parseResizeMsg parses a resize JSON message
func parseResizeMsg(msg []byte, cols, rows *uint16) (bool, error) {
	// Simple parsing without importing encoding/json for this small message
	// Expected format: {"cols":80,"rows":24}
	var c, r int
	n, _ := sscanf(string(msg), `{"cols":%d,"rows":%d}`, &c, &r)
	if n == 2 {
		*cols = uint16(c)
		*rows = uint16(r)
		return true, nil
	}
	return false, nil
}

// sscanf is a simple scanf-like function
func sscanf(s, format string, args ...interface{}) (int, error) {
	var matched int
	var colsPtr, rowsPtr *int

	if len(args) >= 1 {
		colsPtr = args[0].(*int)
	}
	if len(args) >= 2 {
		rowsPtr = args[1].(*int)
	}

	// Find cols value
	colsStart := indexOf(s, `"cols":`)
	if colsStart >= 0 {
		colsStart += 7
		colsEnd := colsStart
		for colsEnd < len(s) && s[colsEnd] >= '0' && s[colsEnd] <= '9' {
			colsEnd++
		}
		if colsEnd > colsStart && colsPtr != nil {
			*colsPtr = atoi(s[colsStart:colsEnd])
			matched++
		}
	}

	// Find rows value
	rowsStart := indexOf(s, `"rows":`)
	if rowsStart >= 0 {
		rowsStart += 7
		rowsEnd := rowsStart
		for rowsEnd < len(s) && s[rowsEnd] >= '0' && s[rowsEnd] <= '9' {
			rowsEnd++
		}
		if rowsEnd > rowsStart && rowsPtr != nil {
			*rowsPtr = atoi(s[rowsStart:rowsEnd])
			matched++
		}
	}

	return matched, nil
}

func indexOf(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

func atoi(s string) int {
	var n int
	for _, c := range s {
		n = n*10 + int(c-'0')
	}
	return n
}
