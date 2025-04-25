package plugins

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"log"
	"net/http"
	"net/url"

	"github.com/eclipse-xfsc/cloud-wallet-plugin-discovery/kong"
)

type PluginInfo struct {
	Name  string `json:"name"`
	Route string `json:"route"`
	URL   string `json:"url"`
}

type PluginsService struct {
}

func (t *PluginsService) AddRoutes(group *gin.RouterGroup) error {

	pluginsGroup := group.Group("/plugins")

	pluginsGroup.GET("/", func(ctx *gin.Context) {
		ListPlugins(ctx)
	})

	return nil
}

func ListPlugins(ctx *gin.Context) {

	scheme := viper.GetString("KONG_SCHEME")
	if scheme == "" {
		scheme = "http"
	}

	kongHost := viper.GetString("KONG_HOST")
	if kongHost == "" {
		log.Fatal("missing kong host (expected to be passed via ENV['KONG_HOST'])")
	}

	pluginTag := viper.GetString("KONG_PLUGIN_TAG")
	if pluginTag == "" {
		pluginTag = "pcm-plugin"
	}

	baseURL := &url.URL{
		Scheme: "http",
		Host:   kongHost,
	}

	client := kong.NewClient(baseURL)
	kongRoutes, _, err := client.ListRoute(pluginTag)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Error sending request")
		log.Printf("%s\t%s\t%v\t%s", "GET", ctx.Request.Host+ctx.Request.URL.Path, http.StatusInternalServerError, err)
	}

	kongServices, _, err := client.ListServices()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Error sending request")
		log.Printf("%s\t%s\t%v\t%s", "GET", ctx.Request.Host+ctx.Request.URL.Path, http.StatusInternalServerError, err)
	}

	plugins := make([]PluginInfo, 0)
	for _, route := range kongRoutes {
		if route.Service != nil {
			plugin := PluginInfo{
				Route: route.Name,
				URL:   route.Paths[0],
			}

			for _, service := range kongServices {
				if route.Service.ID == service.ID {
					plugin.Name = service.Name
				}
			}
			plugins = append(plugins, plugin)
		}
	}

	ctx.JSON(200, plugins)

}
