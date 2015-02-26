{{define "footer"}}

<footer class="inote-footer">
  <p><small>powered by <a href="http://beego.me" target="_blank">beego</a> and <a href="http://amazeui.org" target="_blank">amaze ui</a></small><br/>
    <small>Released under the <a href="http://opensource.org/licenses/MIT" target="_blank">MIT</a> license</small>
  </p>
</footer>

<!--[if lt IE 9]>
<script src="/static/js/jquery-1.11.1.min.js"></script>
<script src="http://cdn.staticfile.org/modernizr/2.8.3/modernizr.js"></script>
<script src="/static/meizi/js/polyfill/rem.min.js"></script>
<script src="/static/meizi/js/polyfill/respond.min.js"></script>
<script src="/static/meizi/js/amazeui.legacy.js"></script>
<![endif]-->

<!--[if (gte IE 9)|!(IE)]><!-->
<script src="/static/meizi/js/jquery.min.js"></script>
<script src="/static/meizi/js/amazeui.min.js"></script>
<!--<![endif]-->

</body>
</html>
{{end}}