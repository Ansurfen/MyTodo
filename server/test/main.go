package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.Static("/", `C:\Users\LENOVO\StudioProjects\my_todo\build\web`)
	r.Run(":7777")
}
