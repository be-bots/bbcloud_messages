package part

// Source tell you the origin of the message
type Source struct {
	Farm  string
	Cell  int32
	Layer int32
	Slot  int32
	Rfid  string
}
