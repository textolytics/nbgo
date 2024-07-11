// //--------------------Methods-------------------

// //--------------------Hands-on exercise #1-------------------

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var counter int64
	max := 1

	var counter2 int64

	const gs = 80
	const gs1 = 40

	var wg sync.WaitGroup
	var wg1 sync.WaitGroup

	wg.Add(gs)
	wg1.Add(gs1)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		if runtime.NumGoroutine() > max {
			max = runtime.NumGoroutine()
			// fmt.Println(max)
		}
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	for j := 0; j < gs1; j++ {
		if runtime.NumGoroutine() > max {
			max = runtime.NumGoroutine()
			// fmt.Println("MaxGoroutines:", max)
		}
		atomic.AddInt64(&counter2, 1)
		runtime.Gosched()
		fmt.Println("Counter-2\t", atomic.LoadInt64(&counter2))
		wg1.Done()
		go func() {

			if j%2 == 0 {
				// atomic.AddInt64(&counter2, 1)
				// runtime.Gosched()
				// fmt.Println("Counter-2\t", atomic.LoadInt64(&counter2))
				// wg1.Done()

			}

		}()

		// fmt.Println("MaxGoroutines:", max)
	}

	wg.Wait()
	wg1.Wait()
	fmt.Println("MaxGoroutines:", max)
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
	fmt.Println("count2:", counter2)

}

//--------------------Atomic-------------------

// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// 	"sync/atomic"
// )

// func main() {
// 	fmt.Println("CPUs:", runtime.NumCPU())
// 	fmt.Println("Goroutines:", runtime.NumGoroutine())

// 	var counter int64
// 	max := 1

// 	const gs = 10
// 	var wg sync.WaitGroup
// 	wg.Add(gs)

// 	for i := 0; i < gs; i++ {
// 		go func() {
// 			atomic.AddInt64(&counter, 1)
// 			runtime.Gosched()
// 			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
// 			wg.Done()
// 		}()
// 		if runtime.NumGoroutine() > max {
// 			max = runtime.NumGoroutine()
// 			fmt.Println(max)
// 		}
// 		fmt.Println("Goroutines:", runtime.NumGoroutine())
// 	}

// 	for j := 0; j < gs; j++ {
// 		if runtime.NumGoroutine() > max {
// 			max = runtime.NumGoroutine()
// 			fmt.Println("MaxGoroutines:", max)
// 		}

// 		go func() {

// 			if j%2 == 0 {
// 				fmt.Println(j)
// 			}

// 		}()

// 		fmt.Println("MaxGoroutines:", max)
// 	}

// 	wg.Wait()
// 	fmt.Println(max)

// 	fmt.Println("Goroutines:", runtime.NumGoroutine())
// 	fmt.Println("count:", counter)
// }

//--------------------Mutex-------------------

// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// )

// func main() {
// 	fmt.Println("CPUs:", runtime.NumCPU())
// 	fmt.Println("Goroutines:", runtime.NumGoroutine())

// 	counter := 0
// 	max := 1

// 	const gs = 10000
// 	var wg sync.WaitGroup
// 	wg.Add(gs)

// 	var mu sync.Mutex

// 	for i := 0; i < gs; i++ {
// 		go func() {
// 			mu.Lock()
// 			v := counter
// 			// time.Sleep(time.Second)
// 			runtime.Gosched()
// 			v++
// 			counter = v
// 			mu.Unlock()
// 			wg.Done()
// 		}()
// 		if runtime.NumGoroutine() > max {
// 			max = runtime.NumGoroutine()
// 			fmt.Println(max)
// 		}
// 		fmt.Println("Goroutines:", runtime.NumGoroutine())
// 	}
// 	wg.Wait()
// 	fmt.Println(max)

// 	fmt.Println("Goroutines:", runtime.NumGoroutine())
// 	fmt.Println("count:", counter)
// }

//--------------------Race condition-------------------

// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// )

// func main() {
// 	fmt.Println("CPUs:", runtime.NumCPU())
// 	fmt.Println("Goroutines:", runtime.NumGoroutine())

// 	counter := 0

// 	const gs = 10000
// 	var wg sync.WaitGroup
// 	wg.Add(gs)
// 	max := 1

// 	for i := 0; i < gs; i++ {
// 		go func() {
// 			v := counter
// 			// time.Sleep(time.Second)
// 			runtime.Gosched()
// 			v++
// 			counter = v
// 			wg.Done()
// 		}()
// 		if runtime.NumGoroutine() > max {
// 			max = runtime.NumGoroutine()
// 			fmt.Println(max)
// 		}
// 		fmt.Println("Goroutines:", runtime.NumGoroutine())
// 	}
// 	wg.Wait()
// 	fmt.Println(max)

// 	fmt.Println("Goroutines:", runtime.NumGoroutine())
// 	fmt.Println("count:", counter)
// }

//--------------------Method sets revisited--------------------

// package main

// import (
// 	"fmt"
// 	"math"
// 	"runtime"
// )

// type circle struct {
// 	radius float64
// }

// type shape interface {
// 	area() float64
// }

// func (c *circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }

// func info(s shape) {
// 	fmt.Println("area", s.area())
// }

// func main() {
// 	c := circle{5}
// 	// go info(c)

// 	fmt.Println("Goroutines\t", runtime.NumGoroutine())

// 	for i := 1; i < 100; i++ {
// 		cc := 1.0
// 		c = circle{cc}
// 		fmt.Println("Goroutines\t", runtime.NumGoroutine())

// 		fmt.Println(c.area())

// 	}
// 	fmt.Println("Goroutines\t", runtime.NumGoroutine())

// 	// info(c)
// 	fmt.Println(c.area())
// }

//--------------------CONCURENCY--------------------

// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// )

// var wg sync.WaitGroup

// func main() {
// 	fmt.Println("OS\t\t", runtime.GOOS)
// 	fmt.Println("ARCH\t\t", runtime.GOARCH)
// 	fmt.Println("CPUs\t\t", runtime.NumCPU())
// 	fmt.Println("Goroutines\t", runtime.NumGoroutine())

// 	wg.Add(1)
// 	go foo()
// 	bar()

// 	fmt.Println("CPUs\t\t", runtime.NumCPU())
// 	fmt.Println("Goroutines\t", runtime.NumGoroutine())
// 	wg.Wait()
// }

// func foo() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println("foo:", i)
// 	}
// 	wg.Done()
// }

// func bar() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println("bar:", i)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// )

// var wg sync.WaitGroup

// func main() {
// 	fmt.Println(runtime.GOOS)
// 	fmt.Println(runtime.GOARCH)
// 	fmt.Println(runtime.NumCPU())
// 	fmt.Println(runtime.NumGoroutine())

// 	wg.Add(1)
// 	foo()
// 	bar()

// 	// fmt.Println(runtime.NumCPU())
// 	fmt.Println(runtime.NumGoroutine())
// 	wg.Wait()
// }

// func foo() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println("foo:", i)
// 	}
// 	wg.Done()

// }

// func bar() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println("bar:", i)
// 	}
// }

//--------------------EX5--------------------

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// type user struct {
// 	First   string
// 	Last    string
// 	Age     int
// 	Sayings []string
// }

// type byAge []user

// func (a byAge) Len() int           { return len(a) }
// func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// type byLast []user

// func (a byLast) Len() int           { return len(a) }
// func (a byLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a byLast) Less(i, j int) bool { return a[i].Last < a[j].Last }

// func main() {
// 	u1 := user{
// 		First: "James",
// 		Last:  "Bond",
// 		Age:   32,
// 		Sayings: []string{
// 			"Shaken, not stirred",
// 			"Youth is no guarantee of innovation",
// 			"In his majesty's royal service",
// 		},
// 	}

// 	u2 := user{
// 		First: "Miss",
// 		Last:  "Moneypenny",
// 		Age:   27,
// 		Sayings: []string{
// 			"James, it is soo good to see you",
// 			"Would you like me to take care of that for you, James?",
// 			"I would really prefer to be a secret agent myself.",
// 		},
// 	}

// 	u3 := user{
// 		First: "M",
// 		Last:  "Hmmmm",
// 		Age:   54,
// 		Sayings: []string{
// 			"Oh, James. You didn't.",
// 			"Dear God, what has James done now?",
// 			"Can someone please tell me where James Bond is?",
// 		},
// 	}

// 	users := []user{u1, u2, u3}

// 	// fmt.Println(users)

// 	for i, user := range users {
// 		fmt.Println("Person # ", i)
// 		fmt.Println("\t", user.First, user.Last, user.Age)
// 		for _, v := range user.Sayings {
// 			fmt.Println("\t Tweet : ", v)
// 		}
// 	}

// 	sort.Sort(byAge([]user(users)))

// 	for i, user := range users {
// 		fmt.Println("Person # ", i)
// 		fmt.Println("\t", user.First, user.Last, user.Age)
// 		for _, v := range user.Sayings {
// 			fmt.Println("\t Tweet : ", v)
// 		}
// 	}

// 	sort.Sort(byLast([]user(users)))

// 	for i, user := range users {
// 		fmt.Println("Person # ", i)
// 		fmt.Println("\t", user.First, user.Last, user.Age)
// 		for _, v := range user.Sayings {
// 			fmt.Println("\t Tweet : ", v)
// 		}
// 	}

// }

//--------------------EX4--------------------

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
// 	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

// 	fmt.Println(xi)
// 	// sort xi
// 	sort.Ints(xi)
// 	fmt.Println(xi)

// 	fmt.Println(xs)
// 	// sort xs
// 	sort.Strings(xs)
// 	fmt.Println(xs)

// }

//--------------------EX3--------------------
// package main

// import (
// 	"encoding/json"
// 	"os"
// )

// type user struct {
// 	First   string
// 	Last    string
// 	Age     int
// 	Sayings []string
// }

// func main() {
// 	u1 := user{
// 		First: "James",
// 		Last:  "Bond",
// 		Age:   32,
// 		Sayings: []string{
// 			"Shaken, not stirred",
// 			"Youth is no guarantee of innovation",
// 			"In his majesty's royal service",
// 		},
// 	}

// 	u2 := user{
// 		First: "Miss",
// 		Last:  "Moneypenny",
// 		Age:   27,
// 		Sayings: []string{
// 			"James, it is soo good to see you",
// 			"Would you like me to take care of that for you, James?",
// 			"I would really prefer to be a secret agent myself.",
// 		},
// 	}

// 	u3 := user{
// 		First: "M",
// 		Last:  "Hmmmm",
// 		Age:   54,
// 		Sayings: []string{
// 			"Oh, James. You didn't.",
// 			"Dear God, what has James done now?",
// 			"Can someone please tell me where James Bond is?",
// 		},
// 	}

// 	users := []user{u1, u2, u3}

// 	// fmt.Println(users)

// 	json.NewEncoder(os.Stdout).Encode(users)

// 	// your code goes here

// }

//--------------------EX2--------------------

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type tweet struct {
// 	First   string   `json:"First"`
// 	Last    string   `json:"Last"`
// 	Age     int      `json:"Age"`
// 	Sayings []string `json:"Sayings"`
// }

// var tweets []tweet

// func main() {
// 	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
// 	err := json.Unmarshal([]byte(s), &tweets)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(tweets)
// 	for i, tweet := range tweets {
// 		fmt.Println("Person # ", i)
// 		fmt.Println("\t", tweet.First, tweet.Last, tweet.Age)
// 		for _, v := range tweet.Sayings {
// 			fmt.Println("\t Tweet : ", v)
// 		}
// 	}

// 	// os.Stdout.Write(tweet)
// }

//--------------------EX1--------------------

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type user struct {
// 	First string
// 	Age   int
// }

// func main() {
// 	u1 := user{
// 		First: "James",
// 		Age:   32,
// 	}

// 	u2 := user{
// 		First: "Moneypenny",
// 		Age:   27,
// 	}

// 	u3 := user{
// 		First: "M",
// 		Age:   54,
// 	}

// 	users := []user{u1, u2, u3}

// 	fmt.Println(users)

// 	bs, err := json.Marshal(users)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println(string(bs))
// }
//--------------------EX1--------------------

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

// type user struct {
// 	First string
// 	Age   int
// }

// func main() {
// 	u1 := user{
// 		First: "James",
// 		Age:   32,
// 	}

// 	u2 := user{
// 		First: "Moneypenny",
// 		Age:   27,
// 	}

// 	u3 := user{
// 		First: "M",
// 		Age:   54,
// 	}

// 	users := []user{u1, u2, u3}

// 	fmt.Println(users)

// 	jsonUsers, _ := json.Marshal(users)
// 	os.Stdout.Write(jsonUsers)

// for _, v := range []users {
// 	json.Marshal(v)

// }
// your code goes here
// }

//--------------------EX1--------------------

// package main

// import (
// 	"fmt"
// )

// type person []struct {
// 	First   string   `json:"First"`
// 	Last    string   `json:"Last"`
// 	Age     int      `json:"Age"`
// 	Sayings []string `json:"Sayings"`
// }

// func main() {
// 	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
// 	fmt.Println(s)

// 	bs := []person{}
// 	fmt.Println(bs)

// }

//--------------------BCRYPT--------------------
// package main

// import (
// 	"fmt"

// 	"golang.org/x/crypto/bcrypt"
// )

// func main() {
// 	s := `password123`
// 	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(s)
// 	fmt.Println(bs)
// 	fmt.Println(string(bs))

// 	loginPword1 := `password1234`

// 	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPword1))
// 	if err != nil {
// 		fmt.Println("YOU CAN'T LOGIN")
// 		return
// 	}

// 	fmt.Println("You're logged in")
// }

//--------------------SORT-CUSTOM--------------------

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// type Person struct {
// 	First string
// 	Age   int
// }

// type ByAge []Person

// func (a ByAge) Len() int           { return len(a) }
// func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// func main() {
// 	p1 := Person{"James", 32}
// 	p2 := Person{"Moneypenny", 27}
// 	p3 := Person{"Q", 64}
// 	p4 := Person{"M", 56}

// 	people := []Person{p1, p2, p3, p4}

// 	fmt.Println(people)
// 	sort.Sort(ByAge(people))
// 	fmt.Println(people)

// }

//--------------------SORT---------------------
// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
// 	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

// 	fmt.Println(xi)
// 	sort.Ints(xi)
// 	fmt.Println(xi)

// 	fmt.Println("------")
// 	fmt.Println(xs)
// 	sort.Strings(xs)
// 	fmt.Println(xs)

// }

//--------------------JSON--Unmarshal---------------------

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type person struct {
// 	First string `json:"First"`
// 	Last  string `json:"Last"`
// 	Age   int    `json:"Age"`
// }

// func main() {

// 	s := `[{"First":"Miss","Last":"MPen","Age":27},{"First":"James","Last":"Bio","Age":32}]`
// 	bs := []byte(s)

// 	fmt.Printf("%T\n", s)
// 	fmt.Printf("%T\n", bs)

// 	var people []person
// 	err := json.Unmarshal(bs, &people)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println("----------ALL-DATA---------\n", people)

// 	for i, v := range people {

// 		fmt.Println("\n PERSON NUMBER : ", i)
// 		fmt.Println("\t", v.First, v.Last, v.Age)
// 	}

// }

//--------------------JSON--Marshal---------------------

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type person struct {
// 	First string
// 	Last  string
// 	Age   int
// }

// func main() {

// 	p1 := person{
// 		First: "Miss",
// 		Last:  "MPen",
// 		Age:   27,
// 	}

// 	p2 := person{
// 		First: "James",
// 		Last:  "Bio",
// 		Age:   32,
// 	}

// 	people := []person{p1, p2}

// 	fmt.Println(people)

// 	bs, err := json.Marshal(people)

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(string(bs))
// }

//--------------------JSON-----------------------

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// func main() {

// 	var jsonBlob = []byte(`[
// 	{"Name": "Platypus", "Order": "Monotremata"},
// 	{"Name": "Quoll",    "Order": "Dasyuromorphia"}
// ]`)
// 	type Animal struct {
// 		Name  string
// 		Order string
// 	}
// 	var animals []Animal
// 	err := json.Unmarshal(jsonBlob, &animals)
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}
// 	fmt.Printf("%+v", animals)
// }

//--------------------JSON-----------------------

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	type ColorGroup struct {
// 		ID     int
// 		Name   string
// 		Colors []string
// 	}
// 	group := ColorGroup{
// 		ID:     1,
// 		Name:   "Reds",
// 		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
// 	}
// 	b, err := json.Marshal(group)
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}
// 	os.Stdout.Write(b)
// }

//--------------------EX4-----------------------
// package main

// import "fmt"

// type vehicle struct {
// 	color string
// 	doors int8
// }

// type truck struct {
// 	vehicle
// 	fourWheel bool
// }

// func main() {

// 	t1 := truck{
// 		vehicle: vehicle{
// 			color: "Blue",
// 			doors: 3,
// 		},
// 		fourWheel: true,
// 	}

// 	change(&t1)
// 	fmt.Println(t1.color)

// }

// func change(v *truck) {

// 	v.color = "Black"

// }

//----------------------EX3---------------------
// package main

// import "fmt"

// type vehicle struct {
// 	color string
// 	doors int8
// }

// type truck struct {
// 	vehicle
// 	fourWheel bool
// }

// type sedan struct {
// 	vehicle
// 	luxury bool
// }

// func main() {

// 	t1 := truck{
// 		vehicle: vehicle{
// 			color: "Blue",
// 			doors: 3,
// 		},
// 		fourWheel: true,
// 	}

// 	s1 := sedan{
// 		vehicle: vehicle{
// 			color: "Green",
// 			doors: 5,
// 		},
// 		luxury: true,
// 	}

// 	t2 := struct {
// 		vehicle
// 		fourWheel bool
// 	}{
// 		vehicle: vehicle{
// 			color: "Cyan",
// 			doors: 3,
// 		},
// 		fourWheel: true,
// 	}
// 	fmt.Println(t2)
// 	fmt.Println(t1)
// 	fmt.Println(s1)
// 	fmt.Println(t1.doors)
// 	fmt.Println(s1.color)
//--------------------EX2-------------------------
// package main

// import (
// 	"fmt"
// )

// type person struct {
// 	first      string
// 	last       string
// 	favFlavors []string
// }

// func main() {
// 	p1 := person{
// 		first: "James",
// 		last:  "Bond",
// 		favFlavors: []string{
// 			"chocolate",
// 			"martini",
// 			"rum and coke",
// 		},
// 	}

// 	p2 := person{
// 		first: "Miss",
// 		last:  "Moneypenny",
// 		favFlavors: []string{
// 			"strawberry",
// 			"vanilla",
// 			"capuccino",
// 		},
// 	}

// 	m := map[string]person{
// 		p1.last: p1,
// 		p2.last: p2,
// 	}

// 	for _, v := range m {
// 		fmt.Println(v.first)
// 		fmt.Println(v.last)
// 		for i, val := range v.favFlavors {
// 			fmt.Println(i, val)
// 		}
// 		fmt.Println("-------")
// 	}

//------------------EX2-------------------------------

// package main
// import "fmt"
// // var testarray = []int{2, 3, 4, 5, 42}
// // const x = 42
// // const y int64 = 32
// // var invar int = 42
// type person struct {
// 	first string
// 	last  string
// 	age   int
// }
// type sSage struct {
// 	person
// 	asd bool
// }
// type iceperson struct {
// 	first  string
// 	last   string
// 	flavor []string
// }
// func main() {
// 	ip1 := iceperson{
// 		first:  "Firstname",
// 		last:   "FirstLast",
// 		flavor: []string{"Vanilla", "Caramel"},
// 	}
// 	ip2 := iceperson{
// 		first:  "Secondtname",
// 		last:   "SecondLast",
// 		flavor: []string{"VaniCaramrlla", "Mint"},
// 	}
// 	fmt.Println("Ice Cream\n", ip1, ip2)
// 	m := map[string][]string{}
// 	m[ip1.last] = ip1.flavor
// 	m[ip2.last] = ip2.flavor
// 	fmt.Println("Range :")
// 	for p, f := range m {
// 		fmt.Printf("%v \t %v\n", p, f)
// 		for i, v := range f {
// 			fmt.Printf("\t %v \t %v\n", i, v)
// 		}
// 	}
// m1 := map[string]person{
// 	p1.last: p1,
// 	p2.last: p2,
// }
// fmt.Println(m1)
// fmt.Println("MAP RANGE\n")

// for _, v := range m1 {
// 	fmt.Println(v.first, v.last)
// 	for i, s := range v.favFlavors {
// 		fmt.Println(i, s)
// 	}
// }

// 	for p, f := range ip1.flavor {
// 		fmt.Printf("\t %v \t %v\t %v \t %v\n", ip1.first, ip1.last, p, f)
// 	}
// 	for p, f := range ip2.flavor {
// 		fmt.Printf("\t %v \t %v\t %v \t %v\n", ip2.first, ip2.last, p, f)
// 	}
// 	p1 := person{
// 		first: "James",
// 		last:  "Bond",
// 		age:   22,
// 	}
// 	p2 := person{
// 		first: "Miss",
// 		last:  "Moneypenny",
// 		age:   44,
// 	}
// 	s1 := sSage{
// 		person: person{
// 			first: "Name",
// 			last:  "Last",
// 			age:   32,
// 		},
// 		asd: true,
// 	}
// 	s2 := struct {
// 		person
// 		asd bool
// 	}{
// 		person: person{
// 			first: "Name2",
// 			last:  "Last2",
// 			age:   36,
// 		},
// 		asd: true,
// 	}
// 	fmt.Println(p1)
// 	fmt.Println(p2)
// 	fmt.Println(s1)
// 	fmt.Println(s2)
// 	fmt.Println(p1.first, p1.last, p1.age)
// 	fmt.Println(p2.first, p2.last, p2.age)
// 	fmt.Println(s1.person.first, s1.last, s1.age, s1.asd)
// 	fmt.Println(s2.person.first, s2.last, s2.age, s2.asd)
//-------------------MAP--------------------------
// m := map[string][]string{
// 	`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
// 	`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
// 	`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
// }
// for k, v := range m {
// 	fmt.Println("This is the record for", k)
// 	for i, v2 := range v {
// 		fmt.Println("\t", i, v2)
// 	}
// }
// m["New one"] = []string{"Bar", "Code", "Sleep"}
// for k, v := range m {
// 	fmt.Println("This is for", k)
// 	for i, v2 := range v {
// 		fmt.Println("\t", i, v2)
// 	}
// }
// if v, ok := m["no_dr"]; ok {
// 	delete(m, "no_dr")
// 	fmt.Println("Deleted", v, ok)
// }
// for k, v := range m {
// 	fmt.Println("This is for", k)
// 	for i, v2 := range v {
// 		fmt.Println("\t", i, v2)
// 	}
// }
// a := []string{`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`, `moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`, `no_dr`, `Being evil`, `Ice cream`, `Sunsets`}
// for i, v := range a {
// 	fmt.Println(i, a[i], v)
// }
// m := map[string]string{}
// fmt.Println(m)
// mapKey := a[0]
// for i, v := range a {
// 	if i%4 == 0 {
// 		mapKey = v
// 	}
// 	m[mapKey] = v
// 	fmt.Println(i, v, mapKey)
// }
// fmt.Println(m)
// m := map[string]int{
// 	"James": 32,
// 	"Miss":  29,
// }
// fmt.Println(m)
// fmt.Println(m["James"])
// fmt.Println(m["Miss"])
// fmt.Println(m["Abc"])
// v, ok := m["Abc"]
// fmt.Println(v)
// fmt.Println(ok)
// m["Todd"] = 33
// if v, ok := m["Miss"]; ok {
// 	fmt.Println("THIS IS IF", v)
// }
// for k, v := range m {
// 	fmt.Println(k, v)
// }
// xi := []int{5, 6, 7, 8, 99, 4}
// for i, v := range xi {
// 	fmt.Println(i, v)
// }
// fmt.Println(m)
// delete(m, "James")
// fmt.Println(m)
// delete(m, "XXX")
// fmt.Println(m)
// if v, ok := m["Miss"]; ok {
// 	fmt.Println("THIS IS IF Miss exist", v)
// 	delete(m, "Miss")
// }
// fmt.Println(m)
// -----------2 Dimentional Slices-------------
// jb := []string{"James", "Bo", "Cho", "Mart"}
// fmt.Println(jb)
// mp := []string{"Miss", "Moneypen", "Jorh", "Jsds"}
// fmt.Println(mp)
// xp := [][]string{jb, mp}
// fmt.Println(xp)
// }
// mslice := make([]int, 10, 100)
// fmt.Println(mslice)
// fmt.Println(len(mslice))
// fmt.Println(cap(mslice))
// mslice[0] = 42
// mslice[9] = 888
// fmt.Println(mslice)
// fmt.Println(len(mslice))
// fmt.Println(cap(mslice))
// mslice = append(mslice, 3454)
// fmt.Println(mslice)
// fmt.Println(len(mslice))
// fmt.Println(cap(mslice))
// ex := mslice
// fmt.Println(ex)
// ex = append(ex, 33, 4, 4, 5, 6, 7)
// fmt.Println(ex)
// y := []int{234, 234, 234, 4}
// ex = append(ex, y...)
// fmt.Println(ex)
// ex = append(ex[8:], ex[10:]...)
// fmt.Println(ex)
// fmt.Println(mslice)
// fmt.Println(len(ex))
// fmt.Println(cap(ex))
// for i, v := range testarray {
// 	fmt.Println(i, v)
// }
// fmt.Printf("%T", testarray)
// a := `erwerw
// erw
// er
// we
// rwe
// r`
// fmt.Printf("%d\t%b\t%#x\t\n", invar, invar, invar)
// // bivar : = fmt.Sprintf("%d", invar)
// bshift := invar << 1
// fmt.Printf("%d\t%b\t%#x", bshift, bshift, bshift)
// // a := (32 == 23)
// b := (22 <= 2)
// c := (3 >= 3)
// d := (33 != 2)
// e := (3 < 3)
// fmt.Println(a, b, c, d)
// fmt.Println(y)
// fmt.Println(a)
//
// }
