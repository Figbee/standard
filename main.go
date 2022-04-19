package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"standard/cmd"
	"standard/internal/global"
	"standard/internal/initialize"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	//初始化配置项组件
	initialize.InitConfig()
	//初始化日志记录组件
	initialize.InitLogger()
	//初始化orm成功
	initialize.InitMysql()
	//初始化casbin
	initialize.InitCasbin()
	//初始化验证器
	initialize.InitValidate()
	//初始化路由
	r := initialize.Routers()
	//执行命令行
	cmd.Execute()
	startHttp(r)

}

//优雅的重启服务器
func startHttp(r *gin.Engine) {
	host := global.Conf.System.Host
	port := global.Conf.System.Port
	srv := http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Error("listen error:", err)
		}
	}()
	global.Logger.Info(fmt.Sprintf("server is rinning at %s:%d", host, port))
	//创建通道
	quit := make(chan os.Signal)
	//signal.Notify 监听和捕获信号量
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.Logger.Info("shutting down server...")
	//控制goRouting的生命周期
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.Logger.Error("Server forced to shutdown:", err)
	}
	global.Logger.Info("Server exiting")
}
