package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var p = fmt.Println
var pf = fmt.Printf

var (
	isPrintMean   bool
	isPrintMedian bool
	isPrintMode   bool
	isPrintSD     bool
)

func init() {
	flag.BoolVar(&isPrintMean, "mean", false, "Print mean value")
	flag.BoolVar(&isPrintMedian, "median", false, "Print median value")
	flag.BoolVar(&isPrintMode, "mode", false, "Print mode value")
	flag.BoolVar(&isPrintSD, "sd", false, "Print standard deviation value")
	flag.Parse()
	if !isPrintMean && !isPrintMedian && !isPrintMode && !isPrintSD {
		isPrintMean = true
		isPrintMedian = true
		isPrintMode = true
		isPrintSD = true
	}
}

func main() {
	var (
		numbers           []int
		frequency         map[int]int
		sum               int
		count             int
		mean              float64
		median            float64
		mode              float64
		standardDeviation float64
	)

	frequency = make(map[int]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputLine := scanner.Text()
		if inputLine == "" {
			gracefulExit("Empty input line", "Valid values in [-100000;100000]")
		}
		num, err := strconv.Atoi(inputLine)
		if err != nil {
			gracefulExit("Error parsing input line as int: "+inputLine, "Valid values in [-100000;100000]")
		}
		if num > 100000 {
			gracefulExit("The number is too big: "+inputLine, "Valid values in [-100000;100000]")
		}
		if num < -100000 {
			gracefulExit("The number is too small: "+inputLine, "Valid values in [-100000;100000]")
		}
		numbers = append(numbers, num)
		if _, ok := frequency[num]; !ok {
			frequency[num] = 1
		} else {
			frequency[num]++
		}
		sum += num
		count++
	}
	if err := scanner.Err(); err != nil {
		gracefulExit(err.Error(), "")

	}
	if count < 1 {
		gracefulExit("Too few numbers", "")

	}
	sort.Ints(numbers)
	mean = float64(sum) / float64(count)
	median = calcMedian(numbers, sum, count)
	mode = calcMode(frequency)
	standardDeviation = calcStandardDeviation(numbers, count, mean)
	if isPrintMean {
		pf("Mean: %.2f\n", mean)
	}
	if isPrintMedian {
		pf("Median: %.2f\n", median)
	}
	if isPrintMode {
		pf("Mode: %.2f\n", mode)
	}
	if isPrintSD {
		pf("SD: %.2f\n", standardDeviation)
	}
}

func gracefulExit(msg string, additions string) {
	p(msg)
	if additions != "" {
		p(additions)
	}
	os.Exit(1)
}

func calcMedian(numbers []int, sum int, count int) float64 {
	var (
		index int
	)

	if count%2 == 0 {
		index = count/2 - 1
		return float64(numbers[index]+numbers[index+1]) / 2
	}
	index = (count - 1) / 2
	return float64(numbers[index])
}

func calcMode(frequency map[int]int) float64 {
	var (
		mode    float64
		maxFreq int
	)

	for num, freq := range frequency {
		if maxFreq < freq {
			maxFreq = freq
			mode = float64(num)
		} else if maxFreq == freq {
			if float64(num) < mode {
				mode = float64(num)
			}
		}
	}
	return mode
}

func calcStandardDeviation(numbers []int, count int, mean float64) float64 {
	var (
		variance float64
		tempSum  float64
	)

	for _, num := range numbers {
		tempSum += math.Pow((float64(num) - mean), 2)
	}
	variance = float64(tempSum) / float64(count)
	return math.Sqrt(variance)
}
