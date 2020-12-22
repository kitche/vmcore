package routers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"panel/config"
	"strconv"
	"strings"
	"net/url"
)

//Server will have two methods, to add a new server, and to get all servers
	type Server interface {
		CreateServer(server *Server) error
		GetServer() ([]*Server, error)
}

	// `dbServer` struct implements the `Store` interface. Variable `db` takes the pointer
	// to the SQL database connection object.
	type dbServer struct {
		db *sql.DB
	}

	//Create global `server` variable
	var server Server

func (server *dbServer) CreateServer(server *Server)  error {
	// 'Server' is a struct which has "hostname", "ip", and "version" attributes.
	// Type SQL query to insert new server into our database.
	// Note: `serverinfo` is the name of the table within our `serverDatabase` postgresql database.
	_, err := server.db.Query(
		"INSERT INTO serverinfo(hostname, ip, version) VALUES ($1,$2,$3)"
		server.Hostname, server.ip, server.Version)
	return err
}

	func (server *dbServer) GetServer() ([]*Server, error) {
		// Query the database for all persons, and return the result to the `rows` object.
		// Note: `peopleinfo` is the name of the table within our `peopleDatabase`
		rows, err := server.db.Query("SELECT hostname, ip, version FROM serverinfo")
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		// Create an empty slice of pointers to `Person` struct. This slice will be returned
		// by this function to its caller.
		serverList := []*Server{}
		for rows.Next() {
			// For each row returned from the database, create a pointer to a `Person` struct.
			server := &Server{}
			// Populate the `Hostname`, `IP`, and `Version` attributes of the person
			if err := rows.Scan(&server.Hostname, &server.IP, &server.Version); err != nil {
				return nil, err
			}
			// Finally, append the new person to the returned slice, and repeat for the next row
			serverList = append(serverList, server)
		}
		return serverList, nil
}

func ServerHandler(w http.ResponseWriter, r *http.Request) {
	//	c, err := r.Cookie("session_token")
	//	if err != nil {
	//		if err == http.ErrNoCookie {
	// If the cookie is not set, return an unauthorized status
	//			w.WriteHeader(http.StatusUnauthorized)
	//			return
	//		}

	serverList, err := server.GetServer()

	serverListBytes, err := json.Marshal(serverList)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write JSON list of servers to response
	w.Write(serverListBytes)
}
	if err := config.Servertpl.Execute(w, nil); err != nil {
		//if err:= templates["index"].Execute(w, "hello"); err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func createServerHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the HTML form data received in the request
	err := r.ParseForm()
	if err != nil {
	fmt.Println(fmt.Errorf("Error: %v", err))
	w.WriteHeader(http.StatusInternalServerError)
	return
}

// Extract the field information about the person from the form info
	server := Server{}
	server.Hostname := r.Form.Get("hostname")
	server.IP := r.Form.Get("ip")
	server.version := r.Form.Get("version")
	server.Host := r.Form.Get("host")
// Write new person details into postgresql database using our `store` interface variable's
// `func (*dbstore) CreatePerson` pointer receiver method defined in `store.go` file
	err = server.CreatePerson(&server)
	if err != nil {
	fmt.Println(err)
}

	endpoint := server.Host
	data := url.Values{}
	data.Set("Hostname", server.Hostname)
	data.Set("IP", server.IP)
	data.Set("Version", server.version)
	client := &http.Client{}
	r, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	res, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.Status)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	//Encode the data
	//postBody, _ := json.Marshal(map[string]string{
	//	"hostname":  server.Hostname,
	//	"ip": server.IP,
	//})
	//responseBody := bytes.NewBuffer(postBody)
	//	resp, err = http.Post(server.Host"/api" server.Hostname, server.IP, server.version)
//Redirect to the originating HTML page
	http.Redirect(w, r, "/", http.StatusFound)
}