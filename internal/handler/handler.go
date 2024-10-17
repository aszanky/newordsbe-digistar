package handler

import (
	"net/http"
	"time"

	"github.com/aszanky/newordsbe-digistar/internal/models"
	"github.com/aszanky/newordsbe-digistar/internal/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	usecase  usecase.Usecase
	router   *gin.Engine
	validate *validator.Validate
}

func NewHandler(
	uc usecase.Usecase,
) *Handler {
	r := gin.Default()
	val := validator.New()

	return &Handler{
		usecase:  uc,
		router:   r,
		validate: val,
	}
}

func (h *Handler) AddNewWord(c *gin.Context) {
	var input models.Word

	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.validate.Struct(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.usecase.AddNewWords(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"word": input,
	})
}

func (h *Handler) GetListWord(c *gin.Context) {

	words, err := h.usecase.GetListWords()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": words,
	})
}

func (r *Handler) Register() {
	r.router.GET("/words/get", r.GetListWord)
	r.router.POST("/words/add", r.AddNewWord)
}

// Start HTTP Server
func (r *Handler) Start(port string) error {
	// url := host + port

	//solve issue You trusted all proxies, this is NOT safe. We recommend you to set a value.
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	r.router.Use(cors.New(config))

	r.router.ForwardedByClientIP = true
	// r.router.SetTrustedProxies([]string{"0.0.0.0:6010"})
	r.Register()

	return r.router.Run(port)
}
