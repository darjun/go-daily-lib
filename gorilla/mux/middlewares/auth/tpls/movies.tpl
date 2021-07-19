<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Movies</title>
</head>
<body>
<ol>
  {{ range . }}
  <li>
    <p>书名: <a href="/movies/{{ .IMDB }}">{{ .Name }}</a></p>
    <p>上映日期: {{ .PublishedAt }}</p>
    <p>时长: {{ .Duration }}分</p>
    <p>语言: {{ .Lang }}</p>
  </li>
  {{ end }}
</ol>
</body>
</html>