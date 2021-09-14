package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"fmt"
    "io/ioutil"
    "strings"
    "bufio"
    "os"
)


type API int



/*Read the vegData file and return the content*/
func (a *API) GetVegData(empty string, reply *[]string) error{

	f, _ := os.OpenFile("VegData.data", os.O_WRONLY|os.O_CREATE, 0666)
	f.Sync()

	data,error := ioutil.ReadFile("VegData.data")
	if error != nil {
        fmt.Println(error)}
	lines := strings.Split(string(data), "\n")
	*reply = lines

	return nil
}

/*Add new vegetable with amount and price*/
func (a *API) AddNewVeg(item string, reply *[] string) error{
	data := []byte(item)

	// OpenFile with more options. Last param is the permission mode
   // Second param is the attributes when opening
	f, err := os.OpenFile("VegData.data", os.O_APPEND|os.O_WRONLY, 0600)
	err =f.Sync()
	if err != nil {
	    panic(err)
	}
	defer f.Close()
	fmt.Fprintf(f, "%s", data)

	return nil
}

/* Search a vegetable entry by name and display its details*/
func (a *API) GetByName(name string, reply *[]string) error {
		f, _ := os.Open("VegData.data")
		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
		    if strings.Contains(scanner.Text(), name) {
		       rdlines := strings.Split(string(scanner.Text()), "\n")
		       //lines := scanner.Text()
		       *reply = rdlines
		    }
		}
	//	f.Close()
	f.Sync()
	return nil
}

/*edit a vegetable's current amount and price*/
func (a *API) EditVegData(item string, reply *[]string) error {


		f, err := os.OpenFile("VegData.data", os.O_WRONLY|os.O_CREATE, 0666)
		f.Sync()

      
        input, err := ioutil.ReadFile("VegData.data")
        if err != nil {
                log.Fatalln(err)
        }

        ln := strings.Split(string(input), "\n")

        for i, line := range ln {
                if strings.Contains(line, item) {
                		split:= strings.Split(line, " ")
                   		*reply = split
                        ln[i] = " "
                }

        }

        output := strings.Join(ln, "\n")
        err = ioutil.WriteFile("VegData.data", []byte(output), 0644)
        if err != nil {
                log.Fatalln(err)
        }
        f.Sync()
	return nil
}

  

func main() {

	api := new(API)

	//Publish the receivers methods
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering API", err)
	}

	//Register HTTP Handler
	rpc.HandleHTTP()

	//Listen to TCP connections on port 4040
	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Fatal("Listener error", err)
	}
	log.Printf("serving rpc on port %d", 4040)
	
	//Start accepting incoming HTTP connections
	http.Serve(listener, nil)

	if err != nil {
		log.Fatal("error serving: ", err)
	}


}