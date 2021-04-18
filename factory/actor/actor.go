/// + build GODEBUG=asyncpreemptoff=1

package actor

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	actorZmq4 "github.com/pebbe/zmq4"
)

// Zmq4pbClientDealer PULL
func Zmq4pbClientDealer() {
	fmt.Println("Zmq4pbClientDealer --- STARTING ....")

	var mu sync.Mutex
	//  This is our client task
	//  It connects to the server, and then sends a request once per second
	//  It collects responses as they arrive, and it prints them out. We will
	//  run several client tasks in parallel, each with a different random ID.

	client, _ := actorZmq4.NewSocket(actorZmq4.DEALER)
	defer client.Close()

	//  Set random identity to make tracing easier
	set_id(client)
	client.Connect("tcp://127.0.0.1:5570")

	go func() {
		for request_nbr := 1; true; request_nbr++ {
			time.Sleep(time.Second)
			mu.Lock()
			client.SendMessage(fmt.Sprintf("request #%d", request_nbr))
			mu.Unlock()
		}
	}()

	for {
		time.Sleep(30 * time.Millisecond)
		mu.Lock()
		msg, err := client.RecvMessage(actorZmq4.DONTWAIT)
		if err == nil {
			id, _ := client.GetIdentity()
			fmt.Println(msg[0], id)
		}
		mu.Unlock()
	}
}

// Zmq4pbRouterDealer
func Zmq4pbRouterDealer() {
	fmt.Println("Zmq4pbRouterDealer --- STARTING ....")
	//  This is our server task.
	//  It uses the multithreaded server model to deal requests out to a pool
	//  of workers and route replies back to clients. One worker can handle
	//  one request at a time but one client can talk to multiple workers at
	//  once.
	//  Frontend socket talks to clients over TCP
	frontend, _ := actorZmq4.NewSocket(actorZmq4.ROUTER)
	defer frontend.Close()
	frontend.Bind("tcp://127.0.0.1:5570")

	//  Backend socket talks to workers over inproc
	backend, _ := actorZmq4.NewSocket(actorZmq4.DEALER)
	defer backend.Close()
	backend.Bind("inproc://backend")

	//  Launch pool of worker threads, precise number is not critical
	for i := 0; i < 5; i++ {
		go Zmq4pbWorkerDealer()
	}

	//  Connect backend to frontend via a proxy
	err := actorZmq4.Proxy(frontend, backend, nil)
	log.Fatalln("Proxy interrupted:", err)
}

//  Each worker task works on one request at a time and sends a random number
//  of replies back, with random delays between replies:

func Zmq4pbWorkerDealer() {
	fmt.Println("Zmq4pbWorkerDealer --- STARTING ....")

	worker, _ := actorZmq4.NewSocket(actorZmq4.DEALER)
	defer worker.Close()
	worker.Connect("inproc://backend")

	for {
		//  The DEALER socket gives us the reply envelope and message
		msg, _ := worker.RecvMessage(0)
		identity, content := pop(msg)

		//  Send 0..4 replies back
		replies := rand.Intn(5)
		for reply := 0; reply < replies; reply++ {
			//  Sleep for some fraction of a second
			time.Sleep(time.Duration(rand.Intn(1000)+1) * time.Millisecond)
			worker.SendMessage(identity, content)
		}
	}
}

//  The main thread simply starts several clients, and a server, and then
//  waits for the server to finish.

func Zmq4pbAsyncsrv() {
	fmt.Println("Zmq4pbWorkerDealer --- STARTING ....")

	rand.Seed(time.Now().UnixNano())
	go Zmq4pbClientDealer()
	go Zmq4pbClientDealer()
	go Zmq4pbClientDealer()
	go Zmq4pbClientDealer()

	//  Run for 5 seconds then quit
	time.Sleep(5 * time.Second)
}

func set_id(soc *actorZmq4.Socket) {
	fmt.Println("set_id --- STARTING ....")

	identity := fmt.Sprintf("%04X-%04X", rand.Intn(0x10000), rand.Intn(0x10000))
	soc.SetIdentity(identity)
}

func pop(msg []string) (head, tail []string) {
	fmt.Println("pop --- STARTING ....")

	if msg[1] == "" {
		head = msg[:2]
		tail = msg[2:]
	} else {
		head = msg[:1]
		tail = msg[1:]
	}
	return
}

// func (Zmq4pbXSUB *actorFactory.Zmq4pbXSUB) Zmq4pbReceiveMessage(topic string) (address string, contents string) {
// 	if Zmq4pbXSUB != nil {
// 		fmt.Println(zmqClient)
// 	}

// 	actorZmq4Client.InitSubscriber.SetSubscribe(string(topic))

// 	for {
// 		//  Read envelope with address
// 		address, _ := zmq4pbXSUB.InitSubscriber.Recv(0)
// 		//  Read message contents
// 		contents, _ := zmq4pbXSUB.InitSubscriber.Recv(0)
// 		fmt.Printf("[%s] %s\n", address, contents)
// 		// spew.Dump("%+v\n", contents)
// 		// return address, contents
// 	}

// 	return address, contents
// }
