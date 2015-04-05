$(function(){
    $.get('/author', function(user){
        var siteTitle = user.siteTitle;
        if(user.headBgPic && user.headBgPic.length > 0){
            $('header').css({'background': 'url('+user.headBgPic+') center top no-repeat'});
        }
        $('#siteTitle').text(siteTitle);
        $('title').text(siteTitle);
        $('#footerCopty').text(user.siteTitle);
        $('#avatar').attr('src', user.thumb);
        $('#author').text(user.userName);
        $('#aboutMe').text(user.aboutMe);
    });

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
                    $('#postTitle').text(post.title);
                    $('#postTag').text(post.tag);
                    $('#postPublishAt').text($.dateFormat(post.publishAt));
                    $('#postContent').html(post.content);
                    if(!firstLoding){
                        window.location = '/#i';
                    }
                    $('#postLoding').plainOverlay('hide');
                    var messages = res.data.messages;
                    if(messages) {
                        var msgCount = messages.length;
                        $('#postMessageCount').text(msgCount);
                        var msgs = [];
                        for(var i = 0; i < msgCount; i++){
                            msgs.push('<blockquote><div class="hvr-sink"><span class="text-primary">'+messages[i].content+'</span><footer><span><em>'+ $.dateFormat(messages[i].createdAt)+'</em>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;'+messages[i].guestName+'</span></footer></div></blockquote>');
                        }
                        $('#postMessages').html(msgs.join(''));
                    }

                } else {
                    $('#postLoding').plainOverlay('hide');
                    alert(res.msg);
                }

            });
        }, 500)
    }
    fetchOnePost(0, true);

    function fetchPosts(page){
        $.get('/i/posts/list/'+page, function(posts){
            if(posts && posts.length > 0){
                var lis = [];
                for(var i = 0; i < posts.length; i++){
                    lis.push('<li class="list-group-item"><input type="hidden" name="" value="'+posts[i].id+'"/><a class="inote-post-list">['+posts[i].tag+']&nbsp;&nbsp;'+posts[i].title+'</a></li>');
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
            var guestName = $("#guestName").val();
            var messageContent = $("#messageContent").val();
            if($.trim(guestName) == '' || $.trim(messageContent) == ''){
                $('#success').html("<div class='alert alert-danger'>").find('.alert-danger').html("<button type='button' class='close' data-dismiss='alert' aria-hidden='true'>&times;")
                    .append("</button>").append("<strong>名字和内容都不能为空</strong>");
                $('#messageForm').trigger("reset");
                return false;
            }

            var message =  {
                postId: $('#postId').val(),
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
                        window.location = '/#i';
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