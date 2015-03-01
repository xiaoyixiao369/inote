<div class="am-g am-g-fixed inote-g-fixed">
  <div class="am-u-md-3">
    <div class="am-panel-group">
      <section class="am-panel am-panel-default">
        <div class="am-panel-hd">关于我</div>
        <div class="am-panel-bd">
          <p class="inote-center"><img id="avatar" class="am-circle" src="" width="100" height="100"/></p>
          <p class="inote-center" id="userName"></p>
          <small id="aboutMe"></small>
        </div>
      </section>
      <section class="am-panel am-panel-default">
              <div class="am-panel-hd">分类</div>
              <ul class="am-list inote-list">
                {{range .Categories}}
                   <li><a href="/i/catetory/{{.Id}}">{{.Name}}</a></li>
                {{end}}
              </ul>
            </section>
    </div>
</div>