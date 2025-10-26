package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/jhaksh-24/Lazy-Foundry/internal/anvil"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Server struct {
	anvilCmd  *exec.Cmd
	mu        sync.Mutex
	wsClients map[*websocket.Conn]bool
	wsMu      sync.Mutex
}

type Request struct {
	Mode    string   `json:"mode"`
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Output  string `json:"output,omitempty"`
}

type StreamMessage struct {
	Type    string `json:"type"`
	Content string `json:"content"`
	Command string `json:"command"`
}

func New() *Server {
	return &Server{
		wsClients: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) Start(port string) error {
	anvil.Initializer()

	fs := http.FileServer(http.Dir("./Frontend"))
	http.Handle("/", fs)

	http.HandleFunc("/api/execute", s.handleExecute)
	http.HandleFunc("/ws", s.handleWebSocket)

	fmt.Printf("🚀 Lazy-Foundry Web Server starting on http://localhost:%s\n", port)
	return http.ListenAndServe(":"+port, nil)
}

func (s *Server) handleExecute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: "Method not allowed",
		})
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: "Invalid request format",
		})
		return
	}

	response := s.executeCommand(req)
	json.NewEncoder(w).Encode(response)
}

func (s *Server) executeCommand(req Request) Response {
	var err error
	var output string

	switch req.Mode {
	case "forge":
		output, err = s.handleForge(req.Command, req.Args)
	case "anvil":
		output, err = s.handleAnvil(req.Command, req.Args)
	default:
		return Response{
			Success: false,
			Message: fmt.Sprintf("Unknown mode: %s", req.Mode),
		}
	}

	if err != nil {
		return Response{
			Success: false,
			Message: err.Error(),
			Output:  output,
		}
	}

	return Response{
		Success: true,
		Message: "Command executed successfully",
		Output:  output,
	}
}

func (s *Server) handleForge(command string, args []string) (string, error) {
	switch command {
	case "build":
		return s.runCommandWithOutput("forge", append([]string{"build"}, args...)...)
	case "test":
		return s.runCommandWithOutput("forge", append([]string{"test"}, args...)...)
	case "coverage":
		return s.runCommandWithOutput("forge", append([]string{"coverage"}, args...)...)
	case "init":
		return s.runCommandWithOutput("forge", append([]string{"init"}, args...)...)
	case "install":
		if len(args) < 1 {
			return "", fmt.Errorf("install requires package name")
		}
		return s.runCommandWithOutput("forge", append([]string{"install"}, args...)...)
	case "script":
		if len(args) < 1 {
			return "", fmt.Errorf("script requires script path")
		}
		return s.runCommandWithOutput("forge", append([]string{"script"}, args...)...)
	case "create":
		if len(args) < 1 {
			return "", fmt.Errorf("create requires contract name")
		}
		return s.runCommandWithOutput("forge", append([]string{"create"}, args...)...)
	default:
		return "", fmt.Errorf("unknown forge command: %s", command)
	}
}

func (s *Server) handleAnvil(command string, args []string) (string, error) {
	switch command {
	case "add":
		if len(args) < 3 {
			return "", fmt.Errorf("usage: anvil add <name> <rpc-url> <chain-id> [fork-url] [private-key]")
		}

		name := args[0]
		anvil.Initializer()
		anvil.ImplementRpcURL(args[1])
		anvil.ImplementChainID(args[2])

		if len(args) > 3 && args[3] != "" {
			anvil.ImplementForkURL(args[3])
		}

		if len(args) > 4 && args[4] != "" {
			if err := anvil.ImplementPrivateKey(args[4]); err != nil {
				return "", fmt.Errorf("invalid private key: %w", err)
			}
		}

		if err := anvil.SavePreset(name); err != nil {
			return "", fmt.Errorf("failed to save preset: %w", err)
		}

		return fmt.Sprintf("✅ Preset '%s' created successfully!", name), nil

	case "list":
		names := anvil.ListPresets()
		if len(names) == 0 {
			return "No presets found.\nCreate one with: anvil add <name> <rpc-url> <chain-id>", nil
		}

		output := "\n=== Available Presets ===\n\n"
		for _, name := range names {
			preset, err := anvil.GetPreset(name)
			if err != nil {
				continue
			}
			output += fmt.Sprintf("📦 %s\n", name)
			output += fmt.Sprintf("   RPC: %s\n", preset.RpcURL)
			output += fmt.Sprintf("   Chain ID: %d\n", preset.ChainID)
			if preset.ForkURL != "" {
				output += fmt.Sprintf("   Fork: %s\n", preset.ForkURL)
			}
			output += "\n"
		}
		return output, nil

	case "show":
		if len(args) < 1 {
			return "", fmt.Errorf("usage: anvil show <preset-name>")
		}

		name := args[0]
		preset, err := anvil.GetPreset(name)
		if err != nil {
			return "", err
		}

		output := fmt.Sprintf("\n📦 Preset: %s\n", name)
		output += "─────────────────────────────────────────\n"
		output += fmt.Sprintf("  RPC URL:      %s\n", preset.RpcURL)
		output += fmt.Sprintf("  Chain ID:     %d\n", preset.ChainID)
		output += fmt.Sprintf("  Gas Limit:    %d\n", preset.GasLimit)
		output += fmt.Sprintf("  Gas Fee:      %d\n", preset.GasFee)
		output += fmt.Sprintf("  Output Dir:   %s\n", preset.OutputDir)

		if preset.ForkURL != "" {
			output += fmt.Sprintf("  Fork URL:     %s\n", preset.ForkURL)
		}

		if preset.PrivateKey != "" {
			output += fmt.Sprintf("  Private Key:  %s... (hidden)\n", preset.PrivateKey[:10])
		}
		output += "─────────────────────────────────────────\n"

		return output, nil

	case "delete":
		if len(args) < 1 {
			return "", fmt.Errorf("usage: anvil delete <preset-name>")
		}

		name := args[0]
		_, err := anvil.GetPreset(name)
		if err != nil {
			return "", err
		}

		err = anvil.DeletePreset(name)
		if err != nil {
			return "", fmt.Errorf("failed to delete preset: %w", err)
		}

		return fmt.Sprintf("✅ Preset '%s' deleted successfully!", name), nil

	case "start":
		return s.startAnvilNode(args)

	case "stop":
		return s.stopAnvilNode()

	default:
		return "", fmt.Errorf("unknown anvil command: %s", command)
	}
}

func (s *Server) runCommandWithOutput(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		return string(output), fmt.Errorf("command failed: %w", err)
	}

	return string(output), nil
}

func (s *Server) startAnvilNode(args []string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.anvilCmd != nil && s.anvilCmd.Process != nil {
		return "", fmt.Errorf("anvil is already running")
	}

	presetName := "local"
	if len(args) > 0 && args[0] != "" {
		presetName = args[0]
	}

	if err := anvil.LoadPreset(presetName); err != nil {
		return "", fmt.Errorf("failed to load preset '%s': %w", presetName, err)
	}

	config := anvil.GetCurrentConfig()

	anvilArgs := []string{
		"--chain-id", fmt.Sprintf("%d", config.ChainID),
		"--gas-limit", fmt.Sprintf("%d", config.GasLimit),
		"--gas-price", fmt.Sprintf("%d", config.GasFee),
	}

	if config.ForkURL != "" {
		anvilArgs = append(anvilArgs, "--fork-url", config.ForkURL)
	}

	s.anvilCmd = exec.Command("anvil", anvilArgs...)

	stdout, _ := s.anvilCmd.StdoutPipe()
	stderr, _ := s.anvilCmd.StderrPipe()

	if err := s.anvilCmd.Start(); err != nil {
		s.anvilCmd = nil
		return "", fmt.Errorf("failed to start anvil: %w", err)
	}

	go s.streamOutput(stdout, "anvil")
	go s.streamOutput(stderr, "anvil")

	output := fmt.Sprintf("✅ Anvil started with preset '%s'\n", presetName)
	output += fmt.Sprintf("Chain ID: %d\n", config.ChainID)
	output += fmt.Sprintf("RPC URL: %s\n", config.RpcURL)
	if config.ForkURL != "" {
		output += fmt.Sprintf("Forking from: %s\n", config.ForkURL)
	}

	return output, nil
}

func (s *Server) stopAnvilNode() (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.anvilCmd == nil || s.anvilCmd.Process == nil {
		return "", fmt.Errorf("no anvil instance running")
	}

	if err := s.anvilCmd.Process.Kill(); err != nil {
		return "", fmt.Errorf("failed to stop anvil: %w", err)
	}

	s.anvilCmd = nil
	return "✅ Anvil stopped successfully", nil
}

func (s *Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade error:", err)
		return
	}

	s.wsMu.Lock()
	s.wsClients[conn] = true
	s.wsMu.Unlock()

	defer func() {
		s.wsMu.Lock()
		delete(s.wsClients, conn)
		s.wsMu.Unlock()
		conn.Close()
	}()

	for {
		var req Request
		err := conn.ReadJSON(&req)
		if err != nil {
			break
		}

		if req.Mode == "forge" && (req.Command == "test" || req.Command == "coverage" || req.Command == "build") {
			s.executeStreamingCommand(conn, req)
		} else {
			response := s.executeCommand(req)
			conn.WriteJSON(response)
		}
	}
}

func (s *Server) executeStreamingCommand(conn *websocket.Conn, req Request) {
	var cmd *exec.Cmd

	switch req.Command {
	case "build":
		cmd = exec.Command("forge", append([]string{"build"}, req.Args...)...)
	case "test":
		cmd = exec.Command("forge", append([]string{"test"}, req.Args...)...)
	case "coverage":
		cmd = exec.Command("forge", append([]string{"coverage"}, req.Args...)...)
	default:
		conn.WriteJSON(Response{Success: false, Message: "Command not supported for streaming"})
		return
	}

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	if err := cmd.Start(); err != nil {
		conn.WriteJSON(Response{Success: false, Message: err.Error()})
		return
	}

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			conn.WriteJSON(StreamMessage{
				Type:    "output",
				Content: scanner.Text(),
				Command: req.Command,
			})
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			conn.WriteJSON(StreamMessage{
				Type:    "error",
				Content: scanner.Text(),
				Command: req.Command,
			})
		}
	}()

	err := cmd.Wait()

	if err != nil {
		conn.WriteJSON(StreamMessage{
			Type:    "complete",
			Content: "Command failed: " + err.Error(),
			Command: req.Command,
		})
	} else {
		conn.WriteJSON(StreamMessage{
			Type:    "complete",
			Content: "✅ Command completed successfully",
			Command: req.Command,
		})
	}
}

func (s *Server) streamOutput(reader io.Reader, source string) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()

		s.wsMu.Lock()
		for client := range s.wsClients {
			client.WriteJSON(StreamMessage{
				Type:    "output",
				Content: line,
				Command: source,
			})
		}
		s.wsMu.Unlock()
	}
}
