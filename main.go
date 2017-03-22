package main

import (
	"fmt"
	"log"
	"net/http"

	device "github.com/gbaeke/go-device/proto"
	"github.com/gorilla/mux"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	"golang.org/x/net/context"
)

// devSvc is the service for the client
var (
	cl device.DevSvcClient
)

func init() {
	// make sure flags are processed
	cmd.Init()

	// initialise a default client for device service
	cl = device.NewDevSvcClient("go.micro.srv.device", client.DefaultClient)

}

func deviceActive(d *device.DeviceName) bool {
	//call Get method from devSvcClient to obtain the device
	fmt.Println("Getting device", d.Name)
	rsp, err := cl.Get(context.TODO(), d)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return rsp.Active
}

// DataGet handles /data/deviceid
func DataGet(w http.ResponseWriter, r *http.Request) {
	// retrieve variables in request
	vars := mux.Vars(r)

	// get device name
	deviceName := vars["device"]

	// print result; no data is actually retrieved
	fmt.Fprintln(w, "Device active: ", deviceActive(&device.DeviceName{Name: deviceName}))
	fmt.Fprintln(w, "Oh and, no data for you!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	// handler for /data/{device}
	router.HandleFunc("/data/{device}", DataGet)

	log.Fatal(http.ListenAndServe(":8080", router))

}
