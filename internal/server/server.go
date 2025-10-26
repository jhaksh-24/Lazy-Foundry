package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/jhaksh-24/Lazy-Foundry/internal/anvil"
	"github.com/jhaksh-24/Lazy-Foundry/internal/forge"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Server struct {
	anvilCmd *exec.Cmd
	mu       sync.Mutex
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

func New() *Server {
	return &Server{}
}

func (s *Server) Start(port string) error {
	// Initialize anvil presets
	anvil.Initializer()

	// Serve static files
	fs := http.FileServer(http.Dir("./Frontend"))
	http.Handle("/", fs)

	// API endpoint
	http.HandleFunc("/api/execute", s.handleExecute)
	
	// WebSocket endpoint for real-time output
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
		return s.captureOutput(func() error {
			return forge.Build(args...)
		})
	case "test":
		return s.captureOutput(func() error {
			return forge.Test(args...)
		})
	case "coverage":
		return s.captureOutput(func() error {
			return forge.Coverage(args...)
		})
	case "init":
		return s.captureOutput(func() error {
			return forge.Init(args...)
		})
	case "install":
		if len(args) < 1 {
			return "", fmt.Errorf("install requires package name")
		}
		return s.captureOutput(func() error {
			return forge.Install(args[0], args[1:]...)
		})
	case "script":
		if len(args) < 1 {
			return "", fmt.Errorf("script requires script path")
		}
		return s.captureOutput(func() error {
			return forge.Script(args[0], args[1:]...)
		})
	case "create":
		if len(args) < 1 {
			return "", fmt.Errorf("create requires contract name")
		}
		return s.captureOutput(func() error {
			return forge.Create(args[0], args[1:]...)
		})
	default:
		return "", fmt.Errorf("unknown forge command: %s", command)
	}
}

func (s *Server) handleAnvil(command string, args []string) (string, error) {
	switch command {
	case "add":
		return s.captureOutput(func() error {
			return anvil.AddPresetCLI(args)
		})
	case "list":
		return s.captureOutput(func() error {
			return anvil.ListPresetsCLI()
		})
	case "show":
		return s.captureOutput(func() error {
			return anvil.ShowPresetCLI(args)
		})
	case "delete":
		return s.captureOutput(func() error {
			return anvil.DeletePresetCLI(args)
		})
	case "start":
		return s.startAnvilNode(args)
	case "stop":
		return s.stopAnvilNode()
	default:
		return "", fmt.Errorf("unknown anvil command: %s", command)
	}
}

func (s *Server) captureOutput(fn func() error) (string, error) {
	// Execute the function and capture any printed output
	// For now, we'll just return error messages
	err := fn()
	if err != nil {
		return "", err
	}
	return "Command executed successfully", nil
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
	
	if err := s.anvilCmd.Start(); err != nil {
		s.anvilCmd = nil
		return "", fmt.Errorf("failed to start anvil: %w", err)
	}

	output := fmt.Sprintf("Anvil started with preset '%s'\nChain ID: %d\nRPC URL: %s",
		presetName, config.ChainID, config.RpcURL)
	
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
	return "Anvil stopped successfully", nil
}

func (s *Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		var req Request
		err := conn.ReadJSON(&req)
		if err != nil {
			break
		}

		response := s.executeCommand(req)
		if err := conn.WriteJSON(response); err != nil {
			break
		}
	}
}
