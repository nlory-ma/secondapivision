package pigeon_phoenix

import (
    "net/http"
    "fmt"
    "google.golang.org/api/vision/v1"
    "golang.org/x/oauth2/google"
    "io/ioutil"
/*    "bytes"
//    "github.com/jmcvetta/napping"
//    "net/url"
//    "log"*/
    "golang.org/x/net/context"
//    "strings"
  //  "path"
   /* "encoding/base64"
    "encoding/json"*/
)

func init() {
    http.HandleFunc("/", static)
    http.HandleFunc("/api/appengine", requeteApi)
}

func static(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "client/"+r.URL.Path)
}

// func lecture(img_64 string) {
// 	fmt.Println("Lecture de l'image")
// 	// fmt.Println(img_64)
// }

func requeteApi(w http.ResponseWriter, r *http.Request) {


	switch r.Method {
		case "GET" :
			/*
			data, err := ioutil.ReadFile("~/projet-pigeon.json")
			if err != nil {
			panic(err)
			}

			config, err := google.JWTConfigFromJSON(data,vision.CloudPlatformScope)
			if err != nil {
			panic(err)
			}*/

			/*donne, err := r.ParseForm()
			if err != nil {
				fmt.Print("error")
			}*/
	/*		r.ParseForm()
			image32 := r.PostFormValue("img")
			//fmt.Fprintf(w, "l'image est %s", image32)
			image64 := image32[23:len(image32)]
			fmt.Fprintf(w, "Cette fois l'image est une string et = %s", image64)*/
			/*ctx := context.Background()
		//	client, err := google.DefaultClient(ctx, vision.CloudPlatformScope)
			client := config.Client(ctx)*/
			/*if err != nil {
				fmt.Print("error")
			}*/
			/*service, err :=  vision.New(client)
			if err != nil {
				fmt.Print("error")
			}
			req := &vision.AnnotateImageRequest{
			Image: &vision.Image{
				//content:([]byte)image64,
				Content:base64.StdEncoding.EncodeToString([]byte(image64)),
				},
			Features: []*vision.Feature{{Type: "LABEL_DETECTION"}},
			}
			batch := &vision.BatchAnnotateImagesRequest{
				Requests: []*vision.AnnotateImageRequest{req},
			}
			res, err := service.Images.Annotate(batch).Do()
			if err != nil {
				panic(err)
			}
			body, err := json.MarshalIndent(res.Responses, "", " ")
			if err != nil {
			panic(err)
			}
			fmt.Println(string(body))

			if annotations := res.Responses[0].LabelAnnotations;len(annotations) > 0 {
				label := annotations[0].Description
				fmt.Printf("{ annotations : %s }\n", label)
			}*/
			//fmt.Printf("no annotation found")
		case "POST" :
			data, err := ioutil.ReadFile("~/projet-pigeon.json")
			if err != nil {
			panic(err)
			}

			config, err := google.JWTConfigFromJSON(data,vision.CloudPlatformScope)
			if err != nil {
			panic(err)
			}

		/*	donne, err := r.ParseForm()
			if err != nil {
				fmt.Print("error")
			}*/
			r.ParseForm()
			image32 := r.PostFormValue("img")
			//fmt.Fprintf(w, "l'image est %s", image32)
			image64 := image32[23:len(image32)]
			fmt.Fprintf(w, "Cette fois l'image est une string et = %s", image64)
			ctx := context.Background()
		//	client, err := google.DefaultClient(ctx, vision.CloudPlatformScope)
			client := config.Client(ctx)
			/*if err != nil {
				fmt.Print("error")
			}*/
			service, err :=  vision.New(client)
			if err != nil {
				fmt.Print("error")
			}
			service.Close
			/*
			req := &vision.AnnotateImageRequest{
			Image: &vision.Image{
				//content:([]byte)image64,
				Content:base64.StdEncoding.EncodeToString([]byte(image64)),
				},
			Features: []*vision.Feature{{Type: "LABEL_DETECTION"}},
			}
			batch := &vision.BatchAnnotateImagesRequest{
				Requests: []*vision.AnnotateImageRequest{req},
			}
			res, err := service.Images.Annotate(batch).Do()
			if err != nil {
				panic(err)
			}
			body, err := json.MarshalIndent(res.Responses, "", " ")
			if err != nil {
			panic(err)
			}
			fmt.Println(string(body))

			if annotations := res.Responses[0].LabelAnnotations;len(annotations) > 0 {
				label := annotations[0].Description
				fmt.Printf("{ annotations : %s }\n", label)
			}
			var jsonstr = []byte(body)
			requete, err := http.NewRequest("POST", "https://pigeon-phoenix.appspot.com/api/appengine", bytes.NewBuffer(jsonstr))
			if err != nil {
				panic(err)
			}
			requete.Header.Set("Content-Type", "application/json")*/
	}
}
