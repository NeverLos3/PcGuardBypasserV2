package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/process"
)

func main() {
	fmt.Println("Made by DevaAndSage.")
	fmt.Println("5초 뒤에 우회가 시작됩니다.")

	// 5초 동안 멈춥니다.
	duration := 5 * time.Second
	time.Sleep(duration)

	fmt.Println("우회 시도중...")

	// 일시 중단할 프로세스들의 PID 리스트를 지정합니다.
	targetPIDs := []int{12124, 12156, 5096, 5008} // 변경해야 할 PID 리스트로 설정하세요.

	// 각 프로세스에 대해 작업을 수행합니다.
	for _, targetPID := range targetPIDs {
		targetProcess, err := process.NewProcess(int32(targetPID))
		if err != nil {
			fmt.Printf("프로세스 %d를 가져오는 데 오류 발생: %v\n", targetPID, err)
			continue // 다음 프로세스를 시도합니다.
		}

		// 프로세스를 일시 중단합니다.
		err = targetProcess.Suspend()
		if err != nil {
			fmt.Printf("프로세스 %d를 일시 중단하는 데 오류 발생: %v\n", targetPID, err)
		} else {
			fmt.Printf("프로세스 %d가 일시 중단되었습니다.\n", targetPID)
		}

	}
	fmt.Println("프로그램을 종료하려면 엔터 키를 누르세요.")
	fmt.Scanln() // 엔터 키를 누를 때까지 대기
}
