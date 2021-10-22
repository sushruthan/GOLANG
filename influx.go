package main

import (
  "context"
  "fmt"
  "github.com/influxdata/influxdb-client-go/v2"
  "time"
)

func main() {
  // You can generate a Token from the "Tokens Tab" in the UI
  const token = "tdgZrkfYzZ_jllhbRxOD2js3aV4zwU23Dr4QULMoBxTln6Sbt0oS0RVnI4eXnd2HkXmkx8XPdNOnT8jrtnITFA=="
  const bucket = "node"
  const org = "demo"

  client := influxdb2.NewClient("http://localhost:8086", token)

  // get non-blocking write client
writeAPI := client.WriteAPI(org, bucket)
var id_1 string = "1"
var id_2 string = "2"
var id_3 string = "3"
var source string = "1"
var target string = "2"
// write line protocol
writeAPI.WriteRecord(fmt.Sprintf("node, id=%s",id_1 ))
writeAPI.WriteRecord(fmt.Sprintf("node, id=%s",id_2 ))
writeAPI.WriteRecord(fmt.Sprintf("edge,id=3 id=%s,source=%s,target=%s", id_3, source, target))
// Flush writes
writeAPI.Flush()
  // always close client at the end
  defer client.Close()
}