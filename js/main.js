/**
 * Created by igor on 01.03.17.
 */
/**
 * Created by igor on 01.03.17.
 */
function loginValidation() {
    var login = document.getElementsByName("userName")[0];
    var loginValue = login.value;
    var incorrectChars = ["%", "'", '"', "$", "~", "#"];

    var validationStatus = false;
    if (loginValue.length >= 3) {
        validationStatus = true;

        for (var i = 0; i < loginValue.length && validationStatus; i++) {
            for (var j = 0; j < incorrectChars.length; j++) {


                if (loginValue.indexOf(incorrectChars[j]) != -1) {
                    validationStatus = false;
                    break;

                }
            }
        }
    }

    if (validationStatus)
        login.style.backgroundColor = "green";
    else

        login.style.backgroundColor = "red";
}