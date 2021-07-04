import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders} from '@angular/common/http';
import { Observable } from 'rxjs';
import {environment} from '../../environments/environment';

const API_URL = environment.baseUrl +  '/api/test/';
const URL = environment.baseUrl + '/users/';

const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' })
};

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(private http: HttpClient) { }

  getPublicContent(): Observable<any> {
    return this.http.get(API_URL + 'all', { responseType: 'text' });
  }

  registerEmployee(user): Observable<any> {
    return this.http.post(URL + 'signupemployee', {
      username: user.username,
      email: user.email,
      password: user.password,
      firstName : user.firstName,
      lastName : user.lastName,
      city : user.city,
      country : user.country,
      address : user.address,
      phoneNumber : user.phoneNumber,
      birthDate: user.birthDate,
      tip: user.role
    }, httpOptions);
  }

}
