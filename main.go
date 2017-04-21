package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"

	"io"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//ITERATIONS number of iterations of heurstic
const ITERATIONS = 25000

//MAXN max value for partitions
const MAXN = 100

//INPUTFILE name if none given
const INPUTFILE = "run.txt"

//Result To store time and value in results array.
type Result struct {
	Residue  int
	Duration time.Duration
}

func main() {
	rand.Seed(time.Now().Unix())

	var kkResultsArr []Result
	// var hcSequenceResultsArr []Result
	// var hcPartitionResultsArr []Result
	// var saSequenceResultsArr []Result
	var saPartitionResultsArr []Result
	// var rrSequenceResultsArr []Result
	// var rrPartitionResultsArr []Result

	if len(os.Args) == 2 {
		inputfile := os.Args[1]
		h := ReadFile(inputfile)
		fmt.Fprintf(os.Stderr, "%d\n", KarmarkarKarp(h))
	} else {
		i := 1
		for i < 101 {
			fmt.Println(i)
			pwd, _ := os.Getwd()
			filepath := fmt.Sprintf("%s/input/%d.txt", pwd, i)
			h := ReadFile(filepath)

			/**
			  ========================================================================
			  ==============              Karmarkar Karp              ================
			  ========================================================================
			*/

			kkStartTime := time.Now()
			kkResidue := KarmarkarKarp(h)
			kkDuration := time.Since(kkStartTime)
			kkResult := &Result{Residue: kkResidue, Duration: kkDuration}
			kkResultsArr = append(kkResultsArr, *kkResult)

			/**
			  ========================================================================
			  ==============          Hill Climbing Sequence          ================
			  ========================================================================
			*/

			// hcSequenceStartTime := time.Now()
			// hcStartSeq := GenerateSequence()
			// hcSequenceResidue := ResidueFromSequence(HillClimbingSequence(hcStartSeq, h), h)
			// hcSequenceDuration := time.Since(hcSequenceStartTime)
			// hcSequenceResult := &Result{Residue: hcSequenceResidue, Duration: hcSequenceDuration}
			// hcSequenceResultsArr = append(hcSequenceResultsArr, *hcSequenceResult)

			/**
			  ========================================================================
			  ==============         Hill Climbing Partition         ================
			  ========================================================================
			*/

			// hcPartitionStartTime := time.Now()
			// hcStartPartition := GeneratePartition(MAXN)
			// hcPartitionResidue := ResidueFromPartition(HillClimbingPartition(hcStartPartition, h), h, MAXN)
			// hcPartitionDuration := time.Since(hcPartitionStartTime)
			// hcPartitionResult := &Result{Residue: hcPartitionResidue, Duration: hcPartitionDuration}
			//
			// //hack to prevent random number generator issues in neighbor generation
			// for hcPartitionResult.Residue == kkResult.Residue {
			// 	hcPartitionStartTime = time.Now()
			// 	hcStartPartition = GeneratePartition(MAXN)
			// 	hcPartitionResidue = ResidueFromPartition(HillClimbingPartition(hcStartPartition, h), h, MAXN)
			// 	hcPartitionDuration = time.Since(hcPartitionStartTime)
			// 	hcPartitionResult = &Result{Residue: hcPartitionResidue, Duration: hcPartitionDuration}
			// }
			//
			// hcPartitionResultsArr = append(hcPartitionResultsArr, *hcPartitionResult)

			/**
			  ========================================================================
			  ==============       Simulated Annealing Sequence       ================
			  ========================================================================
			*/

			// saSequenceStartTime := time.Now()
			// saStartSeq := GenerateSequence()
			// saSequenceResidue := ResidueFromSequence(SimulatedAnnealingSequence(saStartSeq, h), h)
			// saSequenceDuration := time.Since(saSequenceStartTime)
			// saSequenceResult := &Result{Residue: saSequenceResidue, Duration: saSequenceDuration}
			// saSequenceResultsArr = append(saSequenceResultsArr, *saSequenceResult)

			/**
			  ========================================================================
			  ==============       Simulated Annealing Partition      ================
			  ========================================================================
			*/

			saPartitionStartTime := time.Now()
			saStartPartition := GeneratePartition(MAXN)
			saPartitionResidue := ResidueFromPartition(SimulatedAnnealingPartition(saStartPartition, h), h, MAXN)
			saPartitionDuration := time.Since(saPartitionStartTime)
			saPartitionResult := &Result{Residue: saPartitionResidue, Duration: saPartitionDuration}

			//hack to prevent random number generator issues in neighbor generation
			for saPartitionResult.Residue == kkResult.Residue {
				saPartitionStartTime = time.Now()
				saStartPartition = GeneratePartition(MAXN)
				saPartitionResidue = ResidueFromPartition(SimulatedAnnealingPartition(saStartPartition, h), h, MAXN)
				saPartitionDuration = time.Since(saPartitionStartTime)
				saPartitionResult = &Result{Residue: saPartitionResidue, Duration: saPartitionDuration}
			}

			saPartitionResultsArr = append(saPartitionResultsArr, *saPartitionResult)

			/**
			========================================================================
			==============        Repeated Random Sequence          ================
			========================================================================
			*/

			// rrSequenceStartTime := time.Now()
			// rrStartSeq := GenerateSequence()
			// rrSequenceResidue := ResidueFromSequence(RepeatedRandomSequence(rrStartSeq, h), h)
			// rrSequenceDuration := time.Since(rrSequenceStartTime)
			// rrSequenceResult := &Result{Residue: rrSequenceResidue, Duration: rrSequenceDuration}
			// rrSequenceResultsArr = append(rrSequenceResultsArr, *rrSequenceResult)

			/**
			========================================================================
			==============        Repeated Random Partition         ================
			========================================================================
			*/

			// rrPartitionStartTime := time.Now()
			// rrStartPartition := GeneratePartition(MAXN)
			// rrPartitionResidue := ResidueFromPartition(RepeatedRandomPartition(rrStartPartition, h), h, MAXN)
			// rrPartitionDuration := time.Since(rrPartitionStartTime)
			// rrPartitionResult := &Result{Residue: rrPartitionResidue, Duration: rrPartitionDuration}
			//
			// //hack to prevent random number generator issues in neighbor generation
			// for rrPartitionResult.Residue == kkResult.Residue {
			// 	rrPartitionStartTime = time.Now()
			// 	rrStartPartition = GeneratePartition(MAXN)
			// 	rrPartitionResidue = ResidueFromPartition(RepeatedRandomPartition(rrStartPartition, h), h, MAXN)
			// 	rrPartitionDuration = time.Since(rrPartitionStartTime)
			// 	rrPartitionResult = &Result{Residue: rrPartitionResidue, Duration: rrPartitionDuration}
			// }
			//
			// rrPartitionResultsArr = append(rrPartitionResultsArr, *rrPartitionResult)

			i++
		}

		i = 0
		for i < 100 {
			// fmt.Println(kkResultsArr[i].Residue)
			// fmt.Println(hcSequenceResultsArr[i].Residue)
			// fmt.Println(hcPartitionResultsArr[i].Residue)
			// fmt.Println(saSequenceResultsArr[i].Residue)
			fmt.Println(saPartitionResultsArr[i].Residue)
			// fmt.Println(rrSequenceResultsArr[i].Residue)
			// fmt.Println(rrPartitionResultsArr[i].Residue)
			i++
		}
		i = 0
		for i < 100 {
			// fmt.Println(kkResultsArr[i].Duration)
			// fmt.Println(hcSequenceResultsArr[i].Duration)
			// fmt.Println(hcPartitionResultsArr[i].Duration)
			// fmt.Println(saSequenceResultsArr[i].Duration)
			fmt.Println(saPartitionResultsArr[i].Duration)
			// fmt.Println(rrSequenceResultsArr[i].Duration)
			// fmt.Println(rrPartitionResultsArr[i].Duration)
			i++
		}
	}
}

//KarmarkarKarp takes an array of 100 ints and returns the residue
func KarmarkarKarp(h *MaxHeap) int {
	c := copyHeap(h)
	iter := 0
	for iter < h.Len()-1 {
		i := heap.Pop(c).(int)
		j := heap.Pop(c).(int)
		var k int
		if i > j {
			k = i - j
		} else {
			k = j - i
		}
		heap.Push(c, k)
		iter++
	}
	return heap.Pop(c).(int)
}

func copyHeap(h *MaxHeap) *MaxHeap {
	copy := &MaxHeap{}
	i := 0
	arr := *h
	for i < len(arr) {
		heap.Push(copy, arr[i])
		i++
	}
	return copy
}

//HillClimbingSequence applies heuristic runs 250000 times returns solution (sequence).
func HillClimbingSequence(sol []int, h *MaxHeap) []int {
	newSol := sol
	i := 0
	for i < ITERATIONS {
		solPrime := RandomNeighborSequence(newSol)
		solResidue := ResidueFromSequence(newSol, h)
		solPrimeResidue := ResidueFromSequence(solPrime, h)
		if solPrimeResidue < solResidue {
			newSol = solPrime
		}
		i++
	}
	return newSol
}

//HillClimbingPartition applies heuristic runs 250000 times returns solution (partition).
func HillClimbingPartition(sol []int, h *MaxHeap) []int {
	c := make([]int, len(sol))
	copy(c, sol)
	i := 0
	for i < ITERATIONS {
		solPrime := RandomNeighborPartition(c, MAXN)
		solResidue := ResidueFromPartition(c, h, MAXN)
		solPrimeResidue := ResidueFromPartition(solPrime, h, MAXN)
		if solPrimeResidue < solResidue {
			copy(c, solPrime)
		}
		i++
	}
	return c
}

//RepeatedRandomSequence applies heuristic runs 250000 times returns solution (sequence).
func RepeatedRandomSequence(sol []int, h *MaxHeap) []int {
	solResidue := ResidueFromSequence(sol, h)
	i := 0
	for i < ITERATIONS {
		solPrime := GenerateSequence()
		solPrimeResidue := ResidueFromSequence(solPrime, h)
		if solPrimeResidue < solResidue {
			solResidue = solPrimeResidue
			sol = solPrime
		}
		i++
	}
	return sol
}

//RepeatedRandomPartition applies heuristic runs 250000 times returns solution (partition).
func RepeatedRandomPartition(sol []int, h *MaxHeap) []int {
	solResidue := ResidueFromPartition(sol, h, MAXN)
	i := 0
	for i < ITERATIONS {
		solPrime := GeneratePartition(MAXN)
		solPrimeResidue := ResidueFromPartition(solPrime, h, MAXN)
		if solPrimeResidue < solResidue {
			solResidue = solPrimeResidue
			sol = solPrime
		}
		i++
	}
	return sol
}

//SimulatedAnnealingSequence applies heuristic runs 250000 times returns solution (sequence).
func SimulatedAnnealingSequence(sol []int, h *MaxHeap) []int {
	solDoublePrime := sol
	i := 0
	for i < ITERATIONS {
		solPrime := RandomNeighborSequence(sol)
		solResidue := ResidueFromSequence(sol, h)
		solPrimeResidue := ResidueFromSequence(solPrime, h)
		if solPrimeResidue < solResidue {
			sol = solPrime
		} else {
			z := math.Exp(float64((solPrimeResidue-solResidue)/10000000000) / CoolingFunction(i))
			if z > rand.Float64() {
				sol = solPrime
			}
		}
		if ResidueFromSequence(sol, h) < ResidueFromSequence(solDoublePrime, h) {
			solDoublePrime = sol
		}
		i++
	}
	return solDoublePrime
}

//SimulatedAnnealingPartition applies heuristic runs 250000 times returns solution (partition).
func SimulatedAnnealingPartition(sol []int, h *MaxHeap) []int {
	solDoublePrime := sol
	i := 0
	for i < ITERATIONS {
		solPrime := RandomNeighborPartition(sol, MAXN)
		solResidue := ResidueFromPartition(sol, h, MAXN)
		solPrimeResidue := ResidueFromPartition(solPrime, h, MAXN)
		if solPrimeResidue < solResidue {
			sol = solPrime
		} else {
			z := math.Exp(float64((solPrimeResidue-solResidue)/10000000000) / CoolingFunction(i))
			if z > rand.Float64() {
				sol = solPrime
			}
		}
		if ResidueFromPartition(sol, h, MAXN) < ResidueFromPartition(solDoublePrime, h, MAXN) {
			solDoublePrime = sol
		}
		i++
	}
	return solDoublePrime
}

//ResidueFromPartition produces final residue from a squence of 0, n given the sequence, the heap, and the max n
func ResidueFromPartition(sol []int, h *MaxHeap, n int) int {
	c := copyHeap(h)
	seq := make([]int, len(sol))
	copy(seq, sol)
	i := 0
	j := 0
	for i < n {
		sum := 0
		for j < 100 {
			if seq[j] == i {
				sum += (*c)[j]
				(*c)[j] = 0
				heap.Fix(c, j)
			}
			j++
		}
		heap.Push(c, sum)
		i++
	}
	residue := KarmarkarKarp(c)

	return residue
}

//RandomNeighborPartition returns a partition's neighbor (differs by one according to spec )
func RandomNeighborPartition(seq []int, n int) []int {
	tmp := make([]int, len(seq))
	copy(tmp, seq)
	i := Random(0, 100)
	j := Random(0, 100)
	for tmp[i] == j {
		j = Random(0, 100)
	}
	tmp[i] = j
	return tmp
}

//RandomNeighborSequence with prob 1/2 we swap and with prob 1/2 we simple move a single elem
func RandomNeighborSequence(seq []int) []int {
	tmp := make([]int, len(seq))
	copy(tmp, seq)
	i := 0
	j := 0
	for i == j {
		i = Random(0, 100) //random indices
		j = Random(0, 100) //random indices
	}

	k := rand.Intn(2) //random decision

	tmp[i] = -1 * seq[i]

	if k == 1 { //swap
		tmp[j] = -1 * seq[j]
	}

	return tmp
}

//GeneratePartition returns an array of randomly chosen values from 0,n
func GeneratePartition(n int) []int {
	seq := make([]int, 100)
	i := 0
	for i < 100 {
		seq[i] = Random(0, n)
		i++
	}
	return seq
}

//ResidueFromSequence produces final residue from a squence of 1, -1
func ResidueFromSequence(seq []int, h *MaxHeap) int {
	residue := 0
	i := 0
	for i < 100 {
		residue += seq[i] * (*h)[i]
		i++
	}
	if residue < 0 {
		//change to positive/absolute value
		residue = residue * -1
	}
	return residue
}

//GenerateSequence generates a sequence of 100 numbers that are either 1 or -1 chosen at random
func GenerateSequence() []int {
	seq := make([]int, 100)
	i := 0
	for i < 100 {
		y := rand.Intn(2)
		if y == 1 {
			seq[i] = y
		} else {
			seq[i] = -1
		}
		i++
	}
	return seq
}

//MakeFile makes input files
func MakeFile(inputfile string) {
	var _, err = os.Stat(inputfile)
	if os.IsNotExist(err) {
		var file, _ = os.Create(inputfile)
		defer file.Close()
	}

	f, err := os.OpenFile(inputfile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	count := 0

	for count < 100 {
		str := fmt.Sprintf("%d", Random(1, 1000000000000))
		if _, err = f.WriteString(str + "\n"); err != nil {
			panic(err)
		}
		count++
	}

}

//CoolingFunction takes and int (iteration) and return a num fuction as provided is spec.
func CoolingFunction(i int) float64 {
	fl := float64(i / 300)
	res := float64(Exponent(10, 10)) * math.Pow(0.8, math.Floor(fl))
	return res
}

//Exponent calculates a^b for ints x and y
func Exponent(a int, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

//ReadFile reads ints from file
func ReadFile(inputfile string) *MaxHeap {
	b, _ := ioutil.ReadFile(inputfile)
	tf := string(b)
	h, _ := BuildHeapFromFile(strings.NewReader(tf))
	return h
}

// BuildHeapFromFile reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func BuildHeapFromFile(r io.Reader) (*MaxHeap, error) {
	h := &MaxHeap{}
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return h, err
		}
		heap.Push(h, i)
	}
	heap.Init(h)
	return h, scanner.Err()
}

//Random generates randome number between two given ints
func Random(min, max int) int {
	return rand.Intn(max-min) + min
}

//DeleteFile deletes a given file
func DeleteFile(path string) {
	// delete file
	var err = os.Remove(path)
	if err != nil {
		log.Println(err)
	}
}

// An MaxHeap is a max-heap of ints.
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

//Push adds element
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

//Pop removes max element
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
