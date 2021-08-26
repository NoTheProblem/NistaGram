import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {ToastrService} from 'ngx-toastr';
import {Observable} from 'rxjs';
import { of } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class FollowService {

  constructor(private http: HttpClient, private toastr: ToastrService) { }

  followUser(username: string): void{
    this.http.post('http://localhost:8080/api/followers/follow/' + username, null)
      .subscribe(
        res => {
          this.toastr.success('User followed');
        },
        (error => {
          console.log(error);
          this.toastr.error('Failed to follow');
        })
      );
  }

  unFollowUser(username: string): void{
    this.http.put('http://localhost:8080/api/followers/unfollow/' + username, null)
      .subscribe(
        res => {
          this.toastr.success('User followed');
        },
        (error => {
          console.log(error);
          this.toastr.error('Failed to follow');
        })
      );
  }

  blockUser(username: string): void{
    this.http.post('http://localhost:8080/api/followers/block/' + username, null)
      .subscribe(
        res => {
          this.toastr.success('User followed');
        },
        (error => {
          console.log(error);
          this.toastr.error('Failed to follow');
        })
      );
  }

  unBlockUser(username: string): void{
    this.http.post('http://localhost:8080/api/followers/unblock/' + username, null)
      .subscribe(
        res => {
          this.toastr.success('User followed');
        },
        (error => {
          console.log(error);
          this.toastr.error('Failed to follow');
        })
      );
  }

  isFollowing(username: string): Observable<boolean> {
    return of(true);
    // TODO uncomment once its done on bakc
    // return this.http.get<boolean>('http://localhost:8080/api/followers/isFollowing/' + username);
  }

}
