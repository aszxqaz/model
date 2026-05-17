package main

// func main() {
// 	flag.Parse()

// 	m, err := model.Load("testdata/models/model.gob")
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Model params:")
// 	fmt.Printf("Targets: %d..%d\n", m.Params.TargetMin, m.Params.TargetMax)
// 	fmt.Printf("Seconds: %d..%d\n", m.Params.SecondsMin, m.Params.SecondsMax)
// 	fmt.Println("Volumes: ", m.Params.Volumes)
// 	fmt.Println("Takers: ", m.Params.Takers)
// 	fmt.Println("---")
// 	fmt.Fprintln(os.Stderr, "enter: <target> <seconds> <volume> <taker>")

// 	sc := bufio.NewScanner(os.Stdin)
// 	for sc.Scan() {
// 		line := strings.TrimSpace(sc.Text())
// 		if line == "" {
// 			continue
// 		}
// 		fs := strings.Fields(line)
// 		if len(fs) != 4 {
// 			fmt.Fprintln(os.Stderr, "bad input; expected: <target> <seconds> <volume> <taker>")
// 			continue
// 		}

// 		target, err := strconv.Atoi(fs[0])
// 		if err != nil {
// 			fmt.Fprintln(os.Stderr, "bad target:", err)
// 			continue
// 		}

// 		sec, err := strconv.Atoi(fs[1])
// 		if err != nil || sec < 1 {
// 			fmt.Fprintln(os.Stderr, "bad seconds (must be int >= 1)")
// 			continue
// 		}

// 		volume, err := strconv.ParseFloat(fs[2], 64)
// 		if err != nil {
// 			fmt.Fprintln(os.Stderr, "bad volume:", err)
// 			continue
// 		}

// 		taker, err := strconv.ParseFloat(fs[3], 64)
// 		if err != nil {
// 			fmt.Fprintln(os.Stderr, "bad taker:", err)
// 			continue
// 		}

// 		prob, err := m.GetProbability(sec, target, volume, taker)
// 		if err != nil {
// 			fmt.Println("Error: " + err.Error() + "\n")
// 		} else {
// 			fmt.Printf("probability=%.4f\n", prob.Probability)
// 			// fmt.Printf("frequency=%.4f\n", prob.Frequency)
// 		}
// 	}

// 	if err := sc.Err(); err != nil {
// 		fmt.Fprintln(os.Stderr, "read stdin:", err)
// 		os.Exit(1)
// 	}
// }
