package devices

// DeviceFuncs devices supported
var DeviceFuncs = map[string]func(string, map[string]interface{}) (Data, error){
	"pms5003": pms5003,
}

type particles struct {
	PM1  int `json:"pm1"`
	PM25 int `json:"pm2.5"`
	PM10 int `json:"pm10"`
}

// Data data structure for data returned from devices
type Data struct {
	CF            particles `json:"cf"`
	Atmospheric   particles `json:"atmos"`
	ParticleCount struct {
		PointThree   int `json:"0.3um"`
		PointFive    int `json:"0.5um"`
		One          int `json:"1um"`
		TwoPointFive int `json:"2.5um"`
		Five         int `json:"5um"`
		Ten          int `json:"10um"`
	} `json:"particle_count"`
	ConcUnit string `json:"concentration_unit"`
}
