<html>
<head>
<title></title>
</head>
<body>
	<form action="/login" method="post">
		<input type="checkbox" name="interest" value="football">足球
		<input type="checkbox" name="interest" value="basketball">篮球
		<input type="checkbox" name="interest" value="tennis">網球
	    帳號:<input type="text" name="username">
	    密碼:<input type="password" name="password">
	    <input type="submit" value="登入">
	    <input type="hidden" name="token" value="{{.}}">
	</form>
</body>
</html>