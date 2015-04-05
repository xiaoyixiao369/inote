function vlidField(field){
    return field.replace(/[@#\$%\^&\*<>']+/g, '');
}

$.extend({
    dateFormat: function(time){
        var data = new Date(time);
        var year = data.getFullYear();
        var month = data.getMonth() + 1;
        var day = data.getDate();
        var hours = data.getHours();
        var minutes = data.getMinutes();
        if(hours < 10) {
            hours = '0' + hours;
        }
        if(minutes < 10) {
            minutes = '0' + minutes
        }
        return  year + '年' + month + '月' + day + '日' + ' ' + hours + ':' + minutes;
    }
});