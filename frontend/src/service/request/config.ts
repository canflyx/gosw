let BASE_URL = ''
const TIME_OUT = 1500

if (process.env.NODE_ENV === 'development') {
  BASE_URL = 'http://127.0.0.1:8055/api/v1'
} else if (process.env.NODE_ENV === 'production') {
  BASE_URL = 'http://127.0.0.1:8055/api/v1'
} else {
  BASE_URL = 'http://127.0.0.1:8055/api/v1'
}

export { BASE_URL,TIME_OUT }
