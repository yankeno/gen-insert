package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenInsertErrors(t *testing.T) {
	t.Run("file not exists", func(t *testing.T) {
		rootCmd.SetArgs([]string{"not_exists.csv"})

		err := rootCmd.Execute()
		if err == nil {
			t.Fatalf("expected an error but got none")
		}

		if !strings.Contains(err.Error(), "file not exists") {
			t.Fatalf("expected error message to contain %v but got %v", "file not exists", err.Error())
		}
	})

	t.Run("invalid table name", func(t *testing.T) {
		tempDir := t.TempDir()
		csvFile := filepath.Join(tempDir, "test.csv")
		_, err := os.Create(filepath.Clean(csvFile))
		if err != nil {
			t.Fatalf("failed to create test file: %v", err)
		}

		data := []struct {
			tableName string
		}{
			{"123invalid"},
			{"my-table"},
			{"my table"},
		}

		for _, d := range data {
			rootCmd.SetArgs([]string{csvFile, "-t", d.tableName})

			err = rootCmd.Execute()
			if err == nil {
				t.Fatalf("expected an error but got none")
			}

			if !strings.Contains(err.Error(), "table name is invalid") {
				t.Fatalf("expected error message to contain %v but got %v", "table name is invalid", err.Error())
			}
		}
	})

	t.Run("unsupported file format", func(t *testing.T) {
		tempDir := t.TempDir()
		csvFile := filepath.Join(tempDir, "test.unsupported")
		_, err := os.Create(filepath.Clean(csvFile))
		if err != nil {
			t.Fatalf("failed to create test file: %v", err)
		}
		rootCmd.SetArgs([]string{csvFile}) // reset flags with empty strings

		err = rootCmd.Execute()
		if err == nil {
			t.Fatalf("expected an error but got none")
		}

		if !strings.Contains("file format is not supported", "file format is not supported") {
			t.Fatalf("expected error message to contain %v but got %v", "file format is not supported", err.Error())
		}
	})
}

func TestGenInsertOutput(t *testing.T) {
	t.Run("valid CSV input", func(t *testing.T) {
		csvFile := "testdata/sample.csv"
		tempDir := t.TempDir()
		outputFile := filepath.Join(tempDir, "output.sql")
		rootCmd.SetArgs([]string{csvFile, "-o", outputFile, "-t", "sample"})
		err := rootCmd.Execute()

		if err != nil {
			t.Fatalf("command execution failed: %v", err)
		}

		if _, err = os.Stat(outputFile); os.IsNotExist(err) {
			t.Fatalf("expected output file to be created: %v", err)
		}

		outputContent, err := os.ReadFile(filepath.Clean(outputFile))
		if err != nil {
			t.Fatalf("failed to read output file: %v", err)
		}

		expectedContent := "INSERT INTO `sample` (`id`, `name`, `email`) VALUES (1, '佐藤 理紗', 'ayumishiozawa@example.co.jp');\nINSERT INTO `sample` (`id`, `name`, `email`) VALUES (2, '波田野 菜月', NULL);\nINSERT INTO `sample` (`id`, `name`, `email`) VALUES (3, '松田 大樹', 'matsuda1114@example.com');\n"
		if string(outputContent) != expectedContent {
			t.Fatalf("expected output:\n%s\nbut got:\n%s", expectedContent, string(outputContent))
		}
	})

	t.Run("valid TSV input", func(t *testing.T) {
		csvFile := "testdata/sample.tsv"
		tempDir := t.TempDir()
		outputFile := filepath.Join(tempDir, "output.sql")
		rootCmd.SetArgs([]string{csvFile, "-o", outputFile, "-t", "sample"})
		err := rootCmd.Execute()

		if err != nil {
			t.Fatalf("command execution failed: %v", err)
		}

		if _, err = os.Stat(outputFile); os.IsNotExist(err) {
			t.Fatalf("expected output file to be created: %v", err)
		}

		outputContent, err := os.ReadFile(filepath.Clean(outputFile))
		if err != nil {
			t.Fatalf("failed to read output file: %v", err)
		}

		expectedContent := "INSERT INTO `sample` (`id`, `name`, `email`) VALUES (1, '佐藤 理紗', 'ayumishiozawa@example.co.jp');\nINSERT INTO `sample` (`id`, `name`, `email`) VALUES (2, '波田野 菜月', NULL);\nINSERT INTO `sample` (`id`, `name`, `email`) VALUES (3, '松田 大樹', 'matsuda1114@example.com');\n"
		if string(outputContent) != expectedContent {
			t.Fatalf("expected output:\n%s\nbut got:\n%s", expectedContent, string(outputContent))
		}
	})
}
