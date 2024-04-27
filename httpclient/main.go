package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

/*
Demo of create a http client.

- GET request
 url : https://v2.jokeapi.dev/joke/Any?type=twopart
*/

type Joke struct {
	Category string `json:"category"`
	Type     string `json:"type"`
	Setup    string `json:"setup"`
	Delivery string `json:"delivery"`
	Error    bool   `json:"error"`
<<<<<<< HEAD
	Safe     bool   `json:"safe"`
	Flags    struct {
		Explicit bool `json:"explicit"`
	} `json:"flags"`
=======
	Flags    Flags  `json:"flags"`
	Safe     bool   `json:"safe"`
}

type Flags struct {
	Nsfw      bool `json:"nsfw"`
	Religious bool `json:"religious"`
	Political bool `json:"political"`
	Racist    bool `json:"racist"`
	Sexist    bool `json:"sexist"`
	Explicit  bool `json:"explicit"`
>>>>>>> 79ebf7c1fe497e981e5249bf3c9fc1b72688a860
}

type Anime struct {
	Anime     string `json:"anime"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
}

type PostWebhookRequest struct {
	Setup    string `json:"setup"`
	Delivery string `json:"delivery"`
}

type PostWebhookResponse struct {
	Success bool `json:"success"`
}

var jokeURL = "https://v2.jokeapi.dev/joke/Any?type=twopart"
var animeURL = "https://animechan.xyz/api/random"
var webHookURL = "https://webhook.site/cecb0cd8-1209-4083-9bb8-b604752f2929"

func GetJoke() (*Joke, error) {
	// make a GET request to the joke API
	resp, err := http.Get(jokeURL)
	if err != nil {
		fmt.Println("Error while making a GET request to the joke API :", err)
		return nil, err
	}
	defer resp.Body.Close()

	// decode the response from the joke API to a Joke struct
	var joke Joke
	err = json.NewDecoder(resp.Body).Decode(&joke)
	if err != nil {
		fmt.Println("Error while decoding the response from the joke API :", err)
		return nil, err
	}

	return &joke, nil
}

func GetAnime() (*Anime, error) {
	resp, err := http.Get(animeURL)
	if err != nil {
		fmt.Println("Error while making a GET request to the anime API :", err)
		return nil, err
	}

	var anime Anime
	err = json.NewDecoder(resp.Body).Decode(&anime)
	if err != nil {
		fmt.Println("Error while decoding the response from the anime API :", err)
		return nil, err
	}

	return &anime, nil
}

func GetAnimeTellJoke(joke *Joke, anime *Anime) (string, error) {

	if !joke.Safe {
		// check flags true

		flags := []string{}
		if joke.Flags.Nsfw {
			flags = append(flags, "nsfw")
		}
		if joke.Flags.Religious {
			flags = append(flags, "religious")
		}
		if joke.Flags.Political {
			flags = append(flags, "political")
		}
		if joke.Flags.Racist {
			flags = append(flags, "racist")
		}
		if joke.Flags.Sexist {
			flags = append(flags, "sexist")
		}

		strFlags := ""
		for i, flag := range flags {
			if i == 0 {
				strFlags += flag
			} else {
				strFlags += ", " + flag
			}
		}

		str := fmt.Sprintf("This joke is not safe for work, because it contains %v \n", strFlags)
		return str, nil

	} else {
		str := fmt.Sprintf("%v said joke \"%v\", \"%v\" \n", anime.Character, joke.Setup, joke.Delivery)
		return str, nil
	}
}

func PostJokeToWebhook(joke *Joke) (*PostWebhookResponse, error) {

	// create a PostWebhookRequest as the body of the POST request
	req := PostWebhookRequest{
		Setup:    joke.Setup,
		Delivery: joke.Delivery,
	}

	// encode the PostWebhookRequest to JSON
	reqBody, err := json.Marshal(req)
	if err != nil {
		fmt.Println("Error while encoding the PostWebhookRequest to JSON :", err)
		return nil, err
	}

	// make a POST request to the webhook URL
	resp, err := http.Post(webHookURL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("Error while making a POST request to the webhook URL :", err)
		return nil, err
	}
	defer resp.Body.Close()

	// decode the response from the webhook to a PostWebhookResponse struct
	var postWebhookResponse PostWebhookResponse
	err = json.NewDecoder(resp.Body).Decode(&postWebhookResponse)
	if err != nil {
		fmt.Println("Error while decoding the response from the webhook :", err)
		return nil, err
	}

	return &postWebhookResponse, nil
}

func main() {

	// get a joke from the joke API
	joke, err := GetJoke()
	if err != nil {
		fmt.Println("Error while getting a joke from the joke API :", err)
		return
	}

<<<<<<< HEAD
	// print the joke
	if joke.Safe {
		fmt.Printf("Category: %s\n", joke.Category)
		fmt.Printf("Type: %s\n", joke.Type)
		fmt.Printf("Setup: %s\n", joke.Setup)
		fmt.Printf("Delivery: %s\n", joke.Delivery)
	} else {
		fmt.Printf("This joke is not safe for work, because it contains %t\n", joke.Flags.Explicit)
	}

=======
>>>>>>> 79ebf7c1fe497e981e5249bf3c9fc1b72688a860
	// get an anime quote from the anime API
	anime, err := GetAnime()
	if err != nil {
		fmt.Println("Error while getting an anime quote from the anime API :", err)
		return
	}

<<<<<<< HEAD
	// print the anime quote
	fmt.Printf("%s said \"%s\" in %s \n", anime.Character, anime.Quote, anime.Anime)
=======
	/*
			- Format of the output: <anime> said "<setup>", <delivery>
			- add flags and safe in joke struct
		    - check if joke is safe, then print the joke, check if joke is not safe, then print "This joke is not safe for work, because it contains <flags yang true>"

				Example output:
				Fuyou Kaede said joke "I'm not saying my son is ugly...", "But on Halloween he went to tell the neighbors to turn down their TV and they gave him some candy."
	*/

	output, err := GetAnimeTellJoke(joke, anime)
	if err != nil {
		fmt.Println("Error while getting anime tell joke :", err)
		return
	}

	// post the joke to the webhook
	_, err = PostJokeToWebhook(joke)
	if err != nil {
		fmt.Println("Error while posting the joke to the webhook :", err)
		return
	}

	fmt.Println(output)

>>>>>>> 79ebf7c1fe497e981e5249bf3c9fc1b72688a860
}
