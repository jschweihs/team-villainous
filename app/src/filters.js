import Vue from "vue";

// Dates an ISO date string and
// returns a beautified version of a date
Vue.filter("date", function(value) {
  console.log("filter date value", value);
  if (!value) {
    return;
  }
  const date = value.includes(" ") ? value.split(" ")[0] : value.split("T")[0];
  const dates = date.split("-");
  const year = dates[0];
  const months = [
    "January",
    "February",
    "March",
    "April",
    "May",
    "June",
    "July",
    "August",
    "September",
    "October",
    "November",
    "December"
  ];
  const month = months[dates[1] - 1];
  const day = parseInt(dates[2]);
  let suffix = "th";
  if (day == 1 || day == 21) {
    suffix = "st";
  } else if (day == 2 || day == 22) {
    suffix = "nd";
  } else if (day == 3 || day == 23) {
    suffix = "rd";
  }
  return month + " " + day + suffix + ", " + year;
});

// Dates an ISO date string and
// returns a beautified version of a date with time
Vue.filter("isoToDateTime", function(date) {
  let datetime = [];
  if (date.includes(" ")) {
    datetime = date.split(" ");
  } else if (date.includes("T")) {
    datetime = date.split("T");
  } else {
    return "N/A";
  }

  let dates = [];
  if (datetime[0].includes("-")) {
    dates = datetime[0].split("-");
  } else if (datetime[0].includes("/")) {
    dates = datetime[0].split("/");
  } else {
    return "N/A";
  }

  let times = [];
  if (datetime[1].includes(":")) {
    times = datetime[1].split(":");
  } else {
    return "N/A";
  }

  const months = [
    "Jan.",
    "Feb.",
    "Mar.",
    "Apr.",
    "May",
    "Jun.",
    "Jul.",
    "Aug.",
    "Sep.",
    "Oct.",
    "Nov.",
    "Dec."
  ];

  const month = months[parseInt(dates[1]) - 1];
  const day = dates[2];
  const year = dates[0];

  let hour = times[0];
  let ampm = "AM";

  if (hour == 0) {
    hour = 12;
  } else if (hour > 12) {
    // We good
  } else if (hour == 12) {
    ampm = "PM";
  } else if (hour < 24) {
    hour = hour - 12;
    ampm = "PM";
  } else if (hour == 24) {
    hour == 12;
    ampm = "PM";
  }

  const minute = times[1];

  // Seconds not needed for now

  return `${month} ${day} ${year} ${hour}:${minute} ${ampm}`;
});
