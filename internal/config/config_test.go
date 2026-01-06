package config

import (
	"path/filepath"
	"testing"
)

func TestConfig(t *testing.T) {
	t.Run("Reading Config file", func(t *testing.T) {
		setupTestConfig(t)

		initial := Config{
			DBURL:           "postgres://test",
			CurrentUsername: "",
		}

		if err := write(initial); err != nil {
			t.Fatalf("failed to write config: %v", err)
		}

		cfg, err := Read()
		if err != nil {
			t.Fatalf("Read failed: %v", err)
		}

		if cfg.DBURL != initial.DBURL {
			t.Errorf("expected DBURL %s, got %s", initial.DBURL, cfg.DBURL)
		}

		if cfg.CurrentUsername != initial.CurrentUsername {
			t.Errorf(
				"expected CurrentUsername %s, got %s",
				initial.CurrentUsername,
				cfg.CurrentUsername,
			)
		}
	})
	t.Run("Setting new user", func(t *testing.T) {
		setupTestConfig(t)

		cfg := Config{
			DBURL:           "postgres://test",
			CurrentUsername: "",
		}

		err := cfg.SetUser("setuser")
		if err != nil {
			t.Fatalf("SetUser failed: %v", err)
		}

		if cfg.CurrentUsername != "setuser" {
			t.Errorf("expected CurrentUsername 'setuser', got %s", cfg.CurrentUsername)
		}
	})
}

func setupTestConfig(t *testing.T) {
	t.Helper()

	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, configFileName)

	configPathFunc = func() (string, error) {
		return tmpFile, nil
	}

	t.Cleanup(func() {
		configPathFunc = getConfigFilePath
	})
}