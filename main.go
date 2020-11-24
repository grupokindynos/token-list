package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

type ListVersion struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Patch int `json:"patch"`
}

type ListToken struct {
	ChainID  int      `json:"chainId"`
	Address  string   `json:"address"`
	Symbol   string   `json:"symbol"`
	Name     string   `json:"name"`
	Decimals int      `json:"decimals"`
	LogoURI  string   `json:"logoURI"`
	Tags     []string `json:"tags"`
}

type List struct {
	Name      string      `json:"name"`
	LogoURI   string      `json:"logoURI"`
	Keywords  []string    `json:"keywords"`
	Timestamp time.Time   `json:"timestamp"`
	Tokens    []ListToken `json:"tokens"`
	Version   ListVersion `json:"version"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	App := GetApp()
	err := App.Run(":" + port)
	if err != nil {
		panic(err)
	}
}

// GetApp is used to wrap all the additions to the GIN API.
func GetApp() *gin.Engine {
	App := gin.Default()
	App.Use(cors.Default())
	ApplyRoutes(App)
	return App
}

func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/")
	{
		api.GET("list", list)
	}
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})
}

func list(c *gin.Context) {
	list := List{
		Name:     "Kindynos Token List",
		LogoURI:  "https://raw.githubusercontent.com/grupokindynos/kindynos-branding/master/kindynos/kindynos-isotype.png",
		Keywords: []string{"defi", "tokens", "erc20", "kindynos"},
		Timestamp: time.Now(),
		Tokens: []ListToken{
			{
				ChainID: 1,
				Address: "0x71c10040Db9B5Bf6d5Bdd81DAC26B47bb44eB4f0",
				Symbol: "WPOLIS",
				Name:"WrappedPolis",
				Decimals: 18,
				LogoURI: "https://raw.githubusercontent.com/grupokindynos/kindynos-branding/master/polis/polis-isotype-background-round.png",
				Tags: []string{"polis", "erc20", "defi"},
			},
		},
		Version:  ListVersion{Major: 1, Minor: 0, Patch: 0},
	}
	c.JSON(200, list)
	return
}
