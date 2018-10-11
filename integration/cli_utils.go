package integration

import (
	"bytes"
	"os/exec"
	"testing"
)

// InvokeCLISuccess ensure cli invocation is successful
func InvokeCLISuccess(t *testing.T, cmd *exec.Cmd) (*bytes.Buffer, *bytes.Buffer) {
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		t.Fatalf("expected to succeed but failed, %s, %s, %s", err.Error(), stdout.String(), stdout.String())
	}

	t.Log(stdout.String())
	t.Log(stderr.String())
	return &stdout, &stderr
}

// InvokeCLIFailure expect failure from invocation and return failure
func InvokeCLIFailure(t *testing.T, cmd *exec.Cmd) (*bytes.Buffer, *bytes.Buffer, error) {
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		t.Log(err)
		t.Log(stdout.String())
		t.Log(stderr.String())
		return &stdout, &stderr, err
	}

	t.Fatalf("expected to fail but succeeded, %s, %s", stdout.String(), stderr.String())
	return nil, nil, nil
}
