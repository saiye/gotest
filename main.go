package main

//var server engine.FactoryServer
/*func main() {
/*	if len(server.ServerArr) == 0 {
server = engine.FactoryServer{
		ServerArr: []engine.Server{
			engine.HttpServer{},
			engine.WebSocketServer{},
		},
	}
	server.Start()
}*/

/*
	res,err:=MyDiv1(10,0)
	fmt.Println("----res:",res)
	fmt.Println("----err:",err)*/

}*/
func MyDiv1(a int32,b int32) (float64,error)  {
	defer func() {
		if p := recover(); p != nil {
			//
		}
	}()
	return MyDiv(a,b),nil
}

func MyDiv(a int32,b int32) float64{
	if b == 0 {
	    panic("b不能是0")
	}
	return float64(a)/float64(b)
}
