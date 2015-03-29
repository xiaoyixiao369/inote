$(function(){
    $.get('/author', function(user){
        $('#siteTitle').text(user.siteTitle);
        $('#footerCopty').text(user.siteTitle);
        $('#avatar').attr('src', user.thumb);
        $('#author').text(user.userName);
        $('#aboutMe').text(user.aboutMe);
    });
});