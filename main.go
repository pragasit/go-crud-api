package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Structs สำหรับข้อมูลการจองและคิว
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

// จำลองข้อมูลการจองและคิว
var reservations []Reservation
var queues []Queue

func main() {
	// สร้าง router ใหม่
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

	// เริ่มต้น server บนพอร์ต 8000
	http.ListenAndServe(":8000", router)
}

// ฟังก์ชันสำหรับจัดการ CRUD

// Get all reservations
func getReservations(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(reservations)
}

// Get a single reservation by ID
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

// Create a new reservation
func createReservation(w http.ResponseWriter, r *http.Request) {
	var reservation Reservation
	_ = json.NewDecoder(r.Body).Decode(&reservation)
	reservation.ID = len(reservations) + 1
	reservations = append(reservations, reservation)
	json.NewEncoder(w).Encode(reservation)
}

// Update an existing reservation by ID
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

// Delete a reservation by ID
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

// Get all queues
func getQueues(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(queues)
}

// Get a single queue by ID
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

// Create a new queue
func createQueue(w http.ResponseWriter, r *http.Request) {
	var queue Queue
	_ = json.NewDecoder(r.Body).Decode(&queue)
	queue.ID = len(queues) + 1
	queues = append(queues, queue)
	json.NewEncoder(w).Encode(queue)
}

// Delete a queue by ID
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
