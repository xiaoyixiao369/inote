function vlidField(field){
    return field.replace(/[@#\$%\^&\*<>']+/g, '');
}