import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import {environment} from '../../environments/environment';

const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type':  'application/json',
    'Access-Control-Allow-Headers': 'Content-Type',
    'Access-Control-Allow-Methods': 'GET',
    'Access-Control-Allow-Origin': '*'
  })
};

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor(private http: HttpClient) { }

  login(usrn, pw): Observable<any> {
    const loginCredentials = {
      username: usrn,
      password: pw
    };
    return this.http.post('http://localhost:8080/api/auth/login', {
      username: loginCredentials.username,
      password: loginCredentials.password
    }, httpOptions);
  }

  register(user): Observable<any> {
    return this.http.post('http://localhost:8080/api/auth/register', {
      email: user.email,
      username: user.username,
      password: user.password,
    }, httpOptions);
  }
}
