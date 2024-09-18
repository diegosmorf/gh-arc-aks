import http from 'k6/http';

export const options = {
  vus: 50,
  duration: '5s',
  tags: {
    testid: `${getCurrentUTCTimestamp()}-runners`,
  },
};

export function getCurrentUTCTimestamp() {
  const now = new Date();
  return now.toISOString().replace("T", " ").substring(0, 19);
}

export default function () {
  const url = 'http://test.k6.io/login';
  const payload = JSON.stringify({
    email: 'aaa',
    password: 'bbb',
  });

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  http.post(url, payload, params);
}