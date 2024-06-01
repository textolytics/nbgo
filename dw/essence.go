package core

type Event struct{}

type InboundEvent struct {
	// SourceTimestamp      timestamppb.Timestamp
	Source      *Gateway
	Destination *Gateway
	// DestinationTimestamp timestamppb.Timestamp
	Context Event
}

type OutboundEvent struct {
	Source      *Gateway
	Destination *Gateway
	Context     Event
}
