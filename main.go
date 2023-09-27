// Решение для 1 задания 
package main

import "fmt"

func main() {
	var heightPetals, heightStem int
	fmt.Scan(&heightPetals, &heightStem)
	drawFlower(heightPetals, heightStem)
}

func drawFlower(heightPetals, heightStem int) {
	for i := 0; i < heightPetals; i++ {
		for j := 0; j < heightPetals-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 0; i < heightStem-1; i++ {
		fmt.Println("*")
	}
}

// Решение для 2 задания 

package main

import "fmt"

func main() {
    var lastNumber int
    fmt.Scanln(&lastNumber)
    for i := 3; i <= lastNumber; i+=5 {
        if isPrime(i){
            fmt.Print("хоп ")
           } else {
                fmt.Printf("%d ", i)
            }
        }
        fmt.Println()
    }
func isPrime(n int) bool {
    if n<=1 {
        return false 
    }
    if n == 2 || n == 3 {
        return true 
    }
    if n%2 == 0 {
        return false
    }
    for i := 3; i*i <= n; i += 2 {
        if n%i == 0{
            return false
        }
    }
    return true
}
// Решение для 3 задания 
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    firstLetter, _ := reader.ReadString('\n')
    firstLetter = strings.TrimSpace(firstLetter)

    secondLetter, _ := reader.ReadString('\n')
    secondLetter = strings.TrimSpace(secondLetter)

    text, _ := reader.ReadString('\n')
    text = strings.TrimSpace(text)

    firstCount := strings.Count(text, firstLetter)
    secondCount := strings.Count(text, secondLetter)

    if firstCount >= secondCount {
        fmt.Printf("%s %d\n%s %d\n", firstLetter, firstCount, secondLetter, secondCount)
    } else {
        fmt.Printf("%s %d\n%s %d\n", secondLetter, secondCount, firstLetter, firstCount)
    }
}

// Решение для 4 задания 

package main

import (
	"errors"
)

type Grasshopper struct {
	position int
	target   int
}

func (g *Grasshopper) WhereAmI() int {
	return g.position
}

func (g *Grasshopper) Jump() (int, error) {
	if g.position == g.target {
		return 0, errors.New("кузнечик уже достиг зернышка")
	}

	if g.position < g.target {
		nextJump := g.position + 5
		if nextJump > g.target {
			nextJump = g.target
		}
		g.position = nextJump
		return g.position, nil
	}

	if g.position > g.target {
		nextJump := g.position - 5
		if nextJump < g.target {
			nextJump = g.target
		}
		g.position = nextJump
		return g.position, nil
	}

	return g.position, nil
}

func PlaceJumper(place, target int) Jumper {
	return &Grasshopper{
		position: place,
		target:   target,
	}
}

type Jumper interface {
	WhereAmI() int
	Jump() (int, error)
}

// Решение для 5 задания 

package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strings"
)

func main() {
    var n int
    fmt.Scanln(&n)

    names := make([]string, n)
    scanner := bufio.NewScanner(os.Stdin)

    for i := 0; i < n && scanner.Scan(); i++ {
        names[i] = scanner.Text()
    }

    sort.Strings(names)

    for scanner.Scan() {
        prefix := scanner.Text()
        found := false

        for _, name := range names {
            if strings.HasPrefix(name, prefix) {
                fmt.Println(name)
                found = true
                break
            }
        }

        if !found {
            fmt.Println("Не найдено")
        }
    }
}
