package main

import "log"

func main() {

	// Creating the Builder
	bld := newBuilder("Title", "Sub")
	bld.SetImage("image.jpg").SetIcon("icon.png").SetPriority(8)
	bld.SetMessage("Basic notif").SetType("type")

	// Building Notification by using Builder
	notif, err := bld.Build()
	if err != nil {
		log.Fatalln("Error creation Notification", err)
	}

	log.Printf("Notification: %+v\n", *notif)
}
