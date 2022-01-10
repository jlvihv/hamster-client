const sd = require("silly-datetime");

/**
 * convert milliseconds to YYYY-MM-DD
 * @param time
 * @returns {string}
 */
export const timeToDay = time => {
    return sd.format(new Date(time), "YYYY-MM-DD");
};