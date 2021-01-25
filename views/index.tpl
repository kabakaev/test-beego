<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">

  <style type="text/css">
    *,body {
      margin: 0px;
      padding: 0px;
    }

    body {
      margin: 0px;
      font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
      font-size: 14px;
      line-height: 20px;
      background-color: #fff;
    }

    header,
    footer {
      width: 100%;
      margin-left: auto;
      margin-right: auto;
    }

    header {
      padding: 0 0;
    }

    content {
      width: 100%;
      margin: 0;
      position: absolute;
      top: 50%;
      -ms-transform: translateY(-50%);
      transform: translateY(-50%);
    }

    footer {
      line-height: 1.8;
      text-align: center;
      padding: 50px 0;
      color: #999;
    }

    .description {
      text-align: center;
      font-size: 16px;
    }

    a {
      color: #444;
    }

    .emoji {
      text-decoration: none;
      font-size: 32px;
    }

    .backdrop {
      position: absolute;
      width: 100%;
      height: 100%;
      box-shadow: inset 0px 0px 100px #ddd;
      z-index: -1;
      top: 0px;
      left: 0px;
    }

    .hello {
      text-align: center;
      font-size: {{.HeaderFontSize}}px;
      padding: 0 0 70px;
      font-weight: normal;
      text-shadow: 0px 1px 2px #ddd;
    }
  </style>
</head>

<body>
  <content>
  <header>
    <h1 class="hello">Hello {{.Name}}!</h1>

    <div class="description">
      <a class="emoji" href=".">	&#8635; </a>
      <a class="emoji" href="?action=good&h1size={{.HeaderFontSize}}"> &#128077; </a>
      <a class="emoji" href="?action=bad&h1size={{.HeaderFontSize}}"> &#128078; </a>
      <a class="emoji" href="/stat">&#128200;</a>
    </div>

    </header>
    <footer>
      <div class="links">
        <a href="about">Powered by Beego</a> /
        <a href="/very-secret-stat.html">Webserver statistics</a> /
        <a href="https://github.com/kabakaev/test-beego">Source code</a>
      </div>
    </footer>
  </content>
  <div class="backdrop"></div>
</body>
</html>
