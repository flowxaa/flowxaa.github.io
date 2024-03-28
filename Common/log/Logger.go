package logger

import (
	"bytes"
	"container/list"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

type Logger struct {
	level_    int
	type_     int
	filename_ string
	logList_  *list.List
	mutex_    sync.Mutex
	stop_     bool
	rouse_    chan int
}

type tgLogUnit struct {
	FileName string
	FuncName string
	Line     int
	Level    int
	LogStr   string
	AddTime  string
}

const (
	LOG_LEVEL_DEBUG   = iota //调试级别
	LOG_LEVEL_INFO           //信息级别
	LOG_LEVEL_WARNING        //警告基本
	LOG_LEVEL_ERROR          //错误级别
	LOG_LEVEL_NONE           //不记录
)

const (
	LOG_TYPE_NONE   = 0 //不记录
	LOG_TYPE_STDOUT = 1 //标准输出
	LOG_TYPE_FILE   = 2 //写文件
)

const (
	LOG_BUFFER_MAXSIZE = 4096
	LOG_FILE_MAXSIZE   = 300 * 1024 * 1024 //300M
	LOG_SAVE_MAXSIZE   = 7
)

var (
	slogger        = &Logger{type_: LOG_TYPE_STDOUT}
	sLogLevelstr   = [4]string{"DEBUG", "INFO", "WARNING", "ERROR"}
	sLogPrintColor = [4]int{38, 36, 33, 31}
)

func InitLog(level, ntype int, filename string) {
	slogger = &Logger{
		level_:    level,
		type_:     ntype,
		stop_:     false,
		rouse_:    make(chan int),
		logList_:  list.New(),
		filename_: filename,
	}
	go loggerWork()
}

func loggerWork() {
	ticker := time.NewTicker(3 * time.Second)
	var bf bytes.Buffer
	for {
		var ok bool = false
		select {
		case <-slogger.rouse_:
			ok = true
		case <-ticker.C:
		}
		if slogger.stop_ {
			break
		}
		if ok {
			for {
				unit, ok := slogger.Dequeue()
				if ok {
					bf.WriteString(fmt.Sprintf("[%s] [%s] %s(%d), %s: %s\n", unit.AddTime, sLogLevelstr[unit.Level], unit.FileName, unit.Line, unit.FuncName, unit.LogStr))
					if bf.Len() > LOG_BUFFER_MAXSIZE {
						slogger.WriteFile(slogger.filename_, &bf)
					}
				} else {
					break
				}
			}
		} else {
			ok, size := slogger.GetFileByteSize(slogger.filename_)
			if ok && size > LOG_FILE_MAXSIZE {
				slogger.ReuseFile(slogger.filename_, LOG_SAVE_MAXSIZE)
			}
			slogger.WriteFile(slogger.filename_, &bf)
		}
	}
}

func GetFilenameAndLine() (string, string, int) {
	pc, filename, line, ok := runtime.Caller(2)
	if ok {
		funcname := runtime.FuncForPC(pc).Name()
		funcname = filepath.Ext(funcname)
		funcname = strings.TrimPrefix(funcname, ".")
		filename = filepath.Base(filename)
		return filename, funcname, line
	}
	return "", "", 0
}

//记录调试日志（级别最低）
func LogOfDebug(args ...interface{}) {
	if slogger.level_ <= LOG_LEVEL_DEBUG {
		unit := &tgLogUnit{}
		unit.FileName, unit.FuncName, unit.Line = GetFilenameAndLine()
		t := time.Now()
		unit.AddTime = fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d %03d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.UnixNano()/1e6%1000)
		unit.Level = LOG_LEVEL_DEBUG
		unit.LogStr = fmt.Sprint(args...)
		slogger.FormatWriteLogMsg(unit)
	}
}

//记录信息日志（级别一般）
func LogOfInfo(args ...interface{}) {
	if slogger.level_ <= LOG_LEVEL_INFO {
		unit := &tgLogUnit{}
		unit.FileName, unit.FuncName, unit.Line = GetFilenameAndLine()
		t := time.Now()
		unit.AddTime = fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d %03d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.UnixNano()/1e6%1000)
		unit.Level = LOG_LEVEL_INFO
		unit.LogStr = fmt.Sprint(args...)
		slogger.FormatWriteLogMsg(unit)
	}
}

//记录警告日志（级别较高）
func LogOfWarning(args ...interface{}) {
	if slogger.level_ <= LOG_LEVEL_WARNING {
		unit := &tgLogUnit{}
		unit.FileName, unit.FuncName, unit.Line = GetFilenameAndLine()
		t := time.Now()
		unit.AddTime = fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d %03d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.UnixNano()/1e6%1000)
		unit.Level = LOG_LEVEL_WARNING
		unit.LogStr = fmt.Sprint(args...)
		slogger.FormatWriteLogMsg(unit)
	}
}

//记录错误及异常日志（级别最高）
func LogOfError(args ...interface{}) {
	if slogger.level_ <= LOG_LEVEL_ERROR {
		unit := &tgLogUnit{}
		unit.FileName, unit.FuncName, unit.Line = GetFilenameAndLine()
		t := time.Now()
		unit.AddTime = fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d %03d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.UnixNano()/1e6%1000)
		unit.Level = LOG_LEVEL_ERROR
		unit.LogStr = fmt.Sprint(args...)
		slogger.FormatWriteLogMsg(unit)
	}
}

func (l *Logger) FormatWriteLogMsg(unit *tgLogUnit) {
	if (l.type_ & LOG_TYPE_STDOUT) > 0 {
		fmt.Printf("%c[%dm[%s] [%s] %s(%d), %s: %s%c[0m\n", 0x1B, sLogPrintColor[unit.Level], unit.AddTime, sLogLevelstr[unit.Level], unit.FileName, unit.Line, unit.FuncName, unit.LogStr, 0x1B)
	}
	if (l.type_ & LOG_TYPE_FILE) > 0 {
		l.Enqueue(unit)
	}
}

func (l *Logger) Dequeue() (*tgLogUnit, bool) {
	var unit tgLogUnit

	slogger.mutex_.Lock()
	defer slogger.mutex_.Unlock()
	if l.logList_.Len() > 0 {
		front := slogger.logList_.Front()
		unit = front.Value.(tgLogUnit)
		slogger.logList_.Remove(front)
	} else {
		return nil, false
	}
	return &unit, true
}

func (l *Logger) Enqueue(unit *tgLogUnit) {
	l.mutex_.Lock()
	l.logList_.PushBack(*unit)
	l.mutex_.Unlock()
	select {
	case l.rouse_ <- 1:
	default:
	}
}

func (l *Logger) GetFileByteSize(filename string) (bool, int64) {
	bExist, fhandler := l.IsFile(filename)
	if !bExist {
		return false, 0
	}
	return true, fhandler.Size()
}

func (l *Logger) IsFile(filename string) (bool, os.FileInfo) {
	fhandler, err := os.Stat(filename)
	if !(err == nil || os.IsExist(err)) {
		return false, nil
	} else if fhandler.IsDir() {
		return false, nil
	}
	return true, fhandler
}

func (l *Logger) ReuseFile(filename string, maxsize int) {
	suffix := path.Ext(filename)
	prefix := strings.TrimSuffix(filename, suffix)
	oldest := time.Now().Unix()
	var newpath string
	for i := 1; i <= maxsize; i++ {
		fname := fmt.Sprintf("%s.%d%s", prefix, i, suffix)
		bExist, fhandler := l.IsFile(fname)
		if !bExist {
			os.Rename(filename, fname)
			return
		}
		if oldest > fhandler.ModTime().Unix() {
			oldest = fhandler.ModTime().Unix()
			newpath = fname
		}
	}
	os.Rename(filename, newpath)
}

func (l *Logger) WriteFile(filename string, bf *bytes.Buffer) error {
	if bf.Len() == 0 {
		return nil
	}
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	bf.WriteTo(f)
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}
