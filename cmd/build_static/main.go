package main

import (
	"context"
	"io"
	"landing-page/views"
	"log"
	"os"
	"path/filepath"

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

	// Render the Templ Index component directly to index.html file (Vietnamese version)
	f, err := os.Create("dist/index.html")
	if err != nil {
		log.Fatalf("Failed to create index.html: %v", err)
	}
	ctx := context.Background()
	if err := views.Index("vi").Render(ctx, f); err != nil {
		f.Close()
		log.Fatalf("Failed to render index template (vi): %v", err)
	}
	f.Close()

	// Create dist/en directory and render English version to dist/en/index.html
	if err := os.MkdirAll("dist/en", 0755); err != nil {
		log.Fatalf("Failed to create dist/en directory: %v", err)
	}
	fe, err := os.Create("dist/en/index.html")
	if err != nil {
		log.Fatalf("Failed to create dist/en/index.html: %v", err)
	}
	if err := views.Index("en").Render(ctx, fe); err != nil {
		fe.Close()
		log.Fatalf("Failed to render index template (en): %v", err)
	}
	fe.Close()

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
