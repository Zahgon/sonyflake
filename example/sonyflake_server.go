package main

import (
	"net/http"

	"github.com/sony/sonyflake"
	"github.com/sony/sonyflake/awsutil"
)

var sf *sonyflake.Sonyflake

func init() {
	var st sonyflake.Settings
	st.MachineID = awsutil.AmazonEC2MachineID
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
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
