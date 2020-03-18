class Validator {
  static validEmail(str) {
    return /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/.test(str);
  }

  // Checks validity of password
  // "str" takes a string to compare
  // "strength" asks for required strength to be considered valid
  static validPassword(password, strength) {
    // Check length first
    if (password.length < 8) {
      return;
    }
    // Additional criteria
    var matchedCase = new Array();
    matchedCase.push("[$@$!%*#?&]"); // Special Character
    matchedCase.push("[A-Z]"); // Uppercase letter
    matchedCase.push("[0-9]"); // Numbers
    matchedCase.push("[a-z]"); // Lowercase letter

    // Check the conditions
    var ctr = 0;
    for (var i = 0; i < matchedCase.length; i++) {
      if (new RegExp(matchedCase[i]).test(password)) {
        ctr++;
      }
    }

    console.log("ctr", ctr);
    return ctr >= strength;
  }
}

export default Validator;
