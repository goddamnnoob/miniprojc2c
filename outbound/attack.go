package outbound

type Attack struct {
	host             string
	port             int16
	attackType       string
	packetbatchcount int64
}
