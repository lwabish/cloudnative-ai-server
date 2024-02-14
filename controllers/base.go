package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/lwabish/cloudnative-ai-server/models"
	"github.com/lwabish/cloudnative-ai-server/utils"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
)

var (
	BaseCtl = &BaseController{}
)

type BaseController struct {
	DB *gorm.DB
	Q  *utils.TaskQueue
	L  *logrus.Logger
	C  *kubernetes.Clientset
}
type BaseControllerCfg struct {
	DB *gorm.DB
	Q  *utils.TaskQueue
	L  *logrus.Logger
	C  *kubernetes.Clientset
}

func (b *BaseController) Setup(cfg *BaseControllerCfg) {
	b.DB = cfg.DB
	b.Q = cfg.Q
	b.L = cfg.L
	b.C = cfg.C
}

func (b *BaseController) UpdateTaskStatus(uid string, status models.TaskStatus) {
	r := b.DB.
		Model(&models.Task{}).
		Where("uid = ?", uid).
		Update("status", status)
	if r.Error != nil || r.RowsAffected != 1 {
		b.L.Warnf("update task status error: %v", r.Error)
	}
}
