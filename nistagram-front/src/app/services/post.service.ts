import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { ToastrService } from 'ngx-toastr';
import {Observable} from 'rxjs';
import {PostModel} from '../models/post.model';
import {ReportModel} from '../models/report.model';

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

  likePost(form: any): void{
    this.http.put('http://localhost:8080/api/post/likePost', form).subscribe(
      res => {
        this.toastr.success('Post liked!');
      },
      (error => {
        console.log(error);
        this.toastr.error('Post not liked!');
      })
    );
  }

  disLikePost(form: any): void{
    this.http.put('http://localhost:8080/api/post/disLikePost', form).subscribe(
      res => {
      },
      (error => {
        console.log(error);
      })
    );
  }

  commentPost(form: any): void{
    this.http.put('http://localhost:8080/api/post/commentPost', form).subscribe(
      res => {
        this.toastr.success('Comment posted!');
      },
      (error => {
        console.log(error);
        this.toastr.error('Failed to post comment!');
      })
    );
  }

  reportPost(postId: string): void {
    this.http.post('http://localhost:8080/api/post/reportPost', {
      id: postId
    }).subscribe(
        res => {
          this.toastr.success('Post reported!');
        },
        (error => {
          console.log(error);
          this.toastr.error('Failed to report post!');
        })
      );
  }

  public getUnAnsweredReports(): Observable<Array<ReportModel>> {
    return this.http.get<Array<ReportModel>>('http://localhost:8080/api/post/getUnAnsweredReports') ;
  }

  answerReport(body: any): void {
    this.http.put('http://localhost:8080/api/post/answerReport', body).subscribe(
      res => {
        this.toastr.success('Report answered!');
      },
      (error => {
        console.log(error);
        this.toastr.error('Failed to answer to report!');
      })
    );

  }
}
