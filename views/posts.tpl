<div class="am-u-md-8">
    <div class="inote-notice">{{.flash.notice}}</div>
    {{range .Posts}}
    <article class="inote-main">
        <h3 class="am-article-title inote-title">
            <a href="/i/posts/{{.Id}}">{{.Title}}</a>
        </h3>
        <h4 class="am-article-meta inote-meta"><span class="am-icon-calendar"></span> {{.PublishAt}}</h4>
    </article>

    <hr class="am-article-divider">

    {{end}}

    <ul class="am-pagination">
        <li class="am-pagination-prev"><a href="">&laquo; 上一页</a></li>
        <li class="am-pagination-next"><a href="">下一页 &raquo;</a></li>
    </ul>
</div>
