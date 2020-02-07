package file

import (
	"io"
	"os"
	"net/http"
	"os/exec"
)

func Download(filename, url string, headers map[string]string) (int64, error){
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	for k, v:= range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	output, err := os.Create(filename)
	if err != nil {
		return 0, err
	}
	defer output.Close()

	size, err := io.Copy(output, resp.Body)
	if err != nil {
		return 0, err
	}

	return size, nil
}


func Wget(filename, url string) (int64, error) {
	//p := os.Getenv("PATH")
	wget, err := exec.LookPath("wget")
	if err != nil {
		return 0, err
	}
	_, err = exec.Command(wget, url, "-O", filename).Output()
	if err != nil {
		return 0, err
	}

	file, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}
	size := file.Size()
	return size, nil
}

