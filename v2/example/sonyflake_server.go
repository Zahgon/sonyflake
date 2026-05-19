package main

import (
	"net/http"

	"github.com/sony/sonyflake/v2"
	"github.com/sony/sonyflake/v2/awsutil"
)

var sf *sonyflake.Sonyflake

func init() {
	var st sonyflake.Settings
	st.MachineID = awsutil.AmazonEC2MachineID

	var err error
	sf, err = sonyflake.New(st)
	if err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
