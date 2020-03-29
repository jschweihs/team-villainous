class Validator {
  // Checks validity of email
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

    return ctr >= strength;
  }

  // Checks validity of date
  // Does not check for time
  // Assumes 2 digits for day and month
  // Assumes 4 digits for year
  // order expects a 3 digit string use Y, M, and D
  static validDate(date, separator = "/", order = "MDY") {
    if (!date) {
      return false;
    }
    // Check length
    if (date.length != 10) {
      return false;
    }

    // Get components of date
    const dates = date.split(separator);

    // Get expected positions of date elements
    const monthPos = order.indexOf("M");
    const dayPos = order.indexOf("D");
    const yearPos = order.indexOf("Y");

    // Get date values
    const month = dates[monthPos];
    const day = dates[dayPos];
    const year = dates[yearPos];

    // Track validity
    let valid = true;

    // Check for valid year
    // A valid year is a number somewhere between
    // 1900 and the current year
    valid =
      valid &&
      year == parseInt(year) &&
      year >= 1900 &&
      year <= new Date().getFullYear();

    // Check for valid month
    // A valid month is a number somewhere between
    // 1 and 12
    valid = valid && month == parseInt(month) && month >= 1 && month <= 12;

    // Check for valid day
    // A valid day is a number somewhere between
    // 1 and 12
    valid = valid && day == parseInt(day) && day >= 1 && day <= 31;

    return valid;
  }

  // Checks validity of time
  // Assumes 2 digits for hours and minutes
  // HH:MM:SS
  // Can optionally provide seconds
  static validTime(time) {
    if (!time) {
      return false;
    }
    // Get parts of the time
    const times = time.split(":");

    const hour = times[0];
    const minute = times[1];
    const second = times[2];

    // Track validity
    let valid = true;

    // Check for valid hour
    // A valid hour has 2 digits and is a number somewhere between
    // 0 and 23
    valid =
      valid &&
      hour.length == 2 &&
      hour == parseInt(hour) &&
      hour >= 0 &&
      hour <= 23;

    // Check for valid minute
    // A valid minute has 2 digits and is a number somewhere between
    // 0 and 59
    valid =
      valid &&
      minute.length == 2 &&
      minute == parseInt(minute) &&
      minute >= 0 &&
      minute <= 59;

    if (second) {
      // Check for valid second
      // A valid second has 2 digits and is a number somewhere between
      // 0 and 59
      valid =
        valid &&
        second.length == 2 &&
        second == parseInt(second) &&
        second >= 0 &&
        second <= 59;
    }

    return valid;
  }

  // Checks validity of datetime
  // See validDate and validTime for specific requirements
  static validDatetime(
    datetime,
    dateSeparator = "/",
    datetimeSeparator = "T",
    dateOrder = "MDY"
  ) {
    if (!datetime) {
      return false;
    }

    const dateParts = datetime.split(datetimeSeparator);
    if (!dateParts || dateParts.length > 2) {
      return false;
    }

    // Get date
    const date = dateParts[0];

    // Get time
    const leftover = dateParts[1];
    let time;
    if (leftover.includes("Z")) {
      time = leftover.split("Z")[0];
    } else if (leftover.includes("+")) {
      time = leftover.split("-")[0];
    } else if (leftover.includes("-")) {
      time = leftover.split("-")[0];
    } else {
      time = leftover;
    }

    // Return valid date and valid time
    return (
      this.validDate(date, dateSeparator, dateOrder) && this.validTime(time)
    );
  }
}

export default Validator;
