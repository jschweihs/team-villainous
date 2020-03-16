import Vue from 'vue';

Vue.filter('date', function (value) {
    const date = value.includes(' ') ? value.split(' ')[0] : value.split('T')[0];
    const date_array = date.split('-');
    const year = date_array[0];
    const months = [
        'January', 
        'February', 
        'March', 
        'April', 
        'May', 
        'June', 
        'July', 
        'August', 
        'September', 
        'October', 
        'November', 
        'December'
    ];
    const month = months[date_array[1]-1];
    const day = parseInt(date_array[2]);
    let suffix = 'th';
    if (day == 1 || day == 21)		{ suffix = 'st' }
    else if (day == 2 || day == 22) { suffix = 'nd' }
    else if (day == 3 || day == 23) { suffix = 'rd' }
    return month + ' ' + day + suffix + ', ' + year;
});

Vue.filter('isoToDateTime', function (date) {

    let datetime = [];
    if( date.includes(" ") ) {
        datetime = date.split(" ");
    } else if ( date.includes("T") ) {
        datetime = date.split("T");
    } else {
        console.log("Trying to filter non date string");
        return 'N/A';
    }

    let dates = [];
    if( datetime[0].includes("-") ) {
        dates = datetime[0].split("-");
    } else if ( datetime[0].includes("/") ) {
        dates = datetime[0].split("/");
    } else {
        console.log("Trying to filter non date string");
        return 'N/A';
    }

    let times = [];
    if( datetime[1].includes(":") ) {
        times = datetime[1].split(":");
    }  else {
        console.log("Trying to filter non date string");
        return 'N/A';
    }

    const months = [
        'Jan.',
        'Feb.',
        'Mar.',
        'Apr.',
        'May',
        'Jun.',
        'Jul.',
        'Aug.',
        'Sep.',
        'Oct.',
        'Nov.',
        'Dec.'
    ]
    console.log(dates[1]);
    console.log(months[dates[1]])
    const month = months[ parseInt(dates[1]) - 1 ];

    const day = dates[2];

    const  year = dates[0];

    let hour = times[0];
    let ampm = 'AM';

    if(hour == 0) {
        hour = 12;
    } else if(hour > 12) {
        // We good
    } else if (hour == 12) {
        ampm = 'PM';
    } else if (hour < 24) {
        hour = hour - 12;
        ampm = 'PM';
    } else if (hour == 24) {
        hour == 12;
        ampm = 'PM';
    }

    const minute = times[1];
    // Seconds not needed...for now

    return `${month} ${day} ${year} ${hour}:${minute} ${ampm}`;

});