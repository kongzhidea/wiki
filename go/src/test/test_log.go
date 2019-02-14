package main

import (
	"flag"
	"github.com/golang/glog"
)

// glog中 Stderr 修改为 Stdout（默认为error）: init方法中  修改为logging.stderrThreshold = infoLog

// logtostderr    true 仅打印到控制台（默认），false：打印到控制台和文件，
// stderrthreshold 修改成了 默认info级别，超过info级别的日志 打印到控制台
// formatHeader() 修改为 INFO 2017-06-01 15:57:20 test_log.go:22 - message  格式
// warning -> warn， 取消 warnf等操作

// SetLogging**() 设置打印参数，如logtostderr等， 覆盖命令行参数

// glog.V(level)  level 小于等于 -v参数时候 才可以打印日志

// info中会写 warn和error日志
// 不使用 glog.V(level)，强制打印日志

func main() {
	flag.Parse() // 解析命令行参数， 如 -log_dir=/data/logs/go，-logtostderr=false

	//退出时调用，确保日志写入文件中
	defer glog.Flush()

	glog.Info("hello, glog")
	glog.Warning("warning glog")
	glog.Error("error glog")

	glog.Info("info %d", 1)
	glog.Warning("warning %d", 2)
	glog.Error("error %d", 3)



}
