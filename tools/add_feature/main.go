package main

import (
	"errors"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type featureNames struct {
	Raw      string
	Snake    string
	Pascal   string
	Camel    string
	Plural   string
	PathSlug string
}

type scaffoldResult struct {
	Created []string
	Skipped []string
	Updated []string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: go run ./tools/add_feature <feature_name>")
		os.Exit(1)
	}

	names, err := normalizeFeatureName(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid feature name: %v\n", err)
		os.Exit(1)
	}

	result, err := scaffoldFeature(names)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to scaffold feature: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Feature scaffold completed")
	fmt.Printf("Feature: %s\n", names.Snake)

	if len(result.Created) > 0 {
		sort.Strings(result.Created)
		fmt.Println("Created files:")
		for _, file := range result.Created {
			fmt.Printf("  - %s\n", file)
		}
	}

	if len(result.Updated) > 0 {
		sort.Strings(result.Updated)
		fmt.Println("Updated files:")
		for _, file := range result.Updated {
			fmt.Printf("  - %s\n", file)
		}
	}

	if len(result.Skipped) > 0 {
		sort.Strings(result.Skipped)
		fmt.Println("Skipped existing files:")
		for _, file := range result.Skipped {
			fmt.Printf("  - %s\n", file)
		}
	}
}

func scaffoldFeature(names featureNames) (*scaffoldResult, error) {
	result := &scaffoldResult{}

	files := map[string]string{
		filepath.FromSlash("domain/entities/" + names.Snake + ".go"):                                 entityTemplate(names),
		filepath.FromSlash("domain/ports/" + names.Snake + "_repository.go"):                         portTemplate(names),
		filepath.FromSlash("application/services/" + names.Snake + "_service.go"):                    serviceTemplate(names),
		filepath.FromSlash("infrastructure/repositories/" + names.Snake + "_repository_inmemory.go"): repositoryTemplate(names),
		filepath.FromSlash("presentation/handlers/" + names.Snake + "_handler.go"):                   handlerTemplate(names),
		filepath.FromSlash("presentation/routes/" + names.Snake + "_routes.go"):                      routesTemplate(names),
	}

	for file, content := range files {
		created, err := writeGoFileIfMissing(file, content)
		if err != nil {
			return nil, err
		}
		if created {
			result.Created = append(result.Created, filepath.ToSlash(file))
		} else {
			result.Skipped = append(result.Skipped, filepath.ToSlash(file))
		}
	}

	containerUpdated, err := wireContainer(names)
	if err != nil {
		return nil, err
	}
	if containerUpdated {
		result.Updated = append(result.Updated, "presentation/container/container.go")
	}

	routerUpdated, err := wireRouter(names)
	if err != nil {
		return nil, err
	}
	if routerUpdated {
		result.Updated = append(result.Updated, "presentation/routes/router.go")
	}

	return result, nil
}

func writeGoFileIfMissing(path string, content string) (bool, error) {
	if _, err := os.Stat(path); err == nil {
		return false, nil
	} else if !errors.Is(err, os.ErrNotExist) {
		return false, err
	}

	formatted, err := format.Source([]byte(content))
	if err != nil {
		return false, fmt.Errorf("format %s: %w", path, err)
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return false, err
	}

	if err := os.WriteFile(path, formatted, 0o644); err != nil {
		return false, err
	}

	return true, nil
}

func wireContainer(names featureNames) (bool, error) {
	path := filepath.FromSlash("presentation/container/container.go")
	contentBytes, err := os.ReadFile(path)
	if err != nil {
		return false, err
	}
	content := string(contentBytes)

	fieldToken := fmt.Sprintf("\t%sRepository ports.%sRepository", names.Camel, names.Pascal)
	if strings.Contains(content, fieldToken) {
		return false, nil
	}

	structNeedle := "\tproductHandler    *handlers.ProductHandler\n}"
	structInsert := fmt.Sprintf("\tproductHandler    *handlers.ProductHandler\n\t%sRepository ports.%sRepository\n\t%sService    services.%sService\n\t%sHandler    *handlers.%sHandler\n}", names.Camel, names.Pascal, names.Camel, names.Pascal, names.Camel, names.Pascal)
	if !strings.Contains(content, structNeedle) {
		return false, fmt.Errorf("container struct marker not found")
	}
	content = strings.Replace(content, structNeedle, structInsert, 1)

	methods := fmt.Sprintf(`
func (c *Container) Resolve%[1]sRepository() ports.%[1]sRepository {
	if c.%[2]sRepository == nil {
		c.%[2]sRepository = repositories.NewInMemory%[1]sRepository()
	}
	return c.%[2]sRepository
}

func (c *Container) Resolve%[1]sService() services.%[1]sService {
	if c.%[2]sService == nil {
		c.%[2]sService = services.New%[1]sService(c.Resolve%[1]sRepository())
	}
	return c.%[2]sService
}

func (c *Container) Resolve%[1]sHandler() *handlers.%[1]sHandler {
	if c.%[2]sHandler == nil {
		c.%[2]sHandler = handlers.New%[1]sHandler(c)
	}
	return c.%[2]sHandler
}
`, names.Pascal, names.Camel)

	content += methods

	formatted, err := format.Source([]byte(content))
	if err != nil {
		return false, fmt.Errorf("format container wiring: %w", err)
	}

	if err := os.WriteFile(path, formatted, 0o644); err != nil {
		return false, err
	}

	return true, nil
}

func wireRouter(names featureNames) (bool, error) {
	path := filepath.FromSlash("presentation/routes/router.go")
	contentBytes, err := os.ReadFile(path)
	if err != nil {
		return false, err
	}
	content := string(contentBytes)

	handlerLine := fmt.Sprintf("\t%sHandler := appContainer.Resolve%sHandler()", names.Camel, names.Pascal)
	registerLine := fmt.Sprintf("\t\tRegister%sRoutes(v1, %sHandler)", names.Pascal, names.Camel)
	if strings.Contains(content, handlerLine) || strings.Contains(content, registerLine) {
		return false, nil
	}

	productNeedle := "\tproductHandler := appContainer.ResolveProductHandler()"
	if !strings.Contains(content, productNeedle) {
		return false, fmt.Errorf("router product handler marker not found")
	}
	content = strings.Replace(content, productNeedle, productNeedle+"\n"+handlerLine, 1)

	v1Needle := "\tv1 := engine.Group(\"/api/v1\")\n\t{"
	if !strings.Contains(content, v1Needle) {
		return false, fmt.Errorf("router v1 marker not found")
	}
	content = strings.Replace(content, v1Needle, v1Needle+"\n"+registerLine, 1)

	formatted, err := format.Source([]byte(content))
	if err != nil {
		return false, fmt.Errorf("format router wiring: %w", err)
	}

	if err := os.WriteFile(path, formatted, 0o644); err != nil {
		return false, err
	}

	return true, nil
}

func normalizeFeatureName(input string) (featureNames, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return featureNames{}, errors.New("feature name is required")
	}

	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	cleaned := re.ReplaceAllString(trimmed, "_")
	cleaned = strings.Trim(cleaned, "_")
	if cleaned == "" {
		return featureNames{}, errors.New("feature name must contain letters or numbers")
	}

	partsRaw := strings.Split(cleaned, "_")
	parts := make([]string, 0, len(partsRaw))
	for _, part := range partsRaw {
		if part == "" {
			continue
		}
		parts = append(parts, strings.ToLower(part))
	}
	if len(parts) == 0 {
		return featureNames{}, errors.New("feature name is invalid")
	}

	snake := strings.Join(parts, "_")
	pascalParts := make([]string, 0, len(parts))
	for _, part := range parts {
		pascalParts = append(pascalParts, strings.ToUpper(part[:1])+part[1:])
	}
	pascal := strings.Join(pascalParts, "")
	camel := strings.ToLower(pascal[:1]) + pascal[1:]
	plural := pluralize(parts[len(parts)-1])
	pathSlug := strings.ReplaceAll(plural, "_", "-")

	return featureNames{
		Raw:      trimmed,
		Snake:    snake,
		Pascal:   pascal,
		Camel:    camel,
		Plural:   plural,
		PathSlug: pathSlug,
	}, nil
}

func pluralize(word string) string {
	if strings.HasSuffix(word, "s") || strings.HasSuffix(word, "x") || strings.HasSuffix(word, "z") || strings.HasSuffix(word, "ch") || strings.HasSuffix(word, "sh") {
		return word + "es"
	}
	if strings.HasSuffix(word, "y") && len(word) > 1 {
		before := word[len(word)-2]
		if !strings.ContainsRune("aeiou", rune(before)) {
			return word[:len(word)-1] + "ies"
		}
	}
	return word + "s"
}

func entityTemplate(names featureNames) string {
	return fmt.Sprintf(`package entities

import "time"

type %s struct {
	ID        string    `+"`json:\"id\"`"+`
	Name      string    `+"`json:\"name\"`"+`
	CreatedAt time.Time `+"`json:\"created_at\"`"+`
	UpdatedAt time.Time `+"`json:\"updated_at\"`"+`
}

func (e *%s) Validate() error {
	// TODO: Add business validation rules for %s.
	return nil
}
`, names.Pascal, names.Pascal, names.Pascal)
}

func portTemplate(names featureNames) string {
	return fmt.Sprintf(`package ports

import (
	"context"

	"template-go-hexagonal/domain/entities"
)

type %sRepository interface {
	List(ctx context.Context) ([]entities.%s, error)
}
`, names.Pascal, names.Pascal)
}

func serviceTemplate(names featureNames) string {
	return fmt.Sprintf(`package services

import (
	"context"
	"log"

	"template-go-hexagonal/domain/entities"
	"template-go-hexagonal/domain/ports"
)

type %[1]sService interface {
	List(ctx context.Context) ([]entities.%[1]s, error)
}

type %[1]sServiceImpl struct {
	repository ports.%[1]sRepository
}

func New%[1]sService(repository ports.%[1]sRepository) %[1]sService {
	return &%[1]sServiceImpl{repository: repository}
}

func (s *%[1]sServiceImpl) List(ctx context.Context) ([]entities.%[1]s, error) {
	items, err := s.repository.List(ctx)
	if err != nil {
		log.Printf("%[2]s", err)
		return nil, err
	}

	return items, nil
}
`, names.Pascal, "["+strings.ToUpper(names.Snake)+"_SERVICE] operation=List error=%v")
}

func repositoryTemplate(names featureNames) string {
	return fmt.Sprintf(`package repositories

import (
	"context"
	"sync"

	"template-go-hexagonal/domain/entities"
)

type InMemory%sRepository struct {
	mu    sync.RWMutex
	items []entities.%s
}

func NewInMemory%sRepository() *InMemory%sRepository {
	return &InMemory%sRepository{items: make([]entities.%s, 0)}
}

func (r *InMemory%sRepository) List(_ context.Context) ([]entities.%s, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	copied := make([]entities.%s, len(r.items))
	copy(copied, r.items)
	return copied, nil
}
`, names.Pascal, names.Pascal, names.Pascal, names.Pascal, names.Pascal, names.Pascal, names.Pascal, names.Pascal, names.Pascal)
}

func handlerTemplate(names featureNames) string {
	return fmt.Sprintf(`package handlers

import (
	"net/http"

	"template-go-hexagonal/application/services"
	presentationerrors "template-go-hexagonal/presentation/errors"
	"template-go-hexagonal/presentation/logging"
	"template-go-hexagonal/presentation/responses"

	"github.com/gin-gonic/gin"
)

type %[1]sServiceResolver interface {
	Resolve%[1]sService() services.%[1]sService
}

type %[1]sHandler struct {
	serviceResolver %[1]sServiceResolver
}

func New%[1]sHandler(serviceResolver %[1]sServiceResolver) *%[1]sHandler {
	return &%[1]sHandler{serviceResolver: serviceResolver}
}

func (h *%[1]sHandler) List(c *gin.Context) {
	items, err := h.serviceResolver.Resolve%[1]sService().List(c.Request.Context())
	if err != nil {
		h.handleError(c, "List", err)
		return
	}

	responses.Success(c, http.StatusOK, items)
}

func (h *%[1]sHandler) handleError(c *gin.Context, operation string, err error) {
	logging.LogError("%[2]s_HANDLER", operation, err)
	responses.ErrorFromApp(c, presentationerrors.Map(err))
}
`, names.Pascal, strings.ToUpper(names.Snake))
}

func routesTemplate(names featureNames) string {
	return fmt.Sprintf(`package routes

import (
	"template-go-hexagonal/presentation/handlers"

	"github.com/gin-gonic/gin"
)

func Register%sRoutes(v1 *gin.RouterGroup, handler *handlers.%sHandler) {
	resources := v1.Group("/%s")
	{
		resources.GET("", handler.List)
	}
}
`, names.Pascal, names.Pascal, names.PathSlug)
}
