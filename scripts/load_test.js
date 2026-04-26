import http from 'k6/http';
import { check } from 'k6';

export const options = {
  scenarios: {
    stress: {
      executor: 'ramping-vus',
      startVUs: 0,
      stages: [
        { duration: '10s', target: 500 },
        { duration: '30s', target: 500 },
        { duration: '10s', target: 0 },
      ],
    },
  },
};

export default function () {
  const url = 'http://localhost:8080/api/v1/transfer';
  const uniqueId = Math.random().toString(36).substring(2);
  
  const payload = JSON.stringify({
    from_account: 'acc_987654321',
    to_account: 'acc_123456789',
    amount: 25.50
  });

  const params = {
    headers: {
      'Content-Type': 'application/json',
      'Idempotency-Key': uniqueId,
    },
  };

  const response = http.post(url, payload, params);

  check(response, {
    'success or cached': (r) => r.status === 201 || r.status === 200,
    'no server crashes': (r) => r.status !== 500,
  });
}
