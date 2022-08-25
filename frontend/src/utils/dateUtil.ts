/**
 * Independent time operation tool to facilitate subsequent switch to dayjs
 */
import dayjs from 'dayjs';
import { format } from 'date-fns';

type ParsableDateType = Parameters<typeof dayjs>[0];

// Format Rule: https://date-fns.org/v2.28.0/docs/format
const FORMATTERS = {
  normal: 'yyyy-MM-dd',
  datetime: 'yyyy-MM-dd HH:mm:ss',
  datetimeWithoutSec: 'yyyy-MM-dd HH:mm',
  monthOnly: 'MM-dd',
};

const dayjsParseDate = (date: ParsableDateType) => {
  if (!date) return;
  if (date instanceof Date) return date;

  return dayjs(date).toDate();
};

export function formatDate(
  date: ParsableDateType = undefined,
  formatter: string | ((f: typeof FORMATTERS) => string) = FORMATTERS.normal,
): string {
  if (typeof formatter === 'function') formatter = formatter(FORMATTERS);
  const parsedDate = dayjsParseDate(date);
  return parsedDate ? format(parsedDate, formatter) : '';
}

export function formatToDateTime(
  date: ParsableDateType = undefined,
  formatter: string | ((f: typeof FORMATTERS) => string) = FORMATTERS.datetime,
) {
  return formatDate(date, formatter);
}

export const dateUtil = dayjs;

export function formatSeconds(value) {
  let theTime = Number(value);
  let theTime1 = 0;
  let theTime2 = 0;
  let theTime3 = 0;
  if (theTime >= 60) {
    theTime1 = parseInt(theTime / 60);
    theTime = parseInt(theTime % 60);
    if (theTime1 >= 60) {
      theTime2 = parseInt(theTime1 / 60);
      theTime1 = parseInt(theTime1 % 60);
      if (theTime2 >= 24) {
        //大于24小时
        theTime3 = parseInt(theTime2 / 24);
        theTime2 = parseInt(theTime2 % 24);
      }
    }
  }
  let result = '';
  if (theTime1 >= 0) {
    result = '' + parseInt(theTime1) + 'M' + result;
  }
  if (theTime2 > 0) {
    result = '' + parseInt(theTime2) + 'H ' + result;
  }
  if (theTime3 > 0) {
    if (theTime2 == 0) {
      result = '' + parseInt(theTime2) + 'H ' + result;
    }
    result = '' + parseInt(theTime3) + 'D ' + result;
  }
  return result;
}
