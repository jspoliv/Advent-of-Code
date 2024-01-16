package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	startTime := time.Now().UnixMilli()

	res1 := solvePart1("input.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("input.txt", "rx")
	fmt.Println("Part 2:", res2)

	fmt.Printf("Duration : %vms", time.Now().UnixMilli()-startTime)
}

type ModuleType int
type Pulse bool

const (
	UNTYPED ModuleType = iota
	BROADCASTER
	FLIP_FLOP
	CONJUNCTION

	HIGH Pulse = true
	LOW  Pulse = false
)

type Module struct {
	Type          ModuleType
	Destinations  []string
	On            bool
	InputMemories map[string]Pulse
}

type Input struct {
	Pulse Pulse
	Src   string
	Dest  string
}

type Metadata struct {
	Targets        map[string]bool
	Counters       map[string]int64
	ExpectedOutput Pulse
	PressCount     int64
}

func solvePart1(path string) int64 {
	isDefaultState := func(modules map[string]Module) bool {
		for _, module := range modules {
			if module.On {
				return false
			}

			for _, memory := range module.InputMemories {
				if memory == HIGH {
					return false
				}
			}
		}

		return true
	}

	modules := parseInputs("input.txt")
	pulseCounter := make(map[Pulse]int64)
	targetCount := 1000
	for count := 1; count <= targetCount; count++ {
		sendPulse(
			pulseCounter,
			modules,
			[]Input{{Pulse: LOW, Src: "button", Dest: "broadcaster"}},
			nil,
		)

		if isDefaultState(modules) {
			multiplier := int64(targetCount / count)
			for pulse, amount := range pulseCounter {
				pulseCounter[pulse] = amount * multiplier
			}

			count = targetCount - (targetCount % count)
		}
	}

	return pulseCounter[HIGH] * pulseCounter[LOW]
}

func solvePart2(path string, finalDest string) int64 {
	/*
		Current assumptions :
		- there is only one penultimate module
		- penultimate module is guaranteed to be a conjunction module
		- all modules which trigger penultimate module are also guaranteed to be conjunction modules

		e.g. current inputs :
		- &jm -> rx (final)
		- &sg -> jm, &lm -> jm, &dh -> jm, &db -> jm
	*/

	modules := parseInputs("input.txt")
	finalModule, ok := modules[finalDest]
	if !ok {
		fmt.Println("finalDest module not found")
		return -1
	}

	var penultimateModule Module
	for moduleLabel := range finalModule.InputMemories {
		penultimateModule = modules[moduleLabel]
		break
	}

	metadataTargets := map[string]bool{}
	for moduleLabel := range penultimateModule.InputMemories {
		metadataTargets[moduleLabel] = true
	}

	metadata := Metadata{
		Targets:        metadataTargets,
		Counters:       make(map[string]int64),
		ExpectedOutput: HIGH,
		PressCount:     1,
	}

	for {
		sendPulse(
			make(map[Pulse]int64),
			modules,
			[]Input{{Pulse: LOW, Src: "button", Dest: "broadcaster"}},
			&metadata,
		)

		if len(metadata.Targets) == 0 {
			break
		}

		metadata.PressCount++
	}

	numbers := []int64{}
	for _, count := range metadata.Counters {
		numbers = append(numbers, count)
	}
	return findLCM(numbers)
}

func parseInputs(path string) map[string]Module {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	modules := make(map[string]Module)
	for scanner.Scan() {
		line := scanner.Text()

		data := strings.Split(line, " -> ")
		var label string
		var moduleType ModuleType

		switch data[0][0] {
		case '%':
			label = data[0][1:]
			moduleType = FLIP_FLOP
		case '&':
			label = data[0][1:]
			moduleType = CONJUNCTION
		default:
			label = data[0]
			moduleType = UNTYPED
			if label == "broadcaster" {
				moduleType = BROADCASTER
			}
		}

		destinations := strings.Split(data[1], ", ")
		module, ok := modules[label]
		if ok {
			module.Type = moduleType
			module.Destinations = destinations
			modules[label] = module
		} else {
			modules[label] = Module{
				Type:          moduleType,
				Destinations:  destinations,
				InputMemories: make(map[string]Pulse),
			}
		}

		for _, dest := range destinations {
			destModule, ok := modules[dest]
			if ok {
				destModule.InputMemories[label] = LOW
				modules[dest] = destModule
			} else {
				modules[dest] = Module{
					InputMemories: map[string]Pulse{label: LOW},
				}
			}
		}
	}

	return modules
}

func sendPulse(pulseCounter map[Pulse]int64, modules map[string]Module, inputs []Input, metadata *Metadata) {
	newInputs := []Input{}
	for _, input := range inputs {
		pulseCounter[input.Pulse]++

		destModule := modules[input.Dest]
		switch destModule.Type {
		case BROADCASTER:
			for _, newDest := range destModule.Destinations {
				newInputs = append(newInputs, Input{
					Pulse: input.Pulse,
					Src:   input.Dest,
					Dest:  newDest,
				})
			}
		case FLIP_FLOP:
			if input.Pulse == LOW {
				var nextPulse Pulse
				if destModule.On {
					nextPulse = LOW
					destModule.On = false
				} else {
					nextPulse = HIGH
					destModule.On = true
				}
				modules[input.Dest] = destModule

				for _, newDest := range destModule.Destinations {
					newInputs = append(newInputs, Input{
						Pulse: nextPulse,
						Src:   input.Dest,
						Dest:  newDest,
					})
				}
			}
		case CONJUNCTION:
			destModule.InputMemories[input.Src] = input.Pulse
			modules[input.Dest] = destModule

			nextPulse := LOW
			for _, memory := range destModule.InputMemories {
				if memory == LOW {
					nextPulse = HIGH
					break
				}
			}

			for _, newDest := range destModule.Destinations {
				if metadata != nil {
					if _, ok := metadata.Targets[input.Dest]; ok && nextPulse == metadata.ExpectedOutput {
						metadata.Counters[input.Dest] = metadata.PressCount
						delete(metadata.Targets, input.Dest)
					}

					if len(metadata.Targets) == 0 {
						return
					}
				}

				newInputs = append(newInputs, Input{
					Pulse: nextPulse,
					Src:   input.Dest,
					Dest:  newDest,
				})
			}
		}
	}

	if len(newInputs) > 0 {
		sendPulse(pulseCounter, modules, newInputs, metadata)
	}
}

func findLCM(numbers []int64) int64 {
	gcd := func(a, b int64) int64 { //general common divisor
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	lcm := func(a, b int64) int64 { //least common multiple
		return a * b / gcd(a, b)
	}

	result := int64(1)
	for _, num := range numbers {
		result = lcm(result, num)
	}
	return result
}
