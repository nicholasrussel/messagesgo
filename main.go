package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	controllers "github.com/messagesgo/controllers"
)

func main() {
	router := mux.NewRouter()

	// Use handleSend for the /send route
	router.HandleFunc("/send", controllers.HandleSend).Methods("POST")
	// Add more routes as needed...

	svrPort := controllers.LoadEnv("SVR_PORT")
	log.Println("Connected to port " + svrPort)
	addr := ":" + svrPort
	http.ListenAndServe(addr, router)
}


























// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"github.com/go-resty/resty/v2"
// )

// const (
//     apiEndpoint = "https://api.openai.com/v1/chat/completions"
// )

// func main() {

// 	// Use your API KEY here
//     apiKey := "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"
//     client := resty.New()

// 	response, err := client.R().
//         SetAuthToken(apiKey).
//         SetHeader("Content-Type", "application/json").
//         SetBody(map[string]interface{}{
//             "model":      "gpt-3.5-turbo",
//             "messages":   []interface{}{map[string]interface{}{"role": "system", "content": "Hi can you tell me what is the factorial of 10?"}},
//             "max_tokens": 50,
//         }).
//         Post(apiEndpoint)

//     if err != nil {
//         log.Fatalf("Error while sending send the request: %v", err)
//     }

// 	body := response.Body()

//     var data map[string]interface{}
//     err = json.Unmarshal(body, &data)
//     if err != nil {
//         fmt.Println("Error while decoding JSON response:", err)
//         return
//     }

//     // Extract the content from the JSON response
//     content := data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
// 	fmt.Println(content)

// 	// log.Println("RUNNING ...")
// 	// controllers.GocronEvent()
// 	// router := mux.NewRouter()

// 	// router.HandleFunc("/login", controllers.Login).Methods("POST")
// 	// router.HandleFunc("/logout", controllers.Logout).Methods("GET")
// 	// router.HandleFunc("/user", controllers.AddNewUser).Methods("POST")
// 	// router.HandleFunc("/user", controllers.Authenticate(controllers.UpdateUser, 1)).Methods("PUT")
// 	// router.HandleFunc("/user/hotel", controllers.GetHotelList).Methods("GET")
// 	// router.HandleFunc("/user/hotel/room", controllers.GetRoomList).Methods("GET")
// 	// router.HandleFunc("/user/hotel/room", controllers.Authenticate(controllers.AddNewOrder, 1)).Methods("POST")
// 	// router.HandleFunc("/user/flight", controllers.GetFlightList).Methods("GET")
// 	// router.HandleFunc("/user/flight/seat", controllers.GetSeatList).Methods("GET")
// 	// router.HandleFunc("/user/flight", controllers.Authenticate(controllers.AddNewOrder, 1)).Methods("POST")
// 	// router.HandleFunc("/user/bus", controllers.GetBusTripList).Methods("GET")
// 	// router.HandleFunc("/user/bus", controllers.Authenticate(controllers.AddNewOrder, 1)).Methods("POST")
// 	// router.HandleFunc("/user/train", controllers.GetTrainTripList).Methods("GET")
// 	// router.HandleFunc("/user/train", controllers.Authenticate(controllers.AddNewOrder, 1)).Methods("POST")
// 	// router.HandleFunc("/user/tour", controllers.GetTourList).Methods("GET")
// 	// router.HandleFunc("/user/tourschedule", controllers.GetTourScheduleList).Methods("GET")
// 	// router.HandleFunc("/user/tour", controllers.Authenticate(controllers.AddNewOrder, 1)).Methods("POST")
// 	// router.HandleFunc("/user/order", controllers.Authenticate(controllers.GetUserOrder, 1)).Methods("GET")
// 	// router.HandleFunc("/user/order", controllers.Authenticate(controllers.RequestRefund, 1)).Methods("PUT")

// 	// router.HandleFunc("/partner", controllers.AddNewPartner).Methods("POST")
// 	// router.HandleFunc("/partner", controllers.Authenticate(controllers.UpdatePartner, 2)).Methods("PUT")
// 	// router.HandleFunc("/partner/flight", controllers.Authenticate(controllers.GetFlightPartnerList, 2)).Methods("GET")
// 	// router.HandleFunc("/partner/flight", controllers.Authenticate(controllers.AddNewFlight, 2)).Methods("POST")
// 	// router.HandleFunc("/partner/flight", controllers.Authenticate(controllers.DeleteFlight, 2)).Methods("DELETE")

// 	// router.HandleFunc("/partner/airline", controllers.Authenticate(controllers.AddNewAirline, 2)).Methods("POST")
// 	// router.HandleFunc("/partner/airplane", controllers.Authenticate(controllers.AddNewAirplane, 2)).Methods("POST")

// 	// router.HandleFunc("/admin/refund", controllers.Authenticate(controllers.GetRefundList, 0)).Methods("GET")
// 	// router.HandleFunc("/admin/refund/{orderId}", controllers.Authenticate(controllers.ApproveRefund, 0)).Methods("DELETE")

// 	// svrPort := controllers.LoadEnv("SVR_PORT")
// 	// log.Println("Connected to port " + svrPort)
// 	// addr := ":" + svrPort
// 	// http.ListenAndServe(addr, router)





// 	// error := http.ListenAndServe(addr, router)
// 	// if error != nil {
// 	// 	log.Print(error)
// 	// }

// }
