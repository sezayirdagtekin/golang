package main

import "fmt"

func main() {

	var builder = newNotificationBuilder()
	builder.setTitle("welcome")
	builder.SetSubTitle("alert image")
	builder.SetIcon("wlcm.jpg")
	builder.SetImage("hello.jpg")
	builder.SetPriority(3)
	builder.SetMessage("welcome to  go lang course")
	builder.SetType("msg")
	notification, err := builder.Build()

	if err != nil {
		fmt.Println("Error occurred when creation notification", err)
	}

	fmt.Printf(" Notificaton :%+v", notification)
}
