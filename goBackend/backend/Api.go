package main

func api() {
	r := createRoutes()
	r.Static("/photos", "./photos")

	r.Run(":3030")
	//if err := endless.ListenAndServe("localhost:3030", r); err != nil {
	//	log.Fatal(err)
	//}
}
