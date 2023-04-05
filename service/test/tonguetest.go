package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
)

type Client struct {
	client UploadClient
}

func NewClient(conn grpc.ClientConnInterface) Client {
	return Client{
		client: NewUploadClient(conn),
	}
}

//	func UploadTest(file *multipart.FileHeader) (*Response, error) {
//		fileContent, err := file.Open()
//		if err != nil {
//			return nil, err
//		}
//		//image := bytes.NewBuffer(nil)
//		//fmt.Println(image)
//		//if _, err := io.Copy(image, fileContent); err != nil {
//		//	return nil, err
//		//}
//		// Initialise gRPC connection.
//		conn, err := grpc.Dial(viper.GetString("rpc.host"), grpc.WithInsecure())
//		//fmt.Println(viper.GetString("rpc.host"))
//		if err != nil {
//			log.Fatalln(err)
//		}
//		defer conn.Close()
//
//		client := NewClient(conn)
//		ctx := context.Background()
//		ctx, cancel := context.WithDeadline(ctx, time.Now().Add(10*time.Second))
//		defer cancel()
//
//		stream, err := client.client.SendImage(ctx)
//		if err != nil {
//			return nil, err
//		}
//		buf := make([]byte, 1024)
//		for {
//			num, err := fileContent.Read(buf)
//			if err == io.EOF {
//				break
//			}
//			if err != nil {
//				return nil, err
//			}
//			if err := stream.Send(&Data{Image: buf[:num]}); err != nil {
//				return nil, err
//			}
//		}
//
//		resp, err := stream.CloseAndRecv()
//		if err != nil {
//			return nil, err
//		}
//		return resp, nil
//	}
func UploadTest(file *multipart.FileHeader) (map[string]interface{}, error) {
	testUrl := viper.GetString("test.url")
	fmt.Println(file)
	fileContent, err := file.Open()
	if err != nil {
		return nil, err
	}

	fmt.Println(fileContent)
	defer fileContent.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", file.Filename)
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(part, fileContent); err != nil {
		return nil, err
	}
	//fmt.Println("--------------------")
	fmt.Println(part)
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100
	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: t,
	}
	req, err := http.NewRequest("POST", testUrl, body)
	if err != nil {

		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Length", "application/x-www-form-urlencoded")

	var result map[string]interface{}
	req.Close = true
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body1, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("--------------------", err)
		return nil, err
	}
	if err := json.Unmarshal(body1, &result); err != nil {
		return nil, err
	}
	defer client.CloseIdleConnections()
	return result, nil
}
