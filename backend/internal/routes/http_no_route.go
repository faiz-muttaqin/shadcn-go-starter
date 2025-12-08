package routes

import (
	"fmt"
	"io/fs"
	"mime"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/clr"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/util"
	"github.com/gin-gonic/gin"
)

// Shadcn Admin Go Starter

func NoRouteDefaultFiles(embedDistFolder fs.FS, isDevMode bool) []string {
	var useFrontendProxy = false
	var htmlContentType string
	var htmlContent []byte
	routeList := []string{}
	if isDevMode {
		useFrontendProxy = CheckIsProxyAvailable()
	}
	files, _ := fs.Sub(embedDistFolder, "dist")
	subFS, err := fs.Sub(files, ".")
	if err != nil {
		panic(err)
	}
	fs.WalkDir(subFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			file, err := subFS.Open(path)
			if err != nil {
				return err
			}
			// Open file from embedded FS
			defer file.Close()

			// Read file content
			data, err := util.ReadAllFromFile(file)
			if err != nil {
				return err
			}
			contentType := mime.TypeByExtension(filepath.Ext(path))
			if contentType == "" {
				contentType = "application/octet-stream"
			}
			route := strings.ReplaceAll(path, "\\", "/") // jaga-jaga untuk Windows
			if path == "index.html" {
				htmlContentType = contentType
				htmlContent = data
				return nil
			} else if useFrontendProxy {
				return nil
			}
			route = strings.TrimSuffix(route, ".html")
			routeList = append(routeList, route)
			// daftarkan handler untuk setiap file
			route, _ = url.JoinPath(os.Getenv("VITE_BASE_PATH"), route)
			R.GET(route, func(c *gin.Context) {
				c.Header("Cache-Control", "public, max-age=3600")
				c.Data(http.StatusOK, contentType, data)
			})

			// fmt.Println("Serving embedded file:", route)
		}
		return nil
	})

	for i := 0; i < 12; i++ {
		route := ""
		y := 0
		for range i {
			route = route + fmt.Sprintf("/:p%d", y)
			y++
		}
		if route == "" {
			route = "/"
		}
		route, _ = url.JoinPath(os.Getenv("VITE_BASE_PATH"), route)
		R.GET(route, func(c *gin.Context) {
			if useFrontendProxy {
				ProxyToVite(c)
				return
			}
			c.Header("Cache-Control", "public, max-age=1800")
			c.Data(http.StatusOK, htmlContentType, htmlContent)
		})
	}
	R.NoRoute(func(c *gin.Context) {
		if useFrontendProxy {
			ProxyToVite(c)
			return
		}
		c.Header("Cache-Control", "public, max-age=1800")
		c.Data(http.StatusOK, htmlContentType, htmlContent)
	})
	return routeList
}
func CheckIsProxyAvailable() bool {
	if resp, err := http.Get("http://localhost:5173"); err == nil && resp.StatusCode == http.StatusOK {
		// logrus.Println(clr.BgGreen("Use Frontend Via Proxy"))
		fmt.Println(clr.BgMagenta(clr.TextBlack("FRONTEND_PROXY_ENABLED")))
		return true
	}
	return false
}

func ProxyToVite(c *gin.Context) {

	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = "localhost:5173"
		req.URL.Path = c.Request.URL.Path
		req.URL.RawQuery = c.Request.URL.RawQuery
		req.Header = c.Request.Header
	}
	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(c.Writer, c.Request)
}
