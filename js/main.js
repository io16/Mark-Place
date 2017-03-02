/**
 * Created by igor on 01.03.17.
 */
function validation(obj) {
    var objName = obj.getAttribute('name')

    var objValue = obj.value;
    var incorrectChars = ["%", "'", '"', "$", "~", "#"];

    var validationStatus = false;
    if (objValue.length >= 3) {
        validationStatus = true;

        for (var i = 0; i < objValue.length && validationStatus; i++) {
            for (var j = 0; j < incorrectChars.length; j++) {


                if (objValue.indexOf(incorrectChars[j]) != -1) {
                    validationStatus = false;
                    break;

                }
            }
        }
    }

    var objSpan = document.getElementById(objName + "Span");
    if (validationStatus) {

        objSpan.className = "fontawesome-check";
    }
    else
        objSpan.className = "fa fa-times";
}

function passCorrect(obj) {
    var objSpan = document.getElementById(obj.getAttribute('name') + "Span2");
    if (obj.value === document.getElementsByName("pass")[0].value) {
        objSpan.className = "fontawesome-check";
    }
    else
        objSpan.className = "fa fa-times";
}
function emailValidation(obj) {
    var re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
   console.log(re.test(obj.value))
}
/**
 * Created by igor on 01.03.17.
 */
