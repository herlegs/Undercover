package redis

import (
	"testing"
	"os"
	"time"
	"sync"
)

const(
	TestKey = "testkey"
	TestValue = "testvalue"
)

func TestMain(m *testing.M){
	setup()
	code := m.Run()
	Close()
	os.Exit(code)
}

func setup(){
	Set(TestKey, TestValue)
}

func TestConcurrentAccess(t *testing.T){
	closeCh := make(chan struct{})
	workerNum := 11
	wait := &sync.WaitGroup{}
	for i := 0; i < workerNum; i++ {
		id := i
		wait.Add(1)
		go worker(t, id, closeCh, wait)
	}
	time.Sleep(time.Second*3)
	close(closeCh)
	wait.Wait()
}

func worker(t *testing.T, threadID int, closeCh chan struct{}, wait *sync.WaitGroup){
	defer wait.Done()
	for {
		select {
		case <-closeCh:
			return
		default:
			//simulate http request
			time.Sleep(time.Millisecond * 200)
			exist := ExistKey(TestKey)
			if !exist {
				t.Error("thread", threadID, "failed")
				return
			}
		}
	}
}