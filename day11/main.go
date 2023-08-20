package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type MonkeyTest struct {
	divible int
	t       int
	f       int
}

type Monkey struct {
	id    int
	items []int
	op    [3]string
	mtest MonkeyTest
}

func (m *Monkey) SetItems(items []int) *Monkey {
	m.items = items
	return m
}

func (m *Monkey) SetOp(op []string) *Monkey {
	m.op = [3]string{op[0], op[1], op[2]}
	return m
}

func (m *Monkey) SetMtestDiv(div int) *Monkey {
	m.mtest.divible = div
	return m
}

func (m *Monkey) SetMtestTrue(t int) *Monkey {
	m.mtest.t = t
	return m
}

func (m *Monkey) SetMtestFalse(f int) *Monkey {
	m.mtest.f = f
	return m
}

func (m *Monkey) Operation() {

}

func getInputData(input string) (*[]Monkey, error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, err
	}

	scan := bufio.NewScanner(file)

	defer file.Close()

	var (
		monkey  Monkey
		monkeys []Monkey
	)

	for scan.Scan() {

		m := scan.Text()

		switch {
		case strings.Contains(m, "Monkey "):

			id_string := strings.Split(m[:len(m)-1], " ")
			id, err := strconv.Atoi(id_string[1])
			if err != nil {
				return nil, err
			}

			monkey = Monkey{id: id, mtest: MonkeyTest{}}

		case strings.Contains(m, "  Starting items:"):
			items_str := strings.Split(m[18:], ", ")
			items := []int{}
			for _, i := range items_str {

				item, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				items = append(items, item)
				monkey.SetItems(items)
			}

		case strings.Contains(m, "  Operation:"):
			op := strings.Split(m[19:], " ")
			monkey.SetOp(op)

		case strings.Contains(m, "  Test: divisible"):
			div, err := strconv.Atoi(m[21:])
			if err != nil {
				return nil, err
			}
			monkey.SetMtestDiv(div)

		case strings.Contains(m, "true"):
			t, err := strconv.Atoi(m[29:])
			if err != nil {
				return nil, err
			}
			monkey.SetMtestTrue(t)

		case strings.Contains(m, "false"):
			f, err := strconv.Atoi(m[30:])
			if err != nil {
				return nil, err
			}
			monkey.SetMtestFalse(f)

		default:
			monkeys = append(monkeys, monkey)
		}
	}

	// }()
	return &monkeys, nil
}

func main() {
	monkeys, err := getInputData("./day11/input_test")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(monkeys)
}
