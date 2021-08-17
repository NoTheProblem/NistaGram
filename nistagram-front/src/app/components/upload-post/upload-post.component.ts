import { Component, OnInit } from '@angular/core';
import {PostService} from '../../services/post.service';

@Component({
  selector: 'app-upload-post',
  templateUrl: './upload-post.component.html',
  styleUrls: ['./upload-post.component.css']
})
export class UploadPostComponent implements OnInit {

  selectedFile: File = null;
  isUploaded = false;
  fileName = '';
  description: string;
  location: string;
  tags: string;

  constructor(private postService: PostService) { }

  ngOnInit(): void {
  }

  onFileSelected(event): void{
    this.selectedFile = event.target.files[0];
    this.isUploaded = true;
    this.fileName = this.selectedFile.name;
  }

  uploadPost(): void {
    const fd = new FormData();
    fd.append('myFile', this.selectedFile, this.selectedFile.name);
    fd.append('description', this.description);
    fd.append('location', this.location);
    fd.append('tags', this.tags);
    this.postService.uploadPost(fd);
    this.fileName = null;
    this.isUploaded = false;
    this.description = null;
    this.tags = null;
    this.location = null;
  }
}
