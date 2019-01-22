package main

import (
	"fmt"
	"time"

	//l "github.com/weeradej/tvslogs"
	l "github.com/smsdevteam/tvsglobal/common"
)

func main() {
	//app log
	t0 := time.Now()
	//Process
	t1 := time.Now()
	t2 := t1.Sub(t0)

	logs := l.Applog{}
	logs.OrderNo = "aaa"
	logs.TrackingNo = "aaa"
	logs.ApplicationName = "aaa"
	logs.FunctionName = "aaa"
	logs.OrderDate = "aaa"
	logs.TVSNo = "aaa"
	logs.MobileNo = "aaa"
	logs.SerialNo = "aaa"
	logs.Reference1 = "aaa"
	logs.Reference2 = "aaa"
	logs.Reference3 = "aaa"
	logs.Reference4 = "aaa"
	logs.Reference5 = "aaa"
	logs.Request = "aaa"
	logs.Response = "aaa"
	logs.Start = t0.Format(time.RFC3339Nano)
	logs.End = t1.Format(time.RFC3339Nano)
	logs.Duration = t2.String()

	err := logs.InsertappLog("applog01.log", "Hello world!")
	if err != nil {
		fmt.Println(err)
	}

	//workflowlog
	t0 = time.Now()
	t1 = time.Now()
	t2 = t1.Sub(t0)
	logsWF := l.Workflowlog{}
	logsWF.OrderNo = "aaa"
	logsWF.TrackingNo = "aaa"
	logsWF.ApplicationName = "aaa"
	logsWF.FunctionName = "aaa"
	logsWF.OrderDate = "aaa"
	logsWF.TVSNo = "aaa"
	logsWF.MobileNo = "aaa"
	logsWF.SerialNo = "aaa"
	logsWF.Reference1 = "aaa"
	logsWF.Reference2 = "aaa"
	logsWF.Reference3 = "aaa"
	logsWF.Reference4 = "aaa"
	logsWF.Reference5 = "aaa"
	logsWF.Start = t0.Format(time.RFC3339Nano)
	logsWF.End = t1.Format(time.RFC3339Nano)
	logsWF.Duration = t2.String()

	arrpsCfg := []l.Processconfig{}
	psCfg := l.Processconfig{}
	t0 = time.Now()
	t1 = time.Now()
	t2 = t1.Sub(t0)
	psCfg.CallFunction = "bbbb1"
	psCfg.Start = t0.Format(time.RFC3339Nano)
	psCfg.End = t1.Format(time.RFC3339Nano)
	psCfg.ResultCode = "bbbb1"
	psCfg.ResultDesc = "bbbb1"
	arrpsCfg = append(arrpsCfg, psCfg)

	t0 = time.Now()
	t1 = time.Now()
	t2 = t1.Sub(t0)
	psCfg.CallFunction = "bbbb2"
	psCfg.Start = t0.Format(time.RFC3339Nano)
	psCfg.End = t1.Format(time.RFC3339Nano)
	psCfg.ResultCode = "bbbb2"
	psCfg.ResultDesc = "bbbb2"
	arrpsCfg = append(arrpsCfg, psCfg)

	logsWF.ProcessConfig = arrpsCfg

	errWF := logsWF.Insertworkflowlog("workflow01.log")
	if errWF != nil {
		fmt.Println(errWF)
	}
}
