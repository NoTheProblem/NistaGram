import { Injectable } from '@angular/core';
import {Observable} from 'rxjs';
import {PostModel} from '../models/post.model';
import {HttpClient} from '@angular/common/http';
import {UserModel} from '../models/user.model';

@Injectable({
  providedIn: 'root'
})
export class SearchService {

  constructor(private http: HttpClient) { }

  public searchTag(tag: string): Observable<Array<PostModel>> {
      return this.http.get<Array<PostModel>>('http://localhost:8080/api/post/searchTag/' + tag);
  }
  public searchLocation(location: string): Observable<Array<PostModel>> {
    return this.http.get<Array<PostModel>>('http://localhost:8080/api/post/searchLocation/' + location);
  }
  public searchUser(username: string): Observable<Array<UserModel>> {
    return this.http.get<Array<UserModel>>('http://localhost:8080/api/user/search/' + username);
  }

}
