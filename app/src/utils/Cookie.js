class Cookie {
    static setCookie(name, value, expires="TOMORROW", domain=".teamvillainous.com", path="/") {
        // Use string as keyword so we can calculate tomorrow inside function
        if(expires == 'TOMORROW') {
            let d = new Date();
            d.setTime(d.getTime() + (24 * 60 * 60 * 1000));
            expires = d.toUTCString();
        }

        document.cookie = `${name}=${value};expires=${expires};domain=${domain};path=${path}`;
    }

    static getCookie(cname) {
        const name = cname + "=";
        const decodedCookie = decodeURIComponent(document.cookie);
        const ca = decodedCookie.split(';');
        for(var i=0; i<ca.length; i++) {
            let c = ca[i];
            while(c.charAt(0) == ' ') {
                c = c.substring(1);
            }
            if(c.indexOf(name) == 0) {
                return c.substring(name.length, c.length);
            }
        }
        return '';
    }

    static deleteCookie(name) {
        this.setCookie(name, '','Thu, 01 Jan 1970 00:00:00 UTC')
    }
}

export default Cookie;