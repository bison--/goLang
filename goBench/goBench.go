// goBench
package main

import (
	"fmt"
	"time"
)

var timesLowest map[string]float64
var timesHighest map[string]float64
var innerLoops int = 1000000000

func main() {
	// init maps!
	timesLowest = make(map[string]float64)
	timesHighest = make(map[string]float64)

	fmt.Println("###############")
	fmt.Println("# goBench 0.1 #")
	fmt.Println("###############")

	var benchName string
	var benchStartTime time.Time

	for i := 0; i < 5; i++ {
		benchName, benchStartTime = startTimer("for")
		testFor()
		stopTimer(benchName, benchStartTime)

		benchName, benchStartTime = startTimer("while")
		testWhile()
		stopTimer(benchName, benchStartTime)

		benchName, benchStartTime = startTimer("testWhilePlusOne")
		testWhilePlusOne()
		stopTimer(benchName, benchStartTime)

		benchName, benchStartTime = startTimer("whileNot")
		testWhileNot()
		stopTimer(benchName, benchStartTime)

		benchName, benchStartTime = startTimer("testWhileTrue")
		testWhileTrue()
		stopTimer(benchName, benchStartTime)

		benchName, benchStartTime = startTimer("testIfTrueFalse")
		testIfTrueFalse()
		stopTimer(benchName, benchStartTime)

		benchName, benchStartTime = startTimer("testIfTrueElse")
		testIfTrueElse()
		stopTimer(benchName, benchStartTime)

		benchName, benchStartTime = startTimer("testSwitchTrueFalse")
		testSwitchTrueFalse()
		stopTimer(benchName, benchStartTime)

		benchName, benchStartTime = startTimer("testSwitchTrueDefault")
		testSwitchTrueDefault()
		stopTimer(benchName, benchStartTime)
	}

	fmt.Println("### LOWEST TIMES ###")
	for key, value := range timesLowest {
		fmt.Println(key, "Duration:", value)
	}

	fmt.Println("### HIGHEST TIMES ###")
	for key, value := range timesHighest {
		fmt.Println(key, "Duration:", value)
	}

	fmt.Println("fin")
}

func startTimer(s string) (string, time.Time) {
	//log.Printf("start: %s\n", s)
	fmt.Printf("start: %s\n", s)
	return s, time.Now()
}

func stopTimer(s string, startTime time.Time) {
	elapsed := time.Since(startTime)
	//log.Printf("end: %s, elapsed %f secs\n", s, elapsed.Seconds())
	fmt.Printf("end: %s, elapsed %f secs\n", s, elapsed.Seconds())

	val, ok := timesLowest[s]
	if ok == false || elapsed.Seconds() < val {
		timesLowest[s] = elapsed.Seconds()
	}

	val, ok = timesHighest[s]
	if ok == false || elapsed.Seconds() > val {
		timesHighest[s] = elapsed.Seconds()
	}
}

func testFor() {
	var dummyVar bool = false
	for i := 0; i < innerLoops; i++ {
		dummyVar = !dummyVar
	}
}

func testWhile() {
	var dummyVar bool = false
	var countTo int = innerLoops
	var i int
	for i < countTo {
		i++
		dummyVar = !dummyVar
	}
}

func testWhilePlusOne() {
	var dummyVar bool = false
	var countTo int = innerLoops
	var i int
	for i < countTo {
		i += 1
		dummyVar = !dummyVar
	}
}

func testWhileNot() {
	var dummyVar bool = false
	var countTo int = innerLoops
	var i int
	for i != countTo {
		i++
		dummyVar = !dummyVar
	}
}

func testWhileTrue() {
	var dummyVar bool = false
	var countTo int = innerLoops
	var i int
	for true {
		i++
		dummyVar = !dummyVar
		if i == countTo {
			break
		}
	}
}

func testIfTrueFalse() {
	var dummyVar bool = false
	var countTo int = innerLoops
	var i int
	for i < countTo {
		i++
		if dummyVar == true {
			dummyVar = false
		} else if dummyVar == false {
			dummyVar = true
		}
	}
}

func testIfTrueElse() {
	var dummyVar bool = false
	var countTo int = innerLoops
	var i int
	for i < countTo {
		i++
		if dummyVar == true {
			dummyVar = false
		} else {
			dummyVar = true
		}
	}
}

func testSwitchTrueFalse() {
	var dummyVar bool = false
	var countTo int = innerLoops
	var i int
	for i < countTo {
		i++
		switch {
		case dummyVar == true:
			dummyVar = false
		case dummyVar == false:
			dummyVar = true
		}
	}
}

func testSwitchTrueDefault() {
	var dummyVar bool = false
	var countTo int = innerLoops
	var i int
	for i < countTo {
		i++
		switch {
		case dummyVar == true:
			dummyVar = false
		default:
			dummyVar = true
		}
	}
}
