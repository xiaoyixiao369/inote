$(function(){
    $('.inote-copyright-year').text(getCopyrightYear());
    $.get('/author', function(user){
        var siteTitle = user.siteTitle;
        if(user.headBgPic && user.headBgPic.length > 0){
            $('header').css({'background': 'url('+user.headBgPic+') center top no-repeat'});
        }
        $('#siteTitle').text(vlidField(siteTitle));
        $('title').text(vlidField(siteTitle));
        $('#footerCopty').text(vlidField(user.siteTitle));
        $('#avatar').attr('src', user.thumb);
        $('#author').text(vlidField(user.userName));
        $('#aboutMe').text(vlidField(user.aboutMe));
    });
    hljs.initHighlightingOnLoad();

    var hash = window.location.hash;
    if(hash && hash.indexOf('#') == 0){
        var id = hash.substring(1, hash.length);
        if ($.isNumeric(id)){
            fetchOnePost(id, false);
        } else {
            window.location = '/';
        }
    } else {
        fetchOnePost(0, true);
    }


    function fetchOnePost(id, firstLoding){
        if(!firstLoding){
            window.location = '/#postLoding';
        }

        $('#postLoding').plainOverlay('show');
        setTimeout(function () {
            $.get('/i/posts/'+id, function(res){
                if(res.success){
                    var post = res.data.post;
                    $('#postId').val(post.id);
                    $('#postTitle').text(vlidField(post.title));
                    $('#postTag').text(vlidField(post.tag));
                    $('#postPublishAt').text($.dateFormat(post.publishAt));
                    $('#postContent').html(marked(post.content));
                    $('pre code', '#postContent').each(function(i, block) {
                        hljs.highlightBlock(block);
                    });
                    $('img', '#postContent').each(function(){
                        $(this).addClass('img-rounded').wrap('<a class="venobox hvr-float" data-gall="myGallery" href="'+$(this).attr('src')+'">')
                    });
                    $('.venobox').venobox();
                    if(!firstLoding){
                        window.location = '/#'+id;
                    }
                    $('#postLoding').plainOverlay('hide');
                    var messages = res.data.messages;
                    if(messages) {
                        var msgCount = messages.length;
                        $('#postMessageCount').text(msgCount);
                        var msgs = [];
                        for(var i = 0; i < msgCount; i++){
                            var message = '<blockquote><div class="hvr-sink"><span class="text-primary">'+vlidField(messages[i].content)+'</span><footer><span><em>'+ $.dateFormat(messages[i].createdAt)+'</em>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;'+vlidField(messages[i].guestName)+'</span></footer></div><br/>';
                            if(messages[i].reply){
                                message +=  '<footer class="hvr-bounce-to-right">回复:&nbsp;&nbsp;'+messages[i].reply+'</footer>';
                            }
                            message += '</blockquote>';
                            msgs.push(message);
                        }
                        $('#postMessages').html(msgs.join(''));
                    }

                } else {
                    $('#postLoding').plainOverlay('hide');
                    alert(res.msg);
                }

            });
        }, 300)
    }

    function fetchPosts(page){
        $.get('/i/posts/list/'+page, function(posts){
            if(posts && posts.length > 0){
                var lis = [];
                for(var i = 0; i < posts.length; i++){
                    lis.push('<li class="list-group-item"><input type="hidden" name="" value="'+posts[i].id+'"/><a class="inote-post-list">['+posts[i].tag+']&nbsp;&nbsp;'+posts[i].title+'<small class="pull-right"><em>'+ $.dateFormat(posts[i].publishAt)+'</em></small></a></li>');
                }
                $('#postList').append(lis.join(''));
                $('.inote-post-list').on('click', function(){
                    fetchOnePost($(this).prev().val(), false);
                });
            } else {
                $('#loadPageTip').text('已全部加载完毕');
            }
        });
    }

    var $pagIndex = $('#pageIndex');
    $('#loadNextPageBtn').on('click', function(){
        var nextPage = parseInt($pagIndex.val()) + 1;
        $pagIndex.val(nextPage);
        fetchPosts(nextPage);
    });

    fetchPosts(parseInt($pagIndex.val()));

    $("#messageForm input,textarea").jqBootstrapValidation({
        preventSubmit: true,
        submitError: function($form, event, errors) {

        },
        submitSuccess: function($form, event) {
            event.preventDefault();
            var guestName = vlidField($("#guestName").val());
            var messageContent = vlidField($("#messageContent").val());
            if($.trim(guestName) == '' || $.trim(messageContent) == ''){
                $('#success').html("<div class='alert alert-danger'>").find('.alert-danger').html("<button type='button' class='close' data-dismiss='alert' aria-hidden='true'>&times;")
                    .append("</button>").append("<strong>名字和内容都不能为空</strong>");
                $('#messageForm').trigger("reset");
                return false;
            }

            var message =  {
                postId: $('#postId').val(),
                postTitle: $('#postTitle').text(),
                guestName: guestName,
                content: messageContent
            };
            $.ajax({
                url: '/i/submitMsg',
                contentType: 'applcation/json',
                type: 'POST',
                data: JSON.stringify(message),
                cache: false,
                success: function(res) {
                    if(res.success){
                        $('#postMessages').prepend('<blockquote><div class="hvr-sink"><span class="text-primary">'+message.content+'</span><footer><span><em>'+ $.dateFormat(new Date())+'</em>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;'+message.guestName+'</span></footer></div></blockquote>');
                        var $postMsgCount = $('#postMessageCount');
                        $postMsgCount.text(parseInt($postMsgCount.text())+1);
                        window.location = '/#postMessageAnchor';
                        window.location = '/#'+ message.postId;
                    }else{
                        $('#success').html("<div class='alert alert-danger'>").find('.alert-danger').html("<button type='button' class='close' data-dismiss='alert' aria-hidden='true'>&times;")
                            .append("</button>").append("<strong>提交失败</strong>");
                    }
                    $('#messageForm').trigger("reset");
                },
                error: function() {
                    $('#messageForm').trigger("reset");
                }
            })
        },
        filter: function() {
            return $(this).is(":visible");
        }
    });

    $("a[data-toggle=\"tab\"]").click(function(e) {
        e.preventDefault();
        $(this).tab("show");
    });

    $("[data-toggle='tooltip']").tooltip();
});