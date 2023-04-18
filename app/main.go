package main

import (
    "log"
    "net/http"
    "strconv"
    "time"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/spf13/viper"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "Project/sticker/repository"
    "fmt"
)

type Sticker struct {
    ID        int    `json:"id"`
    Name      string `json:"name"`
    ImageURL  string `json:"imageUrl"`
    Priority  int    `json:"priority"`
    StartTime string `json:"startTime"`
    EndTime   string `json:"endTime"`
}

func main() {
    viper.AddConfigPath("../config")
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    err := viper.ReadInConfig()
    if err != nil {
        log.Fatalf("Failed to read configuration: %v", err)
    }

    // Connect to MySQL database using GORM
    user := viper.GetString("db.user")
    password := viper.GetString("db.password")
    host := viper.GetString("db.host")
    port := viper.GetString("db.port")
    dbName := viper.GetString("db.name")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)

    dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }
    sqlDB, err := dbConn.DB()
    if err != nil {
        log.Fatal(err)
    }
    defer sqlDB.Close()

    // Auto-migrate the database
    err = dbConn.AutoMigrate(&Sticker{})
    if err != nil {
        log.Fatalf("Failed to auto-migrate the database: %v", err)
    }

    // Create table if it doesn't exist
    migrator := dbConn.Migrator()
    if !migrator.HasTable(&Sticker{}) {
        err = migrator.CreateTable(&Sticker{})
        if err != nil {
            log.Fatalf("Failed to create table: %v", err)
        }
    }


    // Create StickerRepository instance
    stickerRepo := repository.NewMySQLStickerRepository(dbConn)

    // Set up Echo web framework
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Define endpoint handler
    e.GET("/v1/trendingStickers", func(c echo.Context) error {
        // Get number of stickers to return
        numStickersStr := c.QueryParam("numStickers")
        numStickers, err := strconv.Atoi(numStickersStr)
        if err != nil {
            return c.JSON(http.StatusBadRequest, "Invalid number of stickers")
        }

        // Get current time
        currentTime := time.Now().Format("15:04:05")

        // Query database for trending stickers
        stickers, err := stickerRepo.GetTrendingStickers(currentTime, numStickers)
        if err != nil {
            log.Printf("Failed to fetch stickers from database: %v", err)
            return c.JSON(http.StatusInternalServerError, "Failed to fetch stickers from database")
        }

        return c.JSON(http.StatusOK, stickers)
    })

    // Start server
    localPort := viper.GetString("server.localPort")
    e.Logger.Fatal(e.Start(localPort))
}
