package api

import (
	"net/http"
	"strconv"

	"youtube-studio-v2/config"
	"youtube-studio-v2/db"

	"github.com/labstack/echo/v4"
)

func GetVideosHandler(c echo.Context) error {
	ctx := c.Request().Context()
	logger := config.GetLogger()

	limitParam := c.QueryParam("limit")
	offsetParam := c.QueryParam("offset")

	// default values
	limit := 10
	offset := 0

	if limitParam != "" {
		limitValue, err := strconv.Atoi(limitParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid value for 'limit' parameter"})
		}
		limit = limitValue
	}

	if offsetParam != "" {
		offsetValue, err := strconv.Atoi(offsetParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid value for 'offset' parameter"})
		}

		offset = offsetValue
	}

	videos, err := db.GetPaginatedVideos(ctx, limit, offset)
	if err != nil {
		logger.Error("failed to fetch videos from db, err: %w", err)
		return c.JSON(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusOK, videos)
}

func SearchVideosHandler(c echo.Context) error {
	ctx := c.Request().Context()
	logger := config.GetLogger()

	query := c.QueryParam("query")
	if query == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Query parameter 'query' is required"})
	}

	data, err := db.SearchVideos(ctx, query)
	if err != nil {
		logger.Error("failed to fetch videos from db, err: %w", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Something went wrong"})
	}

	return c.JSON(http.StatusOK, data)
}
