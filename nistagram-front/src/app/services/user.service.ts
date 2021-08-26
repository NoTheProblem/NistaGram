import { Injectable } from '@angular/core';
import {Observable} from 'rxjs';
import {UserModel} from '../models/user.model';
import {HttpClient } from '@angular/common/http';
import {ToastrService} from 'ngx-toastr';
import {PostModel} from '../models/post.model';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(private http: HttpClient, private toastr: ToastrService) {}

  public loadProfile(username: string): Observable<UserModel> {
    return this.http.get<UserModel>('http://localhost:8080/api/user/username/' + username) ;
  }

  public loadProfilePosts(username: string): Observable<Array<PostModel>> {
    return this.http.get<Array<PostModel>>('http://localhost:8080/api/post/username/' + username);
  }


}
