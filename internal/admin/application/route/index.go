package route

import (
	resp "camellia/tool/response"
	"github.com/gin-gonic/gin"
)

func SetupRouter() (e *gin.Engine, err error) {

	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	front := r.Group("/front")
	{
		front.GET("ping", func(context *gin.Context) {
			resp.Success(context, resp.Ok, "pong")
		})

	}

	backend := r.Group("/backend")
	//backend.Use(middleware.AccessTokenMiddleware())
	{
		backend.GET("ping", func(context *gin.Context) {
			resp.Success(context, resp.Ok, "pong")
		})
		BackendGroup(backend)
	}

	//uploadInstance, err := uploader.New(r, uploader.GatherConfig{
	//	Path:      fs.StoragePath() + "upload",
	//	UrlPrefix: "/common",
	//	File: uploader.FileConfig{
	//		Path:      "files",
	//		MaxSize:   10485760,
	//		AllowType: []string{".xls", ".txt"},
	//	},
	//	Image: uploader.ImageConfig{
	//		Path:    "images",
	//		MaxSize: 10485760,
	//		Thumb: uploader.ThumbConfig{
	//			Path:      "thumb",
	//			MaxWidth:  300,
	//			MaxHeight: 300,
	//		},
	//	},
	//})
	//
	//if err != nil {
	//	return &gin.Engine{}, err
	//}
	//
	//uploadInstance.Resolve()

	return r, nil
}

