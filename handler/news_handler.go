package handler

import (
	"github.com/gin-gonic/gin"
	"inspiredlab/domain"
	"log"
	"net/http"
)

type NewsHandler struct {
	newsService domain.NewsService
}

type CreateNewsRequest struct {
	Name         string `json:"name"`
	ThumbnailURL string `json:"thumbnailURL"`
	Content      string `json:"content"`
	Tags         string `json:"tags"`
}

type CreateNewsResponse struct {
	ID string `json:"id"`
}

type UpdateNewsRequest struct {
	Name         string `json:"name"`
	ThumbnailURL string `json:"thumbnailURL"`
	Content      string `json:"content"`
	Tags         string `json:"tags"`
}

type GetNewsResponse struct {
	Name         string `json:"name"`
	ThumbnailURL string `json:"thumbnailURL"`
	Content      string `json:"content"`
	Tags         string `json:"tags"`
	CreatedTime  string `json:"createdTime"`
	UpdatedTime  string `json:"updatedTime"`
}

func NewNewsHandler(r *gin.RouterGroup, service domain.NewsService) {
	handler := &NewsHandler{newsService: service}
	r.POST("/create", handler.CreateItem)
	r.PUT("/:id", handler.UpdateItem)
	r.DELETE("/:id", handler.DeleteItem)
	r.GET("/:id", handler.GetItem)
}

func (h *NewsHandler) GetItem(c *gin.Context) {
	id := c.Param("id")
	newsItem, err := h.newsService.Get(id)

	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, GetNewsResponse{
			Name:         newsItem.Name,
			ThumbnailURL: newsItem.ThumbnailURL,
			Content:      newsItem.Content,
			Tags:         newsItem.Tags,
			CreatedTime:  newsItem.CreatedAt.String(),
			UpdatedTime:  newsItem.UpdatedAt.String(),
		})
	}
}

func (h *NewsHandler) CreateItem(c *gin.Context) {
	var req UpdateNewsRequest
	if err := c.BindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
	}

	created, err := h.newsService.Create(domain.News{
		Name:         req.Name,
		ThumbnailURL: req.ThumbnailURL,
		Content:      req.Content,
		Tags:         req.Tags,
	})
	if err != nil {
		return
	}

	if err != nil {
		c.Status(http.StatusInternalServerError)
		log.Println(err.Error())
	} else {
		c.JSON(http.StatusOK, CreateNewsResponse{
			ID: created.ID,
		})
	}

}

func (h *NewsHandler) UpdateItem(c *gin.Context) {
	id := c.Param("id")
	var req UpdateNewsRequest
	if err := c.BindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
	}

	err := h.newsService.Update(id, domain.News{
		Name:         req.Name,
		ThumbnailURL: req.ThumbnailURL,
		Content:      req.Content,
		Tags:         req.Tags,
	})
	if err != nil {
		return
	}

	if err != nil {
		c.Status(http.StatusInternalServerError)
		log.Println(err.Error())
	} else {
		c.Status(http.StatusOK)
	}

}

func (h *NewsHandler) DeleteItem(c *gin.Context) {
	id := c.Param("id")

	err := h.newsService.Delete(id)
	if err != nil {
		return
	}

	if err != nil {
		c.Status(http.StatusInternalServerError)
		log.Println(err.Error())
	} else {
		c.Status(http.StatusOK)
	}
}
