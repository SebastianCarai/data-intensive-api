import { sleep } from 'k6';
import http from 'k6/http';

export const options = {
    vus: 10000,
    duration: '1m',
    thresholds: {
        http_req_failed: ['rate<0.01'],
        http_req_duration: ['p(95)<500','p(99)<1000']
    }
}

export default function(){
    http.get("http://localhost:8080/feed")
    sleep(.5)
}