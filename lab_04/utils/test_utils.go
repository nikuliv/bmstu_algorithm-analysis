package utils

/*
   #include <Windows.h>
   #include <stdio.h>
   double getCPUTime(void) {
    FILETIME createTime;
    FILETIME exitTime;
    FILETIME kernelTime;
    FILETIME userTime;
    if ( GetProcessTimes( GetCurrentProcess( ),
        &createTime, &exitTime, &kernelTime, &userTime ) != -1 )
    {
        SYSTEMTIME userSystemTime;
        if ( FileTimeToSystemTime( &userTime, &userSystemTime ) != -1 )
            return (double)userSystemTime.wHour * 3600000.0 +
                (double)userSystemTime.wMinute * 60000.0 +
                (double)userSystemTime.wSecond * 1000 +
                (double)userSystemTime.wMilliseconds;
    }
    return -1;
}
*/
import "C"

func GetCPU() float64 {
	result := (float64)(C.getCPUTime())
	return result
}


func Reverse(numbers []int) {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
}
