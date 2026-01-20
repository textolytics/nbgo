package doc

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// Documentation represents generated documentation
type Documentation struct {
	Title       string
	Description string
	Version     string
	GeneratedAt time.Time
	Sections    []Section
}

// Section represents a documentation section
type Section struct {
	Title       string
	Content     string
	Subsections []SubSection
}

// SubSection represents a documentation subsection
type SubSection struct {
	Title   string
	Content string
}

// Generator generates documentation
type Generator struct {
	mu   sync.RWMutex
	docs *Documentation
}

// NewGenerator creates a new documentation generator
func NewGenerator(title, description, version string) *Generator {
	return &Generator{
		docs: &Documentation{
			Title:       title,
			Description: description,
			Version:     version,
			GeneratedAt: time.Now(),
			Sections:    make([]Section, 0),
		},
	}
}

// AddSection adds a section to the documentation
func (g *Generator) AddSection(section Section) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.docs.Sections = append(g.docs.Sections, section)
}

// GenerateMarkdown generates Markdown documentation
func (g *Generator) GenerateMarkdown() string {
	g.mu.RLock()
	defer g.mu.RUnlock()

	var sb strings.Builder

	// Header
	sb.WriteString(fmt.Sprintf("# %s\n\n", g.docs.Title))
	sb.WriteString(fmt.Sprintf("%s\n\n", g.docs.Description))
	sb.WriteString(fmt.Sprintf("**Version:** %s\n", g.docs.Version))
	sb.WriteString(fmt.Sprintf("**Generated:** %s\n\n", g.docs.GeneratedAt.Format(time.RFC3339)))

	// Table of Contents
	sb.WriteString("## Table of Contents\n\n")
	for i, section := range g.docs.Sections {
		sb.WriteString(fmt.Sprintf("%d. [%s](#%s)\n", i+1, section.Title, strings.ToLower(strings.ReplaceAll(section.Title, " ", "-"))))
	}
	sb.WriteString("\n")

	// Sections
	for _, section := range g.docs.Sections {
		sb.WriteString(fmt.Sprintf("## %s\n\n", section.Title))
		sb.WriteString(fmt.Sprintf("%s\n\n", section.Content))

		for _, subsection := range section.Subsections {
			sb.WriteString(fmt.Sprintf("### %s\n\n", subsection.Title))
			sb.WriteString(fmt.Sprintf("%s\n\n", subsection.Content))
		}
	}

	return sb.String()
}

// GenerateHTML generates HTML documentation
func (g *Generator) GenerateHTML() string {
	g.mu.RLock()
	defer g.mu.RUnlock()

	var sb strings.Builder

	sb.WriteString("<!DOCTYPE html>\n<html>\n<head>\n")
	sb.WriteString(fmt.Sprintf("<title>%s</title>\n", g.docs.Title))
	sb.WriteString("<meta charset=\"UTF-8\">\n")
	sb.WriteString("<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n")
	sb.WriteString("</head>\n<body>\n")

	// Header
	sb.WriteString(fmt.Sprintf("<h1>%s</h1>\n", g.docs.Title))
	sb.WriteString(fmt.Sprintf("<p>%s</p>\n", g.docs.Description))
	sb.WriteString(fmt.Sprintf("<p><strong>Version:</strong> %s</p>\n", g.docs.Version))
	sb.WriteString(fmt.Sprintf("<p><strong>Generated:</strong> %s</p>\n", g.docs.GeneratedAt.Format(time.RFC3339)))

	// Sections
	for _, section := range g.docs.Sections {
		sb.WriteString(fmt.Sprintf("<h2>%s</h2>\n", section.Title))
		sb.WriteString(fmt.Sprintf("<p>%s</p>\n", section.Content))

		for _, subsection := range section.Subsections {
			sb.WriteString(fmt.Sprintf("<h3>%s</h3>\n", subsection.Title))
			sb.WriteString(fmt.Sprintf("<p>%s</p>\n", subsection.Content))
		}
	}

	sb.WriteString("</body>\n</html>\n")
	return sb.String()
}

// APIDocGenerator generates API documentation
type APIDocGenerator struct {
	mu        sync.RWMutex
	endpoints []APIEndpoint
}

// APIEndpoint represents an API endpoint
type APIEndpoint struct {
	Method      string
	Path        string
	Description string
	Parameters  []APIParameter
	Response    APIResponse
}

// APIParameter represents an API parameter
type APIParameter struct {
	Name        string
	Type        string
	Required    bool
	Description string
}

// APIResponse represents an API response
type APIResponse struct {
	StatusCode  int
	Type        string
	Description string
}

// NewAPIDocGenerator creates a new API documentation generator
func NewAPIDocGenerator() *APIDocGenerator {
	return &APIDocGenerator{
		endpoints: make([]APIEndpoint, 0),
	}
}

// AddEndpoint adds an API endpoint to the documentation
func (ag *APIDocGenerator) AddEndpoint(endpoint APIEndpoint) {
	ag.mu.Lock()
	defer ag.mu.Unlock()
	ag.endpoints = append(ag.endpoints, endpoint)
}

// GenerateMarkdown generates Markdown API documentation
func (ag *APIDocGenerator) GenerateMarkdown() string {
	ag.mu.RLock()
	defer ag.mu.RUnlock()

	var sb strings.Builder

	sb.WriteString("# API Documentation\n\n")

	for _, endpoint := range ag.endpoints {
		sb.WriteString(fmt.Sprintf("## %s %s\n\n", endpoint.Method, endpoint.Path))
		sb.WriteString(fmt.Sprintf("%s\n\n", endpoint.Description))

		if len(endpoint.Parameters) > 0 {
			sb.WriteString("### Parameters\n\n")
			for _, param := range endpoint.Parameters {
				required := "Required"
				if !param.Required {
					required = "Optional"
				}
				sb.WriteString(fmt.Sprintf("- `%s` (%s, %s): %s\n", param.Name, param.Type, required, param.Description))
			}
			sb.WriteString("\n")
		}

		sb.WriteString("### Response\n\n")
		sb.WriteString(fmt.Sprintf("- Status: %d\n", endpoint.Response.StatusCode))
		sb.WriteString(fmt.Sprintf("- Type: %s\n", endpoint.Response.Type))
		sb.WriteString(fmt.Sprintf("- Description: %s\n\n", endpoint.Response.Description))
	}

	return sb.String()
}
