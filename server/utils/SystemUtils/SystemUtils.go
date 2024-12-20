/*
 * @Descripttion: 系统工具 - System Utils
 */
package SystemUtils

import (
	"VideoHubGo/models/SystemModel"
	"VideoHubGo/utils/LogUtils"
	"bytes"
	"os/exec"
	"runtime"

	"github.com/shirou/gopsutil/disk"
)

/**
 * @Descripttion: 获取当前操作系统环境 - Get current operating system environment
 * @Return: System (int)
 */
func GetSystemInfo() int {
	systemType := runtime.GOOS
	if systemType == "linux" {
		return 1
	} else if systemType == "windows" {
		return 2
	} else {
		return 3
	}
}

/**
 * @Descripttion: 执行Linux命令 - Execute linux commands
 * @Param: cmd (string)
 * @Return: result (string)
 */
func RunLinuxCommand(cmd string) (string, error) {
	order := exec.Command("/bin/bash", "-c", cmd)
	var out bytes.Buffer
	order.Stdout = &out
	errRun := order.Run()
	if errRun != nil {
		LogUtils.Logger("[致命错误 ErrorRun]执行Linux命令-Run步骤中：" + cmd + "中异常：" + errRun.Error())
		return "", errRun
	}
	return out.String(), errRun
}

/**
 * @Descripttion: 执行Windows命令 - Execute windows commands
 * @Param: cmd (string)
 * @Return: result (string)
 */
func RunWindowsCommand(cmd string) (string, error) {
	order := exec.Command("cmd", "/c", cmd)
	var out bytes.Buffer
	order.Stdout = &out
	errRun := order.Run()
	if errRun != nil {
		LogUtils.Logger("[致命错误 ErrorRun]执行Windows命令-Run步骤中：" + cmd + "中异常：" + errRun.Error())
		return "", errRun
	}
	return out.String(), errRun
}

/**
 * @Descripttion: 获取磁盘信息 - Get Disk Info
 * @Param: filePath (string)
 * @Return: disk.UsageStat
 */
func GetDiskInfo(filePath string) (SystemModel.UsageStat, error) {
	info, err := disk.Usage(filePath)
	var ret SystemModel.UsageStat
	ret.Path = info.Path
	ret.Fstype = info.Fstype
	ret.Total = info.Total
	ret.Free = info.Free
	ret.Used = info.Used
	ret.UsedPercent = info.UsedPercent
	ret.InodesTotal = info.InodesTotal
	ret.InodesUsed = info.InodesUsed
	ret.InodesFree = info.InodesFree
	ret.InodesUsedPercent = info.InodesUsedPercent
	return ret, err
}
