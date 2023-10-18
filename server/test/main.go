// package main

// import "github.com/gin-gonic/gin"

// type Test struct {
// 	Port *int `yaml:"port"`
// }

//	func main() {
//		r := gin.Default()
//		r.Static("/", `C:\Users\LENOVO\StudioProjects\my_todo\build\web`)
//		r.Run(":7777")
//	}
package main

import (
	"context"
	"fmt"
	"net/http"

	rkboot "github.com/rookie-ninja/rk-boot/v2"
	rkgrpc "github.com/rookie-ninja/rk-grpc/v2/boot"
)

func main() {
	// Create a new boot instance.
	boot := rkboot.NewBoot()

	// Bootstrap
	boot.Bootstrap(context.Background())

	// Get grpc entry with name
	grpcEntry := rkgrpc.GetGrpcEntry("greeter")

	// Attachment upload from http/s handled manually
	grpcEntry.GwMux.HandlePath("POST", "/v1/files", handleBinaryFileUpload)

	// Wait for shutdown sig
	boot.WaitForShutdownSig(context.Background())
}

func handleBinaryFileUpload(w http.ResponseWriter, req *http.Request, params map[string]string) {
	err := req.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse form: %s", err.Error()), http.StatusBadRequest)
		return
	}

	f, header, err := req.FormFile("attachment")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get file 'attachment': %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer f.Close()

	fmt.Println(header)

	//
	// Now do something with the io.Reader in `f`, i.e. read it into a buffer or stream it to a gRPC client side stream.
	// Also `header` will contain the filename, size etc of the original file.
	//
}
