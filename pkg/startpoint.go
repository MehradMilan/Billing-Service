package pkg

const EndpointsAddress = "./resources/endpoints.json"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Start() {
	CollectData(EndpointsAddress)
	<-ServicesChannel
}
