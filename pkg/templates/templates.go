package templates

import (
	"os"
	"strings"
)

// AssetURL resolves the static asset path.
// If DOMAIN is not defined or is "/", it returns the path relative to root to prevent double-slash host issues.
// Otherwise, it prepends the DOMAIN to the asset path.
func AssetURL(p string) string {
	domain := os.Getenv("DOMAIN")
	if domain == "" || domain == "/" {
		return p
	}
	
	domain = strings.TrimSuffix(domain, "/")
	p = "/" + strings.TrimPrefix(p, "/")
	return domain + p
}
