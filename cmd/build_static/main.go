package main

import (
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
	"landing-page/views"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables if .env exists
	_ = godotenv.Load()

	// Clean and recreate dist directory
	_ = os.RemoveAll("dist")
	if err := os.MkdirAll("dist", 0755); err != nil {
		log.Fatalf("Failed to create dist directory: %v", err)
	}

	// Create output index.html file
	f, err := os.Create("dist/index.html")
	if err != nil {
		log.Fatalf("Failed to create index.html: %v", err)
	}
	defer f.Close()

	// Render the Templ Index component directly to index.html file
	ctx := context.Background()
	if err := views.Index().Render(ctx, f); err != nil {
		log.Fatalf("Failed to render index template: %v", err)
	}

	// Copy public folder resources to dist/static
	if err := copyDir("public", "dist/static"); err != nil {
		log.Fatalf("Failed to copy static files: %v", err)
	}

	log.Println("Static HTML and assets build completed successfully inside dist/")
}

// copyDir recursively copies a directory tree
func copyDir(src string, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)
		if info.IsDir() {
			return os.MkdirAll(target, info.Mode())
		}
		return copyFile(path, target)
	})
}

// copyFile copies a single file from src to dst
func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	if _, err = io.Copy(out, in); err != nil {
		return err
	}
	return out.Sync()
}
