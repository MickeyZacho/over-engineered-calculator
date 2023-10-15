package api

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Username     string
	PasswordHash []byte
	Token        string
	Id           int
}

type LoginRequest struct {
	Username string
	Password string
}

type LoginResponse struct {
	Username string
	Token    string
}

var testUser User
var allUsers []User

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	testUser = User{
		Username: "test",
		// sha256 of "123" + salt
		PasswordHash: []byte{105, 5, 69, 71, 83, 56, 247, 133, 180, 60, 155, 128, 37, 247, 154, 163, 87, 5, 12, 172, 176, 0, 200, 151, 223, 132, 70, 125, 175, 202, 72, 60},
		Id:           4,
	}
	allUsers = append(allUsers, testUser)
	bodybytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Could not read request body")
	}
	var logreq LoginRequest
	json.Unmarshal(bodybytes, &logreq)

	passHash := myHash(logreq.Password)

	for _, u := range allUsers {
		if u.Username != logreq.Username {
			break
		}
		if !bytes.Equal(passHash, u.PasswordHash) {
			break
		}
		var (
			key []byte
			t   *jwt.Token
			s   string
		)

		key = []byte("mysecretkey") /* Load key from somewhere, for example an environment variable */
		t = jwt.New(jwt.SigningMethodHS256)
		s, err = t.SignedString(key)
		if err != nil {
			io.WriteString(w, "error creating token: "+err.Error())
			return
		}

		//Make an effort to store the token "s"
		response := LoginResponse{
			Username: u.Username,
			Token:    s,
		}

		marshalledResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		io.WriteString(w, string(marshalledResponse))
	}
}

func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	bodybytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Could not read request body")
	}
	var createReq LoginRequest
	json.Unmarshal(bodybytes, &createReq)

	for _, u := range allUsers {
		if u.Username == createReq.Username {
			io.WriteString(w, "Username not available")
			return
		}
	}

	newUser := User{
		Username:     createReq.Username,
		PasswordHash: myHash(createReq.Password),
		Id:           len(allUsers),
	}

	allUsers = append(allUsers, newUser)

	io.WriteString(w, "user created - try logging in now")

}

func myHash(pass string) []byte {
	mySecretSalt := "very-secret-salt-that-should-live-in-.env"
	h := sha256.New()
	h.Write([]byte(pass + mySecretSalt))
	passHash := h.Sum(nil)
	return passHash
}
