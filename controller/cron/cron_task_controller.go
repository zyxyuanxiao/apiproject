package c_cron

import (
	"apiproject/log"
	m_cron "apiproject/model/cron"
	s_cron "apiproject/service/cron"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

/**
添加定时任务
*/
func AddCronTask(ctx *gin.Context) {
	para := m_cron.CronTask{}
	if err := ctx.ShouldBind(&para); err != nil {
		log.Logger.Error("绑定请求参数到对象异常", zap.Error(err))
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": nil,
			"msg":  err.Error(),
		})
		return
	}
	log.Logger.Info("绑定请求参数到对象", zap.Any("para", para))

	para.Status = 1
	if err := s_cron.CronTaskService.AddCronTask(&para); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": nil,
			"msg":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": nil,
	})
}

/**
删除定时任务
*/
func DeleteCronTask(ctx *gin.Context) {
	para := m_cron.CronTask{}
	if err := ctx.ShouldBind(&para); err != nil {
		log.Logger.Error("绑定请求参数到对象异常", zap.Error(err))
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": nil,
			"msg":  err.Error(),
		})
		return
	}
	log.Logger.Info("绑定请求参数到对象", zap.Any("para", para))

	if err := s_cron.CronTaskService.DeleteCronTask(para.ID); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": nil,
			"msg":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": nil,
	})
}

/**
启用定时任务
*/
func EnableCronTask(ctx *gin.Context) {
	para := m_cron.CronTask{}
	if err := ctx.ShouldBind(&para); err != nil {
		log.Logger.Error("绑定请求参数到对象异常", zap.Error(err))
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": nil,
			"msg":  err.Error(),
		})
		return
	}
	log.Logger.Info("绑定请求参数到对象", zap.Any("para", para))

	if err := s_cron.CronTaskService.EnableCronTask(para.ID); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": nil,
			"msg":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": nil,
	})
}

/**
禁用定时任务
*/
func DisableCronTask(ctx *gin.Context) {
	para := m_cron.CronTask{}
	if err := ctx.ShouldBind(&para); err != nil {
		log.Logger.Error("绑定请求参数到对象异常", zap.Error(err))
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": nil,
			"msg":  err.Error(),
		})
		return
	}
	log.Logger.Info("绑定请求参数到对象", zap.Any("para", para))

	if err := s_cron.CronTaskService.DisableCronTask(para.ID); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": nil,
			"msg":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": nil,
	})
}

/**
查询定时任务列表
*/
func FindCronTaskList(ctx *gin.Context) {
	cronTaskList, err := s_cron.CronTaskService.FindCronTaskList()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": nil,
			"msg":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": cronTaskList,
	})
}
