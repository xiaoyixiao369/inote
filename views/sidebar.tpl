{{define "sidebar"}}
<div class="am-g am-g-fixed inote-g-fixed">
  <div class="am-u-md-3">
    <div class="am-panel-group">
      <section class="am-panel am-panel-default">
        <div class="am-panel-hd">关于我</div>
        <div class="am-panel-bd">
          <p><img class="am-circle" src="/static/img/1.jpg" width="140" height="140"/></p>
          <p>iNote正在开发中.....</p>
          <a class="am-btn am-btn-success am-btn-sm" href="#">查看更多 →</a>
        </div>
      </section>
      <section class="am-panel am-panel-default">
              <div class="am-panel-hd">分类</div>
              <ul class="am-list inote-list">
                {{range .Categories}}
                   <li><a href="#">{{.Name}}</a></li>
                {{end}}
              </ul>
            </section>
    </div>
  </div>
{{end}}