package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

type Project struct {
	Id int
	Name string
	Project_id int
	Description string
	Pull_count int
	Star_count int
	Tags_count int
	Creation_time string
	Update_time string
}

type Config struct {
	Lables interface{}
}

type Image struct {
	Digest string `json:"digest"`
	Name string `json:"name"`
	Size int `json:"size"`
	Architecture string `json:"architecture"`
	Os string `json:"os"`
	DockerVersion string `json:"docker_version"`
	Author string `json:"author"`
	Created string `json:"created"`
	Config Config `json:"config"`
	Signature string `json:"signature"`
}

func imageCleanerHandler(w http.ResponseWriter, r *http.Request) {
	var username string = os.Getenv("USER")
	var passwd string = os.Getenv("PASS")

	var projects []Project

	tr := &http.Transport {
            TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
        }
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", os.Getenv("HOST_ADDR") + "/api/repositories?project_id=10&sort=creation_time", nil)
	req.SetBasicAuth(username, passwd)
	resp, err := client.Do(req)
	if err != nil{
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(bodyText), &projects)
	for i := range projects {
		fmt.Println("--------------")
		fmt.Println(projects[i].Name)
		req, err = http.NewRequest("GET", os.Getenv("HOST_ADDR") + "/api/repositories/" + projects[i].Name + "/tags?sort=creation_time", nil)
                req.SetBasicAuth(username, passwd)
                resp, err = client.Do(req)
                if err != nil{
                        log.Fatal(err)
                }
	        bodyText, err = ioutil.ReadAll(resp.Body)
		var images []Image
		json.Unmarshal([]byte(bodyText), &images)
	        fmt.Println(len(projects))
		for img := range images {
			sort.Slice(images, func(i, j int) bool {
				return images[i].Created < images[j].Created
			})
			fmt.Println(len(images))
			if img < (len(images) - 5) {
				fmt.Printf("%v\n", os.Getenv("HOST_ADDR") + "/api/repositories/" + projects[i].Name + "/tags/" + images[img].Name)
			}

		}
		fmt.Println("---------------")
	}
}
