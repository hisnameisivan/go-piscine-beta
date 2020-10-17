package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var p = fmt.Println

func main() {
	var (
		numbers           []float64
		frequency         map[float64]int
		sum               float64
		count             int
		mean              float64
		median            float64
		mode              float64
		standardDeviation float64
	)

	frequency = make(map[float64]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputLine := scanner.Text()
		num, err := strconv.ParseFloat(inputLine, 64)
		if err != nil {
			p(err)
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
	sort.Float64s(numbers)
	mean = sum / float64(count)
	median = calcMedian(numbers, sum, count)
	mode = calcMode(frequency)
	standardDeviation = calcStandardDeviation(numbers, count, mean)
	p(numbers)
	p(sum)
	fmt.Printf("Mean: %.2f\nMedian: %.2f\nMode: %.2f\nSD: %.2f\n", mean, median, mode, standardDeviation)
}

func calcMedian(numbers []float64, sum float64, count int) float64 {
	var (
		index int
	)

	if count%2 == 0 {
		index = count/2 - 1
		return (numbers[index] + numbers[index+1]) / 2
	}
	index = (count - 1) / 2
	return numbers[index]
}

func calcMode(frequency map[float64]int) float64 {
	var (
		mode    float64
		maxFreq int
	)

	for num, freq := range frequency {
		if maxFreq < freq {
			freq = maxFreq
			mode = num
		} else if maxFreq == freq {
			if num < mode {
				mode = num
			}
		}
	}
	return mode
}

func calcStandardDeviation(numbers []float64, count int, mean float64) float64 {
	var (
		variance float64
		tempSum  float64
	)

	for _, num := range numbers {
		tempSum += math.Pow((num - mean), 2)
	}
	variance = tempSum / float64(count)
	return math.Sqrt(variance)
}
