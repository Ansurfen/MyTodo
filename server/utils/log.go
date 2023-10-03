package utils

import (
	"os"
	"path/filepath"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const DefaultLogPath = "./log"

type LoggerOpt struct {
	Level          string `yaml:"level"`
	Format         string `yaml:"format"`
	Path           string `yaml:"path"`
	FileName       string `yaml:"filename"`
	FileMaxSize    int    `yaml:"fileMaxSize"`
	FileMaxBackups int    `yaml:"fileMaxBackups"`
	MaxAge         int    `yaml:"maxAge"`
	Compress       bool   `yaml:"compress"`
	Stdout         bool   `yaml:"stdout"`
}

func init() {
	err := InitLogger(LoggerOpt{
		Level:    "info",
		Stdout:   true,
		Path:     DefaultLogPath,
		FileName: "todo.log",
		Format:   "json",
	})
	if err != nil {
		panic(err)
	}
}

func InitLogger(conf LoggerOpt) error {
	logLevel := map[string]zapcore.Level{
		"debug": zapcore.DebugLevel,
		"info":  zapcore.InfoLevel,
		"warn":  zapcore.WarnLevel,
		"error": zapcore.ErrorLevel,
	}
	writeSyncer, err := getLogWriter(conf) // 日志文件配置 文件位置和切割
	if err != nil {
		return err
	}
	encoder := getEncoder(conf)       // 获取日志输出编码
	level, ok := logLevel[conf.Level] // 日志打印级别
	if !ok {
		level = logLevel["info"]
	}
	core := zapcore.NewCore(encoder, writeSyncer, level)
	logger := zap.New(core, zap.AddCaller()) // zap.Addcaller() 输出日志打印文件和行数如： logger/logger_test.go:33
	// 1. zap.ReplaceGlobals 函数将当前初始化的 logger 替换到全局的 logger,
	// 2. 使用 logger 的时候 直接通过 zap.S().Debugf("xxx") or zap.L().Debug("xxx")
	// 3. 使用 zap.S() 和 zap.L() 提供全局锁，保证一个全局的安全访问logger的方式
	zap.ReplaceGlobals(logger)
	return nil
}

// getEncoder 编码器(如何写入日志)
func getEncoder(conf LoggerOpt) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // log 时间格式 例如: 2021-09-11t20:05:54.852+0800
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	// encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	//encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	if conf.Format == "json" {
		return zapcore.NewJSONEncoder(encoderConfig) // 以json格式写入
	}
	return zapcore.NewConsoleEncoder(encoderConfig) // 以logfmt格式写入
}

// getLogWriter 获取日志输出方式  日志文件 控制台
func getLogWriter(conf LoggerOpt) (zapcore.WriteSyncer, error) {

	// 判断日志路径是否存在，如果不存在就创建
	if exist := IsExist(conf.Path); !exist {
		if conf.Path == "" {
			conf.Path = DefaultLogPath
		}
		if err := os.MkdirAll(conf.Path, os.ModePerm); err != nil {
			conf.Path = DefaultLogPath
			if err := os.MkdirAll(conf.Path, os.ModePerm); err != nil {
				return nil, err
			}
		}
	}

	// 日志文件 与 日志切割 配置
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filepath.Join(conf.Path, conf.FileName), // 日志文件路径
		MaxSize:    conf.FileMaxSize,                        // 单个日志文件最大多少 mb
		MaxBackups: conf.FileMaxBackups,                     // 日志备份数量
		MaxAge:     conf.MaxAge,                             // 日志最长保留时间
		Compress:   conf.Compress,                           // 是否压缩日志
	}
	if conf.Stdout {
		// 日志同时输出到控制台和日志文件中
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger), zapcore.AddSync(os.Stdout)), nil
	} else {
		// 日志只输出到日志文件
		return zapcore.AddSync(lumberJackLogger), nil
	}
}
