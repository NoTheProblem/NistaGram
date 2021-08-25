import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {ToastrService} from 'ngx-toastr';
import {Observable} from 'rxjs';
import {VerificationRequestModel} from '../models/verification-request.model';

@Injectable({
  providedIn: 'root'
})
export class RequestService {

  constructor(private http: HttpClient, private toastr: ToastrService) { }


  sendVerificationAnswer(body: any): void{
    this.http.put('http://localhost:8080/api/verification/answer', body)
      .subscribe(
        res => {
          this.toastr.success('Answer sent!');
        },
        (error => {
          this.toastr.error('Failed to send answer!');
        })
      );
  }

  public getUnAnsweredVerificationRequests(): Observable<Array<VerificationRequestModel>> {
    return this.http.get<Array<VerificationRequestModel>>('http://localhost:8080/api/verification/getUnAnswered') ;
  }


}
