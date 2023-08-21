package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type MonkeyTest struct {
	Div   int
	True  int
	False int
}

type Monkey struct {
	id    int
	items []int
	op    [3]string
	mtest MonkeyTest
	evals int
}

func (m *Monkey) receiveItem(item int) *Monkey {
	m.items = append(m.items, item)
	return m
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
	m.mtest.Div = div
	return m
}

func (m *Monkey) SetMtestTrue(t int) *Monkey {
	m.mtest.True = t
	return m
}

func (m *Monkey) SetMtestFalse(f int) *Monkey {
	m.mtest.False = f
	return m
}

func (m *Monkey) Operation(old int) (item int, tomonkey int) {
	i, err := strconv.Atoi(m.op[0])
	if err != nil {
		i = old
	}

	j, err := strconv.Atoi(m.op[2])
	if err != nil {
		j = old
	}

	switch {
	case m.op[1] == "+":
		item = i + j
	case m.op[1] == "-":
		item = i - j
	case m.op[1] == "*":
		item = i * j
	case m.op[1] == "/":
		item = i / j
	}

	item = item / 3

	if item%m.mtest.Div == 0 {
		return item, m.mtest.True
	}
	return item, m.mtest.False
}

func (m *Monkey) Operate(mnks *Monkeys) {
	for _, item := range m.items {
		ev, monkeyId := m.Operation(item)
		(*mnks)[monkeyId].receiveItem(ev)
		m.evals += 1
	}
	m.items = []int{}
}

type Monkeys []*Monkey

func (mnks *Monkeys) ParseMonkeys(input string) (*Monkeys, error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, err
	}

	scan := bufio.NewScanner(file)

	defer file.Close()

	var monkey *Monkey

	for scan.Scan() {

		m := scan.Text()

		switch {
		case strings.Contains(m, "Monkey "):

			id_string := strings.Split(m[:len(m)-1], " ")
			id, err := strconv.Atoi(id_string[1])
			if err != nil {
				return nil, err
			}

			monkey = &Monkey{id: id, mtest: MonkeyTest{}}

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

			*mnks = append(*mnks, monkey)
		}
	}

	return mnks, nil
}

func (mnks *Monkeys) GetMonkeyBusiness() int {

	inspects := []int{}
	for _, monkey := range *mnks {
		inspects = append(inspects, monkey.evals)
	}
	slices.Sort(inspects)
	inspects = inspects[len(inspects)-2:]

	monkey_business := inspects[0] * inspects[1]
	return monkey_business
}

func main() {
	var mnks = new(Monkeys)
	mnks, err := mnks.ParseMonkeys("./day11/input")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 20; i++ {
		for _, monkey := range *mnks {
			monkey.Operate(mnks)
		}
	}
	log.Println(mnks.GetMonkeyBusiness()) // 151312
}
