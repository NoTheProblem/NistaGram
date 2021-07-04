import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import {environment} from '../../environments/environment';


const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' })
};

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor(private http: HttpClient) { }

  login(usrn, pw): Observable<any> {
    this.http.get('http://localhost:8080/api/auth/');
    return this.http.post('http://localhost:8080/api/auth/login', {
      username: usrn,
      password: pw
    }, httpOptions);
  }


}
