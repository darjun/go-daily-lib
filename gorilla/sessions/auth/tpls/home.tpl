<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Home</title>
</head>
<body>
{{ if . }}
<p>Hi, {{ .Username }}</p><br>
<a href="/secret">Goto secret?</a>
{{ else }}
<p>Hi, stranger</p><br>
<a href="/login">Goto login?</a>
{{ end }}
</body>
</html>