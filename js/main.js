/**
 * Created by igor on 01.03.17.
 */
/**
 * Created by igor on 01.03.17.
 */
function loginValidation() {
    var login = document.getElementsByName("userName")[0].value;
    var incorrectChars = ["%", "'", '"', "$", "~", "#"];

    var validationStatus = false;
    if (login.length >= 3) {
        validationStatus = true;

        for (var i = 0; i < login.length && validationStatus; i++) {
            for (var j = 0; j < incorrectChars.length; j++) {



                if (login.indexOf(incorrectChars[j]) != -1) {
                    validationStatus = false;
                    break;

                }
            }
        }
    }

    console.log(validationStatus)

}