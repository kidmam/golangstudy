package main

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type UserList struct {
	XMLName xml.Name
	Body    struct {
		XMLName           xml.Name
		ListUsersResponse struct {
			XMLName xml.Name
			Return  []string `xml:"return"`
		} `xml:"listUsersResponse"`
	}
}

func main() {
	// wsdl service url
	url := fmt.Sprintf("%s%s%s",
		"https://12.34.56.78:9443",
		"/services/RemoteUserStoreManagerService",
		".RemoteUserStoreManagerServiceHttpsSoap11Endpoint",
	)

	// payload
	payload := []byte(strings.TrimSpace(`
		<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.ws.um.carbon.wso2.org">
			<soapenv:Body>
				<ser:listUsers>
					<ser:filter></ser:filter>
					<ser:maxItemLimit>100</ser:maxItemLimit>
				</ser:listUsers>
			</soapenv:Body>
		</soapenv:Envelope>`,
	))

	httpMethod := "POST"

	// soap action
	soapAction := "urn:listUsers"

	// authorization credentials
	username := "admin"
	password := "admin"

	log.Println("-> Preparing the request")

	// prepare the request
	req, err := http.NewRequest(httpMethod, url, bytes.NewReader(payload))
	if err != nil {
		log.Fatal("Error on creating request object. ", err.Error())
		return
	}

	// set the content type header, as well as the oter required headers
	req.Header.Set("Content-type", "text/xml")
	req.Header.Set("SOAPAction", soapAction)
	req.Header.Set("Authorization", fmt.Sprintf(
		"Basic %s",
		base64.StdEncoding.EncodeToString([]byte(
			username+":"+password,
		)),
	))

	// prepare the client request
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	log.Println("-> Dispatching the request")

	// dispatch the request
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Error on dispatching request. ", err.Error())
		return
	}

	log.Println("-> Retrieving and parsing the response")

	// read and parse the response body
	result := new(UserList)
	err = xml.NewDecoder(res.Body).Decode(result)
	if err != nil {
		log.Fatal("Error on unmarshaling xml. ", err.Error())
		return
	}

	log.Println("-> Everything is good, printing users data")

	// print the users data
	users := result.Body.ListUsersResponse.Return
	fmt.Println(strings.Join(users, ", "))
}
