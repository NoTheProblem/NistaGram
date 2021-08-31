import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import {BusinessRequestModel, STATUS_TYPE} from '../models/business-request.model';
import {ToastrService} from 'ngx-toastr';
import {UserModel} from '../models/user.model';

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

  constructor(private http: HttpClient, private toastr: ToastrService) { }

  login(userName, pw): Observable<any> {
    const loginCredentials = {
      username: userName,
      password: pw
    };
    return this.http.post('http://localhost:8080/api/auth/login', {
      username: loginCredentials.username,
      password: loginCredentials.password
    }, httpOptions);
  }

  register(user): Observable<any> {
    return this.http.post('http://localhost:8080/api/auth/register', {
      username: user.username,
      password: user.password,
    }, httpOptions);
  }

  registerBusiness(business): Observable<any> {
    return this.http.post('http://localhost:8080/api/auth/businessRegister', {
      username: business.username,
      email: business.email,
      web: business.web,
      password: business.password,
    });
  }


  getPendingBusinessRequests(): Observable<Array<BusinessRequestModel>> {
    return this.http.get<Array<BusinessRequestModel>>('http://localhost:8080/api/auth/getPendingBusinessRequests');
  }

  answerBusinessRequest(answer): void{
    this.http.post('http://localhost:8080/api/auth/answerBusinessRequest', {
      username: answer.username,
      status: answer.status,
    }).subscribe(
      res => {
        this.toastr.success('Success');
      },
      (error => {
        console.log(error);
        this.toastr.error('Failed to answer');
      })
    );
  }

  getAllUsers(): Observable<Array<UserModel>> {
    return this.http.get<Array<UserModel>>('http://localhost:8080/api/auth/getAllUsers');
  }

  chagneRole(answer): void {
    this.http.put('http://localhost:8080/api/auth/changeRole', {
      username: answer.username,
      role: answer.role,
    }).subscribe(
      res => {
        this.toastr.success('Success');
      },
      (error => {
        console.log(error);
        this.toastr.error('Failed to change role');
      })
    );
  }
}
