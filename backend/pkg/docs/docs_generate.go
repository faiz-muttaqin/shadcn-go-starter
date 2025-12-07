package docs

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"log"
	"mime"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//go:embed *.png *.css *.html *.js
var swaggerFiles embed.FS

var statusCodes = map[string]string{
	"StatusContinue":                      "100",
	"StatusSwitchingProtocols":            "101",
	"StatusProcessing":                    "102",
	"StatusEarlyHints":                    "103",
	"StatusOK":                            "200",
	"StatusCreated":                       "201",
	"StatusAccepted":                      "202",
	"StatusNonAuthoritativeInfo":          "203",
	"StatusNoContent":                     "204",
	"StatusResetContent":                  "205",
	"StatusPartialContent":                "206",
	"StatusMultiStatus":                   "207",
	"StatusAlreadyReported":               "208",
	"StatusIMUsed":                        "226",
	"StatusMultipleChoices":               "300",
	"StatusMovedPermanently":              "301",
	"StatusFound":                         "302",
	"StatusSeeOther":                      "303",
	"StatusNotModified":                   "304",
	"StatusUseProxy":                      "305",
	"_":                                   "306",
	"StatusTemporaryRedirect":             "307",
	"StatusPermanentRedirect":             "308",
	"StatusBadRequest":                    "400",
	"StatusUnauthorized":                  "401",
	"StatusPaymentRequired":               "402",
	"StatusForbidden":                     "403",
	"StatusNotFound":                      "404",
	"StatusMethodNotAllowed":              "405",
	"StatusNotAcceptable":                 "406",
	"StatusProxyAuthRequired":             "407",
	"StatusRequestTimeout":                "408",
	"StatusConflict":                      "409",
	"StatusGone":                          "410",
	"StatusLengthRequired":                "411",
	"StatusPreconditionFailed":            "412",
	"StatusRequestEntityTooLarge":         "413",
	"StatusRequestURITooLong":             "414",
	"StatusUnsupportedMediaType":          "415",
	"StatusRequestedRangeNotSatisfiable":  "416",
	"StatusExpectationFailed":             "417",
	"StatusTeapot":                        "418",
	"StatusMisdirectedRequest":            "421",
	"StatusUnprocessableEntity":           "422",
	"StatusLocked":                        "423",
	"StatusFailedDependency":              "424",
	"StatusTooEarly":                      "425",
	"StatusUpgradeRequired":               "426",
	"StatusPreconditionRequired":          "428",
	"StatusTooManyRequests":               "429",
	"StatusRequestHeaderFieldsTooLarge":   "431",
	"StatusUnavailableForLegalReasons":    "451",
	"StatusInternalServerError":           "500",
	"StatusNotImplemented":                "501",
	"StatusBadGateway":                    "502",
	"StatusServiceUnavailable":            "503",
	"StatusGatewayTimeout":                "504",
	"StatusHTTPVersionNotSupported":       "505",
	"StatusVariantAlsoNegotiates":         "506",
	"StatusInsufficientStorage":           "507",
	"StatusLoopDetected":                  "508",
	"StatusNotExtended":                   "510",
	"StatusNetworkAuthenticationRequired": "511",
}

func ServeSwaggerDocs(r *gin.Engine, docsPath, docsFilePath string, docsFile []byte) []string {
	routeList := []string{}
	subFS, err := fs.Sub(swaggerFiles, ".")
	if err != nil {
		panic(err)
	}
	r.GET(docsPath+"/"+docsFilePath, func(c *gin.Context) { c.Data(200, "application/json", docsFile) })
	// loop semua file embed
	fs.WalkDir(subFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			// tentukan URL route berdasarkan nama file
			route := fmt.Sprintf("%s/%s", docsPath, path)
			route = strings.ReplaceAll(route, "\\", "/") // jaga-jaga untuk Windows
			if path == "index.html" {
				route = docsPath + "/"
			}
			routeList = append(routeList, route)
			// daftarkan handler untuk setiap file
			r.GET(route, func(c *gin.Context) {
				// Open file from embedded FS
				file, err := subFS.Open(path)
				if err != nil {
					c.String(http.StatusNotFound, "file not found: %s", path)
					return
				}
				defer file.Close()

				// Read file content
				data, err := util.ReadAllFromFile(file)
				if err != nil {
					c.String(http.StatusInternalServerError, "error reading file: %s", path)
					return
				}
				if path == "swagger-initializer.js" {
					data = bytes.ReplaceAll(data, []byte("docs.json"), []byte(docsFilePath))
				}
				contentType := mime.TypeByExtension(filepath.Ext(path))
				if contentType == "" {
					contentType = "application/octet-stream"
				}
				c.Data(http.StatusOK, contentType, data)
			})

			// fmt.Println("Serving embedded file:", route)
		}
		return nil
	})
	return routeList
}

func GenerateSwaggerDoc(routes *gin.Engine, docsFilePath string, excludedPath ...string) {
	module_name := getModuleName()

	swaggerDoc := make(map[string]interface{})
	// Read the existing docsFilePath file
	file, err := os.Open(docsFilePath)
	if err != nil {
		logrus.Error(err)
	} else {
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&swaggerDoc); err != nil {
			log.Printf("Failed to decode "+docsFilePath+" file: %v", err)
			swaggerDoc = make(map[string]interface{})
		}
	}
	defer file.Close()

	// Update the swaggerDoc with new information from environment variables
	swaggerDoc["openapi"] = "3.0.4"

	// Build info section from environment variables
	infoSection := map[string]interface{}{
		"title":       util.Getenv("APP_NAME", "API Documentation"),
		"description": util.Getenv("APP_DESCRIPTION", "API Documentation for "+os.Getenv("APP_NAME")),
		"version":     util.Getenv("APP_VERSION", "0.0.1"),
	}

	// Add contact info if provided
	contactName := os.Getenv("APP_CONTACT_NAME")
	contactEmail := os.Getenv("APP_CONTACT_EMAIL")
	contactUrl := os.Getenv("APP_CONTACT_URL")
	if contactName != "" || contactEmail != "" || contactUrl != "" {
		contact := make(map[string]interface{})
		if contactName != "" {
			contact["name"] = contactName
		}
		if contactEmail != "" {
			contact["email"] = contactEmail
		}
		if contactUrl != "" {
			contact["url"] = contactUrl
		}
		infoSection["contact"] = contact
	}

	// Add license info if provided
	licenseName := util.Getenv("APP_LICENSE_NAME", "")
	licenseUrl := util.Getenv("APP_LICENSE_URL", "")
	if licenseName != "" {
		license := map[string]interface{}{
			"name": licenseName,
		}
		if licenseUrl != "" {
			license["url"] = licenseUrl
		}
		infoSection["license"] = license
	}

	// Add terms of service if provided
	termsUrl := util.Getenv("APP_TERMS_URL", "")
	if termsUrl != "" {
		infoSection["termsOfService"] = termsUrl
	}

	swaggerDoc["info"] = infoSection

	// Add external docs if provided
	externalDocsDesc := util.Getenv("APP_EXTERNAL_DOCS_DESC", "Find out more about our API")
	externalDocsUrl := util.Getenv("APP_EXTERNAL_DOCS_URL", "")
	if externalDocsUrl != "" {
		swaggerDoc["externalDocs"] = map[string]interface{}{
			"description": externalDocsDesc,
			"url":         externalDocsUrl,
		}
	}

	// Build servers section from environment variables
	servers := buildServersFromEnv()
	if len(servers) == 0 {
		// Fallback to SwaggerInfo.Host if no servers configured
		// if SwaggerInfo.Host != "" {
		// }
		localhost := util.Getenv("APP_LOCAL_HOST", "http://localhost:8080")
		if !strings.Contains(localhost, "localhost") &&
			strings.Contains(localhost, "127.") &&
			strings.HasPrefix(localhost, ":") {
			localhost = "http://localhost" + localhost
		}
		servers = append(servers, map[string]interface{}{
			"url":         util.Getenv("APP_PUBLIC_URL", localhost),
			"description": "Default server",
		})
	}
	swaggerDoc["servers"] = servers

	// Add security schemes from environment
	swaggerDoc["components"] = buildComponentsFromEnv()

	// List all API routes and update the paths
	for _, route := range routes.Routes() {
		// Convert route.Path to Swagger format
		swaggerPath := route.Path
		// Skip excluded paths
		if util.Contains(excludedPath, swaggerPath) {
			continue
		}
		var req_params []map[string]interface{}
		if strings.Contains(route.Path, ":") {
			parts := strings.Split(route.Path, "/")
			for i, part := range parts {
				if strings.HasPrefix(part, ":") {
					parts[i] = "{" + strings.TrimPrefix(part, ":") + "}"
					req_params = append(req_params, map[string]interface{}{
						"name":     strings.TrimPrefix(part, ":"), // name of the parameter
						"in":       "path",                        // location of the parameter
						"required": true,                          // whether the parameter is required
						"type":     "string",                      // type of the parameter
					})
				}
			}
			// swaggerPath = strings.Join(parts, "/")
			// if !strings.HasPrefix(swaggerPath, "/") {
			// 	swaggerPath = "/" + swaggerPath
			// }
		}
		// fmt.Println("__________________________")
		// fmt.Println("Route:", route.Method, route.Path)
		var (
			parameters, request, response map[string]interface{}
		)
		funcPtr := reflect.ValueOf(route.HandlerFunc).Pointer()
		funcInfo := runtime.FuncForPC(funcPtr)
		if funcInfo != nil {
			file, line := funcInfo.FileLine(funcPtr) // Get file and line number
			// fmt.Println("File:", file)
			// fmt.Println("Handler Line:", line)

			// Parse and analyze the handler function
			parameters, request, response = parseHandlerFunction(file, line)
		}
		// Merge req_params and parameters
		for _, param := range parameters {
			req_params = append(req_params, param.(map[string]interface{}))
		}
		// fmt.Println("parameters")
		// fmt.Println(parameters)
		// fmt.Println("request")
		// fmt.Println(request)
		// fmt.Println("response")
		// fmt.Println(response)
		// // Convert response map to JSON and log it
		// responseJSON, err := json.MarshalIndent(response, "", "  ")
		// if err != nil {
		// logrus.Error(err)
		// 	log.Printf("Failed to marshal response to JSON: %v", err)
		// } else {
		// 	log.Printf("Response JSON: %s", responseJSON)
		// }
		// fmt.Println("__________________________")
		summaryInfo := route.Path
		if strings.HasPrefix(route.Handler, module_name+"/") {
			parts := strings.Split(strings.TrimPrefix(route.Handler, module_name+"/"), ".")
			if len(parts) >= 2 {
				summaryInfo = parts[len(parts)-2]
				var splitSummary []string
				for i, r := range summaryInfo {
					if i > 0 && r >= 'A' && r <= 'Z' {
						splitSummary = append(splitSummary, " ")
					}
					splitSummary = append(splitSummary, string(r))
				}
				summaryInfo = strings.Join(splitSummary, "")
			}
		}

		paths, ok := swaggerDoc["paths"].(map[string]interface{})
		if !ok {
			paths = make(map[string]interface{})
			swaggerDoc["paths"] = paths
		}
		// pathItem, exists := paths[swaggerPath]
		// if !exists {
		// }
		pathItem := map[string]interface{}{}
		paths[swaggerPath] = pathItem
		tags := getTagFromPath(route.Path)
		securityConfig := buildSecurityForRoute(route.Path)

		methodItem := map[string]interface{}{
			"tags":        []string{tags},
			"summary":     generateSummary(summaryInfo, route.Method),
			"description": generateDescription(swaggerPath, route.Method),
			"operationId": generateOperationId(route.Method, swaggerPath),
			"parameters":  req_params,
			"requestBody": request,
			"responses":   enhanceResponses(response, route.Method),
			"security":    securityConfig,
		}

		pathItem[strings.ToLower(route.Method)] = methodItem
	}

	// Add global security schemes
	globalSecurity := buildGlobalSecurity()
	if len(globalSecurity) > 0 {
		swaggerDoc["security"] = globalSecurity
	}

	// Write the updated swaggerDoc to docsFilePath file
	// Create docs directory if it doesn't exist
	// if err := os.MkdirAll("./"+docsFilePath, 0755); err != nil {
	// 	logrus.Error(err)
	// 	log.Fatalf("Failed to create docs directory: %v", err)
	// }

	// Write the updated swaggerDoc to docsFilePath file
	file, err = os.Create(docsFilePath)
	if err != nil {
		logrus.Error(err)
		log.Fatalf("Failed to create "+docsFilePath+" file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(swaggerDoc); err != nil {
		log.Fatalf("Failed to encode swaggerDoc to JSON: %v", err)
	}
}
func getModuleName() string {
	cmd := exec.Command("go", "list", "-m")
	out, err := cmd.Output()
	if err != nil {
		logrus.Error(err)
		return "unknown"
	}
	return strings.TrimSpace(string(out))
}

// Build servers configuration from environment variables
func buildServersFromEnv() []map[string]interface{} {
	var servers []map[string]interface{}

	// Development server
	devUrl := util.Getenv("APP_DEV_SERVER_URL", "", "http://localhost:8080")
	devDesc := util.Getenv("APP_DEV_SERVER_DESC", "", "Development server")
	if devUrl != "" {
		servers = append(servers, map[string]interface{}{
			"url":         devUrl,
			"description": devDesc,
		})
	}

	// Staging server (optional)
	stagingUrl := os.Getenv("APP_STAGING_SERVER_URL")
	stagingDesc := util.Getenv("APP_STAGING_SERVER_DESC", "", "Staging server")
	if stagingUrl != "" {
		servers = append(servers, map[string]interface{}{
			"url":         stagingUrl,
			"description": stagingDesc,
		})
	}

	// Production server (optional)
	prodUrl := os.Getenv("APP_PROD_SERVER_URL")
	prodDesc := os.Getenv("APP_PROD_SERVER_DESC")
	if prodUrl != "" {
		servers = append(servers, map[string]interface{}{
			"url":         prodUrl,
			"description": prodDesc,
		})
	}

	return servers
}

// Build components (security schemes, schemas) from environment variables
func buildComponentsFromEnv() map[string]interface{} {
	components := make(map[string]interface{})

	// Security schemes
	securitySchemes := make(map[string]interface{})

	// API Key authentication
	apiKeyHeader := os.Getenv("APP_API_KEY_HEADER")
	if apiKeyHeader != "" {
		securitySchemes["ApiKeyAuth"] = map[string]interface{}{
			"type":        "apiKey",
			"in":          "header",
			"name":        apiKeyHeader,
			"description": "API key for authentication",
		}
	}

	// Bearer token authentication
	enableBearer := os.Getenv("APP_ENABLE_BEARER")
	if enableBearer == "true" {
		securitySchemes["BearerAuth"] = map[string]interface{}{
			"type":         "http",
			"scheme":       "bearer",
			"bearerFormat": "JWT",
			"description":  "JWT Bearer token",
		}
	}

	// Basic authentication (optional)
	enableBasic := os.Getenv("APP_ENABLE_BASIC")
	if enableBasic == "true" {
		securitySchemes["BasicAuth"] = map[string]interface{}{
			"type":        "http",
			"scheme":      "basic",
			"description": "Basic HTTP authentication",
		}
	}

	// OAuth2 (optional)
	oauthAuthUrl := os.Getenv("APP_OAUTH2_AUTH_URL")
	oauthTokenUrl := os.Getenv("APP_OAUTH2_TOKEN_URL")
	if oauthAuthUrl != "" && oauthTokenUrl != "" {
		flows := map[string]interface{}{
			"authorizationCode": map[string]interface{}{
				"authorizationUrl": oauthAuthUrl,
				"tokenUrl":         oauthTokenUrl,
				"scopes": map[string]interface{}{
					"read":  "Read access",
					"write": "Write access",
				},
			},
		}
		securitySchemes["OAuth2"] = map[string]interface{}{
			"type":        "oauth2",
			"description": "OAuth2 authentication",
			"flows":       flows,
		}
	}

	if len(securitySchemes) > 0 {
		components["securitySchemes"] = securitySchemes
	}

	// Add common schemas
	schemas := buildCommonSchemas()
	if len(schemas) > 0 {
		components["schemas"] = schemas
	}

	return components
}

// Build common reusable schemas
func buildCommonSchemas() map[string]interface{} {
	schemas := make(map[string]interface{})

	// Generic Error Response schema
	schemas["ErrorResponse"] = map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"error": map[string]interface{}{
				"type":        "string",
				"description": "Error message",
				"example":     "Invalid request",
			},
			"code": map[string]interface{}{
				"type":        "integer",
				"description": "Error code",
				"example":     400,
			},
			"details": map[string]interface{}{
				"type":                 "object",
				"description":          "Additional error details",
				"additionalProperties": true,
			},
		},
		"required": []string{"error", "code"},
	}

	// Success Response schema
	schemas["SuccessResponse"] = map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"success": map[string]interface{}{
				"type":        "boolean",
				"description": "Operation success status",
				"example":     true,
			},
			"message": map[string]interface{}{
				"type":        "string",
				"description": "Success message",
				"example":     "Operation completed successfully",
			},
			"data": map[string]interface{}{
				"type":                 "object",
				"description":          "Response data",
				"additionalProperties": true,
			},
		},
		"required": []string{"success"},
	}

	// Pagination schema
	schemas["PaginationMeta"] = map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"page": map[string]interface{}{
				"type":        "integer",
				"description": "Current page number",
				"example":     1,
				"minimum":     1,
			},
			"limit": map[string]interface{}{
				"type":        "integer",
				"description": "Items per page",
				"example":     10,
				"minimum":     1,
				"maximum":     100,
			},
			"total": map[string]interface{}{
				"type":        "integer",
				"description": "Total number of items",
				"example":     100,
			},
			"totalPages": map[string]interface{}{
				"type":        "integer",
				"description": "Total number of pages",
				"example":     10,
			},
		},
		"required": []string{"page", "limit", "total", "totalPages"},
	}

	return schemas
}

// Get tag from path for better organization
func getTagFromPath(path string) string {
	basePath := os.Getenv("VITE_BASE_PATH")
	cleanPath := strings.TrimPrefix(path, basePath)
	if cleanPath == "" || cleanPath == "/" {
		return "default"
	}

	parts := strings.Split(strings.Trim(cleanPath, "/"), "/")
	if len(parts) > 0 && parts[0] != "" {
		// Capitalize first letter and make it human readable
		tag := parts[0]
		if len(tag) > 0 {
			tag = strings.ToUpper(tag[:1]) + tag[1:]
		}
		return tag
	}
	return "default"
}

// Build security configuration for specific route
func buildSecurityForRoute(path string) []map[string]interface{} {
	var security []map[string]interface{}

	// Check if this is a public endpoint (no authentication required)
	publicPaths := strings.Split(os.Getenv("APP_PUBLIC_PATHS"), ",")
	for _, publicPath := range publicPaths {
		if strings.TrimSpace(publicPath) == path {
			return security // Empty security = public endpoint
		}
	}

	// Default security scheme
	defaultAuth := util.Getenv("APP_DEFAULT_AUTH", "", "ApiKeyAuth")
	if defaultAuth != "" {
		security = append(security, map[string]interface{}{
			defaultAuth: []string{},
		})
	}

	return security
}

// Generate human-readable summary
func generateSummary(summaryInfo, method string) string {
	method = strings.ToUpper(method)

	// Custom summaries based on method and path patterns
	switch method {
	case "GET":
		if strings.Contains(summaryInfo, "{") {
			return "Get " + strings.ReplaceAll(summaryInfo, " ", "") + " by ID"
		}
		return "Get all " + strings.ReplaceAll(summaryInfo, " ", "")
	case "POST":
		return "Create new " + strings.ReplaceAll(summaryInfo, " ", "")
	case "PUT":
		return "Update " + strings.ReplaceAll(summaryInfo, " ", "")
	case "PATCH":
		return "Partially update " + strings.ReplaceAll(summaryInfo, " ", "")
	case "DELETE":
		return "Delete " + strings.ReplaceAll(summaryInfo, " ", "")
	default:
		return method + " " + summaryInfo
	}
}

// Generate description for the endpoint
func generateDescription(path, method string) string {
	method = strings.ToUpper(method)

	switch method {
	case "GET":
		if strings.Contains(path, "{") {
			return "Retrieve a specific resource by its unique identifier"
		}
		return "Retrieve a list of resources with optional filtering and pagination"
	case "POST":
		return "Create a new resource with the provided data"
	case "PUT":
		return "Update an existing resource with the provided data"
	case "PATCH":
		return "Partially update an existing resource"
	case "DELETE":
		return "Remove an existing resource"
	default:
		return "Perform " + strings.ToLower(method) + " operation on " + path
	}
}

// Generate operation ID for the endpoint
func generateOperationId(method, path string) string {
	method = strings.ToLower(method)

	// Clean path and create camelCase operation ID
	cleanPath := strings.ReplaceAll(path, "/", "_")
	cleanPath = strings.ReplaceAll(cleanPath, "{", "")
	cleanPath = strings.ReplaceAll(cleanPath, "}", "")
	cleanPath = strings.Trim(cleanPath, "_")

	// Convert to camelCase
	parts := strings.Split(cleanPath, "_")
	var camelParts []string
	camelParts = append(camelParts, method)

	for _, part := range parts {
		if part != "" {
			camelParts = append(camelParts, strings.Title(part))
		}
	}

	return strings.Join(camelParts, "")
}

// Enhance responses with proper status codes and examples
func enhanceResponses(responses map[string]interface{}, method string) map[string]interface{} {
	if responses == nil {
		responses = make(map[string]interface{})
	}

	method = strings.ToUpper(method)

	// Add default success response if not present
	switch method {
	case "GET":
		if _, exists := responses["200"]; !exists {
			responses["200"] = map[string]interface{}{
				"description": "Successful response",
				"content": map[string]interface{}{
					"application/json": map[string]interface{}{
						"schema": map[string]interface{}{
							"$ref": "#/components/schemas/SuccessResponse",
						},
					},
				},
			}
		}
	case "POST":
		if _, exists := responses["201"]; !exists {
			responses["201"] = map[string]interface{}{
				"description": "Resource created successfully",
				"content": map[string]interface{}{
					"application/json": map[string]interface{}{
						"schema": map[string]interface{}{
							"$ref": "#/components/schemas/SuccessResponse",
						},
					},
				},
			}
		}
	case "PUT", "PATCH":
		if _, exists := responses["200"]; !exists {
			responses["200"] = map[string]interface{}{
				"description": "Resource updated successfully",
				"content": map[string]interface{}{
					"application/json": map[string]interface{}{
						"schema": map[string]interface{}{
							"$ref": "#/components/schemas/SuccessResponse",
						},
					},
				},
			}
		}
	case "DELETE":
		if _, exists := responses["204"]; !exists {
			responses["204"] = map[string]interface{}{
				"description": "Resource deleted successfully",
			}
		}
	}

	// Add common error responses
	commonErrors := []string{"400", "401", "403", "404", "500"}
	for _, code := range commonErrors {
		if _, exists := responses[code]; !exists && shouldAddErrorResponse(method, code) {
			responses[code] = map[string]interface{}{
				"description": getErrorDescription(code),
				"content": map[string]interface{}{
					"application/json": map[string]interface{}{
						"schema": map[string]interface{}{
							"$ref": "#/components/schemas/ErrorResponse",
						},
					},
				},
			}
		}
	}

	return responses
}

// Check if error response should be added for this method
func shouldAddErrorResponse(method, code string) bool {
	switch code {
	case "400":
		return method == "POST" || method == "PUT" || method == "PATCH"
	case "401":
		return true // All endpoints can return unauthorized
	case "403":
		return true // All endpoints can return forbidden
	case "404":
		return method == "GET" || method == "PUT" || method == "PATCH" || method == "DELETE"
	case "500":
		return true // All endpoints can return server error
	default:
		return false
	}
}

// Get error description for status code
func getErrorDescription(code string) string {
	descriptions := map[string]string{
		"400": "Bad request - Invalid input data",
		"401": "Unauthorized - Authentication required",
		"403": "Forbidden - Insufficient permissions",
		"404": "Not found - Resource does not exist",
		"500": "Internal server error",
	}

	if desc, exists := descriptions[code]; exists {
		return desc
	}
	return "Error response"
}

// Build global security configuration
func buildGlobalSecurity() []map[string]interface{} {
	var security []map[string]interface{}

	defaultAuth := util.Getenv("APP_DEFAULT_AUTH", "", "ApiKeyAuth")
	if defaultAuth != "" {
		security = append(security, map[string]interface{}{
			defaultAuth: []string{},
		})
	}

	return security
}

func parseHandlerFunction(filePath string, targetLine int) (map[string]interface{}, map[string]interface{}, map[string]interface{}) {
	requestParameters := make(map[string]interface{})
	body := make(map[string]interface{})
	response := make(map[string]interface{})
	// info := getTypesInfo(filePath)

	// Parse the Go source file
	fs := token.NewFileSet()
	node, err := parser.ParseFile(fs, filePath, nil, parser.AllErrors)
	if err != nil {
		logrus.Error(err)
		log.Fatalf("Error parsing file: %v", err)
	}

	// Traverse the AST to find function literals matching the line number
	ast.Inspect(node, func(n ast.Node) bool {
		if fnLit, ok := n.(*ast.FuncLit); ok { // Check for function literals (anonymous functions)
			pos := fs.Position(fnLit.Pos())

			// Match the function literal by line number
			if pos.Line == targetLine {
				// fmt.Printf("Found handler function at line %d\n", targetLine)

				// Analyze function body for `c.Method(...)` calls
				ast.Inspect(fnLit.Body, func(n ast.Node) bool {
					if callExpr, ok := n.(*ast.CallExpr); ok {
						if selExpr, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
							if ident, ok := selExpr.X.(*ast.Ident); ok && ident.Name == "c" {
								methodName := selExpr.Sel.Name
								// fmt.Printf("  - %s()\n", methodName)

								// Analyze request parameters
								if strings.HasPrefix(methodName, "Query") || strings.HasPrefix(methodName, "Header") || strings.HasPrefix(methodName, "Cookie") {
									// Handle both static strings and dynamic parameters
									paramName := ""
									isDynamic := false

									if len(callExpr.Args) > 0 {
										switch arg := callExpr.Args[0].(type) {
										case *ast.BasicLit:
											// Static string: c.Query("name")
											paramName = strings.Trim(arg.Value, "\"")
										case *ast.CallExpr:
											// Dynamic string with fmt.Sprintf: c.Query(fmt.Sprintf("columns[%d][data]", i))
											if fun, ok := arg.Fun.(*ast.SelectorExpr); ok {
												if ident, ok := fun.X.(*ast.Ident); ok && ident.Name == "fmt" && fun.Sel.Name == "Sprintf" {
													if len(arg.Args) > 0 {
														if formatStr, ok := arg.Args[0].(*ast.BasicLit); ok {
															// Use the format string as parameter pattern
															paramName = strings.Trim(formatStr.Value, "\"")
															// Convert "%d" or "%s" to "*" for pattern matching
															paramName = strings.ReplaceAll(paramName, "%d", "*")
															paramName = strings.ReplaceAll(paramName, "%s", "*")
															isDynamic = true
														}
													}
												}
											}
										case *ast.Ident:
											// Variable: c.Query(paramVar)
											paramName = arg.Name
											isDynamic = true
										}
									}

									if paramName != "" {
										description := fmt.Sprintf("%s %s", methodName, paramName)
										if isDynamic {
											description += " (dynamic parameter)"
										}

										requestParameters[paramName] = map[string]interface{}{
											"name":        paramName,
											"in":          strings.ToLower(strings.TrimPrefix(methodName, "Query")),
											"required":    false, // Dynamic params are typically optional
											"schema":      map[string]interface{}{"type": "string"},
											"description": description,
										}
									}
								}

								// Analyze request body
								if strings.HasPrefix(methodName, "BindJSON") || strings.HasPrefix(methodName, "ShouldBindJSON") {
									body["required"] = true
									body["content"] = map[string]interface{}{
										"application/json": map[string]interface{}{
											"schema": parseStructType(callExpr.Args[0]),
										},
									}
								} else if strings.HasPrefix(methodName, "Bind") || strings.HasPrefix(methodName, "ShouldBind") {
									body["required"] = true
									body["content"] = map[string]interface{}{
										"application/x-www-form-urlencoded": map[string]interface{}{
											"schema": parseStructType(callExpr.Args[0]),
										},
									}
								} else if strings.HasPrefix(methodName, "FormFile") || strings.HasPrefix(methodName, "SaveUploadedFile") {
									// Extract parameter name safely
									paramName := ""
									if len(callExpr.Args) > 0 {
										switch arg := callExpr.Args[0].(type) {
										case *ast.BasicLit:
											// Static string: c.FormFile("file")
											paramName = strings.Trim(arg.Value, "\"")
										case *ast.Ident:
											// Variable: c.FormFile(fieldName)
											paramName = arg.Name
										default:
											paramName = "file" // Default fallback
										}
									}

									if paramName == "" {
										paramName = "file"
									}

									// Only add to body if not already present or merge with existing
									if existingContent, ok := body["content"].(map[string]interface{}); ok {
										if multipartData, ok := existingContent["multipart/form-data"].(map[string]interface{}); ok {
											if schema, ok := multipartData["schema"].(map[string]interface{}); ok {
												if properties, ok := schema["properties"].(map[string]interface{}); ok {
													// Add this file parameter to existing properties
													properties[paramName] = map[string]interface{}{
														"type":   "string",
														"format": "binary",
													}
												}
											}
										}
									} else {
										body["required"] = true
										body["content"] = map[string]interface{}{
											"multipart/form-data": map[string]interface{}{
												"schema": map[string]interface{}{
													"type": "object",
													"properties": map[string]interface{}{
														paramName: map[string]interface{}{
															"type":   "string",
															"format": "binary",
														},
													},
												},
											},
										}
									}
								}

								// Detect ContentType() to identify multipart forms
								if methodName == "ContentType" {
									// This indicates the handler checks content type
									// We can infer that it accepts multiple content types
									// Mark body as supporting multiple content types
									if body["content"] == nil {
										body["content"] = make(map[string]interface{})
									}
								}

								// Analyze response
								if strings.HasPrefix(methodName, "JSON") || strings.HasPrefix(methodName, "IndentedJSON") {
									var responseCode string
									switch arg := callExpr.Args[0].(type) {
									case *ast.BasicLit:
										responseCode = strings.Trim(arg.Value, "\"")
									case *ast.SelectorExpr:
										if ident, ok := arg.X.(*ast.Ident); ok {
											if ident.Name == "http" {
												if code, exists := statusCodes[arg.Sel.Name]; exists {
													responseCode = code
												} else {
													responseCode = arg.Sel.Name
												}
											} else {
												responseCode = ident.Name + "." + arg.Sel.Name
											}
										} else {
											responseCode = "200" // Default fallback
										}
									case *ast.Ident:
										responseCode = arg.Name
									default:
										responseCode = "200"
										log.Printf("Unexpected response code type: %T", arg)
									}

									var responseBody map[string]interface{}

									if len(callExpr.Args) > 1 {
										if bodyExpr, ok := callExpr.Args[1].(*ast.CompositeLit); ok {
											responseBody = parseGinH(bodyExpr)
											for key, value := range responseBody {
												if strValue, ok := value.(string); ok {
													responseBody[key] = strings.Trim(strValue, "\"")
												}
											}
										}
									}

									// Ensure response map exists
									if _, exists := response[responseCode]; !exists {
										response[responseCode] = map[string]interface{}{
											"description": "Response " + responseCode,
											"content": map[string]interface{}{
												"application/json": map[string]interface{}{
													"examples": map[string]interface{}{},
												},
											},
										}
									}

									// Navigate to examples map with proper nil checks
									responseMap, ok := response[responseCode].(map[string]interface{})
									if !ok {
										// Skip if response structure is invalid
										log.Printf("Warning: Invalid response structure for code %s", responseCode)
									} else if content, ok := responseMap["content"].(map[string]interface{}); !ok {
										// Skip if content is invalid
										log.Printf("Warning: Invalid content structure for response code %s", responseCode)
									} else if jsonContent, ok := content["application/json"].(map[string]interface{}); !ok {
										// Skip if json content is invalid
										log.Printf("Warning: Invalid JSON content structure for response code %s", responseCode)
									} else {
										examples, ok := jsonContent["examples"].(map[string]interface{})
										if !ok {
											// Initialize examples if it doesn't exist
											examples = make(map[string]interface{})
											jsonContent["examples"] = examples
										}

										// Find next available example key
										exampleIndex := 1
										exampleKey := fmt.Sprintf("Example%d", exampleIndex)
										for {
											if _, exists := examples[exampleKey]; !exists {
												break
											}
											exampleIndex++
											exampleKey = fmt.Sprintf("Example%d", exampleIndex)
										}

										// Append new example
										examples[exampleKey] = map[string]interface{}{
											"value": responseBody,
										}
									}
								} else if strings.HasPrefix(methodName, "HTML") || strings.HasPrefix(methodName, "String") || strings.HasPrefix(methodName, "File") || strings.HasPrefix(methodName, "FileAttachment") || strings.HasPrefix(methodName, "XML") || strings.HasPrefix(methodName, "YAML") || strings.HasPrefix(methodName, "TOML") || strings.HasPrefix(methodName, "Redirect") || strings.HasPrefix(methodName, "Data") {
									var responseCode string
									switch arg := callExpr.Args[0].(type) {
									case *ast.BasicLit:
										responseCode = strings.Trim(arg.Value, "\"")
									case *ast.SelectorExpr:
										if ident, ok := arg.X.(*ast.Ident); ok {
											if ident.Name == "http" {
												if code, exists := statusCodes[arg.Sel.Name]; exists {
													responseCode = code
												} else {
													responseCode = arg.Sel.Name
												}
											} else {
												responseCode = ident.Name + "." + arg.Sel.Name
											}
										} else {
											responseCode = "200" // Default fallback
										}
									case *ast.Ident:
										responseCode = arg.Name
									default:
										responseCode = "200"
										log.Printf("Unexpected response code type: %T", arg)
									}

									// Ensure response map exists
									if _, exists := response[responseCode]; !exists {
										response[responseCode] = map[string]interface{}{
											"description": "Response " + responseCode,
										}
									}

									// Add content type based on method
									contentType := "text/plain"
									if strings.HasPrefix(methodName, "HTML") {
										contentType = "text/html"
									} else if strings.HasPrefix(methodName, "XML") {
										contentType = "application/xml"
									} else if strings.HasPrefix(methodName, "YAML") {
										contentType = "application/x-yaml"
									} else if strings.HasPrefix(methodName, "TOML") {
										contentType = "application/toml"
									} else if strings.HasPrefix(methodName, "Data") {
										if len(callExpr.Args) > 1 {
											if ct, ok := callExpr.Args[1].(*ast.BasicLit); ok {
												contentType = strings.Trim(ct.Value, "\"")
											}
										}
									}

									response[responseCode].(map[string]interface{})["content"] = map[string]interface{}{
										contentType: map[string]interface{}{
											"schema": map[string]interface{}{
												"type": "string",
											},
										},
									}
								}
							}
						}
					}
					return true
				})
			}
		}
		return true
	})

	return requestParameters, body, response
}
func parseStructType(expr ast.Expr) map[string]interface{} {
	structSchema := make(map[string]interface{})
	properties := make(map[string]interface{})
	structSchema["type"] = "object"
	structSchema["properties"] = properties

	switch v := expr.(type) {
	case *ast.CompositeLit:
		if ident, ok := v.Type.(*ast.Ident); ok {
			structName := ident.Name
			structSchema["title"] = structName

			// Parse the struct fields
			for _, elt := range v.Elts {
				if kvExpr, ok := elt.(*ast.KeyValueExpr); ok {
					if key, ok := kvExpr.Key.(*ast.Ident); ok {
						fieldName := key.Name
						fieldType := parseFieldType(kvExpr.Value)
						properties[fieldName] = fieldType
					}
				}
			}
		}
	}

	return structSchema
}

func parseFieldType(expr ast.Expr) map[string]interface{} {
	fieldType := make(map[string]interface{})

	switch v := expr.(type) {
	case *ast.Ident:
		fieldType["type"] = v.Name
	case *ast.ArrayType:
		fieldType["type"] = "array"
		fieldType["items"] = parseFieldType(v.Elt)
	case *ast.MapType:
		fieldType["type"] = "object"
		fieldType["additionalProperties"] = parseFieldType(v.Value)
	case *ast.StarExpr:
		return parseFieldType(v.X)
	case *ast.SelectorExpr:
		if ident, ok := v.X.(*ast.Ident); ok {
			fieldType["type"] = ident.Name + "." + v.Sel.Name
		}
	}

	return fieldType
}
func parseGinH(lit *ast.CompositeLit) map[string]interface{} {
	result := make(map[string]interface{})

	for _, elt := range lit.Elts {
		if kvExpr, ok := elt.(*ast.KeyValueExpr); ok {
			if key, ok := kvExpr.Key.(*ast.BasicLit); ok {
				keyStr := strings.Trim(key.Value, "\"")

				switch v := kvExpr.Value.(type) {
				case *ast.BasicLit: // String, Number, Boolean
					result[keyStr] = v.Value
				case *ast.CompositeLit: // Nested map (e.g., data structures)
					result[keyStr] = "[]interface{}{}"
				case *ast.Ident: // Variable reference
					result[keyStr] = v.Name
				case *ast.BinaryExpr: // Expressions like `err.Error()`
					result[keyStr] = stringifyBinaryExpr(v)
				default:
					result[keyStr] = fmt.Sprintf("%T", v) // Debug unknown types
				}
			}
		}
	}
	return result
}
func stringifyBinaryExpr(expr *ast.BinaryExpr) string {
	left := stringifyExpr(expr.X)
	right := stringifyExpr(expr.Y)
	return left + " " + expr.Op.String() + " " + right
}

func stringifyExpr(expr ast.Expr) string {
	switch v := expr.(type) {
	case *ast.Ident:
		return v.Name
	case *ast.BasicLit:
		return v.Value
	case *ast.CallExpr: // Function call like err.Error()
		if fun, ok := v.Fun.(*ast.SelectorExpr); ok {
			return stringifyExpr(fun.X) + "." + fun.Sel.Name + "()"
		}
	}
	return "unknown"
}
