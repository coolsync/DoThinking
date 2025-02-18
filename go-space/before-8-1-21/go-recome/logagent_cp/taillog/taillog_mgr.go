package taillog

import (
	"fmt"
	"logagent/etcd"
	"time"
)

// 具体管理每一个 tail read log file, write to kafka 进程处理

// global task manager
// var (
// 	tskMgr      *tailLogMgr
// 	tskMap      map[string]*TailTask
// 	newConfChan chan []*etcd.LogEntryConf
// )

var tskMgr *tailLogMgr

type tailLogMgr struct {
	logEntry    []*etcd.LogEntryConf
	tskMap      map[string]*TailTask
	newConfChan chan []*etcd.LogEntryConf
}

func Init(logEntry []*etcd.LogEntryConf) {
	tskMgr = &tailLogMgr{
		logEntry:    logEntry, // 把当前的日志收集项配置信息保存起来
		tskMap:      make(map[string]*TailTask, 16),
		newConfChan: make(chan []*etcd.LogEntryConf), // 无缓冲区的通道
	}

	for _, logEntryObj := range logEntry {
		//conf: *etcd.LogEntryConf
		//logEntry.Path： 要收集的日志文件的路径
		// 初始化的时候起了多少个tailtask 都要记下来，为了后续判断方便
		tailObj := NewTailTask(logEntryObj.Path, logEntryObj.Topic)
		mk := fmt.Sprintf("%s_%s", logEntryObj.Path, logEntryObj.Topic)
		tskMgr.tskMap[mk] = tailObj
	}

	go tskMgr.run() // 启动goroutine, 对应初始化的tskMgr, 可以run tskMgr中多个log文件read
}

// 监听自己的newConfChan，有了新的配置过来之后就做对应的处理
func (t *tailLogMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			for _, conf := range newConf {
				mk := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
				_, ok := t.tskMap[mk]
				if ok {
					// 原来就有，不需要操作
					continue
				} else {
					// 新增的
					tailObj := NewTailTask(conf.Path, conf.Topic)
					t.tskMap[mk] = tailObj
				}
			}
			// 找出原来t.logEntry有，但是newConf中没有的，要删掉
			for _, c1 := range t.logEntry { // 从原来的配置中依次拿出配置项
				isDelete := true
				for _, c2 := range newConf { // 去新的配置中逐一进行比较
					if c2.Path == c1.Path && c2.Topic == c1.Topic {
						isDelete = false
						continue
					}
				}
				if isDelete {
					// 把c1对应的这个tailObj给停掉
					mk := fmt.Sprintf("%s_%s", c1.Path, c1.Topic)
					//t.tskMap[mk] ==> tailObj
					t.tskMap[mk].cancelFunc()
				}
			}
			// 2. 配置删除
			fmt.Println("新的配置来了！", newConf)
		default:
			time.Sleep(time.Second)
		}
	}
}

// 一个函数，向外暴露tskMgr的newConfChan
func NewConfChan() chan<- []*etcd.LogEntryConf {
	return tskMgr.newConfChan
}
