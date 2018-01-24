package main

import (
	"bytes"
	"testing"
)

func Test_pp(t *testing.T) {
	in := `<!doctype html>

<html lang="en">
<head>
  <meta charset="utf-8">
  <title>The HTML5 Herald</title>
  <link rel="stylesheet" href="css/styles.css?v=1.0">
</head>
<body>
  <main>
<!-- this is commented node -->
<!--
multi line
-->
  <div id="content">
    test test
    test test
  </div>
  </main>
</body>
</html>`
	expected := `  <html lang="en" >
    <head >
      <meta charset="utf-8" />
      <title >
        The HTML5 Herald
      </title>
      <link rel="stylesheet" href="css/styles.css?v=1.0" />
    </head>
    <body >
      <main >
        <!--  this is commented node -->
        <!-
        multi line
        -->
        <div id="content" >
          test test
          test test
        </div>
      </main>
    </body>
  </html>
`

	act := new(bytes.Buffer)
	Print(StringReader(in), act)
	if bytes.Compare([]byte(expected), act.Bytes()) != 0 {
		t.Fatal("not matched")
	}
}
