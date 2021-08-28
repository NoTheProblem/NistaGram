import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {ToastrService} from 'ngx-toastr';
import {Observable} from 'rxjs';
import {RelationType} from '../models/relationshipType.model';
import {UsernameListModel} from '../models/username-list.model';

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

  getRelationship(username: string): Observable<RelationType> {
    return this.http.get<RelationType>('http://localhost:8080/api/followers/getRelationship/' + username);
  }

  getRecommendedUsers(): Observable<UsernameListModel> {
    return this.http.get<UsernameListModel>('http://localhost:8080/api/followers/recommendedProfiles');
  }

  getPendingRequests(): Observable<UsernameListModel> {
    return this.http.get<UsernameListModel>('http://localhost:8080/api/followers/getFollowerRequests');
  }

  acceptRequest(username: string): void{
    this.http.put('http://localhost:8080/api/followers/acceptRequest/' + username, null)
      .subscribe(
        res => {
          this.toastr.success('User accepted');
        },
        (error => {
          console.log(error);
          this.toastr.error('Failed to accept user');
        })
      );
  }
}

