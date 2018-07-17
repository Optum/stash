package main

import (
  "log"
	"net/http"
	"io/ioutil"
  "text/template"
  "os"
)

var title = "TITLE"
var subtitle = "SUBTITLE"
var color = "COLOR"

var htmlLandingPage = template.Must(template.New("htmlLandingPage").Parse(`
<!DOCTYPE html>
  <html lang="en">
  <head>
    <link rel="shortcut icon" href="resources/favicon.ico" />
    <title>{{.Title}}</title>
    <style>

      h1, h3, table {
        font-family: arial, sans-serif;
      }

      table {
        border-collapse: collapse;
        width: 100%;
      }

      td, th {
        border: 1px solid #{{.Color}};
        text-align: left;
        padding: 8px;
      }

      tr:nth-child(even) {
        background-color: #{{.Color}};
      }

    </style>
  </head>
  <body>
    <h1>{{.Title}}</h1>
    <h3>{{.SubTitle}}</h3>
    <table>
      <tr>
        <th>
          Download
        </th>
      {{range .Files}}
          <tr>
            <td class="done"><a href="{{.Path}}" download>{{.Name}}</a><br/></td>
          </tr>
      {{end}}
    </table>
    <br/>
    <form enctype="multipart/form-data" action="/resources" method="post">
      <input type="file" name="uploadfile" />
      <input type="submit" value="upload" />
    </form>
  </body>
</html>
`))

type PageData struct {
    Title string
    SubTitle string
    Color string
    Files []File
}

type File struct {
    Path string
    Name string
}

func landingPage(w http.ResponseWriter, req *http.Request) {
  if (logLevel > 1) {
      log.Printf("\n ----- \nPersistent landing page requested \n ----- \n %s\n", req)
  } else {
      log.Printf("Landing page requested")
  }

  var fileData []File

  files, err := ioutil.ReadDir(os.Getenv(sharedDirectory))

  for _, f := range files {
      if(f.Name() != "index.html" && f.Name() != ".snapshot"){
          var next = File{Path: os.Getenv(sharedDirectory) + f.Name(), Name: f.Name()}
          fileData = append(fileData, next)
      }
  }

  pageData := PageData{
      Title: `Stash`,
      SubTitle: `Files`,
      Color: `dddddd`,
      Files: fileData,
  }

  if(os.Getenv(title) != "") {
      pageData.Title = os.Getenv(title)
  }
  if(os.Getenv(subtitle) != "") {
      pageData.SubTitle = os.Getenv(subtitle)
  }
  if(os.Getenv(color) != "") {
      pageData.Color = os.Getenv(color)
  }

  if err != nil {
      log.Fatal(err)
  }

  err = htmlLandingPage.Execute(w, pageData)
  if err != nil {
      log.Fatal(err)
  }

	return
}
