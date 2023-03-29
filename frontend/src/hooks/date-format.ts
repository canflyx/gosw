import dayjs from 'dayjs'
import utc from 'dayjs/plugin/utc'

dayjs.extend(utc)
const DATE_TIME_FORMAT = 'YYYY-MM-DD HH:mm:ss'
// 利用 dayjs 库格式化 utc 字符串
export function formatUtcString(
  utcString: string,
  format: string = DATE_TIME_FORMAT
) {
  if (utcString === '0001-01-01T00:00:00Z') {
    return 'null'
  } else {
    return dayjs.utc(utcString).utcOffset(8).format(format)
  }
}
