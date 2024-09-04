package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Reservation struct {
	ID              int    `json:"id"`
	UserID          int    `json:"user_id"`
	CourtID         int    `json:"court_id"`
	ReservationDate string `json:"reservation_date"`
	ReservationTime string `json:"reservation_time"`
	Status          string `json:"status"`
	CreatedAt       string `json:"created_at"`
}

type Queue struct {
	ID            int    `json:"id"`
	ReservationID int    `json:"reservation_id"`
	QueueNumber   int    `json:"queue_number"`
	CreatedAt     string `json:"created_at"`
}

var reservations []Reservation
var queues []Queue

func main() {
	router := mux.NewRouter()

	// ตั้งค่าเส้นทาง (routes)
	router.HandleFunc("/reservations", getReservations).Methods("GET")
	router.HandleFunc("/reservations/{id}", getReservation).Methods("GET")
	router.HandleFunc("/reservations", createReservation).Methods("POST")
	router.HandleFunc("/reservations/{id}", updateReservation).Methods("PUT")
	router.HandleFunc("/reservations/{id}", deleteReservation).Methods("DELETE")

	router.HandleFunc("/queues", getQueues).Methods("GET")
	router.HandleFunc("/queues/{id}", getQueue).Methods("GET")
	router.HandleFunc("/queues", createQueue).Methods("POST")
	router.HandleFunc("/queues/{id}", deleteQueue).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}

// ฟังก์ชันสำหรับจัดการ CRUD

func getReservations(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(reservations)
}

func getReservation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, item := range reservations {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Reservation{})
}

func createReservation(w http.ResponseWriter, r *http.Request) {
	var reservation Reservation
	_ = json.NewDecoder(r.Body).Decode(&reservation)
	reservation.ID = len(reservations) + 1
	reservations = append(reservations, reservation)
	json.NewEncoder(w).Encode(reservation)
}

func updateReservation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, item := range reservations {
		if item.ID == id {
			reservations = append(reservations[:index], reservations[index+1:]...)
			var reservation Reservation
			_ = json.NewDecoder(r.Body).Decode(&reservation)
			reservation.ID = id
			reservations = append(reservations, reservation)
			json.NewEncoder(w).Encode(reservation)
			return
		}
	}
	json.NewEncoder(w).Encode(reservations)
}

func deleteReservation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, item := range reservations {
		if item.ID == id {
			reservations = append(reservations[:index], reservations[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(reservations)
}

func getQueues(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(queues)
}

func getQueue(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, item := range queues {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Queue{})
}

func createQueue(w http.ResponseWriter, r *http.Request) {
	var queue Queue
	_ = json.NewDecoder(r.Body).Decode(&queue)
	queue.ID = len(queues) + 1
	queues = append(queues, queue)
	json.NewEncoder(w).Encode(queue)
}

func deleteQueue(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, item := range queues {
		if item.ID == id {
			queues = append(queues[:index], queues[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(queues)
}
