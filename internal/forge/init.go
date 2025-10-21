package forge

import (
	"fmt"
)

func Init(flags ...string) error {
	args := {"init"}
	args = append(args, flags...)

	cmd := exec.Command("forge", args...)

	output, err = cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("forge test failed: %w\nOutput: %s", err, string(output))
	}

	fmt.Println("Forge init executed Successfully!")

	if len(output) > 0 {
		fmt.Println(strings.TrimSpace(string(output)))
	}

	return nil
}
