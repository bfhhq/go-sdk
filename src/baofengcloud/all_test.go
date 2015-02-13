package baofengcloud

import (
	//"fmt"
	"testing"
)

var conf = Configure{
	AccessKey: "",
	SecretKey: "",
}

func TestQuery(t *testing.T) {

	fileInfo, err := QueryFile(&conf, Paas, "test.mp4", "")
	if err != nil {
		t.Fatalf("%s", err)
	}

	if fileInfo.FileName != "test.mp4" {
		t.Fatalf("%v", fileInfo)
	}

}

func TestUpdate(t *testing.T) {

	/*
		fileInfo, err := UpdateFile(&conf, Paas, "test.mp4", "", Private)
		if err != nil {
			t.Fatalf("%s", err)
		}

		if fileInfo.FileName != "test.mp4" {
			t.Fatalf("%v", fileInfo)
		}*/

}

func TestDelete(t *testing.T) {

	result, err := DeleteFile(&conf, Paas, "test.mp4", "", "")
	if err != nil {
		t.Fatalf("%s", err)
	}

	if result.Status != 0 {
		t.Fatalf("Status: %d", result.Status)
	}

}

func TestUpload(t *testing.T) {

	err := UploadFile2(&conf, Paas, Public, "C:\\test.mp4", "test.mp4", "", "")
	if err != nil {
		t.Fatalf("%s", err)
	}

}
