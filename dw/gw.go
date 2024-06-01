package core

type Gateway interface {
	Type(interface{})    // Type [ EMS | DW | API | Local | Core | Util | 3rd party Library |File|Schema]
	Pattern(interface{}) // Pattern [ PUB SUB | PULL PUSH | Local | Client | Server ]
	Context(Event)       // Context or Event types , Packet Length, Encoding, Etc.
	Config(struct{})     // Settings, Address , Buffer , Timeout, Heartbit
}
