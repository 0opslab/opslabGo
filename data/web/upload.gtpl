<html>
<head>
    <title>上传文件</title>
</head>
<body>
<form enctype="multipart/form-data" action="/upload" method="post">
  <input type="file" name="file" />
  <input type="hidden" name="token" value="{{.}}"/>
  <input type="submit" value="upload" />
</form>
<form enctype="multipart/form-data" action="/uploadfiles" method="post">
  <input type="file" name="file1" />
  <input type="file" name="file2" />
  <input type="file" name="file3" />
  <input type="hidden" name="token" value="{{.}}"/>
  <input type="submit" value="upload" />
</form>
</body>
</html>