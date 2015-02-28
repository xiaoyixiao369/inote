<div class="am-g am-g-fixed blog-g-fixed">
    {{with .Post}}
    <div class="am-u-md-8 am-u-md-offset-2">
        <br/>
        <article class="am-article">
            <h3 class="am-article-title am-text-center">
                <a href="">{{.Title}}</a>
            </h3>
            <h4 class="am-article-meta"><span class="am-icon-calendar"></span>&nbsp;{{.PublishAt}}&nbsp;&nbsp;&nbsp;&nbsp;<span class="am-icon-book"></span> {{.Category.Name}}</h4>

            <div class="am-article-bd">
                {{str2html .Content}}
            </div>
        </article>

        <hr class="am-article-divider">

    </div>
    {{end}}
</div>