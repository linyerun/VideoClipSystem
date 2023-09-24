package controller

import (
	"VideoClipSystem/app/global"
	"VideoClipSystem/app/resp_msg"
	"VideoClipSystem/app/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

func InitGin() {
	//初始化路由
	r := gin.Default()

	// 添加cors中间件
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))

	//gin整合pprof
	pprof.Register(r)

	//这个是形成swagger文档的命脉所在,位置在"/docs/swagger.json"里面
	url := ginSwagger.URL("/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// 路由分组
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register/", register)
		userGroup.POST("/login/", login)
	}

	emailGroup := r.Group("/email")
	{
		emailGroup.POST("/authCode/", sendAuthCodeByEmail)
	}

	videoGroup := r.Group("/video")
	videoGroup.Use(tokenInterceptor) // 处理token
	{
		videoGroup.POST("/clip/", clipVideo)
		videoGroup.GET("/progress/", clipVideoProgress)
	}

	fileGroup := r.Group("/file")
	{
		fileGroup.GET("/video/:clippedVideoId/", clippedVideoFileReturn)
	}

	// 运行Gin
	if err := r.Run(":" + global.HostPort); err != nil {
		panic(err)
	}
}

func tokenInterceptor(c *gin.Context) {
	// 获取token
	token := c.GetHeader("token")

	// 校验token
	claims, err := utils.ParseToken(token)
	if err != nil {
		code := resp_msg.ErrorAuthCheckTokenFail
		c.JSON(code, resp_msg.NewRespMsg(code, nil))
		return
	}

	if utils.IsTimeoutCheckoutByMilli(claims.ExpiresAt, int64(global.ExpirationSecond)*1000, time.Now().UnixMilli()) {
		code := resp_msg.ErrorAuthCheckTokenTimeout
		c.JSON(code, resp_msg.NewRespMsg(code, nil))
		return
	}

	// 把解析处理的信息加入c中
	c.Set("userId", claims.Id)
	c.Set("userEmail", claims.Username)

	c.Next()

	// 更新token值并装入请求头中
	token, err = utils.GenerateToken(claims.Id, claims.Username, 0)
	if err != nil {
		utils.Logger().Error(err)
		return
	}
	c.Request.Response.Header.Add("token", token)
}
