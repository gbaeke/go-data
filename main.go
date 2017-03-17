package main

import (
	"fmt"
	"log"
	"net/http"

	device "github.com/gbaeke/go-device/proto"
	"github.com/gorilla/mux"
	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	"golang.org/x/net/context"
)

// devSvc is the service for the client
var devSvc micro.Service

// devSvcClient is the client
var devSvcClient device.DevSvcClient

func init() {
	//create the client service & client to check if device is active
	devSvc = micro.NewService(micro.Name("device.client"))
	devSvc.Init()
	devSvcClient = device.NewDevSvcClient("DevSvc", devSvc.Client())

}

func deviceActive(d *device.DeviceName) bool {
	//call Get method from devSvcClient to obtain the device
	rsp, err := devSvcClient.Get(context.TODO(), d)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return rsp.Active
}

// DataGet handles /data/deviceid
func DataGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deviceName := vars["device"]
	fmt.Fprintln(w, "Device active: ", deviceActive(&device.DeviceName{Name: deviceName}))
	fmt.Fprintln(w, "Oh and, no data for you!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/data/{device}", DataGet)

	log.Fatal(http.ListenAndServe(":8080", router))

}
