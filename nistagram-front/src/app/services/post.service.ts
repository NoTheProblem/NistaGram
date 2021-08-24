import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { ToastrService } from 'ngx-toastr';
import {Observable} from 'rxjs';
import {PostModel} from '../models/post.model';

@Injectable({
  providedIn: 'root'
})
export class PostService {

  constructor(private http: HttpClient, private toastr: ToastrService) { }

  uploadPost(fd: FormData): void{
    this.http.post('http://localhost:8080/api/post/uploadPost', fd)
      .subscribe(
        res => {
          this.toastr.success('Post uploaded');
        },
        (error => {
          console.log(error);
          this.toastr.error('Upload failed');
      })
      );
  }

  public HomeFeed(): Observable<Array<PostModel>> {
    return this.http.get<Array<PostModel>>('http://localhost:8080/api/post/homeFeed');
  }
}
