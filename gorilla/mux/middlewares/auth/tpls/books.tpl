<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Books</title>
</head>
<body>
<ol>
{{ range . }}
<li>
<p>书名: <a href="/books/{{ .ISBN }}">{{ .Name }}</a></p>
<p>作者: {{ .Authors }}</p>
<p>出版社: {{ .Press }}</p>
<p>出版日期: {{ .PublishedAt }}</p>
</li>
{{ end }}
</ol>
</body>
</html>