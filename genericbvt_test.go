package httptesting

import (
	http "net/http"
	"testing"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	s "strings"
)

//TestRequest :: Main Function for all requests
func TestSyncRequests(t *testing.T) {
	host := "http://localhost:3131"
	
	var links=make(map[string]string)
	
	dat, err := ioutil.ReadFile("bvtlinks.list")
	
	if(err==nil){
		for _,entry := range s.Split(string(dat),"\n"){
			var temp=s.Split(entry,";")
			links[temp[0]]=s.TrimSpace(temp[1])
		}
	}else{
		t.Fatal("Error reading file:",err)
	}
	
	for name,url := range links{
	t.Run(name, func(t *testing.T) {
		t.Parallel()
		resp, err := httpRequestResponseValues(host + url)

		if err != nil {
			assert.FailNow(t, "Unexpected Error",err)
		}

		// verify http response status
		if resp.StatusCode != http.StatusOK {
			assert.FailNow(t, "Status Code was not 200")
		}
	})
	}
}
