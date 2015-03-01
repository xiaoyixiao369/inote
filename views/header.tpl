<!DOCTYPE html>
<html>
<head lang="en">
  <meta charset="UTF-8">
  <title>iNote</title>
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport"
        content="width=device-width, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no">
  <meta name="format-detection" content="telephone=no">
  <meta name="renderer" content="webkit">
  <meta http-equiv="Cache-Control" content="no-siteapp"/>
  <link rel="alternate icon" type="image/png" href="/static/img/favicon.png">
  <link rel="stylesheet" href="/static/meizi/css/amazeui.min.css"/>
  <link rel="stylesheet" href="/static/css/inote.css"/>
</head>
<body>
<header class="am-topbar am-topbar-inverse am-topbar-fixed-top">
  <h1 class="am-topbar-brand" data-am-scrollspy="{animation:'slide-top', delay: 200}">
    <a href="/"><span class="am-badge am-badge-secondary am-radius brand-i">i</span><span class="am-text-middle brand-note">Note</span></a>
  </h1>

  <button class="am-topbar-btn am-topbar-toggle am-btn am-btn-sm am-btn-success am-show-sm-only"
          data-am-collapse="{target: '#doc-topbar-collapse'}"><span class="am-sr-only">导航切换</span> <span
      class="am-icon-bars"></span></button>

  <div class="am-collapse am-topbar-collapse" id="doc-topbar-collapse">
    <ul class="am-nav am-nav-pills am-topbar-nav">
      <li class="am-active am-animation-fade"><a href="/"><span class="am-icon-home am-icon-md"></span></a></li>
      <li class="inote-words am-vertical-align-middle"><small id="siteWords"></small></li>
    </ul>
  </div>
</header>