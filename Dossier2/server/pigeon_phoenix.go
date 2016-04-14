package pigeon_phoenix

import (
	"net/http"
	"fmt"
	//"google.golang.org/api/vision/v1"
	//"golang.org/x/oauth2/google"
	"bytes"
//	"google.golang.org/appengine"
	//"google.golang.org/appengine/urlfetch"
	//"golang.org/x/net/context"
	"encoding/json"
	"encoding/base64"
//	"flag"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
)

type imageContent struct {
	Content string `json:"content"`
}
type feature struct {
	TypeF      string `json:"type"`
	MaxResults string `json:"maxResults"`
}
type request struct {
	Image    imageContent `json:"image"`
	Features []feature    `json:"features"`
}

type vision struct {
	Requests []request `json:"requests"`
}

type vertice struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type boundingPoly struct {
	Vertices []vertice `json:"vertices"`
}

type fdBoundingPoly struct {
	Vertices []vertice `json:"vertices"`
}

type position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type landmark struct {
	Type     string   `json:"type"`
	Position position `json:"position"`
}

type faceAnnotation struct {
	BoundingPoly           boundingPoly   `json:"boundingPoly"`
	FdBoundingPoly         fdBoundingPoly `json:"fdBoundingPoly"`
	Landmarks              []landmark     `json:"landmarks"`
	RollAngle              float64        `json:"rollAngle"`
	PanAngle               float64        `json:"panAngle"`
	TiltAngle              float64        `json:"tiltAngle"`
	DetectionConfidence    float64        `json:"detectionConfidence"`
	LandmarkingConfidence  float64        `json:"landmarkingConfidence"`
	JoyLikelihood          string         `json:"joyLikelihood"`
	SorrowLikelihood       string         `json:"sorrowLikelihood"`
	AngerLikelihood        string         `json:"angerLikelihood"`
	SurpriseLikelihood     string         `json:"surpriseLikelihood"`
	UnderExposedLikelihood string         `json:"underExposedLikelihood"`
	BlurredLikelihood      string         `json:"blurredLikelihood"`
	HeadwearLikelihood     string         `json:"headwearLikelihood"`
}
type result struct {
	Responses []response `json:"responses"`
}

type response struct {
	FaceAnnotations []faceAnnotation `json:"faceAnnotations"`
}

func encode(bin []byte) []byte {
	e64 := base64.StdEncoding

	maxEncLen := e64.EncodedLen(len(bin))
	encBuf := make([]byte, maxEncLen)

	e64.Encode(encBuf, bin)
	return encBuf
}

func fromLocal(fname string) (string, error) {
	var b bytes.Buffer

	fileExists, _ := exists(fname)
	if !fileExists {
		return "", fmt.Errorf("File does not exist\n")
	}

	file, err := os.Open(fname)
	if err != nil {
		return "", fmt.Errorf("Error opening file\n")
	}

	_, err = b.ReadFrom(file)
	if err != nil {
		return "", fmt.Errorf("Error reading file to buffer\n")
	}

	enc := encode(b.Bytes())
	res := string(enc[:])

	return res, nil
}

func get(url string) ([]byte, string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error getting url.")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	ct := resp.Header.Get("Content-Type")

	if resp.StatusCode == 200 && len(body) > 512 {
		return body, ct
	}

	return []byte(""), ct
}

func fromRemote(url string) string {
	image, _ := get(url)
	enc := encode(image)
	res := string(enc[:])
	return res
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func post(jsonStr []byte, url string) result {
	fmt.Println("URL:>", url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	res := result{}
	json.Unmarshal(body, &res)
	return res
}


/*
func demo42(b64 string, key string) {
	imageV := imageContent{Content: b64}
	featureV := feature{TypeF: "FACE_DETECTION", MaxResults: "10"}
	requestV := request{Image: imageV, Features: []feature{featureV}}
	visionV := vision{Requests: []request{requestV}}

	visionResultJSON, _ := json.Marshal(visionV)
	res := post(visionResultJSON, "https://vision.googleapis.com/v1/images:annotate?key="+key)
	fmt.Println(res)
}*/

/////////////////////////////////*****************************\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\



func init() {
	http.HandleFunc("/", static)
	http.HandleFunc("/api/appengine", requeteApi)
}

func static(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "client/"+r.URL.Path)
}

func demo42(w http.ResponseWriter, r *http.Request) {

	key := "598f6547f1d255211434519b034f8394aa6801e2"
	r.ParseForm()
	image32 := r.PostFormValue("img")
	//ctx := appengine.NewContext(r)
	//client := urlfetch.Client(ctx)
	imageV := imageContent{Content: image32}
	featureV := feature{TypeF: "FACE_DETECTION", MaxResults: "10"}
	requestV := request{Image: imageV, Features: []feature{featureV}}
	visionV := vision{Requests: []request{requestV}}

	visionResultJSON, _ := json.Marshal(visionV)
	res := post(visionResultJSON, "https://vision.googleapis.com/v1/images:annotate?key="+key)
	fmt.Println(res)

}

func requeteApi(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	switch r.Method {
		case "GET":
			get(r.URL.String())
		case "POST":
			json = demo42()
			post(json, r.URL.String)
	}
}
