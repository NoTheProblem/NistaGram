import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { ToastrService } from 'ngx-toastr';

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
          this.toastr.error('Upload failed');
      })
      );
  }
}
