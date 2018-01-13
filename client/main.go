package main
import (
    "net"
    "fmt"
    "bufio"
    "os"
	  "log"
    "bytes"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "strings"
)


// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP
}

func getMacAddr() (addr string) {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				// Don't use random as we have a real address
				addr = i.HardwareAddr.String()
				break
			}
		}
	}
	return
}


func getInfo() (string, string) {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Generating KeyValue")
  KeyValue := getMacAddr()
  fmt.Print("Enter propertyTag: ")
  PropertyTag, _ := reader.ReadString('\n')

  proTag := strings.TrimRight(PropertyTag, "\r\n")
  return KeyValue, proTag
}

func makeHttpPostReq(ipAddress, keyValue, propertyTag string) {
    fmt.Println("Starting to send the POST Request")
    jsonData := map[string]string{"IPaddress": ipAddress, "KeyValue" : keyValue, "PropertyTag" : propertyTag}
    jsonValue, _ := json.Marshal(jsonData)
    request, _ := http.NewRequest("POST", "http://ekstrah.com:8000/serverlist", bytes.NewBuffer(jsonValue))
    request.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    response, err := client.Do(request)

    if err != nil {
      fmt.Println("The HTTP request failed with error %s\n", err)
    } else {
      data, _ := ioutil.ReadAll(response.Body)
      fmt.Println(string(data))
    }
}

func main() {
  IPAddress := GetOutboundIP()
  IPAddressString := IPAddress.String()
  KeyValue, PropertyTag := getInfo()
  makeHttpPostReq(IPAddressString, KeyValue, PropertyTag)

  fmt.Println(PropertyTag)
  // if len(strings.TrimSpace(KeyValue)) == 0 {
}
