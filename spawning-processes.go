package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	fmt.Println("=== Spawning Processes ===")

	// Simple command execution
	fmt.Println("--- Simple Command ---")
	cmd := exec.Command("echo", "Hello from spawned process")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Output: %s", string(output))

	// Command with arguments
	fmt.Println("\n--- Command with Arguments ---")
	cmd = exec.Command("ls", "-la", "/tmp")
	output, err = cmd.Output()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Directory listing:\n%s", string(output))
	}

	// Running command and capturing both output and error
	fmt.Println("\n--- Capture Output and Error ---")
	cmd = exec.Command("ping", "-c", "3", "localhost")
	
	var stdout, stderr strings.Builder
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	fmt.Printf("Stdout:\n%s\n", stdout.String())
	fmt.Printf("Stderr:\n%s\n", stderr.String())
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Working directory
	fmt.Println("\n--- Custom Working Directory ---")
	cmd = exec.Command("pwd")
	cmd.Dir = "/tmp"
	output, err = cmd.Output()
	if err == nil {
		fmt.Printf("Current directory in /tmp: %s", string(output))
	}

	// Environment variables
	fmt.Println("\n--- Custom Environment ---")
	cmd = exec.Command("env")
	cmd.Env = append(os.Environ(), "CUSTOM_VAR=HelloFromGo")
	output, err = cmd.Output()
	if err == nil {
		// Look for our custom variable
		lines := strings.Split(string(output), "\n")
		for _, line := range lines {
			if strings.Contains(line, "CUSTOM_VAR") {
				fmt.Printf("Found: %s\n", line)
				break
			}
		}
	}

	// Long-running process
	fmt.Println("\n--- Long-running Process ---")
	cmd = exec.Command("sleep", "2")
	start := time.Now()
	err = cmd.Run()
	duration := time.Since(start)
	
	if err == nil {
		fmt.Printf("Process completed in %v\n", duration)
	}

	// Process with input
	fmt.Println("\n--- Process with Input ---")
	cmd = exec.Command("wc", "-c")
	cmd.Stdin = strings.NewReader("Hello, World!")
	output, err = cmd.Output()
	if err == nil {
		fmt.Printf("Character count: %s", string(output))
	}

	// Combined output
	fmt.Println("\n--- Combined Output ---")
	cmd = exec.Command("ls", "/nonexistent")
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Combined output (includes error):\n%s\n", string(output))
		fmt.Printf("Error: %v\n", err)
	}

	// Process information
	fmt.Println("\n--- Process Information ---")
	cmd = exec.Command("sleep", "1")
	
	if err := cmd.Start(); err == nil {
		fmt.Printf("Process started with PID: %d\n", cmd.Process.Pid)
		fmt.Printf("Process state: %s\n", cmd.ProcessState.String())
		
		// Wait for process to complete
		err = cmd.Wait()
		if err == nil {
			fmt.Printf("Process completed successfully\n")
			fmt.Printf("Exit code: %d\n", cmd.ProcessState.ExitCode())
		}
	}

	// Cross-platform commands
	fmt.Println("\n--- Cross-platform Commands ---")
	var cmdName string
	if os.Getenv("OS") == "Windows_NT" {
		cmdName = "echo"
	} else {
		cmdName = "echo"
	}
	
	cmd = exec.Command(cmdName, "Cross-platform hello")
	output, err = cmd.Output()
	if err == nil {
		fmt.Printf("Cross-platform output: %s", string(output))
	}
}
