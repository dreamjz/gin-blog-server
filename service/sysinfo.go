package service

import (
	"errors"
	"gin-blog-server/models"
	"gin-blog-server/utils"
	"log"
)

var (
	ErrGetSysInfo = errors.New("get system info error")
)

func SystemInfo() (*models.Server, error) {
	os := utils.OSInfo()
	cpu, err := utils.CPUInfo()
	if err != nil {
		log.Print("Get CPU info error: ", err.Error())
		return nil, ErrGetSysInfo
	}
	ram, err := utils.MemInfo()
	if err != nil {
		log.Print("Get RAM info error: ", err.Error())
		return nil, ErrGetSysInfo
	}
	disk, err := utils.DiskInfo()
	if err != nil {
		log.Print("Get DISK info error: ", err.Error())
		return nil, ErrGetSysInfo
	}
	sys := models.Server{
		Os:   os,
		Cpu:  cpu,
		Ram:  ram,
		Disk: disk,
	}
	return &sys, nil
}
