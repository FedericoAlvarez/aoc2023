package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

////go:embed input.txt test_input
//var input string

//go:embed input.txt
var test_input string

type almanac struct {
	destination_range_start int
	source_range_start      int
	range_lenght            int
}

func main() {
	//part1()
	start := time.Now()

	part2()

	duration := time.Since(start)
	fmt.Println("Duration: ", duration)
}
func part2() {
	lines := strings.Split(test_input, "\n")
	r := regexp.MustCompile(`(\d+)`)
	seeds := r.FindAllString(lines[0], -1)
	index := 2

	seed_to_soil_map := make(map[int]almanac)
	soil_to_fertilizer_map := make(map[int]almanac)
	fertilizer_to_water_map := make(map[int]almanac)
	water_to_light_map := make(map[int]almanac)
	light_to_temperature_map := make(map[int]almanac)
	temperature_to_humidity_map := make(map[int]almanac)
	humidity_to_location_map := make(map[int]almanac)

	if lines[index] == "seed-to-soil map:" {
		index++
		for lines[index] != "" {
			sts := r.FindAllString(lines[index], -1)
			seed_to_soil_map[stringToInt(sts[1])] = almanac{
				destination_range_start: stringToInt(sts[0]),
				source_range_start:      stringToInt(sts[1]),
				range_lenght:            stringToInt(sts[2]),
			}
			index++
		}
	}
	index++
	if lines[index] == "soil-to-fertilizer map:" {
		index++
		for lines[index] != "" {
			sts := r.FindAllString(lines[index], -1)
			soil_to_fertilizer_map[stringToInt(sts[1])] = almanac{
				destination_range_start: stringToInt(sts[0]),
				source_range_start:      stringToInt(sts[1]),
				range_lenght:            stringToInt(sts[2]),
			}
			index++
		}
	}
	index++
	if lines[index] == "fertilizer-to-water map:" {
		index++
		for lines[index] != "" {
			sts := r.FindAllString(lines[index], -1)
			fertilizer_to_water_map[stringToInt(sts[1])] = almanac{
				destination_range_start: stringToInt(sts[0]),
				source_range_start:      stringToInt(sts[1]),
				range_lenght:            stringToInt(sts[2]),
			}
			index++
		}
	}
	index++
	if lines[index] == "water-to-light map:" {
		index++
		for lines[index] != "" {
			sts := r.FindAllString(lines[index], -1)
			water_to_light_map[stringToInt(sts[1])] = almanac{
				destination_range_start: stringToInt(sts[0]),
				source_range_start:      stringToInt(sts[1]),
				range_lenght:            stringToInt(sts[2]),
			}
			index++
		}
	}
	index++
	if lines[index] == "light-to-temperature map:" {
		index++
		for lines[index] != "" {
			sts := r.FindAllString(lines[index], -1)
			light_to_temperature_map[stringToInt(sts[1])] = almanac{
				destination_range_start: stringToInt(sts[0]),
				source_range_start:      stringToInt(sts[1]),
				range_lenght:            stringToInt(sts[2]),
			}
			index++
		}
	}
	index++
	if lines[index] == "temperature-to-humidity map:" {
		index++
		for lines[index] != "" {
			sts := r.FindAllString(lines[index], -1)
			temperature_to_humidity_map[stringToInt(sts[1])] = almanac{
				destination_range_start: stringToInt(sts[0]),
				source_range_start:      stringToInt(sts[1]),
				range_lenght:            stringToInt(sts[2]),
			}
			index++
		}
	}
	index++
	if lines[index] == "humidity-to-location map:" {
		index++
		for lines[index] != "" {
			sts := r.FindAllString(lines[index], -1)
			humidity_to_location_map[stringToInt(sts[1])] = almanac{
				destination_range_start: stringToInt(sts[0]),
				source_range_start:      stringToInt(sts[1]),
				range_lenght:            stringToInt(sts[2]),
			}
			index++
		}
	}

	seedToSoilKeys := sortMapKEys(seed_to_soil_map)
	soilToFertilizerKey := sortMapKEys(soil_to_fertilizer_map)
	fertilizerToWaterKey := sortMapKEys(fertilizer_to_water_map)
	waterToLightKey := sortMapKEys(water_to_light_map)
	lightToTemperatureKey := sortMapKEys(light_to_temperature_map)
	temperatureToHumidityKey := sortMapKEys(temperature_to_humidity_map)
	humidityToLocationKey := sortMapKEys(humidity_to_location_map)

	min := 99999999999999999
	index = 0

	c := make(chan int)

	go func() {
		var wg sync.WaitGroup
		for len(seeds) > index {

			start := stringToInt(seeds[index])
			end := stringToInt(seeds[index]) + stringToInt(seeds[index+1])
			index += 2
			wg.Add(1)
			go func(s, e int) {
				defer wg.Done()
				for i := s; i < e; i++ {
					seedAsInt := i
					seedAsInt = calculate(seedAsInt, seedToSoilKeys, seed_to_soil_map)
					seedAsInt = calculate(seedAsInt, soilToFertilizerKey, soil_to_fertilizer_map)
					seedAsInt = calculate(seedAsInt, fertilizerToWaterKey, fertilizer_to_water_map)
					seedAsInt = calculate(seedAsInt, waterToLightKey, water_to_light_map)
					seedAsInt = calculate(seedAsInt, lightToTemperatureKey, light_to_temperature_map)
					seedAsInt = calculate(seedAsInt, temperatureToHumidityKey, temperature_to_humidity_map)
					seedAsInt = calculate(seedAsInt, humidityToLocationKey, humidity_to_location_map)
					c <- seedAsInt

				}
			}(start, end)
		}
		wg.Wait()
		close(c)
	}()

	for {
		res, ok := <-c
		if !ok {
			break

		}
		if min > res {
			min = res
		}
	}

	fmt.Println(min)
}

func part1() {
	lines := strings.Split(test_input, "\n")
	r := regexp.MustCompile(`(\d+)`)
	seeds := r.FindAllString(lines[0], -1)
	index := 2

	seed_to_soil_map := make(map[int]almanac)
	soil_to_fertilizer_map := make(map[int]almanac)
	fertilizer_to_water_map := make(map[int]almanac)
	water_to_light_map := make(map[int]almanac)
	light_to_temperature_map := make(map[int]almanac)
	temperature_to_humidity_map := make(map[int]almanac)
	humidity_to_location_map := make(map[int]almanac)

	if lines[index] == "seed-to-soil map:" {
		index++
		for lines[index] != "" {
			sts := r.FindAllString(lines[index], -1)
			seed_to_soil_map[stringToInt(sts[1])] = almanac{
				destination_range_start: stringToInt(sts[0]),
				source_range_start:      stringToInt(sts[1]),
				range_lenght:            stringToInt(sts[2]),
			}
			index++
		}
	}
	index++
	if lines[index] == "soil-to-fertilizer map:" {
		index++
		for lines[index] != "" {
			sts := r.FindAllString(lines[index], -1)
			soil_to_fertilizer_map[stringToInt(sts[1])] = almanac{
				destination_range_start: stringToInt(sts[0]),
				source_range_start:      stringToInt(sts[1]),
				range_lenght:            stringToInt(sts[2]),
			}
			index++
		}
	}
	index++
	if lines[index] == "fertilizer-to-water map:" {
		index++
		for lines[index] != "" {
			sts := r.FindAllString(lines[index], -1)
			fertilizer_to_water_map[stringToInt(sts[1])] = almanac{
				destination_range_start: stringToInt(sts[0]),
				source_range_start:      stringToInt(sts[1]),
				range_lenght:            stringToInt(sts[2]),
			}
			index++
		}
	}
	index++
	if lines[index] == "water-to-light map:" {
		index++
		for lines[index] != "" {
			sts := r.FindAllString(lines[index], -1)
			water_to_light_map[stringToInt(sts[1])] = almanac{
				destination_range_start: stringToInt(sts[0]),
				source_range_start:      stringToInt(sts[1]),
				range_lenght:            stringToInt(sts[2]),
			}
			index++
		}
	}
	index++
	if lines[index] == "light-to-temperature map:" {
		index++
		for lines[index] != "" {
			sts := r.FindAllString(lines[index], -1)
			light_to_temperature_map[stringToInt(sts[1])] = almanac{
				destination_range_start: stringToInt(sts[0]),
				source_range_start:      stringToInt(sts[1]),
				range_lenght:            stringToInt(sts[2]),
			}
			index++
		}
	}
	index++
	if lines[index] == "temperature-to-humidity map:" {
		index++
		for lines[index] != "" {
			sts := r.FindAllString(lines[index], -1)
			temperature_to_humidity_map[stringToInt(sts[1])] = almanac{
				destination_range_start: stringToInt(sts[0]),
				source_range_start:      stringToInt(sts[1]),
				range_lenght:            stringToInt(sts[2]),
			}
			index++
		}
	}
	index++
	if lines[index] == "humidity-to-location map:" {
		index++
		for lines[index] != "" {
			sts := r.FindAllString(lines[index], -1)
			humidity_to_location_map[stringToInt(sts[1])] = almanac{
				destination_range_start: stringToInt(sts[0]),
				source_range_start:      stringToInt(sts[1]),
				range_lenght:            stringToInt(sts[2]),
			}
			index++
		}
	}

	fmt.Println(seeds)
	seedToSoilKeys := sortMapKEys(seed_to_soil_map)
	soilToFertilizerKey := sortMapKEys(soil_to_fertilizer_map)
	fertilizerToWaterKey := sortMapKEys(fertilizer_to_water_map)
	waterToLightKey := sortMapKEys(water_to_light_map)
	lightToTemperatureKey := sortMapKEys(light_to_temperature_map)
	temperatureToHumidityKey := sortMapKEys(temperature_to_humidity_map)
	humidityToLocationKey := sortMapKEys(humidity_to_location_map)

	mino := 99999999999999999

	for _, s := range seeds {
		seedAsInt := stringToInt(s)
		seedAsInt = calculate(seedAsInt, seedToSoilKeys, seed_to_soil_map)
		seedAsInt = calculate(seedAsInt, soilToFertilizerKey, soil_to_fertilizer_map)
		seedAsInt = calculate(seedAsInt, fertilizerToWaterKey, fertilizer_to_water_map)
		seedAsInt = calculate(seedAsInt, waterToLightKey, water_to_light_map)
		seedAsInt = calculate(seedAsInt, lightToTemperatureKey, light_to_temperature_map)
		seedAsInt = calculate(seedAsInt, temperatureToHumidityKey, temperature_to_humidity_map)
		seedAsInt = calculate(seedAsInt, humidityToLocationKey, humidity_to_location_map)
		fmt.Println("-----")
		if mino > seedAsInt {
			mino = seedAsInt
		}
	}

	fmt.Println(mino)
}

func stringToInt(s string) int {
	atoi, err := strconv.Atoi(s)
	if err != nil {
		panic("ERROR")
	}
	return atoi
}

func sortMapKEys(m map[int]almanac) []int {

	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})

	return keys
}

func calculate(seedAsInt int, key []int, m map[int]almanac) int {
	for _, k := range key {
		if seedAsInt >= m[k].source_range_start && seedAsInt <= m[k].source_range_start+m[k].range_lenght {
			seedAsInt = (m[k].destination_range_start - m[k].source_range_start) + seedAsInt
			return seedAsInt
		}

	}
	return seedAsInt
}
