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
		mode              int
		standardDeviation float64
	)

	frequency = make(map[int]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputLine := scanner.Text()
		num, err := strconv.Atoi(inputLine)
		if err != nil {
			p("Error parsing input line as int:", inputLine)
			os.Exit(1)
		}
		if num > 100000 {
			p("The number is too big:", inputLine)
			p("Valid values in [-100000;100000]")
			os.Exit(1)
		}
		if num < -100000 {
			p("The number is too small:", inputLine)
			p("Valid values in [-100000;100000]")
			os.Exit(1)
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
		p(err)
		os.Exit(1)
	}
	if count < 1 {
		p("Too few numbers")
		os.Exit(1)
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
		pf("Mode: %d\n", mode)
	}
	if isPrintSD {
		pf("SD: %.2f\n", standardDeviation)
	}
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

func calcMode(frequency map[int]int) int {
	var (
		mode    int
		maxFreq int
	)

	for num, freq := range frequency {
		if maxFreq < freq {
			maxFreq = freq
			mode = num
		} else if maxFreq == freq {
			if num < mode {
				mode = num
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
