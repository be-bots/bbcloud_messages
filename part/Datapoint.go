package part

// Datapoint gives you sensor data
type Datapoint struct {
	Xpos            int32
	Temp            float32
	Hum             float32
	Press           float32
	TopPhotoURI     string
	TopPhotoRedURI  string
	TopPhotoIRURI   string
	SidePhotoURI    string
	SidePhotoRedURI string
	SidePhotoIRURI  string
}
