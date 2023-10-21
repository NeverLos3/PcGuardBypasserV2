package main

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/sys/windows/registry" // 레지스트리 패키지 추가

	"github.com/shirou/gopsutil/process"
)

func main() {
	fmt.Println("닥쳐 스키도")
	// 일시 중지할 프로세스 이름 리스트
	// 5초 동안 멈춥니다.
	duration := 5 * time.Second
	time.Sleep(duration)
	
}
