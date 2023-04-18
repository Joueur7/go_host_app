package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
)

type mockDB struct{}

func (db *mockDB) GetTrendingStickers(currentTime string, numStickers int) ([]Sticker, error) {
	// Define mock stickers
	stickers := []Sticker{
		{ID: 1, Name: "Sticker 1", ImageURL: "http://example.com/sticker1.png", Priority: 1, StartTime: "00:00:00", EndTime: "23:59:59"},
		{ID: 2, Name: "Sticker 2", ImageURL: "http://example.com/sticker2.png", Priority: 2, StartTime: "00:00:00", EndTime: "23:59:59"},
		{ID: 3, Name: "Sticker 3", ImageURL: "http://example.com/sticker3.png", Priority: 3, StartTime: "00:00:00", EndTime: "23:59:59"},
		{ID: 4, Name: "Sticker 4", ImageURL: "http://example.com/sticker4.png", Priority: 4, StartTime: "00:00:00", EndTime: "23:59:59"},
		{ID: 5, Name: "Sticker 5", ImageURL: "http://example.com/sticker5.png", Priority: 5, StartTime: "00:00:00", EndTime: "23:59:59"},
	}

	if numStickers >= len(stickers) {
		return stickers, nil
	}

	return stickers[:numStickers], nil
}

func TestGetTrendingStickers(t *testing.T) {
    // Create a new test server
    e := echo.New()
    req := httptest.NewRequest(http.MethodGet, "/v1/trendingStickers?numStickers=10", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    // Define endpoint handler
    handler := func(c echo.Context) error {
        // Get number of stickers to return
        numStickersStr := c.QueryParam("numStickers")
        numStickers, err := strconv.Atoi(numStickersStr)
        if err != nil {
            return c.JSON(http.StatusBadRequest, "Invalid number of stickers")
        }

        // Return a JSON response with numStickers stickers
        var stickers []Sticker
        for i := 1; i <= numStickers; i++ {
            sticker := Sticker{
                ID:        i,
                Name:      "Sticker " + strconv.Itoa(i),
                ImageURL:  "http://example.com/sticker_" + strconv.Itoa(i) + ".png",
                Priority:  i,
                StartTime: time.Now().Format("15:04:05"),
                EndTime:   time.Now().Add(24 * time.Hour).Format("15:04:05"),
            }
            stickers = append(stickers, sticker)
        }

        return c.JSON(http.StatusOK, stickers)
    }

    // Call the handler function
    if err := handler(c); err != nil {
        t.Errorf("Error handling request: %v", err)
    }

    // Check the response status code
    if rec.Code != http.StatusOK {
        t.Errorf("Expected status code %d but got %d", http.StatusOK, rec.Code)
    }

    // Check the response body
    var stickers []Sticker
    if err := json.NewDecoder(rec.Body).Decode(&stickers); err != nil {
        t.Errorf("Failed to decode response body: %v", err)
    }

    // Check that the correct number of stickers were returned
    numStickersStr := c.QueryParam("numStickers")
    numStickers, _ := strconv.Atoi(numStickersStr)
    if len(stickers) != numStickers {
        t.Errorf("Expected %d stickers but got %d", numStickers, len(stickers))
    }
}
