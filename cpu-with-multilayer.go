package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	cpu    []string
	time   int
	time1  int
	time2  int
	time3  int
	ready1 []string
	ready2 []string
	ready3 []string
	io1    []string
	io2    []string
	io3    []string
	io4    []string
)

func initialized() {
	cpu = make([]string, 2)
	time = 3
	time1 = 0
	time2 = 0
	time3 = 0
	ready1 = make([]string, 10)
	ready2 = make([]string, 10)
	ready3 = make([]string, 10)
	io1 = make([]string, 10)
	io2 = make([]string, 10)
	io3 = make([]string, 10)
	io4 = make([]string, 10)
}

func showProcess() {
	fmt.Printf("\n-----------\n")
	fmt.Printf("Time Quantum = 3\n")
	fmt.Printf("CPU1   -> %s\n", cpu[0])
	fmt.Printf("CPU2   -> %s\n", cpu[1])
	fmt.Printf("1st Ready -> ")
	for i := range ready1 {
		fmt.Printf("%s ", ready1[i])
	}
	fmt.Printf("\n2nd Ready -> ")
	for i := range ready2 {
		fmt.Printf("%s ", ready2[i])
	}
	fmt.Printf("\n3rd Ready -> ")
	for i := range ready3 {
		fmt.Printf("%s ", ready3[i])
	}
	fmt.Printf("\n")
	fmt.Printf("I/O 1 -> ")
	for i := range io1 {
		fmt.Printf("%s ", io1[i])
	}
	fmt.Printf("\n")
	fmt.Printf("I/O 2 -> ")
	for i := range io2 {
		fmt.Printf("%s ", io2[i])
	}
	fmt.Printf("\n")
	fmt.Printf("I/O 3 -> ")
	for i := range io3 {
		fmt.Printf("%s ", io3[i])
	}
	fmt.Printf("\n")
	fmt.Printf("I/O 4 -> ")
	for i := range io4 {
		fmt.Printf("%s ", io4[i])
	}
	fmt.Printf("\n\nCommand > ")
}

func getCommand() string {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	data = strings.Trim(data, "\n")
	return data
}

func command_new(p string) {
	prior := strings.Split(p, "-")
	if cpu[0] == "" && cpu[1] == "" {
		cpu[0] = p
		if prior[1] == "1" {
			time1++
			// fmt.Println(time1)
			// fmt.Println(time2)
			// fmt.Println(time3)
		} else if prior[1] == "2" {
			time2++
			// fmt.Println(time1)
			// fmt.Println(time2)
			// fmt.Println(time3)
		} else {
			time3++
			// fmt.Println(time1)
			// fmt.Println(time2)
			// fmt.Println(time3)
		}
	} else if cpu[0] == "" && cpu[1] != "" {
		cpu[0] = p
		if prior[1] == "1" {
			time1++
			// fmt.Println(time1)
			// fmt.Println(time2)
			// fmt.Println(time3)
		} else if prior[1] == "2" {
			time2++
			// fmt.Println(time1)
			// fmt.Println(time2)
			// fmt.Println(time3)
		} else {
			time3++
			// fmt.Println(time1)
			// fmt.Println(time2)
			// fmt.Println(time3)
		}
	} else if cpu[0] != "" && cpu[1] == "" {
		cpu[1] = p
		if prior[1] == "1" {
			time1++
			// fmt.Println(time1)
			// fmt.Println(time2)
			// fmt.Println(time3)
		} else if prior[1] == "2" {
			time2++
			// fmt.Println(time1)
			// fmt.Println(time2)
			// fmt.Println(time3)
		} else {
			time3++
			// fmt.Println(time1)
			// fmt.Println(time2)
			// fmt.Println(time3)
		}
	} else {
		insertQueue2(p)
	}
}

func command_terminate(i int) {
	if cpu[i-1] != "" {
		if ready1[0] != "" && time1 < 3 {
			cpu[i-1] = deleteQueue(ready1)
			time1++
			// fmt.Println(time1)
			// fmt.Println(time2)
			// fmt.Println(time3)
		} else if (ready1[0] == "" && ready2[0] != "" && time2 < 3) || (time1 == 3 && time2 < 3) {
			cpu[i-1] = deleteQueue(ready2)
			time2++
			time1 = 0
			// fmt.Println(time1)
			// fmt.Println(time2)
			// fmt.Println(time3)
		} else if (ready1[0] == "" && ready2[0] == "" && ready3[0] != "") || ((time1 == 3 || time1 == 0) && time2 == 3 && time3 < 3) {
			cpu[i-1] = deleteQueue(ready3)
			time3++
			time2 = 0
			// fmt.Println(time1)
			// fmt.Println(time2)
			// fmt.Println(time3)
		} else {
			cpu[i-1] = ""
		}
	}
}

func command_expire(i int) {
	if (ready1[0] != "" && time1 < 3) || (ready1[0] != "" && ready2[0] == "" && ready3[0] == "") {
		p := deleteQueue(ready1)
		if p == "" {
			return
		}
		insertQueue2(cpu[i-1])
		cpu[i-1] = p
		time1++
		// fmt.Println(time1)
		// fmt.Println(time2)
		// fmt.Println(time3)
	} else if (ready1[0] == "" && ready2[0] != "" && time2 < 3) || (ready2[0] != "" && time1 == 3 && time2 < 3) {
		p := deleteQueue(ready2)
		if p == "" {
			return
		}
		insertQueue2(cpu[i-1])
		cpu[i-1] = p
		time2++
		time1 = 0
		// fmt.Println(time1)
		// fmt.Println(time2)
		// fmt.Println(time3)
	} else if (ready1[0] == "" && ready2[0] == "" && ready3[0] != "") || (ready3[0] != "" && (time1 == 3 || time1 == 0) && time2 == 3 && time3 < 3) {
		p := deleteQueue(ready3)
		if p == "" {
			return
		}
		insertQueue2(cpu[i-1])
		cpu[i-1] = p
		time3++
		time1 = 0
		time2 = 0
		// fmt.Println(time1)
		// fmt.Println(time2)
		// fmt.Println(time3)
	} else {
		fmt.Println("NaN")
	}
}

func command_expire2(i int) {
	if time1 < 3 && time2 != 3 { //(ready1[0] == "" && time1 < 3) || (ready1[0] != "" && time1 < 3) || (ready1[0] != "" && ready2[0] == "" && ready3[0] == "")
		p := deleteQueue(ready1)
		if p == "" {
			time1++
			// fmt.Println(time1)
			// fmt.Println(time2)
			// fmt.Println(time3)
			insertQueue2(cpu[i-1])
			p := deleteQueue(ready1)
			cpu[i-1] = p
			if cpu[i-1] == "" {
				if ready2[0] != "" {
					p := deleteQueue(ready2)
					insertQueue2(cpu[i-1])
					cpu[i-1] = p
					time2++
					time1 = 0
					// fmt.Println(time1)
					// fmt.Println(time2)
					// fmt.Println(time3)
				} else if ready2[0] == "" && ready3[0] != "" {
					p := deleteQueue(ready2)
					insertQueue2(cpu[i-1])
					cpu[i-1] = p
					time3++
					time2 = 0
					time1 = 0
					// fmt.Println(time1)
					// fmt.Println(time2)
					// fmt.Println(time3)
				}
			}
			return
		}
		insertQueue2(cpu[i-1])
		cpu[i-1] = p
		time1++
		// fmt.Println(time1)
		// fmt.Println(time2)
		// fmt.Println(time3)
	} else if time1 == 3 && time2 < 3 { //(ready1[0] == "" && ready2[0] != "" && time1 == 3 && time2 < 3 ) || (time1 == 3 && time2 < 3)
		p := deleteQueue(ready2)
		if p == "" {
			time2++
			time1 = 0
			// fmt.Println(time1)
			// fmt.Println(time2)
			// fmt.Println(time3)
			insertQueue2(cpu[i-1])
			p := deleteQueue(ready2)
			cpu[i-1] = p
			if cpu[i-1] == "" {
				if ready1[0] != "" {
					p := deleteQueue(ready1)
					insertQueue2(cpu[i-1])
					cpu[i-1] = p
					time1++
					// fmt.Println(time1)
					// fmt.Println(time2)
					// fmt.Println(time3)
				} else if ready1[0] == "" && ready3[0] != "" {
					p := deleteQueue(ready2)
					insertQueue2(cpu[i-1])
					cpu[i-1] = p
					time3++
					time2 = 0
					time1 = 0
					// fmt.Println(time1)
					// fmt.Println(time2)
					// fmt.Println(time3)
				}
			}
			return
		}
		insertQueue2(cpu[i-1])
		cpu[i-1] = p
		time2++
		time1 = 0
		// fmt.Println(time1)
		// fmt.Println(time2)
		// fmt.Println(time3)
	} else if time1 == 0 && time2 >= 3 && ready3[0] != "" { //(ready1[0] == "" && ready2[0] == "" && ready3[0] != "") || ((time1 == 3 || time1 == 0) && time2 == 3 && time3 < 3)
		p := deleteQueue(ready3)
		if p == "" {
			time3++
			time2 = 0
			time1 = 0
			// fmt.Println(time1)
			// fmt.Println(time2)
			// fmt.Println(time3)
			insertQueue2(cpu[i-1])
			p := deleteQueue(ready2)
			cpu[i-1] = p
			if cpu[i-1] == "" {
				if ready1[0] != "" {
					p := deleteQueue(ready1)
					insertQueue2(cpu[i-1])
					cpu[i-1] = p
					time1++
					// fmt.Println(time1)
					// fmt.Println(time2)
					// fmt.Println(time3)
				} else if ready1[0] == "" && ready2[0] != "" {
					p := deleteQueue(ready2)
					insertQueue2(cpu[i-1])
					cpu[i-1] = p
					time2++
					time1 = 0
					// fmt.Println(time1)
					// fmt.Println(time2)
					// fmt.Println(time3)
				}
			}
			return
		}
		insertQueue2(cpu[i-1])
		cpu[i-1] = p
		time3++
		time1 = 0
		time2 = 0
		// fmt.Println(time1)
		// fmt.Println(time2)
		// fmt.Println(time3)
	} else if time1 == 0 && time2 >= 3 && ready3[0] == "" {
		time2 = 0
		if ready1[0] != "" {
			p := deleteQueue(ready1)
			if p == "" {
				time1++
				// fmt.Println(time1)
				// fmt.Println(time2)
				// fmt.Println(time3)
				insertQueue2(cpu[i-1])
				p := deleteQueue(ready1)
				cpu[i-1] = p
				if cpu[i-1] == "" {
					if ready2[0] != "" {
						p := deleteQueue(ready2)
						insertQueue2(cpu[i-1])
						cpu[i-1] = p
						time2++
						time1 = 0
						// fmt.Println(time1)
						// fmt.Println(time2)
						// fmt.Println(time3)
					} else if ready2[0] == "" && ready3[0] != "" {
						p := deleteQueue(ready2)
						insertQueue2(cpu[i-1])
						cpu[i-1] = p
						time3++
						time2 = 0
						time1 = 0
						// fmt.Println(time1)
						// fmt.Println(time2)
						// fmt.Println(time3)
					}
				}
				return
			}
			insertQueue2(cpu[i-1])
			cpu[i-1] = p
			time1++
			// fmt.Println(time1)
			// fmt.Println(time2)
			// fmt.Println(time3)
		} else if ready1[0] == "" && ready2[0] != "" {
			p := deleteQueue(ready2)
			if p == "" {
				time2++
				time1 = 0
				// fmt.Println(time1)
				// fmt.Println(time2)
				// fmt.Println(time3)
				insertQueue2(cpu[i-1])
				p := deleteQueue(ready2)
				cpu[i-1] = p
				return
			}
			insertQueue2(cpu[i-1])
			cpu[i-1] = p
			time2++
			time1 = 0
			// fmt.Println(time1)
			// fmt.Println(time2)
			// fmt.Println(time3)
		}
	} else {
		fmt.Println("NaN")
		// fmt.Println(time1)
		// fmt.Println(time2)
		// fmt.Println(time3)
	}
}

func command_io1(i int) {
	insertQueue(io1, cpu[i-1])
	cpu[i-1] = ""
	command_expire(i)
}

func command_io2(i int) {
	insertQueue(io2, cpu[i-1])
	cpu[i-1] = ""
	command_expire(i)
}

func command_io3(i int) {
	insertQueue(io3, cpu[i-1])
	cpu[i-1] = ""
	command_expire(i)
}

func command_io4(i int) {
	insertQueue(io4, cpu[i-1])
	cpu[i-1] = ""
	command_expire(i)
}

func command_io1x() {
	p := deleteQueue(io1)
	if p == "" {
		return
	}
	if cpu[0] == "" && (cpu[1] == "" || cpu[1] != "") {
		cpu[0] = p
	} else if cpu[0] != "" && cpu[1] == "" {
		cpu[1] = p
	} else {
		insertQueue2(p)
	}
}

func command_io2x() {
	p := deleteQueue(io2)
	if p == "" {
		return
	}
	if cpu[0] == "" && (cpu[1] == "" || cpu[1] != "") {
		cpu[0] = p
	} else if cpu[0] != "" && cpu[1] == "" {
		cpu[1] = p
	} else {
		insertQueue2(p)
	}
}

func command_io3x() {
	p := deleteQueue(io3)
	if p == "" {
		return
	}
	if cpu[0] == "" && (cpu[1] == "" || cpu[1] != "") {
		cpu[0] = p
	} else if cpu[0] != "" && cpu[1] == "" {
		cpu[1] = p
	} else {
		insertQueue2(p)
	}
}

func command_io4x() {
	p := deleteQueue(io4)
	if p == "" {
		return
	}
	if cpu[0] == "" && (cpu[1] == "" || cpu[1] != "") {
		cpu[0] = p
	} else if cpu[0] != "" && cpu[1] == "" {
		cpu[1] = p
	} else {
		insertQueue2(p)
	}
}

func insertQueue(q []string, data string) {
	for i := range q {
		if q[i] == "" {
			q[i] = data
			break
		}
	}
}

func insertQueue2(data string) {
	pr := strings.Split(data, "-")

	if len(pr) > 1 {
		switch pr[1] {
		case "1":
			for i := range ready1 {
				if ready1[i] == "" {
					ready1[i] = data
					break
				}
			}
		case "2":
			for i := range ready2 {
				if ready2[i] == "" {
					ready2[i] = data
					break
				}
			}
		case "3":
			for i := range ready3 {
				if ready3[i] == "" {
					ready3[i] = data
					break
				}
			}
		default:
			fmt.Println("Priority Error.")
		}
	}
}

func deleteQueue(q []string) string {
	result := q[0]
	for i := range q {
		if i == 0 {
			continue
		}
		q[i-1] = q[i]
	}
	q[9] = ""
	return result
}

func main() {
	initialized()
	for {
		showProcess()
		command := getCommand()
		commandx := strings.Split(command, " ")
		switch commandx[0] {
		case "exit":
			return
		case "new":
			for i := range commandx {
				if i == 0 {
					continue
				}
				process := strings.Split(commandx[i], "-")
				if len(process) > 1 {
					if process[1] == "1" || process[1] == "2" || process[1] == "3" {
						command_new(commandx[i])
					} else {
						fmt.Println("Priority Error!!! Please enter the priority within 1-3 range.")
					}
				} else {
					fmt.Println("Syntax Error!!! Please enter the priorty of the process within 1-3 range.")
				}
			}
		case "terminate":
			if commandx[1] != "" {
				index, _ := strconv.Atoi(commandx[1])
				command_terminate(index)
			} else {
				fmt.Println("Syntax Error!!! Please choose the number of cpu that you want to terminate, 1 or 2")
			}
		case "expire":
			if commandx[1] != "" {
				index, _ := strconv.Atoi(commandx[1])
				command_expire2(index)
			} else {
				fmt.Println("Syntax Error!!! Please choose the number of cpu that you want to expire, 1 or 2")
			}
		case "io1":
			if commandx[1] != "" {
				index, _ := strconv.Atoi(commandx[1])
				command_io1(index)
			} else {
				fmt.Println("Syntax Error!!! Please choose the number of cpu, 1 or 2")
			}
		case "io2":
			if commandx[1] != "" {
				index, _ := strconv.Atoi(commandx[1])
				command_io2(index)
			} else {
				fmt.Println("Syntax Error!!! Please choose the number of cpu, 1 or 2")
			}
		case "io3":
			if commandx[1] != "" {
				index, _ := strconv.Atoi(commandx[1])
				command_io3(index)
			} else {
				fmt.Println("Syntax Error!!! Please choose the number of cpu, 1 or 2")
			}
		case "io4":
			if commandx[1] != "" {
				index, _ := strconv.Atoi(commandx[1])
				command_io4(index)
			} else {
				fmt.Println("Syntax Error!!! Please choose the number of cpu, 1 or 2")
			}
		case "io1x":
			command_io1x()
		case "io2x":
			command_io2x()
		case "io3x":
			command_io3x()
		case "io4x":
			command_io4x()
		default:
			fmt.Printf("\nSorry !!! Command Error !!!\n")
		}
	}
}
