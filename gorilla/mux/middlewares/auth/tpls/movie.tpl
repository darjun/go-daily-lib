<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Movies</title>
</head>
<body>
<p>IMDB: {{ .IMDB }}</p>
<p>电影名: {{ .Name }}</p>
<p>上映日期: {{ .PublishedAt }}</p>
<p>时长: {{ .Duration }}分</p>
<p>语言: {{ .Lang }}</p>
</body>
</html>