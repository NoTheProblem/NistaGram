import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {ToastrService} from 'ngx-toastr';
import {Observable} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class FollowService {

  constructor(private http: HttpClient, private toastr: ToastrService) { }

  followUser(username: string): void{
    this.http.post('http://localhost:8080/api/followers/follow/' + username, null)
      .subscribe(
        res => {
          this.toastr.success('Request to follow sent!');
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
          this.toastr.success('User unfollowed');
        },
        (error => {
          console.log(error);
          this.toastr.error('Failed to unfollow');
        })
      );
  }

  blockUser(username: string): void{
    this.http.post('http://localhost:8080/api/followers/block/' + username, null)
      .subscribe(
        res => {
          this.toastr.success('User blocked');
        },
        (error => {
          console.log(error);
          this.toastr.error('Failed to block');
        })
      );
  }

  unBlockUser(username: string): void{
    this.http.put('http://localhost:8080/api/followers/unblock/' + username, null)
      .subscribe(
        res => {
          this.toastr.success('User unblocked');
        },
        (error => {
          console.log(error);
          this.toastr.error('Failed to unblock');
        })
      );
  }

  isFollowing(username: string): Observable<boolean> {
    return this.http.get<boolean>('http://localhost:8080/api/followers/isFollowing/' + username);
  }

  getRecommendedUsers(): Observable<string[]> {
    return this.http.get<string[]>('http://localhost:8080/api/followers/recommendedProfiles');
  }
}
