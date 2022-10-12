import http from 'k6/http';
import { check } from 'k6';

export const options = {
    vus: 3,
    duration: '5s',
};

export default function () {
    var res = http.get('http://localhost:30081/image/jpeg');
    check(res, { 'status was 200': (r) => r.status == 200 });
}
