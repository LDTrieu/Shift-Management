package main

import(
	"os"
	"fmt"
	"log"
	"github.com/joho/godotenv"
	//"database/sql"
	

	"github.com/ldtrieu/staffany-backend/pkg/date"
	"github.com/ldtrieu/staffany-backend/pkg/shift"
	"github.com/ldtrieu/staffany-backend/pkg/user"
	"github.com/ldtrieu/staffany-backend/pkg/week"
	//"github.com/ldtrieu/staffany-backend"
	
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	
)


func main() {
	fmt.Println("Hello, World!")
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	
	dsn := os.Getenv("DB_DSN")
	//fmt.Println("Home: ", os.Getenv("DB_DSN"))
	//dsn := "root:123456aH@@tcp(mysql:3306)/staffany?charset=utf8mb4&parseTime=True&loc=Local"
	//fmt.Println(err)
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//db, err := gorm.Open("mysql","root:123456aH@@tcp(mysql:3306)/staffany?charset=utf8mb4&parseTime=True&loc=Local")

	//db, err := sql.Open("mysql","root:password@tcp(localhost:3306)/staffany")
	//db, err := sql.Open("mysql", "root:123456aH@@tcp(127.0.0.1:3306)/staffany")
	//db, err := gorm.Open("mysql", "root:123456aH@@tcp(mysql:3306)/staffany?charset=utf8&parseTime=True&loc=Local")
	//db, err := gorm.Open("mysql", "root:123456aH@@/staffany?charset=utf8&parseTime=True&loc=Local")
	
	if err != nil {
		panic("failed to connect database") // Kiểm tra kết nối tới database
	  }
	  
	//defer db.Close() // Để đóng cơ sở dữ liệu khi nó không được sử dụng
	//fmt.Println(db)
	/*
	if err != nil {
		log.Fatal(err)
		//panic(err)
	}
	*/
	
	
	// bad, never do this in production
	
	db.AutoMigrate(&user.User{}, &week.Week{}, &date.Date{}, &shift.Shift{})
	r := gin.Default()
	r.Use(cors.Default())

	userRepository := user.NewRepository(db)
	weekRepository := week.NewRepository(db)
	dateRepository := date.NewRepository(db)
	shiftRepository := shift.NewRepository(db)

	userservice := user.NewService(userRepository)
	weekService := week.NewService(weekRepository, dateRepository, shiftRepository)
	shiftService := shift.NewService(shiftRepository)

	v1 := r.Group("/api/v1")
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	userservice.Route(v1)
	weekService.Route(v1)
	shiftService.Route(v1)

	r.Run(":8080")
	
	
}
