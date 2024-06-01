package marketdata

type Route struct {
	Type     (interface{}) // Incoming , Outgoing, Avoiding, Mixed, Bonding, Multicast, Unicat
	Inbound  (Gateway)
	Outbound (Gateway)
}

// type Downstream interface {

// }

// type Upstream interface {
// 	Incoming(g gateway, r OutboundEvent.Request);  (gateway, SourceToRouter)
// 	Outgoing()
// }

// func InboundUpstreamReceiver(o OutboundEvent, s SourceToRouter) (u *Upstream) {
// 	if err := InboundUpstreamReceiver(o, s); err != nil {
// 		return err
// 	}
// 	return u
// }

// func OutboundUpstreamSender(i InboundEvent, r RouterToDestination) (u *Upstream) {
// 	if err := OutboundUpstreamSender(i, r); err != nil {
// 		return err
// 	}
// 	return u
// }

// type SourceToRouter struct {
// 	SourceOutboundGateID     int
// 	SourceOutboundGate       gateway
// 	SourceOutboundGateStatus bool

// 	RouterInboundGateID     int
// 	RouterInboundGate       gateway
// 	RouterInboundGateStatus bool
// }

// type RouterToDestination struct {
// 	RouterOutboundGateID     int
// 	RouterOutboundGate       gateway
// 	RouterOutboundGateStatus bool

// 	DestinationInboundGateID     int
// 	DestinationInboundGate       gateway
// 	DestinationInboundGateStatus bool
// }

// // type path struct {
// 	Status bool
// 	Hops   []int
// 	Route  struct{}
// }

// RouterOutboundGate ID
// RouterOutboundGate Client
// RouterOutboundGate Status

// DestinnationInboundGate ID
// DestinnationInboundGate Client
// DestinnationInboundGate Status

// package core

// type routerConfig struct {
// }

// type route struct {
// }

// type Downstream interface {

// }

// type Upstream interface {
// 	Incoming(g gateway, r OutboundEvent.Request);  (gateway, SourceToRouter)
// 	Outgoing()
// }

// func InboundUpstreamReceiver(o OutboundEvent, s SourceToRouter) (u *Upstream) {
// 	if err := InboundUpstreamReceiver(o, s); err != nil {
// 		return err
// 	}
// 	return u
// }

// func OutboundUpstreamSender(i InboundEvent, r RouterToDestination) (u *Upstream) {
// 	if err := OutboundUpstreamSender(i, r); err != nil {
// 		return err
// 	}
// 	return u
// }

// type SourceToRouter struct {
// 	SourceOutboundGateID     int
// 	SourceOutboundGate       gateway
// 	SourceOutboundGateStatus bool

// 	RouterInboundGateID     int
// 	RouterInboundGate       gateway
// 	RouterInboundGateStatus bool
// }

// type RouterToDestination struct {
// 	RouterOutboundGateID     int
// 	RouterOutboundGate       gateway
// 	RouterOutboundGateStatus bool

// 	DestinationInboundGateID     int
// 	DestinationInboundGate       gateway
// 	DestinationInboundGateStatus bool
// }

// // type path struct {
// // 	Status bool
// // 	Hops   []int
// // 	Route  struct{}
// // }

// // RouterOutboundGate ID
// // RouterOutboundGate Client
// // RouterOutboundGate Status

// // DestinnationInboundGate ID
// // DestinnationInboundGate Client
// // DestinnationInboundGate Status
