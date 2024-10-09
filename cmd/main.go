package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware" // นำเข้า Echo middleware
	router "restaurant/internal/router"
)

func main() {
	e := echo.New()

	// กำหนด CORS middleware อนุญาตทุกโดเมนและทุกวิธีการ
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // อนุญาตทุกโดเมน
		AllowMethods: []string{"*"}, // อนุญาตทุกวิธีการ
		AllowHeaders: []string{"*"}, // อนุญาตทุกหัวข้อ
	}))

	// กำหนด router
	router.CustomerRouter(e)
	router.FoodRouter(e)
	router.OrderRouter(e)
	router.EmployeeRouter(e)

	// เริ่มเซิร์ฟเวอร์
	e.Logger.Fatal(e.Start(":8080"))
}
