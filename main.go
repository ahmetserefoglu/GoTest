package main

/*import (
	"ApiGateway/app"
	"ApiGateway/controllers"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication) // Middleware'e JWT kimlik doğrulaması eklenir
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")

router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	port := os.Getenv("PORT") // Environment dosyasından port bilgisi getirilir
	if port == "" {
		port = "8000" //localhost:8000
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) // Uygulamamız localhost:8000/api altında istekleri dinlemeye başlar
	if err != nil {
		fmt.Print(err)
	}
}*/

import (
	"ApiGateway/app"
	"ApiGateway/controllers"
	u "ApiGateway/utils"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Product struct {
	id int `json:"id"`
	Name  string `json:"Name"`
	Price  string `json:"Price"`
	CreatedAt string `json:"CreatedAt"`
	UpdatedAt string `json:"UpdatedAt"`
}

func returnAllProducts(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp, err := http.Get("http://127.0.0.1:8200/api/product")
	if err != nil {
        fmt.Println("No response from request")
    }
    defer resp.Body.Close()
	//var jsonData []Product
    body, err := ioutil.ReadAll(resp.Body) // response body is []byte
    fmt.Println(string(body))   
	response := u.Message(true, "Veriler Cekildi!")
	response["status"] = http.StatusOK
	w.Header().Add("Content-Type", "application/json")
	response["products"] = string(body)
	u.Respond(w, response)
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/product", returnAllProducts).Methods("GET")

	router.Use(app.JwtAuthentication) // Middleware'e JWT kimlik doğrulaması eklenir
	port := os.Getenv("PORT") // Environment dosyasından port bilgisi getirilir
	if port == "" {
		port = "8000" //localhost:8000
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) // Uygulamamız localhost:8000/api altında istekleri dinlemeye başlar
	if err != nil {
		fmt.Print(err)
	}
	
	/*url := "http://127.0.0.1:8000/api/product"
	var bearer = "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6ImFkY2Q4ZTg2Y2Y3MDBmNzI1NjE0ODQyN2RlZWUwOWE4MzNiMThlMTEwYjViMDUyODMyOWViNTM0YjcwOGE4ZTEyYzQyNmI1MmI3ZjYxZjIzIn0.eyJhdWQiOiI4IiwianRpIjoiYWRjZDhlODZjZjcwMGY3MjU2MTQ4NDI3ZGVlZTA5YTgzM2IxOGUxMTBiNWIwNTI4MzI5ZWI1MzRiNzA4YThlMTJjNDI2YjUyYjdmNjFmMjMiLCJpYXQiOjE2Mzc2OTk4NTcsIm5iZiI6MTYzNzY5OTg1NywiZXhwIjoxNjY5MjM1ODU3LCJzdWIiOiIxIiwic2NvcGVzIjpbXX0.B3qTONGatTS43j_cLogYRkruNcDcJ8vdWpqdU86IERdXLq__aqqxrDWLMtUHpKrxCAPyECqWaWR82rBNfdDU3ds7jb8LauKCKcjIUq1kLwfMUlMvsAxIQY8PX_URAF7KntD_XohGc8e6YuAidqSDDjlWTd7rGijl0BEU8pUHr4nNiYHRzxq2ivTLmTeUc_ZTsOFeCG3NSMsIdBApG0bfLFunaLofzIrc7oNPrTZOzHl-7bAQnrnxcnnICJfZWgCq00G8lbD2eTKkINymYz4uTLX8oA0-P-7g3qi-AiCcOrfdghkkt-vCFInvyPrNwquHRi7XrBF3bz3qBSEC6H9ftRDgsXQE5toKHiKW-E9VSbz_t4tMFvrrH46JEduSPU_mqolG8fJCRrVIdN2KPif044RFk-Yiulvhz5bHYV0nTunOnInJxWWVIjSt79TRggwLrtYiKvtXTa33QlgtFHnYCtasahku7oWoF1kVZ7UyOPRXjEFbGmKSA2Gg54tq_BayIthZ_IMg7p_7s6t7YkeoT6m-4VM7GobiPzwYCtNrjWlkciOnSUIO60osXZ0eHwLoMwwxxxZXsz-Bbb1Y2h_MvckIl0AuAdHzAXodwcG4S0dJs-DYBnDvIE12AnKYcwrEBOeKXyPXFmbYReha6Ol8MWG4W3a7b-ruLXLECGqQXck" 
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", bearer)
	// add authorization header to the req
    req.Header.Add("Authorization", bearer)

    // Send req using http Client
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Println("Error on response.\n[ERROR] -", err)
    }
    defer resp.Body.Close()

	english := func(name string) string { return "Hello, " + name }

        say(english, "ANisus")
        sayTest(english)

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println("Error while reading the response bytes:", err)
    }
    log.Println(string([]byte(body)))*/
}