<!DOCTYPE html>

<html>
  	<head>
    	<title>Beego</title>
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	
		<style type="text/css">
			body {
				margin: 0px;
				font-family: "Helvetica Neue",Helvetica,Arial,sans-serif;
				font-size: 14px;
				line-height: 20px;
				color: rgb(51, 51, 51);
				background-color: rgb(255, 255, 255);
			}

			p {
				margin: 0px 0px 10px;
			}
		</style>
	</head>
  	
  	<body>
  		<header >
  		{{.username}}  {{.password}}
		</header>
		<div>
		<form method="post" action="/user/post">
		     UserName:<input type="text" name="username" /><br/>
		     Passowrd:<input type="password" name="password"/><br/>
		     <input type='submit' value='NEW save' />
		</form>
		<form method="post" action="/user/login">
		     UserName:<input type="text" name="username" /><br/>
		     Passowrd:<input type="password" name="password"/><br/>
		     <input type='submit' value='Login' />
		</form>
		</div>
	</body>
</html>
